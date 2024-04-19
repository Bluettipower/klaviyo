package klaviyo

import (
	"context"
	"fmt"
	"strings"
)

type EventServiceOp struct {
	client *Client
}

type EventService interface {
	Read(context.Context, ReadRequest) (*EventResponse, error)
	Browse(context.Context, BrowseRequest) (*EventResponses, error)
	Create(context.Context, CreateEvent) (*EventResponse, error)
}

type EventResponse struct {
	Data     EventResponseData        `json:"data,omitempty"`
	Included []map[string]interface{} `json:"included,omitempty"`
}

type EventResponses struct {
	Data     []EventResponseData      `json:"data,omitempty"`
	Links    map[string]string        `json:"links,omitempty"`
	Included []map[string]interface{} `json:"included,omitempty"`
}

type EventResponseData struct {
	ID            *string                                `json:"id,omitempty"`
	Type          *string                                `json:"type,omitempty"`
	Attributes    *EventResponseAttributes               `json:"attributes,omitempty"`
	Links         *map[string]string                     `json:"links,omitempty"`
	Relationships *map[string]EventResponseRelationships `json:"relationships,omitempty"`
}

type EventResponseAttributes struct {
	UUID            *string      `json:"uuid,omitempty"`
	Timestamp       *int         `json:"timestamp,omitempty"`
	EventProperties *interface{} `json:"event_properties,omitempty"`
	DateTime        *string      `json:"datetime,omitempty"`
}

type EventResponseRelationships struct {
	Data  *EventResponseData `json:"data,omitempty"`
	Links *map[string]string `json:"links,omitempty"`
}

type ReadRequest struct {
	ID *string
}

type BrowseRequest struct {
	ID *string
}

type CreateEvent struct {
	Data CreateEventData `json:"data,omitempty"`
}

type CreateEventData struct {
	Type       string                 `json:"type,omitempty"`
	Attributes *CreateEventAttributes `json:"attributes,omitempty"`
}

type CreateEventAttributes struct {
	Properties interface{}    `json:"properties,omitempty"`
	Time       *string        `json:"time,omitempty"`
	Value      *float64       `json:"value,omitempty"`
	UniqueID   *string        `json:"unique_id,omitempty"`
	Metric     *CreateMetric  `json:"metric,omitempty"`
	Profile    *CreateProfile `json:"profile,omitempty"`
}

const eventURL = "events/"

func (s *EventServiceOp) Read(ctx context.Context, params ReadRequest) (*EventResponse, error) {

	var resp EventResponse
	paramURL := fmt.Sprintf("%v%v/", eventURL, *params.ID)

	errRequest := s.client.Request("GET", paramURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *EventServiceOp) Browse(ctx context.Context, params BrowseRequest) (*EventResponses, error) {

	var resp EventResponses

	errRequest := s.client.Request("GET", eventURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *EventServiceOp) Create(ctx context.Context, params CreateEvent) (*EventResponse, error) {

	// var resp EventResponse

	errRequest := s.client.Request("POST", eventURL, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return nil, nil
}
