// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateHrisCompanyRequest struct {
	HrisCompany shared.HrisCompany `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateHrisCompanyRequest) GetHrisCompany() shared.HrisCompany {
	if o == nil {
		return shared.HrisCompany{}
	}
	return o.HrisCompany
}

func (o *CreateHrisCompanyRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateHrisCompanyRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateHrisCompanyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisCompany *shared.HrisCompany
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateHrisCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateHrisCompanyResponse) GetHrisCompany() *shared.HrisCompany {
	if o == nil {
		return nil
	}
	return o.HrisCompany
}

func (o *CreateHrisCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateHrisCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
