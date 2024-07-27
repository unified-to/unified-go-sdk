// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCommerceItemRequest struct {
	CommerceItem *shared.CommerceItem `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateCommerceItemRequest) GetCommerceItem() *shared.CommerceItem {
	if o == nil {
		return nil
	}
	return o.CommerceItem
}

func (o *CreateCommerceItemRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateCommerceItemResponse struct {
	// Successful
	CommerceItem *shared.CommerceItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCommerceItemResponse) GetCommerceItem() *shared.CommerceItem {
	if o == nil {
		return nil
	}
	return o.CommerceItem
}

func (o *CreateCommerceItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCommerceItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCommerceItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
