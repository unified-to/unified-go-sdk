// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMetadataMetadataRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Metadata
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMetadataMetadataRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetMetadataMetadataRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetMetadataMetadataRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMetadataMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MetadataMetadata *shared.MetadataMetadata
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMetadataMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMetadataMetadataResponse) GetMetadataMetadata() *shared.MetadataMetadata {
	if o == nil {
		return nil
	}
	return o.MetadataMetadata
}

func (o *GetMetadataMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMetadataMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}