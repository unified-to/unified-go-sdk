// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAtsDocumentRequest struct {
	AtsDocument *shared.AtsDocument `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAtsDocumentRequest) GetAtsDocument() *shared.AtsDocument {
	if o == nil {
		return nil
	}
	return o.AtsDocument
}

func (o *CreateAtsDocumentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAtsDocumentResponse struct {
	// Successful
	AtsDocument *shared.AtsDocument
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAtsDocumentResponse) GetAtsDocument() *shared.AtsDocument {
	if o == nil {
		return nil
	}
	return o.AtsDocument
}

func (o *CreateAtsDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAtsDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAtsDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}