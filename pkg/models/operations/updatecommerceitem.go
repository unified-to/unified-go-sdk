// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateCommerceItemRequest struct {
	CommerceItem *shared.CommerceItem `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Item
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateCommerceItemRequest) GetCommerceItem() *shared.CommerceItem {
	if o == nil {
		return nil
	}
	return o.CommerceItem
}

func (o *UpdateCommerceItemRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCommerceItemRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateCommerceItemResponse struct {
	// Successful
	CommerceItem *shared.CommerceItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCommerceItemResponse) GetCommerceItem() *shared.CommerceItem {
	if o == nil {
		return nil
	}
	return o.CommerceItem
}

func (o *UpdateCommerceItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCommerceItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCommerceItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
