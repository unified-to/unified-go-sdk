// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateGenaiPromptRequest struct {
	GenaiPrompt *shared.GenaiPrompt `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateGenaiPromptRequest) GetGenaiPrompt() *shared.GenaiPrompt {
	if o == nil {
		return nil
	}
	return o.GenaiPrompt
}

func (o *CreateGenaiPromptRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateGenaiPromptResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	GenaiPrompt *shared.GenaiPrompt
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateGenaiPromptResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGenaiPromptResponse) GetGenaiPrompt() *shared.GenaiPrompt {
	if o == nil {
		return nil
	}
	return o.GenaiPrompt
}

func (o *CreateGenaiPromptResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGenaiPromptResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
