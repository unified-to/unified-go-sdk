# Template

## Overview

### Available Operations

* [CreateComputeTemplate](#createcomputetemplate) - Create a template
* [GetComputeTemplate](#getcomputetemplate) - Retrieve a template
* [GetSigningTemplate](#getsigningtemplate) - Retrieve a template
* [ListComputeTemplates](#listcomputetemplates) - List all templates
* [ListSigningTemplates](#listsigningtemplates) - List all templates
* [PatchComputeTemplate](#patchcomputetemplate) - Update a template
* [RemoveComputeTemplate](#removecomputetemplate) - Remove a template
* [UpdateComputeTemplate](#updatecomputetemplate) - Update a template

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

    res, err := s.Template.CreateComputeTemplate(ctx, operations.CreateComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("87214967-c388-454a-b15b-d7ea77a70778"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-23T20:27:16.056Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-12T18:38:51.398Z"),
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

    res, err := s.Template.GetComputeTemplate(ctx, operations.GetComputeTemplateRequest{
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

## GetSigningTemplate

Retrieve a template

### Example Usage

<!-- UsageSnippet language="go" operationID="getSigningTemplate" method="get" path="/signing/{connection_id}/template/{id}" -->
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

    res, err := s.Template.GetSigningTemplate(ctx, operations.GetSigningTemplateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetSigningTemplateRequest](../../pkg/models/operations/getsigningtemplaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetSigningTemplateResponse](../../pkg/models/operations/getsigningtemplateresponse.md), error**

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

    res, err := s.Template.ListComputeTemplates(ctx, operations.ListComputeTemplatesRequest{
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

## ListSigningTemplates

List all templates

### Example Usage

<!-- UsageSnippet language="go" operationID="listSigningTemplates" method="get" path="/signing/{connection_id}/template" -->
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

    res, err := s.Template.ListSigningTemplates(ctx, operations.ListSigningTemplatesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListSigningTemplatesRequest](../../pkg/models/operations/listsigningtemplatesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListSigningTemplatesResponse](../../pkg/models/operations/listsigningtemplatesresponse.md), error**

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

    res, err := s.Template.PatchComputeTemplate(ctx, operations.PatchComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("9cb36833-667b-4fd1-a64c-9bad7d4d2a50"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-23T20:27:16.077Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-12T18:38:51.411Z"),
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

    res, err := s.Template.RemoveComputeTemplate(ctx, operations.RemoveComputeTemplateRequest{
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

    res, err := s.Template.UpdateComputeTemplate(ctx, operations.UpdateComputeTemplateRequest{
        ComputeTemplate: shared.ComputeTemplate{
            CPUCount: unifiedgosdk.Pointer[float64](4.0),
            CreatedAt: types.MustNewTimeFromString("2023-02-23T03:28:59.864Z"),
            DiskMb: unifiedgosdk.Pointer[float64](10240.0),
            Dockerfile: unifiedgosdk.Pointer("FROM ubuntu:24.04\nRUN apt-get update"),
            GpuCount: unifiedgosdk.Pointer[float64](2.0),
            GpuModel: shared.ComputeTemplateGpuModelNvidiaGh200.ToPointer(),
            ID: unifiedgosdk.Pointer("9cb36833-667b-4fd1-a64c-9bad7d4d2a50"),
            Image: unifiedgosdk.Pointer("perfumed-molasses/cohibeo:5.2"),
            IsPublic: unifiedgosdk.Pointer(false),
            LastUsedAt: types.MustNewTimeFromString("2026-03-23T20:27:16.077Z"),
            MemoryMb: unifiedgosdk.Pointer[float64](8192.0),
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("multi-byte-card"),
            Ports: []float64{
                8080.0,
            },
            Status: shared.ComputeTemplateStatusOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-12T18:38:51.411Z"),
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