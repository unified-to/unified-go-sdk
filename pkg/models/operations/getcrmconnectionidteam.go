// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type GetCrmConnectionIDTeamRequest struct {
	// ID of the connection
	ConnectionID string   `pathParam:"style=simple,explode=false,name=connection_id"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (g GetCrmConnectionIDTeamRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCrmConnectionIDTeamRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCrmConnectionIDTeamRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCrmConnectionIDTeamRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetCrmConnectionIDTeamRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetCrmConnectionIDTeamRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetCrmConnectionIDTeamRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetCrmConnectionIDTeamRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetCrmConnectionIDTeamRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetCrmConnectionIDTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmTeams []shared.CrmTeam
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCrmConnectionIDTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCrmConnectionIDTeamResponse) GetCrmTeams() []shared.CrmTeam {
	if o == nil {
		return nil
	}
	return o.CrmTeams
}

func (o *GetCrmConnectionIDTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCrmConnectionIDTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}