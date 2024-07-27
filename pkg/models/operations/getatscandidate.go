// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetAtsCandidateRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Candidate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAtsCandidateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAtsCandidateRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetAtsCandidateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAtsCandidateResponse struct {
	// Successful
	AtsCandidate *shared.AtsCandidate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAtsCandidateResponse) GetAtsCandidate() *shared.AtsCandidate {
	if o == nil {
		return nil
	}
	return o.AtsCandidate
}

func (o *GetAtsCandidateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAtsCandidateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAtsCandidateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
