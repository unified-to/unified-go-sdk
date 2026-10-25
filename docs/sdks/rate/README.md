# Rate

## Overview

### Available Operations

* [CreateShippingRate](#createshippingrate) - Create a rate

## CreateShippingRate

Creates a ShippingRate. Provide the unified ShippingRate fields in the request body; they are mapped to the provider's API. Returns the created ShippingRate.

### Example Usage

<!-- UsageSnippet language="go" operationID="createShippingRate" method="post" path="/shipping/{connection_id}/rate" example="shipping_rate" -->
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

    res, err := s.Rate.CreateShippingRate(ctx, operations.CreateShippingRateRequest{
        ShippingRate: shared.ShippingRate{
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("a332c842-606f-446c-857d-45e9df627233"),
            Metadata: []shared.ShippingMetadata{},
            Rates: []shared.ShippingRateRate{
                shared.ShippingRateRate{
                    Amount: unifiedgosdk.Pointer[float64](44.72),
                    BaseAmount: unifiedgosdk.Pointer[float64](46.01),
                    Currency: unifiedgosdk.Pointer("USD"),
                    DeliveryDays: unifiedgosdk.Pointer[float64](2.0),
                    Description: unifiedgosdk.Pointer("Veniam denego vulgo bos turpis pax."),
                    DiscountAmount: unifiedgosdk.Pointer[float64](4.6),
                    EstimatedDays: unifiedgosdk.Pointer[float64](2.0),
                    EstimatedDeliveryEndAt: types.MustNewTimeFromString("2021-09-03T01:03:36.021Z"),
                    IsNegotiatedRate: unifiedgosdk.Pointer(true),
                    TaxAmount: unifiedgosdk.Pointer[float64](3.31),
                    Title: "Express",
                },
            },
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingRate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateShippingRateRequest](../../pkg/models/operations/createshippingraterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateShippingRateResponse](../../pkg/models/operations/createshippingrateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |