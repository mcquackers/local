package main

import (
  "time"
)

type AppMeta struct {
	BuildTime time.Time `env:"BUILD_TIME" json:"build-time"`
	Version   string    `env:"VERSION" json:"version"`
}

type HealthCheckResponse struct {
	Meta   AppMeta `json:"meta"`
	Status string  `json:"status"`
}
