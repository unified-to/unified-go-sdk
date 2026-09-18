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

Create a metadata

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
            ID: unifiedgosdk.Pointer("da2aee51-87b4-4205-965a-dc91a9c73dd9"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "clubs_group",
            Objects: map[string]any{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-26T08:43:36.161Z"),
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

Retrieve a metadata

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

List all metadatas

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

Update a metadata

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
            ID: unifiedgosdk.Pointer("d4a545d2-0719-493e-a936-8b46198b371c"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "clubs_group",
            Objects: map[string]any{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-26T08:43:36.165Z"),
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

Remove a metadata

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

Update a metadata

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
            ID: unifiedgosdk.Pointer("d4a545d2-0719-493e-a936-8b46198b371c"),
            IsRequired: unifiedgosdk.Pointer(false),
            Name: "autem",
            ObjectType: "clubs_group",
            Objects: map[string]any{

            },
            Options: []string{},
            OriginalFormat: unifiedgosdk.Pointer("advoco"),
            Slug: unifiedgosdk.Pointer("arbustum"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-26T08:43:36.165Z"),
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