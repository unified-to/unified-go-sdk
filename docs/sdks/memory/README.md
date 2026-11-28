# Memory

## Overview

### Available Operations

* [CreateGenaiMemory](#creategenaimemory) - Create a memory
* [GetGenaiMemory](#getgenaimemory) - Retrieve a memory
* [ListGenaiMemories](#listgenaimemories) - List all memories
* [PatchGenaiMemory](#patchgenaimemory) - Update a memory
* [RemoveGenaiMemory](#removegenaimemory) - Remove a memory
* [UpdateGenaiMemory](#updategenaimemory) - Update a memory

## CreateGenaiMemory

Creates a GenaiMemory. Provide the unified GenaiMemory fields in the request body; they are mapped to the provider's API. Returns the created GenaiMemory.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiMemory" method="post" path="/genai/{connection_id}/memory" example="genai_memory" -->
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

    res, err := s.Memory.CreateGenaiMemory(ctx, operations.CreateGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("ff52a677-2ba9-44bc-acce-6051bde2356b"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-07T08:46:48.277Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateGenaiMemoryRequest](../../pkg/models/operations/creategenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateGenaiMemoryResponse](../../pkg/models/operations/creategenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiMemory

Retrieves a single GenaiMemory by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiMemory" method="get" path="/genai/{connection_id}/memory/{id}" -->
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

    res, err := s.Memory.GetGenaiMemory(ctx, operations.GetGenaiMemoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetGenaiMemoryRequest](../../pkg/models/operations/getgenaimemoryrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetGenaiMemoryResponse](../../pkg/models/operations/getgenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiMemories

Returns a list of GenaiMemories. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiMemories" method="get" path="/genai/{connection_id}/memory" -->
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

    res, err := s.Memory.ListGenaiMemories(ctx, operations.ListGenaiMemoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListGenaiMemoriesRequest](../../pkg/models/operations/listgenaimemoriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListGenaiMemoriesResponse](../../pkg/models/operations/listgenaimemoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiMemory

Updates the GenaiMemory identified by `id`. Provide the unified GenaiMemory fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiMemory" method="patch" path="/genai/{connection_id}/memory/{id}" example="genai_memory" -->
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

    res, err := s.Memory.PatchGenaiMemory(ctx, operations.PatchGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("cf18c008-c223-4285-8599-921134017396"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-07T08:46:48.281Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchGenaiMemoryRequest](../../pkg/models/operations/patchgenaimemoryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchGenaiMemoryResponse](../../pkg/models/operations/patchgenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiMemory

Deletes the GenaiMemory identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiMemory" method="delete" path="/genai/{connection_id}/memory/{id}" -->
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

    res, err := s.Memory.RemoveGenaiMemory(ctx, operations.RemoveGenaiMemoryRequest{
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
| `request`                                                                                      | [operations.RemoveGenaiMemoryRequest](../../pkg/models/operations/removegenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveGenaiMemoryResponse](../../pkg/models/operations/removegenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiMemory

Updates the GenaiMemory identified by `id`. Provide the unified GenaiMemory fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiMemory" method="put" path="/genai/{connection_id}/memory/{id}" example="genai_memory" -->
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

    res, err := s.Memory.UpdateGenaiMemory(ctx, operations.UpdateGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("cf18c008-c223-4285-8599-921134017396"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-07T08:46:48.281Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateGenaiMemoryRequest](../../pkg/models/operations/updategenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateGenaiMemoryResponse](../../pkg/models/operations/updategenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |