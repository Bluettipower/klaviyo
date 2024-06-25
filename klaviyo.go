package klaviyo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	endpoint       = "https://a.klaviyo.com/"
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
	Client    ClientService
	List      ListService
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
	c.Client = &ClientServiceOp{client: c}
	c.List = &ListServiceOp{client: c}

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
	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		decoder := json.NewDecoder(resp.Body)
		errDecode := decoder.Decode(&errorResponse)
		if errDecode != nil {
			return errDecode
		}
		return fmt.Errorf("klaviyo API Error: %s", errorResponse.Errors[0].Detail)

	}

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

type ErrorResponse struct {
	Errors []struct {
		ID     string `json:"id"`
		Status int    `json:"status"`
		Code   string `json:"code"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
	} `json:"errors"`
}
