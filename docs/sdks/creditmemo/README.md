# Creditmemo

## Overview

### Available Operations

* [CreateAccountingCreditmemo](#createaccountingcreditmemo) - Create a creditmemo
* [GetAccountingCreditmemo](#getaccountingcreditmemo) - Retrieve a creditmemo
* [ListAccountingCreditmemoes](#listaccountingcreditmemoes) - List all creditmemoes
* [PatchAccountingCreditmemo](#patchaccountingcreditmemo) - Update a creditmemo
* [RemoveAccountingCreditmemo](#removeaccountingcreditmemo) - Remove a creditmemo
* [UpdateAccountingCreditmemo](#updateaccountingcreditmemo) - Update a creditmemo

## CreateAccountingCreditmemo

Create a creditmemo

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingCreditmemo" method="post" path="/accounting/{connection_id}/creditmemo" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.CreateAccountingCreditmemo(ctx, operations.CreateAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://enlightened-chairperson.com/"),
                    ID: unifiedgosdk.Pointer("1488ad93-84a6-49ed-932e-b13b61f4aec9"),
                    MimeType: unifiedgosdk.Pointer("complectus"),
                    Name: unifiedgosdk.Pointer("thesis"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("ulterius"),
            Currency: unifiedgosdk.Pointer("MKD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2023-10-18T04:35:00.543Z"),
            ID: unifiedgosdk.Pointer("aea99f00-408b-4a75-85a6-157c219f02cf"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Dedecor amo adfero torqueo quas."),
            PaymentCollectionMethod: shared.AccountingCreditmemoPaymentCollectionMethodChargeAutomatically.ToPointer(),
            PostedAt: types.MustNewTimeFromString("2025-11-15T11:46:05.682Z"),
            RefundAmount: unifiedgosdk.Pointer[float64](0.0),
            RefundReason: unifiedgosdk.Pointer("Virgo inflammatio quibusdam aestivus magnam."),
            RefundedAt: types.MustNewTimeFromString("2023-10-23T00:35:36.814Z"),
            Send: unifiedgosdk.Pointer(false),
            Status: shared.AccountingCreditmemoStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2024-11-15T13:55:49.062Z"),
            URL: unifiedgosdk.Pointer("https://lighthearted-bandwidth.net/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingCreditmemoRequest](../../pkg/models/operations/createaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateAccountingCreditmemoResponse](../../pkg/models/operations/createaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingCreditmemo

Retrieve a creditmemo

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingCreditmemo" method="get" path="/accounting/{connection_id}/creditmemo/{id}" -->
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

    res, err := s.Creditmemo.GetAccountingCreditmemo(ctx, operations.GetAccountingCreditmemoRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetAccountingCreditmemoRequest](../../pkg/models/operations/getaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAccountingCreditmemoResponse](../../pkg/models/operations/getaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingCreditmemoes

List all creditmemoes

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingCreditmemoes" method="get" path="/accounting/{connection_id}/creditmemo" -->
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

    res, err := s.Creditmemo.ListAccountingCreditmemoes(ctx, operations.ListAccountingCreditmemoesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemoes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListAccountingCreditmemoesRequest](../../pkg/models/operations/listaccountingcreditmemoesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListAccountingCreditmemoesResponse](../../pkg/models/operations/listaccountingcreditmemoesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingCreditmemo

Update a creditmemo

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingCreditmemo" method="patch" path="/accounting/{connection_id}/creditmemo/{id}" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.PatchAccountingCreditmemo(ctx, operations.PatchAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://enlightened-chairperson.com/"),
                    ID: unifiedgosdk.Pointer("92ebacd2-bd76-47ed-a2c3-8354d39cbb0d"),
                    MimeType: unifiedgosdk.Pointer("complectus"),
                    Name: unifiedgosdk.Pointer("thesis"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("ulterius"),
            Currency: unifiedgosdk.Pointer("MKD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2023-10-18T04:35:00.543Z"),
            ID: unifiedgosdk.Pointer("f0ec1da8-b326-4f28-a43e-b4ff7e926358"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Dedecor amo adfero torqueo quas."),
            PaymentCollectionMethod: shared.AccountingCreditmemoPaymentCollectionMethodChargeAutomatically.ToPointer(),
            PostedAt: types.MustNewTimeFromString("2025-11-15T11:46:05.704Z"),
            RefundAmount: unifiedgosdk.Pointer[float64](0.0),
            RefundReason: unifiedgosdk.Pointer("Virgo inflammatio quibusdam aestivus magnam."),
            RefundedAt: types.MustNewTimeFromString("2023-10-23T00:35:36.814Z"),
            Send: unifiedgosdk.Pointer(false),
            Status: shared.AccountingCreditmemoStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2024-11-15T13:55:49.074Z"),
            URL: unifiedgosdk.Pointer("https://lighthearted-bandwidth.net/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchAccountingCreditmemoRequest](../../pkg/models/operations/patchaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PatchAccountingCreditmemoResponse](../../pkg/models/operations/patchaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingCreditmemo

Remove a creditmemo

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingCreditmemo" method="delete" path="/accounting/{connection_id}/creditmemo/{id}" -->
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

    res, err := s.Creditmemo.RemoveAccountingCreditmemo(ctx, operations.RemoveAccountingCreditmemoRequest{
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
| `request`                                                                                                        | [operations.RemoveAccountingCreditmemoRequest](../../pkg/models/operations/removeaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.RemoveAccountingCreditmemoResponse](../../pkg/models/operations/removeaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingCreditmemo

Update a creditmemo

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingCreditmemo" method="put" path="/accounting/{connection_id}/creditmemo/{id}" example="accounting_creditmemo" -->
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

    res, err := s.Creditmemo.UpdateAccountingCreditmemo(ctx, operations.UpdateAccountingCreditmemoRequest{
        AccountingCreditmemo: shared.AccountingCreditmemo{
            Applications: []shared.AccountingCreditApplication{},
            Attachments: []shared.AccountingAttachment{
                shared.AccountingAttachment{
                    DownloadURL: unifiedgosdk.Pointer("https://enlightened-chairperson.com/"),
                    ID: unifiedgosdk.Pointer("92ebacd2-bd76-47ed-a2c3-8354d39cbb0d"),
                    MimeType: unifiedgosdk.Pointer("complectus"),
                    Name: unifiedgosdk.Pointer("thesis"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-09-20T01:47:01.571Z"),
            CreditmemoNumber: unifiedgosdk.Pointer("ulterius"),
            Currency: unifiedgosdk.Pointer("MKD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            DueAt: types.MustNewTimeFromString("2023-10-18T04:35:00.543Z"),
            ID: unifiedgosdk.Pointer("f0ec1da8-b326-4f28-a43e-b4ff7e926358"),
            Lineitems: []shared.AccountingLineitem{},
            Metadata: []shared.AccountingMetadata{},
            Notes: unifiedgosdk.Pointer("Dedecor amo adfero torqueo quas."),
            PaymentCollectionMethod: shared.AccountingCreditmemoPaymentCollectionMethodChargeAutomatically.ToPointer(),
            PostedAt: types.MustNewTimeFromString("2025-11-15T11:46:05.704Z"),
            RefundAmount: unifiedgosdk.Pointer[float64](0.0),
            RefundReason: unifiedgosdk.Pointer("Virgo inflammatio quibusdam aestivus magnam."),
            RefundedAt: types.MustNewTimeFromString("2023-10-23T00:35:36.814Z"),
            Send: unifiedgosdk.Pointer(false),
            Status: shared.AccountingCreditmemoStatusPaid.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            TotalAmount: unifiedgosdk.Pointer[float64](0.0),
            UpdatedAt: types.MustNewTimeFromString("2024-11-15T13:55:49.074Z"),
            URL: unifiedgosdk.Pointer("https://lighthearted-bandwidth.net/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreditmemo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateAccountingCreditmemoRequest](../../pkg/models/operations/updateaccountingcreditmemorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateAccountingCreditmemoResponse](../../pkg/models/operations/updateaccountingcreditmemoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |