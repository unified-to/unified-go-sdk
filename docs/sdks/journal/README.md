# Journal
(*Journal*)

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

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.CreateAccountingJournalRequest{
        ConnectionID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Journal.CreateAccountingJournal(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAccountingJournal

Retrieve a journal

### Example Usage

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.GetAccountingJournalRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Journal.GetAccountingJournal(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListAccountingJournals

List all journals

### Example Usage

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.ListAccountingJournalsRequest{
        ConnectionID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Journal.ListAccountingJournals(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchAccountingJournal

Update a journal

### Example Usage

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.PatchAccountingJournalRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Journal.PatchAccountingJournal(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveAccountingJournal

Remove a journal

### Example Usage

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.RemoveAccountingJournalRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Journal.RemoveAccountingJournal(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateAccountingJournal

Update a journal

### Example Usage

```go
package main

import(
	"os"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(os.Getenv("JWT")),
    )
    request := operations.UpdateAccountingJournalRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Journal.UpdateAccountingJournal(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
