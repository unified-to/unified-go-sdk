# Location

## Overview

### Available Operations

* [CreateCommerceLocation](#createcommercelocation) - Create a location
* [CreateHrisLocation](#createhrislocation) - Create a location
* [GetClubsLocation](#getclubslocation) - Retrieve a location
* [GetCommerceLocation](#getcommercelocation) - Retrieve a location
* [GetHealthcareLocation](#gethealthcarelocation) - Retrieve a location
* [GetHrisLocation](#gethrislocation) - Retrieve a location
* [ListClubsLocations](#listclubslocations) - List all locations
* [ListCommerceLocations](#listcommercelocations) - List all locations
* [ListHealthcareLocations](#listhealthcarelocations) - List all locations
* [ListHrisLocations](#listhrislocations) - List all locations
* [PatchCommerceLocation](#patchcommercelocation) - Update a location
* [PatchHrisLocation](#patchhrislocation) - Update a location
* [RemoveCommerceLocation](#removecommercelocation) - Remove a location
* [RemoveHrisLocation](#removehrislocation) - Remove a location
* [UpdateCommerceLocation](#updatecommercelocation) - Update a location
* [UpdateHrisLocation](#updatehrislocation) - Update a location

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

    res, err := s.Location.CreateCommerceLocation(ctx, operations.CreateCommerceLocationRequest{
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
            ID: unifiedgosdk.Pointer("733600a7-db10-41ff-a05f-4b0cf2a8f43c"),
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
                    ID: unifiedgosdk.Pointer("224dd4c7-28d3-47a6-8379-f8b30e0eaaac"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("3610d93e-217c-4bef-bb56-363de034091f"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-19T17:29:30.242Z"),
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

## CreateHrisLocation

Create a location

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisLocation" method="post" path="/hris/{connection_id}/location" example="hris_location" -->
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

    res, err := s.Location.CreateHrisLocation(ctx, operations.CreateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("20cf8cd4-22a0-401c-b367-9abe58246843"),
            ID: unifiedgosdk.Pointer("a6293150-f21c-4282-ae66-2e0d19cb8a29"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-20T00:26:16.853Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisLocationRequest](../../pkg/models/operations/createhrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisLocationResponse](../../pkg/models/operations/createhrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetClubsLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getClubsLocation" method="get" path="/clubs/{connection_id}/location/{id}" -->
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

    res, err := s.Location.GetClubsLocation(ctx, operations.GetClubsLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetClubsLocationRequest](../../pkg/models/operations/getclubslocationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetClubsLocationResponse](../../pkg/models/operations/getclubslocationresponse.md), error**

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

    res, err := s.Location.GetCommerceLocation(ctx, operations.GetCommerceLocationRequest{
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

## GetHealthcareLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareLocation" method="get" path="/healthcare/{connection_id}/location/{id}" -->
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

    res, err := s.Location.GetHealthcareLocation(ctx, operations.GetHealthcareLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetHealthcareLocationRequest](../../pkg/models/operations/gethealthcarelocationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetHealthcareLocationResponse](../../pkg/models/operations/gethealthcarelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisLocation" method="get" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Location.GetHrisLocation(ctx, operations.GetHrisLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisLocationRequest](../../pkg/models/operations/gethrislocationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisLocationResponse](../../pkg/models/operations/gethrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListClubsLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listClubsLocations" method="get" path="/clubs/{connection_id}/location" -->
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

    res, err := s.Location.ListClubsLocations(ctx, operations.ListClubsLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListClubsLocationsRequest](../../pkg/models/operations/listclubslocationsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListClubsLocationsResponse](../../pkg/models/operations/listclubslocationsresponse.md), error**

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

    res, err := s.Location.ListCommerceLocations(ctx, operations.ListCommerceLocationsRequest{
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

## ListHealthcareLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareLocations" method="get" path="/healthcare/{connection_id}/location" -->
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

    res, err := s.Location.ListHealthcareLocations(ctx, operations.ListHealthcareLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListHealthcareLocationsRequest](../../pkg/models/operations/listhealthcarelocationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListHealthcareLocationsResponse](../../pkg/models/operations/listhealthcarelocationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisLocations" method="get" path="/hris/{connection_id}/location" -->
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

    res, err := s.Location.ListHrisLocations(ctx, operations.ListHrisLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisLocationsRequest](../../pkg/models/operations/listhrislocationsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisLocationsResponse](../../pkg/models/operations/listhrislocationsresponse.md), error**

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

    res, err := s.Location.PatchCommerceLocation(ctx, operations.PatchCommerceLocationRequest{
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
            ID: unifiedgosdk.Pointer("5caf4413-1482-41d7-8044-f72db9e903fe"),
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
                    ID: unifiedgosdk.Pointer("023f7c5e-81db-4474-9536-31e4bd405aba"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("7814e107-2ece-444d-a58e-d77f89d7b745"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-19T17:29:30.255Z"),
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

## PatchHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisLocation" method="patch" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Location.PatchHrisLocation(ctx, operations.PatchHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("52e72984-3d41-4798-aeed-b0e94df76ef1"),
            ID: unifiedgosdk.Pointer("e5aaf781-7bb5-443b-becb-d90f8130b96b"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-20T00:26:16.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisLocationRequest](../../pkg/models/operations/patchhrislocationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisLocationResponse](../../pkg/models/operations/patchhrislocationresponse.md), error**

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

    res, err := s.Location.RemoveCommerceLocation(ctx, operations.RemoveCommerceLocationRequest{
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

## RemoveHrisLocation

Remove a location

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisLocation" method="delete" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Location.RemoveHrisLocation(ctx, operations.RemoveHrisLocationRequest{
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
| `request`                                                                                        | [operations.RemoveHrisLocationRequest](../../pkg/models/operations/removehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisLocationResponse](../../pkg/models/operations/removehrislocationresponse.md), error**

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

    res, err := s.Location.UpdateCommerceLocation(ctx, operations.UpdateCommerceLocationRequest{
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
            ID: unifiedgosdk.Pointer("5caf4413-1482-41d7-8044-f72db9e903fe"),
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
                    ID: unifiedgosdk.Pointer("023f7c5e-81db-4474-9536-31e4bd405aba"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("7814e107-2ece-444d-a58e-d77f89d7b745"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-19T17:29:30.255Z"),
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

## UpdateHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisLocation" method="put" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Location.UpdateHrisLocation(ctx, operations.UpdateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("52e72984-3d41-4798-aeed-b0e94df76ef1"),
            ID: unifiedgosdk.Pointer("e5aaf781-7bb5-443b-becb-d90f8130b96b"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-20T00:26:16.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisLocationRequest](../../pkg/models/operations/updatehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisLocationResponse](../../pkg/models/operations/updatehrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |