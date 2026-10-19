# Compute

## Overview

### Available Operations

* [CreateComputeDeployment](#createcomputedeployment) - Create a deployment
* [CreateComputeDomain](#createcomputedomain) - Create a domain
* [CreateComputeEnvvar](#createcomputeenvvar) - Create an envvar
* [CreateComputeInstance](#createcomputeinstance) - Create an instance
* [CreateComputeKey](#createcomputekey) - Create a key
* [CreateComputeService](#createcomputeservice) - Create a service
* [CreateComputeSnapshot](#createcomputesnapshot) - Create a snapshot
* [CreateComputeTemplate](#createcomputetemplate) - Create a template
* [CreateComputeVolume](#createcomputevolume) - Create a volume
* [GetComputeDeployment](#getcomputedeployment) - Retrieve a deployment
* [GetComputeDomain](#getcomputedomain) - Retrieve a domain
* [GetComputeEnvvar](#getcomputeenvvar) - Retrieve an envvar
* [GetComputeGputype](#getcomputegputype) - Retrieve a gputype
* [GetComputeInstance](#getcomputeinstance) - Retrieve an instance
* [GetComputeKey](#getcomputekey) - Retrieve a key
* [GetComputeRegion](#getcomputeregion) - Retrieve a region
* [GetComputeService](#getcomputeservice) - Retrieve a service
* [GetComputeSnapshot](#getcomputesnapshot) - Retrieve a snapshot
* [GetComputeTemplate](#getcomputetemplate) - Retrieve a template
* [GetComputeVolume](#getcomputevolume) - Retrieve a volume
* [ListComputeDeployments](#listcomputedeployments) - List all deployments
* [ListComputeDomains](#listcomputedomains) - List all domains
* [ListComputeEnvvars](#listcomputeenvvars) - List all envvars
* [ListComputeGputypes](#listcomputegputypes) - List all gputypes
* [ListComputeInstances](#listcomputeinstances) - List all instances
* [ListComputeKeys](#listcomputekeys) - List all keys
* [ListComputeRegions](#listcomputeregions) - List all regions
* [ListComputeServices](#listcomputeservices) - List all services
* [ListComputeSnapshots](#listcomputesnapshots) - List all snapshots
* [ListComputeTemplates](#listcomputetemplates) - List all templates
* [ListComputeUsages](#listcomputeusages) - List all usages
* [ListComputeVolumes](#listcomputevolumes) - List all volumes
* [PatchComputeDeployment](#patchcomputedeployment) - Update a deployment
* [PatchComputeDomain](#patchcomputedomain) - Update a domain
* [PatchComputeEnvvar](#patchcomputeenvvar) - Update an envvar
* [PatchComputeInstance](#patchcomputeinstance) - Update an instance
* [PatchComputeKey](#patchcomputekey) - Update a key
* [PatchComputeService](#patchcomputeservice) - Update a service
* [PatchComputeTemplate](#patchcomputetemplate) - Update a template
* [PatchComputeVolume](#patchcomputevolume) - Update a volume
* [RemoveComputeDeployment](#removecomputedeployment) - Remove a deployment
* [RemoveComputeDomain](#removecomputedomain) - Remove a domain
* [RemoveComputeEnvvar](#removecomputeenvvar) - Remove an envvar
* [RemoveComputeInstance](#removecomputeinstance) - Remove an instance
* [RemoveComputeKey](#removecomputekey) - Remove a key
* [RemoveComputeService](#removecomputeservice) - Remove a service
* [RemoveComputeSnapshot](#removecomputesnapshot) - Remove a snapshot
* [RemoveComputeTemplate](#removecomputetemplate) - Remove a template
* [RemoveComputeVolume](#removecomputevolume) - Remove a volume
* [UpdateComputeDeployment](#updatecomputedeployment) - Update a deployment
* [UpdateComputeDomain](#updatecomputedomain) - Update a domain
* [UpdateComputeEnvvar](#updatecomputeenvvar) - Update an envvar
* [UpdateComputeInstance](#updatecomputeinstance) - Update an instance
* [UpdateComputeKey](#updatecomputekey) - Update a key
* [UpdateComputeService](#updatecomputeservice) - Update a service
* [UpdateComputeTemplate](#updatecomputetemplate) - Update a template
* [UpdateComputeVolume](#updatecomputevolume) - Update a volume

## CreateComputeDeployment

Create a deployment

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeDeployment" method="post" path="/compute/{connection_id}/deployment" example="compute_deployment" -->
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

    res, err := s.Compute.CreateComputeDeployment(ctx, operations.CreateComputeDeploymentRequest{
        ComputeDeployment: shared.ComputeDeployment{
            Branch: unifiedgosdk.Pointer("develop"),
            CommitIdentifier: unifiedgosdk.Pointer("fadfe54eefd8fea670bea6d58b4b15d406c4b02b"),
            CommitMessage: unifiedgosdk.Pointer("copy primary pixel"),
            CreatedAt: types.MustNewTimeFromString("2022-06-07T03:52:04.613Z"),
            Environment: shared.EnvironmentPreview.ToPointer(),
            FinishedAt: types.MustNewTimeFromString("2025-08-27T22:26:19.845Z"),
            ID: unifiedgosdk.Pointer("fcd0132b-8913-4c6e-b9d0-f83feef763b6"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-08T19:42:58.140Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-26T19:04:56.071Z"),
            URL: unifiedgosdk.Pointer("https://comfortable-petal.name"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDeployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateComputeDeploymentRequest](../../pkg/models/operations/createcomputedeploymentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateComputeDeploymentResponse](../../pkg/models/operations/createcomputedeploymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComputeDomain

Create a domain

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeDomain" method="post" path="/compute/{connection_id}/domain" example="compute_domain" -->
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

    res, err := s.Compute.CreateComputeDomain(ctx, operations.CreateComputeDomainRequest{
        ComputeDomain: shared.ComputeDomain{
            CreatedAt: types.MustNewTimeFromString("2019-11-22T16:38:26.009Z"),
            ID: unifiedgosdk.Pointer("284ea11e-a8c9-4b09-9e9d-19f5ca96da5e"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("mysterious-republican.usable-section.info"),
            Status: shared.ComputeDomainStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-02T07:45:13.198Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateComputeDomainRequest](../../pkg/models/operations/createcomputedomainrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateComputeDomainResponse](../../pkg/models/operations/createcomputedomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComputeEnvvar

Create an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeEnvvar" method="post" path="/compute/{connection_id}/envvar" example="compute_envvar" -->
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

    res, err := s.Compute.CreateComputeEnvvar(ctx, operations.CreateComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("4e7d1866-2a07-4ac8-a38b-99978be09162"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-01T21:07:37.137Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateComputeEnvvarRequest](../../pkg/models/operations/createcomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateComputeEnvvarResponse](../../pkg/models/operations/createcomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Compute.CreateComputeInstance(ctx, operations.CreateComputeInstanceRequest{
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
            ID: unifiedgosdk.Pointer("fc7522b0-e4d6-40d3-b0f8-b0082869e2c5"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T13:45:44.784Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T16:09:37.358Z"),
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

## CreateComputeKey

Create a key

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeKey" method="post" path="/compute/{connection_id}/key" example="compute_key" -->
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

    res, err := s.Compute.CreateComputeKey(ctx, operations.CreateComputeKeyRequest{
        ComputeKey: shared.ComputeKey{
            CreatedAt: types.MustNewTimeFromString("2023-08-19T14:31:46.492Z"),
            ID: unifiedgosdk.Pointer("7644ff30-1cfe-49d3-8724-c11cbe729872"),
            Name: unifiedgosdk.Pointer("nienow-secret"),
            Type: shared.ComputeKeyTypeSecret.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-10-06T03:39:17.807Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateComputeKeyRequest](../../pkg/models/operations/createcomputekeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateComputeKeyResponse](../../pkg/models/operations/createcomputekeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComputeService

Create a service

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

    res, err := s.Compute.CreateComputeService(ctx, operations.CreateComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("3a2bd7fa-2a89-4e93-8033-68f79b1b3735"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T13:22:21.819Z"),
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

## CreateComputeSnapshot

Create a snapshot

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeSnapshot" method="post" path="/compute/{connection_id}/snapshot" example="compute_snapshot" -->
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

    res, err := s.Compute.CreateComputeSnapshot(ctx, operations.CreateComputeSnapshotRequest{
        ComputeSnapshot: shared.ComputeSnapshot{
            CreatedAt: types.MustNewTimeFromString("2021-09-24T14:40:06.110Z"),
            ID: unifiedgosdk.Pointer("36d7de56-508d-48be-8c85-126b709060b3"),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("snap-alius"),
            SizeMb: unifiedgosdk.Pointer[float64](169984.0),
            Status: shared.ComputeSnapshotStatusDeleting.ToPointer(),
            Type: shared.ComputeSnapshotTypeDisk.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-28T20:25:34.271Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeSnapshot != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateComputeSnapshotRequest](../../pkg/models/operations/createcomputesnapshotrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateComputeSnapshotResponse](../../pkg/models/operations/createcomputesnapshotresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComputeTemplate

Create a template

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeTemplate" method="post" path="/compute/{connection_id}/template" example="compute_template" -->
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

    res, err := s.Compute.CreateComputeTemplate(ctx, operations.CreateComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("30a8f9b7-0d02-45bf-ad36-04c9fb8b7bba"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-21T22:06:31.039Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-11T12:57:09.050Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateComputeTemplateRequest](../../pkg/models/operations/createcomputetemplaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateComputeTemplateResponse](../../pkg/models/operations/createcomputetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComputeVolume

Create a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeVolume" method="post" path="/compute/{connection_id}/volume" example="compute_volume" -->
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

    res, err := s.Compute.CreateComputeVolume(ctx, operations.CreateComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("0b2f2d44-a8c4-4b19-a3d5-ac91291a532f"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-09T04:32:49.326Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateComputeVolumeRequest](../../pkg/models/operations/createcomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateComputeVolumeResponse](../../pkg/models/operations/createcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeDeployment

Retrieve a deployment

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeDeployment" method="get" path="/compute/{connection_id}/deployment/{id}" -->
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

    res, err := s.Compute.GetComputeDeployment(ctx, operations.GetComputeDeploymentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDeployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetComputeDeploymentRequest](../../pkg/models/operations/getcomputedeploymentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetComputeDeploymentResponse](../../pkg/models/operations/getcomputedeploymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeDomain

Retrieve a domain

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeDomain" method="get" path="/compute/{connection_id}/domain/{id}" -->
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

    res, err := s.Compute.GetComputeDomain(ctx, operations.GetComputeDomainRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeDomainRequest](../../pkg/models/operations/getcomputedomainrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeDomainResponse](../../pkg/models/operations/getcomputedomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeEnvvar

Retrieve an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeEnvvar" method="get" path="/compute/{connection_id}/envvar/{id}" -->
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

    res, err := s.Compute.GetComputeEnvvar(ctx, operations.GetComputeEnvvarRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeEnvvarRequest](../../pkg/models/operations/getcomputeenvvarrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeEnvvarResponse](../../pkg/models/operations/getcomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeGputype

Retrieve a gputype

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeGputype" method="get" path="/compute/{connection_id}/gputype/{id}" -->
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

    res, err := s.Compute.GetComputeGputype(ctx, operations.GetComputeGputypeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeGputype != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetComputeGputypeRequest](../../pkg/models/operations/getcomputegputyperequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetComputeGputypeResponse](../../pkg/models/operations/getcomputegputyperesponse.md), error**

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

    res, err := s.Compute.GetComputeInstance(ctx, operations.GetComputeInstanceRequest{
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

## GetComputeKey

Retrieve a key

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeKey" method="get" path="/compute/{connection_id}/key/{id}" -->
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

    res, err := s.Compute.GetComputeKey(ctx, operations.GetComputeKeyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetComputeKeyRequest](../../pkg/models/operations/getcomputekeyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetComputeKeyResponse](../../pkg/models/operations/getcomputekeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeRegion

Retrieve a region

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeRegion" method="get" path="/compute/{connection_id}/region/{id}" -->
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

    res, err := s.Compute.GetComputeRegion(ctx, operations.GetComputeRegionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeRegion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeRegionRequest](../../pkg/models/operations/getcomputeregionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeRegionResponse](../../pkg/models/operations/getcomputeregionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeService

Retrieve a service

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

    res, err := s.Compute.GetComputeService(ctx, operations.GetComputeServiceRequest{
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

## GetComputeSnapshot

Retrieve a snapshot

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeSnapshot" method="get" path="/compute/{connection_id}/snapshot/{id}" -->
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

    res, err := s.Compute.GetComputeSnapshot(ctx, operations.GetComputeSnapshotRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeSnapshot != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetComputeSnapshotRequest](../../pkg/models/operations/getcomputesnapshotrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetComputeSnapshotResponse](../../pkg/models/operations/getcomputesnapshotresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeTemplate

Retrieve a template

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeTemplate" method="get" path="/compute/{connection_id}/template/{id}" -->
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

    res, err := s.Compute.GetComputeTemplate(ctx, operations.GetComputeTemplateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetComputeTemplateRequest](../../pkg/models/operations/getcomputetemplaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetComputeTemplateResponse](../../pkg/models/operations/getcomputetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeVolume

Retrieve a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeVolume" method="get" path="/compute/{connection_id}/volume/{id}" -->
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

    res, err := s.Compute.GetComputeVolume(ctx, operations.GetComputeVolumeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeVolumeRequest](../../pkg/models/operations/getcomputevolumerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeVolumeResponse](../../pkg/models/operations/getcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeDeployments

List all deployments

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeDeployments" method="get" path="/compute/{connection_id}/deployment" -->
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

    res, err := s.Compute.ListComputeDeployments(ctx, operations.ListComputeDeploymentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDeployments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListComputeDeploymentsRequest](../../pkg/models/operations/listcomputedeploymentsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListComputeDeploymentsResponse](../../pkg/models/operations/listcomputedeploymentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeDomains

List all domains

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeDomains" method="get" path="/compute/{connection_id}/domain" -->
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

    res, err := s.Compute.ListComputeDomains(ctx, operations.ListComputeDomainsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDomains != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeDomainsRequest](../../pkg/models/operations/listcomputedomainsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeDomainsResponse](../../pkg/models/operations/listcomputedomainsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeEnvvars

List all envvars

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeEnvvars" method="get" path="/compute/{connection_id}/envvar" -->
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

    res, err := s.Compute.ListComputeEnvvars(ctx, operations.ListComputeEnvvarsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeEnvvarsRequest](../../pkg/models/operations/listcomputeenvvarsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeEnvvarsResponse](../../pkg/models/operations/listcomputeenvvarsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeGputypes

List all gputypes

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeGputypes" method="get" path="/compute/{connection_id}/gputype" -->
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

    res, err := s.Compute.ListComputeGputypes(ctx, operations.ListComputeGputypesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeGputypes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListComputeGputypesRequest](../../pkg/models/operations/listcomputegputypesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListComputeGputypesResponse](../../pkg/models/operations/listcomputegputypesresponse.md), error**

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

    res, err := s.Compute.ListComputeInstances(ctx, operations.ListComputeInstancesRequest{
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

## ListComputeKeys

List all keys

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeKeys" method="get" path="/compute/{connection_id}/key" -->
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

    res, err := s.Compute.ListComputeKeys(ctx, operations.ListComputeKeysRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListComputeKeysRequest](../../pkg/models/operations/listcomputekeysrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListComputeKeysResponse](../../pkg/models/operations/listcomputekeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeRegions

List all regions

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeRegions" method="get" path="/compute/{connection_id}/region" -->
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

    res, err := s.Compute.ListComputeRegions(ctx, operations.ListComputeRegionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeRegions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeRegionsRequest](../../pkg/models/operations/listcomputeregionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeRegionsResponse](../../pkg/models/operations/listcomputeregionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeServices

List all services

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

    res, err := s.Compute.ListComputeServices(ctx, operations.ListComputeServicesRequest{
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

## ListComputeSnapshots

List all snapshots

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeSnapshots" method="get" path="/compute/{connection_id}/snapshot" -->
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

    res, err := s.Compute.ListComputeSnapshots(ctx, operations.ListComputeSnapshotsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeSnapshots != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListComputeSnapshotsRequest](../../pkg/models/operations/listcomputesnapshotsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListComputeSnapshotsResponse](../../pkg/models/operations/listcomputesnapshotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeTemplates

List all templates

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeTemplates" method="get" path="/compute/{connection_id}/template" -->
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

    res, err := s.Compute.ListComputeTemplates(ctx, operations.ListComputeTemplatesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListComputeTemplatesRequest](../../pkg/models/operations/listcomputetemplatesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListComputeTemplatesResponse](../../pkg/models/operations/listcomputetemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeUsages

List all usages

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeUsages" method="get" path="/compute/{connection_id}/usage" -->
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

    res, err := s.Compute.ListComputeUsages(ctx, operations.ListComputeUsagesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeUsages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListComputeUsagesRequest](../../pkg/models/operations/listcomputeusagesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListComputeUsagesResponse](../../pkg/models/operations/listcomputeusagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeVolumes

List all volumes

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeVolumes" method="get" path="/compute/{connection_id}/volume" -->
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

    res, err := s.Compute.ListComputeVolumes(ctx, operations.ListComputeVolumesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolumes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeVolumesRequest](../../pkg/models/operations/listcomputevolumesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeVolumesResponse](../../pkg/models/operations/listcomputevolumesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeDeployment

Update a deployment

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeDeployment" method="patch" path="/compute/{connection_id}/deployment/{id}" example="compute_deployment" -->
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

    res, err := s.Compute.PatchComputeDeployment(ctx, operations.PatchComputeDeploymentRequest{
        ComputeDeployment: shared.ComputeDeployment{
            Branch: unifiedgosdk.Pointer("develop"),
            CommitIdentifier: unifiedgosdk.Pointer("fadfe54eefd8fea670bea6d58b4b15d406c4b02b"),
            CommitMessage: unifiedgosdk.Pointer("copy primary pixel"),
            CreatedAt: types.MustNewTimeFromString("2022-06-07T03:52:04.613Z"),
            Environment: shared.EnvironmentPreview.ToPointer(),
            FinishedAt: types.MustNewTimeFromString("2025-08-27T22:26:19.861Z"),
            ID: unifiedgosdk.Pointer("ba682ec1-7932-4d1f-a3f1-660e0a8fcb45"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-08T19:42:58.154Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-26T19:04:56.090Z"),
            URL: unifiedgosdk.Pointer("https://comfortable-petal.name"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDeployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchComputeDeploymentRequest](../../pkg/models/operations/patchcomputedeploymentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchComputeDeploymentResponse](../../pkg/models/operations/patchcomputedeploymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeDomain

Update a domain

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeDomain" method="patch" path="/compute/{connection_id}/domain/{id}" example="compute_domain" -->
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

    res, err := s.Compute.PatchComputeDomain(ctx, operations.PatchComputeDomainRequest{
        ComputeDomain: shared.ComputeDomain{
            CreatedAt: types.MustNewTimeFromString("2019-11-22T16:38:26.009Z"),
            ID: unifiedgosdk.Pointer("4af6a842-f884-40f0-8922-444620feedbd"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("mysterious-republican.usable-section.info"),
            Status: shared.ComputeDomainStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-02T07:45:13.203Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchComputeDomainRequest](../../pkg/models/operations/patchcomputedomainrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchComputeDomainResponse](../../pkg/models/operations/patchcomputedomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeEnvvar

Update an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeEnvvar" method="patch" path="/compute/{connection_id}/envvar/{id}" example="compute_envvar" -->
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

    res, err := s.Compute.PatchComputeEnvvar(ctx, operations.PatchComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("bcfbd67c-6a84-43fb-879f-33feb6b6ca49"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-01T21:07:37.141Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchComputeEnvvarRequest](../../pkg/models/operations/patchcomputeenvvarrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchComputeEnvvarResponse](../../pkg/models/operations/patchcomputeenvvarresponse.md), error**

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

    res, err := s.Compute.PatchComputeInstance(ctx, operations.PatchComputeInstanceRequest{
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
            ID: unifiedgosdk.Pointer("2d370752-fea8-4b4d-80f5-2779689fc905"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T13:45:44.809Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T16:09:37.385Z"),
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

## PatchComputeKey

Update a key

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeKey" method="patch" path="/compute/{connection_id}/key/{id}" example="compute_key" -->
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

    res, err := s.Compute.PatchComputeKey(ctx, operations.PatchComputeKeyRequest{
        ComputeKey: shared.ComputeKey{
            CreatedAt: types.MustNewTimeFromString("2023-08-19T14:31:46.492Z"),
            ID: unifiedgosdk.Pointer("f989821c-be54-44e5-ac47-9bb0dd1df77d"),
            Name: unifiedgosdk.Pointer("nienow-secret"),
            Type: shared.ComputeKeyTypeSecret.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-10-06T03:39:17.810Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchComputeKeyRequest](../../pkg/models/operations/patchcomputekeyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchComputeKeyResponse](../../pkg/models/operations/patchcomputekeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeService

Update a service

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

    res, err := s.Compute.PatchComputeService(ctx, operations.PatchComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("90e20fad-1e3f-45d6-8d7b-7d8b3a93a36d"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T13:22:21.831Z"),
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

## PatchComputeTemplate

Update a template

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeTemplate" method="patch" path="/compute/{connection_id}/template/{id}" example="compute_template" -->
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

    res, err := s.Compute.PatchComputeTemplate(ctx, operations.PatchComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("ae84cf0f-a80b-435f-b4e3-af935033937c"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-21T22:06:31.058Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-11T12:57:09.062Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchComputeTemplateRequest](../../pkg/models/operations/patchcomputetemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchComputeTemplateResponse](../../pkg/models/operations/patchcomputetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeVolume

Update a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeVolume" method="patch" path="/compute/{connection_id}/volume/{id}" example="compute_volume" -->
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

    res, err := s.Compute.PatchComputeVolume(ctx, operations.PatchComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("605b9de2-5348-417b-ac0f-cbd800c43b44"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-09T04:32:49.330Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchComputeVolumeRequest](../../pkg/models/operations/patchcomputevolumerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchComputeVolumeResponse](../../pkg/models/operations/patchcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeDeployment

Remove a deployment

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeDeployment" method="delete" path="/compute/{connection_id}/deployment/{id}" -->
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

    res, err := s.Compute.RemoveComputeDeployment(ctx, operations.RemoveComputeDeploymentRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveComputeDeploymentRequest](../../pkg/models/operations/removecomputedeploymentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveComputeDeploymentResponse](../../pkg/models/operations/removecomputedeploymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeDomain

Remove a domain

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeDomain" method="delete" path="/compute/{connection_id}/domain/{id}" -->
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

    res, err := s.Compute.RemoveComputeDomain(ctx, operations.RemoveComputeDomainRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveComputeDomainRequest](../../pkg/models/operations/removecomputedomainrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveComputeDomainResponse](../../pkg/models/operations/removecomputedomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeEnvvar

Remove an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeEnvvar" method="delete" path="/compute/{connection_id}/envvar/{id}" -->
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

    res, err := s.Compute.RemoveComputeEnvvar(ctx, operations.RemoveComputeEnvvarRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveComputeEnvvarRequest](../../pkg/models/operations/removecomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveComputeEnvvarResponse](../../pkg/models/operations/removecomputeenvvarresponse.md), error**

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

    res, err := s.Compute.RemoveComputeInstance(ctx, operations.RemoveComputeInstanceRequest{
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

## RemoveComputeKey

Remove a key

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeKey" method="delete" path="/compute/{connection_id}/key/{id}" -->
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

    res, err := s.Compute.RemoveComputeKey(ctx, operations.RemoveComputeKeyRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveComputeKeyRequest](../../pkg/models/operations/removecomputekeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveComputeKeyResponse](../../pkg/models/operations/removecomputekeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeService

Remove a service

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

    res, err := s.Compute.RemoveComputeService(ctx, operations.RemoveComputeServiceRequest{
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

## RemoveComputeSnapshot

Remove a snapshot

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeSnapshot" method="delete" path="/compute/{connection_id}/snapshot/{id}" -->
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

    res, err := s.Compute.RemoveComputeSnapshot(ctx, operations.RemoveComputeSnapshotRequest{
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
| `request`                                                                                              | [operations.RemoveComputeSnapshotRequest](../../pkg/models/operations/removecomputesnapshotrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveComputeSnapshotResponse](../../pkg/models/operations/removecomputesnapshotresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeTemplate

Remove a template

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeTemplate" method="delete" path="/compute/{connection_id}/template/{id}" -->
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

    res, err := s.Compute.RemoveComputeTemplate(ctx, operations.RemoveComputeTemplateRequest{
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
| `request`                                                                                              | [operations.RemoveComputeTemplateRequest](../../pkg/models/operations/removecomputetemplaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveComputeTemplateResponse](../../pkg/models/operations/removecomputetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeVolume

Remove a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeVolume" method="delete" path="/compute/{connection_id}/volume/{id}" -->
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

    res, err := s.Compute.RemoveComputeVolume(ctx, operations.RemoveComputeVolumeRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveComputeVolumeRequest](../../pkg/models/operations/removecomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveComputeVolumeResponse](../../pkg/models/operations/removecomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeDeployment

Update a deployment

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeDeployment" method="put" path="/compute/{connection_id}/deployment/{id}" example="compute_deployment" -->
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

    res, err := s.Compute.UpdateComputeDeployment(ctx, operations.UpdateComputeDeploymentRequest{
        ComputeDeployment: shared.ComputeDeployment{
            Branch: unifiedgosdk.Pointer("develop"),
            CommitIdentifier: unifiedgosdk.Pointer("fadfe54eefd8fea670bea6d58b4b15d406c4b02b"),
            CommitMessage: unifiedgosdk.Pointer("copy primary pixel"),
            CreatedAt: types.MustNewTimeFromString("2022-06-07T03:52:04.613Z"),
            Environment: shared.EnvironmentPreview.ToPointer(),
            FinishedAt: types.MustNewTimeFromString("2025-08-27T22:26:19.861Z"),
            ID: unifiedgosdk.Pointer("ba682ec1-7932-4d1f-a3f1-660e0a8fcb45"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-08T19:42:58.154Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-26T19:04:56.090Z"),
            URL: unifiedgosdk.Pointer("https://comfortable-petal.name"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDeployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateComputeDeploymentRequest](../../pkg/models/operations/updatecomputedeploymentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateComputeDeploymentResponse](../../pkg/models/operations/updatecomputedeploymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeDomain

Update a domain

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeDomain" method="put" path="/compute/{connection_id}/domain/{id}" example="compute_domain" -->
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

    res, err := s.Compute.UpdateComputeDomain(ctx, operations.UpdateComputeDomainRequest{
        ComputeDomain: shared.ComputeDomain{
            CreatedAt: types.MustNewTimeFromString("2019-11-22T16:38:26.009Z"),
            ID: unifiedgosdk.Pointer("4af6a842-f884-40f0-8922-444620feedbd"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("mysterious-republican.usable-section.info"),
            Status: shared.ComputeDomainStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-02T07:45:13.203Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateComputeDomainRequest](../../pkg/models/operations/updatecomputedomainrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateComputeDomainResponse](../../pkg/models/operations/updatecomputedomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeEnvvar

Update an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeEnvvar" method="put" path="/compute/{connection_id}/envvar/{id}" example="compute_envvar" -->
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

    res, err := s.Compute.UpdateComputeEnvvar(ctx, operations.UpdateComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("bcfbd67c-6a84-43fb-879f-33feb6b6ca49"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-01T21:07:37.141Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateComputeEnvvarRequest](../../pkg/models/operations/updatecomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateComputeEnvvarResponse](../../pkg/models/operations/updatecomputeenvvarresponse.md), error**

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

    res, err := s.Compute.UpdateComputeInstance(ctx, operations.UpdateComputeInstanceRequest{
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
            ID: unifiedgosdk.Pointer("2d370752-fea8-4b4d-80f5-2779689fc905"),
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
            StartedAt: types.MustNewTimeFromString("2025-03-22T13:45:44.809Z"),
            Status: shared.ComputeInstanceStatusError.ToPointer(),
            StatusReason: unifiedgosdk.Pointer("Sponte quasi adversus alioqui reiciendis arca carmen canis."),
            Type: shared.ComputeInstanceTypeBaremetal.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-14T16:09:37.385Z"),
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

## UpdateComputeKey

Update a key

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeKey" method="put" path="/compute/{connection_id}/key/{id}" example="compute_key" -->
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

    res, err := s.Compute.UpdateComputeKey(ctx, operations.UpdateComputeKeyRequest{
        ComputeKey: shared.ComputeKey{
            CreatedAt: types.MustNewTimeFromString("2023-08-19T14:31:46.492Z"),
            ID: unifiedgosdk.Pointer("f989821c-be54-44e5-ac47-9bb0dd1df77d"),
            Name: unifiedgosdk.Pointer("nienow-secret"),
            Type: shared.ComputeKeyTypeSecret.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-10-06T03:39:17.810Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateComputeKeyRequest](../../pkg/models/operations/updatecomputekeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateComputeKeyResponse](../../pkg/models/operations/updatecomputekeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeService

Update a service

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

    res, err := s.Compute.UpdateComputeService(ctx, operations.UpdateComputeServiceRequest{
        ComputeService: shared.ComputeService{
            Branch: unifiedgosdk.Pointer("main"),
            BuildCommand: unifiedgosdk.Pointer("npm run build"),
            CreatedAt: types.MustNewTimeFromString("2023-10-04T02:06:29.391Z"),
            ID: unifiedgosdk.Pointer("90e20fad-1e3f-45d6-8d7b-7d8b3a93a36d"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T13:22:21.831Z"),
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

## UpdateComputeTemplate

Update a template

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeTemplate" method="put" path="/compute/{connection_id}/template/{id}" example="compute_template" -->
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

    res, err := s.Compute.UpdateComputeTemplate(ctx, operations.UpdateComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("ae84cf0f-a80b-435f-b4e3-af935033937c"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-21T22:06:31.058Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-11T12:57:09.062Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateComputeTemplateRequest](../../pkg/models/operations/updatecomputetemplaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateComputeTemplateResponse](../../pkg/models/operations/updatecomputetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeVolume

Update a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeVolume" method="put" path="/compute/{connection_id}/volume/{id}" example="compute_volume" -->
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

    res, err := s.Compute.UpdateComputeVolume(ctx, operations.UpdateComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("605b9de2-5348-417b-ac0f-cbd800c43b44"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-09T04:32:49.330Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateComputeVolumeRequest](../../pkg/models/operations/updatecomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateComputeVolumeResponse](../../pkg/models/operations/updatecomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |