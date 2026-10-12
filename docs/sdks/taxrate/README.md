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

Create a taxrate

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
                    Name: unifiedgosdk.Pointer("cogo"),
                    Rate: unifiedgosdk.Pointer[float64](53.0),
                },
            },
            Country: unifiedgosdk.Pointer("BY"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Atrox tricesimus creator."),
            EffectiveRate: unifiedgosdk.Pointer[float64](53.0),
            ID: unifiedgosdk.Pointer("1b6ac909-f331-4923-8deb-9a2421f46a1a"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(true),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("villa"),
            Rate: unifiedgosdk.Pointer[float64](53.0),
            Region: unifiedgosdk.Pointer("TX"),
            TotalRate: unifiedgosdk.Pointer[float64](53.0),
            Type: shared.AccountingTaxrateTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-04-20T11:37:13.825Z"),
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

Retrieve a taxrate

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

List all taxrates

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

Update a taxrate

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
                    Name: unifiedgosdk.Pointer("cogo"),
                    Rate: unifiedgosdk.Pointer[float64](53.0),
                },
            },
            Country: unifiedgosdk.Pointer("BY"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Atrox tricesimus creator."),
            EffectiveRate: unifiedgosdk.Pointer[float64](53.0),
            ID: unifiedgosdk.Pointer("7ab7fe23-6eee-4cb3-8359-913d3d4e042f"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(true),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("villa"),
            Rate: unifiedgosdk.Pointer[float64](53.0),
            Region: unifiedgosdk.Pointer("TX"),
            TotalRate: unifiedgosdk.Pointer[float64](53.0),
            Type: shared.AccountingTaxrateTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-04-20T11:37:13.832Z"),
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

Remove a taxrate

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

Update a taxrate

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
                    Name: unifiedgosdk.Pointer("cogo"),
                    Rate: unifiedgosdk.Pointer[float64](53.0),
                },
            },
            Country: unifiedgosdk.Pointer("BY"),
            CreatedAt: types.MustNewTimeFromString("2020-01-04T10:37:56.894Z"),
            Description: unifiedgosdk.Pointer("Atrox tricesimus creator."),
            EffectiveRate: unifiedgosdk.Pointer[float64](53.0),
            ID: unifiedgosdk.Pointer("7ab7fe23-6eee-4cb3-8359-913d3d4e042f"),
            IsActive: unifiedgosdk.Pointer(false),
            IsCompound: unifiedgosdk.Pointer(true),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("villa"),
            Rate: unifiedgosdk.Pointer[float64](53.0),
            Region: unifiedgosdk.Pointer("TX"),
            TotalRate: unifiedgosdk.Pointer[float64](53.0),
            Type: shared.AccountingTaxrateTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-04-20T11:37:13.832Z"),
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