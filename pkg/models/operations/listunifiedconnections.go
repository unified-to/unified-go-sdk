// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type Categories string

const (
	CategoriesPassthrough Categories = "passthrough"
	CategoriesHris        Categories = "hris"
	CategoriesAts         Categories = "ats"
	CategoriesAuth        Categories = "auth"
	CategoriesCrm         Categories = "crm"
	CategoriesEnrich      Categories = "enrich"
	CategoriesMartech     Categories = "martech"
	CategoriesTicketing   Categories = "ticketing"
	CategoriesUc          Categories = "uc"
	CategoriesAccounting  Categories = "accounting"
	CategoriesStorage     Categories = "storage"
	CategoriesCommerce    Categories = "commerce"
	CategoriesPayment     Categories = "payment"
	CategoriesGenai       Categories = "genai"
)

func (e Categories) ToPointer() *Categories {
	return &e
}

func (e *Categories) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "passthrough":
		fallthrough
	case "hris":
		fallthrough
	case "ats":
		fallthrough
	case "auth":
		fallthrough
	case "crm":
		fallthrough
	case "enrich":
		fallthrough
	case "martech":
		fallthrough
	case "ticketing":
		fallthrough
	case "uc":
		fallthrough
	case "accounting":
		fallthrough
	case "storage":
		fallthrough
	case "commerce":
		fallthrough
	case "payment":
		fallthrough
	case "genai":
		*e = Categories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Categories: %v", v)
	}
}

type ListUnifiedConnectionsRequest struct {
	// Filter the results on these categories
	Categories []Categories `queryParam:"style=form,explode=true,name=categories"`
	Env        *string      `queryParam:"style=form,explode=true,name=env"`
	// Filter the results to only those integrations for your user referenced by this value
	ExternalXref *string  `queryParam:"style=form,explode=true,name=external_xref"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	Sort         *string  `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListUnifiedConnectionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUnifiedConnectionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUnifiedConnectionsRequest) GetCategories() []Categories {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *ListUnifiedConnectionsRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *ListUnifiedConnectionsRequest) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *ListUnifiedConnectionsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUnifiedConnectionsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUnifiedConnectionsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUnifiedConnectionsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUnifiedConnectionsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListUnifiedConnectionsResponse struct {
	// Successful
	Connections []shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUnifiedConnectionsResponse) GetConnections() []shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *ListUnifiedConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUnifiedConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUnifiedConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
