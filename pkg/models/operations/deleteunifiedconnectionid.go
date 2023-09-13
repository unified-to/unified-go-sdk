// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUnifiedConnectionIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *DeleteUnifiedConnectionIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type DeleteUnifiedConnectionIDRequest struct {
	// ID of the Connection
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteUnifiedConnectionIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteUnifiedConnectionIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	DeleteUnifiedConnectionIDDefaultApplicationJSONString *string
}

func (o *DeleteUnifiedConnectionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUnifiedConnectionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUnifiedConnectionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteUnifiedConnectionIDResponse) GetDeleteUnifiedConnectionIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteUnifiedConnectionIDDefaultApplicationJSONString
}
