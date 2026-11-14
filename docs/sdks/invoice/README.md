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

Creates a AccountingInvoice. Provide the unified AccountingInvoice fields in the request body; they are mapped to the provider's API. Returns the created AccountingInvoice.

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
                    ID: unifiedgosdk.Pointer("e9e921cc-902e-44cb-a6e1-4750e6866f6e"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("09a04c74-48cd-476a-8efb-45d77c06f7ee"),
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
            ID: unifiedgosdk.Pointer("633ebed7-a4c7-4d51-a661-6f24162329e8"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T13:40:57.767Z"),
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

Retrieves a single AccountingInvoice by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of AccountingInvoices. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the AccountingInvoice identified by `id`. Provide the unified AccountingInvoice fields to change in the request body.

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
                    ID: unifiedgosdk.Pointer("dbba6a13-fc62-4ded-8d70-9850ad27614c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("86f0e4c0-41cc-45f8-a7bd-13579e2fed80"),
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
            ID: unifiedgosdk.Pointer("8395be39-d3e8-4e35-9237-f9f84e2a178b"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T13:40:57.800Z"),
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

Deletes the AccountingInvoice identified by `id`.

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

Updates the AccountingInvoice identified by `id`. Provide the unified AccountingInvoice fields to change in the request body.

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
                    ID: unifiedgosdk.Pointer("dbba6a13-fc62-4ded-8d70-9850ad27614c"),
                    MimeType: unifiedgosdk.Pointer("application/pdf"),
                    Name: unifiedgosdk.Pointer("INV-77705.pdf"),
                },
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://earnest-giggle.net"),
                    ID: unifiedgosdk.Pointer("86f0e4c0-41cc-45f8-a7bd-13579e2fed80"),
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
            ID: unifiedgosdk.Pointer("8395be39-d3e8-4e35-9237-f9f84e2a178b"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T13:40:57.800Z"),
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