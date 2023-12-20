// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetAccountingTaxrateRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Taxrate
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAccountingTaxrateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingTaxrateRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetAccountingTaxrateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAccountingTaxrateResponse struct {
	// Successful
	AccountingTaxrate *shared.AccountingTaxrate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingTaxrateResponse) GetAccountingTaxrate() *shared.AccountingTaxrate {
	if o == nil {
		return nil
	}
	return o.AccountingTaxrate
}

func (o *GetAccountingTaxrateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingTaxrateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingTaxrateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}