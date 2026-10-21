# Itemvariant

## Overview

### Available Operations

* [CreateCommerceItemvariant](#createcommerceitemvariant) - Create an itemvariant
* [GetCommerceItemvariant](#getcommerceitemvariant) - Retrieve an itemvariant
* [ListCommerceItemvariants](#listcommerceitemvariants) - List all itemvariants
* [PatchCommerceItemvariant](#patchcommerceitemvariant) - Update an itemvariant
* [RemoveCommerceItemvariant](#removecommerceitemvariant) - Remove an itemvariant
* [UpdateCommerceItemvariant](#updatecommerceitemvariant) - Update an itemvariant

## CreateCommerceItemvariant

Create an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceItemvariant" method="post" path="/commerce/{connection_id}/itemvariant" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.CreateCommerceItemvariant(ctx, operations.CreateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("21db4bbc-5da8-4b7e-8542-9513b516847a"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("ca51d48c-b6b2-4f76-b640-30f89cdc34cb"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("d1e9b8f7-eeff-45d5-9e9e-027461c63701"),
                            Slug: unifiedgosdk.Pointer("ubi"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "textus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2126/721?lock=8820088623807162",
                    Width: unifiedgosdk.Pointer[float64](35.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f1dc1de2-ad01-4062-ad1d-26fcbdb0f39a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "absens",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Chair"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("de487d6a-518f-4143-832a-f0e5d623f17d"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("2d0a0790-a7e6-430c-aa3e-d5168c991f3c"),
                    Name: "Size",
                    Position: unifiedgosdk.Pointer[float64](2.0),
                    Values: []string{
                        "S",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 382.23,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Cheese designed to make you stand out with grave looks"),
            PublicName: unifiedgosdk.Pointer("Chair"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-368-44270-1"),
            Tags: []string{
                "deserunt",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-03-09T16:20:44.119Z"),
            Weight: unifiedgosdk.Pointer[float64](46.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](40.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateCommerceItemvariantRequest](../../pkg/models/operations/createcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateCommerceItemvariantResponse](../../pkg/models/operations/createcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceItemvariant

Retrieve an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceItemvariant" method="get" path="/commerce/{connection_id}/itemvariant/{id}" -->
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

    res, err := s.Itemvariant.GetCommerceItemvariant(ctx, operations.GetCommerceItemvariantRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCommerceItemvariantRequest](../../pkg/models/operations/getcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetCommerceItemvariantResponse](../../pkg/models/operations/getcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceItemvariants

List all itemvariants

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceItemvariants" method="get" path="/commerce/{connection_id}/itemvariant" -->
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

    res, err := s.Itemvariant.ListCommerceItemvariants(ctx, operations.ListCommerceItemvariantsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariants != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCommerceItemvariantsRequest](../../pkg/models/operations/listcommerceitemvariantsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCommerceItemvariantsResponse](../../pkg/models/operations/listcommerceitemvariantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceItemvariant

Update an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceItemvariant" method="patch" path="/commerce/{connection_id}/itemvariant/{id}" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.PatchCommerceItemvariant(ctx, operations.PatchCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("e5b6c614-77df-4640-9851-1b017edac238"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("193d724e-664a-4ff3-9e5e-f06a3a6d6bce"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("be94f989-1203-4816-9c17-6ec4a183fb8c"),
                            Slug: unifiedgosdk.Pointer("ubi"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "textus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2126/721?lock=8820088623807162",
                    Width: unifiedgosdk.Pointer[float64](35.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("029b8520-9463-4356-9fca-a3b2058019a8"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "absens",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Chair"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("e1ec43ce-9696-49ad-9233-4941ec412803"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("9c026072-0592-40da-a00c-9556e7198eed"),
                    Name: "Size",
                    Position: unifiedgosdk.Pointer[float64](2.0),
                    Values: []string{
                        "S",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 382.23,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Cheese designed to make you stand out with grave looks"),
            PublicName: unifiedgosdk.Pointer("Chair"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-368-44270-1"),
            Tags: []string{
                "deserunt",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-03-09T16:20:44.143Z"),
            Weight: unifiedgosdk.Pointer[float64](46.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](40.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCommerceItemvariantRequest](../../pkg/models/operations/patchcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchCommerceItemvariantResponse](../../pkg/models/operations/patchcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceItemvariant

Remove an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceItemvariant" method="delete" path="/commerce/{connection_id}/itemvariant/{id}" -->
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

    res, err := s.Itemvariant.RemoveCommerceItemvariant(ctx, operations.RemoveCommerceItemvariantRequest{
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.RemoveCommerceItemvariantRequest](../../pkg/models/operations/removecommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemoveCommerceItemvariantResponse](../../pkg/models/operations/removecommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceItemvariant

Update an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceItemvariant" method="put" path="/commerce/{connection_id}/itemvariant/{id}" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.UpdateCommerceItemvariant(ctx, operations.UpdateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("e5b6c614-77df-4640-9851-1b017edac238"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("193d724e-664a-4ff3-9e5e-f06a3a6d6bce"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("be94f989-1203-4816-9c17-6ec4a183fb8c"),
                            Slug: unifiedgosdk.Pointer("ubi"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "textus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2126/721?lock=8820088623807162",
                    Width: unifiedgosdk.Pointer[float64](35.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("029b8520-9463-4356-9fca-a3b2058019a8"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "absens",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Chair"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("e1ec43ce-9696-49ad-9233-4941ec412803"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("9c026072-0592-40da-a00c-9556e7198eed"),
                    Name: "Size",
                    Position: unifiedgosdk.Pointer[float64](2.0),
                    Values: []string{
                        "S",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 382.23,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Cheese designed to make you stand out with grave looks"),
            PublicName: unifiedgosdk.Pointer("Chair"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-368-44270-1"),
            Tags: []string{
                "deserunt",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-03-09T16:20:44.143Z"),
            Weight: unifiedgosdk.Pointer[float64](46.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](40.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateCommerceItemvariantRequest](../../pkg/models/operations/updatecommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateCommerceItemvariantResponse](../../pkg/models/operations/updatecommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |