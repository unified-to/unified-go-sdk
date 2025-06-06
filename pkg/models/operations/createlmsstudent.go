// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateLmsStudentRequest struct {
	LmsStudent shared.LmsStudent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *CreateLmsStudentRequest) GetLmsStudent() shared.LmsStudent {
	if o == nil {
		return shared.LmsStudent{}
	}
	return o.LmsStudent
}

func (o *CreateLmsStudentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateLmsStudentRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreateLmsStudentRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type CreateLmsStudentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsStudent *shared.LmsStudent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateLmsStudentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLmsStudentResponse) GetLmsStudent() *shared.LmsStudent {
	if o == nil {
		return nil
	}
	return o.LmsStudent
}

func (o *CreateLmsStudentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLmsStudentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
