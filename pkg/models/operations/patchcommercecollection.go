// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCommerceCollectionRequest struct {
	// A collection of items/products/services
	CommerceCollection *shared.CommerceCollection `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Collection
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCommerceCollectionRequest) GetCommerceCollection() *shared.CommerceCollection {
	if o == nil {
		return nil
	}
	return o.CommerceCollection
}

func (o *PatchCommerceCollectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCommerceCollectionRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchCommerceCollectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCommerceCollectionResponse struct {
	// Successful
	CommerceCollection *shared.CommerceCollection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCommerceCollectionResponse) GetCommerceCollection() *shared.CommerceCollection {
	if o == nil {
		return nil
	}
	return o.CommerceCollection
}

func (o *PatchCommerceCollectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCommerceCollectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCommerceCollectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
