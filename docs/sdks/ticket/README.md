# Ticket

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteTicketingConnectionIDTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.DeleteTicketingConnectionIDTicketID(ctx, operations.DeleteTicketingConnectionIDTicketIDRequest{
        ConnectionID: "vel",
        ID: "e8dbf812-f83b-41ca-aa9f-fc561929cca9",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.DeleteTicketingConnectionIDTicketIDRequest](../../models/operations/deleteticketingconnectionidticketidrequest.md)   | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `security`                                                                                                                       | [operations.DeleteTicketingConnectionIDTicketIDSecurity](../../models/operations/deleteticketingconnectionidticketidsecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetTicketingConnectionIDTicketSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.GetTicketingConnectionIDTicket(ctx, operations.GetTicketingConnectionIDTicketRequest{
        AgentID: unifiedto.String("nemo"),
        ConnectionID: "laboriosam",
        CustomerID: unifiedto.String("eaque"),
        Limit: unifiedto.Float64(6814.58),
        Offset: unifiedto.Float64(977.35),
        Order: unifiedto.String("adipisci"),
        Query: unifiedto.String("occaecati"),
        Sort: unifiedto.String("exercitationem"),
        UpdatedGte: types.MustDateFromString("2022-11-09"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTickets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetTicketingConnectionIDTicketRequest](../../models/operations/getticketingconnectionidticketrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.GetTicketingConnectionIDTicketSecurity](../../models/operations/getticketingconnectionidticketsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetTicketingConnectionIDTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.GetTicketingConnectionIDTicketID(ctx, operations.GetTicketingConnectionIDTicketIDRequest{
        ConnectionID: "quas",
        ID: "da1d48e7-8e3c-4f8e-9143-da9308b27a08",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetTicketingConnectionIDTicketIDRequest](../../models/operations/getticketingconnectionidticketidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.GetTicketingConnectionIDTicketIDSecurity](../../models/operations/getticketingconnectionidticketidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchTicketingConnectionIDTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.PatchTicketingConnectionIDTicketID(ctx, operations.PatchTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("animi"),
            ClosedAt: types.MustDateFromString("2022-06-11"),
            CreatedAt: types.MustDateFromString("2022-11-29"),
            CustomerID: unifiedto.String("voluptatum"),
            Description: unifiedto.String("eius"),
            ID: unifiedto.String("439b3de8-756c-4cce-870c-d2147b6e6152"),
            Priority: unifiedto.String("placeat"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("voluptatibus"),
            SourceRef: unifiedto.String("ipsa"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedto.String("quibusdam"),
            Tags: []string{
                "doloremque",
            },
            UpdatedAt: types.MustDateFromString("2021-05-28"),
        },
        ConnectionID: "eligendi",
        ID: "3a4b9a5b-f935-4dfe-974f-a4b1e9c097ed",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PatchTicketingConnectionIDTicketIDRequest](../../models/operations/patchticketingconnectionidticketidrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.PatchTicketingConnectionIDTicketIDSecurity](../../models/operations/patchticketingconnectionidticketidsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostTicketingConnectionIDTicketSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.PostTicketingConnectionIDTicket(ctx, operations.PostTicketingConnectionIDTicketRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("animi"),
            ClosedAt: types.MustDateFromString("2022-11-09"),
            CreatedAt: types.MustDateFromString("2022-09-13"),
            CustomerID: unifiedto.String("numquam"),
            Description: unifiedto.String("fugit"),
            ID: unifiedto.String("e1a9237e-9984-4c80-b479-e891923c18ca"),
            Priority: unifiedto.String("rem"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("facere"),
            SourceRef: unifiedto.String("vel"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedto.String("porro"),
            Tags: []string{
                "enim",
            },
            UpdatedAt: types.MustDateFromString("2022-06-24"),
        },
        ConnectionID: "cupiditate",
    }, operationSecurity)
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
| `request`                                                                                                                | [operations.PostTicketingConnectionIDTicketRequest](../../models/operations/postticketingconnectionidticketrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PostTicketingConnectionIDTicketSecurity](../../models/operations/postticketingconnectionidticketsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutTicketingConnectionIDTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticket.PutTicketingConnectionIDTicketID(ctx, operations.PutTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("explicabo"),
            ClosedAt: types.MustDateFromString("2022-09-23"),
            CreatedAt: types.MustDateFromString("2021-01-13"),
            CustomerID: unifiedto.String("consequuntur"),
            Description: unifiedto.String("doloremque"),
            ID: unifiedto.String("207e4fae-038c-4d7f-9bc2-cabaf7fc2ccb"),
            Priority: unifiedto.String("id"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("numquam"),
            SourceRef: unifiedto.String("libero"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedto.String("asperiores"),
            Tags: []string{
                "aperiam",
            },
            UpdatedAt: types.MustDateFromString("2020-02-02"),
        },
        ConnectionID: "nisi",
        ID: "8eaedb2e-e70b-4e06-9fb3-6add704080e0",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingTicket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PutTicketingConnectionIDTicketIDRequest](../../models/operations/putticketingconnectionidticketidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.PutTicketingConnectionIDTicketIDSecurity](../../models/operations/putticketingconnectionidticketidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.PutTicketingConnectionIDTicketIDResponse](../../models/operations/putticketingconnectionidticketidresponse.md), error**

