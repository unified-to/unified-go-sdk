# Prompt

## Overview

### Available Operations

* [CreateGenaiPrompt](#creategenaiprompt) - Create a prompt

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

    res, err := s.Prompt.CreateGenaiPrompt(ctx, operations.CreateGenaiPromptRequest{
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