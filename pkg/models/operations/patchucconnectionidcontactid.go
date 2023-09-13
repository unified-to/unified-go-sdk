// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchUcConnectionIDContactIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchUcConnectionIDContactIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchUcConnectionIDContactIDRequest struct {
	// A contact represents a person that optionally is associated with a call
	UcContact *shared.UcContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchUcConnectionIDContactIDRequest) GetUcContact() *shared.UcContact {
	if o == nil {
		return nil
	}
	return o.UcContact
}

func (o *PatchUcConnectionIDContactIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchUcConnectionIDContactIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchUcConnectionIDContactIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	UcContact *shared.UcContact
}

func (o *PatchUcConnectionIDContactIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchUcConnectionIDContactIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchUcConnectionIDContactIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchUcConnectionIDContactIDResponse) GetUcContact() *shared.UcContact {
	if o == nil {
		return nil
	}
	return o.UcContact
}
