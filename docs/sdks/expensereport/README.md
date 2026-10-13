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
            ApprovedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.067Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("20645595-de23-40b3-9f78-c232fd8a2e9e"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("fe2310cc-6b1c-46d9-8094-05250c72b0ee"),
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
                    ID: unifiedgosdk.Pointer("adae7052-d840-4944-9393-5f0b42beeaef"),
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
            PostedAt: types.MustNewTimeFromString("2022-10-30T21:52:33.515Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.067Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.100Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("40aac9fa-8475-47f0-8896-f91aa7626d33"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("0f56ba59-2d7d-44cf-9d2c-0d081eee980f"),
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
                    ID: unifiedgosdk.Pointer("294a2cdf-5df3-44c4-bf4d-da78b3900029"),
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
            PostedAt: types.MustNewTimeFromString("2022-10-30T21:52:33.531Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.100Z"),
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
            ApprovedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.100Z"),
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-05-29T02:04:59.078Z"),
            Currency: unifiedgosdk.Pointer("BMD"),
            ExternalNumber: unifiedgosdk.Pointer("asporto"),
            ID: unifiedgosdk.Pointer("40aac9fa-8475-47f0-8896-f91aa7626d33"),
            Lineitems: []shared.AccountingLineitem{
                shared.AccountingLineitem{
                    ID: unifiedgosdk.Pointer("0f56ba59-2d7d-44cf-9d2c-0d081eee980f"),
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
                    ID: unifiedgosdk.Pointer("294a2cdf-5df3-44c4-bf4d-da78b3900029"),
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
            PostedAt: types.MustNewTimeFromString("2022-10-30T21:52:33.531Z"),
            ReimbursedAmount: unifiedgosdk.Pointer[float64](6696.0),
            Status: shared.AccountingExpensereportStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](983.0),
            TotalAmount: unifiedgosdk.Pointer[float64](13931.0),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T16:17:01.100Z"),
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