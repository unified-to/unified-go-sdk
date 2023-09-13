// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMartechConnectionIDListIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetMartechConnectionIDListIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetMartechConnectionIDListIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the List
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMartechConnectionIDListIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetMartechConnectionIDListIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMartechConnectionIDListIDResponse struct {
	ContentType string
	// Successful
	MarketingList *shared.MarketingList
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetMartechConnectionIDListIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMartechConnectionIDListIDResponse) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *GetMartechConnectionIDListIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMartechConnectionIDListIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
