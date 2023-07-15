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

	getDetailDevice, err := s.repository.GetDeviceByAntares(antaresDeviceID)
	if data.Aerator_status == 1 {
		data.Aerator_status = 11
	} else if data.Aerator_status == 0 {
		data.Aerator_status = 10
	} else {
		data.Aerator_status = 10
	}

	if data.Header == 1 {
		err, _ = s.repository.BindSensorData(getDetailDevice.Device_id, data)
	} else {
		err = fmt.Errorf("Data Is Not From Microcontroller")
	}
	return data, err
}

func (s *service) GetAllDevices(user_id string) ([]models.Device, error) {
	device, _ := s.repository.GetAllDevices(user_id)
	for i, d := range device {
		dd := &device[i]
		dd.Date_updated_formatter = d.Date_updated.Format(time.ANSIC)
	}

	return device, nil
}

func (s *service) GetDeviceHistory(user_id string) ([]models.DeviceHistory, string, error) {
	history, _ := s.repository.GetDeviceHistory(user_id)
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
	DetailDevice, err := s.repository.GetDeviceById(u_id, d_id)
	DetailDevice.Date_updated_formatter = DetailDevice.Date_updated.Format(time.ANSIC)
	return DetailDevice, err
}

func (s *service) GetDeviceHistoryById(d_id string, u_id string) ([]models.DeviceHistory, string, error) {
	history, err := s.repository.GetDeviceHistoryById(d_id, u_id)
	for i, d := range history {
		dd := &history[i]
		dd.History_date_formatter = d.History_date.Format(time.ANSIC)
	}

	conv, _ := s.json.Marshal(history)
	return history, string(conv), err
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

func (s *service) Calibration(input models.PhCalibration) error {
	return s.repository.Calibration(input)
}

func (s *service) PostCalibrationAntares(token string, input models.PhCalibration) error {
	return s.repository.PostCalibrationAntares(token, input)
}
