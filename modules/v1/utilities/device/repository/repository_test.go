package repository

import (
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

	s.repository = NewRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetAllDevices() {
	returnDeviceData := []models.Device{
		{
			Device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			Antares_id:  "ps9t5UiX15TVLxYB",
			Device_name: "Aerator Utama",
		},
	}

	rows := sqlmock.NewRows([]string{"Device_id", "Antares_id", "Device_name"}).
		AddRow(returnDeviceData[0].Device_id, returnDeviceData[0].Antares_id, returnDeviceData[0].Device_name)

	s.mock.ExpectQuery(regexp.QuoteMeta("select * from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id")).
		WillReturnRows(rows)

	resp, err := s.repository.GetAllDevices()
	require.NoError(s.T(), err)
	require.Equal(s.T(), returnDeviceData, resp)
}

func (s *Suite) Test_repository_BindSensorData() {
	var (
		Device_mode      = 1
		Status_device    = 2
		Temperature      = 30
		Ph               = 7
		Dissolved_oxygen = 5
		Device_id        = "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	)

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO device_history (status_id, mode_id, device_id, temperature, ph, dissolved_oxygen, history_date) VALUES ($1,$2,$3,$4,$5,$6,now())")).
		WithArgs(Status_device, Device_mode, Device_id, float64(Temperature), float64(Ph), float64(Dissolved_oxygen)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE devices SET status_id  = $1, mode_id  = $2, date_updated = now() WHERE device_id = $3")).
		WithArgs(Status_device, Device_mode, Device_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err, err2 := s.repository.BindSensorData(Device_id, models.ConnectionDat{Device_mode: Device_mode, Status_device: Status_device, Temperature: float64(Temperature), Ph: float64(Ph), Dissolved_oxygen: float64(Dissolved_oxygen)})
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

	expectHistory := []models.DeviceHistory{
		{
			History_device_name:      "Aerator Utama",
			History_status_name:      "Tidak Aktif",
			History_mode_name:        "Otomatis",
			History_ph:               21.97006,
			History_temperature:      27.4375,
			History_dissolved_oxygen: 49.40756,
			History_date:             time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"device_name", "status_name", "mode_name", "ph", "temperature", "dissolved_oxygen", "history_date"}).
		AddRow(expectHistory[0].History_device_name, expectHistory[0].History_status_name, expectHistory[0].History_mode_name, expectHistory[0].History_ph, expectHistory[0].History_temperature, expectHistory[0].History_dissolved_oxygen, expectHistory[0].History_date)

	s.mock.ExpectQuery(regexp.QuoteMeta("select d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id ORDER BY dh.history_id DESC LIMIT 250")).
		WillReturnRows(rows)

	resp, err := s.repository.GetDeviceHistory()
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
	repo := repository{
		s.DB,
	}
	err := repo.PostControlAntares("ps9t5UiX15TVLxYB", "862b34fe2de548cc:cdf66d91b12db8d2", "1", "2")
	require.NoError(s.T(), err)

	// //error case http.NewRequest()
	// serverCase2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))
	// defer serverCase2.Close()

	// repo = repository{
	// 	s.DB,
	// }
	// err = repo.ManualControl("15", "lga4541000000814")
	// require.Error(s.T(), err)

	// //error case client.Do()
	// serverCase3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))
	// defer serverCase3.Close()
	// repo = repository{
	// 	s.DB,
	// }
	// err = repo.ManualControl("15", "/lga4541000000814")
	// require.Error(s.T(), err)

	// //error case ioutil.ReadAll()
	// serverCase4 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Length", "1")
	// }))
	// defer serverCase4.Close()
	// repo = repository{
	// 	s.DB,
	// 	s.mockUUID,
	// }
	// err = repo.ManualControl("15", "/lga4541000000814")
	// require.Error(s.T(), err)
}
