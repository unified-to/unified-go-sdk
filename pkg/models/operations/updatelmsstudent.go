// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateLmsStudentRequest struct {
	LmsStudent shared.LmsStudent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Student
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateLmsStudentRequest) GetLmsStudent() shared.LmsStudent {
	if o == nil {
		return shared.LmsStudent{}
	}
	return o.LmsStudent
}

func (o *UpdateLmsStudentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateLmsStudentRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateLmsStudentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateLmsStudentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsStudent *shared.LmsStudent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateLmsStudentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLmsStudentResponse) GetLmsStudent() *shared.LmsStudent {
	if o == nil {
		return nil
	}
	return o.LmsStudent
}

func (o *UpdateLmsStudentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLmsStudentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
