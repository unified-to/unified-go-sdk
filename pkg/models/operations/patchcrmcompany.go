// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmCompanyRequest struct {
	// A company represents an organization that optionally is associated with a deal and/or contacts
	CrmCompany shared.CrmCompany `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmCompanyRequest) GetCrmCompany() shared.CrmCompany {
	if o == nil {
		return shared.CrmCompany{}
	}
	return o.CrmCompany
}

func (o *PatchCrmCompanyRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmCompanyRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchCrmCompanyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmCompanyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmCompany *shared.CrmCompany
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCrmCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmCompanyResponse) GetCrmCompany() *shared.CrmCompany {
	if o == nil {
		return nil
	}
	return o.CrmCompany
}

func (o *PatchCrmCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
