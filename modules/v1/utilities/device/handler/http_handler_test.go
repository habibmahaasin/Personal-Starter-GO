package handler

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"
	myJSON "GuppyTech/pkg/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpRouter() *gin.Engine {
	app := gin.Default()
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
							"con" : "{\"aeratorMode\":2,\"temperature\":28.375,\"ph\":21.97006,\"dissolvedOxygen\":50.34506,\"statusDevice\":10}"
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
			response:   `{"meta":{"message":"Success","code":200,"status":"success"},"data":{"aeratorMode":2,"statusDevice":10,"temperature":28.375,"ph":21.97006,"dissolvedOxygen":50.34506,"Device_id":""}}`,
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
	cookie := "message=; GuppyTech=MTY4NjE1MDIwNHxEdi1CQkFFQ180SUFBUkFCRUFBQWNmLUNBQUlHYzNSeWFXNW5EQWdBQm5WelpYSkpSQVp6ZEhKcGJtY01KZ0FrWVRrMk1qTXlNV010Tm1JellTMDBZamt5TFRoaE56QXRPVGN5T1dFeFpqRTFZamMxQm5OMGNtbHVad3dLQUFoMWMyVnlUbUZ0WlFaemRISnBibWNNRVFBUFIzVndjSGxVWldOb0lFRmtiV2x1fHVmHYqw7_906PoNHN; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYxNTExMDQsImZ1bGxfbmFtZSI6Ikd1cHB5VGVjaCBBZG1pbiIsInJvbGVfaWQiOjEsInVzZXJfaWQiOiJhOTYyMzIxYy02YjNhLTRiOTItOGE3MC05NzI5YTFmMTViNzUifQ.yA7MmYzBUjmXRb0m2ftK-WwlQ7CQbNMahok7fZF3TPA"
	basic := "Basic YWRtaW46YWRtaW4="

	tests := []struct {
		name       string
		id         string
		mode       string
		antares_id string
		power      string
		token      string
		statusCode int
	}{
		{
			name:       "Test Controlling Berhasil",
			id:         "e5d415f7-a96b-4dc2-84b8-64a1830b4c01",
			mode:       "2",
			antares_id: "ps9t5UiX15TVLxYB",
			power:      "10",
			token:      "862b34fe2de548cc:cdf66d91b12db8d2",
			statusCode: 302,
		},
		{
			name:       "Test Controlling Gagal",
			id:         "2",
			mode:       "100o",
			antares_id: "ps9t5UiX15TVLxYB",
			power:      "on",
			token:      "862b34fe2de548cc:cdf66d91b12db8d2",
			statusCode: 302,
		},
	}

	r := SetUpRouter()
	r.GET("/control/:id/:antares/:mode/:power", SetupHandler().Control)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/control/"+tt.id+"/"+tt.antares_id+"/"+tt.mode+"/"+tt.power, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Authorization", basic)
			req.Header.Set("Cookie", cookie)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)
			location, err := resp.Result().Location()
			if err != nil {
				fmt.Println(err)
				return
			}
			assert.Equal(t, resp.Code, tt.statusCode)
			assert.Equal(t, location.Path, "/daftar-perangkat")
		})
	}
}
