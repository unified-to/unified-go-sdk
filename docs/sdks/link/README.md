# Link

## Overview

### Available Operations

* [CreateCalendarLink](#createcalendarlink) - Create a link
* [CreatePaymentLink](#createpaymentlink) - Create a link
* [GetCalendarLink](#getcalendarlink) - Retrieve a link
* [GetPaymentLink](#getpaymentlink) - Retrieve a link
* [ListCalendarLinks](#listcalendarlinks) - List all links
* [ListPaymentLinks](#listpaymentlinks) - List all links
* [PatchCalendarLink](#patchcalendarlink) - Update a link
* [PatchPaymentLink](#patchpaymentlink) - Update a link
* [RemoveCalendarLink](#removecalendarlink) - Remove a link
* [RemovePaymentLink](#removepaymentlink) - Remove a link
* [UpdateCalendarLink](#updatecalendarlink) - Update a link
* [UpdatePaymentLink](#updatepaymentlink) - Update a link

## CreateCalendarLink

Creates a CalendarLink. Provide the unified CalendarLink fields in the request body; they are mapped to the provider's API. Returns the created CalendarLink.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCalendarLink" method="post" path="/calendar/{connection_id}/link" example="calendar_link" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Link.CreateCalendarLink(ctx, operations.CreateCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("ee13c8af-b6c1-442e-9d3e-f6c6bc4af431"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-18T00:51:14.008Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCalendarLinkRequest](../../pkg/models/operations/createcalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateCalendarLinkResponse](../../pkg/models/operations/createcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePaymentLink

Creates a PaymentLink. Provide the unified PaymentLink fields in the request body; they are mapped to the provider's API. Returns the created PaymentLink.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPaymentLink" method="post" path="/payment/{connection_id}/link" example="payment_link" -->
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

    res, err := s.Link.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](17001.68),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Savor the tender essence in our Ball, designed for dim culinary adventures"),
            ID: unifiedgosdk.Pointer("e9a5e629-84e3-4982-af82-6724e85c06e4"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("8e40cc80-da54-4cc2-b068-362d334c9ad7"),
                    ItemDescription: unifiedgosdk.Pointer("Discover the worst new Computer with an exciting mix of Steel ingredients"),
                    ItemName: unifiedgosdk.Pointer("Soft Marble Cheese"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](9342.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](519.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](18.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](424.32),
                    ID: unifiedgosdk.Pointer("b7f13fe4-bc49-4066-bb3d-be254af19f44"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Shoes featuring utilized technology and Cotton construction"),
                    ItemName: unifiedgosdk.Pointer("Sleek Aluminum Shoes"),
                    Notes: unifiedgosdk.Pointer("Cervus decretum viridis doloremque."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](6647.68),
                    UnitAmount: unifiedgosdk.Pointer[float64](884.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("b5d84072-21d7-43cb-8dd6-39dce61ace5c"),
                    ItemDescription: unifiedgosdk.Pointer("Featuring Osmium-enhanced technology, our Soap offers unparalleled major performance"),
                    ItemName: unifiedgosdk.Pointer("Elegant Concrete Gloves"),
                    Notes: unifiedgosdk.Pointer("Desipio usque vacuus currus cetera testimonium."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1012.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](506.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](2.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
            },
            Metadata: []shared.PaymentMetadata{},
            SuccessURL: unifiedgosdk.Pointer("https://tinted-jet.info"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-15T15:19:47.034Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/e9a5e629-84e3-4982-af82-6724e85c06e4"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreatePaymentLinkRequest](../../pkg/models/operations/createpaymentlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreatePaymentLinkResponse](../../pkg/models/operations/createpaymentlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarLink

Retrieves a single CalendarLink by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarLink" method="get" path="/calendar/{connection_id}/link/{id}" -->
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

    res, err := s.Link.GetCalendarLink(ctx, operations.GetCalendarLinkRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCalendarLinkRequest](../../pkg/models/operations/getcalendarlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetCalendarLinkResponse](../../pkg/models/operations/getcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPaymentLink

Retrieves a single PaymentLink by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPaymentLink" method="get" path="/payment/{connection_id}/link/{id}" -->
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

    res, err := s.Link.GetPaymentLink(ctx, operations.GetPaymentLinkRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPaymentLinkRequest](../../pkg/models/operations/getpaymentlinkrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetPaymentLinkResponse](../../pkg/models/operations/getpaymentlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarLinks

Returns a list of CalendarLinks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarLinks" method="get" path="/calendar/{connection_id}/link" -->
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

    res, err := s.Link.ListCalendarLinks(ctx, operations.ListCalendarLinksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLinks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCalendarLinksRequest](../../pkg/models/operations/listcalendarlinksrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCalendarLinksResponse](../../pkg/models/operations/listcalendarlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPaymentLinks

Returns a list of PaymentLinks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPaymentLinks" method="get" path="/payment/{connection_id}/link" -->
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

    res, err := s.Link.ListPaymentLinks(ctx, operations.ListPaymentLinksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLinks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPaymentLinksRequest](../../pkg/models/operations/listpaymentlinksrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListPaymentLinksResponse](../../pkg/models/operations/listpaymentlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCalendarLink

Updates the CalendarLink identified by `id`. Provide the unified CalendarLink fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCalendarLink" method="patch" path="/calendar/{connection_id}/link/{id}" example="calendar_link" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Link.PatchCalendarLink(ctx, operations.PatchCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("5cf2365c-4558-4068-93b3-adb8af8ace14"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-18T00:51:14.013Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchCalendarLinkRequest](../../pkg/models/operations/patchcalendarlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchCalendarLinkResponse](../../pkg/models/operations/patchcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPaymentLink

Updates the PaymentLink identified by `id`. Provide the unified PaymentLink fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPaymentLink" method="patch" path="/payment/{connection_id}/link/{id}" example="payment_link" -->
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

    res, err := s.Link.PatchPaymentLink(ctx, operations.PatchPaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](17001.68),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Savor the tender essence in our Ball, designed for dim culinary adventures"),
            ID: unifiedgosdk.Pointer("7e61fbd0-ccef-4f6f-871a-c11df0f1060e"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("4698b3b3-eb21-4675-b700-a233dd98bc4e"),
                    ItemDescription: unifiedgosdk.Pointer("Discover the worst new Computer with an exciting mix of Steel ingredients"),
                    ItemName: unifiedgosdk.Pointer("Soft Marble Cheese"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](9342.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](519.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](18.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](424.32),
                    ID: unifiedgosdk.Pointer("0778b896-34bc-44d2-8386-5236099bbdef"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Shoes featuring utilized technology and Cotton construction"),
                    ItemName: unifiedgosdk.Pointer("Sleek Aluminum Shoes"),
                    Notes: unifiedgosdk.Pointer("Cervus decretum viridis doloremque."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](6647.68),
                    UnitAmount: unifiedgosdk.Pointer[float64](884.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("da5456d2-7500-4c31-b643-b0764026a4c3"),
                    ItemDescription: unifiedgosdk.Pointer("Featuring Osmium-enhanced technology, our Soap offers unparalleled major performance"),
                    ItemName: unifiedgosdk.Pointer("Elegant Concrete Gloves"),
                    Notes: unifiedgosdk.Pointer("Desipio usque vacuus currus cetera testimonium."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1012.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](506.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](2.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
            },
            Metadata: []shared.PaymentMetadata{},
            SuccessURL: unifiedgosdk.Pointer("https://tinted-jet.info"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-15T15:19:47.042Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/7e61fbd0-ccef-4f6f-871a-c11df0f1060e"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchPaymentLinkRequest](../../pkg/models/operations/patchpaymentlinkrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchPaymentLinkResponse](../../pkg/models/operations/patchpaymentlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCalendarLink

Deletes the CalendarLink identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCalendarLink" method="delete" path="/calendar/{connection_id}/link/{id}" -->
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

    res, err := s.Link.RemoveCalendarLink(ctx, operations.RemoveCalendarLinkRequest{
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
| `request`                                                                                        | [operations.RemoveCalendarLinkRequest](../../pkg/models/operations/removecalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveCalendarLinkResponse](../../pkg/models/operations/removecalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePaymentLink

Deletes the PaymentLink identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePaymentLink" method="delete" path="/payment/{connection_id}/link/{id}" -->
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

    res, err := s.Link.RemovePaymentLink(ctx, operations.RemovePaymentLinkRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemovePaymentLinkRequest](../../pkg/models/operations/removepaymentlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemovePaymentLinkResponse](../../pkg/models/operations/removepaymentlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCalendarLink

Updates the CalendarLink identified by `id`. Provide the unified CalendarLink fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCalendarLink" method="put" path="/calendar/{connection_id}/link/{id}" example="calendar_link" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Link.UpdateCalendarLink(ctx, operations.UpdateCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("5cf2365c-4558-4068-93b3-adb8af8ace14"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-18T00:51:14.013Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateCalendarLinkRequest](../../pkg/models/operations/updatecalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateCalendarLinkResponse](../../pkg/models/operations/updatecalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePaymentLink

Updates the PaymentLink identified by `id`. Provide the unified PaymentLink fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePaymentLink" method="put" path="/payment/{connection_id}/link/{id}" example="payment_link" -->
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

    res, err := s.Link.UpdatePaymentLink(ctx, operations.UpdatePaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](17001.68),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("CAD"),
            Description: unifiedgosdk.Pointer("Savor the tender essence in our Ball, designed for dim culinary adventures"),
            ID: unifiedgosdk.Pointer("7e61fbd0-ccef-4f6f-871a-c11df0f1060e"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("4698b3b3-eb21-4675-b700-a233dd98bc4e"),
                    ItemDescription: unifiedgosdk.Pointer("Discover the worst new Computer with an exciting mix of Steel ingredients"),
                    ItemName: unifiedgosdk.Pointer("Soft Marble Cheese"),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](9342.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](519.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](18.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](424.32),
                    ID: unifiedgosdk.Pointer("0778b896-34bc-44d2-8386-5236099bbdef"),
                    ItemDescription: unifiedgosdk.Pointer("Innovative Shoes featuring utilized technology and Cotton construction"),
                    ItemName: unifiedgosdk.Pointer("Sleek Aluminum Shoes"),
                    Notes: unifiedgosdk.Pointer("Cervus decretum viridis doloremque."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](6647.68),
                    UnitAmount: unifiedgosdk.Pointer[float64](884.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("da5456d2-7500-4c31-b643-b0764026a4c3"),
                    ItemDescription: unifiedgosdk.Pointer("Featuring Osmium-enhanced technology, our Soap offers unparalleled major performance"),
                    ItemName: unifiedgosdk.Pointer("Elegant Concrete Gloves"),
                    Notes: unifiedgosdk.Pointer("Desipio usque vacuus currus cetera testimonium."),
                    TaxAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1012.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](506.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](2.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                },
            },
            Metadata: []shared.PaymentMetadata{},
            SuccessURL: unifiedgosdk.Pointer("https://tinted-jet.info"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-15T15:19:47.042Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/7e61fbd0-ccef-4f6f-871a-c11df0f1060e"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePaymentLinkRequest](../../pkg/models/operations/updatepaymentlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdatePaymentLinkResponse](../../pkg/models/operations/updatepaymentlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |