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

Creates a AccountingBankfeedtransaction. Provide the unified AccountingBankfeedtransaction fields in the request body; they are mapped to the provider's API. Returns the created AccountingBankfeedtransaction.

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
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Clothing"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Payment of SRD 92.72 for withdrawal at Halvorson - Ziemann, processed with card ending ****3731 linked to account ***7781."),
            ID: unifiedgosdk.Pointer("b8323a8c-cc6c-4a49-8c79-949ec5e4d438"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Buckridge - Feest"),
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-16T04:42:22.598Z"),
            Reference: unifiedgosdk.Pointer("79574105"),
            TransactionAt: types.MustNewTimeFromString("2022-07-30T11:54:11.959Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-25T03:32:05.495Z"),
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

Retrieves a single AccountingBankfeedtransaction by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of AccountingBankfeedtransactions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the AccountingBankfeedtransaction identified by `id`. Provide the unified AccountingBankfeedtransaction fields to change in the request body.

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
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Clothing"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Payment of SRD 92.72 for withdrawal at Halvorson - Ziemann, processed with card ending ****3731 linked to account ***7781."),
            ID: unifiedgosdk.Pointer("7d725600-a4da-41e4-bf44-a786684bac05"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Buckridge - Feest"),
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-16T04:42:22.612Z"),
            Reference: unifiedgosdk.Pointer("79574105"),
            TransactionAt: types.MustNewTimeFromString("2022-07-30T11:54:11.961Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-25T03:32:05.496Z"),
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

Deletes the AccountingBankfeedtransaction identified by `id`.

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

Updates the AccountingBankfeedtransaction identified by `id`. Provide the unified AccountingBankfeedtransaction fields to change in the request body.

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
            Amount: unifiedgosdk.Pointer[float64](60889.0),
            BankCategory: unifiedgosdk.Pointer("Clothing"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-24T23:41:08.374Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Payment of SRD 92.72 for withdrawal at Halvorson - Ziemann, processed with card ending ****3731 linked to account ***7781."),
            ID: unifiedgosdk.Pointer("7d725600-a4da-41e4-bf44-a786684bac05"),
            IsPending: unifiedgosdk.Pointer(true),
            MerchantName: unifiedgosdk.Pointer("Buckridge - Feest"),
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-16T04:42:22.612Z"),
            Reference: unifiedgosdk.Pointer("79574105"),
            TransactionAt: types.MustNewTimeFromString("2022-07-30T11:54:11.961Z"),
            Type: shared.AccountingBankfeedtransactionTypeCredit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-05-25T03:32:05.496Z"),
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