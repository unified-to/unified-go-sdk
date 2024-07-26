// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateKmsSpaceRequest struct {
	KmsSpace *shared.KmsSpace `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Space
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateKmsSpaceRequest) GetKmsSpace() *shared.KmsSpace {
	if o == nil {
		return nil
	}
	return o.KmsSpace
}

func (o *UpdateKmsSpaceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateKmsSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateKmsSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	KmsSpace *shared.KmsSpace
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateKmsSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateKmsSpaceResponse) GetKmsSpace() *shared.KmsSpace {
	if o == nil {
		return nil
	}
	return o.KmsSpace
}

func (o *UpdateKmsSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateKmsSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
