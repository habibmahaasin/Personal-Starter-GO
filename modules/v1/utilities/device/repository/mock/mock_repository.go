package mock_repository

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Mock BindSensorData
func (mr *MockRepositoryMockRecorder) BindSensorData(Device_id string, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindSensorData", reflect.TypeOf((*MockRepository)(nil).BindSensorData), Device_id, input)
}

func (m *MockRepository) BindSensorData(Device_id string, input models.ConnectionDat) (error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindSensorData", Device_id, input)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mock GetAllDevices
func (mr *MockRepositoryMockRecorder) GetAllDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevices", reflect.TypeOf((*MockRepository)(nil).GetAllDevices))
}

func (m *MockRepository) GetAllDevices() ([]models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevices")
	ret0, _ := ret[0].([]models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mock GetDeviceByAntares
func (mr *MockRepositoryMockRecorder) GetDeviceByAntares(antaresDeviceID string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByAntares", reflect.TypeOf((*MockRepository)(nil).GetDeviceByAntares), antaresDeviceID)
}

func (m *MockRepository) GetDeviceByAntares(antaresDeviceID string) (models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByAntares", antaresDeviceID)
	ret0, _ := ret[0].(models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mock GetDeviceByAntares
func (mr *MockRepositoryMockRecorder) GetDeviceHistory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceHistory", reflect.TypeOf((*MockRepository)(nil).GetDeviceHistory))
}

func (m *MockRepository) GetDeviceHistory() ([]models.DeviceHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceHistory")
	ret0, _ := ret[0].([]models.DeviceHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mock Control
func (mr *MockRepositoryMockRecorder) Control(id string, power string, mode string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Control", reflect.TypeOf((*MockRepository)(nil).Control), id, power, mode)
}

func (m *MockRepository) Control(id string, power string, mode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Control", id, power, mode)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mock PostControlAntares
func (mr *MockRepositoryMockRecorder) PostControlAntares(antares_id string, token string, power string, mode string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostControlAntares", reflect.TypeOf((*MockRepository)(nil).PostControlAntares), antares_id, token, power, mode)
}

func (m *MockRepository) PostControlAntares(antares_id string, token string, power string, mode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostControlAntares", antares_id, token, power, mode)
	ret0, _ := ret[0].(error)
	return ret0
}
