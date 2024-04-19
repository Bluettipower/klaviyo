package klaviyo

import (
	"context"
	"fmt"
	"strings"
)

type MetricServiceOp struct {
	client *Client
}

type MetricService interface {
	Read(context.Context, MetricRequest) (*MetricResponse, error)
	Browse(context.Context, MetricRequest) (*MetricResponses, error)
}

type MetricResponse struct {
	Data MetricResponseData `json:"data,omitempty"`
}

type MetricResponses struct {
	Data  []MetricResponseData `json:"data,omitempty"`
	Links map[string]string    `json:"links,omitempty"`
}

type MetricResponseData struct {
	Type       *string                   `json:"type,omitempty"`
	ID         *string                   `json:"id,omitempty"`
	Attributes *MetricResponseAttributes `json:"attributes,omitempty"`
	Links      *map[string]string        `json:"links,omitempty"`
}

type MetricResponseAttributes struct {
	Name        *string            `json:"name,omitempty"`
	Created     *string            `json:"created,omitempty"`
	Updated     *string            `json:"updated,omitempty"`
	Integration *map[string]string `json:"integration,omitempty"`
}

type MetricRequest struct {
	ID *string
}

type CreateMetric struct {
	Data CreateMetricData `json:"data,omitempty"`
}

type CreateMetricData struct {
	Type       string                  `json:"type,omitempty"`
	Attributes *CreateMetricAttributes `json:"attributes,omitempty"`
}

type CreateMetricAttributes struct {
	Name    *string `json:"name,omitempty"`
	Service *string `json:"service,omitempty"`
}

const metricURL = "metrics/"

func (s *MetricServiceOp) Read(ctx context.Context, params MetricRequest) (*MetricResponse, error) {

	var resp MetricResponse
	url := fmt.Sprintf("%v%v/", metricURL, *params.ID)

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *MetricServiceOp) Browse(ctx context.Context, params MetricRequest) (*MetricResponses, error) {

	var resp MetricResponses

	errRequest := s.client.Request("GET", metricURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
