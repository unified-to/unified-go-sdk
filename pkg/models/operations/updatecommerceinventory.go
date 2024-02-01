// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateCommerceInventoryRequest struct {
	CommerceInventory *shared.CommerceInventory `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Inventory
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateCommerceInventoryRequest) GetCommerceInventory() *shared.CommerceInventory {
	if o == nil {
		return nil
	}
	return o.CommerceInventory
}

func (o *UpdateCommerceInventoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCommerceInventoryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateCommerceInventoryResponse struct {
	// Successful
	CommerceInventory *shared.CommerceInventory
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCommerceInventoryResponse) GetCommerceInventory() *shared.CommerceInventory {
	if o == nil {
		return nil
	}
	return o.CommerceInventory
}

func (o *UpdateCommerceInventoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCommerceInventoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCommerceInventoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
