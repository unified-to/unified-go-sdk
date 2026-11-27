# Expense

## Overview

### Available Operations

* [CreateAccountingExpense](#createaccountingexpense) - Create an expense
* [CreateLegalExpense](#createlegalexpense) - Create an expense
* [GetAccountingExpense](#getaccountingexpense) - Retrieve an expense
* [GetLegalExpense](#getlegalexpense) - Retrieve an expense
* [ListAccountingExpenses](#listaccountingexpenses) - List all expenses
* [ListLegalExpenses](#listlegalexpenses) - List all expenses
* [PatchAccountingExpense](#patchaccountingexpense) - Update an expense
* [PatchLegalExpense](#patchlegalexpense) - Update an expense
* [RemoveAccountingExpense](#removeaccountingexpense) - Remove an expense
* [RemoveLegalExpense](#removelegalexpense) - Remove an expense
* [UpdateAccountingExpense](#updateaccountingexpense) - Update an expense
* [UpdateLegalExpense](#updatelegalexpense) - Update an expense

## CreateAccountingExpense

Creates a AccountingExpense. Provide the unified AccountingExpense fields in the request body; they are mapped to the provider's API. Returns the created AccountingExpense.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingExpense" method="post" path="/accounting/{connection_id}/expense" example="accounting_expense" -->
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

    res, err := s.Expense.CreateAccountingExpense(ctx, operations.CreateAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://avaricious-larva.biz"),
                    ID: unifiedgosdk.Pointer("aae567bb-4651-46a9-befd-90d3ee90897a"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://precious-comparison.net/"),
                    ID: unifiedgosdk.Pointer("478b63f8-6e7d-4653-9139-4896bdc5e58d"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("aa33e4e0-1bbc-4b2e-a6df-2efb3ec17878"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("92d5a6d8-37d1-4a10-9550-4a5a31df1645"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("efa96262-a202-456a-81a4-ab5c4a6ddb5d"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("8117f94d-1246-44be-a111-2da737b73cf0"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Sint spiculum carcer ago charisma."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f4a7116e-86c0-4768-a1aa-8209a732eafb"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f61a01f9-16a8-413f-95f3-4666a23fa0dd"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A deposit for KPW 867.85 was made at Kohler LLC via card ending ****1416 on account ***1040."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-03T15:05:18.801Z"),
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

## CreateLegalExpense

Creates a LegalExpense. Provide the unified LegalExpense fields in the request body; they are mapped to the provider's API. Returns the created LegalExpense.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalExpense" method="post" path="/legal/{connection_id}/expense" example="legal_expense" -->
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

    res, err := s.Expense.CreateLegalExpense(ctx, operations.CreateLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Baby"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Nihil sono villa voluptate cohibeo quas volo succurro."),
            ID: unifiedgosdk.Pointer("4efb6add-1e94-4d16-8540-101fc2177b5f"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-21T01:52:52.384Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLegalExpenseRequest](../../pkg/models/operations/createlegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateLegalExpenseResponse](../../pkg/models/operations/createlegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingExpense

Retrieves a single AccountingExpense by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

## GetLegalExpense

Retrieves a single LegalExpense by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalExpense" method="get" path="/legal/{connection_id}/expense/{id}" -->
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

    res, err := s.Expense.GetLegalExpense(ctx, operations.GetLegalExpenseRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLegalExpenseRequest](../../pkg/models/operations/getlegalexpenserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetLegalExpenseResponse](../../pkg/models/operations/getlegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingExpenses

Returns a list of AccountingExpenses. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

## ListLegalExpenses

Returns a list of LegalExpenses. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalExpenses" method="get" path="/legal/{connection_id}/expense" -->
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

    res, err := s.Expense.ListLegalExpenses(ctx, operations.ListLegalExpensesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpenses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLegalExpensesRequest](../../pkg/models/operations/listlegalexpensesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListLegalExpensesResponse](../../pkg/models/operations/listlegalexpensesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingExpense

Updates the AccountingExpense identified by `id`. Provide the unified AccountingExpense fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingExpense" method="patch" path="/accounting/{connection_id}/expense/{id}" example="accounting_expense" -->
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

    res, err := s.Expense.PatchAccountingExpense(ctx, operations.PatchAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://avaricious-larva.biz"),
                    ID: unifiedgosdk.Pointer("8db790d1-6549-41a3-8027-86ca477ed1ea"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://precious-comparison.net/"),
                    ID: unifiedgosdk.Pointer("a8233b4f-ce05-4f9b-9f4c-d36f7531d4d9"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("1dc07e6a-001d-4741-a966-ca0d3e14cbec"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("cef9117b-0a99-4232-a338-ab19cabc46fd"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("b0248872-f10c-4b4a-933b-94e5936726b6"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("dae87d72-2b7e-4fdb-9256-3e4b85df289a"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Sint spiculum carcer ago charisma."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f062d7ae-1fe0-4136-94c0-710cc14d56cc"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("df549aca-b118-4196-9ed5-b8083323f745"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A deposit for KPW 867.85 was made at Kohler LLC via card ending ****1416 on account ***1040."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-03T15:05:18.802Z"),
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

## PatchLegalExpense

Updates the LegalExpense identified by `id`. Provide the unified LegalExpense fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalExpense" method="patch" path="/legal/{connection_id}/expense/{id}" example="legal_expense" -->
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

    res, err := s.Expense.PatchLegalExpense(ctx, operations.PatchLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Baby"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Nihil sono villa voluptate cohibeo quas volo succurro."),
            ID: unifiedgosdk.Pointer("94ad75f6-af50-4d13-b430-f07500eda8f9"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-21T01:52:52.392Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchLegalExpenseRequest](../../pkg/models/operations/patchlegalexpenserequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchLegalExpenseResponse](../../pkg/models/operations/patchlegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingExpense

Deletes the AccountingExpense identified by `id`.

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

## RemoveLegalExpense

Deletes the LegalExpense identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalExpense" method="delete" path="/legal/{connection_id}/expense/{id}" -->
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

    res, err := s.Expense.RemoveLegalExpense(ctx, operations.RemoveLegalExpenseRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveLegalExpenseRequest](../../pkg/models/operations/removelegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveLegalExpenseResponse](../../pkg/models/operations/removelegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingExpense

Updates the AccountingExpense identified by `id`. Provide the unified AccountingExpense fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingExpense" method="put" path="/accounting/{connection_id}/expense/{id}" example="accounting_expense" -->
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

    res, err := s.Expense.UpdateAccountingExpense(ctx, operations.UpdateAccountingExpenseRequest{
        AccountingExpense: shared.AccountingExpense{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://avaricious-larva.biz"),
                    ID: unifiedgosdk.Pointer("8db790d1-6549-41a3-8027-86ca477ed1ea"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://precious-comparison.net/"),
                    ID: unifiedgosdk.Pointer("a8233b4f-ce05-4f9b-9f4c-d36f7531d4d9"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("1dc07e6a-001d-4741-a966-ca0d3e14cbec"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("cef9117b-0a99-4232-a338-ab19cabc46fd"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("b0248872-f10c-4b4a-933b-94e5936726b6"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("dae87d72-2b7e-4fdb-9256-3e4b85df289a"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Sint spiculum carcer ago charisma."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f062d7ae-1fe0-4136-94c0-710cc14d56cc"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("df549aca-b118-4196-9ed5-b8083323f745"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A deposit for KPW 867.85 was made at Kohler LLC via card ending ****1416 on account ***1040."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-03T15:05:18.802Z"),
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

## UpdateLegalExpense

Updates the LegalExpense identified by `id`. Provide the unified LegalExpense fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalExpense" method="put" path="/legal/{connection_id}/expense/{id}" example="legal_expense" -->
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

    res, err := s.Expense.UpdateLegalExpense(ctx, operations.UpdateLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Baby"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Nihil sono villa voluptate cohibeo quas volo succurro."),
            ID: unifiedgosdk.Pointer("94ad75f6-af50-4d13-b430-f07500eda8f9"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-21T01:52:52.392Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLegalExpenseRequest](../../pkg/models/operations/updatelegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateLegalExpenseResponse](../../pkg/models/operations/updatelegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |