package vendorApi

import (
	"github.com/segmentio/ksuid"
)

type vendor struct {
	Name     string      `json:"name"`
	LogoUrl  string      `json:"logo_url"`
	Address  Location    `json:"address"`
	Contact  ContactInfo `json:"contact_info"`
	ZipCode  uint32      `json:"zip_code"`
	VendorId ksuid.KSUID `json:"vendor_id"`
}

func (this *vendor) Save() error {
	//Validate
	//Generate VendorId
	//Save to database
	return nil
}
