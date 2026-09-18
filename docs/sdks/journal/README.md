# Journal

## Overview

### Available Operations

* [CreateAccountingJournal](#createaccountingjournal) - Create a journal
* [GetAccountingJournal](#getaccountingjournal) - Retrieve a journal
* [ListAccountingJournals](#listaccountingjournals) - List all journals
* [PatchAccountingJournal](#patchaccountingjournal) - Update a journal
* [RemoveAccountingJournal](#removeaccountingjournal) - Remove a journal
* [UpdateAccountingJournal](#updateaccountingjournal) - Update a journal

## CreateAccountingJournal

Create a journal

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingJournal" method="post" path="/accounting/{connection_id}/journal" example="accounting_journal" -->
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

    res, err := s.Journal.CreateAccountingJournal(ctx, operations.CreateAccountingJournalRequest{
        AccountingJournal: shared.AccountingJournal{
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-02-20T15:14:55.881Z"),
            Currency: unifiedgosdk.Pointer("FKP"),
            Description: unifiedgosdk.Pointer("Calco constans adipisci."),
            ID: unifiedgosdk.Pointer("3d0b7527-bb53-4d13-9fe9-a4e3437f68e6"),
            PostedAt: types.MustNewTimeFromString("2023-10-19T02:25:22.678Z"),
            Reference: unifiedgosdk.Pointer("ullam"),
            Source: unifiedgosdk.Pointer("crustulum"),
            TaxAmount: unifiedgosdk.Pointer[float64](78672.0),
            UpdatedAt: types.MustNewTimeFromString("2022-01-01T11:25:54.882Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingJournalRequest](../../pkg/models/operations/createaccountingjournalrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingJournalResponse](../../pkg/models/operations/createaccountingjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingJournal

Retrieve a journal

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingJournal" method="get" path="/accounting/{connection_id}/journal/{id}" -->
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

    res, err := s.Journal.GetAccountingJournal(ctx, operations.GetAccountingJournalRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountingJournalRequest](../../pkg/models/operations/getaccountingjournalrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountingJournalResponse](../../pkg/models/operations/getaccountingjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingJournals

List all journals

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingJournals" method="get" path="/accounting/{connection_id}/journal" -->
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

    res, err := s.Journal.ListAccountingJournals(ctx, operations.ListAccountingJournalsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingJournals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAccountingJournalsRequest](../../pkg/models/operations/listaccountingjournalsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAccountingJournalsResponse](../../pkg/models/operations/listaccountingjournalsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingJournal

Update a journal

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingJournal" method="patch" path="/accounting/{connection_id}/journal/{id}" example="accounting_journal" -->
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

    res, err := s.Journal.PatchAccountingJournal(ctx, operations.PatchAccountingJournalRequest{
        AccountingJournal: shared.AccountingJournal{
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-02-20T15:14:55.881Z"),
            Currency: unifiedgosdk.Pointer("FKP"),
            Description: unifiedgosdk.Pointer("Calco constans adipisci."),
            ID: unifiedgosdk.Pointer("08ebd0be-97e9-4a3d-b153-e0a6d981d560"),
            PostedAt: types.MustNewTimeFromString("2023-10-19T02:25:22.686Z"),
            Reference: unifiedgosdk.Pointer("ullam"),
            Source: unifiedgosdk.Pointer("crustulum"),
            TaxAmount: unifiedgosdk.Pointer[float64](78672.0),
            UpdatedAt: types.MustNewTimeFromString("2022-01-01T11:25:54.885Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAccountingJournalRequest](../../pkg/models/operations/patchaccountingjournalrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAccountingJournalResponse](../../pkg/models/operations/patchaccountingjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingJournal

Remove a journal

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingJournal" method="delete" path="/accounting/{connection_id}/journal/{id}" -->
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

    res, err := s.Journal.RemoveAccountingJournal(ctx, operations.RemoveAccountingJournalRequest{
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
| `request`                                                                                                  | [operations.RemoveAccountingJournalRequest](../../pkg/models/operations/removeaccountingjournalrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAccountingJournalResponse](../../pkg/models/operations/removeaccountingjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingJournal

Update a journal

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingJournal" method="put" path="/accounting/{connection_id}/journal/{id}" example="accounting_journal" -->
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

    res, err := s.Journal.UpdateAccountingJournal(ctx, operations.UpdateAccountingJournalRequest{
        AccountingJournal: shared.AccountingJournal{
            Attachments: []shared.AccountingAttachment{},
            CategoryIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-02-20T15:14:55.881Z"),
            Currency: unifiedgosdk.Pointer("FKP"),
            Description: unifiedgosdk.Pointer("Calco constans adipisci."),
            ID: unifiedgosdk.Pointer("08ebd0be-97e9-4a3d-b153-e0a6d981d560"),
            PostedAt: types.MustNewTimeFromString("2023-10-19T02:25:22.686Z"),
            Reference: unifiedgosdk.Pointer("ullam"),
            Source: unifiedgosdk.Pointer("crustulum"),
            TaxAmount: unifiedgosdk.Pointer[float64](78672.0),
            UpdatedAt: types.MustNewTimeFromString("2022-01-01T11:25:54.885Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAccountingJournalRequest](../../pkg/models/operations/updateaccountingjournalrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAccountingJournalResponse](../../pkg/models/operations/updateaccountingjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |