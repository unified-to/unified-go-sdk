// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutMartechConnectionIDListIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PutMartechConnectionIDListIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PutMartechConnectionIDListIDRequest struct {
	// Mailing List
	MarketingList *shared.MarketingList `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the List
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutMartechConnectionIDListIDRequest) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *PutMartechConnectionIDListIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutMartechConnectionIDListIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutMartechConnectionIDListIDResponse struct {
	ContentType string
	// Successful
	MarketingList *shared.MarketingList
	StatusCode    int
	RawResponse   *http.Response
}

func (o *PutMartechConnectionIDListIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutMartechConnectionIDListIDResponse) GetMarketingList() *shared.MarketingList {
	if o == nil {
		return nil
	}
	return o.MarketingList
}

func (o *PutMartechConnectionIDListIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutMartechConnectionIDListIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
