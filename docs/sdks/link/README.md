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

Create a link

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
            ID: unifiedgosdk.Pointer("ac5a258c-077d-4fdd-988e-bad749ee7965"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T13:01:57.661Z"),
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

Create a link

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
            ID: unifiedgosdk.Pointer("44af6dd3-1389-49e9-8034-8adbc5265f45"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("f9b9e527-2508-4223-bdf6-7fdebfada7ea"),
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
                    ID: unifiedgosdk.Pointer("1009ccc2-d6fe-4236-8c68-d0fdab642432"),
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
                    ID: unifiedgosdk.Pointer("35cf0b54-224b-4149-847c-02116695cc67"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T09:31:26.812Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/44af6dd3-1389-49e9-8034-8adbc5265f45"),
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

Retrieve a link

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

Retrieve a link

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

List all links

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

List all links

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

Update a link

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
            ID: unifiedgosdk.Pointer("d5c1154a-0df1-447f-9a79-a861d4aa4320"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T13:01:57.665Z"),
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

Update a link

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
            ID: unifiedgosdk.Pointer("0a44ac07-e03e-4d91-b1cc-dcb1f91dba36"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("a4992b6e-470c-465d-b2ef-ed9254b71ac6"),
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
                    ID: unifiedgosdk.Pointer("2cbd3403-33b5-4168-a882-a350fd62271b"),
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
                    ID: unifiedgosdk.Pointer("807d9e84-427d-4fd9-9060-9ca6fac75118"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T09:31:26.816Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/0a44ac07-e03e-4d91-b1cc-dcb1f91dba36"),
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

Remove a link

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

Remove a link

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

Update a link

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
            ID: unifiedgosdk.Pointer("d5c1154a-0df1-447f-9a79-a861d4aa4320"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T13:01:57.665Z"),
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

Update a link

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
            ID: unifiedgosdk.Pointer("0a44ac07-e03e-4d91-b1cc-dcb1f91dba36"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("a4992b6e-470c-465d-b2ef-ed9254b71ac6"),
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
                    ID: unifiedgosdk.Pointer("2cbd3403-33b5-4168-a882-a350fd62271b"),
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
                    ID: unifiedgosdk.Pointer("807d9e84-427d-4fd9-9060-9ca6fac75118"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T09:31:26.816Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/0a44ac07-e03e-4d91-b1cc-dcb1f91dba36"),
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