// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetLmsCourseRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Course
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetLmsCourseRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetLmsCourseRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetLmsCourseRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetLmsCourseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsCourse *shared.LmsCourse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetLmsCourseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLmsCourseResponse) GetLmsCourse() *shared.LmsCourse {
	if o == nil {
		return nil
	}
	return o.LmsCourse
}

func (o *GetLmsCourseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLmsCourseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
