// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostCrmConnectionIDDealRequest struct {
	// A deal represents an opportunity with companies and/or contacts
	CrmDeal *shared.CrmDeal `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PostCrmConnectionIDDealRequest) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *PostCrmConnectionIDDealRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostCrmConnectionIDDealResponse struct {
	ContentType string
	// Successful
	CrmDeal     *shared.CrmDeal
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostCrmConnectionIDDealResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostCrmConnectionIDDealResponse) GetCrmDeal() *shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeal
}

func (o *PostCrmConnectionIDDealResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostCrmConnectionIDDealResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
