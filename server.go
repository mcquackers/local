package main

import (
	"github.com/evilwire/go-env"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mcquackers/local/checkinApi"
	"github.com/mcquackers/local/eventApi"
	"github.com/mcquackers/local/imageApi"
	"github.com/mcquackers/local/userApi"
	"github.com/mcquackers/local/vendorApi"
)

type App struct {
	UserApi    *userApi.UserApi
	VendorApi  *vendorApi.VendorApi
	EventApi   *eventApi.EventApi
	CheckinApi *checkinApi.CheckinApi
	ImageApi   *imageApi.ImageApi
	Webserver  *echo.Echo
}

func main() {
	app := mustSetupApp()
}

func mustSetupApp() *App {
	envReader := goenv.NewOsEnvReader()
	app := &App{}
	app.mustSetupApis(envReader)
	app.mustSetupWebserver(envReader)

	return app
}

func (app *App) mustSetupApis(envReader *goenv.OsEnvReader) {
	vendorApi, err := vendorApi.NewVendorApi(envReader)
	if err != nil {
		panic(err)
	}
	app.VendorApi = vendorApi

	userApi, err := userApi.NewUserApi(envReader)
	if err != nil {
		panic(err)
	}
	app.UserApi = userApi

	eventApi, err := eventApi.NewEventApi(envReader)
	if err != nil {
		panic(err)
	}
	app.EventApi = eventApi

	checkinApi, err := checkinApi.NewCheckinApi(envReader)
	if err != nil {
		panic(err)
	}
	app.CheckinApi = checkinApi

	imageApi, err := imageApi.NewImageApi(envReader)
	if err != nil {
		panic(err)
	}
	app.ImageApi = imageApi
}

func (app *App) mustSetupWebserver(envReader *goenv.OsEnvReader) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	//HealthCheck
	e.GET("/healthcheck", healthCheck)

	//End-User REST
	e.POST("/user/new", app.UserApi.SignUpUser)

	//Vendor REST
	e.POST("/vendor/new", app.VendorApi.SignUpVendor)

	//Event REST
	e.POST("/event/new", app.EventApi.NewEvent)
}

func healthCheck(c echo.Context) error {
	meta := AppMeta{}
	marshaler := goenv.DefaultEnvMarshaler{goenv.NewOsEnvReader()}
	err := marshaler.Unmarshal(&meta)
	if err != nil {
		glog.Errorf("Could not parse meta: %v", err)
		return c.JSON(500, HealthCheckResponse{
			Meta: AppMeta{
				Version: "unknown",
			},
			Status: "error",
		})
	}

	return c.JSON(200, HealthCheckResponse{
		Meta:   meta,
		Status: "ok",
	})
}
