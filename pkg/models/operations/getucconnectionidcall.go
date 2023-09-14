// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type GetUcConnectionIDCallSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetUcConnectionIDCallSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetUcConnectionIDCallRequest struct {
	AgentID *string `queryParam:"style=form,explode=true,name=agent_id"`
	// ID of the connection
	ConnectionID string   `pathParam:"style=simple,explode=false,name=connection_id"`
	ContactID    *string  `queryParam:"style=form,explode=true,name=contact_id"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *GetUcConnectionIDCallRequest) GetAgentID() *string {
	if o == nil {
		return nil
	}
	return o.AgentID
}

func (o *GetUcConnectionIDCallRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetUcConnectionIDCallRequest) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *GetUcConnectionIDCallRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUcConnectionIDCallRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUcConnectionIDCallRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetUcConnectionIDCallRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetUcConnectionIDCallRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetUcConnectionIDCallRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetUcConnectionIDCallResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	UcCalls []shared.UcCall
}

func (o *GetUcConnectionIDCallResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUcConnectionIDCallResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUcConnectionIDCallResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUcConnectionIDCallResponse) GetUcCalls() []shared.UcCall {
	if o == nil {
		return nil
	}
	return o.UcCalls
}
