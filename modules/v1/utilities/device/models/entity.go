package models

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
