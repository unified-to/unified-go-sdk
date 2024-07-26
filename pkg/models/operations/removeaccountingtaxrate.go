// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveAccountingTaxrateRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Taxrate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveAccountingTaxrateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveAccountingTaxrateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveAccountingTaxrateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	String *string
}

func (o *RemoveAccountingTaxrateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveAccountingTaxrateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveAccountingTaxrateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveAccountingTaxrateResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}
