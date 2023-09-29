// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutAtsConnectionIDScorecardIDRequest struct {
	AtsScorecard *shared.AtsScorecard `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Document
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutAtsConnectionIDScorecardIDRequest) GetAtsScorecard() *shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecard
}

func (o *PutAtsConnectionIDScorecardIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutAtsConnectionIDScorecardIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutAtsConnectionIDScorecardIDResponse struct {
	// Successful
	AtsScorecard *shared.AtsScorecard
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAtsConnectionIDScorecardIDResponse) GetAtsScorecard() *shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecard
}

func (o *PutAtsConnectionIDScorecardIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAtsConnectionIDScorecardIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAtsConnectionIDScorecardIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
