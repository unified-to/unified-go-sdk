// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteCrmConnectionIDContactIDDealDealIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type DeleteCrmConnectionIDContactIDDealDealIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the deal
	DealID string `pathParam:"style=simple,explode=false,name=deal_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDRequest) GetDealID() string {
	if o == nil {
		return ""
	}
	return o.DealID
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDContactIDDealDealIDResponse struct {
	ContentType string
	// Successful
	CrmContact  *shared.CrmContact
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDResponse) GetCrmContact() *shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContact
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDContactIDDealDealIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
