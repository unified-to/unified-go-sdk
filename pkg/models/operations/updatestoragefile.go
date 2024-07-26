// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateStorageFileRequest struct {
	StorageFile *shared.StorageFile `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the File
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateStorageFileRequest) GetStorageFile() *shared.StorageFile {
	if o == nil {
		return nil
	}
	return o.StorageFile
}

func (o *UpdateStorageFileRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateStorageFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateStorageFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	StorageFile *shared.StorageFile
}

func (o *UpdateStorageFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateStorageFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateStorageFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateStorageFileResponse) GetStorageFile() *shared.StorageFile {
	if o == nil {
		return nil
	}
	return o.StorageFile
}
