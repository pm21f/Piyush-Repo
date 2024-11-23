package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	LowBatteryThreshold    = 20  // Low battery threshold for alerts
	HighRadiationThreshold = 800 // High radiation threshold for alerts
)

// TelemetryData structure to hold telemetry information
type TelemetryData struct {
	BatteryLevel   int    `json:"battery_level"`
	RadiationLevel int    `json:"radiation_level"`
	SignalStrength int    `json:"signal_strength"`
	TransferStatus string `json:"transfer_status"`
	Distance       int    `json:"distance"`
	Alert          string `json:"alert,omitempty"`
}

var telemetryHistory []TelemetryData

// SyncTelemetryData generates random telemetry data
func SyncTelemetryData() TelemetryData {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	batteryLevel := rand.Intn(101)         // Battery level: 0 to 100
	radiationLevel := rand.Intn(1001)      // Radiation level: 0 to 1000
	signalStrength := rand.Intn(26)        // Signal strength: 0 to 25
	distance := 300000 + rand.Intn(200000) // Distance: always greater than 300,000

	transferStatus := "Data successfully reached at base station."
	if signalStrength < 5 { // Simulate an anomaly for low signal strength
		transferStatus = "Anomaly detected! Strength too low, attempting reconnection."
	}

	alert := checkForAlerts(batteryLevel, radiationLevel)

	return TelemetryData{
		BatteryLevel:   batteryLevel,
		RadiationLevel: radiationLevel,
		SignalStrength: signalStrength,
		TransferStatus: transferStatus,
		Distance:       distance,
		Alert:          alert,
	}
}

// Check telemetry data for alert conditions
func checkForAlerts(batteryLevel, radiationLevel int) string {
	var alertMsg string

	if batteryLevel < LowBatteryThreshold {
		alertMsg += fmt.Sprintf("⚠️ Battery level low: %d%%\n", batteryLevel)
	}

	if radiationLevel > HighRadiationThreshold {
		alertMsg += fmt.Sprintf("⚠️ High radiation level detected: %d\n", radiationLevel)
	}

	return alertMsg
}

// Log telemetry data
func logTelemetryData(data TelemetryData) {

	fmt.Println("Starting data analysis:")

	fmt.Println("Battery Level........: ", data.BatteryLevel)
	fmt.Println("Radiation Level............: ", data.RadiationLevel)
	fmt.Println("Signal Strength.................: ", data.SignalStrength)
	fmt.Println("Distance...........................: ", data.Distance)
	fmt.Println("Transfer Status...........................: ", data.TransferStatus)

	if data.Alert != "" {
		fmt.Println("\n🚨 ALERT DETECTED 🚨")
		fmt.Println(data.Alert)
		fmt.Println(".......................................................")
	}

	fmt.Println(drawSignalStrengthGraph(data.SignalStrength))

	// Store the telemetry data
	telemetryHistory = append(telemetryHistory, data)
}

// Draw a terminal-based graph for signal strength
func drawSignalStrengthGraph(strength int) string {
	barHeight := 25
	barWidth := 20
	var graph strings.Builder

	graph.WriteString("Signal Strength Graph:\n")
	for i := barHeight; i > 0; i-- {
		if strength >= i {
			graph.WriteString("| " + strings.Repeat("#", barWidth) + "\n")
		} else {
			graph.WriteString("| " + strings.Repeat(" ", barWidth) + "\n")
		}
	}
	graph.WriteString("+" + strings.Repeat("-", barWidth+1) + "\n")
	graph.WriteString(fmt.Sprintf("Signal Strength.....: %d/25\n\n", strength))
	return graph.String()
}

// Main function
func main() {
	go func() {
		for {
			data := SyncTelemetryData() // Changed the function call to SyncTelemetryData
			logTelemetryData(data)
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/telemetry", func(w http.ResponseWriter, r *http.Request) {
		data := SyncTelemetryData() // Changed the function call to SyncTelemetryData

		response := map[string]interface{}{
			"normal": data,
			"binary": map[string]string{
				"battery_level":   strconv.FormatInt(int64(data.BatteryLevel), 2),
				"radiation_level": strconv.FormatInt(int64(data.RadiationLevel), 2),
				"signal_strength": strconv.FormatInt(int64(data.SignalStrength), 2),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
