package eventApi

import (
	"time"

	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
)

type EventApi struct{}

type event struct {
	Name            string      `json:"name"`
	StartTime       time.Time   `json:"start_time"`
	EndTime         time.Time   `json:"end_time"`
	VendorId        uint64      `json:"vendor_id"`
	ZipCode         uint32      `json:"zip_code"`
	LandingPage     string      `json:"landing_page"`
	DisplayLogoUrl  string      `json:"display_logo_url"`
	BackgroundImage interface{} `json:"background_image"`
}

func NewEventApi(envReader *goenv.OsEnvReader) (*EventApi, error) {
	return &EventApi{}, nil
}

func (api *EventApi) NewEvent(c echo.Context) error {
	event := &event{}
	c.Bind(event)
	//save the event here
	return nil
}
