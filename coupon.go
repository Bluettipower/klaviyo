package klaviyo

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type CouponServiceOp struct {
	client *Client
}

type CouponCodeServiceOp struct {
	client *Client
}

type CouponService interface {
	Read(context.Context, ReadCouponRequest) (*CouponResponse, error)
	Browse(context.Context, BrowseCouponRequest) (*CouponResponses, error)
	Edit(context.Context, EditCoupon) (*CouponResponse, error)
	Create(context.Context, CreateCoupon) (*CouponResponse, error)
	ReadCode(context.Context, ReadCouponCodeRequest) (*CouponCodeResponse, error)
	BrowseCodes(context.Context, BrowseCouponCodeRequest) (*CouponCodeResponses, error)
	EditCode(context.Context, EditCouponCode) (*CouponCodeResponse, error)
	CreateCode(context.Context, CreateCouponCode) (*CouponCodeResponse, error)
}

type BrowseCouponRequest struct {
}

type ReadCouponRequest struct {
	ID string
}

type EditCoupon struct {
	Data EditCouponData `json:"data,omitempty"`
}

type EditCouponData struct {
	Type       string                `json:"type,omitempty"`
	ID         string                `json:"id,omitempty"`
	Attributes *EditCouponAttributes `json:"attributes,omitempty"`
}

type EditCouponAttributes struct {
	Description *string `json:"description,omitempty"`
}

type CreateCoupon struct {
	Data CreateCouponData `json:"data,omitempty"`
}

type CreateCouponData struct {
	Type       string                  `json:"type,omitempty"`
	Attributes *CreateCouponAttributes `json:"attributes,omitempty"`
}

type CreateCouponAttributes struct {
	ExternalID  *string `json:"external_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CouponResponse struct {
	Data CouponResponseData `json:"data,omitempty"`
}

type CouponResponses struct {
	Data  []CouponResponseData `json:"data,omitempty"`
	Links map[string]string    `json:"links,omitempty"`
}

type CouponResponseData struct {
	ID         *string                   `json:"id,omitempty"`
	Type       *string                   `json:"type,omitempty"`
	Attributes *CouponResponseAttributes `json:"attributes,omitempty"`
	Links      *map[string]string        `json:"links,omitempty"`
}

type CouponResponseAttributes struct {
	ExternalID  *string `json:"external_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

type BrowseCouponCodeRequest struct {
	CouponID  *[]string
	ProfileID *[]string
}

type ReadCouponCodeRequest struct {
	ID string
}

type EditCouponCode struct {
	Data EditCouponCodeData `json:"data,omitempty"`
}

type EditCouponCodeData struct {
	Type       string                    `json:"type,omitempty"`
	ID         string                    `json:"id,omitempty"`
	Attributes *EditCouponCodeAttributes `json:"attributes,omitempty"`
}

type EditCouponCodeAttributes struct {
	Status    *string `json:"status,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

type CreateCouponCode struct {
	Data CreateCouponCodeData `json:"data,omitempty"`
}

type CreateCouponCodeData struct {
	Type          string                         `json:"type,omitempty"`
	Attributes    *CreateCouponCodeAttributes    `json:"attributes,omitempty"`
	Relationships *CreateCouponCodeRelationships `json:"relationships,omitempty"`
}

type CreateCouponCodeAttributes struct {
	UniqueCode *string `json:"unique_code,omitempty"`
	ExpiresAt  *string `json:"expires_at,omitempty"`
}

type CreateCouponCodeRelationships struct {
	Coupon *CreateCouponCodeCoupon `json:"coupon,omitempty"`
}

type CreateCouponCodeCoupon struct {
	Data *CreateCouponCodeCouponData `json:"data,omitempty"`
}

type CreateCouponCodeCouponData struct {
	Type string  `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

type CouponCodeResponse struct {
	Data CouponCodeResponseData
}

type CouponCodeResponses struct {
	Data     []CouponCodeResponseData `json:"data,omitempty"`
	Links    map[string]string        `json:"links,omitempty"`
	Included []CouponResponseData     `json:"included,omitempty"`
}

type CouponCodeResponseData struct {
	ID            *string                       `json:"id,omitempty"`
	Type          *string                       `json:"type,omitempty"`
	Attributes    *CouponCodeResponseAttributes `json:"attributes,omitempty"`
	Links         *map[string]string            `json:"links,omitempty"`
	Relationships interface{}                   `json:"relationships,omitempty"`
}

type CouponCodeResponseAttributes struct {
	UniqueCode *string `json:"unique_code,omitempty"`
	ExpiresAt  *string `json:"expires_at,omitempty"`
	Status     *string `json:"status,omitempty"`
}

const couponURL = "coupons/"
const couponCodeURL = "coupon-codes/"

func (s *CouponServiceOp) Read(ctx context.Context, params ReadCouponRequest) (*CouponResponse, error) {

	var resp CouponResponse

	url := fmt.Sprintf("%v%v", couponURL, params.ID)

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Browse(ctx context.Context, params BrowseCouponRequest) (*CouponResponses, error) {

	var resp CouponResponses

	errRequest := s.client.Request("GET", couponURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Edit(ctx context.Context, params EditCoupon) (*CouponResponse, error) {

	var resp CouponResponse
	url := fmt.Sprintf("%v%v", couponURL, params.Data.ID)

	errRequest := s.client.Request("PATCH", url, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Create(ctx context.Context, params CreateCoupon) (*CouponResponse, error) {

	var resp CouponResponse

	errRequest := s.client.Request("POST", couponURL, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) ReadCode(ctx context.Context, params ReadCouponCodeRequest) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse

	url := fmt.Sprintf("%v%v", couponCodeURL, params.ID)

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) BrowseCodes(ctx context.Context, params BrowseCouponCodeRequest) (*CouponCodeResponses, error) {

	var urlString string
	var resp CouponCodeResponses
	uri := url.Values{}

	if params.CouponID != nil {

		uri.Add("filter", "any(coupon.id,[\""+strings.Join(*params.CouponID, "\",\"")+"\"])")
		urlString = fmt.Sprintf("%v?%v", couponCodeURL, uri.Encode())
	}

	if params.ProfileID != nil {
		uri.Add("filter", "any(profile.id,[\""+strings.Join(*params.ProfileID, "\",\"")+"\"])")
		urlString = fmt.Sprintf("%v?%v", couponCodeURL, uri.Encode())
	}

	errRequest := s.client.Request("GET", urlString, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) EditCode(ctx context.Context, params EditCouponCode) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse
	url := fmt.Sprintf("%v%v", couponCodeURL, params.Data.ID)

	errRequest := s.client.Request("PATCH", url, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) CreateCode(ctx context.Context, params CreateCouponCode) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse

	errRequest := s.client.Request("POST", couponCodeURL, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
