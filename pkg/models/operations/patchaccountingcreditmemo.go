// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAccountingCreditmemoRequest struct {
	AccountingCreditmemo shared.AccountingCreditmemo `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Creditmemo
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *PatchAccountingCreditmemoRequest) GetAccountingCreditmemo() shared.AccountingCreditmemo {
	if o == nil {
		return shared.AccountingCreditmemo{}
	}
	return o.AccountingCreditmemo
}

func (o *PatchAccountingCreditmemoRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAccountingCreditmemoRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchAccountingCreditmemoRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchAccountingCreditmemoRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type PatchAccountingCreditmemoResponse struct {
	// Successful
	AccountingCreditmemo *shared.AccountingCreditmemo
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAccountingCreditmemoResponse) GetAccountingCreditmemo() *shared.AccountingCreditmemo {
	if o == nil {
		return nil
	}
	return o.AccountingCreditmemo
}

func (o *PatchAccountingCreditmemoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAccountingCreditmemoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAccountingCreditmemoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
