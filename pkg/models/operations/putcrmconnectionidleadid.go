// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutCrmConnectionIDLeadIDRequest struct {
	CrmLead *shared.CrmLead `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Lead
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutCrmConnectionIDLeadIDRequest) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *PutCrmConnectionIDLeadIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutCrmConnectionIDLeadIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutCrmConnectionIDLeadIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmLead *shared.CrmLead
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutCrmConnectionIDLeadIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCrmConnectionIDLeadIDResponse) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *PutCrmConnectionIDLeadIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCrmConnectionIDLeadIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
