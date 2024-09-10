// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMartechMemberRequest struct {
	// A member represents a person
	MarketingMember *shared.MarketingMember `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateMartechMemberRequest) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *CreateMartechMemberRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateMartechMemberRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateMartechMemberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MarketingMember *shared.MarketingMember
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateMartechMemberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMartechMemberResponse) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *CreateMartechMemberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMartechMemberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
