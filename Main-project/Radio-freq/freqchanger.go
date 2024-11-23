package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"sync"
)

const (
	SPEED_OF_LIGHT = 3e8 // m/s
)

// RFSignal represents the structure of the radio frequency signal.
type RFSignal struct {
	Frequency      float64
	SignalStrength float64
	Timestamp      time.Time
	ID             string
	Transmission   bool
}

// RFHandler handles the sending and receiving of RF signals.
type RFHandler struct {
	mu             sync.Mutex
	signals        []RFSignal
	frequencyPoints []float64
}

// NewRFHandler initializes a new RFHandler with specified frequency points.
func NewRFHandler(frequencyPoints []float64) *RFHandler {
	return &RFHandler{
		signals:        []RFSignal{},
		frequencyPoints: frequencyPoints,
	}
}

// GenerateRandomSignal generates a new RFSignal with a random signal strength.
func (rfh *RFHandler) GenerateRandomSignal() RFSignal {
	rand.Seed(time.Now().UnixNano())
	frequency := rfh.frequencyPoints[rand.Intn(len(rfh.frequencyPoints))]
	signalStrength := rand.Float64()
	return RFSignal{
		Frequency:      frequency,
		SignalStrength: signalStrength,
		Timestamp:      time.Now(),
		ID:             randomString(10),
		Transmission:   true,
	}
}

// AddSignal adds a new RFSignal to the list of signals.
func (rfh *RFHandler) AddSignal(signal RFSignal) {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()
	rfh.signals = append(rfh.signals, signal)
}

// LogSignals logs the current RFSignals.
func (rfh *RFHandler) LogSignals() {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()

	for _, signal := range rfh.signals {
		fmt.Printf("ID: %s | Frequency: %.2f Hz | Signal Strength: %.2f | Timestamp: %v\n",
			signal.ID, signal.Frequency, signal.SignalStrength, signal.Timestamp)
	}
}

// randomString generates a random string of a specified length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// SimulateTransmission simulates the transmission of signals across different frequency points.
func (rfh *RFHandler) SimulateTransmission(duration time.Duration) {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				signal := rfh.GenerateRandomSignal()
				rfh.AddSignal(signal)
				time.Sleep(100 * time.Millisecond) // Simulate transmission interval
			}
		}
	}()

	time.Sleep(duration)
	close(stop)
	fmt.Println("Signal transmission complete.")
}

// SynchronizeFrequencies ensures the transmission and reception frequencies are synchronized.
func (rfh *RFHandler) SynchronizeFrequencies() {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()

	for i := 0; i < len(rfh.signals); i++ {
		for j := i + 1; j < len(rfh.signals); j++ {
			if math.Abs(rfh.signals[i].Frequency-rfh.signals[j].Frequency) > 1e6 { // Synchronize if difference > 1MHz
				fmt.Printf("Synchronizing signal: %s with %s\n", rfh.signals[i].ID, rfh.signals[j].ID)
				// Sync the frequencies here, this is just a placeholder for your logic
				rfh.signals[j].Frequency = rfh.signals[i].Frequency
			}
		}
	}
}

// SimulateSignalShifts applies frequency shifts to simulate RF signal movement.
func (rfh *RFHandler) SimulateSignalShifts() {
	for i := 0; i < len(rfh.signals); i++ {
		rfh.signals[i].Frequency = rfh.frequencyPoints[rand.Intn(len(rfh.frequencyPoints))]
	}
}

// AnalyzeFrequencies calculates statistics on the frequency distribution.
func (rfh *RFHandler) AnalyzeFrequencies() {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()

	var totalFreq float64
	var count int

	for _, signal := range rfh.signals {
		totalFreq += signal.Frequency
		count++
	}

	if count > 0 {
		avgFreq := totalFreq / float64(count)
		fmt.Printf("Average Frequency: %.2f Hz\n", avgFreq)
	} else {
		fmt.Println("No signals for analysis.")
	}
}

// FrequencyAdjuster simulates adjustments made to frequency during transmission.
func (rfh *RFHandler) FrequencyAdjuster() {
	for i := 0; i < len(rfh.signals); i++ {
		rfh.signals[i].Frequency += rand.Float64()*1e5 - 5e4 // Adjust frequency within ±50kHz range
	}
}

// RFAnalyzer provides analysis of the RF signals.
func (rfh *RFHandler) RFAnalyzer() {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()

	// Analyze the frequency shift trend
	var totalShift float64
	var count int
	for _, signal := range rfh.signals {
		shift := signal.Frequency - rfh.frequencyPoints[0] // Compare with the first frequency point
		totalShift += shift
		count++
	}

	if count > 0 {
		avgShift := totalShift / float64(count)
		fmt.Printf("Average Frequency Shift: %.2f Hz\n", avgShift)
	}
}

// MonitorSignalStrength monitors and logs the signal strength at various times.
func (rfh *RFHandler) MonitorSignalStrength() {
	rfh.mu.Lock()
	defer rfh.mu.Unlock()

	for _, signal := range rfh.signals {
		if signal.SignalStrength > 0.8 {
			fmt.Printf("Strong signal detected: %s | Strength: %.2f\n", signal.ID, signal.SignalStrength)
		}
	}
}

// RunTestSimulation runs the entire RF signal transmission and analysis simulation.
func RunTestSimulation(frequencyPoints []float64, duration time.Duration) {
	rfh := NewRFHandler(frequencyPoints)

	rfh.SimulateTransmission(duration)
	rfh.SynchronizeFrequencies()
	rfh.AnalyzeFrequencies()
	rfh.MonitorSignalStrength()

	// Simulate frequency shifts
	rfh.SimulateSignalShifts()

	// Analyze frequency shifts and signal strength again after changes
	rfh.FrequencyAdjuster()
	rfh.RFAnalyzer()
	rfh.LogSignals()
}

func main() {
	// Define frequency points for simulation
	frequencyPoints := []float64{
		2.4e9, // 2.4 GHz
		5.8e9, // 5.8 GHz
		8.5e9, // 8.5 GHz
		12.0e9, // 12 GHz
	}

	// Run simulation for 30 seconds
	RunTestSimulation(frequencyPoints, 30*time.Second)
}
