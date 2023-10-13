// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAtsCandidateRequest struct {
	// A candidate looking for work
	AtsCandidate *shared.AtsCandidate `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Candidate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAtsCandidateRequest) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *PatchAtsCandidateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAtsCandidateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAtsCandidateResponse struct {
	// Successful
	AtsCandidate *shared.AtsCandidate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAtsCandidateResponse) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *PatchAtsCandidateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAtsCandidateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAtsCandidateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}