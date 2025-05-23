// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListRepoRepositoriesRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// The org ID to filter by
	OrgID *string `queryParam:"style=form,explode=true,name=org_id"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw  *string `queryParam:"style=form,explode=true,name=raw"`
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *string `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *ListRepoRepositoriesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListRepoRepositoriesRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListRepoRepositoriesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListRepoRepositoriesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListRepoRepositoriesRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListRepoRepositoriesRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

func (o *ListRepoRepositoriesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListRepoRepositoriesRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ListRepoRepositoriesRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListRepoRepositoriesRequest) GetUpdatedGte() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListRepoRepositoriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	RepoRepositories []shared.RepoRepository
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListRepoRepositoriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRepoRepositoriesResponse) GetRepoRepositories() []shared.RepoRepository {
	if o == nil {
		return nil
	}
	return o.RepoRepositories
}

func (o *ListRepoRepositoriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRepoRepositoriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
