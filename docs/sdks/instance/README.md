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

Create an instance

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
            ID: unifiedgosdk.Pointer("dc70bf59-86bf-4b60-bcd7-2cb992b9ec25"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T03:57:57.474Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T05:26:43.410Z"),
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

Retrieve an instance

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

List all instances

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

Update an instance

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
            ID: unifiedgosdk.Pointer("191c4189-ecd9-4105-9e2a-50148167a68f"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T03:57:57.493Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T05:26:43.431Z"),
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

Remove an instance

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

Update an instance

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
            ID: unifiedgosdk.Pointer("191c4189-ecd9-4105-9e2a-50148167a68f"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T03:57:57.493Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T05:26:43.431Z"),
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