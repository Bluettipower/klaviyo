package klaviyo

import (
	"context"
	"strings"
)

type ProfileServiceOp struct {
	client *Client
}

type ProfileService interface {
	Read(context.Context, ProfileRequest) (*ProfileResponse, error)
	Browse(context.Context, ProfileRequest) (*ProfilesResponse, error)
	Edit(context.Context, EditProfile) (*ProfilesResponse, error)
	Create(context.Context, CreateProfile) (*ProfilesResponse, error)
}

type ProfileRequest struct {
	ID     string
	Emails *[]string
}

type EditProfile struct {
	Data EditProfileData `json:"data,omitempty"`
}

type EditProfileData struct {
	Type       string                 `json:"type,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Attributes *EditProfileAttributes `json:"attributes,omitempty"`
	Meta       *EditProfileMeta       `json:"meta,omitempty"`
}

type EditProfileAttributes struct {
	Email        *string              `json:"email,omitempty"`
	PhoneNumber  *string              `json:"phone_number,omitempty"`
	ExternalID   *string              `json:"external_id,omitempty"`
	AnonymousID  *string              `json:"anonymous_id,omitempty"`
	FirstName    *string              `json:"first_name,omitempty"`
	LastName     *string              `json:"last_name,omitempty"`
	Organization *string              `json:"organization,omitempty"`
	Title        *string              `json:"title,omitempty"`
	Image        *string              `json:"image,omitempty"`
	Location     *EditProfileLocation `json:"location,omitempty"`
	Properties   interface{}          `json:"properties,omitempty"`
}

type EditProfileLocation struct {
	Address1  *string `json:"address1,omitempty"`
	Address2  *string `json:"address2,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Region    *string `json:"region,omitempty"`
	Zip       *string `json:"zip,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
	IP        *string `json:"ip,omitempty"`
}

type EditProfileMeta struct {
	PatchProperties *EditProfileMetaPatchProperties `json:"patch_properties,omitempty"`
}

type EditProfileMetaPatchProperties struct {
	Unset    *string     `json:"unset,omitempty"`
	Append   interface{} `json:"append,omitempty"`
	Unappend interface{} `json:"unappend,omitempty"`
}

type CreateProfile struct {
	Data CreateProfileData `json:"data,omitempty"`
}

type CreateProfileData struct {
	Type       string                   `json:"type,omitempty"`
	Attributes *CreateProfileAttributes `json:"attributes,omitempty"`
}

type CreateProfileAttributes struct {
	Email        *string                `json:"email,omitempty"`
	PhoneNumber  *string                `json:"phone_number,omitempty"`
	FirstName    *string                `json:"first_name,omitempty"`
	LastName     *string                `json:"last_name,omitempty"`
	Organization *string                `json:"organization,omitempty"`
	Title        *string                `json:"title,omitempty"`
	Image        *string                `json:"image,omitempty"`
	Properties   *map[string]string     `json:"properties,omitempty"`
	Location     *CreateProfileLocation `json:"location,omitempty"`
}

type CreateProfileLocation struct {
	Address1  *string `json:"address1,omitempty"`
	Address2  *string `json:"address2,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Region    *string `json:"region,omitempty"`
	Zip       *string `json:"zip,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
	IP        *string `json:"ip,omitempty"`
}

type ProfileResponse struct {
	Data     ProfileResponseData      `json:"data,omitempty"`
	Included []map[string]interface{} `json:"included,omitempty"`
}

type ProfilesResponse struct {
	Links map[string]string     `json:"links,omitempty"`
	Data  []ProfileResponseData `json:"data,omitempty"`
}

type ProfileResponseData struct {
	Type       string                     `json:"type,omitempty"`
	ID         string                     `json:"id,omitempty"`
	Links      map[string]string          `json:"links,omitempty"`
	Attributes *ProfileResponseAttributes `json:"attributes,omitempty"`
}

type ProfileResponseAttributes struct {
	Email               string                              `json:"email,omitempty"`
	PhoneNumber         string                              `json:"phone_number,omitempty"`
	ExternalID          string                              `json:"external_id,omitempty"`
	FirstName           string                              `json:"first_name,omitempty"`
	LastName            string                              `json:"last_name,omitempty"`
	Organization        string                              `json:"organization,omitempty"`
	Title               string                              `json:"title,omitempty"`
	Image               string                              `json:"image,omitempty"`
	Created             string                              `json:"created,omitempty"`
	Updated             string                              `json:"updated,omitempty"`
	LastEventDate       string                              `json:"last_event_date,omitempty"`
	Properties          interface{}                         `json:"properties,omitempty"`
	Location            *ProfileResponseLocation            `json:"location,omitempty"`
	Subscriptions       *ProfileResponseSubscriptions       `json:"subscriptions,omitempty"`
	PredictiveAnalytics *ProfileResponsePredictiveAnalytics `json:"predictive_analytics,omitempty"`
}

type ProfileResponseLocation struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Region    string `json:"region,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	IP        string `json:"ip,omitempty"`
}

