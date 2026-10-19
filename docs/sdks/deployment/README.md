# Deployment

## Overview

### Available Operations

* [CreateComputeDeployment](#createcomputedeployment) - Create a deployment
* [GetComputeDeployment](#getcomputedeployment) - Retrieve a deployment
* [ListComputeDeployments](#listcomputedeployments) - List all deployments
* [PatchComputeDeployment](#patchcomputedeployment) - Update a deployment
* [RemoveComputeDeployment](#removecomputedeployment) - Remove a deployment
* [UpdateComputeDeployment](#updatecomputedeployment) - Update a deployment

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

    res, err := s.Deployment.CreateComputeDeployment(ctx, operations.CreateComputeDeploymentRequest{
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

    res, err := s.Deployment.GetComputeDeployment(ctx, operations.GetComputeDeploymentRequest{
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

    res, err := s.Deployment.ListComputeDeployments(ctx, operations.ListComputeDeploymentsRequest{
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

    res, err := s.Deployment.PatchComputeDeployment(ctx, operations.PatchComputeDeploymentRequest{
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

    res, err := s.Deployment.RemoveComputeDeployment(ctx, operations.RemoveComputeDeploymentRequest{
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

    res, err := s.Deployment.UpdateComputeDeployment(ctx, operations.UpdateComputeDeploymentRequest{
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