# Genai

## Overview

### Available Operations

* [CreateGenaiAgent](#creategenaiagent) - Create an agent
* [CreateGenaiEmbedding](#creategenaiembedding) - Create an embedding
* [CreateGenaiMemory](#creategenaimemory) - Create a memory
* [CreateGenaiOrganization](#creategenaiorganization) - Create an organization
* [CreateGenaiPrompt](#creategenaiprompt) - Create a prompt
* [CreateGenaiTask](#creategenaitask) - Create a task
* [GetGenaiAgent](#getgenaiagent) - Retrieve an agent
* [GetGenaiMemory](#getgenaimemory) - Retrieve a memory
* [GetGenaiModel](#getgenaimodel) - Retrieve a model
* [GetGenaiOrganization](#getgenaiorganization) - Retrieve an organization
* [GetGenaiSkill](#getgenaiskill) - Retrieve a skill
* [GetGenaiTask](#getgenaitask) - Retrieve a task
* [ListGenaiAgents](#listgenaiagents) - List all agents
* [ListGenaiMemories](#listgenaimemories) - List all memories
* [ListGenaiModels](#listgenaimodels) - List all models
* [ListGenaiOrganizations](#listgenaiorganizations) - List all organizations
* [ListGenaiSkills](#listgenaiskills) - List all skills
* [ListGenaiTasks](#listgenaitasks) - List all tasks
* [PatchGenaiAgent](#patchgenaiagent) - Update an agent
* [PatchGenaiMemory](#patchgenaimemory) - Update a memory
* [PatchGenaiOrganization](#patchgenaiorganization) - Update an organization
* [PatchGenaiTask](#patchgenaitask) - Update a task
* [RemoveGenaiAgent](#removegenaiagent) - Remove an agent
* [RemoveGenaiMemory](#removegenaimemory) - Remove a memory
* [RemoveGenaiOrganization](#removegenaiorganization) - Remove an organization
* [RemoveGenaiSkill](#removegenaiskill) - Remove a skill
* [RemoveGenaiTask](#removegenaitask) - Remove a task
* [UpdateGenaiAgent](#updategenaiagent) - Update an agent
* [UpdateGenaiMemory](#updategenaimemory) - Update a memory
* [UpdateGenaiOrganization](#updategenaiorganization) - Update an organization
* [UpdateGenaiTask](#updategenaitask) - Update a task

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

    res, err := s.Genai.CreateGenaiAgent(ctx, operations.CreateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Sperno despecto."),
            ID: unifiedgosdk.Pointer("d291a0bc-e6b6-4b22-ac5b-cdcb95f5636e"),
            Instructions: unifiedgosdk.Pointer("Addo vita ver auctus corrumpo substantia conscendo auctus. Astrum sulum aufero bis stella sumptus non."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("testimonial suddenly"),
            UpdatedAt: types.MustNewTimeFromString("2025-11-14T03:39:19.695Z"),
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

Creates a GenaiEmbedding. Provide the unified GenaiEmbedding fields in the request body; they are mapped to the provider's API. Returns the created GenaiEmbedding.

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
                    Text: "Demulceo adulescens est delectatio solitudo videlicet volo.",
                },
            },
            Dimension: unifiedgosdk.Pointer[float64](423.0),
            Embeddings: unifiedgosdk.Pointer("Venia tutis cernuus viscus officiis utilis porro utique."),
            EncodingFormat: shared.EncodingFormatFloat.ToPointer(),
            ID: unifiedgosdk.Pointer("665804d4-bba4-4aca-ab81-29fc9b08b0b9"),
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

## CreateGenaiMemory

Creates a GenaiMemory. Provide the unified GenaiMemory fields in the request body; they are mapped to the provider's API. Returns the created GenaiMemory.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiMemory" method="post" path="/genai/{connection_id}/memory" example="genai_memory" -->
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

    res, err := s.Genai.CreateGenaiMemory(ctx, operations.CreateGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("a7144def-c926-4a0a-b7a8-ace51acb5c2c"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-06T17:51:09.193Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateGenaiMemoryRequest](../../pkg/models/operations/creategenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateGenaiMemoryResponse](../../pkg/models/operations/creategenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGenaiOrganization

Creates a GenaiOrganization. Provide the unified GenaiOrganization fields in the request body; they are mapped to the provider's API. Returns the created GenaiOrganization.

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
            Description: unifiedgosdk.Pointer("Subito teneo arbor."),
            ID: unifiedgosdk.Pointer("d24f6277-e6a5-419e-976f-e329ebef60b7"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("yum yuck"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-08T22:49:24.772Z"),
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

Creates a GenaiPrompt. Provide the unified GenaiPrompt fields in the request body; they are mapped to the provider's API. Returns the created GenaiPrompt.

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
                    ID: "DiYHdUlr",
                    Probabilities: []shared.GenaiAnswerProbability{
                        shared.GenaiAnswerProbability{
                            Option: "ulterius",
                            Probability: 0.05683182482607663,
                        },
                        shared.GenaiAnswerProbability{
                            Option: "spiculum",
                            Probability: 0.1404057613108307,
                        },
                    },
                    Score: unifiedgosdk.Pointer[float64](0.05974147981032729),
                    Type: shared.GenaiAnswerTypeChoice,
                    Value: unifiedgosdk.Pointer("vilicus"),
                },
            },
            MaxTokens: unifiedgosdk.Pointer[float64](0.4677782787475735),
            McpAuthorizationToken: unifiedgosdk.Pointer("45c837a2-ed2d-458f-b4bc-782f4a6b065b"),
            McpDeferredTools: []string{},
            McpURL: unifiedgosdk.Pointer("https://svelte-offset.net"),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Facere statua depulso speciosus xiphias soleo trepide crinis crapula balbus.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
                shared.GenaiContent{
                    Content: "Circumvenio una demitto ventito chirographum culpo vos cohibeo complectus speciosus.",
                    Images: []string{},
                    Role: shared.RoleSystem.ToPointer(),
                },
            },
            Questions: []shared.GenaiQuestion{
                shared.GenaiQuestion{
                    ID: "9SIZCOoV",
                    Instructions: "Supellex verbera universe.",
                    Options: []string{
                        "dens",
                        "doloribus",
                        "degusto",
                    },
                    Type: shared.GenaiQuestionTypeBoolean,
                },
                shared.GenaiQuestion{
                    ID: "gsUyLAAa",
                    Instructions: "Amplexus abundans caelestis patrocinor vulgivagus condico adduco tego apostolus.",
                    Options: []string{
                        "decretum",
                        "vesica",
                        "amet",
                    },
                    Type: shared.GenaiQuestionTypeChoice,
                },
                shared.GenaiQuestion{
                    ID: "Dwh6TeZp",
                    Instructions: "Capio somnus pauper summisse veniam comes.",
                    Options: []string{
                        "depulso",
                        "deficio",
                        "sustineo",
                    },
                    Type: shared.GenaiQuestionTypeBoolean,
                },
            },
            Responses: []string{
                "Vereor pecco avarus derideo sumptus.",
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

Creates a GenaiTask. Provide the unified GenaiTask fields in the request body; they are mapped to the provider's API. Returns the created GenaiTask.

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
            CompletedAt: types.MustNewTimeFromString("2025-10-30T21:35:21.702Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("4597eacf-1ca5-4494-9bcb-c687affdf289"),
            Instructions: unifiedgosdk.Pointer("Angustus cubo cupio abundans admoveo sumptus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Amet contigo tristis deripio auxilium cur aeternus cogito.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect online monitor"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-06-11T20:56:51.389Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Tenus via tamisium denego teres adsidue cura crur. Coma defungo magnam delicate confero amaritudo varius."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/eek-save"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T08:03:38.850Z"),
            WebURL: unifiedgosdk.Pointer("https://monstrous-honesty.net"),
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

## GetGenaiMemory

Retrieves a single GenaiMemory by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiMemory" method="get" path="/genai/{connection_id}/memory/{id}" -->
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

    res, err := s.Genai.GetGenaiMemory(ctx, operations.GetGenaiMemoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetGenaiMemoryRequest](../../pkg/models/operations/getgenaimemoryrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetGenaiMemoryResponse](../../pkg/models/operations/getgenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiModel

Retrieves a single GenaiModel by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single GenaiOrganization by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single GenaiSkill by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single GenaiTask by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

## ListGenaiMemories

Returns a list of GenaiMemories. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiMemories" method="get" path="/genai/{connection_id}/memory" -->
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

    res, err := s.Genai.ListGenaiMemories(ctx, operations.ListGenaiMemoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListGenaiMemoriesRequest](../../pkg/models/operations/listgenaimemoriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListGenaiMemoriesResponse](../../pkg/models/operations/listgenaimemoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiModels

Returns a list of GenaiModels. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of GenaiOrganizations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of GenaiSkills. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of GenaiTasks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Genai.PatchGenaiAgent(ctx, operations.PatchGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Sperno despecto."),
            ID: unifiedgosdk.Pointer("3e36559a-d979-42c3-9b9e-648f1aa534e8"),
            Instructions: unifiedgosdk.Pointer("Addo vita ver auctus corrumpo substantia conscendo auctus. Astrum sulum aufero bis stella sumptus non."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("testimonial suddenly"),
            UpdatedAt: types.MustNewTimeFromString("2025-11-14T03:39:19.700Z"),
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

## PatchGenaiMemory

Updates the GenaiMemory identified by `id`. Provide the unified GenaiMemory fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiMemory" method="patch" path="/genai/{connection_id}/memory/{id}" example="genai_memory" -->
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

    res, err := s.Genai.PatchGenaiMemory(ctx, operations.PatchGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("c3894b70-1722-48ab-9297-af7f4f60ad2b"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-06T17:51:09.196Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchGenaiMemoryRequest](../../pkg/models/operations/patchgenaimemoryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchGenaiMemoryResponse](../../pkg/models/operations/patchgenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiOrganization

Updates the GenaiOrganization identified by `id`. Provide the unified GenaiOrganization fields to change in the request body.

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
            Description: unifiedgosdk.Pointer("Subito teneo arbor."),
            ID: unifiedgosdk.Pointer("2ad18556-06a5-45b8-bb0c-7638bd042119"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("yum yuck"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-08T22:49:24.774Z"),
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

Updates the GenaiTask identified by `id`. Provide the unified GenaiTask fields to change in the request body.

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
            CompletedAt: types.MustNewTimeFromString("2025-10-30T21:35:21.720Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("63125da0-78ae-417b-8706-0515324f042b"),
            Instructions: unifiedgosdk.Pointer("Angustus cubo cupio abundans admoveo sumptus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Amet contigo tristis deripio auxilium cur aeternus cogito.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect online monitor"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-06-11T20:56:51.402Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Tenus via tamisium denego teres adsidue cura crur. Coma defungo magnam delicate confero amaritudo varius."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/eek-save"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T08:03:38.859Z"),
            WebURL: unifiedgosdk.Pointer("https://monstrous-honesty.net"),
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

## RemoveGenaiMemory

Deletes the GenaiMemory identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiMemory" method="delete" path="/genai/{connection_id}/memory/{id}" -->
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

    res, err := s.Genai.RemoveGenaiMemory(ctx, operations.RemoveGenaiMemoryRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveGenaiMemoryRequest](../../pkg/models/operations/removegenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveGenaiMemoryResponse](../../pkg/models/operations/removegenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiOrganization

Deletes the GenaiOrganization identified by `id`.

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

Deletes the GenaiSkill identified by `id`.

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

Deletes the GenaiTask identified by `id`.

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

    res, err := s.Genai.UpdateGenaiAgent(ctx, operations.UpdateGenaiAgentRequest{
        GenaiAgent: shared.GenaiAgent{
            CreatedAt: types.MustNewTimeFromString("2020-12-18T02:23:33.723Z"),
            Description: unifiedgosdk.Pointer("Sperno despecto."),
            ID: unifiedgosdk.Pointer("3e36559a-d979-42c3-9b9e-648f1aa534e8"),
            Instructions: unifiedgosdk.Pointer("Addo vita ver auctus corrumpo substantia conscendo auctus. Astrum sulum aufero bis stella sumptus non."),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("testimonial suddenly"),
            UpdatedAt: types.MustNewTimeFromString("2025-11-14T03:39:19.700Z"),
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

## UpdateGenaiMemory

Updates the GenaiMemory identified by `id`. Provide the unified GenaiMemory fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiMemory" method="put" path="/genai/{connection_id}/memory/{id}" example="genai_memory" -->
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

    res, err := s.Genai.UpdateGenaiMemory(ctx, operations.UpdateGenaiMemoryRequest{
        GenaiMemory: shared.GenaiMemory{
            Content: unifiedgosdk.Pointer("Spiculum deserunt comitatus despecto. Aveho utroque accedo vinitor repellendus tam curis corona eos vobis."),
            CreatedAt: types.MustNewTimeFromString("2021-10-22T04:14:41.668Z"),
            ID: unifiedgosdk.Pointer("c3894b70-1722-48ab-9297-af7f4f60ad2b"),
            Metadata: []shared.GenaiMetadata{
                shared.GenaiMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateGenaiMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.GenaiMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6d1774f7-a8ad-462d-9ba0-ece333056a26"),
                    Namespace: unifiedgosdk.Pointer("memory"),
                    Slug: unifiedgosdk.Pointer("defleo"),
                    Value: unifiedgosdk.Pointer(shared.CreateGenaiMetadataValueStr(
                        "suffragium",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("tabella.md"),
            Type: shared.GenaiMemoryTypeMemory.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-06T17:51:09.196Z"),
            UserIdentifier: unifiedgosdk.Pointer("70613433-847d-4db6-a8ea-96a2d7513dad"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiMemory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateGenaiMemoryRequest](../../pkg/models/operations/updategenaimemoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateGenaiMemoryResponse](../../pkg/models/operations/updategenaimemoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiOrganization

Updates the GenaiOrganization identified by `id`. Provide the unified GenaiOrganization fields to change in the request body.

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
            Description: unifiedgosdk.Pointer("Subito teneo arbor."),
            ID: unifiedgosdk.Pointer("2ad18556-06a5-45b8-bb0c-7638bd042119"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("yum yuck"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-08T22:49:24.774Z"),
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

Updates the GenaiTask identified by `id`. Provide the unified GenaiTask fields to change in the request body.

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
            CompletedAt: types.MustNewTimeFromString("2025-10-30T21:35:21.720Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("63125da0-78ae-417b-8706-0515324f042b"),
            Instructions: unifiedgosdk.Pointer("Angustus cubo cupio abundans admoveo sumptus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Amet contigo tristis deripio auxilium cur aeternus cogito.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect online monitor"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/felicity.rowe61/eek-save"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-06-11T20:56:51.402Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Tenus via tamisium denego teres adsidue cura crur. Coma defungo magnam delicate confero amaritudo varius."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/eek-save"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T08:03:38.859Z"),
            WebURL: unifiedgosdk.Pointer("https://monstrous-honesty.net"),
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