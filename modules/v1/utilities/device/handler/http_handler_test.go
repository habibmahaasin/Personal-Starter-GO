package handler

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/models"
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"
	myJSON "GuppyTech/pkg/json"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	m_deviceService "GuppyTech/modules/v1/utilities/device/service/mock"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func SetUpRouter() *gin.Engine {
	app := gin.Default()
	cookieStore := cookie.NewStore([]byte("GuPPy_T3ch_5mart_A3raT0rs"))
	app.Use(sessions.Sessions("GuppyTech", cookieStore))

	return app
}

func SetupDB() *gorm.DB {
	dsn := "host=satao.db.elephantsql.com user=dwbejsql password=Kb48I9w7spTcFsiPCP2tPHeR9mhm3Ds1 dbname=dwbejsql port=5432 TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return Db
}

func SetupHandler() *deviceHandler {
	conf, _ := config.Init()
	Repository := repository.NewRepository(SetupDB(), conf)
	Service := service.NewService(Repository, myJSON.Instance())
	Handler := NewDeviceHandler(Service)
	return Handler
}

func Test_SubscribeWebhook(t *testing.T) {
	tests := []struct {
		name       string
		inputJSON  string
		statusCode int
		response   string
	}{
		{
			name: "Success",
			inputJSON: `
			{
				"m2m:sgn" : {
					"m2m:nev" : {
						"m2m:rep" : {
							"m2m:cin" : {
							"rn" : "cin_b7NDksZhTsWEmLgh",
							"ty" : 4,
							"ri" : "/antares-cse/cin-b7NDksZhTsWEmLgh",
							"pi" : "/antares-cse/cnt-ps9t5UiX15TVLxYB",
							"ct" : "20220405T160104",
							"lt" : "20220405T160104",
							"st" : 0,
							"cnf" : "text/plain:0",
							"cs" : 266,
							"con" : "{\"aeratorMode\":2,\"temperature\":28.375,\"ph\":21.97006,\"dissolvedOxygen\":50.34506,\"statusDevice\":10, \"calibration_ph1\":11.5,\"calibration_ph2\":11.5}"
							}
						},
						"m2m:rss" : 1
					},
					"m2m:sud" : false,
					"m2m:sur" : "/antares-cse/sub-HQ5_ZGw6Ts6SUnrz"
				}
			}
			`,
			statusCode: 200,
			response:   ``,
		},
		{
			name: "Format Input Tidak Sesuai",
			inputJSON: `
			{
				"m2m:cin" : {
					"cnf" : "text/plain:0",
					"cs" : 266,
				}
			}
			`,
			statusCode: 220,
			response:   `{"meta":{"message":"Error, Format Input Tidak Sesuai","code":220,"status":"error"},"data":null}`,
		},
	}
	r := SetUpRouter()
	r.POST("/api/v1/webhook", SetupHandler().SubscribeWebhook)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest("POST", "/api/v1/webhook", bytes.NewBufferString(test.inputJSON))
			if err != nil {
				t.Errorf(err.Error())
			}

			response := httptest.NewRecorder()
			r.ServeHTTP(response, request)
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf(err.Error())
			}

			assert.Equal(t, response.Code, test.statusCode)
			assert.Equal(t, string(responseData), test.response)
		})
	}
}

func Test_Control(t *testing.T) {
	ctrler := gomock.NewController(t)
	defer ctrler.Finish()

	cookies := "GuppyTech=MTY4Nzg4OTQyOXxEdi1CQkFFQ180SUFBUkFCRUFBQV81N19nZ0FEQm5OMGNtbHVad3dKQUFkMWMyVnlYMmxrQm5OMGNtbHVad3dtQUNSaE9UWXlNekl4WXkwMllqTmhMVFJpT1RJdE9HRTNNQzA1TnpJNVlURm1NVFZpTnpVR2MzUnlhVzVuREFjQUJXVnRZV2xzQm5OMGNtbHVad3dTQUJCaFpHMXBia0JuZFhCd2VTNTBaV05vQm5OMGNtbHVad3dMQUFsbWRXeHNYMjVoYldVR2MzUnlhVzVuREJFQUQwZDFjSEI1VkdWamFDQkJaRzFwYmc9PXw-fMZDrPRUEhzbOESI0OFERF_CY7HCa7iBnNZfrmK-Yg=="
	id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"
	antares_id := "ps9t5UiX15TVLxYB"
	antares_token := "f784524323f73064:4c0b580400028426"

	tests := []struct {
		name       string
		mode       string
		power      string
		page       string
		statusCode int
		beforeTest func(deviceService *m_deviceService.MockService)
	}{
		{
			name:       "Test Controlling Berhasil Dari Daftar Perangkat",
			mode:       "2",
			power:      "10",
			page:       "daftar_perangkat",
			statusCode: 302,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().Control(id, "10", "2").Return(nil)
				deviceService.EXPECT().GetDeviceById("a962321c-6b3a-4b92-8a70-9729a1f15b75", id).Return(models.Device{
					Antares_id: "ps9t5UiX15TVLxYB",
				}, nil)
				deviceService.EXPECT().PostControlAntares(antares_id, antares_token, "10", "2").Return(nil)
			},
		},
		{
			name:       "Test Controlling Berhasil Dari Detail Perangkat",
			mode:       "1",
			power:      "11",
			page:       "detail_perangkat",
			statusCode: 302,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().Control(id, "11", "1").Return(nil)
				deviceService.EXPECT().GetDeviceById("a962321c-6b3a-4b92-8a70-9729a1f15b75", id).Return(models.Device{
					Antares_id: "ps9t5UiX15TVLxYB",
				}, nil)
				deviceService.EXPECT().PostControlAntares(antares_id, antares_token, "11", "1").Return(nil)
			},
		},
		{
			name:       "Test Controlling Gagal",
			mode:       "1",
			power:      "11",
			page:       "daftar_perangkat",
			statusCode: 200,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().GetDeviceById("a962321c-6b3a-4b92-8a70-9729a1f15b75", id).Return(models.Device{
					Antares_id: "ps9t5UiX15TVLxYB",
				}, nil)
				deviceService.EXPECT().PostControlAntares(antares_id, antares_token, "11", "1").Return(nil)
				deviceService.EXPECT().Control(id, "11", "1").Return(errors.New("Error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceService := m_deviceService.NewMockService(ctrler)

			deviceHandler := &deviceHandler{
				deviceService: deviceService,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(deviceService)
			}

			handler := deviceHandler.Control
			router := SetUpRouter()

			router.GET("/control/:page/:id/:antares/:mode/:power", handler)
			req, err := http.NewRequest("GET", "/control/"+tt.page+"/"+id+"/"+antares_id+"/"+tt.mode+"/"+tt.power, nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, resp.Code, tt.statusCode)

		})
	}
}

