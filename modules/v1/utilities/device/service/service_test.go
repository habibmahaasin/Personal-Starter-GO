package service

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"GuppyTech/modules/v1/utilities/device/models"
	m_deviceRepository "GuppyTech/modules/v1/utilities/device/repository/mock"
	m_json "GuppyTech/pkg/json/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func Test_GetDatafromWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	antares_id := "ps9t5UiX15TVLxYB"

	output := models.Device{
		Device_id:  "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
		Antares_id: "ps9t5UiX15TVLxYB",
		Status_id:  "1",
	}

	tests := []struct {
		nameTest   string
		beforeTest func(service *m_deviceRepository.MockRepository)
		inputJson  string
		err        error
	}{
		{
			nameTest: "Success Post to Antares With Status Device 1",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().GetDeviceByAntares(antares_id).Return(output, nil)
				deviceService.EXPECT().BindSensorData(output.Device_id, models.ConnectionDat{
					Aerator_status:   11,
					Aerator_mode:     2,
					Temperature:      0,
					Ph:               0,
					Dissolved_oxygen: 0,
				}).Return(nil, nil).AnyTimes()
			},
			inputJson: `{"aeratorMode": 2, "statusDevice": 1}`,
			err:       nil,
		},
		{
			nameTest: "Success Post to Antares With Status Device 0",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().GetDeviceByAntares(antares_id).Return(output, nil)
				deviceService.EXPECT().BindSensorData(output.Device_id, models.ConnectionDat{
					Aerator_status:   10,
					Aerator_mode:     1,
					Temperature:      0,
					Ph:               0,
					Dissolved_oxygen: 0,
				}).Return(nil, nil).AnyTimes()
			},
			inputJson: `{"aeratorMode": 1, "statusDevice": 0}`,
			err:       nil,
		},
		{
			nameTest: "Success Post to Antares With Status Device Other Condition",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().GetDeviceByAntares(antares_id).Return(output, nil)
				deviceService.EXPECT().BindSensorData(output.Device_id, models.ConnectionDat{
					Aerator_status:   10,
					Aerator_mode:     1,
					Temperature:      0,
					Ph:               0,
					Dissolved_oxygen: 0,
				}).Return(nil, nil).AnyTimes()
			},
			inputJson: `{"aeratorMode": 1, "statusDevice": 15}`,
			err:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			_, err := w.GetDatafromWebhook(test.inputJson, antares_id)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.err, err)
		})
	}
}

func Test_GetDeviceHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"

	output := []models.DeviceHistory{
		{
			History_device_name: "Aerator Utama",
			History_status_name: "on",
			History_mode_name:   "Otomatis",
			History_date:        time.Now(),
		},
	}

	outputJson, _ := json.Marshal(output)

	tests := []struct {
		nameTest   string
		beforeTest func(deviceRepository *m_deviceRepository.MockRepository, m_json *m_json.MockJSON)
		result     []models.DeviceHistory
		resultjson string
		err        error
	}{
		{
			nameTest: "Success Get Device History",
			beforeTest: func(deviceRepository *m_deviceRepository.MockRepository, m_json *m_json.MockJSON) {
				deviceRepository.EXPECT().GetDeviceHistory(user_id).Return(output, nil)
				m_json.EXPECT().Marshal(output).Return(outputJson, nil)
			},
			result:     output,
			resultjson: string(outputJson),
			err:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)
			my_json := m_json.NewMockJSON(ctrl)

			w := &service{
				repository: mockRepository,
				json:       my_json,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository, my_json)
			}

			resp, jsonresp, err := w.GetDeviceHistory(user_id)
			if err != nil {
				assert.Error(t, test.err, err)
			} else {
				assert.NoError(t, test.err, err)
			}

			assert.Equal(t, test.result, resp)
			assert.Equal(t, test.resultjson, jsonresp)
		})
	}
}

func Test_GetAllDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"

	output := []models.Device{
		{
			Device_id:   "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			Antares_id:  "ps9t5UiX15TVLxYB",
			Device_name: "Aerator Utama",
		},
	}

	tests := []struct {
		nameTest   string
		beforeTest func(service *m_deviceRepository.MockRepository)
		result     []models.Device
		err        error
	}{
		{
			nameTest: "Success Get Devices Data",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().GetAllDevices(user_id).Return(output, nil)
			},
			result: output,
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			resp, err := w.GetAllDevices(user_id)
			if err != nil {
				assert.Error(t, test.err, err)
			} else {
				assert.NoError(t, test.err, err)
			}

			assert.Equal(t, test.result, resp)
		})
	}
}

