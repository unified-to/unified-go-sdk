// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateKmsCommentRequest struct {
	KmsComment *shared.KmsComment `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Comment
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateKmsCommentRequest) GetKmsComment() *shared.KmsComment {
	if o == nil {
		return nil
	}
	return o.KmsComment
}

func (o *UpdateKmsCommentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateKmsCommentRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateKmsCommentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateKmsCommentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	KmsComment *shared.KmsComment
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateKmsCommentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateKmsCommentResponse) GetKmsComment() *shared.KmsComment {
	if o == nil {
		return nil
	}
	return o.KmsComment
}

func (o *UpdateKmsCommentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateKmsCommentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
