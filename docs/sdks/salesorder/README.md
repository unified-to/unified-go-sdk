# Salesorder

## Overview

### Available Operations

* [CreateAccountingSalesorder](#createaccountingsalesorder) - Create a salesorder
* [GetAccountingSalesorder](#getaccountingsalesorder) - Retrieve a salesorder
* [ListAccountingSalesorders](#listaccountingsalesorders) - List all salesorders
* [PatchAccountingSalesorder](#patchaccountingsalesorder) - Update a salesorder
* [RemoveAccountingSalesorder](#removeaccountingsalesorder) - Remove a salesorder
* [UpdateAccountingSalesorder](#updateaccountingsalesorder) - Update a salesorder

## CreateAccountingSalesorder

Create a salesorder

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingSalesorder" method="post" path="/accounting/{connection_id}/salesorder" example="accounting_salesorder" -->
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

    res, err := s.Salesorder.CreateAccountingSalesorder(ctx, operations.CreateAccountingSalesorderRequest{
        AccountingSalesorder: shared.AccountingSalesorder{
            BillingAddress: &shared.PropertyAccountingSalesorderBillingAddress{
                Address1: unifiedgosdk.Pointer("26530 Stroman Rest"),
                Address2: unifiedgosdk.Pointer("Suite 801"),
                City: unifiedgosdk.Pointer("Pocatello"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("05015-8546"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CategoryIds: []string{},
            ClosedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.155Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            Fees: []shared.AccountingFee{},
            FulfillmentType: shared.FulfillmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("2ee1a550-d200-4b90-84de-64861a01cf0a"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            OrderNumber: unifiedgosdk.Pointer("832663"),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            RefundedAmount: unifiedgosdk.Pointer[float64](0.0),
            SalesChannel: unifiedgosdk.Pointer("Wholesale"),
            ServiceChargeAmount: unifiedgosdk.Pointer[float64](0.0),
            ShippingAddress: &shared.PropertyAccountingSalesorderShippingAddress{
                Address1: unifiedgosdk.Pointer("69878 Bradley Mill"),
                City: unifiedgosdk.Pointer("Port Matildestad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("07989-2148"),
                Region: unifiedgosdk.Pointer("Arkansas"),
                RegionCode: unifiedgosdk.Pointer("AK"),
            },
            Status: shared.AccountingSalesorderStatusPaid.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.155Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingSalesorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingSalesorderRequest](../../pkg/models/operations/createaccountingsalesorderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateAccountingSalesorderResponse](../../pkg/models/operations/createaccountingsalesorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingSalesorder

Retrieve a salesorder

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingSalesorder" method="get" path="/accounting/{connection_id}/salesorder/{id}" -->
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

    res, err := s.Salesorder.GetAccountingSalesorder(ctx, operations.GetAccountingSalesorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingSalesorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetAccountingSalesorderRequest](../../pkg/models/operations/getaccountingsalesorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAccountingSalesorderResponse](../../pkg/models/operations/getaccountingsalesorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingSalesorders

List all salesorders

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingSalesorders" method="get" path="/accounting/{connection_id}/salesorder" -->
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

    res, err := s.Salesorder.ListAccountingSalesorders(ctx, operations.ListAccountingSalesordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingSalesorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListAccountingSalesordersRequest](../../pkg/models/operations/listaccountingsalesordersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListAccountingSalesordersResponse](../../pkg/models/operations/listaccountingsalesordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingSalesorder

Update a salesorder

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingSalesorder" method="patch" path="/accounting/{connection_id}/salesorder/{id}" example="accounting_salesorder" -->
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

    res, err := s.Salesorder.PatchAccountingSalesorder(ctx, operations.PatchAccountingSalesorderRequest{
        AccountingSalesorder: shared.AccountingSalesorder{
            BillingAddress: &shared.PropertyAccountingSalesorderBillingAddress{
                Address1: unifiedgosdk.Pointer("26530 Stroman Rest"),
                Address2: unifiedgosdk.Pointer("Suite 801"),
                City: unifiedgosdk.Pointer("Pocatello"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("05015-8546"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CategoryIds: []string{},
            ClosedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.174Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            Fees: []shared.AccountingFee{},
            FulfillmentType: shared.FulfillmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("81444e74-2438-4d48-a12a-446390d8c063"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            OrderNumber: unifiedgosdk.Pointer("832663"),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            RefundedAmount: unifiedgosdk.Pointer[float64](0.0),
            SalesChannel: unifiedgosdk.Pointer("Wholesale"),
            ServiceChargeAmount: unifiedgosdk.Pointer[float64](0.0),
            ShippingAddress: &shared.PropertyAccountingSalesorderShippingAddress{
                Address1: unifiedgosdk.Pointer("69878 Bradley Mill"),
                City: unifiedgosdk.Pointer("Port Matildestad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("07989-2148"),
                Region: unifiedgosdk.Pointer("Arkansas"),
                RegionCode: unifiedgosdk.Pointer("AK"),
            },
            Status: shared.AccountingSalesorderStatusPaid.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.174Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingSalesorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchAccountingSalesorderRequest](../../pkg/models/operations/patchaccountingsalesorderrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PatchAccountingSalesorderResponse](../../pkg/models/operations/patchaccountingsalesorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingSalesorder

Remove a salesorder

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingSalesorder" method="delete" path="/accounting/{connection_id}/salesorder/{id}" -->
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

    res, err := s.Salesorder.RemoveAccountingSalesorder(ctx, operations.RemoveAccountingSalesorderRequest{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.RemoveAccountingSalesorderRequest](../../pkg/models/operations/removeaccountingsalesorderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.RemoveAccountingSalesorderResponse](../../pkg/models/operations/removeaccountingsalesorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingSalesorder

Update a salesorder

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingSalesorder" method="put" path="/accounting/{connection_id}/salesorder/{id}" example="accounting_salesorder" -->
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

    res, err := s.Salesorder.UpdateAccountingSalesorder(ctx, operations.UpdateAccountingSalesorderRequest{
        AccountingSalesorder: shared.AccountingSalesorder{
            BillingAddress: &shared.PropertyAccountingSalesorderBillingAddress{
                Address1: unifiedgosdk.Pointer("26530 Stroman Rest"),
                Address2: unifiedgosdk.Pointer("Suite 801"),
                City: unifiedgosdk.Pointer("Pocatello"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("05015-8546"),
                Region: unifiedgosdk.Pointer("Louisiana"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CategoryIds: []string{},
            ClosedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.174Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            Fees: []shared.AccountingFee{},
            FulfillmentType: shared.FulfillmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("81444e74-2438-4d48-a12a-446390d8c063"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            OrderNumber: unifiedgosdk.Pointer("832663"),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-01-17T16:11:50.310Z"),
            RefundedAmount: unifiedgosdk.Pointer[float64](0.0),
            SalesChannel: unifiedgosdk.Pointer("Wholesale"),
            ServiceChargeAmount: unifiedgosdk.Pointer[float64](0.0),
            ShippingAddress: &shared.PropertyAccountingSalesorderShippingAddress{
                Address1: unifiedgosdk.Pointer("69878 Bradley Mill"),
                City: unifiedgosdk.Pointer("Port Matildestad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("07989-2148"),
                Region: unifiedgosdk.Pointer("Arkansas"),
                RegionCode: unifiedgosdk.Pointer("AK"),
            },
            Status: shared.AccountingSalesorderStatusPaid.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TipAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-25T12:57:23.174Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingSalesorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateAccountingSalesorderRequest](../../pkg/models/operations/updateaccountingsalesorderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateAccountingSalesorderResponse](../../pkg/models/operations/updateaccountingsalesorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |