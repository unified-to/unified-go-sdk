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
            ApprovedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.236Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("fef67b87-0957-42c3-adcf-43c0c46d6107"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("ffa6e9c7-9193-4696-85f6-b19d35be2a3b"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("271f95a7-bbfc-403d-919c-785e59b3c427"),
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
            PostedAt: types.MustNewTimeFromString("2021-06-04T15:33:50.453Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.236Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.291Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("952638a5-c054-405a-9014-96aef0753c77"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("fa30c86b-1dc7-447f-99a5-bb2137a4e4c2"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("8b534c2b-fe41-4494-86df-02ef8117de50"),
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
            PostedAt: types.MustNewTimeFromString("2021-06-04T15:33:50.462Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.291Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.291Z"),
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://ripe-napkin.biz/"),
                    ID: unifiedgosdk.Pointer("952638a5-c054-405a-9014-96aef0753c77"),
                    MimeType: unifiedgosdk.Pointer("annus"),
                    Name: unifiedgosdk.Pointer("cohibeo"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            ExternalNumber: unifiedgosdk.Pointer("necessitatibus"),
            ID: unifiedgosdk.Pointer("fa30c86b-1dc7-447f-99a5-bb2137a4e4c2"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("8b534c2b-fe41-4494-86df-02ef8117de50"),
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
            PostedAt: types.MustNewTimeFromString("2021-06-04T15:33:50.462Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](1833.0),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](2602.0),
            TotalAmount: unifiedgosdk.Pointer[float64](3580.0),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T09:40:54.291Z"),
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