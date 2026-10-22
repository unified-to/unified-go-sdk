# Transaction

## Overview

### Available Operations

* [CreateAccountingTransaction](#createaccountingtransaction) - Create a transaction
* [GetAccountingTransaction](#getaccountingtransaction) - Retrieve a transaction
* [ListAccountingTransactions](#listaccountingtransactions) - List all transactions
* [PatchAccountingTransaction](#patchaccountingtransaction) - Update a transaction
* [RemoveAccountingTransaction](#removeaccountingtransaction) - Remove a transaction
* [UpdateAccountingTransaction](#updateaccountingtransaction) - Update a transaction

## CreateAccountingTransaction

Creates a AccountingTransaction. Provide the unified AccountingTransaction fields in the request body; they are mapped to the provider's API. Returns the created AccountingTransaction.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingTransaction" method="post" path="/accounting/{connection_id}/transaction" example="accounting_transaction" -->
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

    res, err := s.Transaction.CreateAccountingTransaction(ctx, operations.CreateAccountingTransactionRequest{
        AccountingTransaction: shared.AccountingTransaction{
            CreatedAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            ID: unifiedgosdk.Pointer("52dd8af7-7cd7-4625-ab77-a9d341b24d4c"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The turquoise Hat combines Wallis and Futuna aesthetics with Potassium-based durability"),
                    ID: unifiedgosdk.Pointer("d314aa8e-2877-4e1f-8e2d-e1854503d98a"),
                    Name: unifiedgosdk.Pointer("Shirt"),
                    TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](19.0),
                },
            },
            Memo: unifiedgosdk.Pointer("deposit of ISK 792.81 at Fritsch - Dooley charged to account ending in 7879 using card ending in ****3230."),
            Status: shared.AccountingTransactionStatusPending.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2019-10-14T06:18:51.360Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateAccountingTransactionRequest](../../pkg/models/operations/createaccountingtransactionrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateAccountingTransactionResponse](../../pkg/models/operations/createaccountingtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingTransaction

Retrieves a single AccountingTransaction by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingTransaction" method="get" path="/accounting/{connection_id}/transaction/{id}" -->
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

    res, err := s.Transaction.GetAccountingTransaction(ctx, operations.GetAccountingTransactionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAccountingTransactionRequest](../../pkg/models/operations/getaccountingtransactionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetAccountingTransactionResponse](../../pkg/models/operations/getaccountingtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingTransactions

Returns a list of AccountingTransactions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingTransactions" method="get" path="/accounting/{connection_id}/transaction" -->
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

    res, err := s.Transaction.ListAccountingTransactions(ctx, operations.ListAccountingTransactionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListAccountingTransactionsRequest](../../pkg/models/operations/listaccountingtransactionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListAccountingTransactionsResponse](../../pkg/models/operations/listaccountingtransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingTransaction

Updates the AccountingTransaction identified by `id`. Provide the unified AccountingTransaction fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingTransaction" method="patch" path="/accounting/{connection_id}/transaction/{id}" example="accounting_transaction" -->
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

    res, err := s.Transaction.PatchAccountingTransaction(ctx, operations.PatchAccountingTransactionRequest{
        AccountingTransaction: shared.AccountingTransaction{
            CreatedAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            ID: unifiedgosdk.Pointer("22225963-66fb-48f2-9300-11e70830bb12"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The turquoise Hat combines Wallis and Futuna aesthetics with Potassium-based durability"),
                    ID: unifiedgosdk.Pointer("577033d5-6fa9-46c2-938e-3831528b30c8"),
                    Name: unifiedgosdk.Pointer("Shirt"),
                    TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](19.0),
                },
            },
            Memo: unifiedgosdk.Pointer("deposit of ISK 792.81 at Fritsch - Dooley charged to account ending in 7879 using card ending in ****3230."),
            Status: shared.AccountingTransactionStatusPending.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2019-10-14T06:18:51.360Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchAccountingTransactionRequest](../../pkg/models/operations/patchaccountingtransactionrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PatchAccountingTransactionResponse](../../pkg/models/operations/patchaccountingtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingTransaction

Deletes the AccountingTransaction identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingTransaction" method="delete" path="/accounting/{connection_id}/transaction/{id}" -->
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

    res, err := s.Transaction.RemoveAccountingTransaction(ctx, operations.RemoveAccountingTransactionRequest{
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
| `request`                                                                                                          | [operations.RemoveAccountingTransactionRequest](../../pkg/models/operations/removeaccountingtransactionrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.RemoveAccountingTransactionResponse](../../pkg/models/operations/removeaccountingtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingTransaction

Updates the AccountingTransaction identified by `id`. Provide the unified AccountingTransaction fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingTransaction" method="put" path="/accounting/{connection_id}/transaction/{id}" example="accounting_transaction" -->
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

    res, err := s.Transaction.UpdateAccountingTransaction(ctx, operations.UpdateAccountingTransactionRequest{
        AccountingTransaction: shared.AccountingTransaction{
            CreatedAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            ID: unifiedgosdk.Pointer("22225963-66fb-48f2-9300-11e70830bb12"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The turquoise Hat combines Wallis and Futuna aesthetics with Potassium-based durability"),
                    ID: unifiedgosdk.Pointer("577033d5-6fa9-46c2-938e-3831528b30c8"),
                    Name: unifiedgosdk.Pointer("Shirt"),
                    TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](19.0),
                },
            },
            Memo: unifiedgosdk.Pointer("deposit of ISK 792.81 at Fritsch - Dooley charged to account ending in 7879 using card ending in ****3230."),
            Status: shared.AccountingTransactionStatusPending.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](10184.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2019-10-14T06:18:51.360Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateAccountingTransactionRequest](../../pkg/models/operations/updateaccountingtransactionrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateAccountingTransactionResponse](../../pkg/models/operations/updateaccountingtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |