RF Signal Frequency Simulation and Doppler Sorting System
Project Overview
This project provides a simulation and analysis framework for RF (Radio Frequency) signals and their behavior under varying conditions. The system supports dynamic frequency shifts between predefined points, ensuring synchronization between the sending and receiving frequencies. It includes advanced functionalities like Doppler sorting, signal trend analysis, error validation, and real-time logging.

The implementation is designed for research and testing purposes, such as RF communication systems, signal stability, and Doppler shift analysis in real-world applications.

Features
RF Signal Simulation
Dynamic Frequency Changes:

Simulates RF signals alternating between four predefined frequency points (2.4 GHz, 5.8 GHz, 8.5 GHz, and 12 GHz).
Introduces small shifts between these points to mimic real-world behavior.
Real-Time Signal Generation:

Generates RF signals in real-time with associated attributes:
Frequency
Signal Strength
Timestamp
Unique ID
Velocity (for Doppler shift calculations)
Environmental Effects:

Simulates environmental factors like random perturbations and velocity-induced Doppler shifts.
Doppler Sorting and Analysis
Doppler Data Sorting:

Sorts Doppler data based on frequency and velocity using customizable tolerances.
Supports advanced sorting techniques influenced by environmental factors.
Data Validation:

Identifies anomalies in signal data:
Invalid or zero frequencies.
Negative signal strengths.
Statistical Analysis:

Calculates average frequency, velocity, and high-signal percentages.
Logging and Error Handling
Error Logging:

Logs all detected errors (e.g., invalid data) for review.
Provides options to save logs to external files.
Detailed Signal Logs:

Logs each signal's properties for further analysis:
Frequency
Velocity
Signal Strength
Timestamp
Real-Time Data Processing
Continuous Sorting:

Processes and sorts real-time Doppler data streams.
Maintains synchronization between sending and receiving frequencies.
Trend Analysis:

Detects trends in signal frequencies and highlights strong signals (strength > 0.8).
Modularity and Scalability
Easily extensible to include new frequency points, environmental factors, or analysis methods.
Designed for efficient handling of large data sets in real-time.
Setup Instructions
Prerequisites
Ensure the following are installed on your system:

Go (Golang): Version 1.18 or higher.
Git: For cloning the repository.
Installation
Clone the repository:

bash
Copy code
git clone <repository-url>
cd <repository-folder>
Install dependencies (if any).

Run the program:

bash
Copy code
go run main.go
Configuration Options
Frequency Points
The system uses predefined frequency points. You can modify these in the frequencyPoints array in the code:

go
Copy code
frequencyPoints := []float64{
    2.4e9, // 2.4 GHz
    5.8e9, // 5.8 GHz
    8.5e9, // 8.5 GHz
    12.0e9, // 12 GHz
}
Simulation Duration
By default, the simulation runs for 30 seconds. To change this, update the duration in the RunTestSimulation function:

go
Copy code
RunTestSimulation(frequencyPoints, time.Duration(60)*time.Second) // Runs for 60 seconds
Example Output
plaintext
Copy code
Generating real-time RF signals...
ID: AJD89F62K | Frequency: 2400000000.00 Hz | Signal Strength: 0.76 | Timestamp: 2024-11-23 10:00:05
ID: BLF12G45R | Frequency: 5800000000.00 Hz | Signal Strength: 0.81 | Timestamp: 2024-11-23 10:00:07
Strong signal detected: BLF12G45R | Strength: 0.81
Average Frequency: 4100000000.00 Hz
High Signal Strength Percentage: 12.50%
Key Functions
Signal Simulation
GenerateRFSignal(): Generates a random RF signal with dynamic frequency and strength.

DopplerEffect(): Applies the Doppler shift formula to calculate frequency shifts based on velocity and the speed of light.

Doppler Data Processing
SortData(): Sorts Doppler data based on frequency and velocity.

ValidateData(): Checks for anomalies like invalid frequencies or negative signal strengths.

AnalyzeDopplerTrends(): Analyzes average frequency, velocity, and signal trends.

Real-Time Synchronization
SynchronizeFrequencies(): Ensures the sending and receiving frequencies are always aligned.

SimulateRealTimeData(): Continuously generates and processes Doppler data for the defined duration.