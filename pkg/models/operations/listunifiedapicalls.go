// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListUnifiedApicallsRequest struct {
	// Filter the results to just this integration's API calls
	ConnectionID *string `queryParam:"style=form,explode=true,name=connection_id"`
	// Return only results whose updated date is equal or less to this value
	CreatedLte *time.Time `queryParam:"style=form,explode=true,name=created_lte"`
	Env        *string    `queryParam:"style=form,explode=true,name=env"`
	// Filter the results for API Calls with errors
	Error *bool `queryParam:"style=form,explode=true,name=error"`
	// Filter the results to only those integrations for your user referenced by this value
	ExternalXref *string `queryParam:"style=form,explode=true,name=external_xref"`
	// Filter the results for connections with this integration
	IntegrationType *string  `queryParam:"style=form,explode=true,name=integration_type"`
	Limit           *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset          *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order           *string  `queryParam:"style=form,explode=true,name=order"`
	Sort            *string  `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListUnifiedApicallsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUnifiedApicallsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUnifiedApicallsRequest) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *ListUnifiedApicallsRequest) GetCreatedLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedLte
}

func (o *ListUnifiedApicallsRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *ListUnifiedApicallsRequest) GetError() *bool {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListUnifiedApicallsRequest) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *ListUnifiedApicallsRequest) GetIntegrationType() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationType
}

func (o *ListUnifiedApicallsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUnifiedApicallsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUnifiedApicallsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUnifiedApicallsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUnifiedApicallsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListUnifiedApicallsResponse struct {
	// Successful
	APICalls []shared.APICall
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUnifiedApicallsResponse) GetAPICalls() []shared.APICall {
	if o == nil {
		return nil
	}
	return o.APICalls
}

func (o *ListUnifiedApicallsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUnifiedApicallsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUnifiedApicallsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}