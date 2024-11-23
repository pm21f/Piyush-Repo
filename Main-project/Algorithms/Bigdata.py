from concurrent.futures import ThreadPoolExecutor
import pandas as pd
import random
import time


# Reuse the ECC class from previous code

# Batch Encrypt Function
def batch_encrypt(ecc, data_chunk, public_key):
    return data_chunk.apply(lambda x: ecc.encrypt(x, public_key))

# Batch Decrypt Function
def batch_decrypt(ecc, data_chunk, private_key):
    return data_chunk.apply(lambda x: ecc.decrypt(x, private_key))


# Main Function with Parallel Processing
if __name__ == "__main__":
    # ECC Parameters
    a, b, p = 0, 7, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F
    G = (55066263022277343669578718895168534326250603453777594175500187360389116729240,
         32670510020758816978083085130507043184471273380659243275938904335757337470005)
    n = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141

    ecc = ECC(a, b, p, G, n)

    # Generate ECC Keys
    private_key, public_key = ecc.generate_key_pair()

    # Create a Sample Big Dataset
    data = {
        "RecordID": range(1, 10001),  # Simulate 10,000 records
        "SensitiveData": [random.randint(1000, 9999) for _ in range(10000)]
    }
    df = pd.DataFrame(data)
    print("Dataset Created")

    # Divide the dataset into chunks for parallel processing
    num_threads = 4
    chunk_size = len(df) // num_threads
    data_chunks = [df["SensitiveData"][i:i + chunk_size] for i in range(0, len(df), chunk_size)]

    # Encrypt Dataset in Parallel
    start_time = time.time()
    with ThreadPoolExecutor(max_workers=num_threads) as executor:
        encrypted_chunks = list(executor.map(lambda chunk: batch_encrypt(ecc, chunk, public_key), data_chunks))
    end_time = time.time()

    # Combine Encrypted Chunks
    df["Ciphertext"] = pd.concat(encrypted_chunks)
    print(f"Encryption completed in {end_time - start_time:.2f} seconds.")

    # Decrypt Dataset in Parallel
    start_time = time.time()
    with ThreadPoolExecutor(max_workers=num_threads) as executor:
        decrypted_chunks = list(executor.map(lambda chunk: batch_decrypt(ecc, chunk, private_key), df["Ciphertext"]))
    end_time = time.time()

    # Combine Decrypted Chunks
    df["DecryptedData"] = pd.concat(decrypted_chunks)
    print(f"Decryption completed in {end_time - start_time:.2f} seconds.")

    # Check Accuracy
    df["IsDecryptionCorrect"] = df["SensitiveData"] == df["DecryptedData"]
    print(df["IsDecryptionCorrect"].value_counts())

    # Log Results
    df.to_csv("encrypted_dataset.csv", index=False)
    print("Encrypted dataset saved.")
