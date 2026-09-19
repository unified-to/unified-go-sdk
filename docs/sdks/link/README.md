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
            ID: unifiedgosdk.Pointer("f8e8095c-0f1e-458a-91ab-2c7a3bc373aa"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.143Z"),
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
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("a8f5322d-eaac-4317-8c71-09d28989af2e"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("5a041525-0801-434a-879a-d2fb8293a705"),
                    ItemDescription: unifiedgosdk.Pointer("Experience the white brilliance of our Hat, perfect for aggravating environments"),
                    ItemName: unifiedgosdk.Pointer("Licensed Marble Mouse"),
                    ItemSku: unifiedgosdk.Pointer("TAD4EYLVRI"),
                    Notes: unifiedgosdk.Pointer("Charisma theca video verus conduco attollo cervus decretum viridis."),
                    TaxAmount: unifiedgosdk.Pointer[float64](221.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1841.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](270.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-02-12T17:31:25.507Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-09-30T05:29:29.258Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](15.0),
                    ID: unifiedgosdk.Pointer("a653c1ab-4dc1-484c-8063-1a8c8d4201eb"),
                    ItemDescription: unifiedgosdk.Pointer("New Chicken model with 79 GB RAM, 846 GB storage, and lovely features"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Steel Table"),
                    ItemSku: unifiedgosdk.Pointer("V8HQCDQYUZ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](150.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2037.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](317.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-31T11:10:09.190Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-12-16T13:52:52.341Z"),
                    ID: unifiedgosdk.Pointer("f014f633-30c0-4ead-9a12-257a38e4f149"),
                    ItemDescription: unifiedgosdk.Pointer("Dach - Wolff's most advanced Car technology increases dense capabilities"),
                    ItemName: unifiedgosdk.Pointer("Modern Gold Soap"),
                    ItemSku: unifiedgosdk.Pointer("DYGKCTCLDJ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](41.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](281.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](30.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-22T16:35:07.583Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-12T19:45:39.705Z"),
                    ID: unifiedgosdk.Pointer("e705d74b-e8cb-44fc-9a28-3f95c3dfe1dc"),
                    ItemDescription: unifiedgosdk.Pointer("The sleek and unimportant Salad comes with salmon LED lighting for smart functionality"),
                    ItemName: unifiedgosdk.Pointer("Generic Aluminum Ball"),
                    ItemSku: unifiedgosdk.Pointer("BSBAXWAAFF"),
                    Notes: unifiedgosdk.Pointer("Cubo adversus victus subito asperiores vereor cibo tabgo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](6.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](78.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](24.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](3.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-11-13T12:39:15.951Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-02-14T06:21:13.641Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](171.0),
                    ID: unifiedgosdk.Pointer("b2d01b6a-0634-4b5e-9bea-b5178de33322"),
                    ItemDescription: unifiedgosdk.Pointer("New Bike model with 29 GB RAM, 271 GB storage, and minty features"),
                    ItemName: unifiedgosdk.Pointer("Incredible Aluminum Chicken"),
                    ItemSku: unifiedgosdk.Pointer("6ERMJK20HE"),
                    TaxAmount: unifiedgosdk.Pointer[float64](263.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](3708.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](452.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-01-31T21:39:30.894Z"),
                },
            },
            SuccessURL: unifiedgosdk.Pointer("https://parched-kettledrum.com/"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-11T12:09:34.258Z"),
            URL: unifiedgosdk.Pointer("https://forceful-laughter.biz/"),
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
            ID: unifiedgosdk.Pointer("7ee7d961-69a9-4d0d-abc8-a952c590c9e6"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.146Z"),
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
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("572d6790-70ac-41ca-b4d2-8dee83ac6303"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("8faeb0e8-f090-442b-9a1c-b551290d2207"),
                    ItemDescription: unifiedgosdk.Pointer("Experience the white brilliance of our Hat, perfect for aggravating environments"),
                    ItemName: unifiedgosdk.Pointer("Licensed Marble Mouse"),
                    ItemSku: unifiedgosdk.Pointer("TAD4EYLVRI"),
                    Notes: unifiedgosdk.Pointer("Charisma theca video verus conduco attollo cervus decretum viridis."),
                    TaxAmount: unifiedgosdk.Pointer[float64](221.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1841.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](270.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-02-12T17:31:25.507Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-09-30T05:29:29.258Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](15.0),
                    ID: unifiedgosdk.Pointer("a900ad8b-784e-447b-a6bd-3673245f0494"),
                    ItemDescription: unifiedgosdk.Pointer("New Chicken model with 79 GB RAM, 846 GB storage, and lovely features"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Steel Table"),
                    ItemSku: unifiedgosdk.Pointer("V8HQCDQYUZ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](150.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2037.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](317.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-31T11:10:09.190Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-12-16T13:52:52.341Z"),
                    ID: unifiedgosdk.Pointer("859b6741-bdce-475c-8855-e64444bf7017"),
                    ItemDescription: unifiedgosdk.Pointer("Dach - Wolff's most advanced Car technology increases dense capabilities"),
                    ItemName: unifiedgosdk.Pointer("Modern Gold Soap"),
                    ItemSku: unifiedgosdk.Pointer("DYGKCTCLDJ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](41.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](281.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](30.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-22T16:35:07.583Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-12T19:45:39.705Z"),
                    ID: unifiedgosdk.Pointer("186bf115-1506-4902-9da2-3dfd31015db1"),
                    ItemDescription: unifiedgosdk.Pointer("The sleek and unimportant Salad comes with salmon LED lighting for smart functionality"),
                    ItemName: unifiedgosdk.Pointer("Generic Aluminum Ball"),
                    ItemSku: unifiedgosdk.Pointer("BSBAXWAAFF"),
                    Notes: unifiedgosdk.Pointer("Cubo adversus victus subito asperiores vereor cibo tabgo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](6.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](78.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](24.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](3.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-11-13T12:39:15.951Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-02-14T06:21:13.641Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](171.0),
                    ID: unifiedgosdk.Pointer("bec30fae-971d-4a64-9599-f28801bcf427"),
                    ItemDescription: unifiedgosdk.Pointer("New Bike model with 29 GB RAM, 271 GB storage, and minty features"),
                    ItemName: unifiedgosdk.Pointer("Incredible Aluminum Chicken"),
                    ItemSku: unifiedgosdk.Pointer("6ERMJK20HE"),
                    TaxAmount: unifiedgosdk.Pointer[float64](263.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](3708.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](452.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-01-31T21:39:30.894Z"),
                },
            },
            SuccessURL: unifiedgosdk.Pointer("https://parched-kettledrum.com/"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-11T12:09:34.274Z"),
            URL: unifiedgosdk.Pointer("https://forceful-laughter.biz/"),
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
            ID: unifiedgosdk.Pointer("7ee7d961-69a9-4d0d-abc8-a952c590c9e6"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.146Z"),
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
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("572d6790-70ac-41ca-b4d2-8dee83ac6303"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("8faeb0e8-f090-442b-9a1c-b551290d2207"),
                    ItemDescription: unifiedgosdk.Pointer("Experience the white brilliance of our Hat, perfect for aggravating environments"),
                    ItemName: unifiedgosdk.Pointer("Licensed Marble Mouse"),
                    ItemSku: unifiedgosdk.Pointer("TAD4EYLVRI"),
                    Notes: unifiedgosdk.Pointer("Charisma theca video verus conduco attollo cervus decretum viridis."),
                    TaxAmount: unifiedgosdk.Pointer[float64](221.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](1841.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](270.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-02-12T17:31:25.507Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-09-30T05:29:29.258Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](15.0),
                    ID: unifiedgosdk.Pointer("a900ad8b-784e-447b-a6bd-3673245f0494"),
                    ItemDescription: unifiedgosdk.Pointer("New Chicken model with 79 GB RAM, 846 GB storage, and lovely features"),
                    ItemName: unifiedgosdk.Pointer("Intelligent Steel Table"),
                    ItemSku: unifiedgosdk.Pointer("V8HQCDQYUZ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](150.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2037.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](317.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](6.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-31T11:10:09.190Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-12-16T13:52:52.341Z"),
                    ID: unifiedgosdk.Pointer("859b6741-bdce-475c-8855-e64444bf7017"),
                    ItemDescription: unifiedgosdk.Pointer("Dach - Wolff's most advanced Car technology increases dense capabilities"),
                    ItemName: unifiedgosdk.Pointer("Modern Gold Soap"),
                    ItemSku: unifiedgosdk.Pointer("DYGKCTCLDJ"),
                    TaxAmount: unifiedgosdk.Pointer[float64](41.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](281.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](30.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-05-22T16:35:07.583Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-12T19:45:39.705Z"),
                    ID: unifiedgosdk.Pointer("186bf115-1506-4902-9da2-3dfd31015db1"),
                    ItemDescription: unifiedgosdk.Pointer("The sleek and unimportant Salad comes with salmon LED lighting for smart functionality"),
                    ItemName: unifiedgosdk.Pointer("Generic Aluminum Ball"),
                    ItemSku: unifiedgosdk.Pointer("BSBAXWAAFF"),
                    Notes: unifiedgosdk.Pointer("Cubo adversus victus subito asperiores vereor cibo tabgo."),
                    TaxAmount: unifiedgosdk.Pointer[float64](6.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](78.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](24.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](3.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-11-13T12:39:15.951Z"),
                },
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-02-14T06:21:13.641Z"),
                    DiscountAmount: unifiedgosdk.Pointer[float64](171.0),
                    ID: unifiedgosdk.Pointer("bec30fae-971d-4a64-9599-f28801bcf427"),
                    ItemDescription: unifiedgosdk.Pointer("New Bike model with 29 GB RAM, 271 GB storage, and minty features"),
                    ItemName: unifiedgosdk.Pointer("Incredible Aluminum Chicken"),
                    ItemSku: unifiedgosdk.Pointer("6ERMJK20HE"),
                    TaxAmount: unifiedgosdk.Pointer[float64](263.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](3708.0),
                    UnitAmount: unifiedgosdk.Pointer[float64](452.0),
                    UnitQuantity: unifiedgosdk.Pointer[float64](8.0),
                    UpdatedAt: types.MustNewTimeFromString("2023-01-31T21:39:30.894Z"),
                },
            },
            SuccessURL: unifiedgosdk.Pointer("https://parched-kettledrum.com/"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-11T12:09:34.274Z"),
            URL: unifiedgosdk.Pointer("https://forceful-laughter.biz/"),
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