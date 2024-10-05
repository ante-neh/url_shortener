package types

import "time"

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"custom_short"`
	Expire      time.Duration `json:"expire"`
}

type response struct {
	URL            string `json:"url"`
	CustomShort    string `json:"custom_short"`
	Expire         time.Duration `json:"expire"`
	XRateRemaining int `json:"x_rate_remaining"`
	XRateReset     time.Duration `json:"x_rate_reset"`
}