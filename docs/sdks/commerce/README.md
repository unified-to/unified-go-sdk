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

    res, err := s.Commerce.CreateCommerceCollection(ctx, operations.CreateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("73642861-de01-4f41-959c-f609912ee7c0"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("3acdd72f-26ba-48e8-9c58-8e154e034aca"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("58282662-f331-4c7e-bb19-2dbcdb901d17"),
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
                    ID: unifiedgosdk.Pointer("fd7e5bba-cb3c-4311-8c4a-912b3b82d727"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-19T13:22:31.120Z"),
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

Creates a CommerceInventory. Provide the unified CommerceInventory fields in the request body; they are mapped to the provider's API. Returns the created CommerceInventory.

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
            ID: unifiedgosdk.Pointer("4202a291-bd9a-475a-926a-e1caa11c2ede"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T12:22:50.104Z"),
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

    res, err := s.Commerce.CreateCommerceItem(ctx, operations.CreateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("2d589172-9f32-402f-bfe3-c0b423e05242"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("ca17c266-e02b-4da9-a48c-0bccaa9535f3"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("7e27f6dd-e0de-4112-b64a-69b29e919326"),
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
                    ID: unifiedgosdk.Pointer("3e62e0ba-20a4-4098-a6a0-1047c35cc230"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("2359832c-a5e0-4d16-a558-3152ca5056b6"),
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
                    ID: unifiedgosdk.Pointer("c99ac2b3-ff24-4d6e-8d4d-ab9a6024dd31"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-22T21:39:07.435Z"),
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

    res, err := s.Commerce.CreateCommerceItemvariant(ctx, operations.CreateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("4450ef44-f930-42b3-9cb2-35fe77bfc272"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("bcc5ec37-f25f-474c-a166-2d4016954538"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("30bfc68c-eeeb-4185-b3b0-f2c26a6f5172"),
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
                    ID: unifiedgosdk.Pointer("67a73491-6c98-4efe-a71b-1d4db9493ce9"),
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
                    ID: unifiedgosdk.Pointer("15b11fcf-6d7c-49c8-96b3-11376778f520"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("20577685-5366-4b5b-bfa6-99467f99758f"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T08:32:21.778Z"),
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

Creates a CommerceLocation. Provide the unified CommerceLocation fields in the request body; they are mapped to the provider's API. Returns the created CommerceLocation.

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
            ID: unifiedgosdk.Pointer("39278dea-ee45-4259-aca0-69f803651cd5"),
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
                    ID: unifiedgosdk.Pointer("0a36e2f4-2fe3-4931-b9e4-ff3ad0c98296"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("6f0e6dc4-9693-4109-8b56-eb7cb389d80c"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-23T18:44:15.809Z"),
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

Creates a CommerceReservation. Provide the unified CommerceReservation fields in the request body; they are mapped to the provider's API. Returns the created CommerceReservation.

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
            ID: unifiedgosdk.Pointer("c4e43ca4-675d-4ba2-a090-0dc34ea3666c"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-05T19:25:16.176Z"),
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

Creates a CommerceReview. Provide the unified CommerceReview fields in the request body; they are mapped to the provider's API. Returns the created CommerceReview.

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
            ID: unifiedgosdk.Pointer("f94f7da2-99b1-4716-9f34-bb1183dd77c7"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("9ebb6c45-7fcf-4145-ba7f-030163ef08d7"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("3e470a8e-dae8-45f7-8c92-ebd1c9f62926"),
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
                    ID: unifiedgosdk.Pointer("37499c23-7147-4591-ab92-dec9802c7055"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("10260a81-5ceb-4ed6-acad-e145a52e4c71"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-08-28T16:25:25.628Z"),
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

Creates a CommerceSaleschannel. Provide the unified CommerceSaleschannel fields in the request body; they are mapped to the provider's API. Returns the created CommerceSaleschannel.

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
            ID: unifiedgosdk.Pointer("692d412c-0c78-4c57-bc04-1f192aacf9f7"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-02T16:25:02.296Z"),
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

Retrieves a single CommerceInventory by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CommerceLocation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CommerceReservation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CommerceReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CommerceSaleschannel by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of CommerceAvailabilities. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CommerceInventories. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CommerceLocations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CommerceReservations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CommerceReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CommerceSaleschannels. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Commerce.PatchCommerceCollection(ctx, operations.PatchCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("098d3ee7-9b16-4bc5-8d72-1be3ca83bbec"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("d24906c0-b7f1-4f93-84cf-d247d98b0f48"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("e1b2cecb-34e8-408b-ba5e-d35edfd757f1"),
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
                    ID: unifiedgosdk.Pointer("a939594c-7bcd-476d-9704-a80d9687b604"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-19T13:22:31.139Z"),
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

Updates the CommerceInventory identified by `id`. Provide the unified CommerceInventory fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("058cabd3-25ca-4d2c-8b3c-34ca426a1fc9"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T12:22:50.120Z"),
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

    res, err := s.Commerce.PatchCommerceItem(ctx, operations.PatchCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("989c08d1-0175-464f-8afc-a5ecf2b7faf9"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("2d0223fb-5d10-41fe-9757-f1959696c1be"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9353c024-09dc-48e5-bd35-49c9c8cb2573"),
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
                    ID: unifiedgosdk.Pointer("3c6950f9-4207-488c-a32b-0c61027b4f56"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("34a66ed2-357c-4c81-bd11-bd64303a13a3"),
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
                    ID: unifiedgosdk.Pointer("81d99791-bf34-4e8e-9412-29a8f6d1eaa4"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-22T21:39:07.449Z"),
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

    res, err := s.Commerce.PatchCommerceItemvariant(ctx, operations.PatchCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("1f44024e-8ccf-47ce-bc28-51babb67fbd5"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("c6e7cfe9-3afb-4da3-b9e1-93100b918b3c"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("825122ed-0979-4e91-bf2c-91eb47af0494"),
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
                    ID: unifiedgosdk.Pointer("2fada3a1-5f94-48b4-aa5a-b9e35cc87876"),
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
                    ID: unifiedgosdk.Pointer("36429461-db80-40a2-a5c9-876712fdf432"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("eb0a0e17-b887-46c0-bada-2369eafddc8c"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T08:32:21.816Z"),
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

Updates the CommerceLocation identified by `id`. Provide the unified CommerceLocation fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("6489824e-8f90-44fb-a64c-a76321daccc2"),
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
                    ID: unifiedgosdk.Pointer("6ccf67df-cc63-4692-89fb-6090e78d01c8"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("2232653b-e8e1-4c00-a331-a45611b1215e"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-23T18:44:15.827Z"),
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

Updates the CommerceReservation identified by `id`. Provide the unified CommerceReservation fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("319a2d40-429e-43ef-b383-adce260b2372"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-05T19:25:16.180Z"),
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

Updates the CommerceReview identified by `id`. Provide the unified CommerceReview fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("370b1411-7f2a-4808-a238-0ab2cff258c3"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("5c1c8a4d-8d1c-42bc-9a57-8f7a72bcc584"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("851d1bb4-9e1e-4c76-aa0d-fd358513f8c1"),
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
                    ID: unifiedgosdk.Pointer("095cee3f-1f75-4b5b-a45f-e940e092bc8a"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("dd9e78d3-f6cb-4f20-bea9-234b09307324"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-08-28T16:25:25.667Z"),
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

Updates the CommerceSaleschannel identified by `id`. Provide the unified CommerceSaleschannel fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("b3cf5532-3c67-489e-a20b-204fc6090feb"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-02T16:25:02.306Z"),
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

Deletes the CommerceInventory identified by `id`.

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

Deletes the CommerceLocation identified by `id`.

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

Deletes the CommerceReservation identified by `id`.

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

Deletes the CommerceReview identified by `id`.

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

Deletes the CommerceSaleschannel identified by `id`.

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

    res, err := s.Commerce.UpdateCommerceCollection(ctx, operations.UpdateCommerceCollectionRequest{
        CommerceCollection: shared.CommerceCollection{
            CreatedAt: types.MustNewTimeFromString("2023-07-14T00:42:54.742Z"),
            Description: unifiedgosdk.Pointer("The Integrated leading edge website Cheese offers reliable performance and productive design"),
            ID: unifiedgosdk.Pointer("098d3ee7-9b16-4bc5-8d72-1be3ca83bbec"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            ItemMetadata: []shared.CommerceMetadata{},
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Defungo adopto thorax."),
                    Height: unifiedgosdk.Pointer[float64](759.0),
                    ID: unifiedgosdk.Pointer("d24906c0-b7f1-4f93-84cf-d247d98b0f48"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("e1b2cecb-34e8-408b-ba5e-d35edfd757f1"),
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
                    ID: unifiedgosdk.Pointer("a939594c-7bcd-476d-9704-a80d9687b604"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-19T13:22:31.139Z"),
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

Updates the CommerceInventory identified by `id`. Provide the unified CommerceInventory fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("058cabd3-25ca-4d2c-8b3c-34ca426a1fc9"),
            Metadata: []shared.CommerceMetadata{},
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T12:22:50.120Z"),
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

    res, err := s.Commerce.UpdateCommerceItem(ctx, operations.UpdateCommerceItemRequest{
        CommerceItem: shared.CommerceItem{
            Collections: []shared.CommerceReference{},
            CreatedAt: types.MustNewTimeFromString("2019-06-21T20:16:18.628Z"),
            Description: unifiedgosdk.Pointer("Decretum vulnero ustulo abeo."),
            Duration: unifiedgosdk.Pointer[float64](104.0),
            GlobalCode: unifiedgosdk.Pointer("978-0-85270-680-0"),
            ID: unifiedgosdk.Pointer("989c08d1-0175-464f-8afc-a5ecf2b7faf9"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsTaxable: unifiedgosdk.Pointer(true),
            IsVisible: unifiedgosdk.Pointer(true),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Caterva eveniet acies candidus."),
                    Height: unifiedgosdk.Pointer[float64](876.0),
                    ID: unifiedgosdk.Pointer("2d0223fb-5d10-41fe-9757-f1959696c1be"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("9353c024-09dc-48e5-bd35-49c9c8cb2573"),
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
                    ID: unifiedgosdk.Pointer("3c6950f9-4207-488c-a32b-0c61027b4f56"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("34a66ed2-357c-4c81-bd11-bd64303a13a3"),
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
                    ID: unifiedgosdk.Pointer("81d99791-bf34-4e8e-9412-29a8f6d1eaa4"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-04-22T21:39:07.449Z"),
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

    res, err := s.Commerce.UpdateCommerceItemvariant(ctx, operations.UpdateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-06T18:14:41.644Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Stylish Salad designed to make you stand out with little looks"),
            Height: unifiedgosdk.Pointer[float64](54.0),
            ID: unifiedgosdk.Pointer("1f44024e-8ccf-47ce-bc28-51babb67fbd5"),
            IsActive: unifiedgosdk.Pointer(true),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](92.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Hic desipio thorax."),
                    Height: unifiedgosdk.Pointer[float64](961.0),
                    ID: unifiedgosdk.Pointer("c6e7cfe9-3afb-4da3-b9e1-93100b918b3c"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("825122ed-0979-4e91-bf2c-91eb47af0494"),
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
                    ID: unifiedgosdk.Pointer("2fada3a1-5f94-48b4-aa5a-b9e35cc87876"),
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
                    ID: unifiedgosdk.Pointer("36429461-db80-40a2-a5c9-876712fdf432"),
                    Name: "Color",
                    Position: unifiedgosdk.Pointer[float64](1.0),
                    Values: []string{
                        "olive",
                    },
                },
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("eb0a0e17-b887-46c0-bada-2369eafddc8c"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T08:32:21.816Z"),
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

Updates the CommerceLocation identified by `id`. Provide the unified CommerceLocation fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("6489824e-8f90-44fb-a64c-a76321daccc2"),
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
                    ID: unifiedgosdk.Pointer("6ccf67df-cc63-4692-89fb-6090e78d01c8"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("2232653b-e8e1-4c00-a331-a45611b1215e"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-23T18:44:15.827Z"),
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

Updates the CommerceReservation identified by `id`. Provide the unified CommerceReservation fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("319a2d40-429e-43ef-b383-adce260b2372"),
            Metadata: []shared.CommerceMetadata{},
            Notes: unifiedgosdk.Pointer("Adsum textilis ipsum despecto."),
            Size: unifiedgosdk.Pointer[float64](10.0),
            StartAt: types.MustNewTimeFromString("2021-12-18T00:40:25.125Z"),
            Status: shared.CommerceReservationStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-05T19:25:16.180Z"),
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

Updates the CommerceReview identified by `id`. Provide the unified CommerceReview fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("370b1411-7f2a-4808-a238-0ab2cff258c3"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("5c1c8a4d-8d1c-42bc-9a57-8f7a72bcc584"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("851d1bb4-9e1e-4c76-aa0d-fd358513f8c1"),
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
                    ID: unifiedgosdk.Pointer("095cee3f-1f75-4b5b-a45f-e940e092bc8a"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("dd9e78d3-f6cb-4f20-bea9-234b09307324"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-08-28T16:25:25.667Z"),
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

Updates the CommerceSaleschannel identified by `id`. Provide the unified CommerceSaleschannel fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("b3cf5532-3c67-489e-a20b-204fc6090feb"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.CommerceMetadata{},
            Slug: unifiedgosdk.Pointer("amiculum-congregatio-suspendo"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-02T16:25:02.306Z"),
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