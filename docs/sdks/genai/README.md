# Genai

## Overview

### Available Operations

* [CreateGenaiAgent](#creategenaiagent) - Create an agent
* [CreateGenaiEmbedding](#creategenaiembedding) - Create an embedding
* [CreateGenaiOrganization](#creategenaiorganization) - Create an organization
* [CreateGenaiPrompt](#creategenaiprompt) - Create a prompt
* [CreateGenaiTask](#creategenaitask) - Create a task
* [GetGenaiAgent](#getgenaiagent) - Retrieve an agent
* [GetGenaiModel](#getgenaimodel) - Retrieve a model
* [GetGenaiOrganization](#getgenaiorganization) - Retrieve an organization
* [GetGenaiSkill](#getgenaiskill) - Retrieve a skill
* [GetGenaiTask](#getgenaitask) - Retrieve a task
* [ListGenaiAgents](#listgenaiagents) - List all agents
* [ListGenaiModels](#listgenaimodels) - List all models
* [ListGenaiOrganizations](#listgenaiorganizations) - List all organizations
* [ListGenaiSkills](#listgenaiskills) - List all skills
* [ListGenaiTasks](#listgenaitasks) - List all tasks
* [PatchGenaiAgent](#patchgenaiagent) - Update an agent
* [PatchGenaiOrganization](#patchgenaiorganization) - Update an organization
* [PatchGenaiTask](#patchgenaitask) - Update a task
* [RemoveGenaiAgent](#removegenaiagent) - Remove an agent
* [RemoveGenaiOrganization](#removegenaiorganization) - Remove an organization
* [RemoveGenaiSkill](#removegenaiskill) - Remove a skill
* [RemoveGenaiTask](#removegenaitask) - Remove a task
* [UpdateGenaiAgent](#updategenaiagent) - Update an agent
* [UpdateGenaiOrganization](#updategenaiorganization) - Update an organization
* [UpdateGenaiTask](#updategenaitask) - Update a task

## CreateGenaiAgent

Create an agent

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

    res, err := s.Genai.CreateGenaiAgent(ctx, operations.CreateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("62f1f2a6-a08f-4500-83e8-500f7d3d469a"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-13T13:18:49.775Z"),
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

## CreateGenaiEmbedding

Create an embedding

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiEmbedding" method="post" path="/genai/{connection_id}/embedding" example="genai_embedding" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Genai.CreateGenaiEmbedding(ctx, operations.CreateGenaiEmbeddingRequest{
        GenaiEmbedding: shared.GenaiEmbedding{
            Content: []shared.GenaiEmbeddingContent{
                shared.GenaiEmbeddingContent{
                    Text: "Utrimque temptatio pecco demulceo.",
                },
            },
            Dimension: unifiedgosdk.Pointer[float64](423.0),
            Embeddings: unifiedgosdk.Pointer("Est."),
            EncondingFormat: shared.EncondingFormatFloat.ToPointer(),
            ID: unifiedgosdk.Pointer("6718107c-9716-4061-a744-6d653edd6605"),
            MaxTokens: unifiedgosdk.Pointer[float64](223.0),
            TokensUsed: unifiedgosdk.Pointer[float64](836.0),
            Type: unifiedgosdk.Pointer("classification"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiEmbedding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateGenaiEmbeddingRequest](../../pkg/models/operations/creategenaiembeddingrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateGenaiEmbeddingResponse](../../pkg/models/operations/creategenaiembeddingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGenaiOrganization

Create an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiOrganization" method="post" path="/genai/{connection_id}/organization" example="genai_organization" -->
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

    res, err := s.Genai.CreateGenaiOrganization(ctx, operations.CreateGenaiOrganizationRequest{
        GenaiOrganization: shared.GenaiOrganization{
            CreatedAt: types.MustNewTimeFromString("2020-10-27T16:03:47.122Z"),
            Description: unifiedgosdk.Pointer("Voluptates abeo subseco."),
            ID: unifiedgosdk.Pointer("ed2d36b4-edbb-4b14-81d8-73b91574e018"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("officially about"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-25T13:57:32.616Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateGenaiOrganizationRequest](../../pkg/models/operations/creategenaiorganizationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateGenaiOrganizationResponse](../../pkg/models/operations/creategenaiorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGenaiPrompt

Create a prompt

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiPrompt" method="post" path="/genai/{connection_id}/prompt" example="genai_prompt" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Genai.CreateGenaiPrompt(ctx, operations.CreateGenaiPromptRequest{
        GenaiPrompt: shared.GenaiPrompt{
            Answers: []shared.GenaiAnswer{
                shared.GenaiAnswer{
                    Confidence: unifiedgosdk.Pointer[float64](0.5384309098590165),
                    ID: "TeZpueFB",
                    Probabilities: []shared.GenaiAnswerProbability{
                        shared.GenaiAnswerProbability{
                            Option: "veniam",
                            Probability: 0.05683182482607663,
                        },
                        shared.GenaiAnswerProbability{
                            Option: "comes",
                            Probability: 0.1404057613108307,
                        },
                    },
                    Score: unifiedgosdk.Pointer[float64](0.05974147981032729),
                    Type: shared.GenaiAnswerTypeChoice,
                    Value: unifiedgosdk.Pointer("summisse"),
                },
            },
            MaxTokens: unifiedgosdk.Pointer[float64](0.4677782787475735),
            McpAuthorizationToken: unifiedgosdk.Pointer("f45a6e93-7bed-49b4-a5c8-37a2ed2d58f4"),
            McpDeferredTools: []string{},
            McpURL: unifiedgosdk.Pointer("https://unsung-dusk.info/"),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Aegre repudiandae verecundia facere statua.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
                shared.GenaiContent{
                    Content: "Speciosus xiphias soleo trepide crinis.",
                    Images: []string{},
                    Role: shared.RoleSystem.ToPointer(),
                },
            },
            Questions: []shared.GenaiQuestion{
                shared.GenaiQuestion{
                    ID: "KtvbYiEq",
                    Instructions: "Admiratio cura contabesco subvenio venustas architecto vacuus summopere.",
                    Options: []string{
                        "vulnero",
                        "quasi",
                        "tolero",
                    },
                    Type: shared.GenaiQuestionTypeBoolean,
                },
                shared.GenaiQuestion{
                    ID: "oV6IURtv",
                    Instructions: "Celer deleniti ventus ipsam tendo odit.",
                    Options: []string{
                        "nulla",
                        "aspicio",
                        "tero",
                    },
                    Type: shared.GenaiQuestionTypeChoice,
                },
                shared.GenaiQuestion{
                    ID: "60dBZk2K",
                    Instructions: "Decretum vesica amet repudiandae.",
                    Options: []string{
                        "ducimus",
                        "civis",
                        "alter",
                    },
                    Type: shared.GenaiQuestionTypeBoolean,
                },
            },
            Responses: []string{
                "Balbus vobis circumvenio una.",
            },
            Temperature: unifiedgosdk.Pointer[float64](0.0),
            TokensUsed: unifiedgosdk.Pointer[float64](975.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiPrompt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateGenaiPromptRequest](../../pkg/models/operations/creategenaipromptrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateGenaiPromptResponse](../../pkg/models/operations/creategenaipromptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGenaiTask

Create a task

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiTask" method="post" path="/genai/{connection_id}/task" example="genai_task" -->
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

    res, err := s.Genai.CreateGenaiTask(ctx, operations.CreateGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-09-29T09:18:48.378Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("c1da3f04-2362-4df1-85cd-fe0a5e00d662"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-20T01:40:29.541Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-25T09:45:48.981Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateGenaiTaskRequest](../../pkg/models/operations/creategenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateGenaiTaskResponse](../../pkg/models/operations/creategenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiAgent

Retrieve an agent

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

    res, err := s.Genai.GetGenaiAgent(ctx, operations.GetGenaiAgentRequest{
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

## GetGenaiModel

Retrieve a model

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiModel" method="get" path="/genai/{connection_id}/model/{id}" -->
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

    res, err := s.Genai.GetGenaiModel(ctx, operations.GetGenaiModelRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiModel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetGenaiModelRequest](../../pkg/models/operations/getgenaimodelrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetGenaiModelResponse](../../pkg/models/operations/getgenaimodelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiOrganization

Retrieve an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiOrganization" method="get" path="/genai/{connection_id}/organization/{id}" -->
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

    res, err := s.Genai.GetGenaiOrganization(ctx, operations.GetGenaiOrganizationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetGenaiOrganizationRequest](../../pkg/models/operations/getgenaiorganizationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetGenaiOrganizationResponse](../../pkg/models/operations/getgenaiorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiSkill

Retrieve a skill

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiSkill" method="get" path="/genai/{connection_id}/skill/{id}" -->
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

    res, err := s.Genai.GetGenaiSkill(ctx, operations.GetGenaiSkillRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiSkill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetGenaiSkillRequest](../../pkg/models/operations/getgenaiskillrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetGenaiSkillResponse](../../pkg/models/operations/getgenaiskillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiTask

Retrieve a task

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiTask" method="get" path="/genai/{connection_id}/task/{id}" -->
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

    res, err := s.Genai.GetGenaiTask(ctx, operations.GetGenaiTaskRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetGenaiTaskRequest](../../pkg/models/operations/getgenaitaskrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetGenaiTaskResponse](../../pkg/models/operations/getgenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiAgents

List all agents

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

    res, err := s.Genai.ListGenaiAgents(ctx, operations.ListGenaiAgentsRequest{
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

## ListGenaiModels

List all models

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiModels" method="get" path="/genai/{connection_id}/model" -->
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

    res, err := s.Genai.ListGenaiModels(ctx, operations.ListGenaiModelsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiModels != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListGenaiModelsRequest](../../pkg/models/operations/listgenaimodelsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListGenaiModelsResponse](../../pkg/models/operations/listgenaimodelsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiOrganizations

List all organizations

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiOrganizations" method="get" path="/genai/{connection_id}/organization" -->
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

    res, err := s.Genai.ListGenaiOrganizations(ctx, operations.ListGenaiOrganizationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiOrganizations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListGenaiOrganizationsRequest](../../pkg/models/operations/listgenaiorganizationsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListGenaiOrganizationsResponse](../../pkg/models/operations/listgenaiorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiSkills

List all skills

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiSkills" method="get" path="/genai/{connection_id}/skill" -->
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

    res, err := s.Genai.ListGenaiSkills(ctx, operations.ListGenaiSkillsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiSkills != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListGenaiSkillsRequest](../../pkg/models/operations/listgenaiskillsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListGenaiSkillsResponse](../../pkg/models/operations/listgenaiskillsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiTasks

List all tasks

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiTasks" method="get" path="/genai/{connection_id}/task" -->
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

    res, err := s.Genai.ListGenaiTasks(ctx, operations.ListGenaiTasksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListGenaiTasksRequest](../../pkg/models/operations/listgenaitasksrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListGenaiTasksResponse](../../pkg/models/operations/listgenaitasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiAgent

Update an agent

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

    res, err := s.Genai.PatchGenaiAgent(ctx, operations.PatchGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("f3fd9a29-1a01-47a5-bb42-d86528edf993"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-13T13:18:49.782Z"),
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

## PatchGenaiOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiOrganization" method="patch" path="/genai/{connection_id}/organization/{id}" example="genai_organization" -->
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

    res, err := s.Genai.PatchGenaiOrganization(ctx, operations.PatchGenaiOrganizationRequest{
        GenaiOrganization: shared.GenaiOrganization{
            CreatedAt: types.MustNewTimeFromString("2020-10-27T16:03:47.122Z"),
            Description: unifiedgosdk.Pointer("Voluptates abeo subseco."),
            ID: unifiedgosdk.Pointer("7294496b-b4e0-45f7-b39b-afa32d2395f9"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("officially about"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-25T13:57:32.618Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchGenaiOrganizationRequest](../../pkg/models/operations/patchgenaiorganizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchGenaiOrganizationResponse](../../pkg/models/operations/patchgenaiorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiTask

Update a task

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiTask" method="patch" path="/genai/{connection_id}/task/{id}" example="genai_task" -->
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

    res, err := s.Genai.PatchGenaiTask(ctx, operations.PatchGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-09-29T09:18:48.392Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("03af4cc0-2e4e-497b-b29c-912254615cde"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-20T01:40:29.551Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-25T09:45:48.988Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchGenaiTaskRequest](../../pkg/models/operations/patchgenaitaskrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchGenaiTaskResponse](../../pkg/models/operations/patchgenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiAgent

Remove an agent

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

    res, err := s.Genai.RemoveGenaiAgent(ctx, operations.RemoveGenaiAgentRequest{
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

## RemoveGenaiOrganization

Remove an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiOrganization" method="delete" path="/genai/{connection_id}/organization/{id}" -->
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

    res, err := s.Genai.RemoveGenaiOrganization(ctx, operations.RemoveGenaiOrganizationRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveGenaiOrganizationRequest](../../pkg/models/operations/removegenaiorganizationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveGenaiOrganizationResponse](../../pkg/models/operations/removegenaiorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiSkill

Remove a skill

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiSkill" method="delete" path="/genai/{connection_id}/skill/{id}" -->
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

    res, err := s.Genai.RemoveGenaiSkill(ctx, operations.RemoveGenaiSkillRequest{
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
| `request`                                                                                    | [operations.RemoveGenaiSkillRequest](../../pkg/models/operations/removegenaiskillrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveGenaiSkillResponse](../../pkg/models/operations/removegenaiskillresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiTask

Remove a task

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiTask" method="delete" path="/genai/{connection_id}/task/{id}" -->
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

    res, err := s.Genai.RemoveGenaiTask(ctx, operations.RemoveGenaiTaskRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveGenaiTaskRequest](../../pkg/models/operations/removegenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveGenaiTaskResponse](../../pkg/models/operations/removegenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiAgent

Update an agent

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

    res, err := s.Genai.UpdateGenaiAgent(ctx, operations.UpdateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Assentator tero sequi."),
            ID: unifiedgosdk.Pointer("f3fd9a29-1a01-47a5-bb42-d86528edf993"),
            Instructions: unifiedgosdk.Pointer("Delicate tabella addo vita ver auctus corrumpo. Conscendo auctus maiores astrum sulum aufero bis stella. Non varius sto solus admitto vado suggero theatrum."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("daughter with"),
            UpdatedAt: types.MustNewTimeFromString("2025-10-13T13:18:49.782Z"),
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

## UpdateGenaiOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiOrganization" method="put" path="/genai/{connection_id}/organization/{id}" example="genai_organization" -->
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

    res, err := s.Genai.UpdateGenaiOrganization(ctx, operations.UpdateGenaiOrganizationRequest{
        GenaiOrganization: shared.GenaiOrganization{
            CreatedAt: types.MustNewTimeFromString("2020-10-27T16:03:47.122Z"),
            Description: unifiedgosdk.Pointer("Voluptates abeo subseco."),
            ID: unifiedgosdk.Pointer("7294496b-b4e0-45f7-b39b-afa32d2395f9"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("officially about"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-25T13:57:32.618Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateGenaiOrganizationRequest](../../pkg/models/operations/updategenaiorganizationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateGenaiOrganizationResponse](../../pkg/models/operations/updategenaiorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiTask

Update a task

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiTask" method="put" path="/genai/{connection_id}/task/{id}" example="genai_task" -->
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

    res, err := s.Genai.UpdateGenaiTask(ctx, operations.UpdateGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-09-29T09:18:48.392Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("03af4cc0-2e4e-497b-b29c-912254615cde"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-20T01:40:29.551Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-25T09:45:48.988Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateGenaiTaskRequest](../../pkg/models/operations/updategenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateGenaiTaskResponse](../../pkg/models/operations/updategenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |