package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulate Cyclic Redundancy Check (CRC)
func cyclicRedundancyCheck(data string) bool {
	fmt.Println("Performing Cyclic Redundancy Check (CRC)...")
	return rand.Intn(2) == 0 // Simulates random success/failure
}

// Simulate Low-Density Parity Check (LDPC)
func lowDensityParityCheck(data string) bool {
	fmt.Println("Performing Low-Density Parity Check (LDPC)...")
	return rand.Intn(2) == 0 // Simulates random success/failure
}

// Simulate Forward Error Correction (FEC)
func forwardErrorCorrection(data string) bool {
	fmt.Println("Applying Forward Error Correction (FEC)...")
	return rand.Intn(2) == 0 // Simulates random success/failure
}

// Reattempt logic with retries
func reattemptLogic(processFunc func(string) bool, data string, maxRetries int, delay time.Duration) bool {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Attempt %d/%d...\n", attempt, maxRetries)
		if processFunc(data) {
			fmt.Printf("✅ Process succeeded on attempt %d.\n", attempt)
			return true
		}
		fmt.Printf("⚠️ Process failed. Retrying in %d seconds...\n", delay/time.Second)
		time.Sleep(delay)
	}
	fmt.Println("❌ Maximum attempts reached. Process failed.")
	return false
}

// Main communication handler with error recovery
func communicationHandler(data string) bool {
	fmt.Println("\n🌌 Starting Space Communication System 🌌")
	fmt.Printf("Sending data: %s\n", data)

	// Step 1: Perform Cyclic Redundancy Check (CRC)
	if !reattemptLogic(cyclicRedundancyCheck, data, 5, 2*time.Second) {
		fmt.Println("🚨 Critical Failure: CRC reattempts exhausted. Aborting mission.")
		return false
	}

	// Step 2: Perform Low-Density Parity Check (LDPC)
	if !reattemptLogic(lowDensityParityCheck, data, 5, 2*time.Second) {
		fmt.Println("🚨 Critical Failure: LDPC reattempts exhausted. Aborting mission.")
		return false
	}

	// Step 3: Apply Forward Error Correction (FEC)
	if !reattemptLogic(forwardErrorCorrection, data, 5, 2*time.Second) {
		fmt.Println("🚨 Critical Failure: FEC reattempts exhausted. Aborting mission.")
		return false
	}

	fmt.Println("✅ All error checks passed. Data transmission successful.")
	return true
}

// Simulate the space communication system
func simulateSpaceCommunication() {
	testData := "110101101011" // Sample binary data
	if communicationHandler(testData) {
		fmt.Println("\n🚀 Mission Success: Data transmitted without errors.")
	} else {
		fmt.Println("\n🚀 Mission Failure: Could not recover from errors.")
	}
}