type ProfileResponseSubscriptions struct {
	Email *ProfileResponseSubscriptionEmail `json:"email,omitempty"`
	SMS   *ProfileResponseSubscriptionSMS   `json:"sms,omitempty"`
}

type ProfileResponsePredictiveAnalytics struct {
	HistoricCLV              int    `json:"historic_clv,omitempty"`
	PredictedCLV             int    `json:"predicted_clv,omitempty"`
	TotalCLV                 int    `json:"total_clv,omitempty"`
	HistoricNumberOfOrders   int    `json:"historic_number_of_orders,omitempty"`
	PredictedNumberOfOrders  int    `json:"predicted_number_of_orders,omitempty"`
	AverageDaysBetweenOrders int    `json:"average_days_between_orders,omitempty"`
	AverageOrderValue        int    `json:"average_order_value,omitempty"`
	ChurnProbability         int    `json:"churn_probability,omitempty"`
	ExpectedDateOfNextOrder  string `json:"expected_date_of_next_order,omitempty"`
}

type ProfileResponseSubscriptionEmail struct {
	Marketing struct {
		CanReceiveEmailMarketing string                                              `json:"can_receive_email_marketing,omitempty"`
		Consent                  string                                              `json:"consent,omitempty"`
		ConsentTimestamp         string                                              `json:"consent_timestamp,omitempty"`
		LastUpdated              string                                              `json:"last_updated,omitempty"`
		Method                   string                                              `json:"method,omitempty"`
		MethodDetail             string                                              `json:"method_detail,omitempty"`
		CustomMethodDetail       string                                              `json:"custom_method_detail,omitempty"`
		DoubleOptin              string                                              `json:"double_optin,omitempty"`
		Suppressions             *[]ProfileResponseSubscriptionEmailSuppressions     `json:"suppression,omitempty"`
		ListSuppressions         *[]ProfileResponseSubscriptionEmailListSuppressions `json:"list_suppressions,omitempty"`
	} `json:"marketing,omitempty"`
}

type ProfileResponseSubscriptionEmailSuppressions struct {
	Reason    string `json:"reason,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type ProfileResponseSubscriptionEmailListSuppressions struct {
	ListID    string `json:"list_id,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type ProfileResponseSubscriptionSMS struct {
	Marketing struct {
		CanReceiveSMSMarketing string `json:"can_receive_sms_marketing,omitempty"`
		Consent                string `json:"consent,omitempty"`
		ConsentTimestamp       string `json:"consent_timestamp,omitempty"`
		Method                 string `json:"method,omitempty"`
		MethodDetail           string `json:"method_detail,omitempty"`
		LastUpdated            string `json:"last_updated,omitempty"`
	} `json:"marketing,omitempty"`
}

const profileURL = "api/profiles/"

func (s *ProfileServiceOp) Read(ctx context.Context, params ProfileRequest) (*ProfileResponse, error) {

	var resp ProfileResponse

	url := profileURL + params.ID + "/"

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProfileServiceOp) Browse(ctx context.Context, params ProfileRequest) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	errRequest := s.client.Request("GET", profileURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProfileServiceOp) Edit(ctx context.Context, params EditProfile) (*ProfilesResponse, error) {

	var resp ProfilesResponse
	url := profileURL + params.Data.ID

	errRequest := s.client.Request("PATCH", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return nil, nil
}

func (s *ProfileServiceOp) Create(ctx context.Context, params CreateProfile) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	errRequest := s.client.Request("POST", profileURL, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
