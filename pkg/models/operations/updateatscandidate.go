// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateAtsCandidateRequest struct {
	AtsCandidate *shared.AtsCandidate `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Candidate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateAtsCandidateRequest) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *UpdateAtsCandidateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateAtsCandidateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateAtsCandidateResponse struct {
	// Successful
	AtsCandidate *shared.AtsCandidate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAtsCandidateResponse) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *UpdateAtsCandidateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAtsCandidateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAtsCandidateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
