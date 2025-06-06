// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListCommerceCollectionsRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// The parent ID to filter by
	ParentID *string `queryParam:"style=form,explode=true,name=parent_id"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw  *string `queryParam:"style=form,explode=true,name=raw"`
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	Type *string `queryParam:"style=form,explode=true,name=type"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *string `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *ListCommerceCollectionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCommerceCollectionsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListCommerceCollectionsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCommerceCollectionsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListCommerceCollectionsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListCommerceCollectionsRequest) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *ListCommerceCollectionsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListCommerceCollectionsRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ListCommerceCollectionsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListCommerceCollectionsRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListCommerceCollectionsRequest) GetUpdatedGte() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListCommerceCollectionsResponse struct {
	// Successful
	CommerceCollections []shared.CommerceCollection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCommerceCollectionsResponse) GetCommerceCollections() []shared.CommerceCollection {
	if o == nil {
		return nil
	}
	return o.CommerceCollections
}

func (o *ListCommerceCollectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCommerceCollectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCommerceCollectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
