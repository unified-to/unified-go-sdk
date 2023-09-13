// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetAtsConnectionIDScorecardIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetAtsConnectionIDScorecardIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetAtsConnectionIDScorecardIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Document
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAtsConnectionIDScorecardIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAtsConnectionIDScorecardIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAtsConnectionIDScorecardIDResponse struct {
	// Successful
	AtsScorecard *shared.AtsScorecard
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAtsConnectionIDScorecardIDResponse) GetAtsScorecard() *shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecard
}

func (o *GetAtsConnectionIDScorecardIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAtsConnectionIDScorecardIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAtsConnectionIDScorecardIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
