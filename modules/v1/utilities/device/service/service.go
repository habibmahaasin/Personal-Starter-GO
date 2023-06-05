package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"encoding/json"
	"fmt"
)

func (s *service) GetLatestCon(token string) (models.Received, error) {
	getLatestData, err := s.repository.GetLatestCon(token)
	return getLatestData, err
}

func (s *service) GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error) {
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	fmt.Println("Data sensor terbaru :", data)
	err = s.repository.BindSensorData(antaresDeviceID, data)
	return data, err
}

func (s *service) GetAllDevices() ([]models.Device, error) {
	return s.repository.GetAllDevices()
}
