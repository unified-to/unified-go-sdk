// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCrmDealRequest struct {
	// A deal represents an opportunity with companies and/or contacts
	CrmDeal *shared.CrmDeal `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateCrmDealRequest) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *CreateCrmDealRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateCrmDealRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateCrmDealResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmDeal *shared.CrmDeal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCrmDealResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCrmDealResponse) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *CreateCrmDealResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCrmDealResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
