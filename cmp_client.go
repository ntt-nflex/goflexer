package goflexer

import (
	"crypto/tls"

	log "github.com/sirupsen/logrus"
	"github.com/go-resty/resty"
)

// CmpClient type provides an interface to the CMP APIs
type CmpClient struct {
	client *resty.Client
	url    string
}

// NewCmpClient method creates a new Ccmp Client object
func NewCmpClient(conf *Config) *CmpClient {

	restyClient := resty.New()

	c := CmpClient{
		url:    conf.URL,
		client: restyClient,
	}

	c.client.SetHostURL(conf.URL)

	if conf.AccessToken != "" {
		c.client.SetHeader("Cookie", conf.AccessToken)
	}
	if conf.Username != "" && conf.Password != "" {
		c.client.SetBasicAuth(conf.Username, conf.Password)
	}

	c.client.SetHeader("User-Agent", "goflexer-cmp-client")
	c.client.SetHeader("Content-Type", "application/json")

	if !conf.VerifySSL {
		log.Info("Skip Verify SSL")
		c.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	return &c
}

// R method returns a new resty request, to expose full resty functionality
func (c *CmpClient) R() *resty.Request {
	return c.client.R()
}

// Get method provides a simple way to make a get request with query parameters
func (c *CmpClient) Get(url string, params map[string]string) (*resty.Response, error) {
	return c.R().SetQueryParams(params).Get(url)
}
