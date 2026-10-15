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
            ApprovedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.220Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("74812807-3b03-4864-b814-40ed576084de"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("569cbaa2-d9d7-4def-a3f5-da6b7f3e90f5"),
                    ItemDescription: unifiedgosdk.Pointer("The Synchronised motivating software Car offers reliable performance and webbed design"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Gold Cheese"),
                    ItemSku: unifiedgosdk.Pointer("978-0-6550-0228-4"),
                    Notes: unifiedgosdk.Pointer("Utrimque quas vesica spes."),
                    TaxAmount: unifiedgosdk.Pointer[float64](448.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](10875.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](4092.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("3097985d-49fd-4c05-a1a1-3e4b37045097"),
                    ItemDescription: unifiedgosdk.Pointer("Ergonomic Bike made with Marble for all-day coordinated support"),
                    ItemName: unifiedgosdk.Pointer("Practical Steel Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-1-4129-9332-6"),
                    Notes: unifiedgosdk.Pointer("Corrumpo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](875.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](12363.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](6783.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rustic Aluminum Car"),
            PostedAt: types.MustNewTimeFromString("2022-10-31T17:54:00.267Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.220Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.257Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("a75942af-2578-4ed1-a279-dd7ec03f3f10"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("4f9b6182-b488-43dc-8000-8ec2772f46f1"),
                    ItemDescription: unifiedgosdk.Pointer("The Synchronised motivating software Car offers reliable performance and webbed design"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Gold Cheese"),
                    ItemSku: unifiedgosdk.Pointer("978-0-6550-0228-4"),
                    Notes: unifiedgosdk.Pointer("Utrimque quas vesica spes."),
                    TaxAmount: unifiedgosdk.Pointer[float64](448.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](10875.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](4092.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("1157ed49-2e75-4cc1-9521-a1d4aabfaf34"),
                    ItemDescription: unifiedgosdk.Pointer("Ergonomic Bike made with Marble for all-day coordinated support"),
                    ItemName: unifiedgosdk.Pointer("Practical Steel Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-1-4129-9332-6"),
                    Notes: unifiedgosdk.Pointer("Corrumpo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](875.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](12363.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](6783.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rustic Aluminum Car"),
            PostedAt: types.MustNewTimeFromString("2022-10-31T17:54:00.286Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.257Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.257Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("a75942af-2578-4ed1-a279-dd7ec03f3f10"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("4f9b6182-b488-43dc-8000-8ec2772f46f1"),
                    ItemDescription: unifiedgosdk.Pointer("The Synchronised motivating software Car offers reliable performance and webbed design"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Gold Cheese"),
                    ItemSku: unifiedgosdk.Pointer("978-0-6550-0228-4"),
                    Notes: unifiedgosdk.Pointer("Utrimque quas vesica spes."),
                    TaxAmount: unifiedgosdk.Pointer[float64](448.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](10875.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](4092.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("1157ed49-2e75-4cc1-9521-a1d4aabfaf34"),
                    ItemDescription: unifiedgosdk.Pointer("Ergonomic Bike made with Marble for all-day coordinated support"),
                    ItemName: unifiedgosdk.Pointer("Practical Steel Pizza"),
                    ItemSku: unifiedgosdk.Pointer("978-1-4129-9332-6"),
                    Notes: unifiedgosdk.Pointer("Corrumpo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](875.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](12363.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](6783.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rustic Aluminum Car"),
            PostedAt: types.MustNewTimeFromString("2022-10-31T17:54:00.286Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T08:26:37.257Z"),
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