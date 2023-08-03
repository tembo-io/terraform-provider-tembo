package provider

import (
	"net/http"
	"time"
)

// HostURL - Default Tembo API URL.
const HostURL string = "http://localhost:19090"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

type AuthStruct struct {
	AccessToken string `json:"access_token"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func NewClient(host, access_token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if access_token == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		AccessToken: *access_token,
	}

	return &c, nil
}
