# IoT Sensor API and Dashboard Documentation

## Overview

This project provides a system for collecting, storing, and visualizing IoT sensor data. It consists of a Go backend API for handling sensor data operations and an HTML/JavaScript frontend dashboard for real-time visualization.

## System Architecture

The application uses:

- **Go (Golang)** with Gin framework for the backend API
- **GORM** as ORM for database operations
- **MySQL** for data storage
- **HTML/CSS/JavaScript** with Chart.js for the frontend dashboard

## Database Structure

The system uses two primary models:

1. **SensorValue**

   ```go
   type SensorValue struct {
       ID         uint         `json:"id" gorm:"primaryKey"`
       Name       string       `json:"name"`
       Value      int          `json:"value"`
       SensorData []SensorData `json:"sensor_data" gorm:"foreignKey:SensorValueID"`
   }
   ```

2. **SensorData**
   ```go
   type SensorData struct {
       ID            uint        `json:"id" gorm:"primaryKey"`
       SensorValueID uint        `json:"sensor_value_id" gorm:"not null"`
       SensorValue   SensorValue `json:"sensor_value" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
       Value         int         `json:"value"`
       Timestamps    time.Time   `json:"timestamps" gorm:"autoCreateTime"`
   }
   ```

## API Endpoints

### 1. Get All Sensor Data

- **URL:** `GET /`
- **Description:** Retrieves all sensor data with their readings
- **Response:** JSON array of sensor data with historical readings

### 2. Update Sensor Status

- **URL:** `PUT /sensor/:id`
- **Description:** Updates information of a specific sensor
- **Parameters:**
  - `id`: Sensor ID (URL parameter)
- **Request Body:** JSON with sensor properties
- **Response:** Updated sensor data

### 3. Add Sensor Reading

- **URL:** `POST /sensor`
- **Description:** Adds a new sensor reading data point
- **Request Body:** JSON with sensor reading details
- **Response:** Confirmation message with saved data

## Dashboard Features

The dashboard ([dashboard.html](c:\Folder%20Utama\Documents\Kuliah\Semester%204\Tugas\IoT\Bab%204\iot_sensor_api_and_dashboard_with_go\dashboard.html)) provides:

1. **Real-time monitoring** of sensor data, refreshing every 5 seconds
2. **Interactive charts** showing historical sensor values
3. **Statistical information** including minimum, maximum, and average values
4. **Status indicators** for light levels (LDR sensor)

Currently supported sensors:

- Light Sensor (LDR) - values in lux
- Temperature Sensor (DHT) - values in °C

## Installation and Setup

### Prerequisites

- Go (1.24.1 or later)
- MySQL server

### Setup Steps

1. **Clone the repository**

2. **Create MySQL database**

   ```sql
   CREATE DATABASE iot_pertemuan_3;
   ```

3. **Configure database connection**
   The connection string is located in database.go. Default configuration:

   ```go
   dsn := "root:@tcp(127.0.0.1:3306)/iot_pertemuan_3?charset=utf8mb4&parseTime=True&loc=Local"
   ```

   Modify this if your MySQL configuration is different.

4. **Install dependencies**

   ```bash
   go mod download
   ```

5. **Run the application**

   ```bash
   go run main.go
   ```

6. **Access the dashboard**
   Open dashboard.html in your web browser, or serve it through a web server.

## Project Structure

```
.
├── controllers/
│   └── sensor_controller.go   # API endpoint handlers
├── database/
│   └── database.go            # Database connection setup
├── models/
│   ├── sensor.go              # SensorValue model definition
│   └── sensor_data.go         # SensorData model definition
├── routes/
│   └── routes.go              # API route definitions
├── dashboard.html             # Frontend dashboard
├── go.mod                     # Go module definition
├── go.sum                     # Go dependency checksums
└── main.go                    # Application entry point
```

## How It Works

1. The application automatically creates necessary database tables on startup through GORM's AutoMigrate.
2. The API endpoints handle sensor data operations (retrieve, update, add).
3. The dashboard polls the API every 5 seconds to get updated sensor data.
4. Charts and statistics are dynamically updated to visualize the real-time data.

## Customization

- To add new sensor types, update the models and controllers accordingly.
- The dashboard can be extended to support additional sensors by modifying the HTML/JavaScript code.

## Note

This project uses CORS middleware to allow cross-origin requests, making it suitable for development environments where the frontend and API may be served from different origins.
