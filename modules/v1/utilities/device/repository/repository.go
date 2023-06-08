package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func (r *repository) BindSensorData(Device_id string, input models.ConnectionDat) (error, error) {
	err := r.db.Exec("INSERT INTO device_history (status_id, mode_id, device_id, temperature, ph, dissolved_oxygen, history_date) VALUES (?,?,?,?,?,?,now())", input.Status_device, input.Device_mode, Device_id, input.Temperature, input.Ph, input.Dissolved_oxygen).Error
	err2 := r.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, date_updated = now() WHERE device_id = ?", input.Status_device, input.Device_mode, Device_id).Error
	return err, err2
}

func (r *repository) GetAllDevices() ([]models.Device, error) {
	var device []models.Device

	err := r.db.Raw("select * from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id").Scan(&device).Error
	return device, err
}

func (r *repository) GetDeviceByAntares(antaresDeviceID string) (models.Device, error) {
	var device models.Device
	err := r.db.Raw("select * from devices where antares_id = ?", antaresDeviceID).Scan(&device).Error
	return device, err
}

func (r *repository) GetDeviceHistory() ([]models.DeviceHistory, error) {
	var DeviceHistory []models.DeviceHistory
	err := r.db.Raw("select d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id ORDER BY dh.history_id DESC LIMIT 250").Scan(&DeviceHistory).Error
	return DeviceHistory, err
}

func (r *repository) Control(id string, power string, mode string) error {
	err := r.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, date_updated = now() WHERE device_id = ?", power, mode, id).Error
	return err
}

func (r *repository) PostControlAntares(antares_id string, token string, power string, mode string) error {
	data := "\r\n{\r\n  \"m2m:cin\": {\r\n    \"con\": \"{ \\\"aeratorMode\\\":" + mode + ", \\\"statusDevice\\\":" + power + "}\"\r\n    }\r\n}"

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", r.conf.App.Antares_url+"/cnt-"+antares_id, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("X-M2M-Origin", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
