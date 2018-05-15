package imageApi

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
	"github.com/segmentio/ksuid"
)

type ImageApi struct{}

type imageApiResponse struct {
	Status string `json:"status"`
}

type newImageRequest struct {
	Url      string      `json:"url"`
	VendorId ksuid.KSUID `json:"vendor_id"`
}

func NewImageApi(envReader *goenv.OsEnvReader) (*ImageApi, error) {
	return &ImageApi{}, nil
}

func (api *ImageApi) NewImage(c echo.Context) error {
	imageRequest := &newImageRequest{}
	c.Bind(imageRequest)

	if err := imageRequest.Save(); err != nil {
		return c.JSON(500, &imageApiResponse{
			Status: err.Error(),
		})
	}

	return c.JSON(201, &imageApiResponse{
		Status: "ok",
	})
}

func (req *newImageRequest) Save() error {
	//if vendor image directory does not exist, create it
	//pull down image
	//Save it to directory
	return nil
}
