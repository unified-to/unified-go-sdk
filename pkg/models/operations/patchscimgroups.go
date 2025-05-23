// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchScimGroupsRequest struct {
	ScimGroup shared.ScimGroup `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Group
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchScimGroupsRequest) GetScimGroup() shared.ScimGroup {
	if o == nil {
		return shared.ScimGroup{}
	}
	return o.ScimGroup
}

func (o *PatchScimGroupsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchScimGroupsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchScimGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	ScimGroup *shared.ScimGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchScimGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchScimGroupsResponse) GetScimGroup() *shared.ScimGroup {
	if o == nil {
		return nil
	}
	return o.ScimGroup
}

func (o *PatchScimGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchScimGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
