# Expense

## Overview

### Available Operations

* [CreateAccountingExpense](#createaccountingexpense) - Create an expense
* [GetAccountingExpense](#getaccountingexpense) - Retrieve an expense
* [ListAccountingExpenses](#listaccountingexpenses) - List all expenses
* [PatchAccountingExpense](#patchaccountingexpense) - Update an expense
* [RemoveAccountingExpense](#removeaccountingexpense) - Remove an expense
* [UpdateAccountingExpense](#updateaccountingexpense) - Update an expense

## CreateAccountingExpense

Create an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingExpense" method="post" path="/accounting/{connection_id}/expense" example="accounting_expense" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Expense.CreateAccountingExpense(ctx, operations.CreateAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            ApprovedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.446Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("2c89d9c9-f214-4a83-8a49-eb53567dda6e"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("1a4584bd-095a-4b17-a464-9f2b4bc15f7c"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("74fadc00-4e6f-4ef0-a00f-a7a1714a1e23"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Table featuring left technology and Rubber construction"),
                    ItemName: unifiedgosdk.Pointer("Luxurious Cotton Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-0-8324-6620-5"),
                    Notes: unifiedgosdk.Pointer("Degusto conventus defendo valetudo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](2501.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](168.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](3059.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Refined Steel Shoes"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2021-06-04T08:33:42.238Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.446Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingExpenseRequest](../../pkg/models/operations/createaccountingexpenserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingExpenseResponse](../../pkg/models/operations/createaccountingexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingExpense

Retrieve an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingExpense" method="get" path="/accounting/{connection_id}/expense/{id}" -->
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

    res, err := s.Expense.GetAccountingExpense(ctx, operations.GetAccountingExpenseRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountingExpenseRequest](../../pkg/models/operations/getaccountingexpenserequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountingExpenseResponse](../../pkg/models/operations/getaccountingexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingExpenses

List all expenses

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingExpenses" method="get" path="/accounting/{connection_id}/expense" -->
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

    res, err := s.Expense.ListAccountingExpenses(ctx, operations.ListAccountingExpensesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpenses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAccountingExpensesRequest](../../pkg/models/operations/listaccountingexpensesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAccountingExpensesResponse](../../pkg/models/operations/listaccountingexpensesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingExpense

Update an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingExpense" method="patch" path="/accounting/{connection_id}/expense/{id}" example="accounting_expense" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Expense.PatchAccountingExpense(ctx, operations.PatchAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            ApprovedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.488Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("ce6fad02-03df-4bd9-bd53-03c3d01eabfa"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("8a5cb950-617b-4246-bbfa-80448539cab4"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("df246eca-f6fa-4131-997e-73e56dfaa3c3"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Table featuring left technology and Rubber construction"),
                    ItemName: unifiedgosdk.Pointer("Luxurious Cotton Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-0-8324-6620-5"),
                    Notes: unifiedgosdk.Pointer("Degusto conventus defendo valetudo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](2501.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](168.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](3059.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Refined Steel Shoes"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2021-06-04T08:33:42.245Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.488Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAccountingExpenseRequest](../../pkg/models/operations/patchaccountingexpenserequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAccountingExpenseResponse](../../pkg/models/operations/patchaccountingexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingExpense

Remove an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingExpense" method="delete" path="/accounting/{connection_id}/expense/{id}" -->
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

    res, err := s.Expense.RemoveAccountingExpense(ctx, operations.RemoveAccountingExpenseRequest{
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
| `request`                                                                                                  | [operations.RemoveAccountingExpenseRequest](../../pkg/models/operations/removeaccountingexpenserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAccountingExpenseResponse](../../pkg/models/operations/removeaccountingexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingExpense

Update an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingExpense" method="put" path="/accounting/{connection_id}/expense/{id}" example="accounting_expense" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Expense.UpdateAccountingExpense(ctx, operations.UpdateAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            ApprovedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.488Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("ce6fad02-03df-4bd9-bd53-03c3d01eabfa"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("8a5cb950-617b-4246-bbfa-80448539cab4"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("df246eca-f6fa-4131-997e-73e56dfaa3c3"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Table featuring left technology and Rubber construction"),
                    ItemName: unifiedgosdk.Pointer("Luxurious Cotton Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-0-8324-6620-5"),
                    Notes: unifiedgosdk.Pointer("Degusto conventus defendo valetudo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](2501.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](168.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](3059.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Refined Steel Shoes"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2021-06-04T08:33:42.245Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-10T15:28:02.488Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAccountingExpenseRequest](../../pkg/models/operations/updateaccountingexpenserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAccountingExpenseResponse](../../pkg/models/operations/updateaccountingexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |