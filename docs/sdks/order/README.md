# Order

## Overview

### Available Operations

* [CreateAccountingOrder](#createaccountingorder) - Create an order
* [CreateAssessmentOrder](#createassessmentorder) - Create an order
* [GetAccountingOrder](#getaccountingorder) - Retrieve an order
* [GetAssessmentOrder](#getassessmentorder) - Retrieve an order
* [ListAccountingOrders](#listaccountingorders) - List all orders
* [PatchAccountingOrder](#patchaccountingorder) - Update an order
* [PatchAssessmentOrder](#patchassessmentorder) - Update an order
* [RemoveAccountingOrder](#removeaccountingorder) - Remove an order
* [UpdateAccountingOrder](#updateaccountingorder) - Update an order
* [UpdateAssessmentOrder](#updateassessmentorder) - Update an order

## CreateAccountingOrder

Create an order

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
                Address1: unifiedgosdk.Pointer("802 Bechtelar Park"),
                Address2: unifiedgosdk.Pointer("Apt. 436"),
                City: unifiedgosdk.Pointer("Daniellaville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("36947"),
                Region: unifiedgosdk.Pointer("Wisconsin"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("77e39f9e-6f44-41d4-9e9d-e7ecb27a1543"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2022-04-05T00:28:38.595Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("9745 Betty Shore"),
                City: unifiedgosdk.Pointer("South Alainaland"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("25274-7654"),
                Region: unifiedgosdk.Pointer("New Hampshire"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            Status: shared.AccountingOrderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypePurchase.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-17T22:46:34.526Z"),
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

Create an order

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
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("d2c7a88c-4973-4f3a-977c-36e91d6bbb66"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.501Z"),
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

## GetAccountingOrder

Retrieve an order

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

Retrieve an order

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

## ListAccountingOrders

List all orders

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

## PatchAccountingOrder

Update an order

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
                Address1: unifiedgosdk.Pointer("802 Bechtelar Park"),
                Address2: unifiedgosdk.Pointer("Apt. 436"),
                City: unifiedgosdk.Pointer("Daniellaville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("36947"),
                Region: unifiedgosdk.Pointer("Wisconsin"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("d4016fb5-81f7-4acc-af53-e4ff827d86a3"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2022-04-05T00:28:38.601Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("9745 Betty Shore"),
                City: unifiedgosdk.Pointer("South Alainaland"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("25274-7654"),
                Region: unifiedgosdk.Pointer("New Hampshire"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            Status: shared.AccountingOrderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypePurchase.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-17T22:46:34.529Z"),
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

Update an order

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
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("ab8d64eb-a6c2-4128-a202-df4bd71d26a7"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.507Z"),
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

## RemoveAccountingOrder

Remove an order

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

## UpdateAccountingOrder

Update an order

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
                Address1: unifiedgosdk.Pointer("802 Bechtelar Park"),
                Address2: unifiedgosdk.Pointer("Apt. 436"),
                City: unifiedgosdk.Pointer("Daniellaville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("36947"),
                Region: unifiedgosdk.Pointer("Wisconsin"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-11-20T03:46:49.837Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("d4016fb5-81f7-4acc-af53-e4ff827d86a3"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2022-04-05T00:28:38.601Z"),
            ShippingAddress: &shared.PropertyAccountingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("9745 Betty Shore"),
                City: unifiedgosdk.Pointer("South Alainaland"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("25274-7654"),
                Region: unifiedgosdk.Pointer("New Hampshire"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            Status: shared.AccountingOrderStatusSubmitted.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingOrderTypePurchase.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-17T22:46:34.529Z"),
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

Update an order

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
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("ab8d64eb-a6c2-4128-a202-df4bd71d26a7"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.507Z"),
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