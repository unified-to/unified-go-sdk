// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUnifiedConnectionIDRequest struct {
	// ID of the Connection
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteUnifiedConnectionIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteUnifiedConnectionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	DeleteUnifiedConnectionIDDefaultApplicationJSONString *string
}

func (o *DeleteUnifiedConnectionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUnifiedConnectionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUnifiedConnectionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteUnifiedConnectionIDResponse) GetDeleteUnifiedConnectionIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteUnifiedConnectionIDDefaultApplicationJSONString
}
