package models

type Sensor struct {
	AssignedEEP       EEP                   `json:"assigned_eep"`
	TelegramType      string                `json:"telegram_type"`
	SensorDescription string                `json:"sensor_description"`
	Data              map[string]SensorData `json:"data"`
}

type SensorData struct {
	Shortcut    string      `json:"shortcut"`
	Value       interface{} `json:"value"`
	Unit        interface{} `json:"unit"`
	Description string      `json:"description"`
}
