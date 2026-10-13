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
                "da3cafb1-c542-4cc6-ad51-47121e343c42",
                "4a56a105-ae8d-455e-9a33-a0553a9430d5",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-08-08T08:26:35.006Z"),
            ID: unifiedgosdk.Pointer("358b593f-b0c7-4a45-829d-3d2683396328"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("1f2e701b-d366-4707-a86d-51b904ee2fa4"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-06-06T08:54:12.618Z"),
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
                "7469f9ae-a4f9-4ab7-bc61-17bb8317faff",
                "f18bbe6d-b03e-41de-a35b-51810709a3ae",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-08-08T08:26:35.016Z"),
            ID: unifiedgosdk.Pointer("3be4e047-ebf1-4f06-b139-c18428536397"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("ff7539bc-93f5-4d46-b17e-ea2ab59f145a"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-06-06T08:54:12.623Z"),
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
                "7469f9ae-a4f9-4ab7-bc61-17bb8317faff",
                "f18bbe6d-b03e-41de-a35b-51810709a3ae",
            },
            CategoryID: unifiedgosdk.Pointer("vilicus"),
            CreatedAt: types.MustNewTimeFromString("2021-06-25T19:19:31.279Z"),
            Description: unifiedgosdk.Pointer("Cura dignissimos aut clibanus vulgaris patrocinor. Laborum acies curiositas antepono coniuratio. Correptius curiositas sono censura coma. Bestia suus tot cotidie terror subito coniecto beneficium."),
            DueAt: types.MustNewTimeFromString("2025-08-08T08:26:35.016Z"),
            ID: unifiedgosdk.Pointer("3be4e047-ebf1-4f06-b139-c18428536397"),
            Priority: unifiedgosdk.Pointer("LOW"),
            Source: unifiedgosdk.Pointer("atavus"),
            SourceRef: unifiedgosdk.Pointer("ff7539bc-93f5-4d46-b17e-ea2ab59f145a"),
            Status: shared.TicketingTicketStatusActive.ToPointer(),
            Subject: unifiedgosdk.Pointer("Thymbra ratione minus arbitro tricesimus cetera validus."),
            Tags: []string{
                "tamen",
                "vitae",
                "torrens",
            },
            UpdatedAt: types.MustNewTimeFromString("2023-06-06T08:54:12.623Z"),
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