func Test_Control(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	tests := []struct {
		nameTest   string
		beforeTest func(service *m_deviceRepository.MockRepository)
		power      string
		mode       string
		result     error
	}{
		{
			nameTest: "Success Send New Condition",
			power:    "11",
			mode:     "2",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().Control(device_id, "11", "2").Return(nil)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			err := w.Control(device_id, test.power, test.mode)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_PostControlAntares(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	antares_id := "ps9t5UiX15TVLxYB"
	token := "862b34fe2de548cc:cdf66d91b12db8d2"

	tests := []struct {
		nameTest   string
		beforeTest func(service *m_deviceRepository.MockRepository)
		power      string
		mode       string
		result     error
	}{
		{
			nameTest: "Success Post to Antares",
			power:    "11",
			mode:     "2",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().PostControlAntares(antares_id, token, "1", "2").Return(nil)
			},
			result: nil,
		},
		{
			nameTest: "Success Post to Antares",
			power:    "10",
			mode:     "1",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().PostControlAntares(antares_id, token, "0", "1").Return(nil)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			err := w.PostControlAntares(antares_id, token, test.power, test.mode)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_AddDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	input_device := models.DeviceInput{
		Device_name:     "Aerator Utama",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Jl. Raya Bogor KM 30",
		Latitude:        "123",
		Longitude:       "123",
		Brand_id:        "1",
		Mode_id:         "1",
	}

	tests := []struct {
		nameTest    string
		beforeTest  func(service *m_deviceRepository.MockRepository)
		value_input bool
		result      error
	}{
		{
			nameTest:    "Success Add New Device",
			value_input: true,
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().AddDevice(input_device, user_id).Return(nil)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			err := w.AddDevice(input_device, user_id)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetDeviceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	output := models.Device{
		Device_id:              "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
		Antares_id:             "ps9t5UiX15TVLxYB",
		Device_name:            "Aerator Utama",
		Date_updated_formatter: "Mon Jan  1 00:00:00 0001",
	}

	tests := []struct {
		nameTest   string
		beforeTest func(service *m_deviceRepository.MockRepository)
		result     models.Device
		err        error
	}{
		{
			nameTest: "Success Get Devices Data By Id",
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().GetDeviceById(user_id, device_id).Return(output, nil)
			},
			result: output,
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			resp, err := w.GetDeviceById(user_id, device_id)
			if err != nil {
				assert.Error(t, test.err, err)
			} else {
				assert.NoError(t, test.err, err)
			}

			assert.Equal(t, test.result, resp)
		})
	}
}

func Test_GetDeviceHistoryById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"

	output := []models.DeviceHistory{
		{
			History_device_name: "Aerator Utama",
			History_status_name: "on",
			History_mode_name:   "Otomatis",
			History_date:        time.Now(),
		},
	}

	outputJson, _ := json.Marshal(output)

	tests := []struct {
		nameTest   string
		beforeTest func(deviceRepository *m_deviceRepository.MockRepository, m_json *m_json.MockJSON)
		result     []models.DeviceHistory
		resultjson string
		err        error
	}{
		{
			nameTest: "Success Get Device History By Id",
			beforeTest: func(deviceRepository *m_deviceRepository.MockRepository, m_json *m_json.MockJSON) {
				deviceRepository.EXPECT().GetDeviceHistoryById(device_id, user_id).Return(output, nil)
				m_json.EXPECT().Marshal(output).Return(outputJson, nil)
			},
			result:     output,
			resultjson: string(outputJson),
			err:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)
			my_json := m_json.NewMockJSON(ctrl)

			w := &service{
				repository: mockRepository,
				json:       my_json,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository, my_json)
			}

			resp, jsonresp, err := w.GetDeviceHistoryById(device_id, user_id)
			if err != nil {
				assert.Error(t, test.err, err)
			} else {
				assert.NoError(t, test.err, err)
			}

			assert.Equal(t, test.result, resp)
			assert.Equal(t, test.resultjson, jsonresp)
		})
	}
}

func Test_DeleteDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	tests := []struct {
		nameTest    string
		beforeTest  func(service *m_deviceRepository.MockRepository)
		value_input bool
		result      error
	}{
		{
			nameTest:    "Success Delete Device",
			value_input: true,
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().DeleteDevice(device_id).Return(nil)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			err := w.DeleteDevice(device_id)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetDeviceBrands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	output := []models.Device{
		{
			Brand_id:   1,
			Brand_name: "Amara AA-666",
		},
	}

	tests := []struct {
		nameTest   string
		beforeTest func(deviceRepository *m_deviceRepository.MockRepository)
		result     []models.Device
		err        error
	}{
		{
			nameTest: "Success Get Device History By Id",
			beforeTest: func(deviceRepository *m_deviceRepository.MockRepository) {
				deviceRepository.EXPECT().GetDeviceBrands().Return(output, nil)
			},
			result: output,
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			resp, err := w.GetDeviceBrands()
			if err != nil {
				assert.Error(t, test.err, err)
			} else {
				assert.NoError(t, test.err, err)
			}

			assert.Equal(t, test.result, resp)
		})
	}
}

func Test_UpdateDeviceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	input_device := models.DeviceInput{
		Device_name:     "Aerator Utama",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Jl. Raya Bogor KM 30",
		Latitude:        "123",
		Longitude:       "123",
		Brand_id:        "1",
		Mode_id:         "1",
	}

	tests := []struct {
		nameTest    string
		beforeTest  func(service *m_deviceRepository.MockRepository)
		value_input bool
		result      error
	}{
		{
			nameTest:    "Success Add New Device",
			value_input: true,
			beforeTest: func(deviceService *m_deviceRepository.MockRepository) {
				deviceService.EXPECT().UpdateDeviceById(input_device, device_id).Return(nil)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			mockRepository := m_deviceRepository.NewMockRepository(ctrl)

			w := &service{
				repository: mockRepository,
			}

			if test.beforeTest != nil {
				test.beforeTest(mockRepository)
			}

			err := w.UpdateDeviceById(input_device, device_id)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
