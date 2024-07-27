// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetAtsDocumentRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Document
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAtsDocumentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAtsDocumentRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetAtsDocumentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAtsDocumentResponse struct {
	// Successful
	AtsDocument *shared.AtsDocument
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAtsDocumentResponse) GetAtsDocument() *shared.AtsDocument {
	if o == nil {
		return nil
	}
	return o.AtsDocument
}

func (o *GetAtsDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAtsDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAtsDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
