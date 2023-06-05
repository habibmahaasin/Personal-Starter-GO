package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *repository) GetLatestCon(token string) (models.Received, error) {
	data := models.Received{}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://platform.antares.id:8443/~/antares-cse/antares-id/GuppyTech/AllSensor/la", nil)
	req.Header.Set("X-M2M-Origin", token)
	fmt.Println("ini tokennya : ", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer resp.Body.Close()
	isiBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(isiBody, &data)
	return data, err
}

func (r *repository) BindSensorData(DeviceId string, input models.ConnectionDat) error {
	// err := n.db.Exec("INSERT INTO history (device_id, temp, ph, date_updated) VALUES (?,?,?,?)", DeviceId, input.Temp, input.Ph, time.Now()).Error
	return nil
}

func (r *repository) GetAllDevices() ([]models.Device, error) {
	var device []models.Device

	err := r.db.Raw("select * from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id").Scan(&device).Error
	return device, err
}
