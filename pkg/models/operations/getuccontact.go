// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUcContactRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUcContactRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetUcContactRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetUcContactRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUcContactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	UcContact *shared.UcContact
}

func (o *GetUcContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUcContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUcContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUcContactResponse) GetUcContact() *shared.UcContact {
	if o == nil {
		return nil
	}
	return o.UcContact
}
