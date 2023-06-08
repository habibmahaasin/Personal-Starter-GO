package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"encoding/json"
	"fmt"
	"time"
)

func (s *service) GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error) {
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	getDetailDevice, err := s.repository.GetDeviceByAntares(antaresDeviceID)
	if data.Status_device == 1 {
		data.Status_device = 11
	} else if data.Status_device == 0 {
		data.Status_device = 10
	} else {
		data.Status_device = 10
	}

	fmt.Println(data)
	err, _ = s.repository.BindSensorData(getDetailDevice.Device_id, data)
	return data, err
}

func (s *service) GetAllDevices() ([]models.Device, error) {
	return s.repository.GetAllDevices()
}

func (s *service) GetDeviceHistory() ([]models.DeviceHistory, error) {
	history, _ := s.repository.GetDeviceHistory()
	for i, d := range history {
		dd := &history[i]
		dd.History_date_formatter = d.History_date.Format(time.ANSIC)
	}
	return history, nil
}

func (s *service) Control(id string, power string, mode string) error {
	return s.repository.Control(id, power, mode)
}

func (s *service) PostControlAntares(antares_id string, token string, power string, mode string) error {
	if power == "11" {
		power = "1"
	} else if power == "10" {
		power = "0"
	}
	return s.repository.PostControlAntares(antares_id, token, power, mode)
}
