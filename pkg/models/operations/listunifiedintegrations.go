// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListUnifiedIntegrationsQueryParamCategories string

const (
	ListUnifiedIntegrationsQueryParamCategoriesPassthrough ListUnifiedIntegrationsQueryParamCategories = "passthrough"
	ListUnifiedIntegrationsQueryParamCategoriesHris        ListUnifiedIntegrationsQueryParamCategories = "hris"
	ListUnifiedIntegrationsQueryParamCategoriesAts         ListUnifiedIntegrationsQueryParamCategories = "ats"
	ListUnifiedIntegrationsQueryParamCategoriesAuth        ListUnifiedIntegrationsQueryParamCategories = "auth"
	ListUnifiedIntegrationsQueryParamCategoriesCrm         ListUnifiedIntegrationsQueryParamCategories = "crm"
	ListUnifiedIntegrationsQueryParamCategoriesEnrich      ListUnifiedIntegrationsQueryParamCategories = "enrich"
	ListUnifiedIntegrationsQueryParamCategoriesMartech     ListUnifiedIntegrationsQueryParamCategories = "martech"
	ListUnifiedIntegrationsQueryParamCategoriesTicketing   ListUnifiedIntegrationsQueryParamCategories = "ticketing"
	ListUnifiedIntegrationsQueryParamCategoriesUc          ListUnifiedIntegrationsQueryParamCategories = "uc"
	ListUnifiedIntegrationsQueryParamCategoriesAccounting  ListUnifiedIntegrationsQueryParamCategories = "accounting"
)

func (e ListUnifiedIntegrationsQueryParamCategories) ToPointer() *ListUnifiedIntegrationsQueryParamCategories {
	return &e
}

func (e *ListUnifiedIntegrationsQueryParamCategories) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "accounting":
		*e = ListUnifiedIntegrationsQueryParamCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUnifiedIntegrationsQueryParamCategories: %v", v)
	}
}

type ListUnifiedIntegrationsRequest struct {
	// Filter the results for only the workspace's active integrations
	Active *bool `queryParam:"style=form,explode=true,name=active"`
	// Filter the results on these categories
	Categories []ListUnifiedIntegrationsQueryParamCategories `queryParam:"style=form,explode=true,name=categories"`
	Env        *string                                       `queryParam:"style=form,explode=true,name=env"`
	Limit      *float64                                      `queryParam:"style=form,explode=true,name=limit"`
	Offset     *float64                                      `queryParam:"style=form,explode=true,name=offset"`
	Order      *string                                       `queryParam:"style=form,explode=true,name=order"`
	Sort       *string                                       `queryParam:"style=form,explode=true,name=sort"`
	Summary    *bool                                         `queryParam:"style=form,explode=true,name=summary"`
	UpdatedGte *string                                       `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *ListUnifiedIntegrationsRequest) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ListUnifiedIntegrationsRequest) GetCategories() []ListUnifiedIntegrationsQueryParamCategories {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *ListUnifiedIntegrationsRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *ListUnifiedIntegrationsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUnifiedIntegrationsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUnifiedIntegrationsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUnifiedIntegrationsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUnifiedIntegrationsRequest) GetSummary() *bool {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *ListUnifiedIntegrationsRequest) GetUpdatedGte() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListUnifiedIntegrationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	Integrations []shared.Integration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUnifiedIntegrationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUnifiedIntegrationsResponse) GetIntegrations() []shared.Integration {
	if o == nil {
		return nil
	}
	return o.Integrations
}

func (o *ListUnifiedIntegrationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUnifiedIntegrationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
