// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateCrmUserRequest struct {
	CrmUser *shared.CrmUser `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the User
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateCrmUserRequest) GetCrmUser() *shared.CrmUser {
	if o == nil {
		return nil
	}
	return o.CrmUser
}

func (o *UpdateCrmUserRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCrmUserRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateCrmUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateCrmUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmUser *shared.CrmUser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCrmUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCrmUserResponse) GetCrmUser() *shared.CrmUser {
	if o == nil {
		return nil
	}
	return o.CrmUser
}

func (o *UpdateCrmUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCrmUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
