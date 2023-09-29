# Ticket
(*Ticket*)

### Available Operations

* [DeleteTicketingConnectionIDTicketID](#deleteticketingconnectionidticketid) - Remove a ticket
* [GetTicketingConnectionIDTicket](#getticketingconnectionidticket) - List all tickets
* [GetTicketingConnectionIDTicketID](#getticketingconnectionidticketid) - Retrieve a ticket
* [PatchTicketingConnectionIDTicketID](#patchticketingconnectionidticketid) - Update a ticket
* [PostTicketingConnectionIDTicket](#postticketingconnectionidticket) - Create a ticket
* [PutTicketingConnectionIDTicketID](#putticketingconnectionidticketid) - Update a ticket

## DeleteTicketingConnectionIDTicketID

Remove a ticket

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
    res, err := s.Ticket.DeleteTicketingConnectionIDTicketID(ctx, operations.DeleteTicketingConnectionIDTicketIDRequest{
        ConnectionID: "brownie azure payment",
        ID: "<ID>",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteTicketingConnectionIDTicketIDRequest](../../models/operations/deleteticketingconnectionidticketidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.DeleteTicketingConnectionIDTicketIDResponse](../../models/operations/deleteticketingconnectionidticketidresponse.md), error**


## GetTicketingConnectionIDTicket

List all tickets

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
    res, err := s.Ticket.GetTicketingConnectionIDTicket(ctx, operations.GetTicketingConnectionIDTicketRequest{
        AgentID: unifiedgosdk.String("New"),
        ConnectionID: "hertz Savings Steel",
        CustomerID: unifiedgosdk.String("payment biopsy Kids"),
        Limit: unifiedgosdk.Float64(7673.64),
        Offset: unifiedgosdk.Float64(5134.74),
        Order: unifiedgosdk.String("quantifying orange"),
        Query: unifiedgosdk.String("male dynamic"),
        Sort: unifiedgosdk.String("Sedan Tricycle Honda"),
        UpdatedGte: types.MustTimeFromString("2022-08-06T21:30:52.879Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTickets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetTicketingConnectionIDTicketRequest](../../models/operations/getticketingconnectionidticketrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetTicketingConnectionIDTicketResponse](../../models/operations/getticketingconnectionidticketresponse.md), error**


## GetTicketingConnectionIDTicketID

Retrieve a ticket

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
    res, err := s.Ticket.GetTicketingConnectionIDTicketID(ctx, operations.GetTicketingConnectionIDTicketIDRequest{
        ConnectionID: "yellow",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetTicketingConnectionIDTicketIDRequest](../../models/operations/getticketingconnectionidticketidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GetTicketingConnectionIDTicketIDResponse](../../models/operations/getticketingconnectionidticketidresponse.md), error**


## PatchTicketingConnectionIDTicketID

Update a ticket

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
    res, err := s.Ticket.PatchTicketingConnectionIDTicketID(ctx, operations.PatchTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedgosdk.String("Representative calculate"),
            ClosedAt: types.MustTimeFromString("2023-12-03T14:58:54.732Z"),
            CreatedAt: types.MustTimeFromString("2022-09-11T04:52:37.095Z"),
            CustomerID: unifiedgosdk.String("indigo extend given"),
            Description: unifiedgosdk.String("Profound motivating utilisation"),
            ID: unifiedgosdk.String("<ID>"),
            Priority: unifiedgosdk.String("Hill Jazz"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedgosdk.String("West Macedonia City"),
            SourceRef: unifiedgosdk.String("orange West doubtfully"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedgosdk.String("Pizza"),
            Tags: []string{
                "definition",
            },
            UpdatedAt: types.MustTimeFromString("2021-10-05T23:17:22.031Z"),
        },
        ConnectionID: "engage henry",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PatchTicketingConnectionIDTicketIDRequest](../../models/operations/patchticketingconnectionidticketidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PatchTicketingConnectionIDTicketIDResponse](../../models/operations/patchticketingconnectionidticketidresponse.md), error**


## PostTicketingConnectionIDTicket

Create a ticket

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
    res, err := s.Ticket.PostTicketingConnectionIDTicket(ctx, operations.PostTicketingConnectionIDTicketRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedgosdk.String("North"),
            ClosedAt: types.MustTimeFromString("2021-08-03T02:12:35.164Z"),
            CreatedAt: types.MustTimeFromString("2023-05-12T14:26:26.768Z"),
            CustomerID: unifiedgosdk.String("mull hierarchy"),
            Description: unifiedgosdk.String("Triple-buffered solution-oriented info-mediaries"),
            ID: unifiedgosdk.String("<ID>"),
            Priority: unifiedgosdk.String("person Idaho"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedgosdk.String("Convertible whenever feed"),
            SourceRef: unifiedgosdk.String("solid Electric Bespoke"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedgosdk.String("sint uplift"),
            Tags: []string{
                "Idaho",
            },
            UpdatedAt: types.MustTimeFromString("2022-06-24T01:04:15.890Z"),
        },
        ConnectionID: "Oriental outrage",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PostTicketingConnectionIDTicketRequest](../../models/operations/postticketingconnectionidticketrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PostTicketingConnectionIDTicketResponse](../../models/operations/postticketingconnectionidticketresponse.md), error**


## PutTicketingConnectionIDTicketID

Update a ticket

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
    res, err := s.Ticket.PutTicketingConnectionIDTicketID(ctx, operations.PutTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedgosdk.String("North Finland"),
            ClosedAt: types.MustTimeFromString("2023-12-08T00:37:44.739Z"),
            CreatedAt: types.MustTimeFromString("2023-04-01T07:24:49.830Z"),
            CustomerID: unifiedgosdk.String("Marketing"),
            Description: unifiedgosdk.String("Future-proofed high-level system engine"),
            ID: unifiedgosdk.String("<ID>"),
            Priority: unifiedgosdk.String("drat knottily"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedgosdk.String("Upgradable knuckle"),
            SourceRef: unifiedgosdk.String("anenst"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.String("indexing Wooden Crew"),
            Tags: []string{
                "anti",
            },
            UpdatedAt: types.MustTimeFromString("2023-08-10T07:27:15.153Z"),
        },
        ConnectionID: "neural orchestrate",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutTicketingConnectionIDTicketIDRequest](../../models/operations/putticketingconnectionidticketidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.PutTicketingConnectionIDTicketIDResponse](../../models/operations/putticketingconnectionidticketidresponse.md), error**

