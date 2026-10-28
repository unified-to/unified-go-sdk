# Metadata

## Overview

### Available Operations

* [CreateMetadataMetadata](#createmetadatametadata) - Create a metadata
* [GetMetadataMetadata](#getmetadatametadata) - Retrieve a metadata
* [ListMetadataMetadatas](#listmetadatametadatas) - List all metadatas
* [PatchMetadataMetadata](#patchmetadatametadata) - Update a metadata
* [RemoveMetadataMetadata](#removemetadatametadata) - Remove a metadata
* [UpdateMetadataMetadata](#updatemetadatametadata) - Update a metadata

## CreateMetadataMetadata

Creates a MetadataMetadata. Provide the unified MetadataMetadata fields in the request body; they are mapped to the provider's API. Returns the created MetadataMetadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="createMetadataMetadata" method="post" path="/metadata/{connection_id}/metadata" example="metadata_metadata" -->
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

    res, err := s.Metadata.CreateMetadataMetadata(ctx, operations.CreateMetadataMetadataRequest{
        MetadataMetadata: shared.MetadataMetadata{
            CreatedAt: types.MustNewTimeFromString("2021-03-25T03:02:17.656Z"),
            Format: shared.MetadataMetadataFormatPrice.ToPointer(),
            ID: unifiedgosdk.Pointer("c6988a64-6bd1-47cd-837c-827a77b19c85"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "legal_expense",
            Objects: map[string]string{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-27T00:05:22.679Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateMetadataMetadataRequest](../../pkg/models/operations/createmetadatametadatarequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateMetadataMetadataResponse](../../pkg/models/operations/createmetadatametadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMetadataMetadata

Retrieves a single MetadataMetadata by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMetadataMetadata" method="get" path="/metadata/{connection_id}/metadata/{id}" -->
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

    res, err := s.Metadata.GetMetadataMetadata(ctx, operations.GetMetadataMetadataRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetMetadataMetadataRequest](../../pkg/models/operations/getmetadatametadatarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetMetadataMetadataResponse](../../pkg/models/operations/getmetadatametadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMetadataMetadatas

Returns a list of MetadataMetadatas. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMetadataMetadatas" method="get" path="/metadata/{connection_id}/metadata" -->
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

    res, err := s.Metadata.ListMetadataMetadatas(ctx, operations.ListMetadataMetadatasRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataMetadatas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListMetadataMetadatasRequest](../../pkg/models/operations/listmetadatametadatasrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListMetadataMetadatasResponse](../../pkg/models/operations/listmetadatametadatasresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMetadataMetadata

Updates the MetadataMetadata identified by `id`. Provide the unified MetadataMetadata fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMetadataMetadata" method="patch" path="/metadata/{connection_id}/metadata/{id}" example="metadata_metadata" -->
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

    res, err := s.Metadata.PatchMetadataMetadata(ctx, operations.PatchMetadataMetadataRequest{
        MetadataMetadata: shared.MetadataMetadata{
            CreatedAt: types.MustNewTimeFromString("2021-03-25T03:02:17.656Z"),
            Format: shared.MetadataMetadataFormatPrice.ToPointer(),
            ID: unifiedgosdk.Pointer("f0241ea1-9e4b-459e-946c-e09bbc7623d1"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "legal_expense",
            Objects: map[string]string{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-27T00:05:22.685Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchMetadataMetadataRequest](../../pkg/models/operations/patchmetadatametadatarequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchMetadataMetadataResponse](../../pkg/models/operations/patchmetadatametadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMetadataMetadata

Deletes the MetadataMetadata identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMetadataMetadata" method="delete" path="/metadata/{connection_id}/metadata/{id}" -->
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

    res, err := s.Metadata.RemoveMetadataMetadata(ctx, operations.RemoveMetadataMetadataRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveMetadataMetadataRequest](../../pkg/models/operations/removemetadatametadatarequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveMetadataMetadataResponse](../../pkg/models/operations/removemetadatametadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMetadataMetadata

Updates the MetadataMetadata identified by `id`. Provide the unified MetadataMetadata fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMetadataMetadata" method="put" path="/metadata/{connection_id}/metadata/{id}" example="metadata_metadata" -->
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

    res, err := s.Metadata.UpdateMetadataMetadata(ctx, operations.UpdateMetadataMetadataRequest{
        MetadataMetadata: shared.MetadataMetadata{
            CreatedAt: types.MustNewTimeFromString("2021-03-25T03:02:17.656Z"),
            Format: shared.MetadataMetadataFormatPrice.ToPointer(),
            ID: unifiedgosdk.Pointer("f0241ea1-9e4b-459e-946c-e09bbc7623d1"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "legal_expense",
            Objects: map[string]string{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-27T00:05:22.685Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateMetadataMetadataRequest](../../pkg/models/operations/updatemetadatametadatarequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateMetadataMetadataResponse](../../pkg/models/operations/updatemetadatametadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |