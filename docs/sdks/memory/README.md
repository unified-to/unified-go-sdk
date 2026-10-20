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

Create a memory

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
            Content: unifiedgosdk.Pointer("Tabella rem approbo spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis. Ullus ager degero deporto cicuta vulgivagus."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("787f8ed8-aa2d-41b9-a58b-8455734b1f44"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7a8ad62d-ba0e-4ce3-8330-56a2647b5b3e"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("cribro"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "aiunt",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("sponte.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-04-25T14:10:34.096Z"),
            UserIdentifier: unifiedgosdk.Pointer("9d170613-4338-447d-bdb6-8ea96a2d7513"),
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

Retrieve a memory

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

List all memories

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

Update a memory

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
            Content: unifiedgosdk.Pointer("Tabella rem approbo spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis. Ullus ager degero deporto cicuta vulgivagus."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("06002283-556d-4c53-a7aa-f69a1fae53ac"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7a8ad62d-ba0e-4ce3-8330-56a2647b5b3e"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("cribro"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "aiunt",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("sponte.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-04-25T14:10:34.101Z"),
            UserIdentifier: unifiedgosdk.Pointer("9d170613-4338-447d-bdb6-8ea96a2d7513"),
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

Remove a memory

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

Update a memory

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
            Content: unifiedgosdk.Pointer("Tabella rem approbo spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis. Ullus ager degero deporto cicuta vulgivagus."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("06002283-556d-4c53-a7aa-f69a1fae53ac"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7a8ad62d-ba0e-4ce3-8330-56a2647b5b3e"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("cribro"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "aiunt",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("sponte.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-04-25T14:10:34.101Z"),
            UserIdentifier: unifiedgosdk.Pointer("9d170613-4338-447d-bdb6-8ea96a2d7513"),
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