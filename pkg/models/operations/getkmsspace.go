// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetKmsSpaceRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Space
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetKmsSpaceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetKmsSpaceRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetKmsSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetKmsSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	KmsSpace *shared.KmsSpace
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetKmsSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKmsSpaceResponse) GetKmsSpace() *shared.KmsSpace {
	if o == nil {
		return nil
	}
	return o.KmsSpace
}

func (o *GetKmsSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKmsSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
