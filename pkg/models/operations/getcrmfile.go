// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetCrmFileRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the File
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCrmFileRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCrmFileRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetCrmFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetCrmFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmFile *shared.CrmFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCrmFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCrmFileResponse) GetCrmFile() *shared.CrmFile {
	if o == nil {
		return nil
	}
	return o.CrmFile
}

func (o *GetCrmFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCrmFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
