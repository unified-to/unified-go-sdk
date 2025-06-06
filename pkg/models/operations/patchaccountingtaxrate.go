// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAccountingTaxrateRequest struct {
	AccountingTaxrate shared.AccountingTaxrate `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Taxrate
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *PatchAccountingTaxrateRequest) GetAccountingTaxrate() shared.AccountingTaxrate {
	if o == nil {
		return shared.AccountingTaxrate{}
	}
	return o.AccountingTaxrate
}

func (o *PatchAccountingTaxrateRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAccountingTaxrateRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchAccountingTaxrateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchAccountingTaxrateRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type PatchAccountingTaxrateResponse struct {
	// Successful
	AccountingTaxrate *shared.AccountingTaxrate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAccountingTaxrateResponse) GetAccountingTaxrate() *shared.AccountingTaxrate {
	if o == nil {
		return nil
	}
	return o.AccountingTaxrate
}

func (o *PatchAccountingTaxrateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAccountingTaxrateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAccountingTaxrateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
