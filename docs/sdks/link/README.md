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
            ID: unifiedgosdk.Pointer("51e62917-796f-4508-a00e-2b1312725ccc"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T06:11:28.775Z"),
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
            ID: unifiedgosdk.Pointer("ddc00bcf-3103-42d8-bd6b-c7617d2a7d4c"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("a8e524aa-985f-479e-83a4-8dc973cecf7c"),
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
                    ID: unifiedgosdk.Pointer("30b780f9-249d-4bd9-8ec3-df9f34c52614"),
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
                    ID: unifiedgosdk.Pointer("033b7cb5-0e77-4ee5-b614-fec8a9fd659b"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T03:22:14.517Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/ddc00bcf-3103-42d8-bd6b-c7617d2a7d4c"),
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
            ID: unifiedgosdk.Pointer("871cb221-94b6-4339-884a-09cfb75eba6b"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T06:11:28.779Z"),
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
            ID: unifiedgosdk.Pointer("59cfd03e-78b5-4bad-95a3-d17c9fe7f405"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("229cefe0-0145-4e36-afc4-43816f1bc807"),
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
                    ID: unifiedgosdk.Pointer("f15d57a5-10c5-4991-8c14-c444c09187bf"),
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
                    ID: unifiedgosdk.Pointer("7cfc68af-0e58-4434-9fee-25b744138c15"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T03:22:14.528Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/59cfd03e-78b5-4bad-95a3-d17c9fe7f405"),
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
            ID: unifiedgosdk.Pointer("871cb221-94b6-4339-884a-09cfb75eba6b"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CalendarMetadata{},
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-15T06:11:28.779Z"),
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
            ID: unifiedgosdk.Pointer("59cfd03e-78b5-4bad-95a3-d17c9fe7f405"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(true),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
                    ID: unifiedgosdk.Pointer("229cefe0-0145-4e36-afc4-43816f1bc807"),
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
                    ID: unifiedgosdk.Pointer("f15d57a5-10c5-4991-8c14-c444c09187bf"),
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
                    ID: unifiedgosdk.Pointer("7cfc68af-0e58-4434-9fee-25b744138c15"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-04-13T03:22:14.528Z"),
            URL: unifiedgosdk.Pointer("https://pay.example.com/59cfd03e-78b5-4bad-95a3-d17c9fe7f405"),
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