func Test_AddDevice(t *testing.T) {
	cookies := "GuppyTech=MTY4Nzg4OTQyOXxEdi1CQkFFQ180SUFBUkFCRUFBQV81N19nZ0FEQm5OMGNtbHVad3dKQUFkMWMyVnlYMmxrQm5OMGNtbHVad3dtQUNSaE9UWXlNekl4WXkwMllqTmhMVFJpT1RJdE9HRTNNQzA1TnpJNVlURm1NVFZpTnpVR2MzUnlhVzVuREFjQUJXVnRZV2xzQm5OMGNtbHVad3dTQUJCaFpHMXBia0JuZFhCd2VTNTBaV05vQm5OMGNtbHVad3dMQUFsbWRXeHNYMjVoYldVR2MzUnlhVzVuREJFQUQwZDFjSEI1VkdWamFDQkJaRzFwYmc9PXw-fMZDrPRUEhzbOESI0OFERF_CY7HCa7iBnNZfrmK-Yg=="
	ctrler := gomock.NewController(t)
	defer ctrler.Finish()

	user_id := "a962321c-6b3a-4b92-8a70-9729a1f15b75"
	add_device_form := models.DeviceInput{
		Device_name:     "Device 1",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Bandung",
		Latitude:        "123",
		Longitude:       "123",
		Brand_id:        "1",
		Mode_id:         "1",
	}

	tests := []struct {
		name       string
		payload    bool
		statusCode int
		beforeTest func(deviceService *m_deviceService.MockService)
	}{
		{
			name:       "Test Add Device Berhasil",
			payload:    true,
			statusCode: 302,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().AddDevice(models.DeviceInput{
					Device_name:     "Device 1",
					Antares_id:      "ps9t5UiX15TVLxYB",
					Device_location: "Bandung",
					Latitude:        "123",
					Longitude:       "123",
					Brand_id:        "1",
					Mode_id:         "1",
				}, user_id).Return(nil)
			},
		},
		{
			name:       "Test Add Device Gagal",
			payload:    false,
			statusCode: 200,
			beforeTest: func(deviceService *m_deviceService.MockService) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceService := m_deviceService.NewMockService(ctrler)

			deviceHandler := &deviceHandler{
				deviceService: deviceService,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(deviceService)
			}

			router := SetUpRouter()
			router.POST("/add-device", deviceHandler.AddDevice)

			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err := writer.WriteField("device_name", add_device_form.Device_name)
			assert.NoError(t, err)
			err = writer.WriteField("antares_id", add_device_form.Antares_id)
			assert.NoError(t, err)
			err = writer.WriteField("device_location", add_device_form.Device_location)
			assert.NoError(t, err)
			err = writer.WriteField("latitude", add_device_form.Latitude)
			assert.NoError(t, err)
			err = writer.WriteField("longitude", add_device_form.Longitude)
			assert.NoError(t, err)
			err = writer.WriteField("brand_id", add_device_form.Brand_id)
			assert.NoError(t, err)
			err = writer.WriteField("mode_id", add_device_form.Mode_id)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			if tt.payload {
				req, err := http.NewRequest("POST", "/add-device", payload)
				assert.NoError(t, err)

				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Cookie", cookies)

				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				assert.Equal(t, resp.Code, tt.statusCode)
			} else {
				req, err := http.NewRequest("POST", "/add-device", nil)
				assert.NoError(t, err)

				req.Header.Set("Content-Type", writer.FormDataContentType())

				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				assert.Equal(t, resp.Code, tt.statusCode)
			}

		})
	}
}

