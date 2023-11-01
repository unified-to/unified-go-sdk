// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveAtsApplicationRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Application
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveAtsApplicationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveAtsApplicationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveAtsApplicationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *RemoveAtsApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveAtsApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveAtsApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveAtsApplicationResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
