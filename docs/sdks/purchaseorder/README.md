# Purchaseorder

## Overview

### Available Operations

* [CreateAccountingPurchaseorder](#createaccountingpurchaseorder) - Create a purchaseorder
* [GetAccountingPurchaseorder](#getaccountingpurchaseorder) - Retrieve a purchaseorder
* [ListAccountingPurchaseorders](#listaccountingpurchaseorders) - List all purchaseorders
* [PatchAccountingPurchaseorder](#patchaccountingpurchaseorder) - Update a purchaseorder
* [RemoveAccountingPurchaseorder](#removeaccountingpurchaseorder) - Remove a purchaseorder
* [UpdateAccountingPurchaseorder](#updateaccountingpurchaseorder) - Update a purchaseorder

## CreateAccountingPurchaseorder

Create a purchaseorder

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingPurchaseorder" method="post" path="/accounting/{connection_id}/purchaseorder" example="accounting_purchaseorder" -->
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

    res, err := s.Purchaseorder.CreateAccountingPurchaseorder(ctx, operations.CreateAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Tanya Walks"),
                City: unifiedgosdk.Pointer("South Annabelleton"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("30337"),
                Region: unifiedgosdk.Pointer("Nevada"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("ZMW"),
            ID: unifiedgosdk.Pointer("e4b77728-8330-459c-9a39-260ab69e315c"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-25T20:27:36.971Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("649 Maggio Overpass"),
                City: unifiedgosdk.Pointer("Lake Jaylan"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("99211-6547"),
                Region: unifiedgosdk.Pointer("North Carolina"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            Status: shared.AccountingPurchaseorderStatusPartiallyRefunded.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-02-26T04:07:06.164Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateAccountingPurchaseorderRequest](../../pkg/models/operations/createaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateAccountingPurchaseorderResponse](../../pkg/models/operations/createaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingPurchaseorder

Retrieve a purchaseorder

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingPurchaseorder" method="get" path="/accounting/{connection_id}/purchaseorder/{id}" -->
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

    res, err := s.Purchaseorder.GetAccountingPurchaseorder(ctx, operations.GetAccountingPurchaseorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAccountingPurchaseorderRequest](../../pkg/models/operations/getaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetAccountingPurchaseorderResponse](../../pkg/models/operations/getaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingPurchaseorders

List all purchaseorders

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingPurchaseorders" method="get" path="/accounting/{connection_id}/purchaseorder" -->
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

    res, err := s.Purchaseorder.ListAccountingPurchaseorders(ctx, operations.ListAccountingPurchaseordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListAccountingPurchaseordersRequest](../../pkg/models/operations/listaccountingpurchaseordersrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListAccountingPurchaseordersResponse](../../pkg/models/operations/listaccountingpurchaseordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingPurchaseorder

Update a purchaseorder

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingPurchaseorder" method="patch" path="/accounting/{connection_id}/purchaseorder/{id}" example="accounting_purchaseorder" -->
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

    res, err := s.Purchaseorder.PatchAccountingPurchaseorder(ctx, operations.PatchAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Tanya Walks"),
                City: unifiedgosdk.Pointer("South Annabelleton"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("30337"),
                Region: unifiedgosdk.Pointer("Nevada"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("ZMW"),
            ID: unifiedgosdk.Pointer("607bbb4e-9c94-4a95-ab51-72578207bf03"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-25T20:27:37.016Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("649 Maggio Overpass"),
                City: unifiedgosdk.Pointer("Lake Jaylan"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("99211-6547"),
                Region: unifiedgosdk.Pointer("North Carolina"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            Status: shared.AccountingPurchaseorderStatusPartiallyRefunded.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-02-26T04:07:06.167Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchAccountingPurchaseorderRequest](../../pkg/models/operations/patchaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchAccountingPurchaseorderResponse](../../pkg/models/operations/patchaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingPurchaseorder

Remove a purchaseorder

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingPurchaseorder" method="delete" path="/accounting/{connection_id}/purchaseorder/{id}" -->
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

    res, err := s.Purchaseorder.RemoveAccountingPurchaseorder(ctx, operations.RemoveAccountingPurchaseorderRequest{
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.RemoveAccountingPurchaseorderRequest](../../pkg/models/operations/removeaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.RemoveAccountingPurchaseorderResponse](../../pkg/models/operations/removeaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingPurchaseorder

Update a purchaseorder

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingPurchaseorder" method="put" path="/accounting/{connection_id}/purchaseorder/{id}" example="accounting_purchaseorder" -->
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

    res, err := s.Purchaseorder.UpdateAccountingPurchaseorder(ctx, operations.UpdateAccountingPurchaseorderRequest{
        AccountingPurchaseorder: shared.AccountingPurchaseorder{
            BillingAddress: &shared.PropertyAccountingPurchaseorderBillingAddress{
                Address1: unifiedgosdk.Pointer("37214 Tanya Walks"),
                City: unifiedgosdk.Pointer("South Annabelleton"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("30337"),
                Region: unifiedgosdk.Pointer("Nevada"),
                RegionCode: unifiedgosdk.Pointer("MA"),
            },
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-12-12T07:17:47.021Z"),
            Currency: unifiedgosdk.Pointer("ZMW"),
            ID: unifiedgosdk.Pointer("607bbb4e-9c94-4a95-ab51-72578207bf03"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            PostedAt: types.MustNewTimeFromString("2025-04-25T20:27:37.016Z"),
            ShippingAddress: &shared.PropertyAccountingPurchaseorderShippingAddress{
                Address1: unifiedgosdk.Pointer("649 Maggio Overpass"),
                City: unifiedgosdk.Pointer("Lake Jaylan"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("99211-6547"),
                Region: unifiedgosdk.Pointer("North Carolina"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            Status: shared.AccountingPurchaseorderStatusPartiallyRefunded.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2021-02-26T04:07:06.167Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingPurchaseorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateAccountingPurchaseorderRequest](../../pkg/models/operations/updateaccountingpurchaseorderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateAccountingPurchaseorderResponse](../../pkg/models/operations/updateaccountingpurchaseorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |