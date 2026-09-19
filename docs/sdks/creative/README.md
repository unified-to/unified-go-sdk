# Creative

## Overview

### Available Operations

* [CreateAdsCreative](#createadscreative) - Create a creative
* [GetAdsCreative](#getadscreative) - Retrieve a creative
* [ListAdsCreatives](#listadscreatives) - List all creatives
* [PatchAdsCreative](#patchadscreative) - Update a creative
* [RemoveAdsCreative](#removeadscreative) - Remove a creative
* [UpdateAdsCreative](#updateadscreative) - Update a creative

## CreateAdsCreative

Create a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsCreative" method="post" path="/ads/{connection_id}/creative" example="ads_creative" -->
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

    res, err := s.Creative.CreateAdsCreative(ctx, operations.CreateAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("f33457bb-b7b6-44d4-a15f-72ac0a2b61c6"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T05:34:35.455Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAdsCreativeRequest](../../pkg/models/operations/createadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAdsCreativeResponse](../../pkg/models/operations/createadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsCreative

Retrieve a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsCreative" method="get" path="/ads/{connection_id}/creative/{id}" -->
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

    res, err := s.Creative.GetAdsCreative(ctx, operations.GetAdsCreativeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAdsCreativeRequest](../../pkg/models/operations/getadscreativerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAdsCreativeResponse](../../pkg/models/operations/getadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsCreatives

List all creatives

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsCreatives" method="get" path="/ads/{connection_id}/creative" -->
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

    res, err := s.Creative.ListAdsCreatives(ctx, operations.ListAdsCreativesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreatives != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAdsCreativesRequest](../../pkg/models/operations/listadscreativesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAdsCreativesResponse](../../pkg/models/operations/listadscreativesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsCreative

Update a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsCreative" method="patch" path="/ads/{connection_id}/creative/{id}" example="ads_creative" -->
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

    res, err := s.Creative.PatchAdsCreative(ctx, operations.PatchAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("9a027bbe-7d9e-4589-8d2c-3189a4d3a369"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T05:34:35.462Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAdsCreativeRequest](../../pkg/models/operations/patchadscreativerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAdsCreativeResponse](../../pkg/models/operations/patchadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsCreative

Remove a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsCreative" method="delete" path="/ads/{connection_id}/creative/{id}" -->
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

    res, err := s.Creative.RemoveAdsCreative(ctx, operations.RemoveAdsCreativeRequest{
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
| `request`                                                                                      | [operations.RemoveAdsCreativeRequest](../../pkg/models/operations/removeadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAdsCreativeResponse](../../pkg/models/operations/removeadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsCreative

Update a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsCreative" method="put" path="/ads/{connection_id}/creative/{id}" example="ads_creative" -->
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

    res, err := s.Creative.UpdateAdsCreative(ctx, operations.UpdateAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("9a027bbe-7d9e-4589-8d2c-3189a4d3a369"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T05:34:35.462Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAdsCreativeRequest](../../pkg/models/operations/updateadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAdsCreativeResponse](../../pkg/models/operations/updateadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |