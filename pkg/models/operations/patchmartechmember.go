// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchMartechMemberRequest struct {
	// A member represents a person
	MarketingMember *shared.MarketingMember `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Member
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchMartechMemberRequest) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *PatchMartechMemberRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchMartechMemberRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchMartechMemberRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchMartechMemberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MarketingMember *shared.MarketingMember
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchMartechMemberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchMartechMemberResponse) GetMarketingMember() *shared.MarketingMember {
	if o == nil {
		return nil
	}
	return o.MarketingMember
}

func (o *PatchMartechMemberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchMartechMemberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
