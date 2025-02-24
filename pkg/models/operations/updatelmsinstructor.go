// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateLmsInstructorRequest struct {
	LmsInstructor shared.LmsInstructor `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Instructor
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateLmsInstructorRequest) GetLmsInstructor() shared.LmsInstructor {
	if o == nil {
		return shared.LmsInstructor{}
	}
	return o.LmsInstructor
}

func (o *UpdateLmsInstructorRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateLmsInstructorRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateLmsInstructorRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateLmsInstructorResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsInstructor *shared.LmsInstructor
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateLmsInstructorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLmsInstructorResponse) GetLmsInstructor() *shared.LmsInstructor {
	if o == nil {
		return nil
	}
	return o.LmsInstructor
}

func (o *UpdateLmsInstructorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLmsInstructorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
