// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateHrisGroupRequest struct {
	HrisGroup *shared.HrisGroup `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Group
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateHrisGroupRequest) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *UpdateHrisGroupRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateHrisGroupRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateHrisGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateHrisGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisGroup *shared.HrisGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateHrisGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHrisGroupResponse) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *UpdateHrisGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHrisGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
