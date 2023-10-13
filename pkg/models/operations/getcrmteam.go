// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetCrmTeamRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Team
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCrmTeamRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCrmTeamRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetCrmTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmTeam *shared.CrmTeam
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCrmTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCrmTeamResponse) GetCrmTeam() *shared.CrmTeam {
	if o == nil {
		return nil
	}
	return o.CrmTeam
}

func (o *GetCrmTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCrmTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}