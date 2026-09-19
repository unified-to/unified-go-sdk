# Bankfeedtransaction

## Overview

### Available Operations

* [CreateAccountingBankfeedtransaction](#createaccountingbankfeedtransaction) - Create a bankfeedtransaction
* [GetAccountingBankfeedtransaction](#getaccountingbankfeedtransaction) - Retrieve a bankfeedtransaction
* [ListAccountingBankfeedtransactions](#listaccountingbankfeedtransactions) - List all bankfeedtransactions
* [PatchAccountingBankfeedtransaction](#patchaccountingbankfeedtransaction) - Update a bankfeedtransaction
* [RemoveAccountingBankfeedtransaction](#removeaccountingbankfeedtransaction) - Remove a bankfeedtransaction
* [UpdateAccountingBankfeedtransaction](#updateaccountingbankfeedtransaction) - Update a bankfeedtransaction

## CreateAccountingBankfeedtransaction

Create a bankfeedtransaction

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingBankfeedtransaction" method="post" path="/accounting/{connection_id}/bankfeedtransaction" example="accounting_bankfeedtransaction" -->
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

    res, err := s.Bankfeedtransaction.CreateAccountingBankfeedtransaction(ctx, operations.CreateAccountingBankfeedtransactionRequest{
        AccountingBankfeedtransaction: shared.AccountingBankfeedtransaction{
            AccountID: unifiedgosdk.Pointer("b7dc4175-1368-4b89-a700-d621b6666648"),
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Games"),
            BankfeedaccountID: unifiedgosdk.Pointer("34c1d05f-5b62-4bcd-9121-3be8b720941f"),
            CategoryIds: []string{},
            ContactID: unifiedgosdk.Pointer("1ef58ebe-f9c9-46f6-9d9c-2df2658503be"),
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("SRD"),
            Description: unifiedgosdk.Pointer("payment transaction at McLaughlin - Schaden using card ending with ****8233 for DOP 574.03 in account ***9523."),
            ID: unifiedgosdk.Pointer("2b61660c-37e5-4718-8102-2c4a7545f17c"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Reichert, Erdman and Tillman"),
            PostedAt: types.MustNewTimeFromString("2025-03-24T09:30:58.611Z"),
            Reference: unifiedgosdk.Pointer("93642593"),
            TransactionAt: types.MustNewTimeFromString("2022-07-27T21:30:26.661Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-23T21:36:22.811Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedtransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreateAccountingBankfeedtransactionRequest](../../pkg/models/operations/createaccountingbankfeedtransactionrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateAccountingBankfeedtransactionResponse](../../pkg/models/operations/createaccountingbankfeedtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingBankfeedtransaction

Retrieve a bankfeedtransaction

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingBankfeedtransaction" method="get" path="/accounting/{connection_id}/bankfeedtransaction/{id}" -->
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

    res, err := s.Bankfeedtransaction.GetAccountingBankfeedtransaction(ctx, operations.GetAccountingBankfeedtransactionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedtransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetAccountingBankfeedtransactionRequest](../../pkg/models/operations/getaccountingbankfeedtransactionrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetAccountingBankfeedtransactionResponse](../../pkg/models/operations/getaccountingbankfeedtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingBankfeedtransactions

List all bankfeedtransactions

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingBankfeedtransactions" method="get" path="/accounting/{connection_id}/bankfeedtransaction" -->
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

    res, err := s.Bankfeedtransaction.ListAccountingBankfeedtransactions(ctx, operations.ListAccountingBankfeedtransactionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedtransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.ListAccountingBankfeedtransactionsRequest](../../pkg/models/operations/listaccountingbankfeedtransactionsrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.ListAccountingBankfeedtransactionsResponse](../../pkg/models/operations/listaccountingbankfeedtransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingBankfeedtransaction

Update a bankfeedtransaction

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingBankfeedtransaction" method="patch" path="/accounting/{connection_id}/bankfeedtransaction/{id}" example="accounting_bankfeedtransaction" -->
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

    res, err := s.Bankfeedtransaction.PatchAccountingBankfeedtransaction(ctx, operations.PatchAccountingBankfeedtransactionRequest{
        AccountingBankfeedtransaction: shared.AccountingBankfeedtransaction{
            AccountID: unifiedgosdk.Pointer("b7dc4175-1368-4b89-a700-d621b6666648"),
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Games"),
            BankfeedaccountID: unifiedgosdk.Pointer("34c1d05f-5b62-4bcd-9121-3be8b720941f"),
            CategoryIds: []string{},
            ContactID: unifiedgosdk.Pointer("1ef58ebe-f9c9-46f6-9d9c-2df2658503be"),
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("SRD"),
            Description: unifiedgosdk.Pointer("payment transaction at McLaughlin - Schaden using card ending with ****8233 for DOP 574.03 in account ***9523."),
            ID: unifiedgosdk.Pointer("7c97e722-16a3-4103-bb88-c4b3eefe4e58"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Reichert, Erdman and Tillman"),
            PostedAt: types.MustNewTimeFromString("2025-03-24T09:30:58.620Z"),
            Reference: unifiedgosdk.Pointer("93642593"),
            TransactionAt: types.MustNewTimeFromString("2022-07-27T21:30:26.662Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-23T21:36:22.811Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedtransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PatchAccountingBankfeedtransactionRequest](../../pkg/models/operations/patchaccountingbankfeedtransactionrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.PatchAccountingBankfeedtransactionResponse](../../pkg/models/operations/patchaccountingbankfeedtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingBankfeedtransaction

Remove a bankfeedtransaction

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingBankfeedtransaction" method="delete" path="/accounting/{connection_id}/bankfeedtransaction/{id}" -->
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

    res, err := s.Bankfeedtransaction.RemoveAccountingBankfeedtransaction(ctx, operations.RemoveAccountingBankfeedtransactionRequest{
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.RemoveAccountingBankfeedtransactionRequest](../../pkg/models/operations/removeaccountingbankfeedtransactionrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.RemoveAccountingBankfeedtransactionResponse](../../pkg/models/operations/removeaccountingbankfeedtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingBankfeedtransaction

Update a bankfeedtransaction

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingBankfeedtransaction" method="put" path="/accounting/{connection_id}/bankfeedtransaction/{id}" example="accounting_bankfeedtransaction" -->
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

    res, err := s.Bankfeedtransaction.UpdateAccountingBankfeedtransaction(ctx, operations.UpdateAccountingBankfeedtransactionRequest{
        AccountingBankfeedtransaction: shared.AccountingBankfeedtransaction{
            AccountID: unifiedgosdk.Pointer("b7dc4175-1368-4b89-a700-d621b6666648"),
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Games"),
            BankfeedaccountID: unifiedgosdk.Pointer("34c1d05f-5b62-4bcd-9121-3be8b720941f"),
            CategoryIds: []string{},
            ContactID: unifiedgosdk.Pointer("1ef58ebe-f9c9-46f6-9d9c-2df2658503be"),
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("SRD"),
            Description: unifiedgosdk.Pointer("payment transaction at McLaughlin - Schaden using card ending with ****8233 for DOP 574.03 in account ***9523."),
            ID: unifiedgosdk.Pointer("7c97e722-16a3-4103-bb88-c4b3eefe4e58"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Reichert, Erdman and Tillman"),
            PostedAt: types.MustNewTimeFromString("2025-03-24T09:30:58.620Z"),
            Reference: unifiedgosdk.Pointer("93642593"),
            TransactionAt: types.MustNewTimeFromString("2022-07-27T21:30:26.662Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-23T21:36:22.811Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedtransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.UpdateAccountingBankfeedtransactionRequest](../../pkg/models/operations/updateaccountingbankfeedtransactionrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.UpdateAccountingBankfeedtransactionResponse](../../pkg/models/operations/updateaccountingbankfeedtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |