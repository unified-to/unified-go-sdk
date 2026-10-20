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

Create an expensereport

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
            ApprovedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.842Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("e187f895-2ff6-4c87-8d9e-4f0384ca8c56"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("69aadfec-2225-4897-9b97-0e19b5be5cc4"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("7591c890-6304-4981-8c02-16a34136b872"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("2385a09f-376b-4986-ac45-3edf2b178ccb"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("17558b49-df81-48da-a6f0-e8564b6d7bf2"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.842Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.842Z"),
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

Retrieve an expensereport

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

List all expensereports

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

Update an expensereport

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
            ApprovedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("76f16e46-7e5b-43c4-947b-d51ecca30e01"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("106ca4e5-1c4d-4d61-b7db-08a042e57239"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("46e4f4ab-d424-473d-af0f-66dedd2d95fa"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a855ae7d-3616-4c4f-9afa-1ada193fad8d"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("523318e7-b535-4c57-9881-e1c3f75279c3"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
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

Remove an expensereport

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

Update an expensereport

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
            ApprovedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ExternalNumber: unifiedgosdk.Pointer("ER-31240"),
            ID: unifiedgosdk.Pointer("76f16e46-7e5b-43c4-947b-d51ecca30e01"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("106ca4e5-1c4d-4d61-b7db-08a042e57239"),
                    ItemName: unifiedgosdk.Pointer("Conference ticket"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitAmount: unifiedgosdk.Pointer[float64](39.9),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("46e4f4ab-d424-473d-af0f-66dedd2d95fa"),
                    ItemName: unifiedgosdk.Pointer("Flight"),
                    Notes: unifiedgosdk.Pointer("Soluta omnis dolor."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitAmount: unifiedgosdk.Pointer[float64](193.17),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("a855ae7d-3616-4c4f-9afa-1ada193fad8d"),
                    ItemName: unifiedgosdk.Pointer("Client dinner"),
                    Notes: unifiedgosdk.Pointer("Altus caveo capio."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitAmount: unifiedgosdk.Pointer[float64](230.12),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("523318e7-b535-4c57-9881-e1c3f75279c3"),
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
            ReimbursedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](908.46),
            UpdatedAt: types.MustNewTimeFromString("2021-04-19T08:06:52.849Z"),
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