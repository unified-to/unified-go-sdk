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

Create a transaction

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
            ExchangeRate: unifiedgosdk.Pointer[float64](0.5674),
            ID: unifiedgosdk.Pointer("ab063b63-1037-4692-8790-f03b69228bbe"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The Nikolas Table is the latest in a series of downright products from Beier and Sons"),
                    ID: unifiedgosdk.Pointer("575c071a-8215-40e1-8854-3204db736126"),
                    Name: unifiedgosdk.Pointer("Salad"),
                    ObjectType: unifiedgosdk.Pointer("delicate"),
                    TotalAmount: unifiedgosdk.Pointer[float64](58531.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](91.0),
                },
            },
            Memo: unifiedgosdk.Pointer("withdrawal of USD 873.18 at Harber and Sons charged to account ending in 1804 using card ending in ****7022."),
            Status: shared.AccountingTransactionStatusVoid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](94452.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-17T22:22:05.679Z"),
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

Retrieve a transaction

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

List all transactions

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

Update a transaction

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
            ExchangeRate: unifiedgosdk.Pointer[float64](0.5674),
            ID: unifiedgosdk.Pointer("819589e8-319c-46dc-a074-02b5508f4a77"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The Nikolas Table is the latest in a series of downright products from Beier and Sons"),
                    ID: unifiedgosdk.Pointer("24247fbd-74f3-4866-a465-87eaeb700fad"),
                    Name: unifiedgosdk.Pointer("Salad"),
                    ObjectType: unifiedgosdk.Pointer("delicate"),
                    TotalAmount: unifiedgosdk.Pointer[float64](58531.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](91.0),
                },
            },
            Memo: unifiedgosdk.Pointer("withdrawal of USD 873.18 at Harber and Sons charged to account ending in 1804 using card ending in ****7022."),
            Status: shared.AccountingTransactionStatusVoid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](94452.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-17T22:22:05.687Z"),
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

Remove a transaction

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

Update a transaction

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
            ExchangeRate: unifiedgosdk.Pointer[float64](0.5674),
            ID: unifiedgosdk.Pointer("819589e8-319c-46dc-a074-02b5508f4a77"),
            Lineitems: []shared.AccountingTransactionLineItem{
                shared.AccountingTransactionLineItem{
                    CategoryIds: []string{},
                    Description: unifiedgosdk.Pointer("The Nikolas Table is the latest in a series of downright products from Beier and Sons"),
                    ID: unifiedgosdk.Pointer("24247fbd-74f3-4866-a465-87eaeb700fad"),
                    Name: unifiedgosdk.Pointer("Salad"),
                    ObjectType: unifiedgosdk.Pointer("delicate"),
                    TotalAmount: unifiedgosdk.Pointer[float64](58531.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](536.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](91.0),
                },
            },
            Memo: unifiedgosdk.Pointer("withdrawal of USD 873.18 at Harber and Sons charged to account ending in 1804 using card ending in ****7022."),
            Status: shared.AccountingTransactionStatusVoid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](94452.0),
            TransactionAt: types.MustNewTimeFromString("2019-09-25T11:40:42.574Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-17T22:22:05.687Z"),
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