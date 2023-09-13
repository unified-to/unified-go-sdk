// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCrmConnectionIDLeadIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchCrmConnectionIDLeadIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchCrmConnectionIDLeadIDRequest struct {
	CrmLead *shared.CrmLead `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Lead
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCrmConnectionIDLeadIDRequest) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *PatchCrmConnectionIDLeadIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCrmConnectionIDLeadIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCrmConnectionIDLeadIDResponse struct {
	ContentType string
	// Successful
	CrmLead     *shared.CrmLead
	StatusCode  int
	RawResponse *http.Response
}

func (o *PatchCrmConnectionIDLeadIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCrmConnectionIDLeadIDResponse) GetCrmLead() *shared.CrmLead {
	if o == nil {
		return nil
	}
	return o.CrmLead
}

func (o *PatchCrmConnectionIDLeadIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCrmConnectionIDLeadIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
