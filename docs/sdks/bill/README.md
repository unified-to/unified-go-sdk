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

Create a bill

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
                    ID: unifiedgosdk.Pointer("e9be1609-90c4-48d9-bcb6-bb4eba454160"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("4c9f7b6e-eb59-4c60-8da2-a80ddd62b2df"),
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
            ID: unifiedgosdk.Pointer("5243622c-83c2-4232-979e-c8e131d0e6a8"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-06-01T03:07:18.326Z"),
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

Retrieve a bill

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

List all bills

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

Update a bill

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
                    ID: unifiedgosdk.Pointer("e2c60758-75c2-4c9c-8857-0a6582ab0e6c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("0b29040b-8723-4e46-8e8a-be10c3802b8f"),
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
            ID: unifiedgosdk.Pointer("b9408c3b-2328-4dbe-b437-573c4bded241"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-06-01T03:07:18.335Z"),
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

Remove a bill

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

Update a bill

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
                    ID: unifiedgosdk.Pointer("e2c60758-75c2-4c9c-8857-0a6582ab0e6c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("BILL-2466.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://grimy-characterization.net"),
                    ID: unifiedgosdk.Pointer("0b29040b-8723-4e46-8e8a-be10c3802b8f"),
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
            ID: unifiedgosdk.Pointer("b9408c3b-2328-4dbe-b437-573c4bded241"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-06-01T03:07:18.335Z"),
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