// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListEnrichCompaniesRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// The domain of the company to search
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The name of the company to search
	Name *string `queryParam:"style=form,explode=true,name=name"`
}

func (o *ListEnrichCompaniesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListEnrichCompaniesRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *ListEnrichCompaniesRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ListEnrichCompaniesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	EnrichCompany *shared.EnrichCompany
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListEnrichCompaniesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListEnrichCompaniesResponse) GetEnrichCompany() *shared.EnrichCompany {
	if o == nil {
		return nil
	}
	return o.EnrichCompany
}

func (o *ListEnrichCompaniesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListEnrichCompaniesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
