// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type QueryParamCategories string

const (
	QueryParamCategoriesPassthrough QueryParamCategories = "passthrough"
	QueryParamCategoriesHris        QueryParamCategories = "hris"
	QueryParamCategoriesAts         QueryParamCategories = "ats"
	QueryParamCategoriesAuth        QueryParamCategories = "auth"
	QueryParamCategoriesCrm         QueryParamCategories = "crm"
	QueryParamCategoriesEnrich      QueryParamCategories = "enrich"
	QueryParamCategoriesMartech     QueryParamCategories = "martech"
	QueryParamCategoriesTicketing   QueryParamCategories = "ticketing"
	QueryParamCategoriesUc          QueryParamCategories = "uc"
	QueryParamCategoriesAccounting  QueryParamCategories = "accounting"
	QueryParamCategoriesStorage     QueryParamCategories = "storage"
	QueryParamCategoriesCommerce    QueryParamCategories = "commerce"
	QueryParamCategoriesPayment     QueryParamCategories = "payment"
	QueryParamCategoriesGenai       QueryParamCategories = "genai"
	QueryParamCategoriesMessaging   QueryParamCategories = "messaging"
	QueryParamCategoriesKms         QueryParamCategories = "kms"
	QueryParamCategoriesTask        QueryParamCategories = "task"
	QueryParamCategoriesScim        QueryParamCategories = "scim"
	QueryParamCategoriesLms         QueryParamCategories = "lms"
	QueryParamCategoriesRepo        QueryParamCategories = "repo"
)

func (e QueryParamCategories) ToPointer() *QueryParamCategories {
	return &e
}
func (e *QueryParamCategories) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "storage":
		fallthrough
	case "commerce":
		fallthrough
	case "payment":
		fallthrough
	case "genai":
		fallthrough
	case "messaging":
		fallthrough
	case "kms":
		fallthrough
	case "task":
		fallthrough
	case "scim":
		fallthrough
	case "lms":
		fallthrough
	case "repo":
		*e = QueryParamCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamCategories: %v", v)
	}
}

type ListUnifiedIntegrationWorkspacesRequest struct {
	// Filter the results for only the workspace's active integrations
	Active *bool `queryParam:"style=form,explode=true,name=active"`
	// Filter the results on these categories
	Categories []QueryParamCategories `queryParam:"style=form,explode=true,name=categories"`
	Env        *string                `queryParam:"style=form,explode=true,name=env"`
	Limit      *float64               `queryParam:"style=form,explode=true,name=limit"`
	Offset     *float64               `queryParam:"style=form,explode=true,name=offset"`
	Summary    *bool                  `queryParam:"style=form,explode=true,name=summary"`
	UpdatedGte *string                `queryParam:"style=form,explode=true,name=updated_gte"`
	// The ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetCategories() []QueryParamCategories {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetSummary() *bool {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetUpdatedGte() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

func (o *ListUnifiedIntegrationWorkspacesRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type ListUnifiedIntegrationWorkspacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	Integrations []shared.Integration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUnifiedIntegrationWorkspacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUnifiedIntegrationWorkspacesResponse) GetIntegrations() []shared.Integration {
	if o == nil {
		return nil
	}
	return o.Integrations
}

func (o *ListUnifiedIntegrationWorkspacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUnifiedIntegrationWorkspacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
