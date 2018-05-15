package vendorApi

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
)

type VendorApi struct {
}

type Location struct {
	StreetNumber int    `json:"street_number"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type ContactInfo struct {
	PhoneNumber1 string `json:"phone_1"`
	PhoneNumber2 string `json:"phone_2"`
	Email        string `json:"email"`
}

func NewVendorApi(envReader *goenv.OsEnvReader) (*VendorApi, error) {
	return &VendorApi{}, nil
}

func (api *VendorApi) SignUpVendor(c echo.Context) error {
	newVendor := &vendor{}
	c.Bind(newVendor)
	if err := newVendor.Save(); err != nil {
		return c.JSON(500, &struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
	return c.JSON(200, &struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
