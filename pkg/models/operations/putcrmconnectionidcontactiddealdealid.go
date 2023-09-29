// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutCrmConnectionIDContactIDDealDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the deal
	DealID string `pathParam:"style=simple,explode=false,name=deal_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutCrmConnectionIDContactIDDealDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutCrmConnectionIDContactIDDealDealIDRequest) GetDealID() string {
	if o == nil {
		return ""
	}
	return o.DealID
}

func (o *PutCrmConnectionIDContactIDDealDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutCrmConnectionIDContactIDDealDealIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmContact *shared.CrmContact
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutCrmConnectionIDContactIDDealDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCrmConnectionIDContactIDDealDealIDResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *PutCrmConnectionIDContactIDDealDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCrmConnectionIDContactIDDealDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
