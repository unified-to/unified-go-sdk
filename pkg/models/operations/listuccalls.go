// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListUcCallsSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *ListUcCallsSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type ListUcCallsRequest struct {
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	ContactID    *string `queryParam:"style=form,explode=true,name=contact_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
	UserID     *string    `queryParam:"style=form,explode=true,name=user_id"`
}

func (l ListUcCallsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUcCallsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUcCallsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListUcCallsRequest) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *ListUcCallsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListUcCallsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUcCallsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUcCallsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUcCallsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListUcCallsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUcCallsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

func (o *ListUcCallsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type ListUcCallsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	UcCalls []shared.UcCall
}

func (o *ListUcCallsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUcCallsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUcCallsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUcCallsResponse) GetUcCalls() []shared.UcCall {
	if o == nil {
		return nil
	}
	return o.UcCalls
}
