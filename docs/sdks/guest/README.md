# Guest

## Overview

### Available Operations

* [CreatePropertyGuest](#createpropertyguest) - Create a guest
* [GetPropertyGuest](#getpropertyguest) - Retrieve a guest
* [ListPropertyGuests](#listpropertyguests) - List all guests
* [PatchPropertyGuest](#patchpropertyguest) - Update a guest
* [RemovePropertyGuest](#removepropertyguest) - Remove a guest
* [UpdatePropertyGuest](#updatepropertyguest) - Update a guest

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

    res, err := s.Guest.CreatePropertyGuest(ctx, operations.CreatePropertyGuestRequest{
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
            ID: unifiedgosdk.Pointer("18ef5e10-ce0b-4345-94a2-d9a47b126c30"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T22:16:52.037Z"),
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

    res, err := s.Guest.GetPropertyGuest(ctx, operations.GetPropertyGuestRequest{
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

    res, err := s.Guest.ListPropertyGuests(ctx, operations.ListPropertyGuestsRequest{
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

    res, err := s.Guest.PatchPropertyGuest(ctx, operations.PatchPropertyGuestRequest{
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
            ID: unifiedgosdk.Pointer("bcda34cc-aecd-4ce6-845c-58144de03581"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T22:16:52.047Z"),
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

    res, err := s.Guest.RemovePropertyGuest(ctx, operations.RemovePropertyGuestRequest{
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

    res, err := s.Guest.UpdatePropertyGuest(ctx, operations.UpdatePropertyGuestRequest{
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
            ID: unifiedgosdk.Pointer("bcda34cc-aecd-4ce6-845c-58144de03581"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-07-24T22:16:52.047Z"),
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