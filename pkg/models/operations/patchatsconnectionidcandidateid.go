// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAtsConnectionIDCandidateIDRequest struct {
	// A candidate looking for work
	AtsCandidate *shared.AtsCandidate `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Candidate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAtsConnectionIDCandidateIDRequest) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *PatchAtsConnectionIDCandidateIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAtsConnectionIDCandidateIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAtsConnectionIDCandidateIDResponse struct {
	// Successful
	AtsCandidate *shared.AtsCandidate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAtsConnectionIDCandidateIDResponse) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *PatchAtsConnectionIDCandidateIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAtsConnectionIDCandidateIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAtsConnectionIDCandidateIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
