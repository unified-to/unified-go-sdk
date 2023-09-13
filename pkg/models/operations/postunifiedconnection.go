// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostUnifiedConnectionSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PostUnifiedConnectionSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PostUnifiedConnectionResponse struct {
	// Successful
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostUnifiedConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *PostUnifiedConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostUnifiedConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostUnifiedConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
