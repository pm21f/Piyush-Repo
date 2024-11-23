import random
import logging

# Configure logging for error detection and correction
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(message)s')

class ICER:
    def __init__(self):
        self.error_count = 0
        self.corrected_count = 0
        self.transmission_count = 0

    # Basic checksum calculation using parity
    def checksum(self, data):
        return sum(data) % 2

    # Generate Hamming Code for error detection and correction
    def generate_hamming_code(self, data):
        """
        Generates a Hamming code for the given data bits.
        :param data: List of binary data bits.
        :return: Hamming code with parity bits.
        """
        n = len(data)
        r = 1
        while (2 ** r) < (n + r + 1):
            r += 1

        hamming_code = [0] * (n + r)
        j = 0
        for i in range(1, len(hamming_code) + 1):
            if i & (i - 1) == 0:  # If i is a power of 2, leave space for parity
                continue
            hamming_code[i - 1] = data[j]
            j += 1

        # Calculate parity bits
        for i in range(r):
            pos = 2 ** i
            count = 0
            for j in range(1, len(hamming_code) + 1):
                if j & pos and j != pos:
                    count += hamming_code[j - 1]
            hamming_code[pos - 1] = count % 2

        return hamming_code

    # Detect and correct errors using Hamming Code
    def detect_and_correct_hamming(self, code):
        """
        Detect and correct single-bit errors in Hamming code.
        :param code: List of binary Hamming code.
        :return: Corrected code and error position.
        """
        n = len(code)
        r = 0
        while (2 ** r) < n:
            r += 1

        error_position = 0
        for i in range(r):
            pos = 2 ** i
            count = 0
            for j in range(1, n + 1):
                if j & pos:
                    count += code[j - 1]
            if count % 2 != 0:
                error_position += pos

        if error_position:
            self.error_count += 1
            code[error_position - 1] ^= 1  # Flip the erroneous bit
            self.corrected_count += 1
            logging.warning(f"Error detected and corrected at position {error_position}")
        else:
            logging.info("No error detected in Hamming code.")
        return code, error_position

    # Simulate data transmission with random noise
    def simulate_transmission(self, data, error_rate=0.1):
        """
        Simulates data transmission with a given error rate.
        :param data: Original data bits.
        :param error_rate: Probability of bit error.
        :return: Transmitted data with possible errors.
        """
        transmitted = data[:]
        for i in range(len(transmitted)):
            if random.random() < error_rate:  # Introduce error with given probability
                transmitted[i] ^= 1  # Flip the bit
        self.transmission_count += 1
        return transmitted

    # Process data for error detection and correction
    def process(self, data):
        """
        Process the data: encode, simulate transmission, and decode.
        :param data: Original data bits.
        :return: Processed and corrected data.
        """
        # Encode data with Hamming code
        hamming_code = self.generate_hamming_code(data)
        logging.info(f"Original Hamming Code: {hamming_code}")

        # Simulate transmission with errors
        received_code = self.simulate_transmission(hamming_code)
        logging.info(f"Received Hamming Code: {received_code}")

        # Detect and correct errors
        corrected_code, _ = self.detect_and_correct_hamming(received_code)
        logging.info(f"Corrected Hamming Code: {corrected_code}")

        return corrected_code

if __name__ == "__main__":
    icer = ICER()

    # Example data to encode
    original_data = [1, 0, 1, 1]
    logging.info(f"Original Data: {original_data}")

    # Process data
    processed_data = icer.process(original_data)

    # Summary
    logging.info(f"Processed Data: {processed_data}")
    logging.info(f"Total Errors Detected: {icer.error_count}")
    logging.info(f"Total Errors Corrected: {icer.corrected_count}")
    logging.info(f"Total Transmissions: {icer.transmission_count}")
