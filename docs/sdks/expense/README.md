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
                    DownloadURL: unifiedgosdk.Pointer("https://rubbery-couch.biz/"),
                    ID: unifiedgosdk.Pointer("9aabf60d-0097-4632-aaf4-994d747eae41"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://appropriate-smog.biz/"),
                    ID: unifiedgosdk.Pointer("e86d99a2-10c8-4f17-9bad-f9d328ac4f39"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("fc5b3f57-18bd-4021-a137-4246201f6077"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("d2e7ed56-b5e3-45c7-9fa6-e0e7a5a555a3"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a74ba4b7-02a3-4d2c-aec0-9ae6bb4ec6cf"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("04917fce-3562-48b4-b20b-868a6fefebf3"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Casus umquam conicio cinis demitto sint."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("3219eb0f-659b-491e-bc64-9d497ef1055d"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("217d79c7-ec0c-4226-9351-944aa1f6ef6f"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A payment of XPF 714.84 occurred at Bins, Barrows and Koch using a card ending in ****8422 for account ***4614."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-02T08:08:15.486Z"),
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
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("21864e35-8eca-4e7f-81cc-e81c04f58bf5"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-31T12:54:25.595Z"),
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
                    DownloadURL: unifiedgosdk.Pointer("https://rubbery-couch.biz/"),
                    ID: unifiedgosdk.Pointer("d58b40c5-3bab-4eb8-8953-45045eb60e9f"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://appropriate-smog.biz/"),
                    ID: unifiedgosdk.Pointer("77647132-cdbd-485d-a2c5-c2ea5ce89131"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("542e253e-9be2-4769-9e7b-513144533652"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("077d67ff-ee2e-4c5f-ac4f-bb6481113a38"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("0906ed39-7dae-456f-9c82-077543591acd"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f4ae5862-e96c-4dbc-a116-2f9d71403708"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Casus umquam conicio cinis demitto sint."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("9ad8a271-29b8-4c2f-a966-e56e1d5fa671"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("e783cbf5-5d97-4777-8798-596b0362a0c4"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A payment of XPF 714.84 occurred at Bins, Barrows and Koch using a card ending in ****8422 for account ***4614."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-02T08:08:15.489Z"),
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
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("5b3cb8a3-5737-4b18-b126-5266526343a7"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-31T12:54:25.603Z"),
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
                    DownloadURL: unifiedgosdk.Pointer("https://rubbery-couch.biz/"),
                    ID: unifiedgosdk.Pointer("d58b40c5-3bab-4eb8-8953-45045eb60e9f"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://appropriate-smog.biz/"),
                    ID: unifiedgosdk.Pointer("77647132-cdbd-485d-a2c5-c2ea5ce89131"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("receipt-9696-1.pdf"),
                },
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("EXP-38303"),
            ID: unifiedgosdk.Pointer("542e253e-9be2-4769-9e7b-513144533652"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("077d67ff-ee2e-4c5f-ac4f-bb6481113a38"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitAmount: unifiedgosdk.Pointer[float64](364.74),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("0906ed39-7dae-456f-9c82-077543591acd"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitAmount: unifiedgosdk.Pointer[float64](258.35),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("f4ae5862-e96c-4dbc-a116-2f9d71403708"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    Notes: unifiedgosdk.Pointer("Casus umquam conicio cinis demitto sint."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitAmount: unifiedgosdk.Pointer[float64](28.14),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("9ad8a271-29b8-4c2f-a966-e56e1d5fa671"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitAmount: unifiedgosdk.Pointer[float64](157.04),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("e783cbf5-5d97-4777-8798-596b0362a0c4"),
                    ItemName: unifiedgosdk.Pointer("Hotel"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitAmount: unifiedgosdk.Pointer[float64](164.02),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("A payment of XPF 714.84 occurred at Bins, Barrows and Koch using a card ending in ****8422 for account ***4614."),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            PostedAt: types.MustNewTimeFromString("2020-06-11T03:39:37.305Z"),
            Status: shared.AccountingExpenseStatusSubmitted.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](972.29),
            Type: shared.AccountingExpenseTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-09-02T08:08:15.489Z"),
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
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("5b3cb8a3-5737-4b18-b126-5266526343a7"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-31T12:54:25.603Z"),
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