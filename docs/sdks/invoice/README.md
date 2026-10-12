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
                    ID: unifiedgosdk.Pointer("0b2b5623-c86c-48b6-9b52-3032f50b615c"),
                    MimeType: unifiedgosdk.Pointer("benevolentia"),
                    Name: unifiedgosdk.Pointer("vespillo"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](-1.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("RWF"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-11-27T21:25:37.363Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("b9e6d33c-057c-4adf-b443-a7e1a62487fa"),
            InvoiceNumber: unifiedgosdk.Pointer("vinco"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(true),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Auctus comburo clarus ubi."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2022-11-25T15:00:28.871Z"),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodSendInvoice.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2026-04-16T20:56:16.977Z"),
            Reference: unifiedgosdk.Pointer("adinventitias"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusDeleted.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet45.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeCreditmemo.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T19:40:35.053Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/"),
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
                    ID: unifiedgosdk.Pointer("27b32839-b10b-4e14-ba69-3a6390e5e1e9"),
                    MimeType: unifiedgosdk.Pointer("benevolentia"),
                    Name: unifiedgosdk.Pointer("vespillo"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](-1.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("RWF"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-11-27T21:25:37.363Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("78a173f7-dafa-4b6d-9a2b-e84bc65a5176"),
            InvoiceNumber: unifiedgosdk.Pointer("vinco"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(true),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Auctus comburo clarus ubi."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2022-11-25T15:00:28.871Z"),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodSendInvoice.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2026-04-16T20:56:17.032Z"),
            Reference: unifiedgosdk.Pointer("adinventitias"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusDeleted.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet45.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeCreditmemo.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T19:40:35.057Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/"),
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
                    ID: unifiedgosdk.Pointer("27b32839-b10b-4e14-ba69-3a6390e5e1e9"),
                    MimeType: unifiedgosdk.Pointer("benevolentia"),
                    Name: unifiedgosdk.Pointer("vespillo"),
                },
            },
            BalanceAmount: unifiedgosdk.Pointer[float64](-1.0),
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2022-11-07T14:17:29.587Z"),
            Currency: unifiedgosdk.Pointer("RWF"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2022-11-27T21:25:37.363Z"),
            ExtendedNotes: []shared.AccountingExtendedNote{},
            ID: unifiedgosdk.Pointer("78a173f7-dafa-4b6d-9a2b-e84bc65a5176"),
            InvoiceNumber: unifiedgosdk.Pointer("vinco"),
            IsInclusiveOfTax: unifiedgosdk.Pointer(true),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Auctus comburo clarus ubi."),
            PaidAmount: unifiedgosdk.Pointer[float64](0.0),
            PaidAt: types.MustNewTimeFromString("2022-11-25T15:00:28.871Z"),
            PaymentCollectionMethod: shared.AccountingInvoicePaymentCollectionMethodSendInvoice.ToPointer(),
            Payments: []shared.AccountingPaymentReference{},
            PostedAt: types.MustNewTimeFromString("2026-04-16T20:56:17.032Z"),
            Reference: unifiedgosdk.Pointer("adinventitias"),
            SalesorderIds: []string{},
            Send: unifiedgosdk.Pointer(true),
            Status: shared.AccountingInvoiceStatusDeleted.ToPointer(),
            SubTotalAmount: unifiedgosdk.Pointer[float64](0.0),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Term: shared.AccountingInvoiceTermNet45.ToPointer(),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            Type: shared.AccountingInvoiceTypeCreditmemo.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T19:40:35.057Z"),
            URL: unifiedgosdk.Pointer("https://gifted-yarmulke.info/"),
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