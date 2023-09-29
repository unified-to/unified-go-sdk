// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetCrmConnectionIDDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Deal
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCrmConnectionIDDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCrmConnectionIDDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetCrmConnectionIDDealIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmDeal *shared.CrmDeal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCrmConnectionIDDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCrmConnectionIDDealIDResponse) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *GetCrmConnectionIDDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCrmConnectionIDDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
