// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAtsActivityRequest struct {
	AtsActivity *shared.AtsActivity `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAtsActivityRequest) GetAtsActivity() *shared.AtsActivity {
	if o == nil {
		return nil
	}
	return o.AtsActivity
}

func (o *CreateAtsActivityRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAtsActivityResponse struct {
	// Successful
	AtsActivity *shared.AtsActivity
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAtsActivityResponse) GetAtsActivity() *shared.AtsActivity {
	if o == nil {
		return nil
	}
	return o.AtsActivity
}

func (o *CreateAtsActivityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAtsActivityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAtsActivityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
