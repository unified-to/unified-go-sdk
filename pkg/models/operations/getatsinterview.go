// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetAtsInterviewRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Interview
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAtsInterviewRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAtsInterviewRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetAtsInterviewRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAtsInterviewResponse struct {
	// Successful
	AtsInterview *shared.AtsInterview
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAtsInterviewResponse) GetAtsInterview() *shared.AtsInterview {
	if o == nil {
		return nil
	}
	return o.AtsInterview
}

func (o *GetAtsInterviewResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAtsInterviewResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAtsInterviewResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
