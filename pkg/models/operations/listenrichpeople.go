// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListEnrichPeopleRequest struct {
	// The name of the company the person is associated with.  Not valid by itself.
	CompanyName *string `queryParam:"style=form,explode=true,name=company_name"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// The email of the person to search
	Email *string `queryParam:"style=form,explode=true,name=email"`
	// The LinkedIn URL of the person to search
	LinkedinURL *string `queryParam:"style=form,explode=true,name=linkedin_url"`
	// The name of the person to search
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The twitter handle of the person to search
	Twitter *string `queryParam:"style=form,explode=true,name=twitter"`
}

func (o *ListEnrichPeopleRequest) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *ListEnrichPeopleRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListEnrichPeopleRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ListEnrichPeopleRequest) GetLinkedinURL() *string {
	if o == nil {
		return nil
	}
	return o.LinkedinURL
}

func (o *ListEnrichPeopleRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListEnrichPeopleRequest) GetTwitter() *string {
	if o == nil {
		return nil
	}
	return o.Twitter
}

type ListEnrichPeopleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	EnrichPerson *shared.EnrichPerson
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListEnrichPeopleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListEnrichPeopleResponse) GetEnrichPerson() *shared.EnrichPerson {
	if o == nil {
		return nil
	}
	return o.EnrichPerson
}

func (o *ListEnrichPeopleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListEnrichPeopleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
