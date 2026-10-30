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
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.091Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("07416da2-310f-40ab-9f7d-b9c56ca8066e"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("ebe46304-cd84-4d3d-b99e-339d1837cf39"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("09cd8844-abca-45b8-be4a-d528da362785"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a369e845-c07c-477b-bc1d-b408b7d58882"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("30ee46bb-7fa5-4fa3-ad9b-6efee6fd4e51"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.091Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.091Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("c1e9da03-2cfa-4f4f-940b-d4fc9e3a9825"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("aeef158a-6fbf-4070-b384-fedfadd90054"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("134303eb-c1f4-470f-b1cb-180c63bb9f0f"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("2b6e9435-f51b-4539-a4a2-3f1de6254f03"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("9320d24a-67dd-4c19-8d15-26aa2f805760"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("c1e9da03-2cfa-4f4f-940b-d4fc9e3a9825"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("aeef158a-6fbf-4070-b384-fedfadd90054"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("134303eb-c1f4-470f-b1cb-180c63bb9f0f"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("2b6e9435-f51b-4539-a4a2-3f1de6254f03"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("9320d24a-67dd-4c19-8d15-26aa2f805760"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-20T16:58:34.098Z"),
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