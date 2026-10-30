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

Creates a ComputeDeployment. Provide the unified ComputeDeployment fields in the request body; they are mapped to the provider's API. Returns the created ComputeDeployment.

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
            FinishedAt: types.MustNewTimeFromString("2025-09-05T03:27:42.215Z"),
            ID: unifiedgosdk.Pointer("f867241e-ec5d-4b16-98c5-f5a336689682"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-16T01:07:58.759Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-08T06:41:41.011Z"),
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

Retrieves a single ComputeDeployment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of ComputeDeployments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the ComputeDeployment identified by `id`. Provide the unified ComputeDeployment fields to change in the request body.

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
            FinishedAt: types.MustNewTimeFromString("2025-09-05T03:27:42.225Z"),
            ID: unifiedgosdk.Pointer("a6adf8e7-dded-40a2-8aa9-6c95402f8cbc"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-16T01:07:58.768Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-08T06:41:41.022Z"),
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

Deletes the ComputeDeployment identified by `id`.

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

Updates the ComputeDeployment identified by `id`. Provide the unified ComputeDeployment fields to change in the request body.

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
            FinishedAt: types.MustNewTimeFromString("2025-09-05T03:27:42.225Z"),
            ID: unifiedgosdk.Pointer("a6adf8e7-dded-40a2-8aa9-6c95402f8cbc"),
            Metadata: []shared.ComputeMetadata{},
            StartedAt: types.MustNewTimeFromString("2025-04-16T01:07:58.768Z"),
            Status: shared.ComputeDeploymentStatusLive.ToPointer(),
            Trigger: shared.TriggerAPI.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-08T06:41:41.022Z"),
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