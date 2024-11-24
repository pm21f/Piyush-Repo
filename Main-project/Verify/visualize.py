from flask import Flask, render_template, jsonify
import requests
import threading
import time

app = Flask(__name__)

# Global variable to store the latest telemetry data
latest_data = {
    "timestamps": [],
    "battery_levels": [],
    "radiation_levels": [],
    "signal_strengths": [],
    "distances": []  # Add distance to the data structure
}

def fetch_data():
    while True:
        try:
            response = requests.get("http://localhost:8080/telemetry")
            data = response.json()
            current_time = time.strftime("%H:%M:%S", time.localtime())

            # Update the latest_data dictionary
            latest_data["timestamps"].append(current_time)
            latest_data["battery_levels"].append(data["normal"]["battery_level"])
            latest_data["radiation_levels"].append(data["normal"]["radiation_level"])
            latest_data["signal_strengths"].append(data["normal"]["signal_strength"])
            latest_data["distances"].append(data["normal"]["distance"])  # Include distance

            # Limit data history to the last 20 data points
            if len(latest_data["timestamps"]) > 20:
                for key in latest_data:
                    latest_data[key].pop(0)

            # Wait for 1.3 seconds before fetching new data
            time.sleep(1.3)

        except Exception as e:
            print(f"Error fetching data: {e}")

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/data')
def data():
    return jsonify(latest_data)

if __name__ == '__main__':
    # Start data fetching in a separate thread
    threading.Thread(target=fetch_data, daemon=True).start()
    # Start Flask app
    app.run(debug=True, use_reloader=False)
