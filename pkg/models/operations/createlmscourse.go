// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateLmsCourseRequest struct {
	LmsCourse shared.LmsCourse `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateLmsCourseRequest) GetLmsCourse() shared.LmsCourse {
	if o == nil {
		return shared.LmsCourse{}
	}
	return o.LmsCourse
}

func (o *CreateLmsCourseRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateLmsCourseRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateLmsCourseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsCourse *shared.LmsCourse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateLmsCourseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLmsCourseResponse) GetLmsCourse() *shared.LmsCourse {
	if o == nil {
		return nil
	}
	return o.LmsCourse
}

func (o *CreateLmsCourseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLmsCourseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
