# Agent
(*Agent*)

### Available Operations

* [CreateTicketingAgent](#createticketingagent) - Create a agent
* [GetTicketingAgent](#getticketingagent) - Retrieve a agent
* [ListTicketingAgents](#listticketingagents) - List all agents
* [ListUcAgents](#listucagents) - List all agents
* [PatchTicketingAgent](#patchticketingagent) - Update a agent
* [RemoveTicketingAgent](#removeticketingagent) - Remove a agent
* [UpdateTicketingAgent](#updateticketingagent) - Update a agent

## CreateTicketingAgent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.CreateTicketingAgent(ctx, operations.CreateTicketingAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Paolo.Cole8@yahoo.com",
                },
            },
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "scarcely Soap navigating",
                },
            },
        },
        ConnectionID: "smoothly Algeria",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateTicketingAgentRequest](../../models/operations/createticketingagentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateTicketingAgentResponse](../../models/operations/createticketingagentresponse.md), error**


## GetTicketingAgent

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.GetTicketingAgent(ctx, operations.GetTicketingAgentRequest{
        ConnectionID: "Lamborghini",
        Fields: []string{
            "SAS",
        },
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTicketingAgentRequest](../../models/operations/getticketingagentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTicketingAgentResponse](../../models/operations/getticketingagentresponse.md), error**


## ListTicketingAgents

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.ListTicketingAgents(ctx, operations.ListTicketingAgentsRequest{
        ConnectionID: "Mazda",
        Fields: []string{
            "Home",
        },
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTicketingAgentsRequest](../../models/operations/listticketingagentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListTicketingAgentsResponse](../../models/operations/listticketingagentsresponse.md), error**


## ListUcAgents

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.ListUcAgents(ctx, operations.ListUcAgentsRequest{
        ConnectionID: "Representative",
        Fields: []string{
            "olive",
        },
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListUcAgentsRequest](../../models/operations/listucagentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListUcAgentsResponse](../../models/operations/listucagentsresponse.md), error**


## PatchTicketingAgent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.PatchTicketingAgent(ctx, operations.PatchTicketingAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Danyka87@hotmail.com",
                },
            },
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "envious",
                },
            },
        },
        ConnectionID: "by",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchTicketingAgentRequest](../../models/operations/patchticketingagentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PatchTicketingAgentResponse](../../models/operations/patchticketingagentresponse.md), error**


## RemoveTicketingAgent

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.RemoveTicketingAgent(ctx, operations.RemoveTicketingAgentRequest{
        ConnectionID: "monitor",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveTicketingAgentRequest](../../models/operations/removeticketingagentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RemoveTicketingAgentResponse](../../models/operations/removeticketingagentresponse.md), error**


## UpdateTicketingAgent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Agent.UpdateTicketingAgent(ctx, operations.UpdateTicketingAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Lorenz39@hotmail.com",
                },
            },
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "female",
                },
            },
        },
        ConnectionID: "beside Usability Bedfordshire",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTicketingAgentRequest](../../models/operations/updateticketingagentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateTicketingAgentResponse](../../models/operations/updateticketingagentresponse.md), error**

