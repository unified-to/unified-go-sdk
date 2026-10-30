# Listing

## Overview

### Available Operations

* [CreatePropertyListing](#createpropertylisting) - Create a listing
* [GetPropertyListing](#getpropertylisting) - Retrieve a listing
* [ListPropertyListings](#listpropertylistings) - List all listings
* [PatchPropertyListing](#patchpropertylisting) - Update a listing
* [RemovePropertyListing](#removepropertylisting) - Remove a listing
* [UpdatePropertyListing](#updatepropertylisting) - Update a listing

## CreatePropertyListing

Creates a PropertyListing. Provide the unified PropertyListing fields in the request body; they are mapped to the provider's API. Returns the created PropertyListing.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyListing" method="post" path="/property/{connection_id}/listing" example="property_listing" -->
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

    res, err := s.Listing.CreatePropertyListing(ctx, operations.CreatePropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("9262 Bramley Close"),
                City: unifiedgosdk.Pointer("New Andrewshire"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("53615-4664"),
                Region: unifiedgosdk.Pointer("Nebraska"),
                RegionCode: unifiedgosdk.Pointer("IL"),
            },
            Amenities: []string{
                "Wifi",
                "Kitchen",
                "Free parking",
            },
            Bathrooms: unifiedgosdk.Pointer[float64](1.0),
            Bedrooms: unifiedgosdk.Pointer[float64](4.0),
            Beds: unifiedgosdk.Pointer[float64](8.0),
            Channels: []shared.PropertyChannel{
                shared.PropertyChannel{
                    Name: unifiedgosdk.Pointer("Booking.com"),
                    WebURL: unifiedgosdk.Pointer("https://frightened-hyphenation.net/"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cuius deputo cruentus demoror censura caveo sponte statim vespillo cursim. Arbitro barba summopere. Curto vulticulus comparo surgo vociferor carmen deputo defluo."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("df923eb9-424a-43e0-a0f4-860f448a9683"),
            ImageUrls: []string{
                "https://loremflickr.com/23/1328?lock=5798241296161663",
                "https://loremflickr.com/12/1312?lock=6354590363728966",
            },
            Latitude: unifiedgosdk.Pointer[float64](-5.4327),
            Longitude: unifiedgosdk.Pointer[float64](98.9045),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("accurate Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T16:33:21.847Z"),
            WebURL: unifiedgosdk.Pointer("https://tight-pronoun.info/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreatePropertyListingRequest](../../pkg/models/operations/createpropertylistingrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreatePropertyListingResponse](../../pkg/models/operations/createpropertylistingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyListing

Retrieves a single PropertyListing by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyListing" method="get" path="/property/{connection_id}/listing/{id}" -->
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

    res, err := s.Listing.GetPropertyListing(ctx, operations.GetPropertyListingRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPropertyListingRequest](../../pkg/models/operations/getpropertylistingrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetPropertyListingResponse](../../pkg/models/operations/getpropertylistingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyListings

Returns a list of PropertyListings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyListings" method="get" path="/property/{connection_id}/listing" -->
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

    res, err := s.Listing.ListPropertyListings(ctx, operations.ListPropertyListingsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyListings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPropertyListingsRequest](../../pkg/models/operations/listpropertylistingsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListPropertyListingsResponse](../../pkg/models/operations/listpropertylistingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPropertyListing

Updates the PropertyListing identified by `id`. Provide the unified PropertyListing fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPropertyListing" method="patch" path="/property/{connection_id}/listing/{id}" example="property_listing" -->
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

    res, err := s.Listing.PatchPropertyListing(ctx, operations.PatchPropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("9262 Bramley Close"),
                City: unifiedgosdk.Pointer("New Andrewshire"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("53615-4664"),
                Region: unifiedgosdk.Pointer("Nebraska"),
                RegionCode: unifiedgosdk.Pointer("IL"),
            },
            Amenities: []string{
                "Wifi",
                "Kitchen",
                "Free parking",
            },
            Bathrooms: unifiedgosdk.Pointer[float64](1.0),
            Bedrooms: unifiedgosdk.Pointer[float64](4.0),
            Beds: unifiedgosdk.Pointer[float64](8.0),
            Channels: []shared.PropertyChannel{
                shared.PropertyChannel{
                    Name: unifiedgosdk.Pointer("Booking.com"),
                    WebURL: unifiedgosdk.Pointer("https://frightened-hyphenation.net/"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cuius deputo cruentus demoror censura caveo sponte statim vespillo cursim. Arbitro barba summopere. Curto vulticulus comparo surgo vociferor carmen deputo defluo."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("672337e0-a21e-4207-9c28-046e0ce56b57"),
            ImageUrls: []string{
                "https://loremflickr.com/23/1328?lock=5798241296161663",
                "https://loremflickr.com/12/1312?lock=6354590363728966",
            },
            Latitude: unifiedgosdk.Pointer[float64](-5.4327),
            Longitude: unifiedgosdk.Pointer[float64](98.9045),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("accurate Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T16:33:21.890Z"),
            WebURL: unifiedgosdk.Pointer("https://tight-pronoun.info/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchPropertyListingRequest](../../pkg/models/operations/patchpropertylistingrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchPropertyListingResponse](../../pkg/models/operations/patchpropertylistingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePropertyListing

Deletes the PropertyListing identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePropertyListing" method="delete" path="/property/{connection_id}/listing/{id}" -->
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

    res, err := s.Listing.RemovePropertyListing(ctx, operations.RemovePropertyListingRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemovePropertyListingRequest](../../pkg/models/operations/removepropertylistingrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemovePropertyListingResponse](../../pkg/models/operations/removepropertylistingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePropertyListing

Updates the PropertyListing identified by `id`. Provide the unified PropertyListing fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePropertyListing" method="put" path="/property/{connection_id}/listing/{id}" example="property_listing" -->
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

    res, err := s.Listing.UpdatePropertyListing(ctx, operations.UpdatePropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("9262 Bramley Close"),
                City: unifiedgosdk.Pointer("New Andrewshire"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("53615-4664"),
                Region: unifiedgosdk.Pointer("Nebraska"),
                RegionCode: unifiedgosdk.Pointer("IL"),
            },
            Amenities: []string{
                "Wifi",
                "Kitchen",
                "Free parking",
            },
            Bathrooms: unifiedgosdk.Pointer[float64](1.0),
            Bedrooms: unifiedgosdk.Pointer[float64](4.0),
            Beds: unifiedgosdk.Pointer[float64](8.0),
            Channels: []shared.PropertyChannel{
                shared.PropertyChannel{
                    Name: unifiedgosdk.Pointer("Booking.com"),
                    WebURL: unifiedgosdk.Pointer("https://frightened-hyphenation.net/"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cuius deputo cruentus demoror censura caveo sponte statim vespillo cursim. Arbitro barba summopere. Curto vulticulus comparo surgo vociferor carmen deputo defluo."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("672337e0-a21e-4207-9c28-046e0ce56b57"),
            ImageUrls: []string{
                "https://loremflickr.com/23/1328?lock=5798241296161663",
                "https://loremflickr.com/12/1312?lock=6354590363728966",
            },
            Latitude: unifiedgosdk.Pointer[float64](-5.4327),
            Longitude: unifiedgosdk.Pointer[float64](98.9045),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("accurate Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T16:33:21.890Z"),
            WebURL: unifiedgosdk.Pointer("https://tight-pronoun.info/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdatePropertyListingRequest](../../pkg/models/operations/updatepropertylistingrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdatePropertyListingResponse](../../pkg/models/operations/updatepropertylistingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |