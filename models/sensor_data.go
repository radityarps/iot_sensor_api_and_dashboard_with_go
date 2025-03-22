package models

import "time"

type SensorData struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	SensorValueID uint        `json:"sensor_value_id" gorm:"not null"`
	SensorValue   SensorValue `json:"sensor_value" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Value         int         `json:"value"`
	Timestamps    time.Time   `json:"timestamps" gorm:"autoCreateTime"`
}
