package klaviyo

import (
	"context"
	"fmt"
)

const clientBasePath = "client"
const clientSubscriptionResource = "subscriptions"

type ClientSerivceOp struct {
	client *Client
}

type ClientService interface {
	CreateClientSubscription(ctx context.Context, clientSubscription CreateClientSubscription, companyId string) error
}

type CreateClientSubscription struct {
	Data CreateClientSubscriptionData `json:"data,omitempty"`
}

type CreateClientSubscriptionData struct {
	Type          string                                 `json:"type,omitempty"`
	Attributes    *CreateClientSubscriptionAttributes    `json:"attributes,omitempty"`
	Relationships *CreateClientSubscriptionRelationships `json:"relationships,omitempty"`
}

type CreateClientSubscriptionAttributes struct {
	CustomSource *string                          `json:"custom_source,omitempty"`
	Profile      *CreateClientSubscriptionProfile `json:"profile,omitempty"`
}

type CreateClientSubscriptionProfile struct {
	Data CreateClientSubscriptionProfileData `json:"data,omitempty"`
}

type CreateClientSubscriptionProfileData struct {
	Type       string                 `json:"type,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Attributes *EditProfileAttributes `json:"attributes,omitempty"`
}

type CreateClientSubscriptionRelationships struct {
	List *CreateClientSubscriptionList `json:"list,omitempty"`
}

type CreateClientSubscriptionList struct {
	Data CreateClientSubscriptionListData `json:"data,omitempty"`
}

type CreateClientSubscriptionListData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

func (s *ClientSerivceOp) CreateClientSubscription(ctx context.Context, createClientSubscription CreateClientSubscription, companyId string) error {

	resource := fmt.Sprintf("%v/%v/?company_id=%s", clientBasePath, clientSubscriptionResource, companyId)
	err := s.client.Request("POST", resource, createClientSubscription, nil)
	return err
}
