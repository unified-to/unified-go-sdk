// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchRepoOrganizationRequest struct {
	RepoOrganization shared.RepoOrganization `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Organization
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchRepoOrganizationRequest) GetRepoOrganization() shared.RepoOrganization {
	if o == nil {
		return shared.RepoOrganization{}
	}
	return o.RepoOrganization
}

func (o *PatchRepoOrganizationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchRepoOrganizationRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchRepoOrganizationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchRepoOrganizationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	RepoOrganization *shared.RepoOrganization
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchRepoOrganizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchRepoOrganizationResponse) GetRepoOrganization() *shared.RepoOrganization {
	if o == nil {
		return nil
	}
	return o.RepoOrganization
}

func (o *PatchRepoOrganizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchRepoOrganizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
