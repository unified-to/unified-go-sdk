// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateCrmContactSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *UpdateCrmContactSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type UpdateCrmContactRequest struct {
	// A contact represents a person that optionally is associated with a deal and/or a company
	CrmContact *shared.CrmContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateCrmContactRequest) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *UpdateCrmContactRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCrmContactRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateCrmContactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmContact *shared.CrmContact
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCrmContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCrmContactResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *UpdateCrmContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCrmContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