func Test_DeleteDevice(t *testing.T) {
	ctrler := gomock.NewController(t)
	defer ctrler.Finish()

	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	tests := []struct {
		name       string
		statusCode int
		beforeTest func(deviceService *m_deviceService.MockService)
	}{
		{
			name:       "Test Delete Device Berhasil",
			statusCode: 302,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().DeleteDevice(device_id).Return(nil)
			},
		},
		{
			name:       "Test Delete Device Gagal",
			statusCode: 200,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().DeleteDevice(device_id).Return(errors.New("Error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceService := m_deviceService.NewMockService(ctrler)

			deviceHandler := &deviceHandler{
				deviceService: deviceService,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(deviceService)
			}

			router := SetUpRouter()

			router.GET("/delete-device/:id", deviceHandler.DeleteDevice)
			req, err := http.NewRequest("GET", "/delete-device/"+device_id, nil)
			assert.NoError(t, err)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, resp.Code, tt.statusCode)

		})
	}
}

func Test_EditDevice(t *testing.T) {
	ctrler := gomock.NewController(t)
	defer ctrler.Finish()

	cookies := "GuppyTech=MTY4Nzg4OTQyOXxEdi1CQkFFQ180SUFBUkFCRUFBQV81N19nZ0FEQm5OMGNtbHVad3dKQUFkMWMyVnlYMmxrQm5OMGNtbHVad3dtQUNSaE9UWXlNekl4WXkwMllqTmhMVFJpT1RJdE9HRTNNQzA1TnpJNVlURm1NVFZpTnpVR2MzUnlhVzVuREFjQUJXVnRZV2xzQm5OMGNtbHVad3dTQUJCaFpHMXBia0JuZFhCd2VTNTBaV05vQm5OMGNtbHVad3dMQUFsbWRXeHNYMjVoYldVR2MzUnlhVzVuREJFQUQwZDFjSEI1VkdWamFDQkJaRzFwYmc9PXw-fMZDrPRUEhzbOESI0OFERF_CY7HCa7iBnNZfrmK-Yg=="
	device_id := "e5d415f7-a96b-4dc2-84b8-64a1830b4c01"

	edit_device_form := models.DeviceInput{
		Device_name:     "Device 1",
		Antares_id:      "ps9t5UiX15TVLxYB",
		Device_location: "Bandung",
		Latitude:        "123",
		Longitude:       "123",
		Brand_id:        "1",
		Mode_id:         "1",
	}

	tests := []struct {
		name       string
		payload    bool
		statusCode int
		beforeTest func(deviceService *m_deviceService.MockService)
	}{
		{
			name:       "Test Edit Device Berhasil",
			payload:    true,
			statusCode: 302,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().UpdateDeviceById(models.DeviceInput{
					Device_name:     "Device 1",
					Antares_id:      "ps9t5UiX15TVLxYB",
					Device_location: "Bandung",
					Latitude:        "123",
					Longitude:       "123",
					Brand_id:        "1",
					Mode_id:         "1",
				}, device_id).Return(nil)
			},
		},
		{
			name:       "Test Edit Device Berhasil",
			payload:    false,
			statusCode: 200,
			beforeTest: func(deviceService *m_deviceService.MockService) {
				deviceService.EXPECT().UpdateDeviceById(models.DeviceInput{}, device_id).Return(errors.New("Error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceService := m_deviceService.NewMockService(ctrler)

			deviceHandler := &deviceHandler{
				deviceService: deviceService,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(deviceService)
			}

			router := SetUpRouter()
			router.POST("/edit-device/:id", deviceHandler.EditDevice)

			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err := writer.WriteField("device_name", edit_device_form.Device_name)
			assert.NoError(t, err)
			err = writer.WriteField("antares_id", edit_device_form.Antares_id)
			assert.NoError(t, err)
			err = writer.WriteField("device_location", edit_device_form.Device_location)
			assert.NoError(t, err)
			err = writer.WriteField("latitude", edit_device_form.Latitude)
			assert.NoError(t, err)
			err = writer.WriteField("longitude", edit_device_form.Longitude)
			assert.NoError(t, err)
			err = writer.WriteField("brand_id", edit_device_form.Brand_id)
			assert.NoError(t, err)
			err = writer.WriteField("mode_id", edit_device_form.Mode_id)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			if tt.payload {
				req, err := http.NewRequest("POST", "/edit-device/"+device_id, payload)
				assert.NoError(t, err)

				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Cookie", cookies)

				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				assert.Equal(t, resp.Code, tt.statusCode)

			} else {
				req, err := http.NewRequest("POST", "/edit-device/"+device_id, nil)
				assert.NoError(t, err)

				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Cookie", cookies)

				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				assert.Equal(t, resp.Code, tt.statusCode)
			}

		})
	}

}
