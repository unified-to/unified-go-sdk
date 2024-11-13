// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchLmsStudentRequest struct {
	LmsStudent *shared.LmsStudent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Student
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchLmsStudentRequest) GetLmsStudent() *shared.LmsStudent {
	if o == nil {
		return nil
	}
	return o.LmsStudent
}

func (o *PatchLmsStudentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchLmsStudentRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchLmsStudentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchLmsStudentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsStudent *shared.LmsStudent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchLmsStudentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchLmsStudentResponse) GetLmsStudent() *shared.LmsStudent {
	if o == nil {
		return nil
	}
	return o.LmsStudent
}

func (o *PatchLmsStudentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchLmsStudentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}