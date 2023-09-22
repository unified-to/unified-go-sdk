// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutTicketingConnectionIDAgentIDRequest struct {
	TicketingAgent *shared.TicketingAgent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Agent
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutTicketingConnectionIDAgentIDRequest) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}

func (o *PutTicketingConnectionIDAgentIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutTicketingConnectionIDAgentIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutTicketingConnectionIDAgentIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	TicketingAgent *shared.TicketingAgent
}

func (o *PutTicketingConnectionIDAgentIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutTicketingConnectionIDAgentIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutTicketingConnectionIDAgentIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutTicketingConnectionIDAgentIDResponse) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}
