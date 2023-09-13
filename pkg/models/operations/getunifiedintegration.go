// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"net/http"
)

type GetUnifiedIntegrationSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetUnifiedIntegrationSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetUnifiedIntegrationCategories string

const (
	GetUnifiedIntegrationCategoriesPassthrough GetUnifiedIntegrationCategories = "passthrough"
	GetUnifiedIntegrationCategoriesHris        GetUnifiedIntegrationCategories = "hris"
	GetUnifiedIntegrationCategoriesAts         GetUnifiedIntegrationCategories = "ats"
	GetUnifiedIntegrationCategoriesAuth        GetUnifiedIntegrationCategories = "auth"
	GetUnifiedIntegrationCategoriesCrm         GetUnifiedIntegrationCategories = "crm"
	GetUnifiedIntegrationCategoriesEnrich      GetUnifiedIntegrationCategories = "enrich"
	GetUnifiedIntegrationCategoriesMartech     GetUnifiedIntegrationCategories = "martech"
	GetUnifiedIntegrationCategoriesTicketing   GetUnifiedIntegrationCategories = "ticketing"
	GetUnifiedIntegrationCategoriesUc          GetUnifiedIntegrationCategories = "uc"
)

func (e GetUnifiedIntegrationCategories) ToPointer() *GetUnifiedIntegrationCategories {
	return &e
}

func (e *GetUnifiedIntegrationCategories) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "passthrough":
		fallthrough
	case "hris":
		fallthrough
	case "ats":
		fallthrough
	case "auth":
		fallthrough
	case "crm":
		fallthrough
	case "enrich":
		fallthrough
	case "martech":
		fallthrough
	case "ticketing":
		fallthrough
	case "uc":
		*e = GetUnifiedIntegrationCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetUnifiedIntegrationCategories: %v", v)
	}
}

type GetUnifiedIntegrationRequest struct {
	// Filter the results for only the workspace's active integrations
	Active *bool `queryParam:"style=form,explode=true,name=active"`
	// Filter the results on these categories
	Categories []GetUnifiedIntegrationCategories `queryParam:"style=form,explode=true,name=categories"`
	Limit      *float64                          `queryParam:"style=form,explode=true,name=limit"`
	Offset     *float64                          `queryParam:"style=form,explode=true,name=offset"`
	Order      *string                           `queryParam:"style=form,explode=true,name=order"`
	Sort       *string                           `queryParam:"style=form,explode=true,name=sort"`
	Summary    *bool                             `queryParam:"style=form,explode=true,name=summary"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *types.Date `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *GetUnifiedIntegrationRequest) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *GetUnifiedIntegrationRequest) GetCategories() []GetUnifiedIntegrationCategories {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *GetUnifiedIntegrationRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUnifiedIntegrationRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUnifiedIntegrationRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetUnifiedIntegrationRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetUnifiedIntegrationRequest) GetSummary() *bool {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetUnifiedIntegrationRequest) GetUpdatedGte() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetUnifiedIntegrationResponse struct {
	ContentType string
	// Successful
	Integrations []shared.Integration
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetUnifiedIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedIntegrationResponse) GetIntegrations() []shared.Integration {
	if o == nil {
		return nil
	}
	return o.Integrations
}

func (o *GetUnifiedIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
