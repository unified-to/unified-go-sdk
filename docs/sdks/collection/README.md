# Collection

## Overview

### Available Operations

* [CreateCommerceCollection](#createcommercecollection) - Create a collection
* [CreateLmsCollection](#createlmscollection) - Create a collection
* [GetCommerceCollection](#getcommercecollection) - Retrieve a collection
* [GetLmsCollection](#getlmscollection) - Retrieve a collection
* [ListCommerceCollections](#listcommercecollections) - List all collections
* [ListLmsCollections](#listlmscollections) - List all collections
* [PatchCommerceCollection](#patchcommercecollection) - Update a collection
* [PatchLmsCollection](#patchlmscollection) - Update a collection
* [RemoveCommerceCollection](#removecommercecollection) - Remove a collection
* [RemoveLmsCollection](#removelmscollection) - Remove a collection
* [UpdateCommerceCollection](#updatecommercecollection) - Update a collection
* [UpdateLmsCollection](#updatelmscollection) - Update a collection

## CreateCommerceCollection

Creates a CommerceCollection. Provide the unified CommerceCollection fields in the request body; they are mapped to the provider's API. Returns the created CommerceCollection.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceCollection" method="post" path="/commerce/{connection_id}/collection" example="commerce_collection" -->
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

    res, err := s.Collection.CreateCommerceCollection(ctx, operations.CreateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("Ward - Hudson's most advanced Computer technology increases graceful capabilities"),
            ID: unifiedgosdk.Pointer("7b6d55fb-230e-49c3-ba00-59f4b64c101b"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Armarium amiculum amet defendo."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("90dce3d5-c245-458a-927e-7e6bee8692b2"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("dec0b586-2675-45ba-a9e2-d41656cf4270"),
                            Slug: unifiedgosdk.Pointer("cernuus"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "torrens",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/62m4msPNw/3209/1074",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("ef588586-625d-409c-b054-471232c68496"),
                    Slug: unifiedgosdk.Pointer("velit"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "aureus",
                    )),
                },
            },
            Name: "Rustic Gold Mouse",
            PublicDescription: unifiedgosdk.Pointer("Introducing the Sint Maarten-inspired Gloves, blending awful style with local craftsmanship"),
            PublicName: unifiedgosdk.Pointer("Rustic Gold Mouse"),
            Tags: []string{
                "viridis",
                "deserunt",
                "comparo",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T10:59:01.696Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateCommerceCollectionRequest](../../pkg/models/operations/createcommercecollectionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateCommerceCollectionResponse](../../pkg/models/operations/createcommercecollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsCollection

Creates a LmsCollection. Provide the unified LmsCollection fields in the request body; they are mapped to the provider's API. Returns the created LmsCollection.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsCollection" method="post" path="/lms/{connection_id}/collection" example="lms_collection" -->
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

    res, err := s.Collection.CreateLmsCollection(ctx, operations.CreateLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("439785fe-6475-4a32-9d13-d51268ed7f41"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.029Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateLmsCollectionRequest](../../pkg/models/operations/createlmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateLmsCollectionResponse](../../pkg/models/operations/createlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceCollection

Retrieves a single CommerceCollection by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceCollection" method="get" path="/commerce/{connection_id}/collection/{id}" -->
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

    res, err := s.Collection.GetCommerceCollection(ctx, operations.GetCommerceCollectionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCommerceCollectionRequest](../../pkg/models/operations/getcommercecollectionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetCommerceCollectionResponse](../../pkg/models/operations/getcommercecollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsCollection

Retrieves a single LmsCollection by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsCollection" method="get" path="/lms/{connection_id}/collection/{id}" -->
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

    res, err := s.Collection.GetLmsCollection(ctx, operations.GetLmsCollectionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetLmsCollectionRequest](../../pkg/models/operations/getlmscollectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetLmsCollectionResponse](../../pkg/models/operations/getlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceCollections

Returns a list of CommerceCollections. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceCollections" method="get" path="/commerce/{connection_id}/collection" -->
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

    res, err := s.Collection.ListCommerceCollections(ctx, operations.ListCommerceCollectionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceCollections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListCommerceCollectionsRequest](../../pkg/models/operations/listcommercecollectionsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListCommerceCollectionsResponse](../../pkg/models/operations/listcommercecollectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsCollections

Returns a list of LmsCollections. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsCollections" method="get" path="/lms/{connection_id}/collection" -->
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

    res, err := s.Collection.ListLmsCollections(ctx, operations.ListLmsCollectionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLmsCollectionsRequest](../../pkg/models/operations/listlmscollectionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListLmsCollectionsResponse](../../pkg/models/operations/listlmscollectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceCollection

Updates the CommerceCollection identified by `id`. Provide the unified CommerceCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceCollection" method="patch" path="/commerce/{connection_id}/collection/{id}" example="commerce_collection" -->
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

    res, err := s.Collection.PatchCommerceCollection(ctx, operations.PatchCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("Ward - Hudson's most advanced Computer technology increases graceful capabilities"),
            ID: unifiedgosdk.Pointer("fb5292ee-8caa-40c3-9046-bbf0609a17f9"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Armarium amiculum amet defendo."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("52c62c32-90ba-4fee-b995-868f116cdbe5"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("14591a24-0310-46c6-bb5b-6b6cd21b5efd"),
                            Slug: unifiedgosdk.Pointer("cernuus"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "torrens",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/62m4msPNw/3209/1074",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("df094aa0-048a-4e25-b625-895340b670b8"),
                    Slug: unifiedgosdk.Pointer("velit"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "aureus",
                    )),
                },
            },
            Name: "Rustic Gold Mouse",
            PublicDescription: unifiedgosdk.Pointer("Introducing the Sint Maarten-inspired Gloves, blending awful style with local craftsmanship"),
            PublicName: unifiedgosdk.Pointer("Rustic Gold Mouse"),
            Tags: []string{
                "viridis",
                "deserunt",
                "comparo",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T10:59:01.704Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchCommerceCollectionRequest](../../pkg/models/operations/patchcommercecollectionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PatchCommerceCollectionResponse](../../pkg/models/operations/patchcommercecollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsCollection

Updates the LmsCollection identified by `id`. Provide the unified LmsCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsCollection" method="patch" path="/lms/{connection_id}/collection/{id}" example="lms_collection" -->
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

    res, err := s.Collection.PatchLmsCollection(ctx, operations.PatchLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("b3ba5ccf-3f87-4ae8-86a4-6adb7d2b652a"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.038Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchLmsCollectionRequest](../../pkg/models/operations/patchlmscollectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchLmsCollectionResponse](../../pkg/models/operations/patchlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceCollection

Deletes the CommerceCollection identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceCollection" method="delete" path="/commerce/{connection_id}/collection/{id}" -->
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

    res, err := s.Collection.RemoveCommerceCollection(ctx, operations.RemoveCommerceCollectionRequest{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.RemoveCommerceCollectionRequest](../../pkg/models/operations/removecommercecollectionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.RemoveCommerceCollectionResponse](../../pkg/models/operations/removecommercecollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsCollection

Deletes the LmsCollection identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsCollection" method="delete" path="/lms/{connection_id}/collection/{id}" -->
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

    res, err := s.Collection.RemoveLmsCollection(ctx, operations.RemoveLmsCollectionRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveLmsCollectionRequest](../../pkg/models/operations/removelmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveLmsCollectionResponse](../../pkg/models/operations/removelmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceCollection

Updates the CommerceCollection identified by `id`. Provide the unified CommerceCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceCollection" method="put" path="/commerce/{connection_id}/collection/{id}" example="commerce_collection" -->
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

    res, err := s.Collection.UpdateCommerceCollection(ctx, operations.UpdateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("Ward - Hudson's most advanced Computer technology increases graceful capabilities"),
            ID: unifiedgosdk.Pointer("fb5292ee-8caa-40c3-9046-bbf0609a17f9"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Armarium amiculum amet defendo."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("52c62c32-90ba-4fee-b995-868f116cdbe5"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("14591a24-0310-46c6-bb5b-6b6cd21b5efd"),
                            Slug: unifiedgosdk.Pointer("cernuus"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "torrens",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/62m4msPNw/3209/1074",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("df094aa0-048a-4e25-b625-895340b670b8"),
                    Slug: unifiedgosdk.Pointer("velit"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "aureus",
                    )),
                },
            },
            Name: "Rustic Gold Mouse",
            PublicDescription: unifiedgosdk.Pointer("Introducing the Sint Maarten-inspired Gloves, blending awful style with local craftsmanship"),
            PublicName: unifiedgosdk.Pointer("Rustic Gold Mouse"),
            Tags: []string{
                "viridis",
                "deserunt",
                "comparo",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-03T10:59:01.704Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateCommerceCollectionRequest](../../pkg/models/operations/updatecommercecollectionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateCommerceCollectionResponse](../../pkg/models/operations/updatecommercecollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsCollection

Updates the LmsCollection identified by `id`. Provide the unified LmsCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsCollection" method="put" path="/lms/{connection_id}/collection/{id}" example="lms_collection" -->
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

    res, err := s.Collection.UpdateLmsCollection(ctx, operations.UpdateLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("b3ba5ccf-3f87-4ae8-86a4-6adb7d2b652a"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.038Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateLmsCollectionRequest](../../pkg/models/operations/updatelmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateLmsCollectionResponse](../../pkg/models/operations/updatelmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |