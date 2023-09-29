// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmConnectionIDEventIDCompanyCompanyIDRequest struct {
	// ID of the company
	CompanyID string `pathParam:"style=simple,explode=false,name=company_id"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmConnectionIDEventIDCompanyCompanyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmEvent *shared.CrmEvent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDResponse) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmConnectionIDEventIDCompanyCompanyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
