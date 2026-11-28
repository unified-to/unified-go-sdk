# Vehiclesale

## Overview

### Available Operations

* [CreateDmsVehiclesale](#createdmsvehiclesale) - Create a vehiclesale
* [GetDmsVehiclesale](#getdmsvehiclesale) - Retrieve a vehiclesale
* [ListDmsVehiclesales](#listdmsvehiclesales) - List all vehiclesales
* [PatchDmsVehiclesale](#patchdmsvehiclesale) - Update a vehiclesale
* [UpdateDmsVehiclesale](#updatedmsvehiclesale) - Update a vehiclesale

## CreateDmsVehiclesale

Creates a DmsVehiclesale. Provide the unified DmsVehiclesale fields in the request body; they are mapped to the provider's API. Returns the created DmsVehiclesale.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsVehiclesale" method="post" path="/dms/{connection_id}/vehiclesale" example="dms_vehiclesale" -->
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

    res, err := s.Vehiclesale.CreateDmsVehiclesale(ctx, operations.CreateDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T20:04:36.942Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("c1b1798e-c488-44c5-b62b-62f2314f07ca"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-15T04:49:02.655Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-17T15:53:14.000Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDmsVehiclesaleRequest](../../pkg/models/operations/createdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDmsVehiclesaleResponse](../../pkg/models/operations/createdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsVehiclesale

Retrieves a single DmsVehiclesale by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsVehiclesale" method="get" path="/dms/{connection_id}/vehiclesale/{id}" -->
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

    res, err := s.Vehiclesale.GetDmsVehiclesale(ctx, operations.GetDmsVehiclesaleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDmsVehiclesaleRequest](../../pkg/models/operations/getdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDmsVehiclesaleResponse](../../pkg/models/operations/getdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsVehiclesales

Returns a list of DmsVehiclesales. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsVehiclesales" method="get" path="/dms/{connection_id}/vehiclesale" -->
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

    res, err := s.Vehiclesale.ListDmsVehiclesales(ctx, operations.ListDmsVehiclesalesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesales != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDmsVehiclesalesRequest](../../pkg/models/operations/listdmsvehiclesalesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDmsVehiclesalesResponse](../../pkg/models/operations/listdmsvehiclesalesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsVehiclesale

Updates the DmsVehiclesale identified by `id`. Provide the unified DmsVehiclesale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsVehiclesale" method="patch" path="/dms/{connection_id}/vehiclesale/{id}" example="dms_vehiclesale" -->
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

    res, err := s.Vehiclesale.PatchDmsVehiclesale(ctx, operations.PatchDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T20:04:36.953Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("3280b6e0-b0d6-41c5-9219-6bd83e3d2780"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-15T04:49:02.666Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-17T15:53:14.022Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDmsVehiclesaleRequest](../../pkg/models/operations/patchdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDmsVehiclesaleResponse](../../pkg/models/operations/patchdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsVehiclesale

Updates the DmsVehiclesale identified by `id`. Provide the unified DmsVehiclesale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsVehiclesale" method="put" path="/dms/{connection_id}/vehiclesale/{id}" example="dms_vehiclesale" -->
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

    res, err := s.Vehiclesale.UpdateDmsVehiclesale(ctx, operations.UpdateDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T20:04:36.953Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("3280b6e0-b0d6-41c5-9219-6bd83e3d2780"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-15T04:49:02.666Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-17T15:53:14.022Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDmsVehiclesaleRequest](../../pkg/models/operations/updatedmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDmsVehiclesaleResponse](../../pkg/models/operations/updatedmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |