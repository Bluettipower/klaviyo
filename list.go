package klaviyo

import (
	"context"
	"fmt"
)

const listBasePath = "api/lists"

type ListServiceOp struct {
	client *Client
}

type ListService interface {
	Read(context.Context, ReadListRequest) (*ListResponse, error)
	Browse(context.Context, BrowseListRequest) (*ListResponses, error)
	Edit(context.Context, EditList) (*ListResponse, error)
	Create(context.Context, CreateList) (*ListResponse, error)
	Delete(context.Context, DeleteListRequest) error
}

type ListResponse struct {
	Data ListResponseData `json:"data,omitempty"`
}

type ListResponses struct {
	Data []ListResponseData `json:"data,omitempty"`
}

type ListResponseData struct {
	Type       *string                  `json:"type,omitempty"`
	ID         *string                  `json:"id,omitempty"`
	Attributes *ListResponseAttributes  `json:"attributes,omitempty"`
	Links      map[string]string        `json:"links,omitempty"`
	Included   []map[string]interface{} `json:"included,omitempty"`
}

type ListResponseAttributes struct {
	Name         *string `json:"name,omitempty"`
	Created      *string `json:"created,omitempty"`
	Updated      *string `json:"updated,omitempty"`
	OptInProcess *string `json:"opt_in_process,omitempty"`
	ProfileCount *int    `json:"profile_count,omitempty"`
}

type ReadListRequest struct {
	ID string
}

type BrowseListRequest struct {
}

type EditList struct {
	Data EditListData `json:"data,omitempty"`
}

type EditListData struct {
	Type       string              `json:"type,omitempty"`
	ID         string              `json:"id,omitempty"`
	Attributes *EditListAttributes `json:"attributes,omitempty"`
}

type EditListAttributes struct {
	Name *string `json:"name,omitempty"`
}

type CreateList struct {
	Data CreateListData `json:"data,omitempty"`
}

type CreateListData struct {
	Type       string              `json:"type,omitempty"`
	Attributes *EditListAttributes `json:"attributes,omitempty"`
}

type DeleteListRequest struct {
	ID string
}

func (s *ListServiceOp) Read(ctx context.Context, readListRequest ReadListRequest) (*ListResponse, error) {
	resource := fmt.Sprintf("%v/%v/", listBasePath, readListRequest.ID)
	var resp ListResponse
	err := s.client.Request("GET", resource, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ListServiceOp) Browse(ctx context.Context, browseListRequest BrowseListRequest) (*ListResponses, error) {
	resource := fmt.Sprintf("%v/", listBasePath)
	var resp ListResponses
	err := s.client.Request("GET", resource, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ListServiceOp) Edit(ctx context.Context, editList EditList) (*ListResponse, error) {
	resource := fmt.Sprintf("%v/%v/", listBasePath, editList.Data.ID)
	var resp ListResponse
	err := s.client.Request("PATCH", resource, editList, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ListServiceOp) Create(ctx context.Context, createList CreateList) (*ListResponse, error) {
	resource := fmt.Sprintf("%v/", listBasePath)
	var resp ListResponse
	err := s.client.Request("POST", resource, createList, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ListServiceOp) Delete(ctx context.Context, deleteListRequest DeleteListRequest) error {
	resource := fmt.Sprintf("%v/%v/", listBasePath, deleteListRequest.ID)
	err := s.client.Request("DELETE", resource, nil, nil)
	return err
}
