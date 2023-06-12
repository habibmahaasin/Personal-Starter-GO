package view

import (
	"GuppyTech/app/config"
	database "GuppyTech/app/databases"
	"GuppyTech/pkg/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestInit(t *testing.T) {
	os.Chdir("../../../../../")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	cookieStore := cookie.NewStore([]byte("UNITTESTVIEW"))
	router.Use(sessions.Sessions("GuppyTech", cookieStore))
	router.HTMLRender = html.Render("./public/templates")

	router.Static("/assets", "./public/assets")
	router.Static("/img", "./public/assets/img")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/vendor", "./public/assets/vendor")
	return router
}

func Test_Index(t *testing.T) {
	// os.Chdir("../../../../../")
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}

	r := SetUpRouter()
	db := database.Init(conf)
	deviceViewV1 := View(db, conf)
	r.GET("/", deviceViewV1.Index)

	t.Run("Berhasil menampilkan halaman Index", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)

		c := strings.Contains(s, "Beranda")
		assert.Equal(t, c, true)
	})
}

func Test_ListDevice(t *testing.T) {
	//os.Chdir("../../../../../")
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}

	r := SetUpRouter()
	db := database.Init(conf)
	deviceViewV1 := View(db, conf)
	r.GET("/daftar-perangkat", deviceViewV1.ListDevice)

	t.Run("Berhasil menampilkan halaman Daftar Perangkat", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/daftar-perangkat", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)

		c := strings.Contains(s, "Daftar Perangkat")
		assert.Equal(t, c, true)
	})
}

func Test_Report(t *testing.T) {
	//os.Chdir("../../../../../")
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}

	r := SetUpRouter()
	db := database.Init(conf)
	deviceViewV1 := View(db, conf)
	r.GET("/laporan", deviceViewV1.Report)

	t.Run("Berhasil menampilkan halaman Laporan", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/laporan", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)

		c := strings.Contains(s, "Laporan")
		assert.Equal(t, c, true)
	})
}
