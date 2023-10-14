// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateAtsInterviewRequest struct {
	// An interview between a candidate for a specific job
	AtsInterview *shared.AtsInterview `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Interview
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateAtsInterviewRequest) GetAtsInterview() *shared.AtsInterview {
	if o == nil {
		return nil
	}
	return o.AtsInterview
}

func (o *UpdateAtsInterviewRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateAtsInterviewRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateAtsInterviewRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateAtsInterviewResponse struct {
	// Successful
	AtsInterview *shared.AtsInterview
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAtsInterviewResponse) GetAtsInterview() *shared.AtsInterview {
	if o == nil {
		return nil
	}
	return o.AtsInterview
}

func (o *UpdateAtsInterviewResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAtsInterviewResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAtsInterviewResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
