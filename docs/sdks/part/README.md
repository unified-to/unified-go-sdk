# Part

## Overview

### Available Operations

* [CreatePlmPart](#createplmpart) - Create a part
* [GetPlmPart](#getplmpart) - Retrieve a part
* [ListPlmParts](#listplmparts) - List all parts
* [PatchPlmPart](#patchplmpart) - Update a part
* [UpdatePlmPart](#updateplmpart) - Update a part

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
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("882c84da-46f3-422c-b196-166a8dc6bb63"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("df791fad-a251-486b-a2fa-366aec3907c2"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-10T07:36:09.820Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
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
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("eff0c9c1-b2ae-4e04-b478-9abd7f8aa3ff"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("fca6d557-4351-4c11-8c66-c36ee5363204"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-10T07:36:09.824Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
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
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("eff0c9c1-b2ae-4e04-b478-9abd7f8aa3ff"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("fca6d557-4351-4c11-8c66-c36ee5363204"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-10T07:36:09.824Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
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