# Creditmemo

## Overview

### Available Operations

* [CreateAccountingCreditmemo](#createaccountingcreditmemo) - Create a creditmemo
* [GetAccountingCreditmemo](#getaccountingcreditmemo) - Retrieve a creditmemo
* [ListAccountingCreditmemoes](#listaccountingcreditmemoes) - List all creditmemoes
* [PatchAccountingCreditmemo](#patchaccountingcreditmemo) - Update a creditmemo
* [RemoveAccountingCreditmemo](#removeaccountingcreditmemo) - Remove a creditmemo
* [UpdateAccountingCreditmemo](#updateaccountingcreditmemo) - Update a creditmemo

## CreateAccountingCreditmemo

Creates a AccountingCreditmemo. Provide the unified AccountingCreditmemo fields in the request body; they are mapped to the provider's API. Returns the created AccountingCreditmemo.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingCreditmemo" method="post" path="/accounting/{connection_id}/creditmemo" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.CreateAccountingCreditmemo(ctx, operations.CreateAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{},
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("CM-99537"),
            Currency: unifiedgosdk.Pointer("AUD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("ab78ad12-c79a-47fc-bbf8-bb9e3a14e107"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Hic vorax dedecor amo adfero torqueo quas victoria."),
            PostedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingCreditmemoStatusAuthorized.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2025-09-16T23:38:03.135Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingCreditmemoRequest](../../pkg/models/operations/createaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateAccountingCreditmemoResponse](../../pkg/models/operations/createaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingCreditmemo

Retrieves a single AccountingCreditmemo by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingCreditmemo" method="get" path="/accounting/{connection_id}/creditmemo/{id}" -->
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

    res, err := s.Creditmemo.GetAccountingCreditmemo(ctx, operations.GetAccountingCreditmemoRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetAccountingCreditmemoRequest](../../pkg/models/operations/getaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAccountingCreditmemoResponse](../../pkg/models/operations/getaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingCreditmemoes

Returns a list of AccountingCreditmemoes. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingCreditmemoes" method="get" path="/accounting/{connection_id}/creditmemo" -->
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

    res, err := s.Creditmemo.ListAccountingCreditmemoes(ctx, operations.ListAccountingCreditmemoesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemoes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListAccountingCreditmemoesRequest](../../pkg/models/operations/listaccountingcreditmemoesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListAccountingCreditmemoesResponse](../../pkg/models/operations/listaccountingcreditmemoesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingCreditmemo

Updates the AccountingCreditmemo identified by `id`. Provide the unified AccountingCreditmemo fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingCreditmemo" method="patch" path="/accounting/{connection_id}/creditmemo/{id}" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.PatchAccountingCreditmemo(ctx, operations.PatchAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{},
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("CM-99537"),
            Currency: unifiedgosdk.Pointer("AUD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("51f187de-cf66-490a-bf61-e9490693edbd"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Hic vorax dedecor amo adfero torqueo quas victoria."),
            PostedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingCreditmemoStatusAuthorized.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2025-09-16T23:38:03.170Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchAccountingCreditmemoRequest](../../pkg/models/operations/patchaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PatchAccountingCreditmemoResponse](../../pkg/models/operations/patchaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingCreditmemo

Deletes the AccountingCreditmemo identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingCreditmemo" method="delete" path="/accounting/{connection_id}/creditmemo/{id}" -->
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

    res, err := s.Creditmemo.RemoveAccountingCreditmemo(ctx, operations.RemoveAccountingCreditmemoRequest{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RemoveAccountingCreditmemoRequest](../../pkg/models/operations/removeaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.RemoveAccountingCreditmemoResponse](../../pkg/models/operations/removeaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingCreditmemo

Updates the AccountingCreditmemo identified by `id`. Provide the unified AccountingCreditmemo fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingCreditmemo" method="put" path="/accounting/{connection_id}/creditmemo/{id}" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.UpdateAccountingCreditmemo(ctx, operations.UpdateAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{},
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("CM-99537"),
            Currency: unifiedgosdk.Pointer("AUD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("51f187de-cf66-490a-bf61-e9490693edbd"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Hic vorax dedecor amo adfero torqueo quas victoria."),
            PostedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingCreditmemoStatusAuthorized.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2025-09-16T23:38:03.170Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateAccountingCreditmemoRequest](../../pkg/models/operations/updateaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateAccountingCreditmemoResponse](../../pkg/models/operations/updateaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |