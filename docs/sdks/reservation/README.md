# Reservation

## Overview

### Available Operations

* [CreateCommerceReservation](#createcommercereservation) - Create a reservation
* [CreatePropertyReservation](#createpropertyreservation) - Create a reservation
* [GetCommerceReservation](#getcommercereservation) - Retrieve a reservation
* [GetPropertyReservation](#getpropertyreservation) - Retrieve a reservation
* [ListCommerceReservations](#listcommercereservations) - List all reservations
* [ListPropertyReservations](#listpropertyreservations) - List all reservations
* [PatchCommerceReservation](#patchcommercereservation) - Update a reservation
* [PatchPropertyReservation](#patchpropertyreservation) - Update a reservation
* [RemoveCommerceReservation](#removecommercereservation) - Remove a reservation
* [UpdateCommerceReservation](#updatecommercereservation) - Update a reservation
* [UpdatePropertyReservation](#updatepropertyreservation) - Update a reservation

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

    res, err := s.Reservation.CreateCommerceReservation(ctx, operations.CreateCommerceReservationRequest{
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

    res, err := s.Reservation.CreatePropertyReservation(ctx, operations.CreatePropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("GNLTDDQ88Z"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("anya.parisian@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Anya Parisian"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("b212f0be-b6b2-4416-9c6e-fe23de46873d"),
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
            Notes: unifiedgosdk.Pointer("Conculco volaticus vis volva quas trepide vae earum temptatio tertius."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-16T10:17:00.912Z"),
            WebURL: unifiedgosdk.Pointer("https://antique-tennis.biz/"),
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

    res, err := s.Reservation.GetCommerceReservation(ctx, operations.GetCommerceReservationRequest{
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

    res, err := s.Reservation.GetPropertyReservation(ctx, operations.GetPropertyReservationRequest{
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

    res, err := s.Reservation.ListCommerceReservations(ctx, operations.ListCommerceReservationsRequest{
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

    res, err := s.Reservation.ListPropertyReservations(ctx, operations.ListPropertyReservationsRequest{
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

    res, err := s.Reservation.PatchCommerceReservation(ctx, operations.PatchCommerceReservationRequest{
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

    res, err := s.Reservation.PatchPropertyReservation(ctx, operations.PatchPropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("GNLTDDQ88Z"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("anya.parisian@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Anya Parisian"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("eac8c537-6977-4e1b-b1dc-48ed8ecf73db"),
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
            Notes: unifiedgosdk.Pointer("Conculco volaticus vis volva quas trepide vae earum temptatio tertius."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-16T10:17:00.915Z"),
            WebURL: unifiedgosdk.Pointer("https://antique-tennis.biz/"),
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

    res, err := s.Reservation.RemoveCommerceReservation(ctx, operations.RemoveCommerceReservationRequest{
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

    res, err := s.Reservation.UpdateCommerceReservation(ctx, operations.UpdateCommerceReservationRequest{
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

    res, err := s.Reservation.UpdatePropertyReservation(ctx, operations.UpdatePropertyReservationRequest{
        PropertyReservation: shared.PropertyReservation{
            Adults: unifiedgosdk.Pointer[float64](4.0),
            BalanceAmount: unifiedgosdk.Pointer[float64](2181.6),
            Children: unifiedgosdk.Pointer[float64](2.0),
            ConfirmationCode: unifiedgosdk.Pointer("GNLTDDQ88Z"),
            CreatedAt: types.MustNewTimeFromString("2022-06-03T01:40:03.655Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EndAt: types.MustNewTimeFromString("2022-10-06T17:17:08.838Z"),
            EstimatedDepartureTime: unifiedgosdk.Pointer("10:00"),
            GuestEmail: unifiedgosdk.Pointer("anya.parisian@hotmail.com"),
            GuestName: unifiedgosdk.Pointer("Anya Parisian"),
            Guests: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("eac8c537-6977-4e1b-b1dc-48ed8ecf73db"),
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
            Notes: unifiedgosdk.Pointer("Conculco volaticus vis volva quas trepide vae earum temptatio tertius."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            Pets: unifiedgosdk.Pointer[float64](1.0),
            Source: unifiedgosdk.Pointer("Vrbo"),
            StartAt: types.MustNewTimeFromString("2022-09-25T17:17:08.838Z"),
            Status: shared.PropertyReservationStatusPending.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](2181.6),
            UpdatedAt: types.MustNewTimeFromString("2022-11-16T10:17:00.915Z"),
            WebURL: unifiedgosdk.Pointer("https://antique-tennis.biz/"),
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