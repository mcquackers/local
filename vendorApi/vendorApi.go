package vendorApi

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
)

type VendorApi struct {
}

type vendor struct {
	Name    string   `json:"name"`
	LogoUrl string   `json:"logo_url"`
	Address Location `json:"address"`
	Contact
	ZipCode uint32 `json:"zip_code"`
}

type Location struct {
	StreetNumber int    `json:"street_number"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
}

func NewVendorApi(envReader *goenv.OsEnvReader) (*VendorApi, error) {
	return &VendorApi{}, nil
}

func (api *VendorApi) SignUpVendor(c echo.Context) error {
	return c.JSON(200, &struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
