# Commerce

## Overview

### Available Operations

* [CreateCommerceCollection](#createcommercecollection) - Create a collection
* [CreateCommerceInventory](#createcommerceinventory) - Create an inventory
* [CreateCommerceItem](#createcommerceitem) - Create an item
* [CreateCommerceItemvariant](#createcommerceitemvariant) - Create an itemvariant
* [CreateCommerceLocation](#createcommercelocation) - Create a location
* [CreateCommerceReservation](#createcommercereservation) - Create a reservation
* [CreateCommerceReview](#createcommercereview) - Create a review
* [CreateCommerceSaleschannel](#createcommercesaleschannel) - Create a saleschannel
* [GetCommerceCollection](#getcommercecollection) - Retrieve a collection
* [GetCommerceInventory](#getcommerceinventory) - Retrieve an inventory
* [GetCommerceItem](#getcommerceitem) - Retrieve an item
* [GetCommerceItemvariant](#getcommerceitemvariant) - Retrieve an itemvariant
* [GetCommerceLocation](#getcommercelocation) - Retrieve a location
* [GetCommerceReservation](#getcommercereservation) - Retrieve a reservation
* [GetCommerceReview](#getcommercereview) - Retrieve a review
* [GetCommerceSaleschannel](#getcommercesaleschannel) - Retrieve a saleschannel
* [ListCommerceAvailabilities](#listcommerceavailabilities) - List all availabilities
* [ListCommerceCollections](#listcommercecollections) - List all collections
* [ListCommerceInventories](#listcommerceinventories) - List all inventories
* [ListCommerceItems](#listcommerceitems) - List all items
* [ListCommerceItemvariants](#listcommerceitemvariants) - List all itemvariants
* [ListCommerceLocations](#listcommercelocations) - List all locations
* [ListCommerceReservations](#listcommercereservations) - List all reservations
* [ListCommerceReviews](#listcommercereviews) - List all reviews
* [ListCommerceSaleschannels](#listcommercesaleschannels) - List all saleschannels
* [PatchCommerceCollection](#patchcommercecollection) - Update a collection
* [PatchCommerceInventory](#patchcommerceinventory) - Update an inventory
* [PatchCommerceItem](#patchcommerceitem) - Update an item
* [PatchCommerceItemvariant](#patchcommerceitemvariant) - Update an itemvariant
* [PatchCommerceLocation](#patchcommercelocation) - Update a location
* [PatchCommerceReservation](#patchcommercereservation) - Update a reservation
* [PatchCommerceReview](#patchcommercereview) - Update a review
* [PatchCommerceSaleschannel](#patchcommercesaleschannel) - Update a saleschannel
* [RemoveCommerceCollection](#removecommercecollection) - Remove a collection
* [RemoveCommerceInventory](#removecommerceinventory) - Remove an inventory
* [RemoveCommerceItem](#removecommerceitem) - Remove an item
* [RemoveCommerceItemvariant](#removecommerceitemvariant) - Remove an itemvariant
* [RemoveCommerceLocation](#removecommercelocation) - Remove a location
* [RemoveCommerceReservation](#removecommercereservation) - Remove a reservation
* [RemoveCommerceReview](#removecommercereview) - Remove a review
* [RemoveCommerceSaleschannel](#removecommercesaleschannel) - Remove a saleschannel
* [UpdateCommerceCollection](#updatecommercecollection) - Update a collection
* [UpdateCommerceInventory](#updatecommerceinventory) - Update an inventory
* [UpdateCommerceItem](#updatecommerceitem) - Update an item
* [UpdateCommerceItemvariant](#updatecommerceitemvariant) - Update an itemvariant
* [UpdateCommerceLocation](#updatecommercelocation) - Update a location
* [UpdateCommerceReservation](#updatecommercereservation) - Update a reservation
* [UpdateCommerceReview](#updatecommercereview) - Update a review
* [UpdateCommerceSaleschannel](#updatecommercesaleschannel) - Update a saleschannel

## CreateCommerceCollection

Create a collection

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

    res, err := s.Commerce.CreateCommerceCollection(ctx, operations.CreateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("1991e4a2-1f34-47c0-8486-bcee77e5935d"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("d299b709-a77d-4d3b-b544-ec9239c78317"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("6df6510d-b37f-41f1-8606-9f92778811de"),
                            Slug: unifiedgosdk.Pointer("censura"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "toties",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/1319/1257?lock=7280448425732025",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("25c38385-e463-4e8f-a923-ecef43f8e92b"),
                    Slug: unifiedgosdk.Pointer("aetas"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "consuasor",
                    )),
                },
            },
            Name: "Small Marble Chips",
            PublicDescription: unifiedgosdk.Pointer("Generic Gloves designed with Cotton for miserable performance"),
            PublicName: unifiedgosdk.Pointer("Small Marble Chips"),
            Tags: []string{
                "ambulo",
                "adeptio",
                "contego",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-15T01:18:21.480Z"),
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

## CreateCommerceInventory

Create an inventory

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceInventory" method="post" path="/commerce/{connection_id}/inventory" example="commerce_inventory" -->
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

    res, err := s.Commerce.CreateCommerceInventory(ctx, operations.CreateCommerceInventoryRequest{
        CommerceInventory: shared.CommerceInventory{
            Available: unifiedgosdk.Pointer[float64](337.0),
            ID: unifiedgosdk.Pointer("b3ba51df-a271-4054-9b23-658b4ddf92d9"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T09:48:39.178Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateCommerceInventoryRequest](../../pkg/models/operations/createcommerceinventoryrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateCommerceInventoryResponse](../../pkg/models/operations/createcommerceinventoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCommerceItem

Create an item

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

    res, err := s.Commerce.CreateCommerceItem(ctx, operations.CreateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("b932dc7d-d4ba-4c4f-9234-98804f973add"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("d49b8ce9-d4c0-4aca-9fa5-bd21a100e401"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a65c394d-c43b-4776-98e5-71eb1f9fdb41"),
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
                    ID: unifiedgosdk.Pointer("8fff6f50-742f-453a-800c-20bda1039b3a"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("1552eab6-e15b-4322-9fea-dad4178a11e9"),
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
                    ID: unifiedgosdk.Pointer("85f2a1e8-f760-443f-ac8b-c6a1c1831565"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-19T12:01:33.154Z"),
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

    res, err := s.Commerce.CreateCommerceItemvariant(ctx, operations.CreateCommerceItemvariantRequest{
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

## CreateCommerceLocation

Create a location

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceLocation" method="post" path="/commerce/{connection_id}/location" example="commerce_location" -->
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

    res, err := s.Commerce.CreateCommerceLocation(ctx, operations.CreateCommerceLocationRequest{
        CommerceLocation: shared.CommerceLocation{
            Address: &shared.PropertyCommerceLocationAddress{
                Address1: unifiedgosdk.Pointer("29896 The Limes"),
                City: unifiedgosdk.Pointer("New Kenny"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("14490-0609"),
                Region: unifiedgosdk.Pointer("Virginia"),
                RegionCode: unifiedgosdk.Pointer("MS"),
            },
            Categories: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-12-29T04:15:21.195Z"),
            Currency: unifiedgosdk.Pointer("XCD"),
            Description: unifiedgosdk.Pointer("Adsidue audentia."),
            ID: unifiedgosdk.Pointer("cfb9fd29-22d5-44ed-9d15-7b8865a2d9d9"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/hjFt1/1036/2220"),
            IsActive: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("vulgaris"),
            Latitude: unifiedgosdk.Pointer[float64](0.0),
            LocationType: shared.LocationTypeRestaurant.ToPointer(),
            Longitude: unifiedgosdk.Pointer[float64](0.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Addo."),
                    Height: unifiedgosdk.Pointer[float64](283.0),
                    ID: unifiedgosdk.Pointer("1bc20b62-1ed8-468f-a07a-1270eb5f886d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("2db7ffe1-096a-4d3b-ae95-79468699a9df"),
                            Slug: unifiedgosdk.Pointer("abutor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "damno",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](40.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/QVh7ViTV/3964/1567",
                    Width: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Name: unifiedgosdk.Pointer("Olson - Mraz"),
            PriceLevel: unifiedgosdk.Pointer(""),
            Rating: unifiedgosdk.Pointer[float64](0.0),
            ReviewCount: unifiedgosdk.Pointer[float64](0.0),
            Telephones: []shared.CommerceTelephone{
                shared.CommerceTelephone{
                    Telephone: "(872) 522-3201",
                    Type: shared.CommerceTelephoneTypeOther.ToPointer(),
                },
                shared.CommerceTelephone{
                    Telephone: "(236) 274-2445",
                    Type: shared.CommerceTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-04-20T17:56:50.699Z"),
            WebURL: unifiedgosdk.Pointer("https://chilly-edge.info"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateCommerceLocationRequest](../../pkg/models/operations/createcommercelocationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateCommerceLocationResponse](../../pkg/models/operations/createcommercelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCommerceReservation

Create a reservation

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceReservation" method="post" path="/commerce/{connection_id}/reservation" example="commerce_reservation" -->
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

    res, err := s.Commerce.CreateCommerceReservation(ctx, operations.CreateCommerceReservationRequest{
        CommerceReservation: shared.CommerceReservation{
            CreatedAt: types.MustNewTimeFromString("2021-12-14T19:50:31.151Z"),
            EndAt: types.MustNewTimeFromString("2022-01-01T22:00:17.868Z"),
            GuestEmail: unifiedgosdk.Pointer("julien.bartoletti@icloud.com"),
            GuestName: unifiedgosdk.Pointer("Julien Bartoletti"),
            GuestPhone: unifiedgosdk.Pointer("(728) 596-2620"),
            ID: unifiedgosdk.Pointer("0ef71e6c-d9d9-4c9f-b62a-3f0f0f56fb0f"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-03T21:23:50.117Z"),
            URL: unifiedgosdk.Pointer("https://expensive-gift.org"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateCommerceReservationRequest](../../pkg/models/operations/createcommercereservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateCommerceReservationResponse](../../pkg/models/operations/createcommercereservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCommerceReview

Create a review

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceReview" method="post" path="/commerce/{connection_id}/review" example="commerce_review" -->
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

    res, err := s.Commerce.CreateCommerceReview(ctx, operations.CreateCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("a701a537-140d-43ca-82c9-18ce8f785886"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("7d289800-5f77-4fbe-92b2-b1eb3fa9e31e"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("fdc2101f-e2db-4cd0-be49-c031a306184a"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("6ee6be79-1ed9-45eb-a395-3bd13d182f01"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("8b9b79fe-905f-496d-9674-d3111c203fe4"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T08:42:36.697Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateCommerceReviewRequest](../../pkg/models/operations/createcommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateCommerceReviewResponse](../../pkg/models/operations/createcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCommerceSaleschannel

Create a saleschannel

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceSaleschannel" method="post" path="/commerce/{connection_id}/saleschannel" example="commerce_saleschannel" -->
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

    res, err := s.Commerce.CreateCommerceSaleschannel(ctx, operations.CreateCommerceSaleschannelRequest{
        CommerceSaleschannel: shared.CommerceSaleschannel{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2021-12-12T06:19:55.421Z"),
            Description: unifiedgosdk.Pointer("Utroque denuncio solutio."),
            ID: unifiedgosdk.Pointer("cc3de979-f65e-4409-9afc-44337c992d3c"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-01-28T00:01:54.549Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceSaleschannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateCommerceSaleschannelRequest](../../pkg/models/operations/createcommercesaleschannelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateCommerceSaleschannelResponse](../../pkg/models/operations/createcommercesaleschannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceCollection

Retrieve a collection

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

    res, err := s.Commerce.GetCommerceCollection(ctx, operations.GetCommerceCollectionRequest{
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

## GetCommerceInventory

Retrieve an inventory

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceInventory" method="get" path="/commerce/{connection_id}/inventory/{id}" -->
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

    res, err := s.Commerce.GetCommerceInventory(ctx, operations.GetCommerceInventoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCommerceInventoryRequest](../../pkg/models/operations/getcommerceinventoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetCommerceInventoryResponse](../../pkg/models/operations/getcommerceinventoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceItem

Retrieve an item

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

    res, err := s.Commerce.GetCommerceItem(ctx, operations.GetCommerceItemRequest{
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

    res, err := s.Commerce.GetCommerceItemvariant(ctx, operations.GetCommerceItemvariantRequest{
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

## GetCommerceLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceLocation" method="get" path="/commerce/{connection_id}/location/{id}" -->
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

    res, err := s.Commerce.GetCommerceLocation(ctx, operations.GetCommerceLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCommerceLocationRequest](../../pkg/models/operations/getcommercelocationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetCommerceLocationResponse](../../pkg/models/operations/getcommercelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceReservation

Retrieve a reservation

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceReservation" method="get" path="/commerce/{connection_id}/reservation/{id}" -->
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

    res, err := s.Commerce.GetCommerceReservation(ctx, operations.GetCommerceReservationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCommerceReservationRequest](../../pkg/models/operations/getcommercereservationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetCommerceReservationResponse](../../pkg/models/operations/getcommercereservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceReview

Retrieve a review

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceReview" method="get" path="/commerce/{connection_id}/review/{id}" -->
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

    res, err := s.Commerce.GetCommerceReview(ctx, operations.GetCommerceReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetCommerceReviewRequest](../../pkg/models/operations/getcommercereviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetCommerceReviewResponse](../../pkg/models/operations/getcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceSaleschannel

Retrieve a saleschannel

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceSaleschannel" method="get" path="/commerce/{connection_id}/saleschannel/{id}" -->
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

    res, err := s.Commerce.GetCommerceSaleschannel(ctx, operations.GetCommerceSaleschannelRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceSaleschannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCommerceSaleschannelRequest](../../pkg/models/operations/getcommercesaleschannelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetCommerceSaleschannelResponse](../../pkg/models/operations/getcommercesaleschannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceAvailabilities

List all availabilities

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceAvailabilities" method="get" path="/commerce/{connection_id}/availability" -->
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

    res, err := s.Commerce.ListCommerceAvailabilities(ctx, operations.ListCommerceAvailabilitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceAvailabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListCommerceAvailabilitiesRequest](../../pkg/models/operations/listcommerceavailabilitiesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListCommerceAvailabilitiesResponse](../../pkg/models/operations/listcommerceavailabilitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceCollections

List all collections

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

    res, err := s.Commerce.ListCommerceCollections(ctx, operations.ListCommerceCollectionsRequest{
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

## ListCommerceInventories

List all inventories

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceInventories" method="get" path="/commerce/{connection_id}/inventory" -->
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

    res, err := s.Commerce.ListCommerceInventories(ctx, operations.ListCommerceInventoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceInventories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListCommerceInventoriesRequest](../../pkg/models/operations/listcommerceinventoriesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListCommerceInventoriesResponse](../../pkg/models/operations/listcommerceinventoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceItems

List all items

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

    res, err := s.Commerce.ListCommerceItems(ctx, operations.ListCommerceItemsRequest{
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

    res, err := s.Commerce.ListCommerceItemvariants(ctx, operations.ListCommerceItemvariantsRequest{
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

## ListCommerceLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceLocations" method="get" path="/commerce/{connection_id}/location" -->
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

    res, err := s.Commerce.ListCommerceLocations(ctx, operations.ListCommerceLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCommerceLocationsRequest](../../pkg/models/operations/listcommercelocationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListCommerceLocationsResponse](../../pkg/models/operations/listcommercelocationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceReservations

List all reservations

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceReservations" method="get" path="/commerce/{connection_id}/reservation" -->
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

    res, err := s.Commerce.ListCommerceReservations(ctx, operations.ListCommerceReservationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReservations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCommerceReservationsRequest](../../pkg/models/operations/listcommercereservationsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCommerceReservationsResponse](../../pkg/models/operations/listcommercereservationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceReviews

List all reviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceReviews" method="get" path="/commerce/{connection_id}/review" -->
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

    res, err := s.Commerce.ListCommerceReviews(ctx, operations.ListCommerceReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListCommerceReviewsRequest](../../pkg/models/operations/listcommercereviewsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListCommerceReviewsResponse](../../pkg/models/operations/listcommercereviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceSaleschannels

List all saleschannels

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceSaleschannels" method="get" path="/commerce/{connection_id}/saleschannel" -->
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

    res, err := s.Commerce.ListCommerceSaleschannels(ctx, operations.ListCommerceSaleschannelsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceSaleschannels != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListCommerceSaleschannelsRequest](../../pkg/models/operations/listcommercesaleschannelsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListCommerceSaleschannelsResponse](../../pkg/models/operations/listcommercesaleschannelsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceCollection

Update a collection

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

    res, err := s.Commerce.PatchCommerceCollection(ctx, operations.PatchCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("b9e54251-c92b-4782-b1bb-926437414e65"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("b5d8356e-a5c4-4e37-88f1-f0f6c4baf524"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("be896d4e-ad51-4055-86be-b2da90d4641b"),
                            Slug: unifiedgosdk.Pointer("censura"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "toties",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/1319/1257?lock=7280448425732025",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("e6a20238-d170-498a-be5d-6832a5df7d28"),
                    Slug: unifiedgosdk.Pointer("aetas"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "consuasor",
                    )),
                },
            },
            Name: "Small Marble Chips",
            PublicDescription: unifiedgosdk.Pointer("Generic Gloves designed with Cotton for miserable performance"),
            PublicName: unifiedgosdk.Pointer("Small Marble Chips"),
            Tags: []string{
                "ambulo",
                "adeptio",
                "contego",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-15T01:18:21.493Z"),
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

## PatchCommerceInventory

Update an inventory

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceInventory" method="patch" path="/commerce/{connection_id}/inventory/{id}" example="commerce_inventory" -->
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

    res, err := s.Commerce.PatchCommerceInventory(ctx, operations.PatchCommerceInventoryRequest{
        CommerceInventory: shared.CommerceInventory{
            Available: unifiedgosdk.Pointer[float64](337.0),
            ID: unifiedgosdk.Pointer("505d8fdb-6d3e-4cc6-8aca-7a9b4ef5be26"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T09:48:39.185Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchCommerceInventoryRequest](../../pkg/models/operations/patchcommerceinventoryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchCommerceInventoryResponse](../../pkg/models/operations/patchcommerceinventoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceItem

Update an item

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

    res, err := s.Commerce.PatchCommerceItem(ctx, operations.PatchCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("c13a38fe-fad4-4198-bf4a-7b5dc5f77b0d"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("9bda371f-7ce3-4f50-b725-587283bf4f2d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("79b75953-de77-4893-9da0-5b030a19f949"),
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
                    ID: unifiedgosdk.Pointer("99807e35-840c-4c23-bb17-9c5b9c3fd475"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("78ff5b2d-8ab5-4eab-a0d4-abb192ef8de7"),
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
                    ID: unifiedgosdk.Pointer("10b73839-3f28-45d3-b97e-b2c5fe46fa8f"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-19T12:01:33.164Z"),
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

    res, err := s.Commerce.PatchCommerceItemvariant(ctx, operations.PatchCommerceItemvariantRequest{
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

## PatchCommerceLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceLocation" method="patch" path="/commerce/{connection_id}/location/{id}" example="commerce_location" -->
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

    res, err := s.Commerce.PatchCommerceLocation(ctx, operations.PatchCommerceLocationRequest{
        CommerceLocation: shared.CommerceLocation{
            Address: &shared.PropertyCommerceLocationAddress{
                Address1: unifiedgosdk.Pointer("29896 The Limes"),
                City: unifiedgosdk.Pointer("New Kenny"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("14490-0609"),
                Region: unifiedgosdk.Pointer("Virginia"),
                RegionCode: unifiedgosdk.Pointer("MS"),
            },
            Categories: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-12-29T04:15:21.195Z"),
            Currency: unifiedgosdk.Pointer("XCD"),
            Description: unifiedgosdk.Pointer("Adsidue audentia."),
            ID: unifiedgosdk.Pointer("40a4434d-2376-4620-9003-349464d19ca7"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/hjFt1/1036/2220"),
            IsActive: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("vulgaris"),
            Latitude: unifiedgosdk.Pointer[float64](0.0),
            LocationType: shared.LocationTypeRestaurant.ToPointer(),
            Longitude: unifiedgosdk.Pointer[float64](0.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Addo."),
                    Height: unifiedgosdk.Pointer[float64](283.0),
                    ID: unifiedgosdk.Pointer("3e6802f0-02a0-4d54-b944-d6a4b261d631"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("0b3ca168-ec82-4454-aaac-eff637558b65"),
                            Slug: unifiedgosdk.Pointer("abutor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "damno",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](40.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/QVh7ViTV/3964/1567",
                    Width: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Name: unifiedgosdk.Pointer("Olson - Mraz"),
            PriceLevel: unifiedgosdk.Pointer(""),
            Rating: unifiedgosdk.Pointer[float64](0.0),
            ReviewCount: unifiedgosdk.Pointer[float64](0.0),
            Telephones: []shared.CommerceTelephone{
                shared.CommerceTelephone{
                    Telephone: "(872) 522-3201",
                    Type: shared.CommerceTelephoneTypeOther.ToPointer(),
                },
                shared.CommerceTelephone{
                    Telephone: "(236) 274-2445",
                    Type: shared.CommerceTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-04-20T17:56:50.709Z"),
            WebURL: unifiedgosdk.Pointer("https://chilly-edge.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchCommerceLocationRequest](../../pkg/models/operations/patchcommercelocationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchCommerceLocationResponse](../../pkg/models/operations/patchcommercelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceReservation

Update a reservation

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceReservation" method="patch" path="/commerce/{connection_id}/reservation/{id}" example="commerce_reservation" -->
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

    res, err := s.Commerce.PatchCommerceReservation(ctx, operations.PatchCommerceReservationRequest{
        CommerceReservation: shared.CommerceReservation{
            CreatedAt: types.MustNewTimeFromString("2021-12-14T19:50:31.151Z"),
            EndAt: types.MustNewTimeFromString("2022-01-01T22:00:17.868Z"),
            GuestEmail: unifiedgosdk.Pointer("julien.bartoletti@icloud.com"),
            GuestName: unifiedgosdk.Pointer("Julien Bartoletti"),
            GuestPhone: unifiedgosdk.Pointer("(728) 596-2620"),
            ID: unifiedgosdk.Pointer("4c9dd695-6c67-4102-9022-4ea5300139ad"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-03T21:23:50.121Z"),
            URL: unifiedgosdk.Pointer("https://expensive-gift.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCommerceReservationRequest](../../pkg/models/operations/patchcommercereservationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchCommerceReservationResponse](../../pkg/models/operations/patchcommercereservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceReview" method="patch" path="/commerce/{connection_id}/review/{id}" example="commerce_review" -->
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

    res, err := s.Commerce.PatchCommerceReview(ctx, operations.PatchCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("578fe127-194b-4d01-a25f-c54fc1709240"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("65be1798-462d-44e5-84c9-515cac1619eb"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9f916c18-3b91-4219-8d26-0639d2363609"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("c1af0dea-75ef-4e0b-a7aa-d6d3b840ec10"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9ad8d0f3-c8b8-4fd9-a4a7-53a7062bf74d"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T08:42:36.722Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchCommerceReviewRequest](../../pkg/models/operations/patchcommercereviewrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchCommerceReviewResponse](../../pkg/models/operations/patchcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceSaleschannel

Update a saleschannel

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceSaleschannel" method="patch" path="/commerce/{connection_id}/saleschannel/{id}" example="commerce_saleschannel" -->
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

    res, err := s.Commerce.PatchCommerceSaleschannel(ctx, operations.PatchCommerceSaleschannelRequest{
        CommerceSaleschannel: shared.CommerceSaleschannel{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2021-12-12T06:19:55.421Z"),
            Description: unifiedgosdk.Pointer("Utroque denuncio solutio."),
            ID: unifiedgosdk.Pointer("dcf669f2-72fe-4f86-bf18-a0776d1f5b30"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-01-28T00:01:54.554Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceSaleschannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCommerceSaleschannelRequest](../../pkg/models/operations/patchcommercesaleschannelrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PatchCommerceSaleschannelResponse](../../pkg/models/operations/patchcommercesaleschannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceCollection

Remove a collection

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

    res, err := s.Commerce.RemoveCommerceCollection(ctx, operations.RemoveCommerceCollectionRequest{
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

## RemoveCommerceInventory

Remove an inventory

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceInventory" method="delete" path="/commerce/{connection_id}/inventory/{id}" -->
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

    res, err := s.Commerce.RemoveCommerceInventory(ctx, operations.RemoveCommerceInventoryRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveCommerceInventoryRequest](../../pkg/models/operations/removecommerceinventoryrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveCommerceInventoryResponse](../../pkg/models/operations/removecommerceinventoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceItem

Remove an item

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

    res, err := s.Commerce.RemoveCommerceItem(ctx, operations.RemoveCommerceItemRequest{
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

    res, err := s.Commerce.RemoveCommerceItemvariant(ctx, operations.RemoveCommerceItemvariantRequest{
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

## RemoveCommerceLocation

Remove a location

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceLocation" method="delete" path="/commerce/{connection_id}/location/{id}" -->
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

    res, err := s.Commerce.RemoveCommerceLocation(ctx, operations.RemoveCommerceLocationRequest{
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
| `request`                                                                                                | [operations.RemoveCommerceLocationRequest](../../pkg/models/operations/removecommercelocationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveCommerceLocationResponse](../../pkg/models/operations/removecommercelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceReservation

Remove a reservation

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceReservation" method="delete" path="/commerce/{connection_id}/reservation/{id}" -->
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

    res, err := s.Commerce.RemoveCommerceReservation(ctx, operations.RemoveCommerceReservationRequest{
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
| `request`                                                                                                      | [operations.RemoveCommerceReservationRequest](../../pkg/models/operations/removecommercereservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemoveCommerceReservationResponse](../../pkg/models/operations/removecommercereservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceReview

Remove a review

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceReview" method="delete" path="/commerce/{connection_id}/review/{id}" -->
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

    res, err := s.Commerce.RemoveCommerceReview(ctx, operations.RemoveCommerceReviewRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveCommerceReviewRequest](../../pkg/models/operations/removecommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveCommerceReviewResponse](../../pkg/models/operations/removecommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceSaleschannel

Remove a saleschannel

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceSaleschannel" method="delete" path="/commerce/{connection_id}/saleschannel/{id}" -->
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

    res, err := s.Commerce.RemoveCommerceSaleschannel(ctx, operations.RemoveCommerceSaleschannelRequest{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RemoveCommerceSaleschannelRequest](../../pkg/models/operations/removecommercesaleschannelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.RemoveCommerceSaleschannelResponse](../../pkg/models/operations/removecommercesaleschannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceCollection

Update a collection

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

    res, err := s.Commerce.UpdateCommerceCollection(ctx, operations.UpdateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("b9e54251-c92b-4782-b1bb-926437414e65"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("b5d8356e-a5c4-4e37-88f1-f0f6c4baf524"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("be896d4e-ad51-4055-86be-b2da90d4641b"),
                            Slug: unifiedgosdk.Pointer("censura"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "toties",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](80.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/1319/1257?lock=7280448425732025",
                    Width: unifiedgosdk.Pointer[float64](40.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ID: unifiedgosdk.Pointer("e6a20238-d170-498a-be5d-6832a5df7d28"),
                    Slug: unifiedgosdk.Pointer("aetas"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "consuasor",
                    )),
                },
            },
            Name: "Small Marble Chips",
            PublicDescription: unifiedgosdk.Pointer("Generic Gloves designed with Cotton for miserable performance"),
            PublicName: unifiedgosdk.Pointer("Small Marble Chips"),
            Tags: []string{
                "ambulo",
                "adeptio",
                "contego",
            },
            Type: shared.CommerceCollectionTypeCollection.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-15T01:18:21.493Z"),
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

## UpdateCommerceInventory

Update an inventory

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceInventory" method="put" path="/commerce/{connection_id}/inventory/{id}" example="commerce_inventory" -->
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

    res, err := s.Commerce.UpdateCommerceInventory(ctx, operations.UpdateCommerceInventoryRequest{
        CommerceInventory: shared.CommerceInventory{
            Available: unifiedgosdk.Pointer[float64](337.0),
            ID: unifiedgosdk.Pointer("505d8fdb-6d3e-4cc6-8aca-7a9b4ef5be26"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T09:48:39.185Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateCommerceInventoryRequest](../../pkg/models/operations/updatecommerceinventoryrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateCommerceInventoryResponse](../../pkg/models/operations/updatecommerceinventoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceItem

Update an item

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

    res, err := s.Commerce.UpdateCommerceItem(ctx, operations.UpdateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("c13a38fe-fad4-4198-bf4a-7b5dc5f77b0d"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("9bda371f-7ce3-4f50-b725-587283bf4f2d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("79b75953-de77-4893-9da0-5b030a19f949"),
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
                    ID: unifiedgosdk.Pointer("99807e35-840c-4c23-bb17-9c5b9c3fd475"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("78ff5b2d-8ab5-4eab-a0d4-abb192ef8de7"),
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
                    ID: unifiedgosdk.Pointer("10b73839-3f28-45d3-b97e-b2c5fe46fa8f"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-19T12:01:33.164Z"),
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

    res, err := s.Commerce.UpdateCommerceItemvariant(ctx, operations.UpdateCommerceItemvariantRequest{
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

## UpdateCommerceLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceLocation" method="put" path="/commerce/{connection_id}/location/{id}" example="commerce_location" -->
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

    res, err := s.Commerce.UpdateCommerceLocation(ctx, operations.UpdateCommerceLocationRequest{
        CommerceLocation: shared.CommerceLocation{
            Address: &shared.PropertyCommerceLocationAddress{
                Address1: unifiedgosdk.Pointer("29896 The Limes"),
                City: unifiedgosdk.Pointer("New Kenny"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("14490-0609"),
                Region: unifiedgosdk.Pointer("Virginia"),
                RegionCode: unifiedgosdk.Pointer("MS"),
            },
            Categories: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-12-29T04:15:21.195Z"),
            Currency: unifiedgosdk.Pointer("XCD"),
            Description: unifiedgosdk.Pointer("Adsidue audentia."),
            ID: unifiedgosdk.Pointer("40a4434d-2376-4620-9003-349464d19ca7"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/hjFt1/1036/2220"),
            IsActive: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("vulgaris"),
            Latitude: unifiedgosdk.Pointer[float64](0.0),
            LocationType: shared.LocationTypeRestaurant.ToPointer(),
            Longitude: unifiedgosdk.Pointer[float64](0.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Addo."),
                    Height: unifiedgosdk.Pointer[float64](283.0),
                    ID: unifiedgosdk.Pointer("3e6802f0-02a0-4d54-b944-d6a4b261d631"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("0b3ca168-ec82-4454-aaac-eff637558b65"),
                            Slug: unifiedgosdk.Pointer("abutor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "damno",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](40.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/QVh7ViTV/3964/1567",
                    Width: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Name: unifiedgosdk.Pointer("Olson - Mraz"),
            PriceLevel: unifiedgosdk.Pointer(""),
            Rating: unifiedgosdk.Pointer[float64](0.0),
            ReviewCount: unifiedgosdk.Pointer[float64](0.0),
            Telephones: []shared.CommerceTelephone{
                shared.CommerceTelephone{
                    Telephone: "(872) 522-3201",
                    Type: shared.CommerceTelephoneTypeOther.ToPointer(),
                },
                shared.CommerceTelephone{
                    Telephone: "(236) 274-2445",
                    Type: shared.CommerceTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-04-20T17:56:50.709Z"),
            WebURL: unifiedgosdk.Pointer("https://chilly-edge.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateCommerceLocationRequest](../../pkg/models/operations/updatecommercelocationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateCommerceLocationResponse](../../pkg/models/operations/updatecommercelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceReservation

Update a reservation

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceReservation" method="put" path="/commerce/{connection_id}/reservation/{id}" example="commerce_reservation" -->
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

    res, err := s.Commerce.UpdateCommerceReservation(ctx, operations.UpdateCommerceReservationRequest{
        CommerceReservation: shared.CommerceReservation{
            CreatedAt: types.MustNewTimeFromString("2021-12-14T19:50:31.151Z"),
            EndAt: types.MustNewTimeFromString("2022-01-01T22:00:17.868Z"),
            GuestEmail: unifiedgosdk.Pointer("julien.bartoletti@icloud.com"),
            GuestName: unifiedgosdk.Pointer("Julien Bartoletti"),
            GuestPhone: unifiedgosdk.Pointer("(728) 596-2620"),
            ID: unifiedgosdk.Pointer("4c9dd695-6c67-4102-9022-4ea5300139ad"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-03T21:23:50.121Z"),
            URL: unifiedgosdk.Pointer("https://expensive-gift.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateCommerceReservationRequest](../../pkg/models/operations/updatecommercereservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateCommerceReservationResponse](../../pkg/models/operations/updatecommercereservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceReview" method="put" path="/commerce/{connection_id}/review/{id}" example="commerce_review" -->
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

    res, err := s.Commerce.UpdateCommerceReview(ctx, operations.UpdateCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("578fe127-194b-4d01-a25f-c54fc1709240"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("65be1798-462d-44e5-84c9-515cac1619eb"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9f916c18-3b91-4219-8d26-0639d2363609"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("c1af0dea-75ef-4e0b-a7aa-d6d3b840ec10"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9ad8d0f3-c8b8-4fd9-a4a7-53a7062bf74d"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T08:42:36.722Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateCommerceReviewRequest](../../pkg/models/operations/updatecommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateCommerceReviewResponse](../../pkg/models/operations/updatecommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceSaleschannel

Update a saleschannel

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceSaleschannel" method="put" path="/commerce/{connection_id}/saleschannel/{id}" example="commerce_saleschannel" -->
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

    res, err := s.Commerce.UpdateCommerceSaleschannel(ctx, operations.UpdateCommerceSaleschannelRequest{
        CommerceSaleschannel: shared.CommerceSaleschannel{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2021-12-12T06:19:55.421Z"),
            Description: unifiedgosdk.Pointer("Utroque denuncio solutio."),
            ID: unifiedgosdk.Pointer("dcf669f2-72fe-4f86-bf18-a0776d1f5b30"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-01-28T00:01:54.554Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceSaleschannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateCommerceSaleschannelRequest](../../pkg/models/operations/updatecommercesaleschannelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateCommerceSaleschannelResponse](../../pkg/models/operations/updatecommercesaleschannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |