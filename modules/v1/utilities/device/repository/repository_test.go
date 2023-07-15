package repository

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/models"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	config     config.Conf
	mock       sqlmock.Sqlmock
	repository *repository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(1)
	s.config = config.Conf{
		App: config.App{
			Name:       "GuppyTech",
			Port:       "8080",
			Mode:       "Test",
			Url:        "http://localhost",
			Secret_key: "GuppyTest",
		},
	}

	s.repository = NewRepository(s.DB, s.config)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetAllDevices() {
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	returnDeviceData := []models.Device{
		{
			Device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			Antares_id:  "ps9t5UiX15TVLxYB",
			Device_name: "Aerator Utama",
		},
	}

	rows := sqlmock.NewRows([]string{"Device_id", "Antares_id", "Device_name"}).
		AddRow(returnDeviceData[0].Device_id, returnDeviceData[0].Antares_id, returnDeviceData[0].Device_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select d.device_id, d.antares_id, d.device_name, d.device_location, d.mode_id, dm.mode_name, d.status_id, ds.status_name, d.brand_id,b.brand_name, d.user_id, d.latitude, d.longitude, d.date_created, d.date_updated from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id inner join brand b on b.brand_id = d.brand_id where d.user_id = $1")).
		WithArgs(user_id).
		WillReturnRows(rows)

	resp, err := s.repository.GetAllDevices(user_id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), returnDeviceData, resp)
}

