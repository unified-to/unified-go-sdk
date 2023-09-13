// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostTicketingConnectionIDAgentSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PostTicketingConnectionIDAgentSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PostTicketingConnectionIDAgentRequest struct {
	TicketingAgent *shared.TicketingAgent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PostTicketingConnectionIDAgentRequest) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}

func (o *PostTicketingConnectionIDAgentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostTicketingConnectionIDAgentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	TicketingAgent *shared.TicketingAgent
}

func (o *PostTicketingConnectionIDAgentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostTicketingConnectionIDAgentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostTicketingConnectionIDAgentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostTicketingConnectionIDAgentResponse) GetTicketingAgent() *shared.TicketingAgent {
	if o == nil {
		return nil
	}
	return o.TicketingAgent
}
