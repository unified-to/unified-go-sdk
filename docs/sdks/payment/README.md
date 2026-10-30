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

    res, err := s.Payment.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
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

## CreatePaymentPayment

Creates a PaymentPayment. Provide the unified PaymentPayment fields in the request body; they are mapped to the provider's API. Returns the created PaymentPayment.

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
            CardBrand: unifiedgosdk.Pointer("MASTERCARD"),
            CardLast4: unifiedgosdk.Pointer("1529"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            FeeAmount: unifiedgosdk.Pointer[float64](2055.82),
            ID: unifiedgosdk.Pointer("afaf821d-80d8-4354-b5a3-88d59e8b3b7f"),
            Metadata: []shared.PaymentMetadata{},
            NetAmount: unifiedgosdk.Pointer[float64](68824.18),
            Notes: unifiedgosdk.Pointer("Bibo tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            Reference: unifiedgosdk.Pointer("PMT-119080"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCash.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](70880.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-14T20:49:54.160Z"),
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

Creates a PaymentSubscription. Provide the unified PaymentSubscription fields in the request body; they are mapped to the provider's API. Returns the created PaymentSubscription.

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
            Currency: unifiedgosdk.Pointer("USD"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2024-02-08T10:05:22.912Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](8.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            ID: unifiedgosdk.Pointer("583de0bb-62b4-4192-ba06-bbd87356f3a2"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitYear.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            Metadata: []shared.PaymentMetadata{},
            Month: unifiedgosdk.Pointer[float64](5.0),
            StartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-10-05T13:54:10.622Z"),
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

Retrieves a single PaymentPayment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single PaymentPayout by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single PaymentRefund by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single PaymentSubscription by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of PaymentPayments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of PaymentPayouts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of PaymentRefunds. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of PaymentSubscriptions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Payment.PatchPaymentLink(ctx, operations.PatchPaymentLinkRequest{
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

## PatchPaymentPayment

Updates the PaymentPayment identified by `id`. Provide the unified PaymentPayment fields to change in the request body.

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
            CardBrand: unifiedgosdk.Pointer("MASTERCARD"),
            CardLast4: unifiedgosdk.Pointer("1529"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            FeeAmount: unifiedgosdk.Pointer[float64](2055.82),
            ID: unifiedgosdk.Pointer("beeb730f-3813-4ce7-8057-4b35214f4c04"),
            Metadata: []shared.PaymentMetadata{},
            NetAmount: unifiedgosdk.Pointer[float64](68824.18),
            Notes: unifiedgosdk.Pointer("Bibo tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            Reference: unifiedgosdk.Pointer("PMT-119080"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCash.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](70880.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-14T20:49:54.201Z"),
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

Updates the PaymentSubscription identified by `id`. Provide the unified PaymentSubscription fields to change in the request body.

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
            Currency: unifiedgosdk.Pointer("USD"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2024-02-08T10:05:22.912Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](8.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            ID: unifiedgosdk.Pointer("e2fabf5d-21ae-4203-93fd-5d5bfcca0c35"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitYear.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            Metadata: []shared.PaymentMetadata{},
            Month: unifiedgosdk.Pointer[float64](5.0),
            StartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-10-05T13:54:10.627Z"),
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

Deletes the PaymentPayment identified by `id`.

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

Deletes the PaymentSubscription identified by `id`.

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

    res, err := s.Payment.UpdatePaymentLink(ctx, operations.UpdatePaymentLinkRequest{
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

## UpdatePaymentPayment

Updates the PaymentPayment identified by `id`. Provide the unified PaymentPayment fields to change in the request body.

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
            CardBrand: unifiedgosdk.Pointer("MASTERCARD"),
            CardLast4: unifiedgosdk.Pointer("1529"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            FeeAmount: unifiedgosdk.Pointer[float64](2055.82),
            ID: unifiedgosdk.Pointer("beeb730f-3813-4ce7-8057-4b35214f4c04"),
            Metadata: []shared.PaymentMetadata{},
            NetAmount: unifiedgosdk.Pointer[float64](68824.18),
            Notes: unifiedgosdk.Pointer("Bibo tactus vilicus."),
            PaidAt: types.MustNewTimeFromString("2022-03-10T00:19:42.086Z"),
            PaymentMethod: unifiedgosdk.Pointer("CASH"),
            Reference: unifiedgosdk.Pointer("PMT-119080"),
            Status: shared.PaymentPaymentStatusSucceeded.ToPointer(),
            TenderType: shared.TenderTypeCash.ToPointer(),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](70880.0),
            Type: shared.PaymentPaymentTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-14T20:49:54.201Z"),
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

Updates the PaymentSubscription identified by `id`. Provide the unified PaymentSubscription fields to change in the request body.

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
            Currency: unifiedgosdk.Pointer("USD"),
            CurrentPeriodEndAt: types.MustNewTimeFromString("2024-02-08T10:05:22.912Z"),
            CurrentPeriodStartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            DayOfMonth: unifiedgosdk.Pointer[float64](8.0),
            Description: unifiedgosdk.Pointer("Innovative Mouse featuring important technology and Bamboo construction"),
            ID: unifiedgosdk.Pointer("e2fabf5d-21ae-4203-93fd-5d5bfcca0c35"),
            Interval: unifiedgosdk.Pointer[float64](1.0),
            IntervalUnit: shared.IntervalUnitYear.ToPointer(),
            Lineitems: []shared.PaymentLineitem{},
            Metadata: []shared.PaymentMetadata{},
            Month: unifiedgosdk.Pointer[float64](5.0),
            StartAt: types.MustNewTimeFromString("2023-05-08T10:11:03.414Z"),
            Status: shared.PaymentSubscriptionStatusActive.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-10-05T13:54:10.627Z"),
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