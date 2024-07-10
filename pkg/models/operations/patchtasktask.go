// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchTaskTaskRequest struct {
	TaskTask *shared.TaskTask `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Task
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchTaskTaskRequest) GetTaskTask() *shared.TaskTask {
	if o == nil {
		return nil
	}
	return o.TaskTask
}

func (o *PatchTaskTaskRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchTaskTaskRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchTaskTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TaskTask *shared.TaskTask
}

func (o *PatchTaskTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTaskTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTaskTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTaskTaskResponse) GetTaskTask() *shared.TaskTask {
	if o == nil {
		return nil
	}
	return o.TaskTask
}