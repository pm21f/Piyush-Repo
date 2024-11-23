import numpy as np
import random

class TurboCodes:
    def __init__(self, constraint_length=3, rate=1/2, iterations=5):
        """
        Initialize TurboCodes with given parameters.
        :param constraint_length: Constraint length for convolutional encoding.
        :param rate: Rate of the code.
        :param iterations: Number of iterations for decoding (simplified).
        """
        self.constraint_length = constraint_length
        self.rate = rate
        self.iterations = iterations

    def convolutional_encode(self, data):
        """
        Simplified convolutional encoder.
        :param data: List of binary data.
        :return: Encoded binary data.
        """
        encoded = []
        n = len(data)
        for i in range(n):
            if i == 0:
                encoded.append(data[i])
                encoded.append(data[i])
            else:
                encoded.append(data[i] ^ data[i-1])  # XOR with the previous bit
                encoded.append(data[i])
        return encoded

    def interleave(self, data, method="block"):
        """
        Apply interleaving to data.
        :param data: List of binary data.
        :param method: Interleaving method ('block' or 'random').
        :return: Interleaved data.
        """
        n = len(data)
        interleaved = data[:]
        if method == "block":
            interleaved = data[::-1]  # Reverse the data
        elif method == "random":
            random.seed(42)  # Fixed seed for reproducibility
            interleaved = random.sample(data, len(data))
        return interleaved

    def encode(self, data):
        """
        Perform Turbo encoding by combining two convolutional encoders with interleaving.
        :param data: List of binary data.
        :return: Turbo encoded data.
        """
        first_encoded = self.convolutional_encode(data)
        interleaved_data = self.interleave(data, method="block")
        second_encoded = self.convolutional_encode(interleaved_data)
        return first_encoded + second_encoded

    def simulate_error(self, encoded, error_rate=0.1):
        """
        Simulate errors by flipping bits randomly based on an error rate.
        :param encoded: Encoded binary data.
        :param error_rate: Probability of flipping each bit.
        :return: Noisy encoded data.
        """
        noisy_encoded = encoded[:]
        for i in range(len(noisy_encoded)):
            if np.random.rand() < error_rate:
                noisy_encoded[i] ^= 1  # Flip the bit
        return noisy_encoded

    def decode(self, encoded):
        """
        Simplified decoding by returning the first half of the encoded data.
        :param encoded: Encoded binary data.
        :return: Decoded binary data.
        """
        n = len(encoded) // 2
        return encoded[:n]  # Return the first half for simplicity

    def calculate_bit_error_rate(self, original, decoded):
        """
        Calculate Bit Error Rate (BER) between original and decoded data.
        :param original: Original binary data.
        :param decoded: Decoded binary data.
        :return: BER as a float.
        """
        errors = sum(o != d for o, d in zip(original, decoded))
        return errors / len(original)

if __name__ == "__main__":
    # Example binary data
    data = [1, 0, 1, 1, 0, 1, 0, 0]
    
    turbo = TurboCodes()
    
    # Encoding
    encoded = turbo.encode(data)
    print("Encoded Data:", encoded)
    
    # Simulating errors
    noisy_encoded = turbo.simulate_error(encoded, error_rate=0.2)
    print("Noisy Encoded Data:", noisy_encoded)
    
    # Decoding
    decoded = turbo.decode(noisy_encoded)
    print("Decoded Data:", decoded)
    
    # BER Calculation
    ber = turbo.calculate_bit_error_rate(data, decoded)
    print("Bit Error Rate (BER):", ber)
