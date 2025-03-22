package models

type SensorValue struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	SensorData  []SensorData `json:"sensor_data" gorm:"foreignKey:SensorValueID"` 
}
