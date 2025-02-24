// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchRepoRepositoryRequest struct {
	RepoRepository shared.RepoRepository `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Repository
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchRepoRepositoryRequest) GetRepoRepository() shared.RepoRepository {
	if o == nil {
		return shared.RepoRepository{}
	}
	return o.RepoRepository
}

func (o *PatchRepoRepositoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchRepoRepositoryRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchRepoRepositoryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchRepoRepositoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	RepoRepository *shared.RepoRepository
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchRepoRepositoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchRepoRepositoryResponse) GetRepoRepository() *shared.RepoRepository {
	if o == nil {
		return nil
	}
	return o.RepoRepository
}

func (o *PatchRepoRepositoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchRepoRepositoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