func (s *Suite) Test_repository_BindSensorData() {
	var (
		Aerator_mode            = 1
		Aerator_status          = 2
		Temperature             = 30
		Ph                      = 7
		Dissolved_oxygen        = 5
		Ph_calibration_secval   = 10
		Ph_calibration_firstval = 11
		Device_id               = "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	)

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO device_history (status_id, mode_id, device_id, temperature, ph, dissolved_oxygen, ph_calibration_firstval, ph_calibration_secval, history_date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,now())")).
		WithArgs(Aerator_status, Aerator_mode, Device_id, float64(Temperature), float64(Ph), float64(Dissolved_oxygen), float64(Ph_calibration_firstval), float64(Ph_calibration_secval)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE devices SET status_id  = $1, mode_id  = $2, ph_calibration_firstval = $3, ph_calibration_secval = $4, date_updated = now() WHERE device_id = $5")).
		WithArgs(Aerator_status, Aerator_mode, float64(Ph_calibration_firstval), float64(Ph_calibration_secval), Device_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err, err2 := s.repository.BindSensorData(Device_id, models.ConnectionDat{Aerator_mode: Aerator_mode, Aerator_status: Aerator_status, Temperature: float64(Temperature), Ph: float64(Ph), Dissolved_oxygen: float64(Dissolved_oxygen), Ph_calibration_firstval: float64(Ph_calibration_firstval), Ph_calibration_secval: float64(Ph_calibration_secval)})
	require.NoError(s.T(), err)
	require.NoError(s.T(), err2)
	require.Nil(s.T(), deep.Equal(nil, err))
}

func (s *Suite) Test_repository_GetDeviceByAntares() {
	antaresIdInput := "ps9t5UiX15TVLxYB"
	returnDeviceData := models.Device{
		Device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
		Antares_id:  "ps9t5UiX15TVLxYB",
		Device_name: "Aerator Utama",
	}

	rows := sqlmock.NewRows([]string{"Device_id", "Antares_id", "Device_name"}).
		AddRow(returnDeviceData.Device_id, returnDeviceData.Antares_id, returnDeviceData.Device_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select * from devices where antares_id = $1")).
		WithArgs(antaresIdInput).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceByAntares(antaresIdInput)
	require.NoError(s.T(), err)
	require.Equal(s.T(), returnDeviceData, resp)
}

func (s *Suite) Test_repository_GetDeviceHistory() {
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	expectHistory := []models.DeviceHistory{
		{
			History_device_id:        "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			History_device_name:      "Aerator Utama",
			History_status_name:      "Tidak Aktif",
			History_mode_name:        "Otomatis",
			History_ph:               21.97006,
			History_temperature:      27.4375,
			History_dissolved_oxygen: 49.40756,
			History_date:             time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"device_id", "device_name", "status_name", "mode_name", "ph", "temperature", "dissolved_oxygen", "history_date"}).
		AddRow(expectHistory[0].History_device_id, expectHistory[0].History_device_name, expectHistory[0].History_status_name, expectHistory[0].History_mode_name, expectHistory[0].History_ph, expectHistory[0].History_temperature, expectHistory[0].History_dissolved_oxygen, expectHistory[0].History_date)

	s.mock.ExpectQuery(regexp.QuoteMeta("select d.device_id, d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id where d.user_id = $1 ORDER BY dh.history_id DESC LIMIT 250")).
		WithArgs(user_id).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceHistory(user_id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectHistory, resp)
}

func (s *Suite) Test_repository_Control() {
	id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	power := "10"
	mode := "2"

	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE devices SET status_id = $1, mode_id = $2, date_updated = now() WHERE device_id = $3")).
		WithArgs(power, mode, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Control(id, power, mode)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_PostControlAntares() {
	//success case
	serverCase1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer serverCase1.Close()
	s.config.App.Antares_url = serverCase1.URL
	repo := repository{
		s.DB,
		s.config,
	}
	err := repo.PostControlAntares("ps9t5UiX15TVLxYB", "862b34fe2de548cc:cdf66d91b12db8d2", "1", "2")
	require.NoError(s.T(), err)

	//error case http.NewRequest()
	serverCase2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer serverCase2.Close()
	s.config.App.Antares_url = serverCase1.URL + "error"
	repo = repository{
		s.DB,
		s.config,
	}
	err = repo.PostControlAntares("ps9t5UiX15TVLxYB", "862b34fe2de548cc:cdf66d91b12db8d2", "1", "2")
	require.Error(s.T(), err)

	//error case client.Do()
	serverCase3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer serverCase3.Close()
	s.config.App.Antares_url = "serverCase3.URL"
	repo = repository{
		s.DB,
		s.config,
	}
	err = repo.PostControlAntares("ps9t5UiX15TVLxYB", "862b34fe2de548cc:cdf66d91b12db8d2", "1", "2")
	require.Error(s.T(), err)

	//error case ioutil.ReadAll()
	serverCase4 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer serverCase4.Close()
	s.config.App.Antares_url = serverCase4.URL
	repo = repository{
		s.DB,
		s.config,
	}
	err = repo.PostControlAntares("ps9t5UiX15TVLxYB", "862b34fe2de548cc:cdf66d91b12db8d2", "1", "2")
	require.Error(s.T(), err)
}

func (s *Suite) Test_repository_AddDevice() {
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	deviceInput := models.DeviceInput{
		Device_name:     "Aerator Utama",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Jl. Raya Bogor",
		Latitude:        "1.1",
		Longitude:       "1.1",
		Brand_id:        "1",
		Mode_id:         "2",
	}

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO devices (device_id, device_name, antares_id, device_location, status_id, latitude, longitude, brand_id, user_id, mode_id, date_created, date_updated) VALUES ($1,$2,$3,$4,10,$5,$6,$7,$8,$9,now(),now())")).
		WithArgs(
			sqlmock.AnyArg(),
			deviceInput.Device_name,
			deviceInput.Antares_id,
			deviceInput.Device_location,
			deviceInput.Latitude,
			deviceInput.Longitude,
			deviceInput.Brand_id,
			user_id,
			deviceInput.Mode_id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.AddDevice(deviceInput, user_id)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_GetDeviceById() {
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	expected_result := models.Device{
		Device_id:       "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
		Device_name:     "Aerator Utama",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Aquarium Kiri Hitam",
		Status_name:     "Tidak Aktif",
		Mode_name:       "Otomatis",
		Brand_name:      "Amara-666",
	}
	rows := sqlmock.NewRows([]string{"device_id", "device_name", "antares_id", "device_location", "status_name", "mode_name", "brand_name"}).AddRow(expected_result.Device_id, expected_result.Device_name, expected_result.Antares_id, expected_result.Device_location, expected_result.Status_name, expected_result.Mode_name, expected_result.Brand_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select d.device_id, d.antares_id, d.device_name, d.device_location, d.mode_id, dm.mode_name, d.status_id, ds.status_name, d.brand_id,b.brand_name, d.user_id, d.latitude, d.longitude, d.ph_calibration_firstval, d.ph_calibration_secval, d.date_created, d.date_updated from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id inner join brand b on b.brand_id = d.brand_id where d.device_id = $1 and user_id = $2")).
		WithArgs(device_id, user_id).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceById(user_id, device_id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expected_result, resp)
}

func (s *Suite) Test_repository_GetDeviceHistoryById() {
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	expected_result := []models.DeviceHistory{
		{
			History_device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			History_device_name: "Aerator Utama",
			History_status_name: "Tidak Aktif",
			History_mode_name:   "Otomatis",
		},
	}

	rows := sqlmock.NewRows([]string{"device_id", "device_name", "status_name", "mode_name"}).AddRow(expected_result[0].History_device_id, expected_result[0].History_device_name, expected_result[0].History_status_name, expected_result[0].History_mode_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select d.device_id, d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id where d.device_id = $1 and d.user_id = $2 ORDER BY dh.history_id DESC")).
		WithArgs(device_id, user_id).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceHistoryById(device_id, user_id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expected_result, resp)
}

func (s *Suite) Test_repository_GetDeviceBrands() {
	expected_result := []models.Device{
		{
			Brand_id:   1,
			Brand_name: "Amara AA-666",
		},
	}

	rows := sqlmock.NewRows([]string{"brand_id", "brand_name"}).AddRow(expected_result[0].Brand_id, expected_result[0].Brand_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select * from brand")).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceBrands()
	require.NoError(s.T(), err)
	require.Equal(s.T(), expected_result, resp)
}

func (s *Suite) Test_repository_DeleteDevice() {
	d_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM device_history WHERE device_id = $1")).
		WithArgs(d_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM devices WHERE device_id = $1")).
		WithArgs(d_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.DeleteDevice(d_id)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_UpdateDeviceById() {
	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	deviceInput := models.DeviceInput{
		Device_name:     "Aerator Utama",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Jl. Raya Bogor",
		Latitude:        "1.1",
		Longitude:       "1.1",
		Brand_id:        "1",
		Mode_id:         "2",
	}

	s.mock.ExpectExec(regexp.QuoteMeta(`update devices set device_name = $1, antares_id = $2, device_location = $3, latitude = $4, longitude = $5, brand_id = $6, mode_id = $7, date_updated = now() where device_id = $8`)).
		WithArgs(
			deviceInput.Device_name,
			deviceInput.Antares_id,
			deviceInput.Device_location,
			deviceInput.Latitude,
			deviceInput.Longitude,
			deviceInput.Brand_id,
			deviceInput.Mode_id,
			device_id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.UpdateDeviceById(deviceInput, device_id)
	require.NoError(s.T(), err)
}
