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

Creates a CommerceItemvariant. Provide the unified CommerceItemvariant fields in the request body; they are mapped to the provider's API. Returns the created CommerceItemvariant.

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
            Description: unifiedgosdk.Pointer("Our dog-friendly Bacon ensures pessimistic comfort for your pets"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("08935593-8768-40a3-aa1c-02960548704f"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("4d366b77-5e20-4305-853d-8ffd9f3fb0ce"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a62310ca-6e33-4478-a777-568735070a5f"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/3208/3394",
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
                    ID: unifiedgosdk.Pointer("dddfe6d7-8659-4fe3-b081-84d59cd6132f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "colligo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Table"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("528cac4f-1fb3-4b94-b491-acde585dbd92"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "purple",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("cc4af086-6ab2-4ed8-b9d0-184ff70fefbe"),
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
            PublicDescription: unifiedgosdk.Pointer("New teal Gloves with ergonomic design for favorite comfort"),
            PublicName: unifiedgosdk.Pointer("Table"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-05-109555-9"),
            Tags: []string{
                "labore",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-04-11T18:02:54.307Z"),
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

Retrieves a single CommerceItemvariant by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of CommerceItemvariants. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the CommerceItemvariant identified by `id`. Provide the unified CommerceItemvariant fields to change in the request body.

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
            Description: unifiedgosdk.Pointer("Our dog-friendly Bacon ensures pessimistic comfort for your pets"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("c02317bd-4ad9-4ca1-a116-0087da141996"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("3f5ba19e-a44b-4348-bd6d-b96ccd287a07"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("bb50eab0-4427-428b-a16e-ee03874088bd"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/3208/3394",
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
                    ID: unifiedgosdk.Pointer("ba903446-bd33-4129-9319-6621aa4d8fd9"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "colligo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Table"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("36854b1f-2993-4e60-a66c-e0e937f9bfdd"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "purple",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("9129e291-b79c-4b1e-aaca-31b8c59a4009"),
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
            PublicDescription: unifiedgosdk.Pointer("New teal Gloves with ergonomic design for favorite comfort"),
            PublicName: unifiedgosdk.Pointer("Table"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-05-109555-9"),
            Tags: []string{
                "labore",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-04-11T18:02:54.332Z"),
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

Deletes the CommerceItemvariant identified by `id`.

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

Updates the CommerceItemvariant identified by `id`. Provide the unified CommerceItemvariant fields to change in the request body.

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
            Description: unifiedgosdk.Pointer("Our dog-friendly Bacon ensures pessimistic comfort for your pets"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("c02317bd-4ad9-4ca1-a116-0087da141996"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("3f5ba19e-a44b-4348-bd6d-b96ccd287a07"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("bb50eab0-4427-428b-a16e-ee03874088bd"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/3208/3394",
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
                    ID: unifiedgosdk.Pointer("ba903446-bd33-4129-9319-6621aa4d8fd9"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "colligo",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Table"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("36854b1f-2993-4e60-a66c-e0e937f9bfdd"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "purple",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("9129e291-b79c-4b1e-aaca-31b8c59a4009"),
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
            PublicDescription: unifiedgosdk.Pointer("New teal Gloves with ergonomic design for favorite comfort"),
            PublicName: unifiedgosdk.Pointer("Table"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-1-05-109555-9"),
            Tags: []string{
                "labore",
            },
            TotalStock: unifiedgosdk.Pointer[float64](436.0),
            UpdatedAt: types.MustNewTimeFromString("2026-04-11T18:02:54.332Z"),
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