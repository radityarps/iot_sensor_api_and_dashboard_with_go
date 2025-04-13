package controllers

import (
	"iot_api_with_go/database"
	"iot_api_with_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetStatusSensor(c *gin.Context) {
    var sensors []models.SensorValue

    // Ambil data sensor dari database + Preload SensorData
    if err := database.GetDB().Preload("SensorData").Find(&sensors).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
        return
    }

    // Menambahkan status berdasarkan nilai sensor
    var response []gin.H
    for _, sensor := range sensors {
        status := ""
        if sensor.Name == "ldr" {
            switch {
            case sensor.Value > 1500:
                status = "Terang"
            case sensor.Value > 1000:
                status = "Cerah"
            case sensor.Value > 500:
                status = "Redup"
            default:
                status = "Gelap"
            }
        } else if sensor.Name == "dht_kelembapan" {
            switch {
            case sensor.Value > 80:
                status = "Sangat Lembap"
            case sensor.Value > 60:
                status = "Lembap"
            case sensor.Value > 40:
                status = "Normal"
            case sensor.Value > 20:
                status = "Kering"
            default:
                status = "Sangat Kering"
            }
        }

        // Ambil semua sensor data berdasarkan ID sensor ini
        var sensorDataList []gin.H
        for _, data := range sensor.SensorData {
            sensorDataList = append(sensorDataList, gin.H{
                "id":         data.ID,
                "value":      data.Value,
                "timestamps": data.Timestamps,
            })
        }

        // Tambahkan response utama
        data := gin.H{
            "id":          sensor.ID,
            "name":        sensor.Name,
            "value":       sensor.Value,
            "sensor_data": sensorDataList, // Tambahkan sensor data
        }

        // Tambahkan status jika sensor adalah ldr atau dht_kelembapan
        if sensor.Name == "ldr" || sensor.Name == "dht_kelembapan" {
            data["status"] = status
        }

        response = append(response, data)
    }

    // Kirim response dalam format JSON
    c.JSON(http.StatusOK, response)
}



func UpdateStatusSensor(c *gin.Context) {
	id := c.Param("id")

	var sensor models.SensorValue
	if err := database.DB.First(&sensor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
		return
	}

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&sensor)
	c.JSON(http.StatusOK, sensor)
}

func AddDataSensor(c *gin.Context) {
	var sensorData models.SensorData

	// Bind JSON dari request ke struct SensorData
	if err := c.ShouldBindJSON(&sensorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah sensor dengan ID yang diberikan ada di tabel SensorValue
	var sensorValue models.SensorValue
	if err := database.DB.First(&sensorValue, sensorData.SensorValueID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SensorValue not found"})
		return
	}

	// Simpan data ke database
	if err := database.DB.Create(&sensorData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data sensor berhasil ditambahkan",
		"data":    sensorData,
	})
}

