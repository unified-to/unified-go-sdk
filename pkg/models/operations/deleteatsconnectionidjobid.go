// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAtsConnectionIDJobIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *DeleteAtsConnectionIDJobIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type DeleteAtsConnectionIDJobIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Job
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteAtsConnectionIDJobIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteAtsConnectionIDJobIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteAtsConnectionIDJobIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	DeleteAtsConnectionIDJobIDDefaultApplicationJSONString *string
}

func (o *DeleteAtsConnectionIDJobIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAtsConnectionIDJobIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAtsConnectionIDJobIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteAtsConnectionIDJobIDResponse) GetDeleteAtsConnectionIDJobIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteAtsConnectionIDJobIDDefaultApplicationJSONString
}
