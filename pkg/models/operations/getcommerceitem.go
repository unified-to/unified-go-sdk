// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetCommerceItemSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetCommerceItemSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetCommerceItemRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Item
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCommerceItemRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCommerceItemRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetCommerceItemRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetCommerceItemResponse struct {
	// Successful
	CommerceItem *shared.CommerceItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCommerceItemResponse) GetCommerceItem() *shared.CommerceItem {
	if o == nil {
		return nil
	}
	return o.CommerceItem
}

func (o *GetCommerceItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCommerceItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCommerceItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
