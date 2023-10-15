// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAtsScorecardRequest struct {
	// A scorecard is feedback/assessment of a candidate's interview
	AtsScorecard *shared.AtsScorecard `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAtsScorecardRequest) GetAtsScorecard() *shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecard
}

func (o *CreateAtsScorecardRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAtsScorecardResponse struct {
	// Successful
	AtsScorecard *shared.AtsScorecard
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAtsScorecardResponse) GetAtsScorecard() *shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecard
}

func (o *CreateAtsScorecardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAtsScorecardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAtsScorecardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
