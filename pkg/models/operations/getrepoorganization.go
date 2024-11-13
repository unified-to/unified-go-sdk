// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRepoOrganizationRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Organization
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRepoOrganizationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetRepoOrganizationRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetRepoOrganizationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRepoOrganizationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	RepoOrganization *shared.RepoOrganization
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRepoOrganizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRepoOrganizationResponse) GetRepoOrganization() *shared.RepoOrganization {
	if o == nil {
		return nil
	}
	return o.RepoOrganization
}

func (o *GetRepoOrganizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRepoOrganizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}