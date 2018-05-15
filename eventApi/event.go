package eventApi

import (
	"github.com/segmentio/ksuid"
	"time"
)

type event struct {
	Name            string      `json:"name"`
	StartTime       time.Time   `json:"start_time"`
	EndTime         time.Time   `json:"end_time"`
	VendorId        ksuid.KSUID `json:"vendor_id"`
	EventId         ksuid.KSUID `json:"event_id"`
	ZipCode         uint32      `json:"zip_code"`
	LandingPage     string      `json:"landing_page"`
	DisplayLogoUrl  string      `json:"display_logo_url"`
	BackgroundImage interface{} `json:"background_image"`
}
