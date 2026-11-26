# Vehicle

## Overview

### Available Operations

* [CreateDmsVehicle](#createdmsvehicle) - Create a vehicle
* [GetDmsVehicle](#getdmsvehicle) - Retrieve a vehicle
* [ListDmsVehicles](#listdmsvehicles) - List all vehicles
* [PatchDmsVehicle](#patchdmsvehicle) - Update a vehicle
* [UpdateDmsVehicle](#updatedmsvehicle) - Update a vehicle

## CreateDmsVehicle

Creates a DmsVehicle. Provide the unified DmsVehicle fields in the request body; they are mapped to the provider's API. Returns the created DmsVehicle.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsVehicle" method="post" path="/dms/{connection_id}/vehicle" example="dms_vehicle" -->
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

    res, err := s.Vehicle.CreateDmsVehicle(ctx, operations.CreateDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("e02606d0-aeac-4d15-872c-6ad4c43411b3"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T01:42:43.733Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateDmsVehicleRequest](../../pkg/models/operations/createdmsvehiclerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateDmsVehicleResponse](../../pkg/models/operations/createdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsVehicle

Retrieves a single DmsVehicle by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsVehicle" method="get" path="/dms/{connection_id}/vehicle/{id}" -->
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

    res, err := s.Vehicle.GetDmsVehicle(ctx, operations.GetDmsVehicleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDmsVehicleRequest](../../pkg/models/operations/getdmsvehiclerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetDmsVehicleResponse](../../pkg/models/operations/getdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsVehicles

Returns a list of DmsVehicles. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsVehicles" method="get" path="/dms/{connection_id}/vehicle" -->
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

    res, err := s.Vehicle.ListDmsVehicles(ctx, operations.ListDmsVehiclesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListDmsVehiclesRequest](../../pkg/models/operations/listdmsvehiclesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListDmsVehiclesResponse](../../pkg/models/operations/listdmsvehiclesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsVehicle

Updates the DmsVehicle identified by `id`. Provide the unified DmsVehicle fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsVehicle" method="patch" path="/dms/{connection_id}/vehicle/{id}" example="dms_vehicle" -->
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

    res, err := s.Vehicle.PatchDmsVehicle(ctx, operations.PatchDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("79085b50-74a7-4f1d-8d69-d4979c6610bf"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T01:42:43.749Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchDmsVehicleRequest](../../pkg/models/operations/patchdmsvehiclerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchDmsVehicleResponse](../../pkg/models/operations/patchdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsVehicle

Updates the DmsVehicle identified by `id`. Provide the unified DmsVehicle fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsVehicle" method="put" path="/dms/{connection_id}/vehicle/{id}" example="dms_vehicle" -->
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

    res, err := s.Vehicle.UpdateDmsVehicle(ctx, operations.UpdateDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("79085b50-74a7-4f1d-8d69-d4979c6610bf"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T01:42:43.749Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateDmsVehicleRequest](../../pkg/models/operations/updatedmsvehiclerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateDmsVehicleResponse](../../pkg/models/operations/updatedmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |