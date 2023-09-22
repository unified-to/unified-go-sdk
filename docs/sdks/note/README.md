# Note

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
        ConnectionID: "recusandae",
        ID: "9085075b-c253-4825-b343-fb0a4e66ea47",
        TicketID: "exercitationem",
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
        ConnectionID: "molestiae",
        Limit: unifiedgosdk.Float64(5096.38),
        Offset: unifiedgosdk.Float64(8540.88),
        Order: unifiedgosdk.String("ab"),
        Query: unifiedgosdk.String("voluptate"),
        Sort: unifiedgosdk.String("et"),
        TicketID: "recusandae",
        UpdatedGte: types.MustTimeFromString("2022-05-19T12:21:22.703Z"),
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
        ConnectionID: "eius",
        ID: "1818fc67-9b6b-42f2-9359-b855d015b62c",
        TicketID: "quos",
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
            AgentID: unifiedgosdk.String("cum"),
            CreatedAt: types.MustTimeFromString("2022-07-23T08:55:47.914Z"),
            CustomerID: unifiedgosdk.String("dolorum"),
            Description: unifiedgosdk.String("amet"),
            ID: unifiedgosdk.String("8a8a88c1-4420-40c2-8aeb-1ae1ecf8c349"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("modi"),
        },
        ConnectionID: "commodi",
        ID: "bba7a05a-8b4a-49ec-9b36-88cca3632727",
        TicketID: "iure",
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
            AgentID: unifiedgosdk.String("sit"),
            CreatedAt: types.MustTimeFromString("2021-04-03T18:27:00.342Z"),
            CustomerID: unifiedgosdk.String("vel"),
            Description: unifiedgosdk.String("autem"),
            ID: unifiedgosdk.String("e97e0541-0334-47d7-8ff2-491145fab9e5"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("perspiciatis"),
        },
        ConnectionID: "laborum",
        TicketID: "incidunt",
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
            AgentID: unifiedgosdk.String("fuga"),
            CreatedAt: types.MustTimeFromString("2022-05-09T22:54:29.590Z"),
            CustomerID: unifiedgosdk.String("amet"),
            Description: unifiedgosdk.String("nisi"),
            ID: unifiedgosdk.String("664eaa6b-f2ff-414e-8c1b-352accedacc5"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedgosdk.String("consequuntur"),
        },
        ConnectionID: "qui",
        ID: "7814eca0-16bc-441e-a134-2d4104a25ef7",
        TicketID: "quasi",
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

