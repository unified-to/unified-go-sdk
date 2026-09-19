# Genai

## Overview

### Available Operations

* [CreateGenaiEmbedding](#creategenaiembedding) - Create an embedding
* [CreateGenaiPrompt](#creategenaiprompt) - Create a prompt
* [GetGenaiModel](#getgenaimodel) - Retrieve a model
* [ListGenaiModels](#listgenaimodels) - List all models

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
            ID: unifiedgosdk.Pointer("f71681e0-f4f7-4689-af6f-68ff10d75b79"),
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
            MaxTokens: unifiedgosdk.Pointer[float64](0.4677782787475735),
            McpAuthorizationToken: unifiedgosdk.Pointer("f45a6e93-7bed-49b4-a5c8-37a2ed2d58f4"),
            McpDeferredTools: []string{},
            McpURL: unifiedgosdk.Pointer("https://unsung-dusk.info/"),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Aegre repudiandae verecundia facere statua.",
                    Role: shared.RoleAssistant.ToPointer(),
                },
                shared.GenaiContent{
                    Content: "Speciosus xiphias soleo trepide crinis.",
                    Role: shared.RoleSystem.ToPointer(),
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