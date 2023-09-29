// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type GetTicketingConnectionIDAgentRequest struct {
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

func (g GetTicketingConnectionIDAgentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTicketingConnectionIDAgentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTicketingConnectionIDAgentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetTicketingConnectionIDAgentRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetTicketingConnectionIDAgentRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetTicketingConnectionIDAgentRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetTicketingConnectionIDAgentRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetTicketingConnectionIDAgentRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetTicketingConnectionIDAgentRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetTicketingConnectionIDAgentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingAgents []shared.TicketingAgent
}

func (o *GetTicketingConnectionIDAgentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTicketingConnectionIDAgentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTicketingConnectionIDAgentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTicketingConnectionIDAgentResponse) GetTicketingAgents() []shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgents
}
