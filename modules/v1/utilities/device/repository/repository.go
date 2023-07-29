package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *repository) BindSensorData(Device_id string, input models.ConnectionDat) (error, error) {
	err := r.db.Exec("INSERT INTO device_history (status_id, mode_id, device_id, temperature, ph, dissolved_oxygen, ph_calibration_firstval, ph_calibration_secval, ph_adc, history_date) VALUES (?,?,?,?,?,?,?,?,?,now())", input.Aerator_status, input.Aerator_mode, Device_id, input.Temperature, input.Ph, input.Dissolved_oxygen, input.Ph_calibration_firstval, input.Ph_calibration_secval, input.Ph_adc).Error
	err2 := r.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, ph_calibration_firstval = ?, ph_calibration_secval = ?, date_updated = now() WHERE device_id = ?", input.Aerator_status, input.Aerator_mode, input.Ph_calibration_firstval, input.Ph_calibration_secval, Device_id).Error
	return err, err2
}

func (r *repository) GetAllDevices(user_id string) ([]models.Device, error) {
	var device []models.Device

	err := r.db.Raw("select d.device_id, d.antares_id, d.device_name, d.device_location, d.mode_id, dm.mode_name, d.status_id, ds.status_name, d.brand_id,b.brand_name, d.user_id, d.latitude, d.longitude, d.date_created, d.date_updated from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id inner join brand b on b.brand_id = d.brand_id where d.user_id = ?", user_id).Scan(&device).Error
	return device, err
}

func (r *repository) GetDeviceByAntares(antaresDeviceID string) (models.Device, error) {
	var device models.Device
	err := r.db.Raw("select * from devices where antares_id = ?", antaresDeviceID).Scan(&device).Error
	return device, err
}

func (r *repository) GetDeviceHistory(user_id string) ([]models.DeviceHistory, error) {
	var DeviceHistory []models.DeviceHistory
	err := r.db.Raw("select d.device_id, d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.ph_adc, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id where d.user_id = ? and NOT(dh.dissolved_oxygen = 0) ORDER BY dh.history_id DESC LIMIT 250", user_id).Scan(&DeviceHistory).Error
	return DeviceHistory, err
}

func (r *repository) Control(id string, power string, mode string) error {
	err := r.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, date_updated = now() WHERE device_id = ?", power, mode, id).Error
	return err
}

func (r *repository) PostControlAntares(antares_id string, token string, power string, mode string) error {
	data := "\r\n{\r\n  \"m2m:cin\": {\r\n    \"con\": \"{ \\\"source_data\\\":\\\"website\\\" , \\\"header\\\": 2 , \\\"aerator_mode\\\":" + mode + ", \\\"aerator_status\\\":" + power + "}\"\r\n    }\r\n}"

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

func (r *repository) AddDevice(input models.DeviceInput, user_id string) error {
	bindUuid := uuid.New()
	err := r.db.Exec("INSERT INTO devices (device_id, device_name, antares_id, device_location, status_id, latitude, longitude, brand_id, user_id, mode_id, date_created, date_updated) VALUES (?,?,?,?,10,?,?,?,?,?,now(),now())", bindUuid, input.Device_name, input.Antares_id, input.Device_location, input.Latitude, input.Longitude, input.Brand_id, user_id, input.Mode_id).Error
	return err
}

func (r *repository) GetDeviceById(u_id string, d_id string) (models.Device, error) {
	var device models.Device
	err := r.db.Raw("select d.device_id, d.antares_id, d.device_name, d.device_location, d.mode_id, dm.mode_name, d.status_id, ds.status_name, d.brand_id,b.brand_name, d.user_id, d.latitude, d.longitude, d.ph_calibration_firstval, d.ph_calibration_secval, d.date_created, d.date_updated from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id inner join brand b on b.brand_id = d.brand_id where d.device_id = ? and user_id = ?", d_id, u_id).Scan(&device).Error
	return device, err
}

func (r *repository) GetDeviceHistoryById(d_id string, u_id string) ([]models.DeviceHistory, error) {
	var DeviceHistory []models.DeviceHistory
	err := r.db.Raw("select d.device_id, d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.ph_calibration_firstval, dh.ph_calibration_secval, dh.ph_adc, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id where d.device_id = ? and d.user_id = ? and NOT(dh.dissolved_oxygen = 0) ORDER BY dh.history_id DESC LIMIT 1000", d_id, u_id).Scan(&DeviceHistory).Error
	return DeviceHistory, err
}

func (r *repository) DeleteDevice(device_id string) error {
	err := r.db.Exec("DELETE FROM device_history WHERE device_id = ?", device_id).Error
	err = r.db.Exec("DELETE FROM devices WHERE device_id = ?", device_id).Error
	return err
}

func (r *repository) GetDeviceBrands() ([]models.Device, error) {
	var brand []models.Device
	err := r.db.Raw("select * from brand").Scan(&brand).Error
	return brand, err
}

func (r *repository) UpdateDeviceById(up_input models.DeviceInput, device_id string) error {
	err := r.db.Exec("update devices set device_name = ?, antares_id = ?, device_location = ?, latitude = ?, longitude = ?, brand_id = ?, mode_id = ?, date_updated = now() where device_id = ?", up_input.Device_name, up_input.Antares_id, up_input.Device_location, up_input.Latitude, up_input.Longitude, up_input.Brand_id, up_input.Mode_id, device_id).Error
	return err
}

func (r *repository) Calibration(input models.PhCalibration) error {
	err := r.db.Exec("update devices set ph_calibration_firstval = ?, ph_calibration_secval = ?, date_updated = now() where antares_id = ?", input.Ph_calibration_firstval, input.Ph_calibration_secval, input.Antares_id).Error
	return err
}

func (r *repository) PostCalibrationAntares(token string, input models.PhCalibration) error {
	data := "\r\n{\r\n  \"m2m:cin\": {\r\n    \"con\": \"{ \\\"source_data\\\":\\\"website\\\" , \\\"header\\\": 3 , \\\"value_ph_k\\\":" + input.Ph_calibration_firstval + ", \\\"value_ph_x\\\":" + input.Ph_calibration_secval + "}\"\r\n    }\r\n}"

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", r.conf.App.Antares_url+"/cnt-"+input.Antares_id, bytes.NewBuffer([]byte(data)))
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
