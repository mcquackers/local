package checkinApi

import (
	"github.com/evilwire/go-env"
)

type CheckinApi struct{}

func NewCheckinApi(envReader *goenv.OsEnvReader) (*CheckinApi, error) {
	return &CheckinApi{}, nil
}
