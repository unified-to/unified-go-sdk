# Expensereport

## Overview

### Available Operations

* [CreateAccountingExpensereport](#createaccountingexpensereport) - Create an expensereport
* [GetAccountingExpensereport](#getaccountingexpensereport) - Retrieve an expensereport
* [ListAccountingExpensereports](#listaccountingexpensereports) - List all expensereports
* [PatchAccountingExpensereport](#patchaccountingexpensereport) - Update an expensereport
* [RemoveAccountingExpensereport](#removeaccountingexpensereport) - Remove an expensereport
* [UpdateAccountingExpensereport](#updateaccountingexpensereport) - Update an expensereport

## CreateAccountingExpensereport

Creates a AccountingExpensereport. Provide the unified AccountingExpensereport fields in the request body; they are mapped to the provider's API. Returns the created AccountingExpensereport.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingExpensereport" method="post" path="/accounting/{connection_id}/expensereport" example="accounting_expensereport" -->
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

    res, err := s.Expensereport.CreateAccountingExpensereport(ctx, operations.CreateAccountingExpensereportRequest{
        AccountingExpensereport: shared.AccountingExpensereport{
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.505Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("71457120-e20c-4755-a1d3-e99052c65559"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("88954ec2-0dd9-45c0-bf38-3e22fd795169"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("2f85c060-bcdd-4bc1-90e4-43c0b4594d90"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a220cca8-75a0-4a9d-a818-803a3899453f"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("1de01da8-bd3a-4301-bf72-50a98b6f529a"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Lolita Kunde - October expenses"),
            PostedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](908.46),
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.505Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.505Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpensereport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateAccountingExpensereportRequest](../../pkg/models/operations/createaccountingexpensereportrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateAccountingExpensereportResponse](../../pkg/models/operations/createaccountingexpensereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingExpensereport

Retrieves a single AccountingExpensereport by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingExpensereport" method="get" path="/accounting/{connection_id}/expensereport/{id}" -->
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

    res, err := s.Expensereport.GetAccountingExpensereport(ctx, operations.GetAccountingExpensereportRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpensereport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAccountingExpensereportRequest](../../pkg/models/operations/getaccountingexpensereportrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetAccountingExpensereportResponse](../../pkg/models/operations/getaccountingexpensereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingExpensereports

Returns a list of AccountingExpensereports. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingExpensereports" method="get" path="/accounting/{connection_id}/expensereport" -->
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

    res, err := s.Expensereport.ListAccountingExpensereports(ctx, operations.ListAccountingExpensereportsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpensereports != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListAccountingExpensereportsRequest](../../pkg/models/operations/listaccountingexpensereportsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListAccountingExpensereportsResponse](../../pkg/models/operations/listaccountingexpensereportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingExpensereport

Updates the AccountingExpensereport identified by `id`. Provide the unified AccountingExpensereport fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingExpensereport" method="patch" path="/accounting/{connection_id}/expensereport/{id}" example="accounting_expensereport" -->
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

    res, err := s.Expensereport.PatchAccountingExpensereport(ctx, operations.PatchAccountingExpensereportRequest{
        AccountingExpensereport: shared.AccountingExpensereport{
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("1a49fdf7-5ce6-40bb-8e20-e9c5e519e435"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("7b84854d-ec6e-41ff-a41b-22bf6ffd2613"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("eddf9644-5aa8-4893-b42a-e242cd7cef7d"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a02ba3dd-4c37-411d-8213-625b8d879974"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("fe9cda4e-50b1-4666-88ee-3d071bb135a9"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Lolita Kunde - October expenses"),
            PostedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](908.46),
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpensereport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchAccountingExpensereportRequest](../../pkg/models/operations/patchaccountingexpensereportrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchAccountingExpensereportResponse](../../pkg/models/operations/patchaccountingexpensereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingExpensereport

Deletes the AccountingExpensereport identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingExpensereport" method="delete" path="/accounting/{connection_id}/expensereport/{id}" -->
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

    res, err := s.Expensereport.RemoveAccountingExpensereport(ctx, operations.RemoveAccountingExpensereportRequest{
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.RemoveAccountingExpensereportRequest](../../pkg/models/operations/removeaccountingexpensereportrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.RemoveAccountingExpensereportResponse](../../pkg/models/operations/removeaccountingexpensereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingExpensereport

Updates the AccountingExpensereport identified by `id`. Provide the unified AccountingExpensereport fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingExpensereport" method="put" path="/accounting/{connection_id}/expensereport/{id}" example="accounting_expensereport" -->
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

    res, err := s.Expensereport.UpdateAccountingExpensereport(ctx, operations.UpdateAccountingExpensereportRequest{
        AccountingExpensereport: shared.AccountingExpensereport{
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("1a49fdf7-5ce6-40bb-8e20-e9c5e519e435"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("7b84854d-ec6e-41ff-a41b-22bf6ffd2613"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("eddf9644-5aa8-4893-b42a-e242cd7cef7d"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a02ba3dd-4c37-411d-8213-625b8d879974"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("fe9cda4e-50b1-4666-88ee-3d071bb135a9"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitAmount: unifiedgosdk.Pointer[float64](445.27),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Lolita Kunde - October expenses"),
            PostedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](908.46),
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T07:08:41.508Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingExpensereport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateAccountingExpensereportRequest](../../pkg/models/operations/updateaccountingexpensereportrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateAccountingExpensereportResponse](../../pkg/models/operations/updateaccountingexpensereportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |