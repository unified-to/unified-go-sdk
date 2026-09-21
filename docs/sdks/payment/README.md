# Payment

## Overview

### Available Operations

* [CreatePaymentLink](#createpaymentlink) - Create a link
* [CreatePaymentPayment](#createpaymentpayment) - Create a payment
* [CreatePaymentSubscription](#createpaymentsubscription) - Create a subscription
* [GetPaymentLink](#getpaymentlink) - Retrieve a link
* [GetPaymentPayment](#getpaymentpayment) - Retrieve a payment
* [GetPaymentPayout](#getpaymentpayout) - Retrieve a payout
* [GetPaymentRefund](#getpaymentrefund) - Retrieve a refund
* [GetPaymentSubscription](#getpaymentsubscription) - Retrieve a subscription
* [ListPaymentLinks](#listpaymentlinks) - List all links
* [ListPaymentPayments](#listpaymentpayments) - List all payments
* [ListPaymentPayouts](#listpaymentpayouts) - List all payouts
* [ListPaymentRefunds](#listpaymentrefunds) - List all refunds
* [ListPaymentSubscriptions](#listpaymentsubscriptions) - List all subscriptions
* [PatchPaymentLink](#patchpaymentlink) - Update a link
* [PatchPaymentPayment](#patchpaymentpayment) - Update a payment
* [PatchPaymentSubscription](#patchpaymentsubscription) - Update a subscription
* [RemovePaymentLink](#removepaymentlink) - Remove a link
* [RemovePaymentPayment](#removepaymentpayment) - Remove a payment
* [RemovePaymentSubscription](#removepaymentsubscription) - Remove a subscription
* [UpdatePaymentLink](#updatepaymentlink) - Update a link
* [UpdatePaymentPayment](#updatepaymentpayment) - Update a payment
* [UpdatePaymentSubscription](#updatepaymentsubscription) - Update a subscription

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

    res, err := s.Payment.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("3c37a830-c4f4-467c-ac72-474742a7aba1"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("3aa3f836-7f85-4321-bb1c-f744771ff3b8"),
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
                    ID: unifiedgosdk.Pointer("a501a1ce-cae5-4448-831e-d60f9dc2c05b"),
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
                    ID: unifiedgosdk.Pointer("7cc7b666-ebd2-4a1a-9293-b1a6b262bb7b"),
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
                    ID: unifiedgosdk.Pointer("19c55e3a-fc1e-4715-9a20-ac2d2d04c514"),
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
                    ID: unifiedgosdk.Pointer("109d1de1-fcc2-4dc8-8411-92736c142fb6"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T22:27:29.725Z"),
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

## CreatePaymentPayment

Create a payment

### Example Usage

<!-- UsageSnippet language="go" operationID="createPaymentPayment" method="post" path="/payment/{connection_id}/payment" example="payment_payment" -->
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

    res, err := s.Payment.CreatePaymentPayment(ctx, operations.CreatePaymentPaymentRequest{
        PaymentPayment: shared.PaymentPayment{
            Allocations: []shared.PaymentAllocation{},
            CardBrand: unifiedgosdk.Pointer("AMEX"),
            CardLast4: unifiedgosdk.Pointer("0819"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("BIF"),
            ExchangeRate: unifiedgosdk.Pointer[float64](1.203),
            FeeAmount: unifiedgosdk.Pointer[float64](3.0),
            ID: unifiedgosdk.Pointer("7591d080-0abe-434e-bc81-d2daac72f966"),
            LocationID: unifiedgosdk.Pointer("94f7c68e-07de-40d1-9d6f-a0896363913f"),
            Notes: unifiedgosdk.Pointer("Tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("BANK_TRANSFER"),
            Reference: unifiedgosdk.Pointer("auctus"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCheck.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](2.0),
            TotalAmount: unifiedgosdk.Pointer[float64](44219.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-05-26T14:45:50.944Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreatePaymentPaymentRequest](../../pkg/models/operations/createpaymentpaymentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreatePaymentPaymentResponse](../../pkg/models/operations/createpaymentpaymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePaymentSubscription

Create a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="createPaymentSubscription" method="post" path="/payment/{connection_id}/subscription" example="payment_subscription" -->
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

    res, err := s.Payment.CreatePaymentSubscription(ctx, operations.CreatePaymentSubscriptionRequest{
        PaymentSubscription: shared.PaymentSubscription{
            CreatedAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Currency: unifiedgosdk.Pointer("WST"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2023-06-03T04:20:29.157Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-21T03:55:58.846Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](1.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            EndAt: types.MustNewTimeFromString("2023-05-21T12:36:09.234Z"),
            ID: unifiedgosdk.Pointer("5646479f-e014-4863-aac9-b2417de0de31"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitMonth.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            StartAt: types.MustNewTimeFromString("2023-05-29T06:04:51.030Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](75616.0),
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T10:39:39.614Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreatePaymentSubscriptionRequest](../../pkg/models/operations/createpaymentsubscriptionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreatePaymentSubscriptionResponse](../../pkg/models/operations/createpaymentsubscriptionresponse.md), error**

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

    res, err := s.Payment.GetPaymentLink(ctx, operations.GetPaymentLinkRequest{
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

## GetPaymentPayment

Retrieve a payment

### Example Usage

<!-- UsageSnippet language="go" operationID="getPaymentPayment" method="get" path="/payment/{connection_id}/payment/{id}" -->
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

    res, err := s.Payment.GetPaymentPayment(ctx, operations.GetPaymentPaymentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPaymentPaymentRequest](../../pkg/models/operations/getpaymentpaymentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetPaymentPaymentResponse](../../pkg/models/operations/getpaymentpaymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPaymentPayout

Retrieve a payout

### Example Usage

<!-- UsageSnippet language="go" operationID="getPaymentPayout" method="get" path="/payment/{connection_id}/payout/{id}" -->
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

    res, err := s.Payment.GetPaymentPayout(ctx, operations.GetPaymentPayoutRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayout != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPaymentPayoutRequest](../../pkg/models/operations/getpaymentpayoutrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPaymentPayoutResponse](../../pkg/models/operations/getpaymentpayoutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPaymentRefund

Retrieve a refund

### Example Usage

<!-- UsageSnippet language="go" operationID="getPaymentRefund" method="get" path="/payment/{connection_id}/refund/{id}" -->
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

    res, err := s.Payment.GetPaymentRefund(ctx, operations.GetPaymentRefundRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentRefund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPaymentRefundRequest](../../pkg/models/operations/getpaymentrefundrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPaymentRefundResponse](../../pkg/models/operations/getpaymentrefundresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPaymentSubscription

Retrieve a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="getPaymentSubscription" method="get" path="/payment/{connection_id}/subscription/{id}" -->
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

    res, err := s.Payment.GetPaymentSubscription(ctx, operations.GetPaymentSubscriptionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetPaymentSubscriptionRequest](../../pkg/models/operations/getpaymentsubscriptionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetPaymentSubscriptionResponse](../../pkg/models/operations/getpaymentsubscriptionresponse.md), error**

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

    res, err := s.Payment.ListPaymentLinks(ctx, operations.ListPaymentLinksRequest{
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

## ListPaymentPayments

List all payments

### Example Usage

<!-- UsageSnippet language="go" operationID="listPaymentPayments" method="get" path="/payment/{connection_id}/payment" -->
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

    res, err := s.Payment.ListPaymentPayments(ctx, operations.ListPaymentPaymentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListPaymentPaymentsRequest](../../pkg/models/operations/listpaymentpaymentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListPaymentPaymentsResponse](../../pkg/models/operations/listpaymentpaymentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPaymentPayouts

List all payouts

### Example Usage

<!-- UsageSnippet language="go" operationID="listPaymentPayouts" method="get" path="/payment/{connection_id}/payout" -->
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

    res, err := s.Payment.ListPaymentPayouts(ctx, operations.ListPaymentPayoutsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayouts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPaymentPayoutsRequest](../../pkg/models/operations/listpaymentpayoutsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPaymentPayoutsResponse](../../pkg/models/operations/listpaymentpayoutsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPaymentRefunds

List all refunds

### Example Usage

<!-- UsageSnippet language="go" operationID="listPaymentRefunds" method="get" path="/payment/{connection_id}/refund" -->
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

    res, err := s.Payment.ListPaymentRefunds(ctx, operations.ListPaymentRefundsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentRefunds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPaymentRefundsRequest](../../pkg/models/operations/listpaymentrefundsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPaymentRefundsResponse](../../pkg/models/operations/listpaymentrefundsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPaymentSubscriptions

List all subscriptions

### Example Usage

<!-- UsageSnippet language="go" operationID="listPaymentSubscriptions" method="get" path="/payment/{connection_id}/subscription" -->
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

    res, err := s.Payment.ListPaymentSubscriptions(ctx, operations.ListPaymentSubscriptionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentSubscriptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListPaymentSubscriptionsRequest](../../pkg/models/operations/listpaymentsubscriptionsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListPaymentSubscriptionsResponse](../../pkg/models/operations/listpaymentsubscriptionsresponse.md), error**

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

    res, err := s.Payment.PatchPaymentLink(ctx, operations.PatchPaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("5f68a3c9-283f-4106-859c-1a2e94c31bbf"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("e568df18-438d-4cbc-8a32-d961573e8479"),
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
                    ID: unifiedgosdk.Pointer("bb1d36cf-c8d2-4f2b-83f9-540e43fa9537"),
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
                    ID: unifiedgosdk.Pointer("6ed3ceb3-b514-4183-b461-d5998cfb96b8"),
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
                    ID: unifiedgosdk.Pointer("f8139bf0-f860-4ae2-86ee-fcdeea614862"),
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
                    ID: unifiedgosdk.Pointer("4b1c1e89-98a1-4957-b1cc-3195dc1e1fd1"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T22:27:29.740Z"),
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

## PatchPaymentPayment

Update a payment

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPaymentPayment" method="patch" path="/payment/{connection_id}/payment/{id}" example="payment_payment" -->
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

    res, err := s.Payment.PatchPaymentPayment(ctx, operations.PatchPaymentPaymentRequest{
        PaymentPayment: shared.PaymentPayment{
            Allocations: []shared.PaymentAllocation{},
            CardBrand: unifiedgosdk.Pointer("AMEX"),
            CardLast4: unifiedgosdk.Pointer("0819"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("BIF"),
            ExchangeRate: unifiedgosdk.Pointer[float64](1.203),
            FeeAmount: unifiedgosdk.Pointer[float64](3.0),
            ID: unifiedgosdk.Pointer("f9af39a6-6246-4573-83d8-93a5bf0bdd12"),
            LocationID: unifiedgosdk.Pointer("94f7c68e-07de-40d1-9d6f-a0896363913f"),
            Notes: unifiedgosdk.Pointer("Tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("BANK_TRANSFER"),
            Reference: unifiedgosdk.Pointer("auctus"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCheck.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](2.0),
            TotalAmount: unifiedgosdk.Pointer[float64](44219.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-05-26T14:45:50.955Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchPaymentPaymentRequest](../../pkg/models/operations/patchpaymentpaymentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchPaymentPaymentResponse](../../pkg/models/operations/patchpaymentpaymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPaymentSubscription

Update a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPaymentSubscription" method="patch" path="/payment/{connection_id}/subscription/{id}" example="payment_subscription" -->
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

    res, err := s.Payment.PatchPaymentSubscription(ctx, operations.PatchPaymentSubscriptionRequest{
        PaymentSubscription: shared.PaymentSubscription{
            CreatedAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Currency: unifiedgosdk.Pointer("WST"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2023-06-03T04:20:29.157Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-21T03:55:58.846Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](1.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            EndAt: types.MustNewTimeFromString("2023-05-21T12:36:09.234Z"),
            ID: unifiedgosdk.Pointer("e5056088-527a-4374-b90f-a4778a2eb421"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitMonth.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            StartAt: types.MustNewTimeFromString("2023-05-29T06:04:51.030Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](75616.0),
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T10:39:39.617Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchPaymentSubscriptionRequest](../../pkg/models/operations/patchpaymentsubscriptionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchPaymentSubscriptionResponse](../../pkg/models/operations/patchpaymentsubscriptionresponse.md), error**

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

    res, err := s.Payment.RemovePaymentLink(ctx, operations.RemovePaymentLinkRequest{
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

## RemovePaymentPayment

Remove a payment

### Example Usage

<!-- UsageSnippet language="go" operationID="removePaymentPayment" method="delete" path="/payment/{connection_id}/payment/{id}" -->
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

    res, err := s.Payment.RemovePaymentPayment(ctx, operations.RemovePaymentPaymentRequest{
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
| `request`                                                                                            | [operations.RemovePaymentPaymentRequest](../../pkg/models/operations/removepaymentpaymentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemovePaymentPaymentResponse](../../pkg/models/operations/removepaymentpaymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePaymentSubscription

Remove a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="removePaymentSubscription" method="delete" path="/payment/{connection_id}/subscription/{id}" -->
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

    res, err := s.Payment.RemovePaymentSubscription(ctx, operations.RemovePaymentSubscriptionRequest{
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
| `request`                                                                                                      | [operations.RemovePaymentSubscriptionRequest](../../pkg/models/operations/removepaymentsubscriptionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemovePaymentSubscriptionResponse](../../pkg/models/operations/removepaymentsubscriptionresponse.md), error**

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

    res, err := s.Payment.UpdatePaymentLink(ctx, operations.UpdatePaymentLinkRequest{
        PaymentLink: shared.PaymentLink{
            Amount: unifiedgosdk.Pointer[float64](81211.0),
            CreatedAt: types.MustNewTimeFromString("2023-06-04T16:11:45.685Z"),
            Currency: unifiedgosdk.Pointer("GYD"),
            Description: unifiedgosdk.Pointer("Adfero ipsa terreo benevolentia utrum."),
            ID: unifiedgosdk.Pointer("5f68a3c9-283f-4106-859c-1a2e94c31bbf"),
            IsActive: unifiedgosdk.Pointer(true),
            IsChargeableNow: unifiedgosdk.Pointer(false),
            Lineitems: []shared.PaymentLineitem{
                shared.PaymentLineitem{
                    CreatedAt: types.MustNewTimeFromString("2023-08-21T00:45:53.202Z"),
                    ID: unifiedgosdk.Pointer("e568df18-438d-4cbc-8a32-d961573e8479"),
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
                    ID: unifiedgosdk.Pointer("bb1d36cf-c8d2-4f2b-83f9-540e43fa9537"),
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
                    ID: unifiedgosdk.Pointer("6ed3ceb3-b514-4183-b461-d5998cfb96b8"),
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
                    ID: unifiedgosdk.Pointer("f8139bf0-f860-4ae2-86ee-fcdeea614862"),
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
                    ID: unifiedgosdk.Pointer("4b1c1e89-98a1-4957-b1cc-3195dc1e1fd1"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-12-12T22:27:29.740Z"),
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

## UpdatePaymentPayment

Update a payment

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePaymentPayment" method="put" path="/payment/{connection_id}/payment/{id}" example="payment_payment" -->
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

    res, err := s.Payment.UpdatePaymentPayment(ctx, operations.UpdatePaymentPaymentRequest{
        PaymentPayment: shared.PaymentPayment{
            Allocations: []shared.PaymentAllocation{},
            CardBrand: unifiedgosdk.Pointer("AMEX"),
            CardLast4: unifiedgosdk.Pointer("0819"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("BIF"),
            ExchangeRate: unifiedgosdk.Pointer[float64](1.203),
            FeeAmount: unifiedgosdk.Pointer[float64](3.0),
            ID: unifiedgosdk.Pointer("f9af39a6-6246-4573-83d8-93a5bf0bdd12"),
            LocationID: unifiedgosdk.Pointer("94f7c68e-07de-40d1-9d6f-a0896363913f"),
            Notes: unifiedgosdk.Pointer("Tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("BANK_TRANSFER"),
            Reference: unifiedgosdk.Pointer("auctus"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCheck.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](2.0),
            TotalAmount: unifiedgosdk.Pointer[float64](44219.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-05-26T14:45:50.955Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdatePaymentPaymentRequest](../../pkg/models/operations/updatepaymentpaymentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdatePaymentPaymentResponse](../../pkg/models/operations/updatepaymentpaymentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePaymentSubscription

Update a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePaymentSubscription" method="put" path="/payment/{connection_id}/subscription/{id}" example="payment_subscription" -->
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

    res, err := s.Payment.UpdatePaymentSubscription(ctx, operations.UpdatePaymentSubscriptionRequest{
        PaymentSubscription: shared.PaymentSubscription{
            CreatedAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Currency: unifiedgosdk.Pointer("WST"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2023-06-03T04:20:29.157Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-21T03:55:58.846Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](1.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            EndAt: types.MustNewTimeFromString("2023-05-21T12:36:09.234Z"),
            ID: unifiedgosdk.Pointer("e5056088-527a-4374-b90f-a4778a2eb421"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitMonth.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            StartAt: types.MustNewTimeFromString("2023-05-29T06:04:51.030Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](75616.0),
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T10:39:39.617Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdatePaymentSubscriptionRequest](../../pkg/models/operations/updatepaymentsubscriptionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdatePaymentSubscriptionResponse](../../pkg/models/operations/updatepaymentsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |