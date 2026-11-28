# Purchaseorder

## Overview

### Available Operations

* [CreateAccountingPurchaseorder](#createaccountingpurchaseorder) - Create a purchaseorder
* [GetAccountingPurchaseorder](#getaccountingpurchaseorder) - Retrieve a purchaseorder
* [ListAccountingPurchaseorders](#listaccountingpurchaseorders) - List all purchaseorders
* [PatchAccountingPurchaseorder](#patchaccountingpurchaseorder) - Update a purchaseorder
* [RemoveAccountingPurchaseorder](#removeaccountingpurchaseorder) - Remove a purchaseorder
* [UpdateAccountingPurchaseorder](#updateaccountingpurchaseorder) - Update a purchaseorder

## CreateAccountingPurchaseorder

Creates a AccountingPurchaseorder. Provide the unified AccountingPurchaseorder fields in the request body; they are mapped to the provider's API. Returns the created AccountingPurchaseorder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingPurchaseorder" method="post" path="/accounting/{connection_id}/purchaseorder" example="accounting_purchaseorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.CreateAccountingPurchaseorder(ctx, operations.CreateAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            Attachments: []shared.AccountingAttachment{},
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Golden Via"),
                Address2: unifiedgosdk.Pointer("Suite 713"),
                City: unifiedgosdk.Pointer("North Eltontown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("75366"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            ID: unifiedgosdk.Pointer("380dd529-e33b-469c-b7d7-e68c46e19959"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("9455 Brad Forks"),
                City: unifiedgosdk.Pointer("Washington"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16547"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            Status: shared.AccountingPurchaseorderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-08-15T13:12:09.603Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateAccountingPurchaseorderRequest](../../pkg/models/operations/createaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateAccountingPurchaseorderResponse](../../pkg/models/operations/createaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingPurchaseorder

Retrieves a single AccountingPurchaseorder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingPurchaseorder" method="get" path="/accounting/{connection_id}/purchaseorder/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.GetAccountingPurchaseorder(ctx, operations.GetAccountingPurchaseorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAccountingPurchaseorderRequest](../../pkg/models/operations/getaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetAccountingPurchaseorderResponse](../../pkg/models/operations/getaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingPurchaseorders

Returns a list of AccountingPurchaseorders. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingPurchaseorders" method="get" path="/accounting/{connection_id}/purchaseorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.ListAccountingPurchaseorders(ctx, operations.ListAccountingPurchaseordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListAccountingPurchaseordersRequest](../../pkg/models/operations/listaccountingpurchaseordersrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListAccountingPurchaseordersResponse](../../pkg/models/operations/listaccountingpurchaseordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingPurchaseorder

Updates the AccountingPurchaseorder identified by `id`. Provide the unified AccountingPurchaseorder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingPurchaseorder" method="patch" path="/accounting/{connection_id}/purchaseorder/{id}" example="accounting_purchaseorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.PatchAccountingPurchaseorder(ctx, operations.PatchAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            Attachments: []shared.AccountingAttachment{},
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Golden Via"),
                Address2: unifiedgosdk.Pointer("Suite 713"),
                City: unifiedgosdk.Pointer("North Eltontown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("75366"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            ID: unifiedgosdk.Pointer("b5b6f8ef-449e-46a6-becf-9775ead44411"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("9455 Brad Forks"),
                City: unifiedgosdk.Pointer("Washington"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16547"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            Status: shared.AccountingPurchaseorderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-08-15T13:12:09.606Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchAccountingPurchaseorderRequest](../../pkg/models/operations/patchaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchAccountingPurchaseorderResponse](../../pkg/models/operations/patchaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingPurchaseorder

Deletes the AccountingPurchaseorder identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingPurchaseorder" method="delete" path="/accounting/{connection_id}/purchaseorder/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.RemoveAccountingPurchaseorder(ctx, operations.RemoveAccountingPurchaseorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.RemoveAccountingPurchaseorderRequest](../../pkg/models/operations/removeaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.RemoveAccountingPurchaseorderResponse](../../pkg/models/operations/removeaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingPurchaseorder

Updates the AccountingPurchaseorder identified by `id`. Provide the unified AccountingPurchaseorder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingPurchaseorder" method="put" path="/accounting/{connection_id}/purchaseorder/{id}" example="accounting_purchaseorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Purchaseorder.UpdateAccountingPurchaseorder(ctx, operations.UpdateAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            Attachments: []shared.AccountingAttachment{},
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Golden Via"),
                Address2: unifiedgosdk.Pointer("Suite 713"),
                City: unifiedgosdk.Pointer("North Eltontown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("75366"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            ID: unifiedgosdk.Pointer("b5b6f8ef-449e-46a6-becf-9775ead44411"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("9455 Brad Forks"),
                City: unifiedgosdk.Pointer("Washington"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16547"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            Status: shared.AccountingPurchaseorderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-08-15T13:12:09.606Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateAccountingPurchaseorderRequest](../../pkg/models/operations/updateaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateAccountingPurchaseorderResponse](../../pkg/models/operations/updateaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |