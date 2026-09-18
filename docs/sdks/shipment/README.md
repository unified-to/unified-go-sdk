# Shipment

## Overview

### Available Operations

* [CreateShippingShipment](#createshippingshipment) - Create a shipment
* [GetShippingShipment](#getshippingshipment) - Retrieve a shipment
* [ListShippingShipments](#listshippingshipments) - List all shipments
* [PatchShippingShipment](#patchshippingshipment) - Update a shipment
* [RemoveShippingShipment](#removeshippingshipment) - Remove a shipment
* [UpdateShippingShipment](#updateshippingshipment) - Update a shipment

## CreateShippingShipment

Create a shipment

### Example Usage

<!-- UsageSnippet language="go" operationID="createShippingShipment" method="post" path="/shipping/{connection_id}/shipment" example="shipping_shipment" -->
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

    res, err := s.Shipment.CreateShippingShipment(ctx, operations.CreateShippingShipmentRequest{
        ShippingShipment: shared.ShippingShipment{
            CarrierName: unifiedgosdk.Pointer("Bogisich, Franey and Koelpin"),
            CreatedAt: types.MustNewTimeFromString("2022-09-12T03:11:28.960Z"),
            ID: unifiedgosdk.Pointer("04eb90d8-c4d6-4a12-a014-b29fe96672d2"),
            RateAmount: unifiedgosdk.Pointer[float64](8.86546263936907),
            RateCurrency: unifiedgosdk.Pointer("USD"),
            RateEstimatedDays: unifiedgosdk.Pointer[float64](8.0),
            RateServiceName: unifiedgosdk.Pointer("Fisher - Kilback"),
            ServiceCode: unifiedgosdk.Pointer("F7U"),
            ShippedAt: types.MustNewTimeFromString("2025-08-24T19:04:21.740Z"),
            Status: shared.ShippingShipmentStatusPending.ToPointer(),
            TrackingURL: unifiedgosdk.Pointer("https://shallow-secrecy.info/"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-03T02:49:16.232Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingShipment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateShippingShipmentRequest](../../pkg/models/operations/createshippingshipmentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateShippingShipmentResponse](../../pkg/models/operations/createshippingshipmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetShippingShipment

Retrieve a shipment

### Example Usage

<!-- UsageSnippet language="go" operationID="getShippingShipment" method="get" path="/shipping/{connection_id}/shipment/{id}" -->
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

    res, err := s.Shipment.GetShippingShipment(ctx, operations.GetShippingShipmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingShipment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetShippingShipmentRequest](../../pkg/models/operations/getshippingshipmentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetShippingShipmentResponse](../../pkg/models/operations/getshippingshipmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListShippingShipments

List all shipments

### Example Usage

<!-- UsageSnippet language="go" operationID="listShippingShipments" method="get" path="/shipping/{connection_id}/shipment" -->
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

    res, err := s.Shipment.ListShippingShipments(ctx, operations.ListShippingShipmentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingShipments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListShippingShipmentsRequest](../../pkg/models/operations/listshippingshipmentsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListShippingShipmentsResponse](../../pkg/models/operations/listshippingshipmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchShippingShipment

Update a shipment

### Example Usage

<!-- UsageSnippet language="go" operationID="patchShippingShipment" method="patch" path="/shipping/{connection_id}/shipment/{id}" example="shipping_shipment" -->
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

    res, err := s.Shipment.PatchShippingShipment(ctx, operations.PatchShippingShipmentRequest{
        ShippingShipment: shared.ShippingShipment{
            CarrierName: unifiedgosdk.Pointer("Bogisich, Franey and Koelpin"),
            CreatedAt: types.MustNewTimeFromString("2022-09-12T03:11:28.960Z"),
            ID: unifiedgosdk.Pointer("998c4b1e-2e8c-4fd7-abde-fb56d6f2caa8"),
            RateAmount: unifiedgosdk.Pointer[float64](8.86546263936907),
            RateCurrency: unifiedgosdk.Pointer("USD"),
            RateEstimatedDays: unifiedgosdk.Pointer[float64](8.0),
            RateServiceName: unifiedgosdk.Pointer("Fisher - Kilback"),
            ServiceCode: unifiedgosdk.Pointer("F7U"),
            ShippedAt: types.MustNewTimeFromString("2025-08-24T19:04:21.789Z"),
            Status: shared.ShippingShipmentStatusPending.ToPointer(),
            TrackingURL: unifiedgosdk.Pointer("https://shallow-secrecy.info/"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-03T02:49:16.278Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingShipment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchShippingShipmentRequest](../../pkg/models/operations/patchshippingshipmentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchShippingShipmentResponse](../../pkg/models/operations/patchshippingshipmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveShippingShipment

Remove a shipment

### Example Usage

<!-- UsageSnippet language="go" operationID="removeShippingShipment" method="delete" path="/shipping/{connection_id}/shipment/{id}" -->
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

    res, err := s.Shipment.RemoveShippingShipment(ctx, operations.RemoveShippingShipmentRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveShippingShipmentRequest](../../pkg/models/operations/removeshippingshipmentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveShippingShipmentResponse](../../pkg/models/operations/removeshippingshipmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateShippingShipment

Update a shipment

### Example Usage

<!-- UsageSnippet language="go" operationID="updateShippingShipment" method="put" path="/shipping/{connection_id}/shipment/{id}" example="shipping_shipment" -->
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

    res, err := s.Shipment.UpdateShippingShipment(ctx, operations.UpdateShippingShipmentRequest{
        ShippingShipment: shared.ShippingShipment{
            CarrierName: unifiedgosdk.Pointer("Bogisich, Franey and Koelpin"),
            CreatedAt: types.MustNewTimeFromString("2022-09-12T03:11:28.960Z"),
            ID: unifiedgosdk.Pointer("998c4b1e-2e8c-4fd7-abde-fb56d6f2caa8"),
            RateAmount: unifiedgosdk.Pointer[float64](8.86546263936907),
            RateCurrency: unifiedgosdk.Pointer("USD"),
            RateEstimatedDays: unifiedgosdk.Pointer[float64](8.0),
            RateServiceName: unifiedgosdk.Pointer("Fisher - Kilback"),
            ServiceCode: unifiedgosdk.Pointer("F7U"),
            ShippedAt: types.MustNewTimeFromString("2025-08-24T19:04:21.789Z"),
            Status: shared.ShippingShipmentStatusPending.ToPointer(),
            TrackingURL: unifiedgosdk.Pointer("https://shallow-secrecy.info/"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-03T02:49:16.278Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingShipment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateShippingShipmentRequest](../../pkg/models/operations/updateshippingshipmentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateShippingShipmentResponse](../../pkg/models/operations/updateshippingshipmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |