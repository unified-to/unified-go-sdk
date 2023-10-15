// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmPipelineRequest struct {
	CrmPipeline *shared.CrmPipeline `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Pipeline
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmPipelineRequest) GetCrmPipeline() *shared.CrmPipeline {
	if o == nil {
		return nil
	}
	return o.CrmPipeline
}

func (o *PatchCrmPipelineRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmPipelineRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmPipelineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmPipeline *shared.CrmPipeline
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCrmPipelineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmPipelineResponse) GetCrmPipeline() *shared.CrmPipeline {
	if o == nil {
		return nil
	}
	return o.CrmPipeline
}

func (o *PatchCrmPipelineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmPipelineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
