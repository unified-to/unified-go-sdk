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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
        Limit: unifiedgosdk.Float64(2310.88),
        Offset: unifiedgosdk.Float64(6688.82),
        Order: unifiedgosdk.String("AI"),
        Query: unifiedgosdk.String("stanch Investor attitude"),
        Sort: unifiedgosdk.String("Cotton"),
        TicketID: "Handmade Kia",
        UpdatedGte: types.MustTimeFromString("2022-05-26T17:12:11.333Z"),
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            AgentID: unifiedgosdk.String("compress Oganesson"),
            CreatedAt: types.MustTimeFromString("2022-02-16T08:13:19.991Z"),
            CustomerID: unifiedgosdk.String("demystify"),
            Description: unifiedgosdk.String("Fundamental demand-driven workforce"),
            ID: unifiedgosdk.String("<ID>"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("Nissan"),
        },
        ConnectionID: "Chicken",
        ID: "<ID>",
        TicketID: "frictionless convergence officia",
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            AgentID: unifiedgosdk.String("Plantation blue"),
            CreatedAt: types.MustTimeFromString("2021-06-11T06:54:31.529Z"),
            CustomerID: unifiedgosdk.String("asymmetric"),
            Description: unifiedgosdk.String("Expanded intermediate attitude"),
            ID: unifiedgosdk.String("<ID>"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("naturally"),
        },
        ConnectionID: "Wagon Sulfur",
        TicketID: "digital",
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            AgentID: unifiedgosdk.String("SMTP Cis"),
            CreatedAt: types.MustTimeFromString("2022-07-27T18:14:06.584Z"),
            CustomerID: unifiedgosdk.String("Carolina"),
            Description: unifiedgosdk.String("Integrated asymmetric strategy"),
            ID: unifiedgosdk.String("<ID>"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("Northeast Morocco supposing"),
        },
        ConnectionID: "DNS Fermium",
        ID: "<ID>",
        TicketID: "Southwest round",
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

