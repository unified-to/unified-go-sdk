# Note
(*Note*)

### Available Operations

* [DeleteTicketingConnectionIDNoteTicketIDID](#deleteticketingconnectionidnoteticketidid) - Remove a note
* [GetTicketingConnectionIDNoteTicketID](#getticketingconnectionidnoteticketid) - List all notes
* [GetTicketingConnectionIDNoteTicketIDID](#getticketingconnectionidnoteticketidid) - Retrieve a note
* [PatchTicketingConnectionIDNoteTicketIDID](#patchticketingconnectionidnoteticketidid) - Update a note
* [PostTicketingConnectionIDNoteTicketID](#postticketingconnectionidnoteticketid) - Create a note
* [PutTicketingConnectionIDNoteTicketIDID](#putticketingconnectionidnoteticketidid) - Update a note

## DeleteTicketingConnectionIDNoteTicketIDID

Remove a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.DeleteTicketingConnectionIDNoteTicketIDID(ctx, operations.DeleteTicketingConnectionIDNoteTicketIDIDRequest{
        ConnectionID: "DRAM Liaison",
        ID: "<ID>",
        TicketID: "Tasty exploit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.DeleteTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/deleteticketingconnectionidnoteticketididrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.DeleteTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/deleteticketingconnectionidnoteticketididresponse.md), error**


## GetTicketingConnectionIDNoteTicketID

List all notes

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.GetTicketingConnectionIDNoteTicketID(ctx, operations.GetTicketingConnectionIDNoteTicketIDRequest{
        ConnectionID: "Account revolutionary",
        TicketID: "Associate",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetTicketingConnectionIDNoteTicketIDRequest](../../models/operations/getticketingconnectionidnoteticketidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetTicketingConnectionIDNoteTicketIDResponse](../../models/operations/getticketingconnectionidnoteticketidresponse.md), error**


## GetTicketingConnectionIDNoteTicketIDID

Retrieve a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.GetTicketingConnectionIDNoteTicketIDID(ctx, operations.GetTicketingConnectionIDNoteTicketIDIDRequest{
        ConnectionID: "for",
        ID: "<ID>",
        TicketID: "female",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/getticketingconnectionidnoteticketididrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/getticketingconnectionidnoteticketididresponse.md), error**


## PatchTicketingConnectionIDNoteTicketIDID

Update a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.PatchTicketingConnectionIDNoteTicketIDID(ctx, operations.PatchTicketingConnectionIDNoteTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            Raw: shared.PropertyTicketingNoteRaw{},
        },
        ConnectionID: "compress Oganesson",
        ID: "<ID>",
        TicketID: "York Fantastic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PatchTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/patchticketingconnectionidnoteticketididrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PatchTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/patchticketingconnectionidnoteticketididresponse.md), error**


## PostTicketingConnectionIDNoteTicketID

Create a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.PostTicketingConnectionIDNoteTicketID(ctx, operations.PostTicketingConnectionIDNoteTicketIDRequest{
        TicketingNote: &shared.TicketingNote{
            Raw: shared.PropertyTicketingNoteRaw{},
        },
        ConnectionID: "Plantation blue",
        TicketID: "Ford",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PostTicketingConnectionIDNoteTicketIDRequest](../../models/operations/postticketingconnectionidnoteticketidrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.PostTicketingConnectionIDNoteTicketIDResponse](../../models/operations/postticketingconnectionidnoteticketidresponse.md), error**


## PutTicketingConnectionIDNoteTicketIDID

Update a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Note.PutTicketingConnectionIDNoteTicketIDID(ctx, operations.PutTicketingConnectionIDNoteTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            Raw: shared.PropertyTicketingNoteRaw{},
        },
        ConnectionID: "SMTP Cis",
        ID: "<ID>",
        TicketID: "East benchmark",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PutTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/putticketingconnectionidnoteticketididrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PutTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/putticketingconnectionidnoteticketididresponse.md), error**

