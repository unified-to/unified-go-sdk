// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTaskProjectRequest struct {
	TaskProject *shared.TaskProject `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateTaskProjectRequest) GetTaskProject() *shared.TaskProject {
	if o == nil {
		return nil
	}
	return o.TaskProject
}

func (o *CreateTaskProjectRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateTaskProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TaskProject *shared.TaskProject
}

func (o *CreateTaskProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTaskProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTaskProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTaskProjectResponse) GetTaskProject() *shared.TaskProject {
	if o == nil {
		return nil
	}
	return o.TaskProject
}
