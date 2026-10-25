# Paymentterm

## Overview

### Available Operations

* [CreateAccountingPaymentterm](#createaccountingpaymentterm) - Create a paymentterm
* [GetAccountingPaymentterm](#getaccountingpaymentterm) - Retrieve a paymentterm
* [ListAccountingPaymentterms](#listaccountingpaymentterms) - List all paymentterms
* [PatchAccountingPaymentterm](#patchaccountingpaymentterm) - Update a paymentterm
* [RemoveAccountingPaymentterm](#removeaccountingpaymentterm) - Remove a paymentterm
* [UpdateAccountingPaymentterm](#updateaccountingpaymentterm) - Update a paymentterm

## CreateAccountingPaymentterm

Creates a AccountingPaymentterm. Provide the unified AccountingPaymentterm fields in the request body; they are mapped to the provider's API. Returns the created AccountingPaymentterm.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingPaymentterm" method="post" path="/accounting/{connection_id}/paymentterm" example="accounting_paymentterm" -->
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

    res, err := s.Paymentterm.CreateAccountingPaymentterm(ctx, operations.CreateAccountingPaymenttermRequest{
        AccountingPaymentterm: shared.AccountingPaymentterm{
            Category: shared.CategoryStandard.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2021-08-22T22:42:42.265Z"),
            DayOfMonthDue: unifiedgosdk.Pointer[float64](4.0),
            Description: unifiedgosdk.Pointer("Cogito pecco eos cultura."),
            DiscountDayOfMonth: unifiedgosdk.Pointer[float64](13.0),
            DiscountDays: unifiedgosdk.Pointer[float64](4.0),
            DiscountPercent: unifiedgosdk.Pointer[float64](5.0),
            DueDays: unifiedgosdk.Pointer[float64](57.0),
            DueNextMonthDays: unifiedgosdk.Pointer[float64](9.0),
            ID: unifiedgosdk.Pointer("f3ade6d3-f270-416c-8368-af85cca0150c"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-11T17:51:46.725Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPaymentterm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateAccountingPaymenttermRequest](../../pkg/models/operations/createaccountingpaymenttermrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateAccountingPaymenttermResponse](../../pkg/models/operations/createaccountingpaymenttermresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingPaymentterm

Retrieves a single AccountingPaymentterm by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingPaymentterm" method="get" path="/accounting/{connection_id}/paymentterm/{id}" -->
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

    res, err := s.Paymentterm.GetAccountingPaymentterm(ctx, operations.GetAccountingPaymenttermRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPaymentterm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAccountingPaymenttermRequest](../../pkg/models/operations/getaccountingpaymenttermrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetAccountingPaymenttermResponse](../../pkg/models/operations/getaccountingpaymenttermresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingPaymentterms

Returns a list of AccountingPaymentterms. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingPaymentterms" method="get" path="/accounting/{connection_id}/paymentterm" -->
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

    res, err := s.Paymentterm.ListAccountingPaymentterms(ctx, operations.ListAccountingPaymenttermsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPaymentterms != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListAccountingPaymenttermsRequest](../../pkg/models/operations/listaccountingpaymenttermsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListAccountingPaymenttermsResponse](../../pkg/models/operations/listaccountingpaymenttermsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingPaymentterm

Updates the AccountingPaymentterm identified by `id`. Provide the unified AccountingPaymentterm fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingPaymentterm" method="patch" path="/accounting/{connection_id}/paymentterm/{id}" example="accounting_paymentterm" -->
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

    res, err := s.Paymentterm.PatchAccountingPaymentterm(ctx, operations.PatchAccountingPaymenttermRequest{
        AccountingPaymentterm: shared.AccountingPaymentterm{
            Category: shared.CategoryStandard.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2021-08-22T22:42:42.265Z"),
            DayOfMonthDue: unifiedgosdk.Pointer[float64](4.0),
            Description: unifiedgosdk.Pointer("Cogito pecco eos cultura."),
            DiscountDayOfMonth: unifiedgosdk.Pointer[float64](13.0),
            DiscountDays: unifiedgosdk.Pointer[float64](4.0),
            DiscountPercent: unifiedgosdk.Pointer[float64](5.0),
            DueDays: unifiedgosdk.Pointer[float64](57.0),
            DueNextMonthDays: unifiedgosdk.Pointer[float64](9.0),
            ID: unifiedgosdk.Pointer("6f9b3d95-318a-462a-aec1-f9d19c7789ef"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-11T17:51:46.736Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPaymentterm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchAccountingPaymenttermRequest](../../pkg/models/operations/patchaccountingpaymenttermrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PatchAccountingPaymenttermResponse](../../pkg/models/operations/patchaccountingpaymenttermresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingPaymentterm

Deletes the AccountingPaymentterm identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingPaymentterm" method="delete" path="/accounting/{connection_id}/paymentterm/{id}" -->
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

    res, err := s.Paymentterm.RemoveAccountingPaymentterm(ctx, operations.RemoveAccountingPaymenttermRequest{
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.RemoveAccountingPaymenttermRequest](../../pkg/models/operations/removeaccountingpaymenttermrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.RemoveAccountingPaymenttermResponse](../../pkg/models/operations/removeaccountingpaymenttermresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingPaymentterm

Updates the AccountingPaymentterm identified by `id`. Provide the unified AccountingPaymentterm fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingPaymentterm" method="put" path="/accounting/{connection_id}/paymentterm/{id}" example="accounting_paymentterm" -->
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

    res, err := s.Paymentterm.UpdateAccountingPaymentterm(ctx, operations.UpdateAccountingPaymenttermRequest{
        AccountingPaymentterm: shared.AccountingPaymentterm{
            Category: shared.CategoryStandard.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2021-08-22T22:42:42.265Z"),
            DayOfMonthDue: unifiedgosdk.Pointer[float64](4.0),
            Description: unifiedgosdk.Pointer("Cogito pecco eos cultura."),
            DiscountDayOfMonth: unifiedgosdk.Pointer[float64](13.0),
            DiscountDays: unifiedgosdk.Pointer[float64](4.0),
            DiscountPercent: unifiedgosdk.Pointer[float64](5.0),
            DueDays: unifiedgosdk.Pointer[float64](57.0),
            DueNextMonthDays: unifiedgosdk.Pointer[float64](9.0),
            ID: unifiedgosdk.Pointer("6f9b3d95-318a-462a-aec1-f9d19c7789ef"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-11T17:51:46.736Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPaymentterm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateAccountingPaymenttermRequest](../../pkg/models/operations/updateaccountingpaymenttermrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateAccountingPaymenttermResponse](../../pkg/models/operations/updateaccountingpaymenttermresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |