# Property

## Overview

### Available Operations

* [CreateAnalyticsProperty](#createanalyticsproperty) - Create a property
* [CreatePropertyConversation](#createpropertyconversation) - Create a conversation
* [CreatePropertyGuest](#createpropertyguest) - Create a guest
* [CreatePropertyListing](#createpropertylisting) - Create a listing
* [CreatePropertyQuote](#createpropertyquote) - Create a quote
* [CreatePropertyReservation](#createpropertyreservation) - Create a reservation
* [GetAnalyticsProperty](#getanalyticsproperty) - Retrieve a property
* [GetPropertyConversation](#getpropertyconversation) - Retrieve a conversation
* [GetPropertyGuest](#getpropertyguest) - Retrieve a guest
* [GetPropertyListing](#getpropertylisting) - Retrieve a listing
* [GetPropertyReservation](#getpropertyreservation) - Retrieve a reservation
* [GetPropertyReview](#getpropertyreview) - Retrieve a review
* [GetPropertyUnittype](#getpropertyunittype) - Retrieve an unittype
* [ListAnalyticsProperties](#listanalyticsproperties) - List all properties
* [ListPropertyAvailabilities](#listpropertyavailabilities) - List all availabilities
* [ListPropertyConversations](#listpropertyconversations) - List all conversations
* [ListPropertyGuests](#listpropertyguests) - List all guests
* [ListPropertyListings](#listpropertylistings) - List all listings
* [ListPropertyReservations](#listpropertyreservations) - List all reservations
* [ListPropertyReviews](#listpropertyreviews) - List all reviews
* [ListPropertyUnittypes](#listpropertyunittypes) - List all unittypes
* [PatchAnalyticsProperty](#patchanalyticsproperty) - Update a property
* [PatchPropertyGuest](#patchpropertyguest) - Update a guest
* [PatchPropertyListing](#patchpropertylisting) - Update a listing
* [PatchPropertyReservation](#patchpropertyreservation) - Update a reservation
* [PatchPropertyReview](#patchpropertyreview) - Update a review
* [RemoveAnalyticsProperty](#removeanalyticsproperty) - Remove a property
* [RemovePropertyGuest](#removepropertyguest) - Remove a guest
* [RemovePropertyListing](#removepropertylisting) - Remove a listing
* [RemovePropertyReservation](#removepropertyreservation) - Remove a reservation
* [UpdateAnalyticsProperty](#updateanalyticsproperty) - Update a property
* [UpdatePropertyGuest](#updatepropertyguest) - Update a guest
* [UpdatePropertyListing](#updatepropertylisting) - Update a listing
* [UpdatePropertyReservation](#updatepropertyreservation) - Update a reservation
* [UpdatePropertyReview](#updatepropertyreview) - Update a review

## CreateAnalyticsProperty

Creates a AnalyticsProperty. Provide the unified AnalyticsProperty fields in the request body; they are mapped to the provider's API. Returns the created AnalyticsProperty.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAnalyticsProperty" method="post" path="/analytics/{connection_id}/property" example="analytics_property" -->
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

    res, err := s.Property.CreateAnalyticsProperty(ctx, operations.CreateAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("f67b3f0d-235e-4e26-a4b0-96f210a5d040"),
            Name: unifiedgosdk.Pointer("Graham - Glover"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-15T00:44:17.431Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAnalyticsPropertyRequest](../../pkg/models/operations/createanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAnalyticsPropertyResponse](../../pkg/models/operations/createanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePropertyConversation

Creates a PropertyConversation. Provide the unified PropertyConversation fields in the request body; they are mapped to the provider's API. Returns the created PropertyConversation.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyConversation" method="post" path="/property/{connection_id}/conversation" example="property_conversation" -->
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

    res, err := s.Property.CreatePropertyConversation(ctx, operations.CreatePropertyConversationRequest{
        PropertyConversation: shared.PropertyConversation{
            CreatedAt: types.MustNewTimeFromString("2022-06-20T00:02:46.530Z"),
            GuestName: unifiedgosdk.Pointer("Maida Funk"),
            ID: unifiedgosdk.Pointer("b791b126-7a4a-4157-879d-f0b49f445a90"),
            IsUnread: unifiedgosdk.Pointer(false),
            Messages: []shared.PropertyMessage{
                shared.PropertyMessage{
                    Attachments: []shared.PropertyAttachment{},
                    CreatedAt: types.MustNewTimeFromString("2025-05-10T15:20:52.207Z"),
                    ID: unifiedgosdk.Pointer("b5381264-6d7a-4c9c-bb53-59814aa42f36"),
                    IsAutomatic: unifiedgosdk.Pointer(false),
                    Message: unifiedgosdk.Pointer("Sophismata admitto colo."),
                    SenderName: unifiedgosdk.Pointer("Maida Funk"),
                    SenderType: shared.SenderTypeGuest.ToPointer(),
                },
                shared.PropertyMessage{
                    Attachments: []shared.PropertyAttachment{},
                    CreatedAt: types.MustNewTimeFromString("2026-05-30T13:20:29.195Z"),
                    ID: unifiedgosdk.Pointer("d69313b5-f752-436e-b51e-b8cdf586813a"),
                    Message: unifiedgosdk.Pointer("Aufero torrens tonsor aequitas torqueo curiositas cimentarius repellendus auditor."),
                    SenderName: unifiedgosdk.Pointer("Alysson Sanford"),
                    SenderType: shared.SenderTypeHost.ToPointer(),
                },
            },
            Metadata: []shared.PropertyMetadata{},
            Source: unifiedgosdk.Pointer("Airbnb"),
            UpdatedAt: types.MustNewTimeFromString("2026-05-30T13:20:29.195Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreatePropertyConversationRequest](../../pkg/models/operations/createpropertyconversationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreatePropertyConversationResponse](../../pkg/models/operations/createpropertyconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePropertyGuest

Creates a PropertyGuest. Provide the unified PropertyGuest fields in the request body; they are mapped to the provider's API. Returns the created PropertyGuest.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyGuest" method="post" path="/property/{connection_id}/guest" example="property_guest" -->
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

    res, err := s.Property.CreatePropertyGuest(ctx, operations.CreatePropertyGuestRequest{
        PropertyGuest: shared.PropertyGuest{
            Address: &shared.PropertyPropertyGuestAddress{
                Address1: unifiedgosdk.Pointer("5781 Jules Rest"),
                City: unifiedgosdk.Pointer("Brandiborough"),
                Country: unifiedgosdk.Pointer("Uruguay"),
                CountryCode: unifiedgosdk.Pointer("IS"),
                PostalCode: unifiedgosdk.Pointer("29167"),
                Region: unifiedgosdk.Pointer("Mississippi"),
            },
            CompanyName: unifiedgosdk.Pointer("Hettinger - Lockman"),
            CreatedAt: types.MustNewTimeFromString("2021-10-17T05:59:05.070Z"),
            Emails: []shared.PropertyEmail{
                shared.PropertyEmail{
                    Email: "ed.pacocha@icloud.com",
                    Type: shared.PropertyEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ed"),
            ID: unifiedgosdk.Pointer("dc967abc-29e4-4dc8-be0d-23476de95d1b"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/48.jpg"),
            Language: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Pacocha"),
            Metadata: []shared.PropertyMetadata{},
            Name: unifiedgosdk.Pointer("Ed Pacocha"),
            Nationality: unifiedgosdk.Pointer("AW"),
            Notes: unifiedgosdk.Pointer("Officia cinis corrumpo conspergo delibero contra bos repellat."),
            Telephones: []shared.PropertyTelephone{
                shared.PropertyTelephone{
                    Telephone: "+14843285877",
                    Type: shared.PropertyTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T04:43:46.669Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyGuest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePropertyGuestRequest](../../pkg/models/operations/createpropertyguestrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePropertyGuestResponse](../../pkg/models/operations/createpropertyguestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Property.CreatePropertyListing(ctx, operations.CreatePropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("7424 1st Avenue"),
                City: unifiedgosdk.Pointer("Porterville"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("95594-3059"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("ID"),
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
                    WebURL: unifiedgosdk.Pointer("https://gentle-flat.name"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cruentus demoror censura caveo sponte statim. Cursim angulus arbitro barba summopere subnecto curto vulticulus comparo surgo. Carmen deputo defluo abscido conatus arbustum sequi abbas complectus demum."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("7d410b8b-9b57-410e-bcf0-061bbb1d3002"),
            ImageUrls: []string{
                "https://picsum.photos/seed/cCdY57x3Vw/2822/2061",
                "https://picsum.photos/seed/nH6zsBCr/1004/2259",
            },
            Latitude: unifiedgosdk.Pointer[float64](-73.4369),
            Longitude: unifiedgosdk.Pointer[float64](97.314),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("ironclad Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-04T04:17:55.256Z"),
            WebURL: unifiedgosdk.Pointer("https://oval-junior.info"),
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

## CreatePropertyQuote

Creates a PropertyQuote. Provide the unified PropertyQuote fields in the request body; they are mapped to the provider's API. Returns the created PropertyQuote.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyQuote" method="post" path="/property/{connection_id}/quote" example="property_quote" -->
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

    res, err := s.Property.CreatePropertyQuote(ctx, operations.CreatePropertyQuoteRequest{
        PropertyQuote: shared.PropertyQuote{
            Adults: unifiedgosdk.Pointer[float64](1.0),
            Children: unifiedgosdk.Pointer[float64](1.0),
            CreatedAt: types.MustNewTimeFromString("2020-12-18T19:14:11.415Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DepositAmount: unifiedgosdk.Pointer[float64](594.54),
            EndAt: types.MustNewTimeFromString("2026-12-11"),
            Guests: unifiedgosdk.Pointer[float64](2.0),
            ID: unifiedgosdk.Pointer("05e096b0-8b34-41c4-85e5-348019304f7d"),
            LineItems: []shared.PropertyLineitem{
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](1665.0),
                    Description: unifiedgosdk.Pointer("9 nights"),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Accommodation"),
                    Type: shared.PropertyLineitemTypeBase.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](117.0),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Cleaning fee"),
                    Type: shared.PropertyLineitemTypeFee.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](199.8),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lodging tax"),
                    Type: shared.PropertyLineitemTypeTax.ToPointer(),
                },
            },
            StartAt: types.MustNewTimeFromString("2026-12-02"),
            TotalAmount: unifiedgosdk.Pointer[float64](1981.8),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePropertyQuoteRequest](../../pkg/models/operations/createpropertyquoterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePropertyQuoteResponse](../../pkg/models/operations/createpropertyquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePropertyReservation

Creates a PropertyReservation. Provide the unified PropertyReservation fields in the request body; they are mapped to the provider's API. Returns the created PropertyReservation.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyReservation" method="post" path="/property/{connection_id}/reservation" example="property_reservation" -->
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

    res, err := s.Property.CreatePropertyReservation(ctx, operations.CreatePropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("TDDQ88ZYCY"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("mariam.metz@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Mariam Metz"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("123af2d2-4a1e-49f9-a249-f52def39d501"),
            LineItems: []shared.PropertyLineitem{
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](1980.0),
                    Description: unifiedgosdk.Pointer("11 nights"),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Accommodation"),
                    Type: shared.PropertyLineitemTypeBase.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](162.0),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Cleaning fee"),
                    Type: shared.PropertyLineitemTypeFee.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](237.6),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lodging tax"),
                    Type: shared.PropertyLineitemTypeTax.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](-198.0),
                    IsMandatory: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Weekly discount"),
                    Type: shared.PropertyLineitemTypeDiscount.ToPointer(),
                },
            },
            Metadata: []shared.PropertyMetadata{},
            Nights: unifiedgosdk.Pointer[float64](11.0),
            Notes: unifiedgosdk.Pointer("Volva quas trepide vae earum temptatio tertius dedico adfectus veniam."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-19T08:13:00.549Z"),
            WebURL: unifiedgosdk.Pointer("https://winged-lox.name/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreatePropertyReservationRequest](../../pkg/models/operations/createpropertyreservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreatePropertyReservationResponse](../../pkg/models/operations/createpropertyreservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAnalyticsProperty

Retrieves a single AnalyticsProperty by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsProperty" method="get" path="/analytics/{connection_id}/property/{id}" -->
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

    res, err := s.Property.GetAnalyticsProperty(ctx, operations.GetAnalyticsPropertyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAnalyticsPropertyRequest](../../pkg/models/operations/getanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAnalyticsPropertyResponse](../../pkg/models/operations/getanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyConversation

Retrieves a single PropertyConversation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyConversation" method="get" path="/property/{connection_id}/conversation/{id}" -->
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

    res, err := s.Property.GetPropertyConversation(ctx, operations.GetPropertyConversationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetPropertyConversationRequest](../../pkg/models/operations/getpropertyconversationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetPropertyConversationResponse](../../pkg/models/operations/getpropertyconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyGuest

Retrieves a single PropertyGuest by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyGuest" method="get" path="/property/{connection_id}/guest/{id}" -->
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

    res, err := s.Property.GetPropertyGuest(ctx, operations.GetPropertyGuestRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyGuest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPropertyGuestRequest](../../pkg/models/operations/getpropertyguestrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPropertyGuestResponse](../../pkg/models/operations/getpropertyguestresponse.md), error**

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

    res, err := s.Property.GetPropertyListing(ctx, operations.GetPropertyListingRequest{
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

## GetPropertyReservation

Retrieves a single PropertyReservation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyReservation" method="get" path="/property/{connection_id}/reservation/{id}" -->
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

    res, err := s.Property.GetPropertyReservation(ctx, operations.GetPropertyReservationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetPropertyReservationRequest](../../pkg/models/operations/getpropertyreservationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetPropertyReservationResponse](../../pkg/models/operations/getpropertyreservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyReview

Retrieves a single PropertyReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyReview" method="get" path="/property/{connection_id}/review/{id}" -->
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

    res, err := s.Property.GetPropertyReview(ctx, operations.GetPropertyReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPropertyReviewRequest](../../pkg/models/operations/getpropertyreviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetPropertyReviewResponse](../../pkg/models/operations/getpropertyreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyUnittype

Retrieves a single PropertyUnittype by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyUnittype" method="get" path="/property/{connection_id}/unittype/{id}" -->
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

    res, err := s.Property.GetPropertyUnittype(ctx, operations.GetPropertyUnittypeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyUnittype != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPropertyUnittypeRequest](../../pkg/models/operations/getpropertyunittyperequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPropertyUnittypeResponse](../../pkg/models/operations/getpropertyunittyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsProperties

Returns a list of AnalyticsProperties. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsProperties" method="get" path="/analytics/{connection_id}/property" -->
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

    res, err := s.Property.ListAnalyticsProperties(ctx, operations.ListAnalyticsPropertiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperties != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListAnalyticsPropertiesRequest](../../pkg/models/operations/listanalyticspropertiesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListAnalyticsPropertiesResponse](../../pkg/models/operations/listanalyticspropertiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyAvailabilities

Returns a list of PropertyAvailabilities. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyAvailabilities" method="get" path="/property/{connection_id}/availability" -->
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

    res, err := s.Property.ListPropertyAvailabilities(ctx, operations.ListPropertyAvailabilitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyAvailabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListPropertyAvailabilitiesRequest](../../pkg/models/operations/listpropertyavailabilitiesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListPropertyAvailabilitiesResponse](../../pkg/models/operations/listpropertyavailabilitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyConversations

Returns a list of PropertyConversations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyConversations" method="get" path="/property/{connection_id}/conversation" -->
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

    res, err := s.Property.ListPropertyConversations(ctx, operations.ListPropertyConversationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPropertyConversationsRequest](../../pkg/models/operations/listpropertyconversationsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListPropertyConversationsResponse](../../pkg/models/operations/listpropertyconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyGuests

Returns a list of PropertyGuests. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyGuests" method="get" path="/property/{connection_id}/guest" -->
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

    res, err := s.Property.ListPropertyGuests(ctx, operations.ListPropertyGuestsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyGuests != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPropertyGuestsRequest](../../pkg/models/operations/listpropertyguestsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPropertyGuestsResponse](../../pkg/models/operations/listpropertyguestsresponse.md), error**

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

    res, err := s.Property.ListPropertyListings(ctx, operations.ListPropertyListingsRequest{
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

## ListPropertyReservations

Returns a list of PropertyReservations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyReservations" method="get" path="/property/{connection_id}/reservation" -->
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

    res, err := s.Property.ListPropertyReservations(ctx, operations.ListPropertyReservationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReservations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListPropertyReservationsRequest](../../pkg/models/operations/listpropertyreservationsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListPropertyReservationsResponse](../../pkg/models/operations/listpropertyreservationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyReviews

Returns a list of PropertyReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyReviews" method="get" path="/property/{connection_id}/review" -->
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

    res, err := s.Property.ListPropertyReviews(ctx, operations.ListPropertyReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListPropertyReviewsRequest](../../pkg/models/operations/listpropertyreviewsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListPropertyReviewsResponse](../../pkg/models/operations/listpropertyreviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyUnittypes

Returns a list of PropertyUnittypes. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyUnittypes" method="get" path="/property/{connection_id}/unittype" -->
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

    res, err := s.Property.ListPropertyUnittypes(ctx, operations.ListPropertyUnittypesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyUnittypes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListPropertyUnittypesRequest](../../pkg/models/operations/listpropertyunittypesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListPropertyUnittypesResponse](../../pkg/models/operations/listpropertyunittypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAnalyticsProperty

Updates the AnalyticsProperty identified by `id`. Provide the unified AnalyticsProperty fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAnalyticsProperty" method="patch" path="/analytics/{connection_id}/property/{id}" example="analytics_property" -->
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

    res, err := s.Property.PatchAnalyticsProperty(ctx, operations.PatchAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("afbcafee-5bff-4869-a34b-242584f839f8"),
            Name: unifiedgosdk.Pointer("Graham - Glover"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-15T00:44:17.431Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAnalyticsPropertyRequest](../../pkg/models/operations/patchanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAnalyticsPropertyResponse](../../pkg/models/operations/patchanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPropertyGuest

Updates the PropertyGuest identified by `id`. Provide the unified PropertyGuest fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPropertyGuest" method="patch" path="/property/{connection_id}/guest/{id}" example="property_guest" -->
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

    res, err := s.Property.PatchPropertyGuest(ctx, operations.PatchPropertyGuestRequest{
        PropertyGuest: shared.PropertyGuest{
            Address: &shared.PropertyPropertyGuestAddress{
                Address1: unifiedgosdk.Pointer("5781 Jules Rest"),
                City: unifiedgosdk.Pointer("Brandiborough"),
                Country: unifiedgosdk.Pointer("Uruguay"),
                CountryCode: unifiedgosdk.Pointer("IS"),
                PostalCode: unifiedgosdk.Pointer("29167"),
                Region: unifiedgosdk.Pointer("Mississippi"),
            },
            CompanyName: unifiedgosdk.Pointer("Hettinger - Lockman"),
            CreatedAt: types.MustNewTimeFromString("2021-10-17T05:59:05.070Z"),
            Emails: []shared.PropertyEmail{
                shared.PropertyEmail{
                    Email: "ed.pacocha@icloud.com",
                    Type: shared.PropertyEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ed"),
            ID: unifiedgosdk.Pointer("81f0c4b7-a735-4d97-90a6-5cce60250f6c"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/48.jpg"),
            Language: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Pacocha"),
            Metadata: []shared.PropertyMetadata{},
            Name: unifiedgosdk.Pointer("Ed Pacocha"),
            Nationality: unifiedgosdk.Pointer("AW"),
            Notes: unifiedgosdk.Pointer("Officia cinis corrumpo conspergo delibero contra bos repellat."),
            Telephones: []shared.PropertyTelephone{
                shared.PropertyTelephone{
                    Telephone: "+14843285877",
                    Type: shared.PropertyTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T04:43:46.688Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyGuest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchPropertyGuestRequest](../../pkg/models/operations/patchpropertyguestrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchPropertyGuestResponse](../../pkg/models/operations/patchpropertyguestresponse.md), error**

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

    res, err := s.Property.PatchPropertyListing(ctx, operations.PatchPropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("7424 1st Avenue"),
                City: unifiedgosdk.Pointer("Porterville"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("95594-3059"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("ID"),
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
                    WebURL: unifiedgosdk.Pointer("https://gentle-flat.name"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cruentus demoror censura caveo sponte statim. Cursim angulus arbitro barba summopere subnecto curto vulticulus comparo surgo. Carmen deputo defluo abscido conatus arbustum sequi abbas complectus demum."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("9b01c2ae-b6cb-4c93-b48b-16497474727c"),
            ImageUrls: []string{
                "https://picsum.photos/seed/cCdY57x3Vw/2822/2061",
                "https://picsum.photos/seed/nH6zsBCr/1004/2259",
            },
            Latitude: unifiedgosdk.Pointer[float64](-73.4369),
            Longitude: unifiedgosdk.Pointer[float64](97.314),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("ironclad Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-04T04:17:55.290Z"),
            WebURL: unifiedgosdk.Pointer("https://oval-junior.info"),
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

## PatchPropertyReservation

Updates the PropertyReservation identified by `id`. Provide the unified PropertyReservation fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPropertyReservation" method="patch" path="/property/{connection_id}/reservation/{id}" example="property_reservation" -->
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

    res, err := s.Property.PatchPropertyReservation(ctx, operations.PatchPropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("TDDQ88ZYCY"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("mariam.metz@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Mariam Metz"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("3e1af200-3035-4318-99a4-a613a99382b7"),
            LineItems: []shared.PropertyLineitem{
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](1980.0),
                    Description: unifiedgosdk.Pointer("11 nights"),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Accommodation"),
                    Type: shared.PropertyLineitemTypeBase.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](162.0),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Cleaning fee"),
                    Type: shared.PropertyLineitemTypeFee.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](237.6),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lodging tax"),
                    Type: shared.PropertyLineitemTypeTax.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](-198.0),
                    IsMandatory: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Weekly discount"),
                    Type: shared.PropertyLineitemTypeDiscount.ToPointer(),
                },
            },
            Metadata: []shared.PropertyMetadata{},
            Nights: unifiedgosdk.Pointer[float64](11.0),
            Notes: unifiedgosdk.Pointer("Volva quas trepide vae earum temptatio tertius dedico adfectus veniam."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-19T08:13:00.551Z"),
            WebURL: unifiedgosdk.Pointer("https://winged-lox.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchPropertyReservationRequest](../../pkg/models/operations/patchpropertyreservationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchPropertyReservationResponse](../../pkg/models/operations/patchpropertyreservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPropertyReview

Updates the PropertyReview identified by `id`. Provide the unified PropertyReview fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPropertyReview" method="patch" path="/property/{connection_id}/review/{id}" example="property_review" -->
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

    res, err := s.Property.PatchPropertyReview(ctx, operations.PatchPropertyReviewRequest{
        PropertyReview: shared.PropertyReview{
            AuthorName: unifiedgosdk.Pointer("Hilma Volkman"),
            CategoryRatings: []shared.PropertyCategoryrating{
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCleanliness.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Sum blanditiis terreo adinventitias capitulus dolorem verus studio voluptates temperantia."),
                    Rating: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryAccuracy.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Avarus caelum curvo utor odit arcesso quisquam."),
                    Rating: unifiedgosdk.Pointer[float64](4.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCommunication.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCheckin.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryLocation.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](2.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryValue.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Iusto artificiose odit nemo sui."),
                    Rating: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Content: unifiedgosdk.Pointer("Attonbitus comitatus auditor. Tergo nostrum deleo vomica comprehendo tertius summopere. Apparatus incidunt tabula pauci confido colo fugiat umquam perferendis amitto."),
            CreatedAt: types.MustNewTimeFromString("2022-10-17T11:17:07.682Z"),
            ID: unifiedgosdk.Pointer("208376ca-6432-4013-96d4-fd7e5a361ec7"),
            IsPublic: unifiedgosdk.Pointer(true),
            Metadata: []shared.PropertyMetadata{},
            PrivateFeedback: unifiedgosdk.Pointer("Patria stella magni."),
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Source: unifiedgosdk.Pointer("Airbnb"),
            UpdatedAt: types.MustNewTimeFromString("2026-03-13T13:43:51.779Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchPropertyReviewRequest](../../pkg/models/operations/patchpropertyreviewrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchPropertyReviewResponse](../../pkg/models/operations/patchpropertyreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAnalyticsProperty

Deletes the AnalyticsProperty identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAnalyticsProperty" method="delete" path="/analytics/{connection_id}/property/{id}" -->
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

    res, err := s.Property.RemoveAnalyticsProperty(ctx, operations.RemoveAnalyticsPropertyRequest{
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
| `request`                                                                                                  | [operations.RemoveAnalyticsPropertyRequest](../../pkg/models/operations/removeanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAnalyticsPropertyResponse](../../pkg/models/operations/removeanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePropertyGuest

Deletes the PropertyGuest identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePropertyGuest" method="delete" path="/property/{connection_id}/guest/{id}" -->
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

    res, err := s.Property.RemovePropertyGuest(ctx, operations.RemovePropertyGuestRequest{
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
| `request`                                                                                          | [operations.RemovePropertyGuestRequest](../../pkg/models/operations/removepropertyguestrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemovePropertyGuestResponse](../../pkg/models/operations/removepropertyguestresponse.md), error**

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

    res, err := s.Property.RemovePropertyListing(ctx, operations.RemovePropertyListingRequest{
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

## RemovePropertyReservation

Deletes the PropertyReservation identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePropertyReservation" method="delete" path="/property/{connection_id}/reservation/{id}" -->
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

    res, err := s.Property.RemovePropertyReservation(ctx, operations.RemovePropertyReservationRequest{
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
| `request`                                                                                                      | [operations.RemovePropertyReservationRequest](../../pkg/models/operations/removepropertyreservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemovePropertyReservationResponse](../../pkg/models/operations/removepropertyreservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAnalyticsProperty

Updates the AnalyticsProperty identified by `id`. Provide the unified AnalyticsProperty fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAnalyticsProperty" method="put" path="/analytics/{connection_id}/property/{id}" example="analytics_property" -->
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

    res, err := s.Property.UpdateAnalyticsProperty(ctx, operations.UpdateAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("afbcafee-5bff-4869-a34b-242584f839f8"),
            Name: unifiedgosdk.Pointer("Graham - Glover"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-15T00:44:17.431Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAnalyticsPropertyRequest](../../pkg/models/operations/updateanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAnalyticsPropertyResponse](../../pkg/models/operations/updateanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePropertyGuest

Updates the PropertyGuest identified by `id`. Provide the unified PropertyGuest fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePropertyGuest" method="put" path="/property/{connection_id}/guest/{id}" example="property_guest" -->
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

    res, err := s.Property.UpdatePropertyGuest(ctx, operations.UpdatePropertyGuestRequest{
        PropertyGuest: shared.PropertyGuest{
            Address: &shared.PropertyPropertyGuestAddress{
                Address1: unifiedgosdk.Pointer("5781 Jules Rest"),
                City: unifiedgosdk.Pointer("Brandiborough"),
                Country: unifiedgosdk.Pointer("Uruguay"),
                CountryCode: unifiedgosdk.Pointer("IS"),
                PostalCode: unifiedgosdk.Pointer("29167"),
                Region: unifiedgosdk.Pointer("Mississippi"),
            },
            CompanyName: unifiedgosdk.Pointer("Hettinger - Lockman"),
            CreatedAt: types.MustNewTimeFromString("2021-10-17T05:59:05.070Z"),
            Emails: []shared.PropertyEmail{
                shared.PropertyEmail{
                    Email: "ed.pacocha@icloud.com",
                    Type: shared.PropertyEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ed"),
            ID: unifiedgosdk.Pointer("81f0c4b7-a735-4d97-90a6-5cce60250f6c"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/48.jpg"),
            Language: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Pacocha"),
            Metadata: []shared.PropertyMetadata{},
            Name: unifiedgosdk.Pointer("Ed Pacocha"),
            Nationality: unifiedgosdk.Pointer("AW"),
            Notes: unifiedgosdk.Pointer("Officia cinis corrumpo conspergo delibero contra bos repellat."),
            Telephones: []shared.PropertyTelephone{
                shared.PropertyTelephone{
                    Telephone: "+14843285877",
                    Type: shared.PropertyTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T04:43:46.688Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyGuest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdatePropertyGuestRequest](../../pkg/models/operations/updatepropertyguestrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdatePropertyGuestResponse](../../pkg/models/operations/updatepropertyguestresponse.md), error**

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

    res, err := s.Property.UpdatePropertyListing(ctx, operations.UpdatePropertyListingRequest{
        PropertyListing: shared.PropertyListing{
            Address: &shared.PropertyPropertyListingAddress{
                Address1: unifiedgosdk.Pointer("7424 1st Avenue"),
                City: unifiedgosdk.Pointer("Porterville"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("95594-3059"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("ID"),
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
                    WebURL: unifiedgosdk.Pointer("https://gentle-flat.name"),
                },
            },
            CheckInEndTime: unifiedgosdk.Pointer("22:00"),
            CheckInStartTime: unifiedgosdk.Pointer("15:00"),
            CheckOutTime: unifiedgosdk.Pointer("11:00"),
            CleaningFeeAmount: unifiedgosdk.Pointer[float64](134.0),
            CreatedAt: types.MustNewTimeFromString("2021-01-08T09:27:35.485Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cruentus demoror censura caveo sponte statim. Cursim angulus arbitro barba summopere subnecto curto vulticulus comparo surgo. Carmen deputo defluo abscido conatus arbustum sequi abbas complectus demum."),
            HouseRules: unifiedgosdk.Pointer("No smoking. No parties or events."),
            ID: unifiedgosdk.Pointer("9b01c2ae-b6cb-4c93-b48b-16497474727c"),
            ImageUrls: []string{
                "https://picsum.photos/seed/cCdY57x3Vw/2822/2061",
                "https://picsum.photos/seed/nH6zsBCr/1004/2259",
            },
            Latitude: unifiedgosdk.Pointer[float64](-73.4369),
            Longitude: unifiedgosdk.Pointer[float64](97.314),
            MaxNights: unifiedgosdk.Pointer[float64](15.0),
            MaxOccupancy: unifiedgosdk.Pointer[float64](5.0),
            Metadata: []shared.PropertyMetadata{},
            MinNights: unifiedgosdk.Pointer[float64](1.0),
            Name: unifiedgosdk.Pointer("ironclad Loft"),
            Rating: unifiedgosdk.Pointer[float64](4.9),
            ReviewCount: unifiedgosdk.Pointer[float64](343.0),
            StartPriceAmount: unifiedgosdk.Pointer[float64](470.0),
            Timezone: unifiedgosdk.Pointer("America/Los_Angeles"),
            Type: shared.PropertyListingTypeRoom.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-04T04:17:55.290Z"),
            WebURL: unifiedgosdk.Pointer("https://oval-junior.info"),
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

## UpdatePropertyReservation

Updates the PropertyReservation identified by `id`. Provide the unified PropertyReservation fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePropertyReservation" method="put" path="/property/{connection_id}/reservation/{id}" example="property_reservation" -->
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

    res, err := s.Property.UpdatePropertyReservation(ctx, operations.UpdatePropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("TDDQ88ZYCY"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("mariam.metz@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Mariam Metz"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("3e1af200-3035-4318-99a4-a613a99382b7"),
            LineItems: []shared.PropertyLineitem{
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](1980.0),
                    Description: unifiedgosdk.Pointer("11 nights"),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Accommodation"),
                    Type: shared.PropertyLineitemTypeBase.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](162.0),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Cleaning fee"),
                    Type: shared.PropertyLineitemTypeFee.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](237.6),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lodging tax"),
                    Type: shared.PropertyLineitemTypeTax.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](-198.0),
                    IsMandatory: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Weekly discount"),
                    Type: shared.PropertyLineitemTypeDiscount.ToPointer(),
                },
            },
            Metadata: []shared.PropertyMetadata{},
            Nights: unifiedgosdk.Pointer[float64](11.0),
            Notes: unifiedgosdk.Pointer("Volva quas trepide vae earum temptatio tertius dedico adfectus veniam."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-19T08:13:00.551Z"),
            WebURL: unifiedgosdk.Pointer("https://winged-lox.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdatePropertyReservationRequest](../../pkg/models/operations/updatepropertyreservationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdatePropertyReservationResponse](../../pkg/models/operations/updatepropertyreservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePropertyReview

Updates the PropertyReview identified by `id`. Provide the unified PropertyReview fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePropertyReview" method="put" path="/property/{connection_id}/review/{id}" example="property_review" -->
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

    res, err := s.Property.UpdatePropertyReview(ctx, operations.UpdatePropertyReviewRequest{
        PropertyReview: shared.PropertyReview{
            AuthorName: unifiedgosdk.Pointer("Hilma Volkman"),
            CategoryRatings: []shared.PropertyCategoryrating{
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCleanliness.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Sum blanditiis terreo adinventitias capitulus dolorem verus studio voluptates temperantia."),
                    Rating: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryAccuracy.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Avarus caelum curvo utor odit arcesso quisquam."),
                    Rating: unifiedgosdk.Pointer[float64](4.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCommunication.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryCheckin.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](1.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryLocation.ToPointer(),
                    Rating: unifiedgosdk.Pointer[float64](2.0),
                },
                shared.PropertyCategoryrating{
                    Category: shared.PropertyCategoryratingCategoryValue.ToPointer(),
                    Comment: unifiedgosdk.Pointer("Iusto artificiose odit nemo sui."),
                    Rating: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Content: unifiedgosdk.Pointer("Attonbitus comitatus auditor. Tergo nostrum deleo vomica comprehendo tertius summopere. Apparatus incidunt tabula pauci confido colo fugiat umquam perferendis amitto."),
            CreatedAt: types.MustNewTimeFromString("2022-10-17T11:17:07.682Z"),
            ID: unifiedgosdk.Pointer("208376ca-6432-4013-96d4-fd7e5a361ec7"),
            IsPublic: unifiedgosdk.Pointer(true),
            Metadata: []shared.PropertyMetadata{},
            PrivateFeedback: unifiedgosdk.Pointer("Patria stella magni."),
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Source: unifiedgosdk.Pointer("Airbnb"),
            UpdatedAt: types.MustNewTimeFromString("2026-03-13T13:43:51.779Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdatePropertyReviewRequest](../../pkg/models/operations/updatepropertyreviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdatePropertyReviewResponse](../../pkg/models/operations/updatepropertyreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |