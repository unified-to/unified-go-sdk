# Vendorcredit

## Overview

### Available Operations

* [CreateAccountingVendorcredit](#createaccountingvendorcredit) - Create a vendorcredit
* [GetAccountingVendorcredit](#getaccountingvendorcredit) - Retrieve a vendorcredit
* [ListAccountingVendorcredits](#listaccountingvendorcredits) - List all vendorcredits
* [PatchAccountingVendorcredit](#patchaccountingvendorcredit) - Update a vendorcredit
* [RemoveAccountingVendorcredit](#removeaccountingvendorcredit) - Remove a vendorcredit
* [UpdateAccountingVendorcredit](#updateaccountingvendorcredit) - Update a vendorcredit

## CreateAccountingVendorcredit

Create a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingVendorcredit" method="post" path="/accounting/{connection_id}/vendorcredit" example="accounting_vendorcredit" -->
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

    res, err := s.Vendorcredit.CreateAccountingVendorcredit(ctx, operations.CreateAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{
            Applications: []shared.AccountingCreditApplication{},
            ApplyAmount: unifiedgosdk.Pointer[float64](1.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-04-15T21:14:08.197Z"),
            Currency: unifiedgosdk.Pointer("KGS"),
            DueAt: types.MustNewTimeFromString("2023-05-06T20:38:46.775Z"),
            ID: unifiedgosdk.Pointer("b6823602-0b54-43b4-aed9-7f713d059bca"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Conatus cruciamentum decor avaritia tantum."),
            PostedAt: types.MustNewTimeFromString("2023-09-28T19:41:28.758Z"),
            Status: shared.AccountingVendorcreditStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-11-26T18:36:43.880Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateAccountingVendorcreditRequest](../../pkg/models/operations/createaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CreateAccountingVendorcreditResponse](../../pkg/models/operations/createaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingVendorcredit

Retrieve a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingVendorcredit" method="get" path="/accounting/{connection_id}/vendorcredit/{id}" -->
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

    res, err := s.Vendorcredit.GetAccountingVendorcredit(ctx, operations.GetAccountingVendorcreditRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAccountingVendorcreditRequest](../../pkg/models/operations/getaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetAccountingVendorcreditResponse](../../pkg/models/operations/getaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingVendorcredits

List all vendorcredits

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingVendorcredits" method="get" path="/accounting/{connection_id}/vendorcredit" -->
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

    res, err := s.Vendorcredit.ListAccountingVendorcredits(ctx, operations.ListAccountingVendorcreditsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListAccountingVendorcreditsRequest](../../pkg/models/operations/listaccountingvendorcreditsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListAccountingVendorcreditsResponse](../../pkg/models/operations/listaccountingvendorcreditsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingVendorcredit

Update a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingVendorcredit" method="patch" path="/accounting/{connection_id}/vendorcredit/{id}" example="accounting_vendorcredit" -->
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

    res, err := s.Vendorcredit.PatchAccountingVendorcredit(ctx, operations.PatchAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{
            Applications: []shared.AccountingCreditApplication{},
            ApplyAmount: unifiedgosdk.Pointer[float64](1.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-04-15T21:14:08.197Z"),
            Currency: unifiedgosdk.Pointer("KGS"),
            DueAt: types.MustNewTimeFromString("2023-05-06T20:38:46.775Z"),
            ID: unifiedgosdk.Pointer("7ae560d6-73a0-4e00-abb7-9718c6933434"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Conatus cruciamentum decor avaritia tantum."),
            PostedAt: types.MustNewTimeFromString("2023-09-28T19:41:28.763Z"),
            Status: shared.AccountingVendorcreditStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-11-26T18:36:43.887Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchAccountingVendorcreditRequest](../../pkg/models/operations/patchaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PatchAccountingVendorcreditResponse](../../pkg/models/operations/patchaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingVendorcredit

Remove a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingVendorcredit" method="delete" path="/accounting/{connection_id}/vendorcredit/{id}" -->
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

    res, err := s.Vendorcredit.RemoveAccountingVendorcredit(ctx, operations.RemoveAccountingVendorcreditRequest{
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.RemoveAccountingVendorcreditRequest](../../pkg/models/operations/removeaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.RemoveAccountingVendorcreditResponse](../../pkg/models/operations/removeaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingVendorcredit

Update a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingVendorcredit" method="put" path="/accounting/{connection_id}/vendorcredit/{id}" example="accounting_vendorcredit" -->
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

    res, err := s.Vendorcredit.UpdateAccountingVendorcredit(ctx, operations.UpdateAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{
            Applications: []shared.AccountingCreditApplication{},
            ApplyAmount: unifiedgosdk.Pointer[float64](1.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-04-15T21:14:08.197Z"),
            Currency: unifiedgosdk.Pointer("KGS"),
            DueAt: types.MustNewTimeFromString("2023-05-06T20:38:46.775Z"),
            ID: unifiedgosdk.Pointer("7ae560d6-73a0-4e00-abb7-9718c6933434"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Conatus cruciamentum decor avaritia tantum."),
            PostedAt: types.MustNewTimeFromString("2023-09-28T19:41:28.763Z"),
            Status: shared.AccountingVendorcreditStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-11-26T18:36:43.887Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.UpdateAccountingVendorcreditRequest](../../pkg/models/operations/updateaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpdateAccountingVendorcreditResponse](../../pkg/models/operations/updateaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |