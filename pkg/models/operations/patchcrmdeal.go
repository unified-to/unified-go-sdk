// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmDealRequest struct {
	// A deal represents an opportunity with companies and/or contacts
	CrmDeal *shared.CrmDeal `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Deal
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmDealRequest) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *PatchCrmDealRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmDealRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmDealResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmDeal *shared.CrmDeal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCrmDealResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmDealResponse) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *PatchCrmDealResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmDealResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}