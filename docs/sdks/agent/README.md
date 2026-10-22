# Agent

## Overview

### Available Operations

* [CreateGenaiAgent](#creategenaiagent) - Create an agent
* [GetGenaiAgent](#getgenaiagent) - Retrieve an agent
* [ListGenaiAgents](#listgenaiagents) - List all agents
* [PatchGenaiAgent](#patchgenaiagent) - Update an agent
* [RemoveGenaiAgent](#removegenaiagent) - Remove an agent
* [UpdateGenaiAgent](#updategenaiagent) - Update an agent

## CreateGenaiAgent

Creates a GenaiAgent. Provide the unified GenaiAgent fields in the request body; they are mapped to the provider's API. Returns the created GenaiAgent.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiAgent" method="post" path="/genai/{connection_id}/agent" example="genai_agent" -->
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

    res, err := s.Agent.CreateGenaiAgent(ctx, operations.CreateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("e6383f2d-b4d2-489e-9384-0cca840ad6f3"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-16T06:26:17.706Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateGenaiAgentRequest](../../pkg/models/operations/creategenaiagentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateGenaiAgentResponse](../../pkg/models/operations/creategenaiagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiAgent

Retrieves a single GenaiAgent by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiAgent" method="get" path="/genai/{connection_id}/agent/{id}" -->
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

    res, err := s.Agent.GetGenaiAgent(ctx, operations.GetGenaiAgentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetGenaiAgentRequest](../../pkg/models/operations/getgenaiagentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetGenaiAgentResponse](../../pkg/models/operations/getgenaiagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiAgents

Returns a list of GenaiAgents. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiAgents" method="get" path="/genai/{connection_id}/agent" -->
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

    res, err := s.Agent.ListGenaiAgents(ctx, operations.ListGenaiAgentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListGenaiAgentsRequest](../../pkg/models/operations/listgenaiagentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListGenaiAgentsResponse](../../pkg/models/operations/listgenaiagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiAgent

Updates the GenaiAgent identified by `id`. Provide the unified GenaiAgent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiAgent" method="patch" path="/genai/{connection_id}/agent/{id}" example="genai_agent" -->
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

    res, err := s.Agent.PatchGenaiAgent(ctx, operations.PatchGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("691117f4-f6f3-456b-90b6-e5ab838c6c45"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-16T06:26:17.713Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchGenaiAgentRequest](../../pkg/models/operations/patchgenaiagentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchGenaiAgentResponse](../../pkg/models/operations/patchgenaiagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiAgent

Deletes the GenaiAgent identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiAgent" method="delete" path="/genai/{connection_id}/agent/{id}" -->
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

    res, err := s.Agent.RemoveGenaiAgent(ctx, operations.RemoveGenaiAgentRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveGenaiAgentRequest](../../pkg/models/operations/removegenaiagentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveGenaiAgentResponse](../../pkg/models/operations/removegenaiagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiAgent

Updates the GenaiAgent identified by `id`. Provide the unified GenaiAgent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiAgent" method="put" path="/genai/{connection_id}/agent/{id}" example="genai_agent" -->
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

    res, err := s.Agent.UpdateGenaiAgent(ctx, operations.UpdateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("691117f4-f6f3-456b-90b6-e5ab838c6c45"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-16T06:26:17.713Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiAgent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateGenaiAgentRequest](../../pkg/models/operations/updategenaiagentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateGenaiAgentResponse](../../pkg/models/operations/updategenaiagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |