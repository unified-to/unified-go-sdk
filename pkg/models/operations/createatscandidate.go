// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAtsCandidateRequest struct {
	AtsCandidate *shared.AtsCandidate `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAtsCandidateRequest) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *CreateAtsCandidateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAtsCandidateResponse struct {
	// Successful
	AtsCandidate *shared.AtsCandidate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAtsCandidateResponse) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *CreateAtsCandidateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAtsCandidateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAtsCandidateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
