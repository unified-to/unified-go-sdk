// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchHrisCompanyRequest struct {
	HrisCompany *shared.HrisCompany `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchHrisCompanyRequest) GetHrisCompany() *shared.HrisCompany {
	if o == nil {
		return nil
	}
	return o.HrisCompany
}

func (o *PatchHrisCompanyRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchHrisCompanyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchHrisCompanyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisCompany *shared.HrisCompany
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchHrisCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchHrisCompanyResponse) GetHrisCompany() *shared.HrisCompany {
	if o == nil {
		return nil
	}
	return o.HrisCompany
}

func (o *PatchHrisCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchHrisCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
