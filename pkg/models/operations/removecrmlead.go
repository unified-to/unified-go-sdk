// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveCrmLeadRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Lead
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveCrmLeadRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveCrmLeadRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveCrmLeadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	RemoveCrmLeadDefaultApplicationJSONString *string
}

func (o *RemoveCrmLeadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveCrmLeadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveCrmLeadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveCrmLeadResponse) GetRemoveCrmLeadDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.RemoveCrmLeadDefaultApplicationJSONString
}