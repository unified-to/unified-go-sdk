# Invoice

## Overview

### Available Operations

* [CreateAccountingInvoice](#createaccountinginvoice) - Create an invoice
* [GetAccountingInvoice](#getaccountinginvoice) - Retrieve an invoice
* [ListAccountingInvoices](#listaccountinginvoices) - List all invoices
* [PatchAccountingInvoice](#patchaccountinginvoice) - Update an invoice
* [RemoveAccountingInvoice](#removeaccountinginvoice) - Remove an invoice
* [UpdateAccountingInvoice](#updateaccountinginvoice) - Update an invoice

## CreateAccountingInvoice

Create an invoice

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingInvoice" method="post" path="/accounting/{connection_id}/invoice" example="accounting_invoice" -->
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

    res, err := s.Invoice.CreateAccountingInvoice(ctx, operations.CreateAccountingInvoiceRequest{
        AccountingInvoice: shared.AccountingInvoice{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://glossy-markup.net/"),
                    ID: unifiedgosdk.Pointer("8785a531-c170-4e0a-a23b-6026aeb2808f"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("795784ec-1677-4171-87a9-191b7246647a"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-12-07T14:17:29.587Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("96f21cc2-029b-451f-bf63-27b0da867a12"),
            InvoiceNumber: unifiedgosdk.Pointer("INV-77705"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(false),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Adinventitias sursum temptatio auctus comburo clarus ubi deprimo."),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Reference: unifiedgosdk.Pointer("PO-2331"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusOverdue.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet25.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-23T13:03:42.762Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/invoices/INV-77705"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingInvoiceRequest](../../pkg/models/operations/createaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingInvoiceResponse](../../pkg/models/operations/createaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingInvoice

Retrieve an invoice

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingInvoice" method="get" path="/accounting/{connection_id}/invoice/{id}" -->
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

    res, err := s.Invoice.GetAccountingInvoice(ctx, operations.GetAccountingInvoiceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountingInvoiceRequest](../../pkg/models/operations/getaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountingInvoiceResponse](../../pkg/models/operations/getaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingInvoices

List all invoices

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingInvoices" method="get" path="/accounting/{connection_id}/invoice" -->
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

    res, err := s.Invoice.ListAccountingInvoices(ctx, operations.ListAccountingInvoicesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingInvoices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAccountingInvoicesRequest](../../pkg/models/operations/listaccountinginvoicesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAccountingInvoicesResponse](../../pkg/models/operations/listaccountinginvoicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingInvoice

Update an invoice

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingInvoice" method="patch" path="/accounting/{connection_id}/invoice/{id}" example="accounting_invoice" -->
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

    res, err := s.Invoice.PatchAccountingInvoice(ctx, operations.PatchAccountingInvoiceRequest{
        AccountingInvoice: shared.AccountingInvoice{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://glossy-markup.net/"),
                    ID: unifiedgosdk.Pointer("21a65956-7898-49fe-954e-657be9079de1"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("0de07f5c-0422-4bf1-8312-b38a6e63a87d"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-12-07T14:17:29.587Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("896edd87-b896-4ec1-9035-22e2a216a31b"),
            InvoiceNumber: unifiedgosdk.Pointer("INV-77705"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(false),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Adinventitias sursum temptatio auctus comburo clarus ubi deprimo."),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Reference: unifiedgosdk.Pointer("PO-2331"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusOverdue.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet25.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-23T13:03:42.816Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/invoices/INV-77705"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAccountingInvoiceRequest](../../pkg/models/operations/patchaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAccountingInvoiceResponse](../../pkg/models/operations/patchaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingInvoice

Remove an invoice

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingInvoice" method="delete" path="/accounting/{connection_id}/invoice/{id}" -->
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

    res, err := s.Invoice.RemoveAccountingInvoice(ctx, operations.RemoveAccountingInvoiceRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAccountingInvoiceRequest](../../pkg/models/operations/removeaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAccountingInvoiceResponse](../../pkg/models/operations/removeaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingInvoice

Update an invoice

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingInvoice" method="put" path="/accounting/{connection_id}/invoice/{id}" example="accounting_invoice" -->
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

    res, err := s.Invoice.UpdateAccountingInvoice(ctx, operations.UpdateAccountingInvoiceRequest{
        AccountingInvoice: shared.AccountingInvoice{
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://glossy-markup.net/"),
                    ID: unifiedgosdk.Pointer("21a65956-7898-49fe-954e-657be9079de1"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("0de07f5c-0422-4bf1-8312-b38a6e63a87d"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705-1.pdf"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](0.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-12-07T14:17:29.587Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("896edd87-b896-4ec1-9035-22e2a216a31b"),
            InvoiceNumber: unifiedgosdk.Pointer("INV-77705"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(false),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Adinventitias sursum temptatio auctus comburo clarus ubi deprimo."),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodChargeAutomatically.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Reference: unifiedgosdk.Pointer("PO-2331"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusOverdue.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet25.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeInvoice.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-23T13:03:42.816Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/invoices/INV-77705"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAccountingInvoiceRequest](../../pkg/models/operations/updateaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAccountingInvoiceResponse](../../pkg/models/operations/updateaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |