# Ad

## Overview

### Available Operations

* [CreateAdsAd](#createadsad) - Create an ad
* [GetAdsAd](#getadsad) - Retrieve an ad
* [ListAdsAds](#listadsads) - List all ads
* [PatchAdsAd](#patchadsad) - Update an ad
* [RemoveAdsAd](#removeadsad) - Remove an ad
* [UpdateAdsAd](#updateadsad) - Update an ad

## CreateAdsAd

Creates a AdsAd. Provide the unified AdsAd fields in the request body; they are mapped to the provider's API. Returns the created AdsAd.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsAd" method="post" path="/ads/{connection_id}/ad" example="ads_ad" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ad.CreateAdsAd(ctx, operations.CreateAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("cb34f714-8f62-403c-a382-00c242559806"),
            Metadata: []shared.AdsMetadata{},
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-19T17:40:38.164Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateAdsAdRequest](../../pkg/models/operations/createadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateAdsAdResponse](../../pkg/models/operations/createadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsAd

Retrieves a single AdsAd by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsAd" method="get" path="/ads/{connection_id}/ad/{id}" -->
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

    res, err := s.Ad.GetAdsAd(ctx, operations.GetAdsAdRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAdsAdRequest](../../pkg/models/operations/getadsadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetAdsAdResponse](../../pkg/models/operations/getadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsAds

Returns a list of AdsAds. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsAds" method="get" path="/ads/{connection_id}/ad" -->
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

    res, err := s.Ad.ListAdsAds(ctx, operations.ListAdsAdsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAdsAdsRequest](../../pkg/models/operations/listadsadsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAdsAdsResponse](../../pkg/models/operations/listadsadsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsAd

Updates the AdsAd identified by `id`. Provide the unified AdsAd fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsAd" method="patch" path="/ads/{connection_id}/ad/{id}" example="ads_ad" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ad.PatchAdsAd(ctx, operations.PatchAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("7d5ce116-3ff5-4d88-8ee6-f4894faf7853"),
            Metadata: []shared.AdsMetadata{},
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-19T17:40:38.180Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchAdsAdRequest](../../pkg/models/operations/patchadsadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PatchAdsAdResponse](../../pkg/models/operations/patchadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsAd

Deletes the AdsAd identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsAd" method="delete" path="/ads/{connection_id}/ad/{id}" -->
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

    res, err := s.Ad.RemoveAdsAd(ctx, operations.RemoveAdsAdRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveAdsAdRequest](../../pkg/models/operations/removeadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.RemoveAdsAdResponse](../../pkg/models/operations/removeadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsAd

Updates the AdsAd identified by `id`. Provide the unified AdsAd fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsAd" method="put" path="/ads/{connection_id}/ad/{id}" example="ads_ad" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ad.UpdateAdsAd(ctx, operations.UpdateAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("7d5ce116-3ff5-4d88-8ee6-f4894faf7853"),
            Metadata: []shared.AdsMetadata{},
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-19T17:40:38.180Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateAdsAdRequest](../../pkg/models/operations/updateadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateAdsAdResponse](../../pkg/models/operations/updateadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |