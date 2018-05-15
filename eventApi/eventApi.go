package eventApi

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
)

type EventApi struct{}

func NewEventApi(envReader *goenv.OsEnvReader) (*EventApi, error) {
	return &EventApi{}, nil
}

func (api *EventApi) NewEvent(c echo.Context) error {
	event := &event{}
	c.Bind(event)
	//save the event here
	return nil
}
