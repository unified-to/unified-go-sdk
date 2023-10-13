// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateMartechListRequest struct {
	// Mailing List
	MarketingList *shared.MarketingList `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the List
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateMartechListRequest) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *UpdateMartechListRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateMartechListRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateMartechListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MarketingList *shared.MarketingList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateMartechListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMartechListResponse) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *UpdateMartechListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMartechListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}