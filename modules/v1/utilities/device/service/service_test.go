package service

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	"GuppyTech/modules/v1/utilities/device/models"
// 	m_deviceRepository "GuppyTech/modules/v1/utilities/device/repository/mock"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestInit(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	os.Chdir("../../../../../")
// }

// func Test_GetDatafromWebhook(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	data := models.ConnectionDat{
// 		Status_device:    10,
// 		Device_mode:      2,
// 		Temperature:      0,
// 		Ph:               0,
// 		Dissolved_oxygen: 0,
// 	}
// 	antares_id := "ps9t5UiX15TVLxYB"
// 	inputJSON := `{"aeratorMode": 2, "statusDevice": 0}`
// 	output := models.Device{
// 		Device_id:  "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
// 		Antares_id: "ps9t5UiX15TVLxYB",
// 		Status_id:  "1",
// 	}

// 	tests := []struct {
// 		nameTest   string
// 		beforeTest func(service *m_deviceRepository.MockRepository)
// 		err        error
// 	}{
// 		{
// 			nameTest: "Success Post to Antares",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().GetDeviceByAntares(antares_id).Return(output, nil)
// 				deviceService.EXPECT().BindSensorData(output.Device_id, data).Return(nil, nil).AnyTimes()
// 			},
// 			err: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t *testing.T) {
// 			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

// 			w := &service{
// 				repository: mockRepository,
// 			}

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockRepository)
// 			}

// 			_, err := w.GetDatafromWebhook(inputJSON, antares_id)
// 			if err != nil {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}

// 			assert.Equal(t, test.err, err)
// 		})
// 	}
// }

// func Test_GetAllDevices(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	output := []models.Device{
// 		{
// 			Device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
// 			Antares_id:  "ps9t5UiX15TVLxYB",
// 			Device_name: "Aerator Utama",
// 		},
// 	}

// 	tests := []struct {
// 		nameTest   string
// 		beforeTest func(service *m_deviceRepository.MockRepository)
// 		result     []models.Device
// 		err        error
// 	}{
// 		{
// 			nameTest: "Success Get Devices Data",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().GetAllDevices().Return(output, nil)
// 			},
// 			result: output,
// 			err:    nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t *testing.T) {
// 			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

// 			w := &service{
// 				repository: mockRepository,
// 			}

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockRepository)
// 			}

// 			resp, err := w.GetAllDevices()
// 			if err != nil {
// 				assert.Error(t, test.err, err)
// 			} else {
// 				assert.NoError(t, test.err, err)
// 			}

// 			assert.Equal(t, test.result, resp)
// 		})
// 	}
// }

// func Test_GetDeviceHistory(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	output := []models.DeviceHistory{
// 		{
// 			History_device_name: "Aerator Utama",
// 			History_status_name: "on",
// 			History_mode_name:   "Otomatis",
// 			History_date:        time.Now(),
// 		},
// 	}

// 	tests := []struct {
// 		nameTest   string
// 		beforeTest func(service *m_deviceRepository.MockRepository)
// 		result     []models.DeviceHistory
// 		err        error
// 	}{
// 		{
// 			nameTest: "Success Get Device History",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().GetDeviceHistory().Return(output, nil)
// 			},
// 			result: output,
// 			err:    nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t *testing.T) {
// 			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

// 			w := &service{
// 				repository: mockRepository,
// 			}

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockRepository)
// 			}

// 			resp, err := w.GetDeviceHistory()
// 			if err != nil {
// 				assert.Error(t, test.err, err)
// 			} else {
// 				assert.NoError(t, test.err, err)
// 			}

// 			assert.Equal(t, test.result, resp)
// 		})
// 	}
// }

// func Test_Control(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
// 	power := "11"
// 	mode := "2"

// 	tests := []struct {
// 		nameTest   string
// 		beforeTest func(service *m_deviceRepository.MockRepository)
// 		result     error
// 	}{
// 		{
// 			nameTest: "Success Controlling Device",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().Control(id, power, mode).Return(nil)
// 			},
// 			result: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t *testing.T) {
// 			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

// 			w := &service{
// 				repository: mockRepository,
// 			}

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockRepository)
// 			}

// 			err := w.Control(id, power, mode)
// 			if err != nil {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func Test_PostControlAntares(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	antares_id := "ps9t5UiX15TVLxYB"
// 	token := "862b34fe2de548cc:cdf66d91b12db8d2"
// 	power := "0"
// 	mode := "2"

// 	tests := []struct {
// 		nameTest   string
// 		beforeTest func(service *m_deviceRepository.MockRepository)
// 		result     error
// 	}{
// 		{
// 			nameTest: "Success Post to Antares",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().PostControlAntares(antares_id, token, power, mode).Return(nil)
// 			},
// 			result: nil,
// 		},
// 		{
// 			nameTest: "Success Post to Antares",
// 			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
// 				deviceService.EXPECT().PostControlAntares(antares_id, token, power, mode).Return(nil)
// 			},
// 			result: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t *testing.T) {
// 			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

// 			w := &service{
// 				repository: mockRepository,
// 			}

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockRepository)
// 			}

// 			err := w.PostControlAntares(antares_id, token, power, mode)
// 			if err != nil {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
