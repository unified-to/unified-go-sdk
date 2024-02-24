// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveAtsDocumentSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *RemoveAtsDocumentSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type RemoveAtsDocumentRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Document
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveAtsDocumentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveAtsDocumentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveAtsDocumentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *RemoveAtsDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveAtsDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveAtsDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveAtsDocumentResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
