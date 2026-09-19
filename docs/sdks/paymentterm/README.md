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

Create a paymentterm

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
            ID: unifiedgosdk.Pointer("99ffcec4-46cb-431c-96c8-c6fcb9d51288"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T06:05:16.527Z"),
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

Retrieve a paymentterm

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

List all paymentterms

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

Update a paymentterm

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
            ID: unifiedgosdk.Pointer("d4f9a147-86b6-4b4e-8e39-ff8011ffcf83"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T06:05:16.542Z"),
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

Remove a paymentterm

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

Update a paymentterm

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
            ID: unifiedgosdk.Pointer("d4f9a147-86b6-4b4e-8e39-ff8011ffcf83"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Net 30"),
            Type: shared.AccountingPaymenttermTypeNet15.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T06:05:16.542Z"),
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