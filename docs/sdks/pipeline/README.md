# Pipeline

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDPipelineIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.DeleteCrmConnectionIDPipelineID(ctx, operations.DeleteCrmConnectionIDPipelineIDRequest{
        ConnectionID: "laboriosam",
        ID: "73d522b8-28a9-4030-a60f-024c79b4cc64",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteCrmConnectionIDPipelineIDRequest](../../models/operations/deletecrmconnectionidpipelineidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.DeleteCrmConnectionIDPipelineIDSecurity](../../models/operations/deletecrmconnectionidpipelineidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDPipelineSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.GetCrmConnectionIDPipeline(ctx, operations.GetCrmConnectionIDPipelineRequest{
        ConnectionID: "eligendi",
        Limit: unifiedto.Float64(1687.36),
        Offset: unifiedto.Float64(7276.04),
        Order: unifiedto.String("sequi"),
        Query: unifiedto.String("culpa"),
        Sort: unifiedto.String("ratione"),
        UpdatedGte: types.MustTimeFromString("2022-03-30T02:59:59.063Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCrmConnectionIDPipelineRequest](../../models/operations/getcrmconnectionidpipelinerequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetCrmConnectionIDPipelineSecurity](../../models/operations/getcrmconnectionidpipelinesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDPipelineIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.GetCrmConnectionIDPipelineID(ctx, operations.GetCrmConnectionIDPipelineIDRequest{
        ConnectionID: "labore",
        ID: "88ade62f-6aa5-458a-a5e2-083016ca34bb",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetCrmConnectionIDPipelineIDRequest](../../models/operations/getcrmconnectionidpipelineidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.GetCrmConnectionIDPipelineIDSecurity](../../models/operations/getcrmconnectionidpipelineidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDPipelineIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.PatchCrmConnectionIDPipelineID(ctx, operations.PatchCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedto.Bool(false),
            CreatedAt: types.MustTimeFromString("2022-01-06T01:33:11.339Z"),
            DealProbability: unifiedto.Bool(false),
            DisplayOrder: unifiedto.Float64(8302.16),
            ID: unifiedto.String("4f62127a-607d-4160-a294-514c3db9ca9f"),
            Name: unifiedto.String("Brandy Powlowski"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2021-03-30T20:29:28.566Z"),
        },
        ConnectionID: "quos",
        ID: "78703493-f49a-4a84-a5a3-283279b719d1",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchCrmConnectionIDPipelineIDRequest](../../models/operations/patchcrmconnectionidpipelineidrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PatchCrmConnectionIDPipelineIDSecurity](../../models/operations/patchcrmconnectionidpipelineidsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDPipelineSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.PostCrmConnectionIDPipeline(ctx, operations.PostCrmConnectionIDPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedto.Bool(false),
            CreatedAt: types.MustTimeFromString("2020-05-07T22:58:48.615Z"),
            DealProbability: unifiedto.Bool(false),
            DisplayOrder: unifiedto.Float64(6422.68),
            ID: unifiedto.String("673d86e3-b35e-449a-b135-778ce54cacb0"),
            Name: unifiedto.String("Chris Terry"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2022-09-03T20:01:11.409Z"),
        },
        ConnectionID: "voluptatem",
    }, operationSecurity)
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
| `request`                                                                                                        | [operations.PostCrmConnectionIDPipelineRequest](../../models/operations/postcrmconnectionidpipelinerequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PostCrmConnectionIDPipelineSecurity](../../models/operations/postcrmconnectionidpipelinesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDPipelineIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Pipeline.PutCrmConnectionIDPipelineID(ctx, operations.PutCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedto.Bool(false),
            CreatedAt: types.MustTimeFromString("2022-08-22T17:28:32.263Z"),
            DealProbability: unifiedto.Bool(false),
            DisplayOrder: unifiedto.Float64(6880.36),
            ID: unifiedto.String("acf63b21-5186-4ab5-a3a0-22614315d156"),
            Name: unifiedto.String("Victor Mayer"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2022-11-25T09:18:50.894Z"),
        },
        ConnectionID: "officia",
        ID: "fc7186ff-20b7-4a73-9f40-ca0d7657c164",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutCrmConnectionIDPipelineIDRequest](../../models/operations/putcrmconnectionidpipelineidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PutCrmConnectionIDPipelineIDSecurity](../../models/operations/putcrmconnectionidpipelineidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PutCrmConnectionIDPipelineIDResponse](../../models/operations/putcrmconnectionidpipelineidresponse.md), error**

