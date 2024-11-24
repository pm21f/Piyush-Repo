package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RFConnection simulates a radio frequency connection
type RFConnection struct {
	isConnected bool
}

func NewRFConnection() *RFConnection {
	return &RFConnection{isConnected: true} // Start with a successful connection
}

func (rf *RFConnection) SendData(data TelemetryData) error {
	if !rf.isConnected {
		return fmt.Errorf("RF Connection lost, unable to send data")
	}

	// Simulate a transmission delay
	time.Sleep(500 * time.Millisecond)

	// Simulate random success/failure of data transmission
	if rand.Intn(100) < 10 { // 10% chance of failure
		rf.isConnected = false // Mark connection as lost on failure
		return fmt.Errorf("failed to send data: connection unstable")
	}

	// Successfully sent data
	fmt.Printf("RF Transmission successful: %+v\n", data)
	return nil
}

// Attempt to reconnect
func (rf *RFConnection) Reconnect() {
	fmt.Println("Attempting to reconnect...")
	time.Sleep(2 * time.Second) // Simulate time taken to reconnect
	rf.isConnected = true
	fmt.Println("Reconnection successful!")
}

// Monitor RF connection status
func (rf *RFConnection) Monitor() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if !rf.isConnected {
				rf.Reconnect()
			}
		}
	}()
}
