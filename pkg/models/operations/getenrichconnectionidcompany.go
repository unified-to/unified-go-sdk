// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEnrichConnectionIDCompanySecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetEnrichConnectionIDCompanySecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetEnrichConnectionIDCompanyRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// The domain of the company to search
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The name of the company to search
	Name *string `queryParam:"style=form,explode=true,name=name"`
}

func (o *GetEnrichConnectionIDCompanyRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetEnrichConnectionIDCompanyRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetEnrichConnectionIDCompanyRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetEnrichConnectionIDCompanyResponse struct {
	ContentType string
	// Successful
	EnrichCompany *shared.EnrichCompany
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetEnrichConnectionIDCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnrichConnectionIDCompanyResponse) GetEnrichCompany() *shared.EnrichCompany {
	if o == nil {
		return nil
	}
	return o.EnrichCompany
}

func (o *GetEnrichConnectionIDCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnrichConnectionIDCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
