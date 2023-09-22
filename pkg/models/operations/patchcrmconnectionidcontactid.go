// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmConnectionIDContactIDRequest struct {
	// A contact represents a person that optionally is associated with a deal and/or a company
	CrmContact *shared.CrmContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmConnectionIDContactIDRequest) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *PatchCrmConnectionIDContactIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmConnectionIDContactIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmConnectionIDContactIDResponse struct {
	ContentType string
	// Successful
	CrmContact  *shared.CrmContact
	StatusCode  int
	RawResponse *http.Response
}

func (o *PatchCrmConnectionIDContactIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmConnectionIDContactIDResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *PatchCrmConnectionIDContactIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmConnectionIDContactIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
