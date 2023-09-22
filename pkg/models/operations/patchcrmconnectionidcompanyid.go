// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmConnectionIDCompanyIDRequest struct {
	// A company represents an organization that optionally is associated with a deal and/or contacts
	CrmCompany *shared.CrmCompany `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmConnectionIDCompanyIDRequest) GetCrmCompany() *shared.CrmCompany {
	if o == nil {
		return nil
	}
	return o.CrmCompany
}

func (o *PatchCrmConnectionIDCompanyIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmConnectionIDCompanyIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmConnectionIDCompanyIDResponse struct {
	ContentType string
	// Successful
	CrmCompany  *shared.CrmCompany
	StatusCode  int
	RawResponse *http.Response
}

func (o *PatchCrmConnectionIDCompanyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmConnectionIDCompanyIDResponse) GetCrmCompany() *shared.CrmCompany {
	if o == nil {
		return nil
	}
	return o.CrmCompany
}

func (o *PatchCrmConnectionIDCompanyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmConnectionIDCompanyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
