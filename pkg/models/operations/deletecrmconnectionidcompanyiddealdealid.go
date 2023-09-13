// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteCrmConnectionIDCompanyIDDealDealIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type DeleteCrmConnectionIDCompanyIDDealDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the deal
	DealID string `pathParam:"style=simple,explode=false,name=deal_id"`
	// ID of the Company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDRequest) GetDealID() string {
	if o == nil {
		return ""
	}
	return o.DealID
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDCompanyIDDealDealIDResponse struct {
	ContentType string
	// Successful
	CrmCompany  *shared.CrmCompany
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDResponse) GetCrmCompany() *shared.CrmCompany {
	if o == nil {
		return nil
	}
	return o.CrmCompany
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDCompanyIDDealDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
