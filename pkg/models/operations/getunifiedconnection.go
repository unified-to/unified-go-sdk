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

type GetUnifiedConnectionCategories string

const (
	GetUnifiedConnectionCategoriesPassthrough GetUnifiedConnectionCategories = "passthrough"
	GetUnifiedConnectionCategoriesHris        GetUnifiedConnectionCategories = "hris"
	GetUnifiedConnectionCategoriesAts         GetUnifiedConnectionCategories = "ats"
	GetUnifiedConnectionCategoriesAuth        GetUnifiedConnectionCategories = "auth"
	GetUnifiedConnectionCategoriesCrm         GetUnifiedConnectionCategories = "crm"
	GetUnifiedConnectionCategoriesEnrich      GetUnifiedConnectionCategories = "enrich"
	GetUnifiedConnectionCategoriesMartech     GetUnifiedConnectionCategories = "martech"
	GetUnifiedConnectionCategoriesTicketing   GetUnifiedConnectionCategories = "ticketing"
	GetUnifiedConnectionCategoriesUc          GetUnifiedConnectionCategories = "uc"
)

func (e GetUnifiedConnectionCategories) ToPointer() *GetUnifiedConnectionCategories {
	return &e
}

func (e *GetUnifiedConnectionCategories) UnmarshalJSON(data []byte) error {
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
		*e = GetUnifiedConnectionCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetUnifiedConnectionCategories: %v", v)
	}
}

type GetUnifiedConnectionRequest struct {
	// Filter the results on these categories
	Categories []GetUnifiedConnectionCategories `queryParam:"style=form,explode=true,name=categories"`
	Env        *string                          `queryParam:"style=form,explode=true,name=env"`
	// Filter the results to only those integrations for your user referenced by this value
	ExternalXref *string  `queryParam:"style=form,explode=true,name=external_xref"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	Sort         *string  `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (g GetUnifiedConnectionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetUnifiedConnectionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetUnifiedConnectionRequest) GetCategories() []GetUnifiedConnectionCategories {
	if o == nil {
		return nil
	}
	return o.Categories
}

func (o *GetUnifiedConnectionRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *GetUnifiedConnectionRequest) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *GetUnifiedConnectionRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUnifiedConnectionRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUnifiedConnectionRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetUnifiedConnectionRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetUnifiedConnectionRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetUnifiedConnectionResponse struct {
	// Successful
	Connections []shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetUnifiedConnectionResponse) GetConnections() []shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *GetUnifiedConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
