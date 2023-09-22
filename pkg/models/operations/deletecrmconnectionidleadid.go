// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteCrmConnectionIDLeadIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Lead
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDLeadIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDLeadIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDLeadIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	DeleteCrmConnectionIDLeadIDDefaultApplicationJSONString *string
}

func (o *DeleteCrmConnectionIDLeadIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDLeadIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDLeadIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteCrmConnectionIDLeadIDResponse) GetDeleteCrmConnectionIDLeadIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteCrmConnectionIDLeadIDDefaultApplicationJSONString
}
