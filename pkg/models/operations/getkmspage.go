// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetKmsPageRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Page
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetKmsPageRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetKmsPageRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetKmsPageRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetKmsPageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	KmsPage *shared.KmsPage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetKmsPageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKmsPageResponse) GetKmsPage() *shared.KmsPage {
	if o == nil {
		return nil
	}
	return o.KmsPage
}

func (o *GetKmsPageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKmsPageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}