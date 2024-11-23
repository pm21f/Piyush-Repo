package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Configuration Parameters
const (
	minStableSignalStrength = 2.0
	maxSignalStrength       = 10.0
	maxNoiseLevel           = 5.0
	maxPacketLoss           = 20.0 // Percentage
	maxJitter               = 50.0 // Milliseconds
	maxErrorRate            = 1.0  // Percentage
	retryLimit              = 3
	checkInterval           = 2 * time.Second
	healthSummaryInterval   = 10 * time.Second
)

// Log Severity Levels
const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

// Global Metrics for Analytics
var reconnectionAttempts int
var successfulReconnections int
var totalHealthChecks int
var failedConnections int

// Log an event with a severity level and timestamp
func logEvent(level, message string) {
	fmt.Printf("[%s] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), level, message)
}

// Generate random RF metrics
func generateRFMetrics() (signalStrength, noiseLevel, packetLoss, jitter, errorRate float64) {
	signalStrength = rand.Float64() * maxSignalStrength
	noiseLevel = rand.Float64() * maxNoiseLevel
	packetLoss = rand.Float64() * 100 // Packet loss in percentage
	jitter = rand.Float64() * maxJitter
	errorRate = rand.Float64() * maxErrorRate
	return
}

// Evaluate RF metrics against thresholds
func evaluateMetrics(signalStrength, noiseLevel, packetLoss, jitter, errorRate float64) bool {
	logEvent(INFO, fmt.Sprintf("Metrics - Signal: %.2f, Noise: %.2f, Packet Loss: %.2f%%, Jitter: %.2fms, Error Rate: %.2f%%",
		signalStrength, noiseLevel, packetLoss, jitter, errorRate))

	if signalStrength < minStableSignalStrength {
		logEvent(WARNING, fmt.Sprintf("Low Signal Strength: %.2f (Threshold: %.2f)", signalStrength, minStableSignalStrength))
		return false
	}
	if noiseLevel > maxNoiseLevel {
		logEvent(WARNING, fmt.Sprintf("High Noise Level: %.2f (Threshold: %.2f)", noiseLevel, maxNoiseLevel))
		return false
	}
	if packetLoss > maxPacketLoss {
		logEvent(WARNING, fmt.Sprintf("Excessive Packet Loss: %.2f%% (Threshold: %.2f%%)", packetLoss, maxPacketLoss))
		return false
	}
	if jitter > maxJitter {
		logEvent(WARNING, fmt.Sprintf("High Jitter: %.2fms (Threshold: %.2fms)", jitter, maxJitter))
		return false
	}
	if errorRate > maxErrorRate {
		logEvent(WARNING, fmt.Sprintf("High Error Rate: %.2f%% (Threshold: %.2f%%)", errorRate, maxErrorRate))
		return false
	}
	return true
}

// Attempt reconnection in case of instability
func attemptReconnection() {
	logEvent(INFO, "Initiating reconnection attempts...")
	for attempt := 1; attempt <= retryLimit; attempt++ {
		reconnectionAttempts++
		logEvent(INFO, fmt.Sprintf("Reconnection Attempt %d/%d...", attempt, retryLimit))
		time.Sleep(1 * time.Second)
		if rand.Float64()*10 > 5 { // Simulated success rate
			logEvent(INFO, "Reconnection successful!")
			successfulReconnections++
			return
		}
		logEvent(WARNING, "Reconnection attempt failed.")
	}
	logEvent(ERROR, "All reconnection attempts failed. Manual intervention required.")
	failedConnections++
}

// Check RF stability and manage reconnection logic
func checkRFStability() {
	signalStrength, noiseLevel, packetLoss, jitter, errorRate := generateRFMetrics()

	if !evaluateMetrics(signalStrength, noiseLevel, packetLoss, jitter, errorRate) {
		logEvent(ERROR, "RF Signal Unstable! Attempting to reconnect...")
		attemptReconnection()
	} else {
		logEvent(INFO, "RF Signal Stable.")
	}
}

// Display periodic health summary
func displayHealthSummary() {
	for {
		time.Sleep(healthSummaryInterval)
		totalHealthChecks++
		logEvent(INFO, strings.Repeat("=", 50))
		logEvent(INFO, "System Health Summary")
		logEvent(INFO, fmt.Sprintf("Total Health Checks: %d", totalHealthChecks))
		logEvent(INFO, fmt.Sprintf("Reconnection Attempts: %d", reconnectionAttempts))
		logEvent(INFO, fmt.Sprintf("Successful Reconnections: %d", successfulReconnections))
		logEvent(INFO, fmt.Sprintf("Failed Connections: %d", failedConnections))
		logEvent(INFO, strings.Repeat("=", 50))
	}
}

// Simulate background diagnostics
func runDiagnostics() {
	for {
		time.Sleep(5 * time.Second)
		logEvent(INFO, "Running background diagnostics...")
		time.Sleep(1 * time.Second)
		logEvent(INFO, "Diagnostics completed. All systems nominal.")
	}
}
