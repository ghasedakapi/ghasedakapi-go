// Package ghasedakapi is a library for interacting with http://www.ghasedakapi.com/ API.
package ghasedakapi

import(
	"net/http"
)

const(
	baseUrl="http://ghasedakapi.com"
)
type ghasedakapi struct {
	ApiKey string
	BaseUrl    string
	HTTPClient *http.Client
}

// Create a new ghasedakapi struct.
func NewClient(apiKey string) *ghasedakapi {
	return NewHttpClient(apiKey, nil)
}

// Create a new ghasedakapi client, using a http.Client
func NewHttpClient(apiKey string, HTTPClient *http.Client) *ghasedakapi {
	if HTTPClient == nil {
		HTTPClient = http.DefaultClient
	}
	return &ghasedakapi{apiKey,baseUrl, HTTPClient}
}
