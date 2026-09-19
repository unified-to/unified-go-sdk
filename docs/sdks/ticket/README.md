# Ticket

## Overview

### Available Operations

* [CreateTicketingTicket](#createticketingticket) - Create a ticket
* [GetTicketingTicket](#getticketingticket) - Retrieve a ticket
* [ListTicketingTickets](#listticketingtickets) - List all tickets
* [PatchTicketingTicket](#patchticketingticket) - Update a ticket
* [RemoveTicketingTicket](#removeticketingticket) - Remove a ticket
* [UpdateTicketingTicket](#updateticketingticket) - Update a ticket

## CreateTicketingTicket

Create a ticket

### Example Usage

<!-- UsageSnippet language="go" operationID="createTicketingTicket" method="post" path="/ticketing/{connection_id}/ticket" example="ticketing_ticket" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ticket.CreateTicketingTicket(ctx, operations.CreateTicketingTicketRequest{
        TicketingTicket: shared.TicketingTicket{
            AttachmentIds: []string{
                "95ed4534-71ee-477b-a10d-651c42c33a51",
                "7bff6d6d-8393-4b66-8e1d-3c1e4c0844d1",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T21:20:40.361Z"),
            ID: unifiedgosdk.Pointer("8fb30a8f-0ca1-49ec-bfdc-7bdb5a081c51"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("b9bae82f-ddaf-431d-9b43-e196d1e91da5"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T15:38:14.571Z"),
            URL: unifiedgosdk.Pointer("https://yellowish-testimonial.biz"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateTicketingTicketRequest](../../pkg/models/operations/createticketingticketrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateTicketingTicketResponse](../../pkg/models/operations/createticketingticketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTicketingTicket

Retrieve a ticket

### Example Usage

<!-- UsageSnippet language="go" operationID="getTicketingTicket" method="get" path="/ticketing/{connection_id}/ticket/{id}" -->
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

    res, err := s.Ticket.GetTicketingTicket(ctx, operations.GetTicketingTicketRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetTicketingTicketRequest](../../pkg/models/operations/getticketingticketrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetTicketingTicketResponse](../../pkg/models/operations/getticketingticketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTicketingTickets

List all tickets

### Example Usage

<!-- UsageSnippet language="go" operationID="listTicketingTickets" method="get" path="/ticketing/{connection_id}/ticket" -->
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

    res, err := s.Ticket.ListTicketingTickets(ctx, operations.ListTicketingTicketsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListTicketingTicketsRequest](../../pkg/models/operations/listticketingticketsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListTicketingTicketsResponse](../../pkg/models/operations/listticketingticketsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchTicketingTicket

Update a ticket

### Example Usage

<!-- UsageSnippet language="go" operationID="patchTicketingTicket" method="patch" path="/ticketing/{connection_id}/ticket/{id}" example="ticketing_ticket" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ticket.PatchTicketingTicket(ctx, operations.PatchTicketingTicketRequest{
        TicketingTicket: shared.TicketingTicket{
            AttachmentIds: []string{
                "3270bdeb-240b-46f7-a3c0-6b8ca0f467e7",
                "df48ad30-1f07-414f-981f-88380d04b872",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T21:20:40.368Z"),
            ID: unifiedgosdk.Pointer("3389b15d-555e-4863-bfc9-8139655b2140"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("792d53ee-5438-4d13-bb31-fa6123dc48af"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T15:38:14.574Z"),
            URL: unifiedgosdk.Pointer("https://yellowish-testimonial.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchTicketingTicketRequest](../../pkg/models/operations/patchticketingticketrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchTicketingTicketResponse](../../pkg/models/operations/patchticketingticketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveTicketingTicket

Remove a ticket

### Example Usage

<!-- UsageSnippet language="go" operationID="removeTicketingTicket" method="delete" path="/ticketing/{connection_id}/ticket/{id}" -->
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

    res, err := s.Ticket.RemoveTicketingTicket(ctx, operations.RemoveTicketingTicketRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveTicketingTicketRequest](../../pkg/models/operations/removeticketingticketrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveTicketingTicketResponse](../../pkg/models/operations/removeticketingticketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTicketingTicket

Update a ticket

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTicketingTicket" method="put" path="/ticketing/{connection_id}/ticket/{id}" example="ticketing_ticket" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Ticket.UpdateTicketingTicket(ctx, operations.UpdateTicketingTicketRequest{
        TicketingTicket: shared.TicketingTicket{
            AttachmentIds: []string{
                "3270bdeb-240b-46f7-a3c0-6b8ca0f467e7",
                "df48ad30-1f07-414f-981f-88380d04b872",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T21:20:40.368Z"),
            ID: unifiedgosdk.Pointer("3389b15d-555e-4863-bfc9-8139655b2140"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("792d53ee-5438-4d13-bb31-fa6123dc48af"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T15:38:14.574Z"),
            URL: unifiedgosdk.Pointer("https://yellowish-testimonial.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateTicketingTicketRequest](../../pkg/models/operations/updateticketingticketrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateTicketingTicketResponse](../../pkg/models/operations/updateticketingticketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |