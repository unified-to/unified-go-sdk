# Note

### Available Operations

* [DeleteTicketingConnectionIDNotesTicketIDID](#deleteticketingconnectionidnotesticketidid) - Remove a note
* [GetTicketingConnectionIDNotesTicketID](#getticketingconnectionidnotesticketid) - List all notes
* [GetTicketingConnectionIDNotesTicketIDID](#getticketingconnectionidnotesticketidid) - Retrieve a note
* [PatchTicketingConnectionIDNotesTicketIDID](#patchticketingconnectionidnotesticketidid) - Update a note
* [PostTicketingConnectionIDNotesTicketID](#postticketingconnectionidnotesticketid) - Create a note
* [PutTicketingConnectionIDNotesTicketIDID](#putticketingconnectionidnotesticketidid) - Update a note

## DeleteTicketingConnectionIDNotesTicketIDID

Remove a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteTicketingConnectionIDNotesTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.DeleteTicketingConnectionIDNotesTicketIDID(ctx, operations.DeleteTicketingConnectionIDNotesTicketIDIDRequest{
        ConnectionID: "recusandae",
        ID: "9085075b-c253-4825-b343-fb0a4e66ea47",
        TicketID: "exercitationem",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.DeleteTicketingConnectionIDNotesTicketIDIDRequest](../../models/operations/deleteticketingconnectionidnotesticketididrequest.md)   | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `security`                                                                                                                                     | [operations.DeleteTicketingConnectionIDNotesTicketIDIDSecurity](../../models/operations/deleteticketingconnectionidnotesticketididsecurity.md) | :heavy_check_mark:                                                                                                                             | The security requirements to use for the request.                                                                                              |


### Response

**[*operations.DeleteTicketingConnectionIDNotesTicketIDIDResponse](../../models/operations/deleteticketingconnectionidnotesticketididresponse.md), error**


## GetTicketingConnectionIDNotesTicketID

List all notes

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetTicketingConnectionIDNotesTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.GetTicketingConnectionIDNotesTicketID(ctx, operations.GetTicketingConnectionIDNotesTicketIDRequest{
        ConnectionID: "molestiae",
        Limit: unifiedto.Float64(5096.38),
        Offset: unifiedto.Float64(8540.88),
        Order: unifiedto.String("ab"),
        Query: unifiedto.String("voluptate"),
        Sort: unifiedto.String("et"),
        TicketID: "recusandae",
        UpdatedGte: types.MustDateFromString("2022-05-19"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetTicketingConnectionIDNotesTicketIDRequest](../../models/operations/getticketingconnectionidnotesticketidrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.GetTicketingConnectionIDNotesTicketIDSecurity](../../models/operations/getticketingconnectionidnotesticketidsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.GetTicketingConnectionIDNotesTicketIDResponse](../../models/operations/getticketingconnectionidnotesticketidresponse.md), error**


## GetTicketingConnectionIDNotesTicketIDID

Retrieve a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetTicketingConnectionIDNotesTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.GetTicketingConnectionIDNotesTicketIDID(ctx, operations.GetTicketingConnectionIDNotesTicketIDIDRequest{
        ConnectionID: "eius",
        ID: "1818fc67-9b6b-42f2-9359-b855d015b62c",
        TicketID: "quos",
    }, operationSecurity)
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
| `request`                                                                                                                                | [operations.GetTicketingConnectionIDNotesTicketIDIDRequest](../../models/operations/getticketingconnectionidnotesticketididrequest.md)   | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.GetTicketingConnectionIDNotesTicketIDIDSecurity](../../models/operations/getticketingconnectionidnotesticketididsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |


### Response

**[*operations.GetTicketingConnectionIDNotesTicketIDIDResponse](../../models/operations/getticketingconnectionidnotesticketididresponse.md), error**


## PatchTicketingConnectionIDNotesTicketIDID

Update a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchTicketingConnectionIDNotesTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.PatchTicketingConnectionIDNotesTicketIDID(ctx, operations.PatchTicketingConnectionIDNotesTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("cum"),
            CreatedAt: types.MustDateFromString("2022-07-23"),
            CustomerID: unifiedto.String("dolorum"),
            Description: unifiedto.String("amet"),
            ID: unifiedto.String("8a8a88c1-4420-40c2-8aeb-1ae1ecf8c349"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("modi"),
        },
        ConnectionID: "commodi",
        ID: "bba7a05a-8b4a-49ec-9b36-88cca3632727",
        TicketID: "iure",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PatchTicketingConnectionIDNotesTicketIDIDRequest](../../models/operations/patchticketingconnectionidnotesticketididrequest.md)   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.PatchTicketingConnectionIDNotesTicketIDIDSecurity](../../models/operations/patchticketingconnectionidnotesticketididsecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |


### Response

**[*operations.PatchTicketingConnectionIDNotesTicketIDIDResponse](../../models/operations/patchticketingconnectionidnotesticketididresponse.md), error**


## PostTicketingConnectionIDNotesTicketID

Create a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostTicketingConnectionIDNotesTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.PostTicketingConnectionIDNotesTicketID(ctx, operations.PostTicketingConnectionIDNotesTicketIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("sit"),
            CreatedAt: types.MustDateFromString("2021-04-03"),
            CustomerID: unifiedto.String("vel"),
            Description: unifiedto.String("autem"),
            ID: unifiedto.String("e97e0541-0334-47d7-8ff2-491145fab9e5"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("perspiciatis"),
        },
        ConnectionID: "laborum",
        TicketID: "incidunt",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PostTicketingConnectionIDNotesTicketIDRequest](../../models/operations/postticketingconnectionidnotesticketidrequest.md)   | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.PostTicketingConnectionIDNotesTicketIDSecurity](../../models/operations/postticketingconnectionidnotesticketidsecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


### Response

**[*operations.PostTicketingConnectionIDNotesTicketIDResponse](../../models/operations/postticketingconnectionidnotesticketidresponse.md), error**


## PutTicketingConnectionIDNotesTicketIDID

Update a note

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutTicketingConnectionIDNotesTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Note.PutTicketingConnectionIDNotesTicketIDID(ctx, operations.PutTicketingConnectionIDNotesTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("fuga"),
            CreatedAt: types.MustDateFromString("2022-05-09"),
            CustomerID: unifiedto.String("amet"),
            Description: unifiedto.String("nisi"),
            ID: unifiedto.String("664eaa6b-f2ff-414e-8c1b-352accedacc5"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("consequuntur"),
        },
        ConnectionID: "qui",
        ID: "7814eca0-16bc-441e-a134-2d4104a25ef7",
        TicketID: "quasi",
    }, operationSecurity)
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
| `request`                                                                                                                                | [operations.PutTicketingConnectionIDNotesTicketIDIDRequest](../../models/operations/putticketingconnectionidnotesticketididrequest.md)   | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.PutTicketingConnectionIDNotesTicketIDIDSecurity](../../models/operations/putticketingconnectionidnotesticketididsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |


### Response

**[*operations.PutTicketingConnectionIDNotesTicketIDIDResponse](../../models/operations/putticketingconnectionidnotesticketididresponse.md), error**

