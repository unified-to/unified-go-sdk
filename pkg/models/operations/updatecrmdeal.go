// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateCrmDealRequest struct {
	// A deal represents an opportunity with companies and/or contacts
	CrmDeal *shared.CrmDeal `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Deal
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateCrmDealRequest) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *UpdateCrmDealRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCrmDealRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateCrmDealRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateCrmDealResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmDeal *shared.CrmDeal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCrmDealResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCrmDealResponse) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *UpdateCrmDealResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCrmDealResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
