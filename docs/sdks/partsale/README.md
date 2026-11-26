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
            ID: unifiedgosdk.Pointer("8c8fcfa6-dca2-491c-9275-b6a5936d328a"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("3675b135-1418-494c-bbb8-9b768f22c5cd"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("99850d2e-ecde-4846-bc73-e9e1b93c4214"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("49f90062-4f84-478b-a7c0-a8caa949741a"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-04-17T16:03:18.260Z"),
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
            ID: unifiedgosdk.Pointer("c577212b-856c-4afa-b866-510023c35189"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("95290c13-4f44-4bb9-92a2-e5dab8e5e43f"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("08f9368b-c157-4968-bdd5-4af302744207"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("8a489684-dbcc-4342-ab28-1caac8e76023"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-04-17T16:03:18.289Z"),
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
            ID: unifiedgosdk.Pointer("c577212b-856c-4afa-b866-510023c35189"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("95290c13-4f44-4bb9-92a2-e5dab8e5e43f"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("08f9368b-c157-4968-bdd5-4af302744207"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("8a489684-dbcc-4342-ab28-1caac8e76023"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-04-17T16:03:18.289Z"),
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