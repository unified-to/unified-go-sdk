# Pipeline
(*Pipeline*)

### Available Operations

* [DeleteCrmConnectionIDPipelineID](#deletecrmconnectionidpipelineid) - Remove a pipeline
* [GetCrmConnectionIDPipeline](#getcrmconnectionidpipeline) - List all pipelines
* [GetCrmConnectionIDPipelineID](#getcrmconnectionidpipelineid) - Retrieve a pipeline
* [PatchCrmConnectionIDPipelineID](#patchcrmconnectionidpipelineid) - Update a pipeline
* [PostCrmConnectionIDPipeline](#postcrmconnectionidpipeline) - Create a pipeline
* [PutCrmConnectionIDPipelineID](#putcrmconnectionidpipelineid) - Update a pipeline

## DeleteCrmConnectionIDPipelineID

Remove a pipeline

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.DeleteCrmConnectionIDPipelineID(ctx, operations.DeleteCrmConnectionIDPipelineIDRequest{
        ConnectionID: "Customer",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteCrmConnectionIDPipelineIDRequest](../../models/operations/deletecrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.DeleteCrmConnectionIDPipelineIDResponse](../../models/operations/deletecrmconnectionidpipelineidresponse.md), error**


## GetCrmConnectionIDPipeline

List all pipelines

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.GetCrmConnectionIDPipeline(ctx, operations.GetCrmConnectionIDPipelineRequest{
        ConnectionID: "dirty Awesome Checking",
        Limit: unifiedgosdk.Float64(9055.88),
        Offset: unifiedgosdk.Float64(3443.76),
        Order: unifiedgosdk.String("glom"),
        Query: unifiedgosdk.String("panel"),
        Sort: unifiedgosdk.String("Latin tightly"),
        UpdatedGte: types.MustTimeFromString("2022-03-01T15:47:43.244Z"),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCrmConnectionIDPipelineRequest](../../models/operations/getcrmconnectionidpipelinerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetCrmConnectionIDPipelineResponse](../../models/operations/getcrmconnectionidpipelineresponse.md), error**


## GetCrmConnectionIDPipelineID

Retrieve a pipeline

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.GetCrmConnectionIDPipelineID(ctx, operations.GetCrmConnectionIDPipelineIDRequest{
        ConnectionID: "Tricycle roughly markets",
        ID: "<ID>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCrmConnectionIDPipelineIDRequest](../../models/operations/getcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetCrmConnectionIDPipelineIDResponse](../../models/operations/getcrmconnectionidpipelineidresponse.md), error**


## PatchCrmConnectionIDPipelineID

Update a pipeline

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.PatchCrmConnectionIDPipelineID(ctx, operations.PatchCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2023-08-24T17:39:51.183Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(664.58),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("bandwidth"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2023-11-27T01:55:15.440Z"),
        },
        ConnectionID: "Chips",
        ID: "<ID>",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchCrmConnectionIDPipelineIDRequest](../../models/operations/patchcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PatchCrmConnectionIDPipelineIDResponse](../../models/operations/patchcrmconnectionidpipelineidresponse.md), error**


## PostCrmConnectionIDPipeline

Create a pipeline

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.PostCrmConnectionIDPipeline(ctx, operations.PostCrmConnectionIDPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2023-12-10T23:55:22.206Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(3879.73),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("upward Mayaguez"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2021-09-25T10:43:23.679Z"),
        },
        ConnectionID: "Lead Health",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostCrmConnectionIDPipelineRequest](../../models/operations/postcrmconnectionidpipelinerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostCrmConnectionIDPipelineResponse](../../models/operations/postcrmconnectionidpipelineresponse.md), error**


## PutCrmConnectionIDPipelineID

Update a pipeline

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pipeline.PutCrmConnectionIDPipelineID(ctx, operations.PutCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2021-05-16T17:24:47.805Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(5470.76),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("West"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2022-02-28T07:49:31.151Z"),
        },
        ConnectionID: "optimizing",
        ID: "<ID>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutCrmConnectionIDPipelineIDRequest](../../models/operations/putcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PutCrmConnectionIDPipelineIDResponse](../../models/operations/putcrmconnectionidpipelineidresponse.md), error**

