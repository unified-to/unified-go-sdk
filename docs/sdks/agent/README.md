# Agent

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Agent.DeleteTicketingConnectionIDAgentID(ctx, operations.DeleteTicketingConnectionIDAgentIDRequest{
        ConnectionID: "perferendis",
        ID: "5dfc2ddf-7cc7-48ca-9ba9-28fc816742cb",
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
    res, err := s.Agent.GetTicketingConnectionIDAgent(ctx, operations.GetTicketingConnectionIDAgentRequest{
        ConnectionID: "esse",
        Limit: unifiedto.Float64(2165.5),
        Offset: unifiedto.Float64(5684.34),
        Order: unifiedto.String("aspernatur"),
        Query: unifiedto.String("perferendis"),
        Sort: unifiedto.String("ad"),
        UpdatedGte: types.MustDateFromString("2022-09-13"),
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
    res, err := s.Agent.GetTicketingConnectionIDAgentID(ctx, operations.GetTicketingConnectionIDAgentIDRequest{
        ConnectionID: "iste",
        ID: "396fea75-96eb-410f-aaa2-352c5955907a",
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


## GetUcConnectionIDAgent

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
    operationSecurity := operations.GetUcConnectionIDAgentSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Agent.GetUcConnectionIDAgent(ctx, operations.GetUcConnectionIDAgentRequest{
        ConnectionID: "doloribus",
        ContactID: unifiedto.String("sapiente"),
        Limit: unifiedto.Float64(1020.44),
        Offset: unifiedto.Float64(6527.9),
        Order: unifiedto.String("dolorem"),
        Query: unifiedto.String("culpa"),
        Sort: unifiedto.String("consequuntur"),
        UpdatedGte: types.MustDateFromString("2021-01-15"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetUcConnectionIDAgentRequest](../../models/operations/getucconnectionidagentrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetUcConnectionIDAgentSecurity](../../models/operations/getucconnectionidagentsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
    res, err := s.Agent.PatchTicketingConnectionIDAgentID(ctx, operations.PatchTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustDateFromString("2022-06-30"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Jamil62@yahoo.com",
                    Type: shared.TicketingEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("51aa52c3-f5ad-4019-9a1f-fe78f097b007"),
            Name: unifiedto.String("Shawna Carter"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "iusto",
                    Type: shared.TicketingTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2022-05-13"),
        },
        ConnectionID: "accusamus",
        ID: "6e13b99d-488e-41e9-9e45-0ad2abd44269",
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
    res, err := s.Agent.PostTicketingConnectionIDAgent(ctx, operations.PostTicketingConnectionIDAgentRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustDateFromString("2022-12-17"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Rhoda14@gmail.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("94bb4f63-c969-4e9a-befa-77dfb14cd66a"),
            Name: unifiedto.String("Alfred McClure"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "delectus",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2021-07-20"),
        },
        ConnectionID: "id",
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
    res, err := s.Agent.PutTicketingConnectionIDAgentID(ctx, operations.PutTicketingConnectionIDAgentIDRequest{
        TicketingAgent: &shared.TicketingAgent{
            CreatedAt: types.MustDateFromString("2021-12-07"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Daren.Nolan@hotmail.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("97074ba4-469b-46e2-9419-59890afa563e"),
            Name: unifiedto.String("Vivian Boyle"),
            Raw: shared.PropertyTicketingAgentRaw{},
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "debitis",
                    Type: shared.TicketingTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2021-05-22"),
        },
        ConnectionID: "facilis",
        ID: "711e5b7f-d2ed-4028-921c-ddc692601fb5",
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

