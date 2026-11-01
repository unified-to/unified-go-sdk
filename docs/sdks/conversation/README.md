# Conversation

## Overview

### Available Operations

* [CreatePropertyConversation](#createpropertyconversation) - Create a conversation
* [GetPropertyConversation](#getpropertyconversation) - Retrieve a conversation
* [ListPropertyConversations](#listpropertyconversations) - List all conversations

## CreatePropertyConversation

Creates a PropertyConversation. Provide the unified PropertyConversation fields in the request body; they are mapped to the provider's API. Returns the created PropertyConversation.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyConversation" method="post" path="/property/{connection_id}/conversation" example="property_conversation" -->
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

    res, err := s.Conversation.CreatePropertyConversation(ctx, operations.CreatePropertyConversationRequest{
        PropertyConversation: shared.PropertyConversation{
            CreatedAt: types.MustNewTimeFromString("2022-06-20T00:02:46.530Z"),
            GuestName: unifiedgosdk.Pointer("Evie Okuneva"),
            ID: unifiedgosdk.Pointer("b8d35ad4-3814-428a-8a30-57cc15b7ed19"),
            IsUnread: unifiedgosdk.Pointer(false),
            Messages: []shared.PropertyMessage{
                shared.PropertyMessage{
                    Attachments: []shared.PropertyAttachment{},
                    CreatedAt: types.MustNewTimeFromString("2025-04-23T16:38:22.920Z"),
                    ID: unifiedgosdk.Pointer("c4d9d075-9ee0-4965-bc8c-18f94da835c7"),
                    IsAutomatic: unifiedgosdk.Pointer(false),
                    Message: unifiedgosdk.Pointer("Vinitor caterva amor sophismata admitto colo tempore aufero torrens."),
                    SenderName: unifiedgosdk.Pointer("Evie Okuneva"),
                    SenderType: shared.SenderTypeGuest.ToPointer(),
                },
                shared.PropertyMessage{
                    Attachments: []shared.PropertyAttachment{},
                    CreatedAt: types.MustNewTimeFromString("2026-05-07T10:20:08.159Z"),
                    ID: unifiedgosdk.Pointer("39fdae71-4706-4e7f-b2c7-d26a271df258"),
                    Message: unifiedgosdk.Pointer("Aequitas torqueo curiositas cimentarius repellendus auditor degero ademptio adnuo."),
                    SenderName: unifiedgosdk.Pointer("Ms. Patti Lynch"),
                    SenderType: shared.SenderTypeHost.ToPointer(),
                },
            },
            Metadata: []shared.PropertyMetadata{},
            Source: unifiedgosdk.Pointer("Airbnb"),
            UpdatedAt: types.MustNewTimeFromString("2026-05-07T10:20:08.159Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreatePropertyConversationRequest](../../pkg/models/operations/createpropertyconversationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreatePropertyConversationResponse](../../pkg/models/operations/createpropertyconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPropertyConversation

Retrieves a single PropertyConversation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPropertyConversation" method="get" path="/property/{connection_id}/conversation/{id}" -->
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

    res, err := s.Conversation.GetPropertyConversation(ctx, operations.GetPropertyConversationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetPropertyConversationRequest](../../pkg/models/operations/getpropertyconversationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetPropertyConversationResponse](../../pkg/models/operations/getpropertyconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPropertyConversations

Returns a list of PropertyConversations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPropertyConversations" method="get" path="/property/{connection_id}/conversation" -->
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

    res, err := s.Conversation.ListPropertyConversations(ctx, operations.ListPropertyConversationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyConversations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListPropertyConversationsRequest](../../pkg/models/operations/listpropertyconversationsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListPropertyConversationsResponse](../../pkg/models/operations/listpropertyconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |