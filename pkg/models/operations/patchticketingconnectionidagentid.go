// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchTicketingConnectionIDAgentIDRequest struct {
	TicketingAgent *shared.TicketingAgent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Agent
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchTicketingConnectionIDAgentIDRequest) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}

func (o *PatchTicketingConnectionIDAgentIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchTicketingConnectionIDAgentIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchTicketingConnectionIDAgentIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingAgent *shared.TicketingAgent
}

func (o *PatchTicketingConnectionIDAgentIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTicketingConnectionIDAgentIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTicketingConnectionIDAgentIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTicketingConnectionIDAgentIDResponse) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}
