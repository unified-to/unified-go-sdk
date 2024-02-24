// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAtsApplicationSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchAtsApplicationSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchAtsApplicationRequest struct {
	AtsApplication *shared.AtsApplication `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Application
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAtsApplicationRequest) GetAtsApplication() *shared.AtsApplication {
	if o == nil {
		return nil
	}
	return o.AtsApplication
}

func (o *PatchAtsApplicationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAtsApplicationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAtsApplicationResponse struct {
	// Successful
	AtsApplication *shared.AtsApplication
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAtsApplicationResponse) GetAtsApplication() *shared.AtsApplication {
	if o == nil {
		return nil
	}
	return o.AtsApplication
}

func (o *PatchAtsApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAtsApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAtsApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
