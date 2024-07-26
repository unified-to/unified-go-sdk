// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListAccountingJournalsRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListAccountingJournalsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingJournalsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingJournalsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListAccountingJournalsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListAccountingJournalsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingJournalsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountingJournalsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListAccountingJournalsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListAccountingJournalsResponse struct {
	// Successful
	AccountingJournals []shared.AccountingJournal
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountingJournalsResponse) GetAccountingJournals() []shared.AccountingJournal {
	if o == nil {
		return nil
	}
	return o.AccountingJournals
}

func (o *ListAccountingJournalsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountingJournalsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountingJournalsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
