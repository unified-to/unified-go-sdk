// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAtsConnectionIDJobIDRequest struct {
	AtsJob *shared.AtsJob `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Job
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAtsConnectionIDJobIDRequest) GetAtsJob() *shared.AtsJob {
	if o == nil {
		return nil
	}
	return o.AtsJob
}

func (o *PatchAtsConnectionIDJobIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAtsConnectionIDJobIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAtsConnectionIDJobIDResponse struct {
	// Successful
	AtsJob *shared.AtsJob
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAtsConnectionIDJobIDResponse) GetAtsJob() *shared.AtsJob {
	if o == nil {
		return nil
	}
	return o.AtsJob
}

func (o *PatchAtsConnectionIDJobIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAtsConnectionIDJobIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAtsConnectionIDJobIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
