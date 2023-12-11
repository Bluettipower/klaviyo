package klaviyo

import (
	"context"
	"fmt"
	"math"
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
	CreateEventProperties *map[string]string
	Time                  *string
	CreateEventMetric     *struct {
		Name    *string
		Service *string
	}
	CreateEventProfile *struct {
		Email *string
		Value *int
	}
	UniqueID *string
}

const eventURL = "https://a.klaviyo.com/api/events/"

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

	var resp EventResponse
	var eventBuild []string

	// Property
	if params.CreateEventProperties != nil {
		createEventProperties := *params.CreateEventProperties
		var eventPropertiesBuild []string
		for createEventPropertyKey, createEventProperty := range createEventProperties {
			eventPropertiesBuild = append(eventPropertiesBuild, "\""+createEventPropertyKey+"\":\""+createEventProperty+"\"")
		}

		eventBuild = append(eventBuild,
			fmt.Sprintf("\"properties\":{%v}", strings.Join(eventPropertiesBuild, ",")))
	}

	// Metric
	if params.CreateEventMetric != nil {
		createEventMetric := params.CreateEventMetric

		var eventMetricBuild []string
		if createEventMetric.Name != nil {
			eventMetricBuild = append(eventMetricBuild, "\"name\":\""+*createEventMetric.Name+"\"")
		}
		if createEventMetric.Service != nil {
			eventMetricBuild = append(eventMetricBuild, "\"service\":\""+*createEventMetric.Service+"\"")
		}
		eventBuild = append(eventBuild,
			fmt.Sprintf("\"metric\":{\"data\":{\"type\":\"metric\",\"attributes\":{%v}}}", strings.Join(eventMetricBuild, ",")))
	}

	// Profile
	if params.CreateEventProfile != nil {
		createEventProfile := params.CreateEventProfile

		var eventProfileBuild []string
		if createEventProfile.Email != nil {
			eventProfileBuild = append(eventProfileBuild, "\"email\":\""+*createEventProfile.Email+"\"")
		}
		if createEventProfile.Value != nil {
			ratio := math.Pow(10, float64(4))
			value := math.Round(float64(*createEventProfile.Value)/100*ratio) / ratio
			// eventBuild = append(eventBuild, "\"value\":100.00")
			eventBuild = append(eventBuild, fmt.Sprintf("\"value\":%.2f", value))
		}
		eventBuild = append(eventBuild,
			fmt.Sprintf("\"profile\":{\"data\":{\"type\":\"profile\",\"attributes\":{%v}}}", strings.Join(eventProfileBuild, ",")))
	}

	eventBuild = append(eventBuild, "\"unique_id\":\"1702060556\"")

	payloadString := fmt.Sprintf("{\"data\":{\"type\":\"event\",\"attributes\":{%v}}}", strings.Join(eventBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("POST", eventURL, *payload, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}