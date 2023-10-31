// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveAccountingCustomerRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Customer
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveAccountingCustomerRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveAccountingCustomerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveAccountingCustomerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	RemoveAccountingCustomerDefaultApplicationJSONString *string
}

func (o *RemoveAccountingCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveAccountingCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveAccountingCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveAccountingCustomerResponse) GetRemoveAccountingCustomerDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.RemoveAccountingCustomerDefaultApplicationJSONString
}
