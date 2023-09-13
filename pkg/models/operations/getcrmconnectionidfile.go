// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"net/http"
)

type GetCrmConnectionIDFileSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetCrmConnectionIDFileSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetCrmConnectionIDFileRequest struct {
	// The company ID to filter results
	CompanyID *string `queryParam:"style=form,explode=true,name=company_id"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// The contact ID to filter results
	ContactID *string `queryParam:"style=form,explode=true,name=contact_id"`
	// The deal ID to filter results
	DealID *string  `queryParam:"style=form,explode=true,name=deal_id"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *types.Date `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *GetCrmConnectionIDFileRequest) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *GetCrmConnectionIDFileRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCrmConnectionIDFileRequest) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *GetCrmConnectionIDFileRequest) GetDealID() *string {
	if o == nil {
		return nil
	}
	return o.DealID
}

func (o *GetCrmConnectionIDFileRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetCrmConnectionIDFileRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetCrmConnectionIDFileRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetCrmConnectionIDFileRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetCrmConnectionIDFileRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetCrmConnectionIDFileRequest) GetUpdatedGte() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetCrmConnectionIDFileResponse struct {
	ContentType string
	// Successful
	CrmFiles    []shared.CrmFile
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCrmConnectionIDFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCrmConnectionIDFileResponse) GetCrmFiles() []shared.CrmFile {
	if o == nil {
		return nil
	}
	return o.CrmFiles
}

func (o *GetCrmConnectionIDFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCrmConnectionIDFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
