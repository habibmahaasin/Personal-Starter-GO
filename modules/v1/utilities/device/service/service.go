package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"encoding/json"
	"fmt"
)

func (n *service) GetLatestCon(token string) (models.Received, error) {
	getLatestData, err := n.repository.GetLatestCon(token)
	return getLatestData, err
}

func (n *service) GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error) {
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	fmt.Println("Data sensor terbaru :", data)
	err = n.repository.BindSensorData(antaresDeviceID, data)
	return data, err
}
