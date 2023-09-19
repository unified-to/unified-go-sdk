# Ticketing

### Available Operations

* [DeleteTicketingConnectionIDAgentID](#deleteticketingconnectionidagentid) - Remove a agent
* [DeleteTicketingConnectionIDCustomerID](#deleteticketingconnectionidcustomerid) - Remove a customer
* [DeleteTicketingConnectionIDNoteTicketIDID](#deleteticketingconnectionidnoteticketidid) - Remove a note
* [DeleteTicketingConnectionIDTicketID](#deleteticketingconnectionidticketid) - Remove a ticket
* [GetTicketingConnectionIDAgent](#getticketingconnectionidagent) - List all agents
* [GetTicketingConnectionIDAgentID](#getticketingconnectionidagentid) - Retrieve a agent
* [GetTicketingConnectionIDCustomer](#getticketingconnectionidcustomer) - List all customers
* [GetTicketingConnectionIDCustomerID](#getticketingconnectionidcustomerid) - Retrieve a customer
* [GetTicketingConnectionIDNoteTicketID](#getticketingconnectionidnoteticketid) - List all notes
* [GetTicketingConnectionIDNoteTicketIDID](#getticketingconnectionidnoteticketidid) - Retrieve a note
* [GetTicketingConnectionIDTicket](#getticketingconnectionidticket) - List all tickets
* [GetTicketingConnectionIDTicketID](#getticketingconnectionidticketid) - Retrieve a ticket
* [PatchTicketingConnectionIDAgentID](#patchticketingconnectionidagentid) - Update a agent
* [PatchTicketingConnectionIDCustomerID](#patchticketingconnectionidcustomerid) - Update a customer
* [PatchTicketingConnectionIDNoteTicketIDID](#patchticketingconnectionidnoteticketidid) - Update a note
* [PatchTicketingConnectionIDTicketID](#patchticketingconnectionidticketid) - Update a ticket
* [PostTicketingConnectionIDAgent](#postticketingconnectionidagent) - Create a agent
* [PostTicketingConnectionIDCustomer](#postticketingconnectionidcustomer) - Create a customer
* [PostTicketingConnectionIDNoteTicketID](#postticketingconnectionidnoteticketid) - Create a note
* [PostTicketingConnectionIDTicket](#postticketingconnectionidticket) - Create a ticket
* [PutTicketingConnectionIDAgentID](#putticketingconnectionidagentid) - Update a agent
* [PutTicketingConnectionIDCustomerID](#putticketingconnectionidcustomerid) - Update a customer
* [PutTicketingConnectionIDNoteTicketIDID](#putticketingconnectionidnoteticketidid) - Update a note
* [PutTicketingConnectionIDTicketID](#putticketingconnectionidticketid) - Update a ticket

## DeleteTicketingConnectionIDAgentID

Remove a agent

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
    operationSecurity := operations.DeleteTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.DeleteTicketingConnectionIDAgentID(ctx, operations.DeleteTicketingConnectionIDAgentIDRequest{
        ConnectionID: "fuga",
        ID: "3fc73a5a-034b-4114-9924-3afa6987a472",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteTicketingConnectionIDAgentIDRequest](../../models/operations/deleteticketingconnectionidagentidrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.DeleteTicketingConnectionIDAgentIDSecurity](../../models/operations/deleteticketingconnectionidagentidsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.DeleteTicketingConnectionIDAgentIDResponse](../../models/operations/deleteticketingconnectionidagentidresponse.md), error**


## DeleteTicketingConnectionIDCustomerID

Remove a customer

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
    operationSecurity := operations.DeleteTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.DeleteTicketingConnectionIDCustomerID(ctx, operations.DeleteTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "libero",
        ID: "709a153e-2230-4106-8539-ce0932d10acd",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.DeleteTicketingConnectionIDCustomerIDRequest](../../models/operations/deleteticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.DeleteTicketingConnectionIDCustomerIDSecurity](../../models/operations/deleteticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.DeleteTicketingConnectionIDCustomerIDResponse](../../models/operations/deleteticketingconnectionidcustomeridresponse.md), error**


## DeleteTicketingConnectionIDNoteTicketIDID

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
    operationSecurity := operations.DeleteTicketingConnectionIDNoteTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.DeleteTicketingConnectionIDNoteTicketIDID(ctx, operations.DeleteTicketingConnectionIDNoteTicketIDIDRequest{
        ConnectionID: "vitae",
        ID: "5d8cc306-b786-4b3d-b7bd-204a1f340bb3",
        TicketID: "ex",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.DeleteTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/deleteticketingconnectionidnoteticketididrequest.md)   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.DeleteTicketingConnectionIDNoteTicketIDIDSecurity](../../models/operations/deleteticketingconnectionidnoteticketididsecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |


### Response

**[*operations.DeleteTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/deleteticketingconnectionidnoteticketididresponse.md), error**


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
    res, err := s.Ticketing.DeleteTicketingConnectionIDTicketID(ctx, operations.DeleteTicketingConnectionIDTicketIDRequest{
        ConnectionID: "voluptatibus",
        ID: "677a4851-9c33-4749-8284-8826bb03c7fd",
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


## GetTicketingConnectionIDAgent

List all agents

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
    operationSecurity := operations.GetTicketingConnectionIDAgentSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDAgent(ctx, operations.GetTicketingConnectionIDAgentRequest{
        ConnectionID: "consequuntur",
        Limit: unifiedto.Float64(1427.69),
        Offset: unifiedto.Float64(3147.32),
        Order: unifiedto.String("debitis"),
        Query: unifiedto.String("dolore"),
        Sort: unifiedto.String("in"),
        UpdatedGte: types.MustTimeFromString("2022-01-15T17:11:52.134Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetTicketingConnectionIDAgentRequest](../../models/operations/getticketingconnectionidagentrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetTicketingConnectionIDAgentSecurity](../../models/operations/getticketingconnectionidagentsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.GetTicketingConnectionIDAgentResponse](../../models/operations/getticketingconnectionidagentresponse.md), error**


## GetTicketingConnectionIDAgentID

Retrieve a agent

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
    operationSecurity := operations.GetTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDAgentID(ctx, operations.GetTicketingConnectionIDAgentIDRequest{
        ConnectionID: "architecto",
        ID: "a88ed72a-2d4a-4f41-98ac-2d0f0f58c3b8",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetTicketingConnectionIDAgentIDRequest](../../models/operations/getticketingconnectionidagentidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.GetTicketingConnectionIDAgentIDSecurity](../../models/operations/getticketingconnectionidagentidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.GetTicketingConnectionIDAgentIDResponse](../../models/operations/getticketingconnectionidagentidresponse.md), error**


## GetTicketingConnectionIDCustomer

List all customers

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
    operationSecurity := operations.GetTicketingConnectionIDCustomerSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDCustomer(ctx, operations.GetTicketingConnectionIDCustomerRequest{
        ConnectionID: "in",
        Limit: unifiedto.Float64(7360.32),
        Offset: unifiedto.Float64(2850.04),
        Order: unifiedto.String("molestiae"),
        Query: unifiedto.String("eaque"),
        Sort: unifiedto.String("tempora"),
        UpdatedGte: types.MustTimeFromString("2022-03-06T22:27:54.190Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetTicketingConnectionIDCustomerRequest](../../models/operations/getticketingconnectionidcustomerrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.GetTicketingConnectionIDCustomerSecurity](../../models/operations/getticketingconnectionidcustomersecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.GetTicketingConnectionIDCustomerResponse](../../models/operations/getticketingconnectionidcustomerresponse.md), error**


## GetTicketingConnectionIDCustomerID

Retrieve a customer

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
    operationSecurity := operations.GetTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDCustomerID(ctx, operations.GetTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "aut",
        ID: "d98e9d82-c5e3-406f-9576-f5cdeb0286d0",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetTicketingConnectionIDCustomerIDRequest](../../models/operations/getticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.GetTicketingConnectionIDCustomerIDSecurity](../../models/operations/getticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.GetTicketingConnectionIDCustomerIDResponse](../../models/operations/getticketingconnectionidcustomeridresponse.md), error**


## GetTicketingConnectionIDNoteTicketID

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
    operationSecurity := operations.GetTicketingConnectionIDNoteTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDNoteTicketID(ctx, operations.GetTicketingConnectionIDNoteTicketIDRequest{
        ConnectionID: "distinctio",
        Limit: unifiedto.Float64(8031.14),
        Offset: unifiedto.Float64(3028.92),
        Order: unifiedto.String("adipisci"),
        Query: unifiedto.String("harum"),
        Sort: unifiedto.String("veritatis"),
        TicketID: "quas",
        UpdatedGte: types.MustTimeFromString("2021-07-31T21:20:45.941Z"),
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetTicketingConnectionIDNoteTicketIDRequest](../../models/operations/getticketingconnectionidnoteticketidrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.GetTicketingConnectionIDNoteTicketIDSecurity](../../models/operations/getticketingconnectionidnoteticketidsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetTicketingConnectionIDNoteTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.GetTicketingConnectionIDNoteTicketIDID(ctx, operations.GetTicketingConnectionIDNoteTicketIDIDRequest{
        ConnectionID: "ipsum",
        ID: "78f2fcff-81dd-4f7e-888f-74ef54c9216e",
        TicketID: "atque",
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
| `request`                                                                                                                              | [operations.GetTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/getticketingconnectionidnoteticketididrequest.md)   | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.GetTicketingConnectionIDNoteTicketIDIDSecurity](../../models/operations/getticketingconnectionidnoteticketididsecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


### Response

**[*operations.GetTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/getticketingconnectionidnoteticketididresponse.md), error**


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
    res, err := s.Ticketing.GetTicketingConnectionIDTicket(ctx, operations.GetTicketingConnectionIDTicketRequest{
        AgentID: unifiedto.String("unde"),
        ConnectionID: "qui",
        CustomerID: unifiedto.String("aliquid"),
        Limit: unifiedto.Float64(1977.7),
        Offset: unifiedto.Float64(641.35),
        Order: unifiedto.String("velit"),
        Query: unifiedto.String("libero"),
        Sort: unifiedto.String("soluta"),
        UpdatedGte: types.MustTimeFromString("2022-01-21T15:01:10.881Z"),
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
    res, err := s.Ticketing.GetTicketingConnectionIDTicketID(ctx, operations.GetTicketingConnectionIDTicketIDRequest{
        ConnectionID: "quo",
        ID: "2c8d2701-096b-466a-96e3-e1d9d3b66033",
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


## PatchTicketingConnectionIDAgentID

Update a agent

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
    operationSecurity := operations.PatchTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PatchTicketingConnectionIDAgentID(ctx, operations.PatchTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-05-08T09:12:54.892Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Brad.Olson@gmail.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("5d2247de-9b3d-4461-b0e7-68a96bb39878"),
            Name: unifiedto.String("Jimmy Metz"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "tempore",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-15T11:30:38.762Z"),
        },
        ConnectionID: "cum",
        ID: "f7143356-f634-49a1-a424-9b211ce46b95",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PatchTicketingConnectionIDAgentIDRequest](../../models/operations/patchticketingconnectionidagentidrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.PatchTicketingConnectionIDAgentIDSecurity](../../models/operations/patchticketingconnectionidagentidsecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.PatchTicketingConnectionIDAgentIDResponse](../../models/operations/patchticketingconnectionidagentidresponse.md), error**


## PatchTicketingConnectionIDCustomerID

Update a customer

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
    operationSecurity := operations.PatchTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PatchTicketingConnectionIDCustomerID(ctx, operations.PatchTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustTimeFromString("2022-07-29T08:43:35.611Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Brandi35@gmail.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("ca9142f0-5263-42b3-9cad-692ffc874500"),
            Name: unifiedto.String("Lorena Moore"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "temporibus",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "excepturi",
                    Type: shared.TicketingTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-02-14T22:17:14.340Z"),
        },
        ConnectionID: "aperiam",
        ID: "36f5c388-664f-4698-9530-a2e2aed6aaf8",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PatchTicketingConnectionIDCustomerIDRequest](../../models/operations/patchticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PatchTicketingConnectionIDCustomerIDSecurity](../../models/operations/patchticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


### Response

**[*operations.PatchTicketingConnectionIDCustomerIDResponse](../../models/operations/patchticketingconnectionidcustomeridresponse.md), error**


## PatchTicketingConnectionIDNoteTicketIDID

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
    operationSecurity := operations.PatchTicketingConnectionIDNoteTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PatchTicketingConnectionIDNoteTicketIDID(ctx, operations.PatchTicketingConnectionIDNoteTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("ea"),
            CreatedAt: types.MustTimeFromString("2022-03-24T22:55:36.292Z"),
            CustomerID: unifiedto.String("eos"),
            Description: unifiedto.String("praesentium"),
            ID: unifiedto.String("d040c69a-3d90-46f6-abd5-ad7ec7394f25"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("repellat"),
        },
        ConnectionID: "ex",
        ID: "34b37307-14e6-4be8-83e0-9c64d342ac29",
        TicketID: "provident",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PatchTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/patchticketingconnectionidnoteticketididrequest.md)   | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `security`                                                                                                                                 | [operations.PatchTicketingConnectionIDNoteTicketIDIDSecurity](../../models/operations/patchticketingconnectionidnoteticketididsecurity.md) | :heavy_check_mark:                                                                                                                         | The security requirements to use for the request.                                                                                          |


### Response

**[*operations.PatchTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/patchticketingconnectionidnoteticketididresponse.md), error**


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
    res, err := s.Ticketing.PatchTicketingConnectionIDTicketID(ctx, operations.PatchTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("officia"),
            ClosedAt: types.MustTimeFromString("2022-01-28T05:10:10.473Z"),
            CreatedAt: types.MustTimeFromString("2022-02-09T11:50:25.266Z"),
            CustomerID: unifiedto.String("quam"),
            Description: unifiedto.String("dolorum"),
            ID: unifiedto.String("ef13402e-945f-4537-83ef-de1198221f9b"),
            Priority: unifiedto.String("inventore"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("tenetur"),
            SourceRef: unifiedto.String("nihil"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedto.String("iste"),
            Tags: []string{
                "deserunt",
            },
            UpdatedAt: types.MustTimeFromString("2020-02-21T15:04:19.759Z"),
        },
        ConnectionID: "repudiandae",
        ID: "69682ace-efb0-44f8-8512-caabea708ed5",
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


## PostTicketingConnectionIDAgent

Create a agent

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
    operationSecurity := operations.PostTicketingConnectionIDAgentSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PostTicketingConnectionIDAgent(ctx, operations.PostTicketingConnectionIDAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-05-21T16:53:45.567Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Raven.Frami@yahoo.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("460599d5-c334-4957-ad55-209e9a2253b6"),
            Name: unifiedto.String("Clinton Huel"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "laudantium",
                    Type: shared.TicketingTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2020-03-13T17:02:39.583Z"),
        },
        ConnectionID: "similique",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PostTicketingConnectionIDAgentRequest](../../models/operations/postticketingconnectionidagentrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PostTicketingConnectionIDAgentSecurity](../../models/operations/postticketingconnectionidagentsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.PostTicketingConnectionIDAgentResponse](../../models/operations/postticketingconnectionidagentresponse.md), error**


## PostTicketingConnectionIDCustomer

Create a customer

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
    operationSecurity := operations.PostTicketingConnectionIDCustomerSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PostTicketingConnectionIDCustomer(ctx, operations.PostTicketingConnectionIDCustomerRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustTimeFromString("2021-11-22T08:46:26.211Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Sabina.Gulgowski59@hotmail.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("8a149067-8f13-4c68-ad83-9fc9e175ffa9"),
            Name: unifiedto.String("Ella Murazik"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "corporis",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "error",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-11-08T14:43:03.666Z"),
        },
        ConnectionID: "debitis",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PostTicketingConnectionIDCustomerRequest](../../models/operations/postticketingconnectionidcustomerrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.PostTicketingConnectionIDCustomerSecurity](../../models/operations/postticketingconnectionidcustomersecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.PostTicketingConnectionIDCustomerResponse](../../models/operations/postticketingconnectionidcustomerresponse.md), error**


## PostTicketingConnectionIDNoteTicketID

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
    operationSecurity := operations.PostTicketingConnectionIDNoteTicketIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PostTicketingConnectionIDNoteTicketID(ctx, operations.PostTicketingConnectionIDNoteTicketIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("quidem"),
            CreatedAt: types.MustTimeFromString("2022-07-05T14:59:15.588Z"),
            CustomerID: unifiedto.String("magnam"),
            Description: unifiedto.String("vel"),
            ID: unifiedto.String("030fe183-76c2-4bed-ae76-790ed0c16a7b"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("id"),
        },
        ConnectionID: "dolore",
        TicketID: "quam",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PostTicketingConnectionIDNoteTicketIDRequest](../../models/operations/postticketingconnectionidnoteticketidrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.PostTicketingConnectionIDNoteTicketIDSecurity](../../models/operations/postticketingconnectionidnoteticketidsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.PostTicketingConnectionIDNoteTicketIDResponse](../../models/operations/postticketingconnectionidnoteticketidresponse.md), error**


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
    res, err := s.Ticketing.PostTicketingConnectionIDTicket(ctx, operations.PostTicketingConnectionIDTicketRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("rem"),
            ClosedAt: types.MustTimeFromString("2022-12-31T16:32:53.167Z"),
            CreatedAt: types.MustTimeFromString("2022-09-20T04:19:27.059Z"),
            CustomerID: unifiedto.String("totam"),
            Description: unifiedto.String("unde"),
            ID: unifiedto.String("f6770ef0-4809-41a2-ba25-ee6c75af8a60"),
            Priority: unifiedto.String("laborum"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("quam"),
            SourceRef: unifiedto.String("laborum"),
            Status: shared.TicketingTicketStatusClosed.ToPointer(),
            Subject: unifiedto.String("dolor"),
            Tags: []string{
                "dolore",
            },
            UpdatedAt: types.MustTimeFromString("2022-01-27T16:39:01.281Z"),
        },
        ConnectionID: "perferendis",
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


## PutTicketingConnectionIDAgentID

Update a agent

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
    operationSecurity := operations.PutTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PutTicketingConnectionIDAgentID(ctx, operations.PutTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-02-01T01:30:25.348Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Sigrid_Hilll@yahoo.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("60acaca6-45de-4986-b551-a9cce61ec2c7"),
            Name: unifiedto.String("Gerard Feest"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "officiis",
                    Type: shared.TicketingTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-05-28T17:33:55.487Z"),
        },
        ConnectionID: "assumenda",
        ID: "5a65b4d5-562d-49b7-99e2-d6fcf557629d",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutTicketingConnectionIDAgentIDRequest](../../models/operations/putticketingconnectionidagentidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PutTicketingConnectionIDAgentIDSecurity](../../models/operations/putticketingconnectionidagentidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.PutTicketingConnectionIDAgentIDResponse](../../models/operations/putticketingconnectionidagentidresponse.md), error**


## PutTicketingConnectionIDCustomerID

Update a customer

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
    operationSecurity := operations.PutTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PutTicketingConnectionIDCustomerID(ctx, operations.PutTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustTimeFromString("2021-11-27T14:28:08.554Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Enola63@yahoo.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("90282a51-f41c-4f67-96ed-3d724c18f581"),
            Name: unifiedto.String("Terrence Langworth"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "officiis",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "consectetur",
                    Type: shared.TicketingTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-11-17T00:52:06.013Z"),
        },
        ConnectionID: "autem",
        ID: "600da0e3-aa61-4c6f-a09d-852b53b32c8c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PutTicketingConnectionIDCustomerIDRequest](../../models/operations/putticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.PutTicketingConnectionIDCustomerIDSecurity](../../models/operations/putticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.PutTicketingConnectionIDCustomerIDResponse](../../models/operations/putticketingconnectionidcustomeridresponse.md), error**


## PutTicketingConnectionIDNoteTicketIDID

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
    operationSecurity := operations.PutTicketingConnectionIDNoteTicketIDIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ticketing.PutTicketingConnectionIDNoteTicketIDID(ctx, operations.PutTicketingConnectionIDNoteTicketIDIDRequest{
        TicketingNote: &shared.TicketingNote{
            AgentID: unifiedto.String("dignissimos"),
            CreatedAt: types.MustTimeFromString("2022-06-05T09:20:14.169Z"),
            CustomerID: unifiedto.String("eligendi"),
            Description: unifiedto.String("esse"),
            ID: unifiedto.String("10e1673d-905c-4b4b-adef-3c127c390995"),
            Raw: shared.PropertyTicketingNoteRaw{},
            UpdatedAt: unifiedto.String("veniam"),
        },
        ConnectionID: "consequuntur",
        ID: "8250dcbb-cd3b-4121-b88c-1ee5e7a06139",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PutTicketingConnectionIDNoteTicketIDIDRequest](../../models/operations/putticketingconnectionidnoteticketididrequest.md)   | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.PutTicketingConnectionIDNoteTicketIDIDSecurity](../../models/operations/putticketingconnectionidnoteticketididsecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


### Response

**[*operations.PutTicketingConnectionIDNoteTicketIDIDResponse](../../models/operations/putticketingconnectionidnoteticketididresponse.md), error**


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
    res, err := s.Ticketing.PutTicketingConnectionIDTicketID(ctx, operations.PutTicketingConnectionIDTicketIDRequest{
        TicketingTicket: &shared.TicketingTicket{
            Category: unifiedto.String("placeat"),
            ClosedAt: types.MustTimeFromString("2021-06-17T09:59:15.139Z"),
            CreatedAt: types.MustTimeFromString("2021-01-22T11:03:32.954Z"),
            CustomerID: unifiedto.String("aut"),
            Description: unifiedto.String("soluta"),
            ID: unifiedto.String("7d176492-6b0c-4f5e-acb6-ebabe5e0b99f"),
            Priority: unifiedto.String("velit"),
            Raw: shared.PropertyTicketingTicketRaw{},
            Source: unifiedto.String("nobis"),
            SourceRef: unifiedto.String("illo"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedto.String("enim"),
            Tags: []string{
                "quas",
            },
            UpdatedAt: types.MustTimeFromString("2021-10-30T17:09:24.055Z"),
        },
        ConnectionID: "deserunt",
        ID: "87bb7aec-be56-49d7-8cb0-69907f989441",
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

