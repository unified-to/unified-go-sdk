# Pipeline

## Overview

### Available Operations

* [CreateCrmPipeline](#createcrmpipeline) - Create a pipeline
* [GetCrmPipeline](#getcrmpipeline) - Retrieve a pipeline
* [ListCrmPipelines](#listcrmpipelines) - List all pipelines
* [PatchCrmPipeline](#patchcrmpipeline) - Update a pipeline
* [RemoveCrmPipeline](#removecrmpipeline) - Remove a pipeline
* [UpdateCrmPipeline](#updatecrmpipeline) - Update a pipeline

## CreateCrmPipeline

Create a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmPipeline" method="post" path="/crm/{connection_id}/pipeline" example="crm_pipeline" -->
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

    res, err := s.Pipeline.CreateCrmPipeline(ctx, operations.CreateCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("29f58939-f7f8-4a3c-b30b-5450a91d92ac"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("910995bc-d696-4327-963b-38bfe0d869fb"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-04T10:29:29.743Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-26T00:48:41.394Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateCrmPipelineRequest](../../pkg/models/operations/createcrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateCrmPipelineResponse](../../pkg/models/operations/createcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmPipeline

Retrieve a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmPipeline" method="get" path="/crm/{connection_id}/pipeline/{id}" -->
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

    res, err := s.Pipeline.GetCrmPipeline(ctx, operations.GetCrmPipelineRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCrmPipelineRequest](../../pkg/models/operations/getcrmpipelinerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetCrmPipelineResponse](../../pkg/models/operations/getcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmPipelines

List all pipelines

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmPipelines" method="get" path="/crm/{connection_id}/pipeline" -->
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

    res, err := s.Pipeline.ListCrmPipelines(ctx, operations.ListCrmPipelinesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmPipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListCrmPipelinesRequest](../../pkg/models/operations/listcrmpipelinesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListCrmPipelinesResponse](../../pkg/models/operations/listcrmpipelinesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmPipeline

Update a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmPipeline" method="patch" path="/crm/{connection_id}/pipeline/{id}" example="crm_pipeline" -->
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

    res, err := s.Pipeline.PatchCrmPipeline(ctx, operations.PatchCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("50bfa9a9-9435-4f13-b4a5-6979f482baae"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("bf41da9d-c06a-4516-98b8-4d792e4ad923"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-04T10:29:29.752Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-26T00:48:41.403Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchCrmPipelineRequest](../../pkg/models/operations/patchcrmpipelinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchCrmPipelineResponse](../../pkg/models/operations/patchcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmPipeline

Remove a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmPipeline" method="delete" path="/crm/{connection_id}/pipeline/{id}" -->
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

    res, err := s.Pipeline.RemoveCrmPipeline(ctx, operations.RemoveCrmPipelineRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveCrmPipelineRequest](../../pkg/models/operations/removecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveCrmPipelineResponse](../../pkg/models/operations/removecrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmPipeline

Update a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmPipeline" method="put" path="/crm/{connection_id}/pipeline/{id}" example="crm_pipeline" -->
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

    res, err := s.Pipeline.UpdateCrmPipeline(ctx, operations.UpdateCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("50bfa9a9-9435-4f13-b4a5-6979f482baae"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("bf41da9d-c06a-4516-98b8-4d792e4ad923"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-04T10:29:29.752Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-26T00:48:41.403Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateCrmPipelineRequest](../../pkg/models/operations/updatecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateCrmPipelineResponse](../../pkg/models/operations/updatecrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |