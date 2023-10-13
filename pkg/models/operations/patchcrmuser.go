// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmUserRequest struct {
	CrmUser *shared.CrmUser `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the User
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmUserRequest) GetCrmUser() *shared.CrmUser {
	if o == nil {
		return nil
	}
	return o.CrmUser
}

func (o *PatchCrmUserRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmUser *shared.CrmUser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCrmUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmUserResponse) GetCrmUser() *shared.CrmUser {
	if o == nil {
		return nil
	}
	return o.CrmUser
}

func (o *PatchCrmUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}