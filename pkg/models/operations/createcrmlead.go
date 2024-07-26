// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCrmLeadRequest struct {
	CrmLead *shared.CrmLead `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateCrmLeadRequest) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *CreateCrmLeadRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateCrmLeadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmLead *shared.CrmLead
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCrmLeadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCrmLeadResponse) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *CreateCrmLeadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCrmLeadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
