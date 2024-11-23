package communication

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// DopplerData represents a single data packet with its frequency and velocity.
type DopplerData struct {
	Frequency    float64
	Velocity     float64
	SignalStrength float64
	Timestamp    time.Time
	ID           string
}

// DopplerSorter contains methods for sorting Doppler data.
type DopplerSorter struct {
	data      []DopplerData
	mutex     sync.Mutex
	errorLog  []string
	tolerance float64
}

// NewDopplerSorter initializes a new DopplerSorter with a given tolerance for frequency shifts.
func NewDopplerSorter(tolerance float64) *DopplerSorter {
	return &DopplerSorter{
		data:      []DopplerData{},
		tolerance: tolerance,
	}
}

// AddData adds a new DopplerData entry to the sorter.
func (ds *DopplerSorter) AddData(data DopplerData) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.data = append(ds.data, data)
}

// SortData sorts Doppler data based on frequency and velocity.
func (ds *DopplerSorter) SortData() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	sort.SliceStable(ds.data, func(i, j int) bool {
		if math.Abs(ds.data[i].Frequency-ds.data[j].Frequency) <= ds.tolerance {
			return ds.data[i].Velocity < ds.data[j].Velocity
		}
		return ds.data[i].Frequency < ds.data[j].Frequency
	})
}

// ValidateData checks for anomalies in the Doppler data and logs errors if found.
func (ds *DopplerSorter) ValidateData() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	for _, entry := range ds.data {
		if entry.Frequency <= 0 {
			ds.errorLog = append(ds.errorLog, fmt.Sprintf("Invalid frequency for data ID %s at %v", entry.ID, entry.Timestamp))
		}
		if entry.SignalStrength < 0 {
			ds.errorLog = append(ds.errorLog, fmt.Sprintf("Negative signal strength for data ID %s at %v", entry.ID, entry.Timestamp))
		}
	}
}

// GetErrorLog returns the error log for inspection.
func (ds *DopplerSorter) GetErrorLog() []string {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	return ds.errorLog
}

// SimulateRealTimeData generates mock Doppler data in real-time.
func (ds *DopplerSorter) SimulateRealTimeData(generator func() DopplerData, count int) {
	for i := 0; i < count; i++ {
		newData := generator()
		ds.AddData(newData)
		time.Sleep(50 * time.Millisecond) // Simulate incoming data interval
	}
}

// DopplerEffect applies the Doppler shift formula to calculate shifted frequency.
func DopplerEffect(originalFreq, velocity, speedOfLight float64) float64 {
	return originalFreq * ((speedOfLight + velocity) / speedOfLight)
}

// GenerateMockData creates mock DopplerData for testing.
func GenerateMockData() DopplerData {
	velocity := randomFloat(-5000, 5000) // in m/s
	frequency := DopplerEffect(2.4e9, velocity, 3e8) // 2.4 GHz base frequency
	signalStrength := randomFloat(0, 1)
	return DopplerData{
		Frequency:    frequency,
		Velocity:     velocity,
		SignalStrength: signalStrength,
		Timestamp:    time.Now(),
		ID:           randomString(10),
	}
}

// AnalyzeDopplerTrends calculates statistics on the Doppler data set.
func (ds *DopplerSorter) AnalyzeDopplerTrends() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	var (
		totalFreq       float64
		totalVelocity   float64
		count           int
		highSignalCount int
	)

	for _, entry := range ds.data {
		totalFreq += entry.Frequency
		totalVelocity += entry.Velocity
		count++
		if entry.SignalStrength > 0.8 {
			highSignalCount++
		}
	}

	if count > 0 {
		fmt.Printf("Average Frequency: %.2f Hz\n", totalFreq/float64(count))
		fmt.Printf("Average Velocity: %.2f m/s\n", totalVelocity/float64(count))
		fmt.Printf("High Signal Strength Percentage: %.2f%%\n", float64(highSignalCount)/float64(count)*100)
	} else {
		fmt.Println("No data available for analysis.")
	}
}

// randomFloat generates a random float between min and max.
func randomFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
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

// LogDopplerData logs detailed information of each DopplerData entry.
func (ds *DopplerSorter) LogDopplerData() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	for _, entry := range ds.data {
		fmt.Printf("ID: %s | Frequency: %.2f Hz | Velocity: %.2f m/s | Signal Strength: %.2f | Timestamp: %v\n",
			entry.ID, entry.Frequency, entry.Velocity, entry.SignalStrength, entry.Timestamp)
	}
}

// ClearData clears all the Doppler data from the sorter.
func (ds *DopplerSorter) ClearData() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.data = []DopplerData{}
	ds.errorLog = []string{}
}

// ProcessRealTimeSorting simulates continuous data processing with sorting and validation.
func (ds *DopplerSorter) ProcessRealTimeSorting(generator func() DopplerData, duration time.Duration) {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				ds.AddData(generator())
				ds.SortData()
				ds.ValidateData()
				time.Sleep(100 * time.Millisecond) // Simulate processing time
			}
		}
	}()

	time.Sleep(duration)
	close(stop)
	fmt.Println("Real-time sorting process stopped.")
}

// SaveErrorLog saves error logs to a file for post-mission analysis.
func (ds *DopplerSorter) SaveErrorLog(filename string) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, log := range ds.errorLog {
		file.WriteString(log + "\n")
	}
	return nil
}

// AdvancedSorting applies dynamic sorting based on environmental parameters.
func (ds *DopplerSorter) AdvancedSorting(environmentalFactor float64) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	sort.SliceStable(ds.data, func(i, j int) bool {
		scoreI := ds.data[i].Frequency*environmentalFactor - ds.data[i].Velocity
		scoreJ := ds.data[j].Frequency*environmentalFactor - ds.data[j].Velocity
		return scoreI > scoreJ
	})
}