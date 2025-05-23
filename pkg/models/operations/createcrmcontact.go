// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCrmContactRequest struct {
	// A contact represents a person that optionally is associated with a deal and/or a company
	CrmContact shared.CrmContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *CreateCrmContactRequest) GetCrmContact() shared.CrmContact {
	if o == nil {
		return shared.CrmContact{}
	}
	return o.CrmContact
}

func (o *CreateCrmContactRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateCrmContactRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreateCrmContactRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type CreateCrmContactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmContact *shared.CrmContact
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCrmContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCrmContactResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *CreateCrmContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCrmContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
