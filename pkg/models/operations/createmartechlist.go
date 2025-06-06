// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMartechListRequest struct {
	// Mailing List
	MarketingList shared.MarketingList `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *CreateMartechListRequest) GetMarketingList() shared.MarketingList {
	if o == nil {
		return shared.MarketingList{}
	}
	return o.MarketingList
}

func (o *CreateMartechListRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateMartechListRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreateMartechListRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type CreateMartechListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MarketingList *shared.MarketingList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateMartechListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMartechListResponse) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *CreateMartechListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMartechListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
