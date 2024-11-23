package communication

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"sync"
	"time"
)

// CarrierSignal represents a signal in the system.
type CarrierSignal struct {
	Frequency     float64 // in Hz
	Phase         float64 // in radians
	Amplitude     float64 // in arbitrary units
	NoiseLevel    float64 // in dB
	SignalID      string
	Timestamp     time.Time
}

// CarrierSync handles synchronization of carrier signals.
type CarrierSync struct {
	signals      []CarrierSignal
	mutex        sync.Mutex
	errorLog     []string
	syncTolerance float64 // Tolerance in Hz
}

// NewCarrierSync initializes a new CarrierSync object.
func NewCarrierSync(tolerance float64) *CarrierSync {
	return &CarrierSync{
		signals:      []CarrierSignal{},
		syncTolerance: tolerance,
	}
}

// AddSignal adds a new CarrierSignal for synchronization.
func (cs *CarrierSync) AddSignal(signal CarrierSignal) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.signals = append(cs.signals, signal)
}

// SimulateSignal generates a simulated carrier signal.
func SimulateSignal() CarrierSignal {
	baseFrequency := 1.8e9 // 1.8 GHz
	noise := randomFloat(0.01, 0.1) // Simulated noise
	phase := randomFloat(0, 2*math.Pi)
	amplitude := randomFloat(0.8, 1.2)

	return CarrierSignal{
		Frequency:  baseFrequency + randomFloat(-500, 500), // ±500 Hz deviation
		Phase:      phase,
		Amplitude:  amplitude,
		NoiseLevel: noise,
		SignalID:   randomString(10),
		Timestamp:  time.Now(),
	}
}

// AlignPhase aligns the phase of the given signals.
func (cs *CarrierSync) AlignPhase() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for i := range cs.signals {
		signal := &cs.signals[i]
		alignedPhase := math.Mod(signal.Phase, 2*math.Pi)
		if alignedPhase < 0 {
			alignedPhase += 2 * math.Pi
		}
		fmt.Printf("Signal %s phase aligned: %.2f -> %.2f radians\n", signal.SignalID, signal.Phase, alignedPhase)
		signal.Phase = alignedPhase
	}
}

// CorrectFrequency attempts to correct frequency deviations within tolerance.
func (cs *CarrierSync) CorrectFrequency() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for i := range cs.signals {
		signal := &cs.signals[i]
		if math.Abs(signal.Frequency-1.8e9) > cs.syncTolerance {
			correctedFrequency := 1.8e9
			fmt.Printf("Signal %s frequency corrected: %.2f Hz -> %.2f Hz\n", signal.SignalID, signal.Frequency, correctedFrequency)
			signal.Frequency = correctedFrequency
		} else {
			fmt.Printf("Signal %s frequency within tolerance: %.2f Hz\n", signal.SignalID, signal.Frequency)
		}
	}
}

// CalculateSignalStrength calculates signal-to-noise ratio (SNR) for a signal.
func (cs *CarrierSync) CalculateSignalStrength() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for _, signal := range cs.signals {
		snr := 10 * math.Log10(signal.Amplitude/signal.NoiseLevel)
		fmt.Printf("Signal %s SNR: %.2f dB\n", signal.SignalID, snr)
		if snr < 10.0 {
			cs.errorLog = append(cs.errorLog, fmt.Sprintf("Low SNR for signal %s at %v", signal.SignalID, signal.Timestamp))
		}
	}
}

// Synchronize synchronizes all carrier signals.
func (cs *CarrierSync) Synchronize() {
	fmt.Println("\nStarting Carrier Synchronization...")
	cs.AlignPhase()
	cs.CorrectFrequency()
	cs.CalculateSignalStrength()
	fmt.Println("Carrier Synchronization Completed.")
}

// SaveErrorLog saves the error log to a file for post-analysis.
func (cs *CarrierSync) SaveErrorLog(filename string) error {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, log := range cs.errorLog {
		file.WriteString(log + "\n")
	}
	return nil
}

// AdvancedNoiseReduction applies noise reduction techniques to signals.
func (cs *CarrierSync) AdvancedNoiseReduction() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for i := range cs.signals {
		signal := &cs.signals[i]
		originalNoise := signal.NoiseLevel
		signal.NoiseLevel = originalNoise * 0.8 // Example: reduce noise by 20%
		fmt.Printf("Signal %s noise reduced: %.2f -> %.2f\n", signal.SignalID, originalNoise, signal.NoiseLevel)
	}
}

// LogSignalDetails logs details of all signals.
func (cs *CarrierSync) LogSignalDetails() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for _, signal := range cs.signals {
		fmt.Printf("SignalID: %s | Frequency: %.2f Hz | Phase: %.2f radians | Amplitude: %.2f | Noise: %.2f dB | Timestamp: %v\n",
			signal.SignalID, signal.Frequency, signal.Phase, signal.Amplitude, signal.NoiseLevel, signal.Timestamp)
	}
}

// Helper Functions

// randomFloat generates a random float between min and max.
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// randomString generates a random alphanumeric string of given length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// SimulateRealTimeProcessing simulates continuous signal synchronization over time.
func (cs *CarrierSync) SimulateRealTimeProcessing(duration time.Duration) {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				newSignal := SimulateSignal()
				cs.AddSignal(newSignal)
				cs.Synchronize()
				time.Sleep(100 * time.Millisecond) // Simulate processing interval
			}
		}
	}()

	time.Sleep(duration)
	close(stop)
	fmt.Println("Real-time processing simulation ended.")
}

// TestCarrierSync runs tests on the CarrierSync system.
func TestCarrierSync() {
	cs := NewCarrierSync(50.0) // 50 Hz tolerance
	fmt.Println("Simulating Signal Generation...")
	for i := 0; i < 20; i++ {
		cs.AddSignal(SimulateSignal())
	}
	fmt.Println("Initial Signal Details:")
	cs.LogSignalDetails()
	cs.Synchronize()
	cs.AdvancedNoiseReduction()
	cs.Synchronize()

	fmt.Println("Saving Error Log...")
	if err := cs.SaveErrorLog("carrier_sync_errors.log"); err != nil {
		fmt.Printf("Error saving log: %v\n", err)
	}
	fmt.Println("Error log saved successfully.")
}