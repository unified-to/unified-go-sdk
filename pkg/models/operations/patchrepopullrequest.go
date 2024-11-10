// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchRepoPullrequestRequest struct {
	RepoPullrequest *shared.RepoPullrequest `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Pullrequest
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchRepoPullrequestRequest) GetRepoPullrequest() *shared.RepoPullrequest {
	if o == nil {
		return nil
	}
	return o.RepoPullrequest
}

func (o *PatchRepoPullrequestRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchRepoPullrequestRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchRepoPullrequestRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchRepoPullrequestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	RepoPullrequest *shared.RepoPullrequest
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchRepoPullrequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchRepoPullrequestResponse) GetRepoPullrequest() *shared.RepoPullrequest {
	if o == nil {
		return nil
	}
	return o.RepoPullrequest
}

func (o *PatchRepoPullrequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchRepoPullrequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
