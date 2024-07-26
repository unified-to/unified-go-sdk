// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateAtsActivityRequest struct {
	AtsActivity *shared.AtsActivity `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Activity
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateAtsActivityRequest) GetAtsActivity() *shared.AtsActivity {
	if o == nil {
		return nil
	}
	return o.AtsActivity
}

func (o *UpdateAtsActivityRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateAtsActivityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateAtsActivityResponse struct {
	// Successful
	AtsActivity *shared.AtsActivity
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAtsActivityResponse) GetAtsActivity() *shared.AtsActivity {
	if o == nil {
		return nil
	}
	return o.AtsActivity
}

func (o *UpdateAtsActivityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAtsActivityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAtsActivityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
