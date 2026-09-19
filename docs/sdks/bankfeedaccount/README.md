# Bankfeedaccount

## Overview

### Available Operations

* [CreateAccountingBankfeedaccount](#createaccountingbankfeedaccount) - Create a bankfeedaccount
* [GetAccountingBankfeedaccount](#getaccountingbankfeedaccount) - Retrieve a bankfeedaccount
* [ListAccountingBankfeedaccounts](#listaccountingbankfeedaccounts) - List all bankfeedaccounts
* [PatchAccountingBankfeedaccount](#patchaccountingbankfeedaccount) - Update a bankfeedaccount
* [RemoveAccountingBankfeedaccount](#removeaccountingbankfeedaccount) - Remove a bankfeedaccount
* [UpdateAccountingBankfeedaccount](#updateaccountingbankfeedaccount) - Update a bankfeedaccount

## CreateAccountingBankfeedaccount

Create a bankfeedaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingBankfeedaccount" method="post" path="/accounting/{connection_id}/bankfeedaccount" example="accounting_bankfeedaccount" -->
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

    res, err := s.Bankfeedaccount.CreateAccountingBankfeedaccount(ctx, operations.CreateAccountingBankfeedaccountRequest{
        AccountingBankfeedaccount: shared.AccountingBankfeedaccount{
            AccountID: unifiedgosdk.Pointer("baa0e9a4-65e9-4bf5-856a-ce46fc36ebb1"),
            AccountNumber: unifiedgosdk.Pointer("30369722"),
            AccountNumberLast4: unifiedgosdk.Pointer("9722"),
            AccountType: shared.AccountTypeLoan.ToPointer(),
            Balance: unifiedgosdk.Pointer[float64](90358.0),
            BankName: unifiedgosdk.Pointer("Weissnat Inc"),
            CreatedAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            FeedStartAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            ID: unifiedgosdk.Pointer("f66a8b38-7d8d-4ac7-8a5c-bde4b6ea64b0"),
            Name: unifiedgosdk.Pointer("Corwin, Donnelly and Connelly Savings"),
            RoutingNumber: unifiedgosdk.Pointer("667753156"),
            Status: shared.AccountingBankfeedaccountStatusActive.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-11T21:16:29.668Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.CreateAccountingBankfeedaccountRequest](../../pkg/models/operations/createaccountingbankfeedaccountrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateAccountingBankfeedaccountResponse](../../pkg/models/operations/createaccountingbankfeedaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingBankfeedaccount

Retrieve a bankfeedaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingBankfeedaccount" method="get" path="/accounting/{connection_id}/bankfeedaccount/{id}" -->
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

    res, err := s.Bankfeedaccount.GetAccountingBankfeedaccount(ctx, operations.GetAccountingBankfeedaccountRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAccountingBankfeedaccountRequest](../../pkg/models/operations/getaccountingbankfeedaccountrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GetAccountingBankfeedaccountResponse](../../pkg/models/operations/getaccountingbankfeedaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingBankfeedaccounts

List all bankfeedaccounts

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingBankfeedaccounts" method="get" path="/accounting/{connection_id}/bankfeedaccount" -->
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

    res, err := s.Bankfeedaccount.ListAccountingBankfeedaccounts(ctx, operations.ListAccountingBankfeedaccountsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedaccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListAccountingBankfeedaccountsRequest](../../pkg/models/operations/listaccountingbankfeedaccountsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListAccountingBankfeedaccountsResponse](../../pkg/models/operations/listaccountingbankfeedaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingBankfeedaccount

Update a bankfeedaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingBankfeedaccount" method="patch" path="/accounting/{connection_id}/bankfeedaccount/{id}" example="accounting_bankfeedaccount" -->
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

    res, err := s.Bankfeedaccount.PatchAccountingBankfeedaccount(ctx, operations.PatchAccountingBankfeedaccountRequest{
        AccountingBankfeedaccount: shared.AccountingBankfeedaccount{
            AccountID: unifiedgosdk.Pointer("baa0e9a4-65e9-4bf5-856a-ce46fc36ebb1"),
            AccountNumber: unifiedgosdk.Pointer("30369722"),
            AccountNumberLast4: unifiedgosdk.Pointer("9722"),
            AccountType: shared.AccountTypeLoan.ToPointer(),
            Balance: unifiedgosdk.Pointer[float64](90358.0),
            BankName: unifiedgosdk.Pointer("Weissnat Inc"),
            CreatedAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            FeedStartAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            ID: unifiedgosdk.Pointer("c6e2408d-5646-4f3e-ac84-7b2adb89ad93"),
            Name: unifiedgosdk.Pointer("Corwin, Donnelly and Connelly Savings"),
            RoutingNumber: unifiedgosdk.Pointer("667753156"),
            Status: shared.AccountingBankfeedaccountStatusActive.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-11T21:16:29.678Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchAccountingBankfeedaccountRequest](../../pkg/models/operations/patchaccountingbankfeedaccountrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PatchAccountingBankfeedaccountResponse](../../pkg/models/operations/patchaccountingbankfeedaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingBankfeedaccount

Remove a bankfeedaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingBankfeedaccount" method="delete" path="/accounting/{connection_id}/bankfeedaccount/{id}" -->
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

    res, err := s.Bankfeedaccount.RemoveAccountingBankfeedaccount(ctx, operations.RemoveAccountingBankfeedaccountRequest{
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.RemoveAccountingBankfeedaccountRequest](../../pkg/models/operations/removeaccountingbankfeedaccountrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.RemoveAccountingBankfeedaccountResponse](../../pkg/models/operations/removeaccountingbankfeedaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingBankfeedaccount

Update a bankfeedaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingBankfeedaccount" method="put" path="/accounting/{connection_id}/bankfeedaccount/{id}" example="accounting_bankfeedaccount" -->
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

    res, err := s.Bankfeedaccount.UpdateAccountingBankfeedaccount(ctx, operations.UpdateAccountingBankfeedaccountRequest{
        AccountingBankfeedaccount: shared.AccountingBankfeedaccount{
            AccountID: unifiedgosdk.Pointer("baa0e9a4-65e9-4bf5-856a-ce46fc36ebb1"),
            AccountNumber: unifiedgosdk.Pointer("30369722"),
            AccountNumberLast4: unifiedgosdk.Pointer("9722"),
            AccountType: shared.AccountTypeLoan.ToPointer(),
            Balance: unifiedgosdk.Pointer[float64](90358.0),
            BankName: unifiedgosdk.Pointer("Weissnat Inc"),
            CreatedAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            Currency: unifiedgosdk.Pointer("SSP"),
            FeedStartAt: types.MustNewTimeFromString("2022-10-31T16:42:19.277Z"),
            ID: unifiedgosdk.Pointer("c6e2408d-5646-4f3e-ac84-7b2adb89ad93"),
            Name: unifiedgosdk.Pointer("Corwin, Donnelly and Connelly Savings"),
            RoutingNumber: unifiedgosdk.Pointer("667753156"),
            Status: shared.AccountingBankfeedaccountStatusActive.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-04-11T21:16:29.678Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBankfeedaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.UpdateAccountingBankfeedaccountRequest](../../pkg/models/operations/updateaccountingbankfeedaccountrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateAccountingBankfeedaccountResponse](../../pkg/models/operations/updateaccountingbankfeedaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |