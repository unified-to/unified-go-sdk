# Part

## Overview

### Available Operations

* [CreateDmsPart](#createdmspart) - Create a part
* [CreatePlmPart](#createplmpart) - Create a part
* [GetDmsPart](#getdmspart) - Retrieve a part
* [GetPlmPart](#getplmpart) - Retrieve a part
* [ListDmsParts](#listdmsparts) - List all parts
* [ListPlmParts](#listplmparts) - List all parts
* [PatchDmsPart](#patchdmspart) - Update a part
* [PatchPlmPart](#patchplmpart) - Update a part
* [RemovePlmPart](#removeplmpart) - Remove a part
* [UpdateDmsPart](#updatedmspart) - Update a part
* [UpdatePlmPart](#updateplmpart) - Update a part

## CreateDmsPart

Creates a DmsPart. Provide the unified DmsPart fields in the request body; they are mapped to the provider's API. Returns the created DmsPart.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsPart" method="post" path="/dms/{connection_id}/part" example="dms_part" -->
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

    res, err := s.Part.CreateDmsPart(ctx, operations.CreateDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44573705-82b3-43dc-a21f-6a8d3e8fdce2"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.426Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.247Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateDmsPartRequest](../../pkg/models/operations/createdmspartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateDmsPartResponse](../../pkg/models/operations/createdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePlmPart

Creates a PlmPart. Provide the unified PlmPart fields in the request body; they are mapped to the provider's API. Returns the created PlmPart.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlmPart" method="post" path="/plm/{connection_id}/part" example="plm_part" -->
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

    res, err := s.Part.CreatePlmPart(ctx, operations.CreatePlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Aestivus talis tero ullam cruciamentum voluptatibus cumque degusto cui."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("d698083d-a16c-41eb-a14e-6045b25bdab5"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("7851c138-7d32-4c26-9cc5-249d7b2d2306"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("speedy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-15T20:32:28.533Z"),
            URL: unifiedgosdk.Pointer("https://digital-pop.com"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreatePlmPartRequest](../../pkg/models/operations/createplmpartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreatePlmPartResponse](../../pkg/models/operations/createplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsPart

Retrieves a single DmsPart by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsPart" method="get" path="/dms/{connection_id}/part/{id}" -->
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

    res, err := s.Part.GetDmsPart(ctx, operations.GetDmsPartRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDmsPartRequest](../../pkg/models/operations/getdmspartrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetDmsPartResponse](../../pkg/models/operations/getdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmPart

Retrieves a single PlmPart by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmPart" method="get" path="/plm/{connection_id}/part/{id}" -->
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

    res, err := s.Part.GetPlmPart(ctx, operations.GetPlmPartRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetPlmPartRequest](../../pkg/models/operations/getplmpartrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetPlmPartResponse](../../pkg/models/operations/getplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsParts

Returns a list of DmsParts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsParts" method="get" path="/dms/{connection_id}/part" -->
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

    res, err := s.Part.ListDmsParts(ctx, operations.ListDmsPartsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsParts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDmsPartsRequest](../../pkg/models/operations/listdmspartsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListDmsPartsResponse](../../pkg/models/operations/listdmspartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmParts

Returns a list of PlmParts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmParts" method="get" path="/plm/{connection_id}/part" -->
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

    res, err := s.Part.ListPlmParts(ctx, operations.ListPlmPartsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmParts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPlmPartsRequest](../../pkg/models/operations/listplmpartsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPlmPartsResponse](../../pkg/models/operations/listplmpartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsPart

Updates the DmsPart identified by `id`. Provide the unified DmsPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsPart" method="patch" path="/dms/{connection_id}/part/{id}" example="dms_part" -->
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

    res, err := s.Part.PatchDmsPart(ctx, operations.PatchDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44358212-7c73-4818-bd09-04036b6311db"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.427Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.256Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchDmsPartRequest](../../pkg/models/operations/patchdmspartrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchDmsPartResponse](../../pkg/models/operations/patchdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPlmPart

Updates the PlmPart identified by `id`. Provide the unified PlmPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPlmPart" method="patch" path="/plm/{connection_id}/part/{id}" example="plm_part" -->
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

    res, err := s.Part.PatchPlmPart(ctx, operations.PatchPlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Aestivus talis tero ullam cruciamentum voluptatibus cumque degusto cui."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("42a2e8db-72b6-492a-a77a-0652980bbdbd"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("10a75b36-e6ae-4479-9bd4-3a22af0c0fd3"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("speedy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-15T20:32:28.538Z"),
            URL: unifiedgosdk.Pointer("https://digital-pop.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchPlmPartRequest](../../pkg/models/operations/patchplmpartrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchPlmPartResponse](../../pkg/models/operations/patchplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePlmPart

Deletes the PlmPart identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePlmPart" method="delete" path="/plm/{connection_id}/part/{id}" -->
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

    res, err := s.Part.RemovePlmPart(ctx, operations.RemovePlmPartRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemovePlmPartRequest](../../pkg/models/operations/removeplmpartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemovePlmPartResponse](../../pkg/models/operations/removeplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsPart

Updates the DmsPart identified by `id`. Provide the unified DmsPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsPart" method="put" path="/dms/{connection_id}/part/{id}" example="dms_part" -->
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

    res, err := s.Part.UpdateDmsPart(ctx, operations.UpdateDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44358212-7c73-4818-bd09-04036b6311db"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.427Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.256Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateDmsPartRequest](../../pkg/models/operations/updatedmspartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateDmsPartResponse](../../pkg/models/operations/updatedmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlmPart

Updates the PlmPart identified by `id`. Provide the unified PlmPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlmPart" method="put" path="/plm/{connection_id}/part/{id}" example="plm_part" -->
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

    res, err := s.Part.UpdatePlmPart(ctx, operations.UpdatePlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Aestivus talis tero ullam cruciamentum voluptatibus cumque degusto cui."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("42a2e8db-72b6-492a-a77a-0652980bbdbd"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("10a75b36-e6ae-4479-9bd4-3a22af0c0fd3"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("speedy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-15T20:32:28.538Z"),
            URL: unifiedgosdk.Pointer("https://digital-pop.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdatePlmPartRequest](../../pkg/models/operations/updateplmpartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdatePlmPartResponse](../../pkg/models/operations/updateplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |