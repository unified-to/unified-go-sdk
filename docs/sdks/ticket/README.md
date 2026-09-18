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
                "69425dd4-f539-4065-b5bc-480248461157",
                "c86192d9-3698-4f71-b0d2-6b66f4814c84",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T04:44:17.194Z"),
            ID: unifiedgosdk.Pointer("a5de714a-b3fd-40d9-b93e-506f6876e3ea"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("6861883d-19d8-4462-bf4c-9c74cbf2cd43"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T07:47:40.236Z"),
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
                "fd0c224e-47d7-4b83-86e9-ea7749c489e0",
                "b29bb217-5c4c-47f8-b85a-138890e9b662",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T04:44:17.204Z"),
            ID: unifiedgosdk.Pointer("9d7a3dfb-281a-454f-92eb-122a5a1b5f6c"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("5963f755-1720-41d8-92fe-7b2d6d2c852e"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T07:47:40.241Z"),
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
                "fd0c224e-47d7-4b83-86e9-ea7749c489e0",
                "b29bb217-5c4c-47f8-b85a-138890e9b662",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-07-20T04:44:17.204Z"),
            ID: unifiedgosdk.Pointer("9d7a3dfb-281a-454f-92eb-122a5a1b5f6c"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("5963f755-1720-41d8-92fe-7b2d6d2c852e"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-05-28T07:47:40.241Z"),
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