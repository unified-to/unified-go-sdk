# Item

## Overview

### Available Operations

* [CreateCommerceItem](#createcommerceitem) - Create an item
* [GetCommerceItem](#getcommerceitem) - Retrieve an item
* [ListCommerceItems](#listcommerceitems) - List all items
* [PatchCommerceItem](#patchcommerceitem) - Update an item
* [RemoveCommerceItem](#removecommerceitem) - Remove an item
* [UpdateCommerceItem](#updatecommerceitem) - Update an item

## CreateCommerceItem

Creates a CommerceItem. Provide the unified CommerceItem fields in the request body; they are mapped to the provider's API. Returns the created CommerceItem.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceItem" method="post" path="/commerce/{connection_id}/item" example="commerce_item" -->
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

    res, err := s.Item.CreateCommerceItem(ctx, operations.CreateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("f9f0b334-236d-4049-be22-8bee59c597eb"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("32f9cfd0-81b0-4915-812e-0292b74ed36d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("31c0fc85-1c62-4b0c-aace-e76f2308366b"),
                            Slug: unifiedgosdk.Pointer("doloremque"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "allatus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](21.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/v73y0uzyK/3001/972",
                    Width: unifiedgosdk.Pointer[float64](67.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Comedo."),
                    Height: unifiedgosdk.Pointer[float64](651.0),
                    ID: unifiedgosdk.Pointer("41f4423a-2421-4a44-9e92-5b7669ff354e"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("dc8378ed-b157-4b42-9722-67b1a0424486"),
                            Slug: unifiedgosdk.Pointer("bis"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "somniculosus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](51.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/Ao4iatfO/771/3906",
                    Width: unifiedgosdk.Pointer[float64](3.0),
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
                    ID: unifiedgosdk.Pointer("805f9f5c-4633-488f-869a-d56125e6aa5a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "abstergo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 78.11,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Custodia ventus solio compono."),
            PublicName: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            RequiresShipping: unifiedgosdk.Pointer(true),
            Slug: unifiedgosdk.Pointer("handcrafted-rubber-tuna"),
            Tags: []string{
                "beatae",
                "blanditiis",
            },
            TotalStock: unifiedgosdk.Pointer[float64](682.0),
            Type: unifiedgosdk.Pointer("SERVICE"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-24T21:50:05.889Z"),
            VendorName: unifiedgosdk.Pointer("Mayer - Flatley"),
            Weight: unifiedgosdk.Pointer[float64](37.0),
            WeightUnit: shared.WeightUnitOz.ToPointer(),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCommerceItemRequest](../../pkg/models/operations/createcommerceitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateCommerceItemResponse](../../pkg/models/operations/createcommerceitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceItem

Retrieves a single CommerceItem by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceItem" method="get" path="/commerce/{connection_id}/item/{id}" -->
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

    res, err := s.Item.GetCommerceItem(ctx, operations.GetCommerceItemRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCommerceItemRequest](../../pkg/models/operations/getcommerceitemrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetCommerceItemResponse](../../pkg/models/operations/getcommerceitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceItems

Returns a list of CommerceItems. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceItems" method="get" path="/commerce/{connection_id}/item" -->
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

    res, err := s.Item.ListCommerceItems(ctx, operations.ListCommerceItemsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItems != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCommerceItemsRequest](../../pkg/models/operations/listcommerceitemsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCommerceItemsResponse](../../pkg/models/operations/listcommerceitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceItem

Updates the CommerceItem identified by `id`. Provide the unified CommerceItem fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceItem" method="patch" path="/commerce/{connection_id}/item/{id}" example="commerce_item" -->
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

    res, err := s.Item.PatchCommerceItem(ctx, operations.PatchCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("a5fdb470-24be-4417-8221-9887546eb7b0"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("d4bab427-6ee5-4956-9a32-c281ec2146fd"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("5c9a251f-d7ff-4e84-9936-7787147909e6"),
                            Slug: unifiedgosdk.Pointer("doloremque"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "allatus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](21.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/v73y0uzyK/3001/972",
                    Width: unifiedgosdk.Pointer[float64](67.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Comedo."),
                    Height: unifiedgosdk.Pointer[float64](651.0),
                    ID: unifiedgosdk.Pointer("f8270a1e-5c85-40c8-8818-bfd06498fe96"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("301bca97-ecad-4a01-82e8-53359186fae3"),
                            Slug: unifiedgosdk.Pointer("bis"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "somniculosus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](51.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/Ao4iatfO/771/3906",
                    Width: unifiedgosdk.Pointer[float64](3.0),
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
                    ID: unifiedgosdk.Pointer("a23d1cb5-9e56-4f71-9cef-08adaf36a0cb"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "abstergo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 78.11,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Custodia ventus solio compono."),
            PublicName: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            RequiresShipping: unifiedgosdk.Pointer(true),
            Slug: unifiedgosdk.Pointer("handcrafted-rubber-tuna"),
            Tags: []string{
                "beatae",
                "blanditiis",
            },
            TotalStock: unifiedgosdk.Pointer[float64](682.0),
            Type: unifiedgosdk.Pointer("SERVICE"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-24T21:50:05.902Z"),
            VendorName: unifiedgosdk.Pointer("Mayer - Flatley"),
            Weight: unifiedgosdk.Pointer[float64](37.0),
            WeightUnit: shared.WeightUnitOz.ToPointer(),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchCommerceItemRequest](../../pkg/models/operations/patchcommerceitemrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchCommerceItemResponse](../../pkg/models/operations/patchcommerceitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceItem

Deletes the CommerceItem identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceItem" method="delete" path="/commerce/{connection_id}/item/{id}" -->
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

    res, err := s.Item.RemoveCommerceItem(ctx, operations.RemoveCommerceItemRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveCommerceItemRequest](../../pkg/models/operations/removecommerceitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveCommerceItemResponse](../../pkg/models/operations/removecommerceitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceItem

Updates the CommerceItem identified by `id`. Provide the unified CommerceItem fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceItem" method="put" path="/commerce/{connection_id}/item/{id}" example="commerce_item" -->
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

    res, err := s.Item.UpdateCommerceItem(ctx, operations.UpdateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("a5fdb470-24be-4417-8221-9887546eb7b0"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("d4bab427-6ee5-4956-9a32-c281ec2146fd"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("5c9a251f-d7ff-4e84-9936-7787147909e6"),
                            Slug: unifiedgosdk.Pointer("doloremque"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "allatus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](21.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/v73y0uzyK/3001/972",
                    Width: unifiedgosdk.Pointer[float64](67.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Comedo."),
                    Height: unifiedgosdk.Pointer[float64](651.0),
                    ID: unifiedgosdk.Pointer("f8270a1e-5c85-40c8-8818-bfd06498fe96"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("301bca97-ecad-4a01-82e8-53359186fae3"),
                            Slug: unifiedgosdk.Pointer("bis"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "somniculosus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](51.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/Ao4iatfO/771/3906",
                    Width: unifiedgosdk.Pointer[float64](3.0),
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
                    ID: unifiedgosdk.Pointer("a23d1cb5-9e56-4f71-9cef-08adaf36a0cb"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "abstergo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    Currency: unifiedgosdk.Pointer("USD"),
                    Price: 78.11,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Custodia ventus solio compono."),
            PublicName: unifiedgosdk.Pointer("Handcrafted Rubber Tuna"),
            RequiresShipping: unifiedgosdk.Pointer(true),
            Slug: unifiedgosdk.Pointer("handcrafted-rubber-tuna"),
            Tags: []string{
                "beatae",
                "blanditiis",
            },
            TotalStock: unifiedgosdk.Pointer[float64](682.0),
            Type: unifiedgosdk.Pointer("SERVICE"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-24T21:50:05.902Z"),
            VendorName: unifiedgosdk.Pointer("Mayer - Flatley"),
            Weight: unifiedgosdk.Pointer[float64](37.0),
            WeightUnit: shared.WeightUnitOz.ToPointer(),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateCommerceItemRequest](../../pkg/models/operations/updatecommerceitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateCommerceItemResponse](../../pkg/models/operations/updatecommerceitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |