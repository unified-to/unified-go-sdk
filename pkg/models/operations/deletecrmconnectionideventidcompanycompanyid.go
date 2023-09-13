// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteCrmConnectionIDEventIDCompanyCompanyIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest struct {
	// ID of the company
	CompanyID string `pathParam:"style=simple,explode=false,name=company_id"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse struct {
	ContentType string
	// Successful
	CrmEvent    *shared.CrmEvent
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
