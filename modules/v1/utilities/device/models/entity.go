package models

import "time"

type Product struct {
	Id     int
	Name   string
	Amount int
	Price  int
}

type Sensor_Data struct {
	Device_id    string
	Temp         int
	Ph           int
	Date_updated string
}

type Device struct {
	Device_id               string
	Antares_id              string
	Device_name             string
	Device_location         string
	Status_id               string
	Status_name             string
	Mode_id                 string
	Mode_name               string
	Brand_id                int
	Brand_name              string
	User_id                 string
	Latitude                string
	Longitude               string
	Ph_calibration_firstval string
	Ph_calibration_secval   string
	Date_created            time.Time
	Date_updated            time.Time
	Date_updated_formatter  string
}

type DeviceHistory struct {
	History_device_id        string    `gorm:"column:device_id"`
	History_device_name      string    `gorm:"column:device_name"`
	History_status_name      string    `gorm:"column:status_name"`
	History_mode_name        string    `gorm:"column:mode_name"`
	History_ph               float32   `gorm:"column:ph"`
	History_temperature      float32   `gorm:"column:temperature"`
	History_dissolved_oxygen float32   `gorm:"column:dissolved_oxygen"`
	History_date             time.Time `gorm:"column:history_date"`
	History_ph_firstval      float32   `gorm:"column:ph_calibration_firstval"`
	History_ph_secval        float32   `gorm:"column:ph_calibration_secval"`
	History_ph_adc           int       `gorm:"column:ph_adc"`
	History_date_formatter   string
}
