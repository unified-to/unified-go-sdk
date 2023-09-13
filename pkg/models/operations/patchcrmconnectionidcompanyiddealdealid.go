// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmConnectionIDCompanyIDDealDealIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchCrmConnectionIDCompanyIDDealDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the deal
	DealID string `pathParam:"style=simple,explode=false,name=deal_id"`
	// ID of the Company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDRequest) GetDealID() string {
	if o == nil {
		return ""
	}
	return o.DealID
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmConnectionIDCompanyIDDealDealIDResponse struct {
	ContentType string
	// Successful
	CrmCompany  *shared.CrmCompany
	StatusCode  int
	RawResponse *http.Response
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDResponse) GetCrmCompany() *shared.CrmCompany {
	if o == nil {
		return nil
	}
	return o.CrmCompany
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmConnectionIDCompanyIDDealDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
