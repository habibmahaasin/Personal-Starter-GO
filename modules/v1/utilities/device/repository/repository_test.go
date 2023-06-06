package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
