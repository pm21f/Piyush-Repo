import os
import struct
from typing import List, Tuple

# ================================================================
# AES Implementation in Python
# ================================================================

class AES:
    """
    Advanced Encryption Standard (AES) Implementation
    with 128-bit, 192-bit, and 256-bit key support.
    """
    Nb = 4  # Block size in 32-bit words
    Nk = None  # Number of 32-bit words in the key
    Nr = None  # Number of rounds

    SBox = [
        # S-Box table (hex values for substitution)
        0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
        0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
        0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
        # ... (remaining entries for brevity)
    ]

    Rcon = [
        # Round constant for key schedule
        0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36, 0x6c, 0xd8, 0xab, 0x4d, 0x9a, 0x2f
    ]

    def _init_(self, key: bytes):
        """
        Initializes the AES object with the provided key.
        """
        self.key = key
        self.Nk = len(key) // 4
        self.Nr = self.Nk + 6
        self.w = self.key_expansion()

    def key_expansion(self) -> List[List[int]]:
        """
        Expands the cipher key into a key schedule.
        """
        # Initialize the key schedule with the original key
        key_schedule = [[0] * self.Nb for _ in range(self.Nb * (self.Nr + 1))]
        for i in range(self.Nk):
            for j in range(4):
                key_schedule[i][j] = self.key[i * 4 + j]

        # Generate the rest of the key schedule
        for i in range(self.Nk, self.Nb * (self.Nr + 1)):
            temp = key_schedule[i - 1][:]
            if i % self.Nk == 0:
                temp = self.sub_word(self.rot_word(temp))
                temp[0] ^= self.Rcon[i // self.Nk - 1]
            elif self.Nk > 6 and i % self.Nk == 4:
                temp = self.sub_word(temp)
            for j in range(4):
                key_schedule[i][j] ^= key_schedule[i - self.Nk][j]
        return key_schedule

    def sub_word(self, word: List[int]) -> List[int]:
        """
        Applies the SBox to each byte in the word.
        """
        return [self.SBox[b] for b in word]

    def rot_word(self, word: List[int]) -> List[int]:
        """
        Rotates a word (left shift by 1 byte).
        """
        return word[1:] + word[:1]

    def encrypt(self, plaintext: bytes) -> bytes:
        """
        Encrypts a single block (16 bytes) of plaintext.
        """
        state = [list(plaintext[i:i + self.Nb]) for i in range(0, len(plaintext), self.Nb)]
        self.add_round_key(state, self.w[:self.Nb])

        for round in range(1, self.Nr):
            self.sub_bytes(state)
            self.shift_rows(state)
            self.mix_columns(state)
            self.add_round_key(state, self.w[round * self.Nb:(round + 1) * self.Nb])

        self.sub_bytes(state)
        self.shift_rows(state)
        self.add_round_key(state, self.w[self.Nr * self.Nb:])

        return b''.join(bytes(row) for row in state)

    def decrypt(self, ciphertext: bytes) -> bytes:
        """
        Decrypts a single block (16 bytes) of ciphertext.
        """
        state = [list(ciphertext[i:i + self.Nb]) for i in range(0, len(ciphertext), self.Nb)]
        self.add_round_key(state, self.w[self.Nr * self.Nb:])

        for round in range(self.Nr - 1, 0, -1):
            self.inv_shift_rows(state)
            self.inv_sub_bytes(state)
            self.add_round_key(state, self.w[round * self.Nb:(round + 1) * self.Nb])
            self.inv_mix_columns(state)

        self.inv_shift_rows(state)
        self.inv_sub_bytes(state)
        self.add_round_key(state, self.w[:self.Nb])

        return b''.join(bytes(row) for row in state)

    # Add round key logic
    def add_round_key(self, state: List[List[int]], round_key: List[List[int]]):
        """
        Adds (XORs) the round key to the state.
        """
        for i in range(self.Nb):
            for j in range(4):
                state[i][j] ^= round_key[i][j]

