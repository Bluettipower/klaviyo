package klaviyo

import (
	"context"
	"strings"
)

type TagServiceOp struct {
	client *Client
}

type TagService interface {
	Read(context.Context, ReadTagRequest) (*TagResponse, error)
	Browse(context.Context, BrowseTagRequest) (*TagResponses, error)
	Edit(context.Context, EditTag) (*TagResponse, error)
	Create(context.Context, CreateTag) (*TagResponse, error)
}

type BrowseTagRequest struct {
	Name *string
}

type ReadTagRequest struct {
	ID string
}

type EditTag struct {
	Data EditTagData `json:"data,omitempty"`
}

type EditTagData struct {
	Type       string             `json:"type,omitempty"`
	ID         string             `json:"id,omitempty"`
	Attributes *EditTagAttributes `json:"attributes,omitempty"`
}

type EditTagAttributes struct {
	Name *string `json:"name,omitempty"`
}

type CreateTag struct {
	Data CreateTagData `json:"data,omitempty"`
}

type CreateTagData struct {
	Type          string                  `json:"type,omitempty"`
	Attributes    *CreateTagAttributes    `json:"attributes,omitempty"`
	Relationships *CreateTagRelationships `json:"relationships,omitempty"`
}

type CreateTagAttributes struct {
	Name *string `json:"name,omitempty"`
}

type CreateTagRelationships struct {
	TagGroup *CreateTagTagGroup `json:"tag-group,omitempty"`
}

type CreateTagTagGroup struct {
	Data *CreateTagTagGroupData `json:"data,omitempty"`
}

type CreateTagTagGroupData struct {
	Type string  `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

type TagResponse struct {
	Data     TagResponseData `json:"data,omitempty"`
	Included []TagIncluded   `json:"included,omitempty"`
}

type TagResponses struct {
	Links    map[string]string `json:"links,omitempty"`
	Data     []TagResponseData `json:"data,omitempty"`
	Included []TagIncluded     `json:"included,omitempty"`
}

type TagResponseData struct {
	ID            *string                `json:"id,omitempty"`
	Type          *string                `json:"type,omitempty"`
	Attributes    *TagResponseAttributes `json:"attributes,omitempty"`
	Links         *map[string]string     `json:"links,omitempty"`
	Relationships interface{}            `json:"relationships,omitempty"`
}

type TagIncluded struct {
	ID        *string                `json:"id,omitempty"`
	Type      *string                `json:"type,omitempty"`
	Attibutes *TagResponseAttributes `json:"attributes,omitempty"`
	Links     map[string]string      `json:"links,omitempty"`
}

type TagResponseAttributes struct {
	Name      *string `json:"name,omitempty"`
	Exclusive *bool   `json:"exclusive,omitempty"`
	Default   *bool   `json:"default,omitempty"`
}

const tagURL = "api/tags/"

func (s *TagServiceOp) Read(ctx context.Context, params ReadTagRequest) (*TagResponse, error) {

	var resp TagResponse

	url := tagURL + params.ID

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Browse(ctx context.Context, params BrowseTagRequest) (*TagResponses, error) {

	var resp TagResponses

	errRequest := s.client.Request("GET", tagURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Edit(ctx context.Context, params EditTag) (*TagResponse, error) {

	var resp TagResponse
	url := tagURL + params.Data.ID

	errRequest := s.client.Request("PATCH", url, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Create(ctx context.Context, params CreateTag) (*TagResponse, error) {

	var resp TagResponse

	errRequest := s.client.Request("POST", tagURL, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
