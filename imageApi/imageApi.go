package imageApi

import (
	"github.com/evilwire/go-env"
)

type ImageApi struct{}

func NewImageApi(envReader *goenv.OsEnvReader) (*ImageApi, error) {
	return &ImageApi{}, nil
}
