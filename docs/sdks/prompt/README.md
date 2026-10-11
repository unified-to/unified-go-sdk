# Prompt

## Overview

### Available Operations

* [CreateGenaiPrompt](#creategenaiprompt) - Create a prompt

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

    res, err := s.Prompt.CreateGenaiPrompt(ctx, operations.CreateGenaiPromptRequest{
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