# Embedding

## Overview

### Available Operations

* [CreateGenaiEmbedding](#creategenaiembedding) - Create an embedding

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

    res, err := s.Embedding.CreateGenaiEmbedding(ctx, operations.CreateGenaiEmbeddingRequest{
        GenaiEmbedding: shared.GenaiEmbedding{
            Content: []shared.GenaiEmbeddingContent{
                shared.GenaiEmbeddingContent{
                    Text: "Demulceo adulescens est delectatio solitudo videlicet volo.",
                },
            },
            Dimension: unifiedgosdk.Pointer[float64](423.0),
            Embeddings: unifiedgosdk.Pointer("Venia tutis cernuus viscus officiis utilis porro utique."),
            EncodingFormat: shared.EncodingFormatFloat.ToPointer(),
            ID: unifiedgosdk.Pointer("11db6f2b-f8ae-4068-829e-6b1af0b40ea6"),
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