// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateMartechMemberRequest struct {
	// A member represents a person
	MarketingMember *shared.MarketingMember `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Member
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the list
	ListID string `pathParam:"style=simple,explode=false,name=list_id"`
}

func (o *UpdateMartechMemberRequest) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *UpdateMartechMemberRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateMartechMemberRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateMartechMemberRequest) GetListID() string {
	if o == nil {
		return ""
	}
	return o.ListID
}

type UpdateMartechMemberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MarketingMember *shared.MarketingMember
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateMartechMemberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMartechMemberResponse) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *UpdateMartechMemberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMartechMemberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
