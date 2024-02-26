// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveCommerceInventorySecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *RemoveCommerceInventorySecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type RemoveCommerceInventoryRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Inventory
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveCommerceInventoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveCommerceInventoryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveCommerceInventoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *RemoveCommerceInventoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveCommerceInventoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveCommerceInventoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveCommerceInventoryResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
