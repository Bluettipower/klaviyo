package klaviyo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	endpoint       = "https://a.klaviyo.com/api/"
	defaultVersion = "2024-02-15"
)

type Client struct {
	APIKey    string
	APISecret string
	Version   string
	Profile   ProfileService
	Event     EventService
	Metric    MetricService
	Tag       TagService
	Coupon    CouponService
}

type Option func(c *Client)

func WithVersion(version string) Option {
	return func(c *Client) {
		c.Version = version
	}
}

func NewClient(apiKey string, opts ...Option) *Client {

	c := &Client{
		APIKey: apiKey,
	}

	c.Profile = &ProfileServiceOp{client: c}
	c.Event = &EventServiceOp{client: c}
	c.Metric = &MetricServiceOp{client: c}
	c.Tag = &TagServiceOp{client: c}
	c.Coupon = &CouponServiceOp{client: c}

	for _, opt := range opts {
		opt(c)
	}

	if c.Version == "" {
		c.Version = defaultVersion
	}

	return c

}

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {

	var bodyReader io.Reader
	if body != nil {
		requestJson, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}

		bodyReader = bytes.NewBuffer(requestJson)
	}

	url = fmt.Sprintf("%v%v", endpoint, url)

	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}

	// Content Type
	httpReq.Header.Add("accept", "application/json")
	httpReq.Header.Add("revision", c.Version)
	httpReq.Header.Add("content-type", "application/json")
	httpReq.Header.Add("Authorization", "Klaviyo-API-Key "+c.APIKey)

	//Client
	client := &http.Client{}
	resp, errDo := client.Do(httpReq)

	// fmt.Println(resp)
	// fmt.Println()
	// fmt.Println(errDo)

	if resp != nil {
		defer resp.Body.Close()
	}
	if errDo != nil {
		return errDo
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		errDecode := decoder.Decode(&v)
		if errDecode != nil {
			return errDecode
		}
	}
	return nil
}
