// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostAtsConnectionIDJobRequest struct {
	AtsJob *shared.AtsJob `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PostAtsConnectionIDJobRequest) GetAtsJob() *shared.AtsJob {
	if o == nil {
		return nil
	}
	return o.AtsJob
}

func (o *PostAtsConnectionIDJobRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostAtsConnectionIDJobResponse struct {
	// Successful
	AtsJob *shared.AtsJob
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAtsConnectionIDJobResponse) GetAtsJob() *shared.AtsJob {
	if o == nil {
		return nil
	}
	return o.AtsJob
}

func (o *PostAtsConnectionIDJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAtsConnectionIDJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAtsConnectionIDJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
