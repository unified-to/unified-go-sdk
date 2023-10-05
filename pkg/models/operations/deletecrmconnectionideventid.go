// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteCrmConnectionIDEventIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDEventIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDEventIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDEventIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	DeleteCrmConnectionIDEventIDDefaultApplicationJSONString *string
}

func (o *DeleteCrmConnectionIDEventIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDEventIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDEventIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteCrmConnectionIDEventIDResponse) GetDeleteCrmConnectionIDEventIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteCrmConnectionIDEventIDDefaultApplicationJSONString
}