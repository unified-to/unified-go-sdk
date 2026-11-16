# Bill

## Overview

### Available Operations

* [CreateAccountingBill](#createaccountingbill) - Create a bill
* [GetAccountingBill](#getaccountingbill) - Retrieve a bill
* [ListAccountingBills](#listaccountingbills) - List all bills
* [PatchAccountingBill](#patchaccountingbill) - Update a bill
* [RemoveAccountingBill](#removeaccountingbill) - Remove a bill
* [UpdateAccountingBill](#updateaccountingbill) - Update a bill

## CreateAccountingBill

Creates a AccountingBill. Provide the unified AccountingBill fields in the request body; they are mapped to the provider's API. Returns the created AccountingBill.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingBill" method="post" path="/accounting/{connection_id}/bill" example="accounting_bill" -->
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

    res, err := s.Bill.CreateAccountingBill(ctx, operations.CreateAccountingBillRequest{
        AccountingBill: shared.AccountingBill{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://puny-foodstuffs.name"),
                    ID: unifiedgosdk.Pointer("ae6ff06a-1ab0-4bfc-ab21-bdb787038982"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("8557a9d0-b757-4d46-bef8-700b8ae2fa07"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            BillNumber: unifiedgosdk.Pointer("BILL-2466"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2019-09-07T23:03:14.104Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("ecd751b5-5db1-4e6f-8984-dad240aec28b"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Admiratio deporto tutamen cilicium infit amissio ante demulceo."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2019-08-15T17:24:28.313Z"),
            PaymentCollectionMethod: shared.PaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            PurchaseorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingBillStatusPartiallyPaid.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.TermNet90.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2020-06-04T02:54:10.246Z"),
            URL: unifiedgosdk.Pointer("https://cluttered-importance.org/bills/BILL-2466"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateAccountingBillRequest](../../pkg/models/operations/createaccountingbillrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateAccountingBillResponse](../../pkg/models/operations/createaccountingbillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingBill

Retrieves a single AccountingBill by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingBill" method="get" path="/accounting/{connection_id}/bill/{id}" -->
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

    res, err := s.Bill.GetAccountingBill(ctx, operations.GetAccountingBillRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAccountingBillRequest](../../pkg/models/operations/getaccountingbillrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetAccountingBillResponse](../../pkg/models/operations/getaccountingbillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingBills

Returns a list of AccountingBills. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingBills" method="get" path="/accounting/{connection_id}/bill" -->
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

    res, err := s.Bill.ListAccountingBills(ctx, operations.ListAccountingBillsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBills != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAccountingBillsRequest](../../pkg/models/operations/listaccountingbillsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAccountingBillsResponse](../../pkg/models/operations/listaccountingbillsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingBill

Updates the AccountingBill identified by `id`. Provide the unified AccountingBill fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingBill" method="patch" path="/accounting/{connection_id}/bill/{id}" example="accounting_bill" -->
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

    res, err := s.Bill.PatchAccountingBill(ctx, operations.PatchAccountingBillRequest{
        AccountingBill: shared.AccountingBill{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://puny-foodstuffs.name"),
                    ID: unifiedgosdk.Pointer("e0366807-5fcf-4743-b139-0b391a8aef8c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("23508760-298c-4dbd-a166-9e2ab827ac9f"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            BillNumber: unifiedgosdk.Pointer("BILL-2466"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2019-09-07T23:03:14.104Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("b11d39d9-80ba-4481-a5ec-718182f7ad97"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Admiratio deporto tutamen cilicium infit amissio ante demulceo."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2019-08-15T17:24:28.313Z"),
            PaymentCollectionMethod: shared.PaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            PurchaseorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingBillStatusPartiallyPaid.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.TermNet90.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2020-06-04T02:54:10.252Z"),
            URL: unifiedgosdk.Pointer("https://cluttered-importance.org/bills/BILL-2466"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchAccountingBillRequest](../../pkg/models/operations/patchaccountingbillrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchAccountingBillResponse](../../pkg/models/operations/patchaccountingbillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingBill

Deletes the AccountingBill identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingBill" method="delete" path="/accounting/{connection_id}/bill/{id}" -->
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

    res, err := s.Bill.RemoveAccountingBill(ctx, operations.RemoveAccountingBillRequest{
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
| `request`                                                                                            | [operations.RemoveAccountingBillRequest](../../pkg/models/operations/removeaccountingbillrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveAccountingBillResponse](../../pkg/models/operations/removeaccountingbillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingBill

Updates the AccountingBill identified by `id`. Provide the unified AccountingBill fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingBill" method="put" path="/accounting/{connection_id}/bill/{id}" example="accounting_bill" -->
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

    res, err := s.Bill.UpdateAccountingBill(ctx, operations.UpdateAccountingBillRequest{
        AccountingBill: shared.AccountingBill{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://puny-foodstuffs.name"),
                    ID: unifiedgosdk.Pointer("e0366807-5fcf-4743-b139-0b391a8aef8c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("23508760-298c-4dbd-a166-9e2ab827ac9f"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            BillNumber: unifiedgosdk.Pointer("BILL-2466"),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2019-09-07T23:03:14.104Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("b11d39d9-80ba-4481-a5ec-718182f7ad97"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Admiratio deporto tutamen cilicium infit amissio ante demulceo."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2019-08-15T17:24:28.313Z"),
            PaymentCollectionMethod: shared.PaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2019-08-08T23:03:14.104Z"),
            PurchaseorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingBillStatusPartiallyPaid.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.TermNet90.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2020-06-04T02:54:10.252Z"),
            URL: unifiedgosdk.Pointer("https://cluttered-importance.org/bills/BILL-2466"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingBill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateAccountingBillRequest](../../pkg/models/operations/updateaccountingbillrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateAccountingBillResponse](../../pkg/models/operations/updateaccountingbillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |