// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchHrisGroupRequest struct {
	HrisGroup *shared.HrisGroup `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Group
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchHrisGroupRequest) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *PatchHrisGroupRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchHrisGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchHrisGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisGroup *shared.HrisGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchHrisGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchHrisGroupResponse) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *PatchHrisGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchHrisGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
