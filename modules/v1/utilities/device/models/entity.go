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
	Device_id       string
	Antares_id      string
	Device_name     string
	Device_location string
	Device_label    string
	Status_id       string
	Status_name     string
	Mode_id         string
	Mode_name       string
	Brand           string
	User_id         string
	Latitude        string
	Longitude       string
	Date_created    time.Time
	Date_updated    time.Time
}
