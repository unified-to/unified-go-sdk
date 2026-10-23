# Taxrate

## Overview

### Available Operations

* [CreateAccountingTaxrate](#createaccountingtaxrate) - Create a taxrate
* [GetAccountingTaxrate](#getaccountingtaxrate) - Retrieve a taxrate
* [ListAccountingTaxrates](#listaccountingtaxrates) - List all taxrates
* [PatchAccountingTaxrate](#patchaccountingtaxrate) - Update a taxrate
* [RemoveAccountingTaxrate](#removeaccountingtaxrate) - Remove a taxrate
* [UpdateAccountingTaxrate](#updateaccountingtaxrate) - Update a taxrate

## CreateAccountingTaxrate

Creates a AccountingTaxrate. Provide the unified AccountingTaxrate fields in the request body; they are mapped to the provider's API. Returns the created AccountingTaxrate.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingTaxrate" method="post" path="/accounting/{connection_id}/taxrate" example="accounting_taxrate" -->
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

    res, err := s.Taxrate.CreateAccountingTaxrate(ctx, operations.CreateAccountingTaxrateRequest{
        AccountingTaxrate: shared.AccountingTaxrate{
            Components: []shared.AccountingTaxrateComponent{
                shared.AccountingTaxrateComponent{
                    IsCompound: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Exempt"),
                    Rate: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Exempt (0%)"),
            EffectiveRate: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("6c2902dd-2610-46bb-b2c1-17bbb9e5c82c"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Exempt"),
            Rate: unifiedgosdk.Pointer[float64](0.0),
            Region: unifiedgosdk.Pointer("KS"),
            TotalRate: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingTaxrateTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-13T05:56:57.458Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTaxrate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingTaxrateRequest](../../pkg/models/operations/createaccountingtaxraterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingTaxrateResponse](../../pkg/models/operations/createaccountingtaxrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingTaxrate

Retrieves a single AccountingTaxrate by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingTaxrate" method="get" path="/accounting/{connection_id}/taxrate/{id}" -->
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

    res, err := s.Taxrate.GetAccountingTaxrate(ctx, operations.GetAccountingTaxrateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTaxrate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountingTaxrateRequest](../../pkg/models/operations/getaccountingtaxraterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountingTaxrateResponse](../../pkg/models/operations/getaccountingtaxrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingTaxrates

Returns a list of AccountingTaxrates. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingTaxrates" method="get" path="/accounting/{connection_id}/taxrate" -->
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

    res, err := s.Taxrate.ListAccountingTaxrates(ctx, operations.ListAccountingTaxratesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTaxrates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAccountingTaxratesRequest](../../pkg/models/operations/listaccountingtaxratesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAccountingTaxratesResponse](../../pkg/models/operations/listaccountingtaxratesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingTaxrate

Updates the AccountingTaxrate identified by `id`. Provide the unified AccountingTaxrate fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingTaxrate" method="patch" path="/accounting/{connection_id}/taxrate/{id}" example="accounting_taxrate" -->
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

    res, err := s.Taxrate.PatchAccountingTaxrate(ctx, operations.PatchAccountingTaxrateRequest{
        AccountingTaxrate: shared.AccountingTaxrate{
            Components: []shared.AccountingTaxrateComponent{
                shared.AccountingTaxrateComponent{
                    IsCompound: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Exempt"),
                    Rate: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Exempt (0%)"),
            EffectiveRate: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("cd5dc689-498f-4ed0-9f7f-03897b6a3a05"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Exempt"),
            Rate: unifiedgosdk.Pointer[float64](0.0),
            Region: unifiedgosdk.Pointer("KS"),
            TotalRate: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingTaxrateTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-13T05:56:57.458Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTaxrate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAccountingTaxrateRequest](../../pkg/models/operations/patchaccountingtaxraterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAccountingTaxrateResponse](../../pkg/models/operations/patchaccountingtaxrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingTaxrate

Deletes the AccountingTaxrate identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingTaxrate" method="delete" path="/accounting/{connection_id}/taxrate/{id}" -->
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

    res, err := s.Taxrate.RemoveAccountingTaxrate(ctx, operations.RemoveAccountingTaxrateRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAccountingTaxrateRequest](../../pkg/models/operations/removeaccountingtaxraterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAccountingTaxrateResponse](../../pkg/models/operations/removeaccountingtaxrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingTaxrate

Updates the AccountingTaxrate identified by `id`. Provide the unified AccountingTaxrate fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingTaxrate" method="put" path="/accounting/{connection_id}/taxrate/{id}" example="accounting_taxrate" -->
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

    res, err := s.Taxrate.UpdateAccountingTaxrate(ctx, operations.UpdateAccountingTaxrateRequest{
        AccountingTaxrate: shared.AccountingTaxrate{
            Components: []shared.AccountingTaxrateComponent{
                shared.AccountingTaxrateComponent{
                    IsCompound: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Exempt"),
                    Rate: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Exempt (0%)"),
            EffectiveRate: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("cd5dc689-498f-4ed0-9f7f-03897b6a3a05"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Exempt"),
            Rate: unifiedgosdk.Pointer[float64](0.0),
            Region: unifiedgosdk.Pointer("KS"),
            TotalRate: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingTaxrateTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-13T05:56:57.458Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTaxrate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAccountingTaxrateRequest](../../pkg/models/operations/updateaccountingtaxraterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAccountingTaxrateResponse](../../pkg/models/operations/updateaccountingtaxrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |