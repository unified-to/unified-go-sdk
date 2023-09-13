// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"net/http"
)

type GetMartechConnectionIDListSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetMartechConnectionIDListSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetMartechConnectionIDListRequest struct {
	// ID of the connection
	ConnectionID string   `pathParam:"style=simple,explode=false,name=connection_id"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *types.Date `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *GetMartechConnectionIDListRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetMartechConnectionIDListRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetMartechConnectionIDListRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetMartechConnectionIDListRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetMartechConnectionIDListRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetMartechConnectionIDListRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetMartechConnectionIDListRequest) GetUpdatedGte() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetMartechConnectionIDListResponse struct {
	ContentType string
	// Successful
	MarketingLists []shared.MarketingList
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetMartechConnectionIDListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMartechConnectionIDListResponse) GetMarketingLists() []shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingLists
}

func (o *GetMartechConnectionIDListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMartechConnectionIDListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
