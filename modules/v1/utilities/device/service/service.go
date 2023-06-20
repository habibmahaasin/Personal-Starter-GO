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

	err, _ = s.repository.BindSensorData(getDetailDevice.Device_id, data)
	return data, err
}

func (s *service) GetAllDevices() ([]models.Device, error) {
	return s.repository.GetAllDevices()
}

func (s *service) GetDeviceHistory() ([]models.DeviceHistory, string, error) {
	history, _ := s.repository.GetDeviceHistory()
	for i, d := range history {
		dd := &history[i]
		dd.History_date_formatter = d.History_date.Format(time.ANSIC)
	}

	conv, _ := s.json.Marshal(history)
	return history, string(conv), nil
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

func (s *service) AddDevice(input models.DeviceInput, user_id string) error {
	return s.repository.AddDevice(input, user_id)
}

func (s *service) GetDeviceById(u_id string, d_id string) (models.Device, error) {
	DetailDevice, _ := s.repository.GetDeviceById(u_id, d_id)
	DetailDevice.Date_updated_formatter = DetailDevice.Date_updated.Format(time.ANSIC)
	return DetailDevice, nil
}

func (s *service) GetDeviceHistoryById(d_id string) ([]models.DeviceHistory, string, error) {
	history, _ := s.repository.GetDeviceHistoryById(d_id)
	for i, d := range history {
		dd := &history[i]
		dd.History_date_formatter = d.History_date.Format(time.ANSIC)
	}

	conv, _ := s.json.Marshal(history)
	return history, string(conv), nil
}

func (s *service) DeleteDevice(device_id string) error {
	return s.repository.DeleteDevice(device_id)
}

func (s *service) GetDeviceBrands() ([]models.Device, error) {
	return s.repository.GetDeviceBrands()
}

func (s *service) UpdateDeviceById(up_input models.DeviceInput, device_id string) error {
	return s.repository.UpdateDeviceById(up_input, device_id)
}
