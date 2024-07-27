// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetHrisGroupRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Group
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetHrisGroupRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetHrisGroupRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetHrisGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetHrisGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisGroup *shared.HrisGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHrisGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHrisGroupResponse) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *GetHrisGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHrisGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
