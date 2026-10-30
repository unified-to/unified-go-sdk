# Service

## Overview

### Available Operations

* [CreateComputeService](#createcomputeservice) - Create a service
* [GetComputeService](#getcomputeservice) - Retrieve a service
* [ListComputeServices](#listcomputeservices) - List all services
* [PatchComputeService](#patchcomputeservice) - Update a service
* [RemoveComputeService](#removecomputeservice) - Remove a service
* [UpdateComputeService](#updatecomputeservice) - Update a service

## CreateComputeService

Creates a ComputeService. Provide the unified ComputeService fields in the request body; they are mapped to the provider's API. Returns the created ComputeService.

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeService" method="post" path="/compute/{connection_id}/service" example="compute_service" -->
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

    res, err := s.Service.CreateComputeService(ctx, operations.CreateComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("2409720a-c3bb-49a2-99f4-b74f9529e442"),
            InstanceCount: unifiedgosdk.Pointer[float64](6.0),
            MaxInstanceCount: unifiedgosdk.Pointer[float64](6.0),
            Metadata: []shared.ComputeMetadata{},
            MinInstanceCount: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("tame-sideboard"),
            ProjectIdentifier: unifiedgosdk.Pointer("4nyWICvPfi"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/Kimberly_Muller95/tame-sideboard"),
            Runtime: unifiedgosdk.Pointer("python"),
            Status: shared.ComputeServiceStatusSuspended.ToPointer(),
            Type: shared.ComputeServiceTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T00:51:40.824Z"),
            URL: unifiedgosdk.Pointer("https://tame-sideboard.rotating-shoulder.biz"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateComputeServiceRequest](../../pkg/models/operations/createcomputeservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateComputeServiceResponse](../../pkg/models/operations/createcomputeserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeService

Retrieves a single ComputeService by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeService" method="get" path="/compute/{connection_id}/service/{id}" -->
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

    res, err := s.Service.GetComputeService(ctx, operations.GetComputeServiceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetComputeServiceRequest](../../pkg/models/operations/getcomputeservicerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetComputeServiceResponse](../../pkg/models/operations/getcomputeserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeServices

Returns a list of ComputeServices. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeServices" method="get" path="/compute/{connection_id}/service" -->
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

    res, err := s.Service.ListComputeServices(ctx, operations.ListComputeServicesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeServices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListComputeServicesRequest](../../pkg/models/operations/listcomputeservicesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListComputeServicesResponse](../../pkg/models/operations/listcomputeservicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeService

Updates the ComputeService identified by `id`. Provide the unified ComputeService fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeService" method="patch" path="/compute/{connection_id}/service/{id}" example="compute_service" -->
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

    res, err := s.Service.PatchComputeService(ctx, operations.PatchComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("a9792133-7680-4698-8ade-71d0cee4144e"),
            InstanceCount: unifiedgosdk.Pointer[float64](6.0),
            MaxInstanceCount: unifiedgosdk.Pointer[float64](6.0),
            Metadata: []shared.ComputeMetadata{},
            MinInstanceCount: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("tame-sideboard"),
            ProjectIdentifier: unifiedgosdk.Pointer("4nyWICvPfi"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/Kimberly_Muller95/tame-sideboard"),
            Runtime: unifiedgosdk.Pointer("python"),
            Status: shared.ComputeServiceStatusSuspended.ToPointer(),
            Type: shared.ComputeServiceTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T00:51:40.834Z"),
            URL: unifiedgosdk.Pointer("https://tame-sideboard.rotating-shoulder.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchComputeServiceRequest](../../pkg/models/operations/patchcomputeservicerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchComputeServiceResponse](../../pkg/models/operations/patchcomputeserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeService

Deletes the ComputeService identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeService" method="delete" path="/compute/{connection_id}/service/{id}" -->
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

    res, err := s.Service.RemoveComputeService(ctx, operations.RemoveComputeServiceRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveComputeServiceRequest](../../pkg/models/operations/removecomputeservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveComputeServiceResponse](../../pkg/models/operations/removecomputeserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeService

Updates the ComputeService identified by `id`. Provide the unified ComputeService fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeService" method="put" path="/compute/{connection_id}/service/{id}" example="compute_service" -->
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

    res, err := s.Service.UpdateComputeService(ctx, operations.UpdateComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("a9792133-7680-4698-8ade-71d0cee4144e"),
            InstanceCount: unifiedgosdk.Pointer[float64](6.0),
            MaxInstanceCount: unifiedgosdk.Pointer[float64](6.0),
            Metadata: []shared.ComputeMetadata{},
            MinInstanceCount: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("tame-sideboard"),
            ProjectIdentifier: unifiedgosdk.Pointer("4nyWICvPfi"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/Kimberly_Muller95/tame-sideboard"),
            Runtime: unifiedgosdk.Pointer("python"),
            Status: shared.ComputeServiceStatusSuspended.ToPointer(),
            Type: shared.ComputeServiceTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-09T00:51:40.834Z"),
            URL: unifiedgosdk.Pointer("https://tame-sideboard.rotating-shoulder.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateComputeServiceRequest](../../pkg/models/operations/updatecomputeservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateComputeServiceResponse](../../pkg/models/operations/updatecomputeserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |