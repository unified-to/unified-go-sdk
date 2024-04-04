// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetHrisTimeoffRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Timeoff
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetHrisTimeoffRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetHrisTimeoffRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetHrisTimeoffRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetHrisTimeoffResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisTimeoff *shared.HrisTimeoff
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHrisTimeoffResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHrisTimeoffResponse) GetHrisTimeoff() *shared.HrisTimeoff {
	if o == nil {
		return nil
	}
	return o.HrisTimeoff
}

func (o *GetHrisTimeoffResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHrisTimeoffResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
