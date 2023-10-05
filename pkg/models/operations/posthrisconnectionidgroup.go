// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostHrisConnectionIDGroupRequest struct {
	HrisGroup *shared.HrisGroup `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PostHrisConnectionIDGroupRequest) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *PostHrisConnectionIDGroupRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostHrisConnectionIDGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisGroup *shared.HrisGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostHrisConnectionIDGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostHrisConnectionIDGroupResponse) GetHrisGroup() *shared.HrisGroup {
	if o == nil {
		return nil
	}
	return o.HrisGroup
}

func (o *PostHrisConnectionIDGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostHrisConnectionIDGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}