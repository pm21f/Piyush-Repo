import random
import time

class DataVerification:
    def __init__(self):
        self.received_data = set()  # Set to hold received data for comparison
        self.retry_limit = 3        # Maximum number of retries
        self.transmission_attempts = 0
        self.successful_transmissions = 0
        self.failed_transmissions = 0
        self.log = []               # To store log of actions

    def send_data(self, data, client_id):
        """
        Simulate sending data and storing the received data.
        :param data: Data being sent.
        :param client_id: Unique identifier for the client.
        :return: None
        """
        self.transmission_attempts += 1
        print(f"Client {client_id} Sending Data: {data}")
        self.log.append(f"Client {client_id} Sending Data: {data}")
        
        if random.random() < 0.1:  # 10% chance to simulate an error
            noisy_data = self.simulate_error(data)
            print(f"Error in transmission. Data after error: {noisy_data}")
            self.log.append(f"Error in transmission. Data after error: {noisy_data}")
            self.received_data.add(noisy_data)
        else:
            self.received_data.add(data)  # Successfully received the data

    def simulate_error(self, data):
        """
        Randomly introduce errors into the data.
        :param data: Original data.
        :return: Data with simulated error.
        """
        noisy_data = list(data)
        error_position = random.randint(0, len(data) - 1)
        noisy_data[error_position] = '1' if data[error_position] == '0' else '0'  # Flip one bit
        return ''.join(noisy_data)

    def verify_data(self, retry_count=0):
        """
        Verify if the data sent is correct by checking if all three transmissions match.
        :param retry_count: Current retry attempt number.
        :return: True if data is consistent, False otherwise.
        """
        if len(self.received_data) == 1:
            print(f"Data is consistent across all transmissions after {retry_count} retries.")
            self.successful_transmissions += 1
            return True
        else:
            self.failed_transmissions += 1
            print(f"Data inconsistency detected after {retry_count} retries.")
            if retry_count < self.retry_limit:
                print(f"Retrying data verification... Attempt {retry_count + 1}")
                self.received_data.clear()
                return False
            else:
                print("Max retries reached. Data is inconsistent.")
                return False

    def process_data(self, data, client_id):
        """
        Processes the data for transmission, simulates errors, and verifies consistency.
        :param data: Data to be sent.
        :param client_id: Identifier of the client sending the data.
        :return: None
        """
        retry_count = 0
        while retry_count <= self.retry_limit:
            self.send_data(data, client_id)
            if self.verify_data(retry_count):
                break
            retry_count += 1
            time.sleep(2)  # Delay before retrying

    def generate_report(self):
        """
        Generate a report of the transmission attempts, successes, and failures.
        :return: None
        """
        print("\n--- Transmission Report ---")
        print(f"Total Transmission Attempts: {self.transmission_attempts}")
        print(f"Successful Transmissions: {self.successful_transmissions}")
        print(f"Failed Transmissions: {self.failed_transmissions}")
        print(f"Log of Actions: {self.log}")
        print("----------------------------")


# Simulate multiple clients
def simulate_multiple_clients():
    data = "101010101010"  # Example data to send
    client_ids = [1, 2, 3, 4, 5]
    
    verifier = DataVerification()
    
    for client_id in client_ids:
        verifier.process_data(data, client_id)
    
    verifier.generate_report()


if __name__ == "__main__":
    simulate_multiple_clients()
