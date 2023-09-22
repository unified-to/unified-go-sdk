// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteCrmConnectionIDEventIDDealDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the deal
	DealID string `pathParam:"style=simple,explode=false,name=deal_id"`
	// ID of the Event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDRequest) GetDealID() string {
	if o == nil {
		return ""
	}
	return o.DealID
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDEventIDDealDealIDResponse struct {
	ContentType string
	// Successful
	CrmEvent    *shared.CrmEvent
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDResponse) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDEventIDDealDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
