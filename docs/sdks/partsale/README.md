# Partsale

## Overview

### Available Operations

* [CreateDmsPartsale](#createdmspartsale) - Create a partsale
* [GetDmsPartsale](#getdmspartsale) - Retrieve a partsale
* [ListDmsPartsales](#listdmspartsales) - List all partsales
* [PatchDmsPartsale](#patchdmspartsale) - Update a partsale
* [UpdateDmsPartsale](#updatedmspartsale) - Update a partsale

## CreateDmsPartsale

Creates a DmsPartsale. Provide the unified DmsPartsale fields in the request body; they are mapped to the provider's API. Returns the created DmsPartsale.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsPartsale" method="post" path="/dms/{connection_id}/partsale" example="dms_partsale" -->
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

    res, err := s.Partsale.CreateDmsPartsale(ctx, operations.CreateDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("bd08d54d-448b-4d62-8768-b6502973e230"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("caca2735-8164-423c-84f4-488aa8c75ac9"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("8e6e84b8-fb04-45e5-a51e-06a1cd4dc803"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("dc2e64c1-cc98-4b72-a4fa-f67dd4941ab9"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.883Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateDmsPartsaleRequest](../../pkg/models/operations/createdmspartsalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateDmsPartsaleResponse](../../pkg/models/operations/createdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsPartsale

Retrieves a single DmsPartsale by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsPartsale" method="get" path="/dms/{connection_id}/partsale/{id}" -->
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

    res, err := s.Partsale.GetDmsPartsale(ctx, operations.GetDmsPartsaleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetDmsPartsaleRequest](../../pkg/models/operations/getdmspartsalerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetDmsPartsaleResponse](../../pkg/models/operations/getdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsPartsales

Returns a list of DmsPartsales. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsPartsales" method="get" path="/dms/{connection_id}/partsale" -->
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

    res, err := s.Partsale.ListDmsPartsales(ctx, operations.ListDmsPartsalesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsales != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListDmsPartsalesRequest](../../pkg/models/operations/listdmspartsalesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListDmsPartsalesResponse](../../pkg/models/operations/listdmspartsalesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsPartsale

Updates the DmsPartsale identified by `id`. Provide the unified DmsPartsale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsPartsale" method="patch" path="/dms/{connection_id}/partsale/{id}" example="dms_partsale" -->
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

    res, err := s.Partsale.PatchDmsPartsale(ctx, operations.PatchDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("c8f3d6d4-f9f0-4bf3-9721-4d7e489e98ba"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("13b3eda3-f590-4de6-a13d-499d910cfc23"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("939aa1f5-3dfd-4b82-838b-240da4c49ea6"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("f65e363c-3828-4710-909c-2b319567d354"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.920Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchDmsPartsaleRequest](../../pkg/models/operations/patchdmspartsalerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchDmsPartsaleResponse](../../pkg/models/operations/patchdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsPartsale

Updates the DmsPartsale identified by `id`. Provide the unified DmsPartsale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsPartsale" method="put" path="/dms/{connection_id}/partsale/{id}" example="dms_partsale" -->
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

    res, err := s.Partsale.UpdateDmsPartsale(ctx, operations.UpdateDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("c8f3d6d4-f9f0-4bf3-9721-4d7e489e98ba"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("13b3eda3-f590-4de6-a13d-499d910cfc23"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("939aa1f5-3dfd-4b82-838b-240da4c49ea6"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("f65e363c-3828-4710-909c-2b319567d354"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.920Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateDmsPartsaleRequest](../../pkg/models/operations/updatedmspartsalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateDmsPartsaleResponse](../../pkg/models/operations/updatedmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |