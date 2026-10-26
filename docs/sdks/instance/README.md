# Instance

## Overview

### Available Operations

* [CreateComputeInstance](#createcomputeinstance) - Create an instance
* [GetComputeInstance](#getcomputeinstance) - Retrieve an instance
* [ListComputeInstances](#listcomputeinstances) - List all instances
* [PatchComputeInstance](#patchcomputeinstance) - Update an instance
* [RemoveComputeInstance](#removecomputeinstance) - Remove an instance
* [UpdateComputeInstance](#updatecomputeinstance) - Update an instance

## CreateComputeInstance

Creates a ComputeInstance. Provide the unified ComputeInstance fields in the request body; they are mapped to the provider's API. Returns the created ComputeInstance.

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeInstance" method="post" path="/compute/{connection_id}/instance" example="compute_instance" -->
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

    res, err := s.Instance.CreateComputeInstance(ctx, operations.CreateComputeInstanceRequest{
        ComputeInstance: shared.ComputeInstance{
            AllowedActions: []shared.PropertyComputeInstanceAllowedActions{
                shared.PropertyComputeInstanceAllowedActionsTerminate,
            },
            Architecture: shared.ComputeInstanceArchitectureOther.ToPointer(),
            CPUCount: unifiedgosdk.Pointer[float64](8.0),
            CreatedAt: types.MustNewTimeFromString("2022-10-07T14:18:53.874Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiskMb: unifiedgosdk.Pointer[float64](1024000.0),
            Endpoints: []shared.ComputeEndpoint{
                shared.ComputeEndpoint{
                    IsAuthRequired: unifiedgosdk.Pointer(true),
                    Port: unifiedgosdk.Pointer[float64](8080.0),
                    Protocol: shared.ProtocolHTTPS.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://serene-shark.biz/"),
                },
            },
            GpuCount: unifiedgosdk.Pointer[float64](0.0),
            HasInternetAccess: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("3fba269d-11bd-4d23-b96f-40911f91bdf6"),
            Image: unifiedgosdk.Pointer("ubuntu:24.04"),
            Market: shared.MarketSpot.ToPointer(),
            MemoryMb: unifiedgosdk.Pointer[float64](2048.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("digital-panel"),
            OnIdle: shared.OnIdleTerminate.ToPointer(),
            PriceAmount: unifiedgosdk.Pointer[float64](24.90336107776966),
            PriceUnit: shared.PriceUnitPerInstanceHour.ToPointer(),
            ProjectIdentifier: unifiedgosdk.Pointer("g547KrivqY"),
            SSH: &shared.PropertyComputeInstanceSSH{
                Host: unifiedgosdk.Pointer("116.42.165.89"),
                Port: unifiedgosdk.Pointer[float64](22.0),
                Username: unifiedgosdk.Pointer("root"),
            },
            StartedAt: types.MustNewTimeFromString("2025-03-26T21:49:12.989Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-19T09:58:28.558Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeInstance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateComputeInstanceRequest](../../pkg/models/operations/createcomputeinstancerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateComputeInstanceResponse](../../pkg/models/operations/createcomputeinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeInstance

Retrieves a single ComputeInstance by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeInstance" method="get" path="/compute/{connection_id}/instance/{id}" -->
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

    res, err := s.Instance.GetComputeInstance(ctx, operations.GetComputeInstanceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeInstance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetComputeInstanceRequest](../../pkg/models/operations/getcomputeinstancerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetComputeInstanceResponse](../../pkg/models/operations/getcomputeinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeInstances

Returns a list of ComputeInstances. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeInstances" method="get" path="/compute/{connection_id}/instance" -->
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

    res, err := s.Instance.ListComputeInstances(ctx, operations.ListComputeInstancesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeInstances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListComputeInstancesRequest](../../pkg/models/operations/listcomputeinstancesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListComputeInstancesResponse](../../pkg/models/operations/listcomputeinstancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeInstance

Updates the ComputeInstance identified by `id`. Provide the unified ComputeInstance fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeInstance" method="patch" path="/compute/{connection_id}/instance/{id}" example="compute_instance" -->
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

    res, err := s.Instance.PatchComputeInstance(ctx, operations.PatchComputeInstanceRequest{
        ComputeInstance: shared.ComputeInstance{
            AllowedActions: []shared.PropertyComputeInstanceAllowedActions{
                shared.PropertyComputeInstanceAllowedActionsTerminate,
            },
            Architecture: shared.ComputeInstanceArchitectureOther.ToPointer(),
            CPUCount: unifiedgosdk.Pointer[float64](8.0),
            CreatedAt: types.MustNewTimeFromString("2022-10-07T14:18:53.874Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiskMb: unifiedgosdk.Pointer[float64](1024000.0),
            Endpoints: []shared.ComputeEndpoint{
                shared.ComputeEndpoint{
                    IsAuthRequired: unifiedgosdk.Pointer(true),
                    Port: unifiedgosdk.Pointer[float64](8080.0),
                    Protocol: shared.ProtocolHTTPS.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://serene-shark.biz/"),
                },
            },
            GpuCount: unifiedgosdk.Pointer[float64](0.0),
            HasInternetAccess: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f65f02aa-7092-4370-baf7-d42f02876606"),
            Image: unifiedgosdk.Pointer("ubuntu:24.04"),
            Market: shared.MarketSpot.ToPointer(),
            MemoryMb: unifiedgosdk.Pointer[float64](2048.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("digital-panel"),
            OnIdle: shared.OnIdleTerminate.ToPointer(),
            PriceAmount: unifiedgosdk.Pointer[float64](24.90336107776966),
            PriceUnit: shared.PriceUnitPerInstanceHour.ToPointer(),
            ProjectIdentifier: unifiedgosdk.Pointer("g547KrivqY"),
            SSH: &shared.PropertyComputeInstanceSSH{
                Host: unifiedgosdk.Pointer("116.42.165.89"),
                Port: unifiedgosdk.Pointer[float64](22.0),
                Username: unifiedgosdk.Pointer("root"),
            },
            StartedAt: types.MustNewTimeFromString("2025-03-26T21:49:13.022Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-19T09:58:28.594Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeInstance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchComputeInstanceRequest](../../pkg/models/operations/patchcomputeinstancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchComputeInstanceResponse](../../pkg/models/operations/patchcomputeinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeInstance

Deletes the ComputeInstance identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeInstance" method="delete" path="/compute/{connection_id}/instance/{id}" -->
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

    res, err := s.Instance.RemoveComputeInstance(ctx, operations.RemoveComputeInstanceRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveComputeInstanceRequest](../../pkg/models/operations/removecomputeinstancerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveComputeInstanceResponse](../../pkg/models/operations/removecomputeinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeInstance

Updates the ComputeInstance identified by `id`. Provide the unified ComputeInstance fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeInstance" method="put" path="/compute/{connection_id}/instance/{id}" example="compute_instance" -->
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

    res, err := s.Instance.UpdateComputeInstance(ctx, operations.UpdateComputeInstanceRequest{
        ComputeInstance: shared.ComputeInstance{
            AllowedActions: []shared.PropertyComputeInstanceAllowedActions{
                shared.PropertyComputeInstanceAllowedActionsTerminate,
            },
            Architecture: shared.ComputeInstanceArchitectureOther.ToPointer(),
            CPUCount: unifiedgosdk.Pointer[float64](8.0),
            CreatedAt: types.MustNewTimeFromString("2022-10-07T14:18:53.874Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiskMb: unifiedgosdk.Pointer[float64](1024000.0),
            Endpoints: []shared.ComputeEndpoint{
                shared.ComputeEndpoint{
                    IsAuthRequired: unifiedgosdk.Pointer(true),
                    Port: unifiedgosdk.Pointer[float64](8080.0),
                    Protocol: shared.ProtocolHTTPS.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://serene-shark.biz/"),
                },
            },
            GpuCount: unifiedgosdk.Pointer[float64](0.0),
            HasInternetAccess: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f65f02aa-7092-4370-baf7-d42f02876606"),
            Image: unifiedgosdk.Pointer("ubuntu:24.04"),
            Market: shared.MarketSpot.ToPointer(),
            MemoryMb: unifiedgosdk.Pointer[float64](2048.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("digital-panel"),
            OnIdle: shared.OnIdleTerminate.ToPointer(),
            PriceAmount: unifiedgosdk.Pointer[float64](24.90336107776966),
            PriceUnit: shared.PriceUnitPerInstanceHour.ToPointer(),
            ProjectIdentifier: unifiedgosdk.Pointer("g547KrivqY"),
            SSH: &shared.PropertyComputeInstanceSSH{
                Host: unifiedgosdk.Pointer("116.42.165.89"),
                Port: unifiedgosdk.Pointer[float64](22.0),
                Username: unifiedgosdk.Pointer("root"),
            },
            StartedAt: types.MustNewTimeFromString("2025-03-26T21:49:13.022Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-19T09:58:28.594Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeInstance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateComputeInstanceRequest](../../pkg/models/operations/updatecomputeinstancerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateComputeInstanceResponse](../../pkg/models/operations/updatecomputeinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |