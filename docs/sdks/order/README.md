# Order

## Overview

### Available Operations

* [~~CreateAccountingOrder~~](#createaccountingorder) - Create an order :warning: **Deprecated**
* [CreateAssessmentOrder](#createassessmentorder) - Create an order
* [CreatePrintingOrder](#createprintingorder) - Create an order
* [~~GetAccountingOrder~~](#getaccountingorder) - Retrieve an order :warning: **Deprecated**
* [GetAssessmentOrder](#getassessmentorder) - Retrieve an order
* [GetPrintingOrder](#getprintingorder) - Retrieve an order
* [~~ListAccountingOrders~~](#listaccountingorders) - List all orders :warning: **Deprecated**
* [ListPrintingOrders](#listprintingorders) - List all orders
* [~~PatchAccountingOrder~~](#patchaccountingorder) - Update an order :warning: **Deprecated**
* [PatchAssessmentOrder](#patchassessmentorder) - Update an order
* [PatchPrintingOrder](#patchprintingorder) - Update an order
* [~~RemoveAccountingOrder~~](#removeaccountingorder) - Remove an order :warning: **Deprecated**
* [RemovePrintingOrder](#removeprintingorder) - Remove an order
* [~~UpdateAccountingOrder~~](#updateaccountingorder) - Update an order :warning: **Deprecated**
* [UpdateAssessmentOrder](#updateassessmentorder) - Update an order
* [UpdatePrintingOrder](#updateprintingorder) - Update an order

## ~~CreateAccountingOrder~~

Creates a AccountingOrder. Provide the unified AccountingOrder fields in the request body; they are mapped to the provider's API. Returns the created AccountingOrder.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingOrder" method="post" path="/accounting/{connection_id}/order" example="accounting_order" -->
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

    res, err := s.Order.CreateAccountingOrder(ctx, operations.CreateAccountingOrderRequest{
        AccountingOrder: shared.AccountingOrder{
            BillingAddress: &shared.PropertyAccountingOrderBillingAddress{
                Address1: unifiedgosdk.Pointer("125 Hayden Highway"),
                City: unifiedgosdk.Pointer("Altoona"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("94706-7691"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("VT"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ID: unifiedgosdk.Pointer("2979e2a1-ffd4-4322-8fc3-477ee3de2622"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("450 Garett Mall"),
                City: unifiedgosdk.Pointer("New Kamronburgh"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("76546"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("VA"),
            },
            Status: shared.AccountingOrderStatusAuthorized.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-02-15T17:02:42.994Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAccountingOrderRequest](../../pkg/models/operations/createaccountingorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateAccountingOrderResponse](../../pkg/models/operations/createaccountingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAssessmentOrder

Creates a AssessmentOrder. Provide the unified AssessmentOrder fields in the request body; they are mapped to the provider's API. Returns the created AssessmentOrder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAssessmentOrder" method="post" path="/assessment/{connection_id}/order" example="assessment_order" -->
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

    res, err := s.Order.CreateAssessmentOrder(ctx, operations.CreateAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("0e9aae88-f7b1-44db-9eb7-bae5b1c01c18"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "jordan.braun@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Jordan"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Braun"),
            ProfileName: unifiedgosdk.Pointer("Jordan Braun"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://astonishing-bug.com/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "322.658.2806 x3159",
            },
            Reference: unifiedgosdk.Pointer("calcar"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://nervous-daughter.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://idolized-suv.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-04T10:05:17.586Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAssessmentOrderRequest](../../pkg/models/operations/createassessmentorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateAssessmentOrderResponse](../../pkg/models/operations/createassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePrintingOrder

Creates a PrintingOrder. Provide the unified PrintingOrder fields in the request body; they are mapped to the provider's API. Returns the created PrintingOrder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPrintingOrder" method="post" path="/printing/{connection_id}/order" example="printing_order" -->
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

    res, err := s.Order.CreatePrintingOrder(ctx, operations.CreatePrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("9d6516d5-784b-4442-9751-127e629dba39"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.017Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePrintingOrderRequest](../../pkg/models/operations/createprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePrintingOrderResponse](../../pkg/models/operations/createprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~GetAccountingOrder~~

Retrieves a single AccountingOrder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingOrder" method="get" path="/accounting/{connection_id}/order/{id}" -->
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

    res, err := s.Order.GetAccountingOrder(ctx, operations.GetAccountingOrderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAccountingOrderRequest](../../pkg/models/operations/getaccountingorderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAccountingOrderResponse](../../pkg/models/operations/getaccountingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAssessmentOrder

Retrieves a single AssessmentOrder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAssessmentOrder" method="get" path="/assessment/{connection_id}/order/{id}" -->
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

    res, err := s.Order.GetAssessmentOrder(ctx, operations.GetAssessmentOrderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAssessmentOrderRequest](../../pkg/models/operations/getassessmentorderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAssessmentOrderResponse](../../pkg/models/operations/getassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPrintingOrder

Retrieves a single PrintingOrder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrintingOrder" method="get" path="/printing/{connection_id}/order/{id}" -->
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

    res, err := s.Order.GetPrintingOrder(ctx, operations.GetPrintingOrderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPrintingOrderRequest](../../pkg/models/operations/getprintingorderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPrintingOrderResponse](../../pkg/models/operations/getprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~ListAccountingOrders~~

Returns a list of AccountingOrders. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingOrders" method="get" path="/accounting/{connection_id}/order" -->
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

    res, err := s.Order.ListAccountingOrders(ctx, operations.ListAccountingOrdersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingOrders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAccountingOrdersRequest](../../pkg/models/operations/listaccountingordersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAccountingOrdersResponse](../../pkg/models/operations/listaccountingordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPrintingOrders

Returns a list of PrintingOrders. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPrintingOrders" method="get" path="/printing/{connection_id}/order" -->
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

    res, err := s.Order.ListPrintingOrders(ctx, operations.ListPrintingOrdersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPrintingOrdersRequest](../../pkg/models/operations/listprintingordersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPrintingOrdersResponse](../../pkg/models/operations/listprintingordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~PatchAccountingOrder~~

Updates the AccountingOrder identified by `id`. Provide the unified AccountingOrder fields to change in the request body.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingOrder" method="patch" path="/accounting/{connection_id}/order/{id}" example="accounting_order" -->
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

    res, err := s.Order.PatchAccountingOrder(ctx, operations.PatchAccountingOrderRequest{
        AccountingOrder: shared.AccountingOrder{
            BillingAddress: &shared.PropertyAccountingOrderBillingAddress{
                Address1: unifiedgosdk.Pointer("125 Hayden Highway"),
                City: unifiedgosdk.Pointer("Altoona"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("94706-7691"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("VT"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ID: unifiedgosdk.Pointer("b048f651-5d88-4b31-b8fd-14f1914ea89f"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("450 Garett Mall"),
                City: unifiedgosdk.Pointer("New Kamronburgh"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("76546"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("VA"),
            },
            Status: shared.AccountingOrderStatusAuthorized.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-02-15T17:02:43.001Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchAccountingOrderRequest](../../pkg/models/operations/patchaccountingorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchAccountingOrderResponse](../../pkg/models/operations/patchaccountingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAssessmentOrder

Updates the AssessmentOrder identified by `id`. Provide the unified AssessmentOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAssessmentOrder" method="patch" path="/assessment/{connection_id}/order/{id}" example="assessment_order" -->
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

    res, err := s.Order.PatchAssessmentOrder(ctx, operations.PatchAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("1fcae33f-aaf4-4fbf-ada4-05639d7c5869"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "jordan.braun@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Jordan"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Braun"),
            ProfileName: unifiedgosdk.Pointer("Jordan Braun"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://astonishing-bug.com/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "322.658.2806 x3159",
            },
            Reference: unifiedgosdk.Pointer("calcar"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://nervous-daughter.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://idolized-suv.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-04T10:05:17.593Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchAssessmentOrderRequest](../../pkg/models/operations/patchassessmentorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchAssessmentOrderResponse](../../pkg/models/operations/patchassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPrintingOrder

Updates the PrintingOrder identified by `id`. Provide the unified PrintingOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPrintingOrder" method="patch" path="/printing/{connection_id}/order/{id}" example="printing_order" -->
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

    res, err := s.Order.PatchPrintingOrder(ctx, operations.PatchPrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("af220e4d-5540-45c4-bb37-aeb092251e06"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.037Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchPrintingOrderRequest](../../pkg/models/operations/patchprintingorderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchPrintingOrderResponse](../../pkg/models/operations/patchprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~RemoveAccountingOrder~~

Deletes the AccountingOrder identified by `id`.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingOrder" method="delete" path="/accounting/{connection_id}/order/{id}" -->
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

    res, err := s.Order.RemoveAccountingOrder(ctx, operations.RemoveAccountingOrderRequest{
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
| `request`                                                                                              | [operations.RemoveAccountingOrderRequest](../../pkg/models/operations/removeaccountingorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveAccountingOrderResponse](../../pkg/models/operations/removeaccountingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePrintingOrder

Deletes the PrintingOrder identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePrintingOrder" method="delete" path="/printing/{connection_id}/order/{id}" -->
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

    res, err := s.Order.RemovePrintingOrder(ctx, operations.RemovePrintingOrderRequest{
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
| `request`                                                                                          | [operations.RemovePrintingOrderRequest](../../pkg/models/operations/removeprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemovePrintingOrderResponse](../../pkg/models/operations/removeprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~UpdateAccountingOrder~~

Updates the AccountingOrder identified by `id`. Provide the unified AccountingOrder fields to change in the request body.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingOrder" method="put" path="/accounting/{connection_id}/order/{id}" example="accounting_order" -->
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

    res, err := s.Order.UpdateAccountingOrder(ctx, operations.UpdateAccountingOrderRequest{
        AccountingOrder: shared.AccountingOrder{
            BillingAddress: &shared.PropertyAccountingOrderBillingAddress{
                Address1: unifiedgosdk.Pointer("125 Hayden Highway"),
                City: unifiedgosdk.Pointer("Altoona"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("94706-7691"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("VT"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            ID: unifiedgosdk.Pointer("b048f651-5d88-4b31-b8fd-14f1914ea89f"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("450 Garett Mall"),
                City: unifiedgosdk.Pointer("New Kamronburgh"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("76546"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("VA"),
            },
            Status: shared.AccountingOrderStatusAuthorized.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypeSales.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-02-15T17:02:43.001Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAccountingOrderRequest](../../pkg/models/operations/updateaccountingorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAccountingOrderResponse](../../pkg/models/operations/updateaccountingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAssessmentOrder

Updates the AssessmentOrder identified by `id`. Provide the unified AssessmentOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAssessmentOrder" method="put" path="/assessment/{connection_id}/order/{id}" example="assessment_order" -->
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

    res, err := s.Order.UpdateAssessmentOrder(ctx, operations.UpdateAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("1fcae33f-aaf4-4fbf-ada4-05639d7c5869"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "jordan.braun@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Jordan"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Braun"),
            ProfileName: unifiedgosdk.Pointer("Jordan Braun"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://astonishing-bug.com/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "322.658.2806 x3159",
            },
            Reference: unifiedgosdk.Pointer("calcar"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://nervous-daughter.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://idolized-suv.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-04T10:05:17.593Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAssessmentOrderRequest](../../pkg/models/operations/updateassessmentorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAssessmentOrderResponse](../../pkg/models/operations/updateassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePrintingOrder

Updates the PrintingOrder identified by `id`. Provide the unified PrintingOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePrintingOrder" method="put" path="/printing/{connection_id}/order/{id}" example="printing_order" -->
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

    res, err := s.Order.UpdatePrintingOrder(ctx, operations.UpdatePrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("af220e4d-5540-45c4-bb37-aeb092251e06"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.037Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdatePrintingOrderRequest](../../pkg/models/operations/updateprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdatePrintingOrderResponse](../../pkg/models/operations/updateprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |