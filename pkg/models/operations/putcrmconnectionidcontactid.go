// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutCrmConnectionIDContactIDRequest struct {
	// A contact represents a person that optionally is associated with a deal and/or a company
	CrmContact *shared.CrmContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutCrmConnectionIDContactIDRequest) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *PutCrmConnectionIDContactIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutCrmConnectionIDContactIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutCrmConnectionIDContactIDResponse struct {
	ContentType string
	// Successful
	CrmContact  *shared.CrmContact
	StatusCode  int
	RawResponse *http.Response
}

func (o *PutCrmConnectionIDContactIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCrmConnectionIDContactIDResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *PutCrmConnectionIDContactIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCrmConnectionIDContactIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
