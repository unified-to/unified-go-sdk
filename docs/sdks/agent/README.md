# Agent
(*Agent*)

### Available Operations

* [DeleteTicketingConnectionIDAgentID](#deleteticketingconnectionidagentid) - Remove a agent
* [GetTicketingConnectionIDAgent](#getticketingconnectionidagent) - List all agents
* [GetTicketingConnectionIDAgentID](#getticketingconnectionidagentid) - Retrieve a agent
* [GetUcConnectionIDAgent](#getucconnectionidagent) - List all agents
* [PatchTicketingConnectionIDAgentID](#patchticketingconnectionidagentid) - Update a agent
* [PostTicketingConnectionIDAgent](#postticketingconnectionidagent) - Create a agent
* [PutTicketingConnectionIDAgentID](#putticketingconnectionidagentid) - Update a agent

## DeleteTicketingConnectionIDAgentID

Remove a agent

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
    res, err := s.Agent.DeleteTicketingConnectionIDAgentID(ctx, operations.DeleteTicketingConnectionIDAgentIDRequest{
        ConnectionID: "navigate",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.DeleteTicketingConnectionIDAgentIDRequest](../../models/operations/deleteticketingconnectionidagentidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.DeleteTicketingConnectionIDAgentIDResponse](../../models/operations/deleteticketingconnectionidagentidresponse.md), error**


## GetTicketingConnectionIDAgent

List all agents

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
    res, err := s.Agent.GetTicketingConnectionIDAgent(ctx, operations.GetTicketingConnectionIDAgentRequest{
        ConnectionID: "East Steel Frozen",
        Limit: unifiedgosdk.Float64(8285.04),
        Offset: unifiedgosdk.Float64(5507.07),
        Order: unifiedgosdk.String("Korea West Ryan"),
        Query: unifiedgosdk.String("invoice coulomb soluta"),
        Sort: unifiedgosdk.String("adored"),
        UpdatedGte: types.MustTimeFromString("2023-11-15T19:25:12.859Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetTicketingConnectionIDAgentRequest](../../models/operations/getticketingconnectionidagentrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


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
    res, err := s.Agent.GetTicketingConnectionIDAgentID(ctx, operations.GetTicketingConnectionIDAgentIDRequest{
        ConnectionID: "Hat gas Cisgender",
        ID: "<ID>",
    })
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
| `request`                                                                                                              | [operations.GetTicketingConnectionIDAgentIDRequest](../../models/operations/getticketingconnectionidagentidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetTicketingConnectionIDAgentIDResponse](../../models/operations/getticketingconnectionidagentidresponse.md), error**


## GetUcConnectionIDAgent

List all agents

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
    res, err := s.Agent.GetUcConnectionIDAgent(ctx, operations.GetUcConnectionIDAgentRequest{
        ConnectionID: "Regional East Sedan",
        ContactID: unifiedgosdk.String("blue"),
        Limit: unifiedgosdk.Float64(7827.68),
        Offset: unifiedgosdk.Float64(2116.69),
        Order: unifiedgosdk.String("Bicycle"),
        Query: unifiedgosdk.String("Bacon officia iterate"),
        Sort: unifiedgosdk.String("sticky vote lumen"),
        UpdatedGte: types.MustTimeFromString("2021-07-05T19:53:29.041Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUcConnectionIDAgentRequest](../../models/operations/getucconnectionidagentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetUcConnectionIDAgentResponse](../../models/operations/getucconnectionidagentresponse.md), error**


## PatchTicketingConnectionIDAgentID

Update a agent

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
    res, err := s.Agent.PatchTicketingConnectionIDAgentID(ctx, operations.PatchTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-06-01T22:24:40.372Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Antonette63@gmail.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Hop"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "driver",
                    Type: shared.TicketingTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-07-09T08:35:36.354Z"),
        },
        ConnectionID: "Soft Diesel Springs",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PatchTicketingConnectionIDAgentIDRequest](../../models/operations/patchticketingconnectionidagentidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PatchTicketingConnectionIDAgentIDResponse](../../models/operations/patchticketingconnectionidagentidresponse.md), error**


## PostTicketingConnectionIDAgent

Create a agent

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
    res, err := s.Agent.PostTicketingConnectionIDAgent(ctx, operations.PostTicketingConnectionIDAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-12-14T10:20:29.412Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Eleazar_Beatty22@gmail.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("exploit our wireless"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "Korea wireless Ferrari",
                    Type: shared.TicketingTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-24T13:41:54.208Z"),
        },
        ConnectionID: "capacity copy Blues",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostTicketingConnectionIDAgentRequest](../../models/operations/postticketingconnectionidagentrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostTicketingConnectionIDAgentResponse](../../models/operations/postticketingconnectionidagentresponse.md), error**


## PutTicketingConnectionIDAgentID

Update a agent

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
    res, err := s.Agent.PutTicketingConnectionIDAgentID(ctx, operations.PutTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustTimeFromString("2022-12-19T19:47:13.993Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Augustus_Kessler34@hotmail.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("reintermediate impression Refined"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "asynchronous",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-04-29T18:13:42.824Z"),
        },
        ConnectionID: "synergistic Uzbekistan green",
        ID: "<ID>",
    })
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
| `request`                                                                                                              | [operations.PutTicketingConnectionIDAgentIDRequest](../../models/operations/putticketingconnectionidagentidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PutTicketingConnectionIDAgentIDResponse](../../models/operations/putticketingconnectionidagentidresponse.md), error**

