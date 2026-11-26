# Message

## Overview

### Available Operations

* [CreateMessagingMessage](#createmessagingmessage) - Create a message
* [GetMessagingMessage](#getmessagingmessage) - Retrieve a message
* [ListMessagingMessages](#listmessagingmessages) - List all messages
* [PatchMessagingMessage](#patchmessagingmessage) - Update a message
* [RemoveMessagingMessage](#removemessagingmessage) - Remove a message
* [UpdateMessagingMessage](#updatemessagingmessage) - Update a message

## CreateMessagingMessage

Creates a MessagingMessage. Provide the unified MessagingMessage fields in the request body; they are mapped to the provider's API. Returns the created MessagingMessage.

### Example Usage

<!-- UsageSnippet language="go" operationID="createMessagingMessage" method="post" path="/messaging/{connection_id}/message" example="messaging_message" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Message.CreateMessagingMessage(ctx, operations.CreateMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("e6dee3be-c95f-4420-9998-59dc1d498b08"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("ccc8e064-67ed-40b1-bf67-5b62146202e6"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "e69b1a7d-c34e-465f-b35e-47f7d6543734",
                    Text: unifiedgosdk.Pointer("comparo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("ccc8e064-67ed-40b1-bf67-5b62146202e6"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Metadata: []shared.MessagingMetadata{},
            Reactions: []shared.MessagingReaction{},
            Subject: unifiedgosdk.Pointer("Summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-29T12:16:51.448Z"),
            WebURL: unifiedgosdk.Pointer("https://grumpy-kit.net"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingMessage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateMessagingMessageRequest](../../pkg/models/operations/createmessagingmessagerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateMessagingMessageResponse](../../pkg/models/operations/createmessagingmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMessagingMessage

Retrieves a single MessagingMessage by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMessagingMessage" method="get" path="/messaging/{connection_id}/message/{id}" -->
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

    res, err := s.Message.GetMessagingMessage(ctx, operations.GetMessagingMessageRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingMessage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetMessagingMessageRequest](../../pkg/models/operations/getmessagingmessagerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetMessagingMessageResponse](../../pkg/models/operations/getmessagingmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMessagingMessages

Returns a list of MessagingMessages. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMessagingMessages" method="get" path="/messaging/{connection_id}/message" -->
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

    res, err := s.Message.ListMessagingMessages(ctx, operations.ListMessagingMessagesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListMessagingMessagesRequest](../../pkg/models/operations/listmessagingmessagesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListMessagingMessagesResponse](../../pkg/models/operations/listmessagingmessagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMessagingMessage

Updates the MessagingMessage identified by `id`. Provide the unified MessagingMessage fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMessagingMessage" method="patch" path="/messaging/{connection_id}/message/{id}" example="messaging_message" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Message.PatchMessagingMessage(ctx, operations.PatchMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("192b9469-be67-4bf9-a511-53c478ff22d8"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("55f15382-9fb0-470c-988f-3e032c0deb4a"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "18bbaf51-6a05-4d7b-9ec1-128b1c9cb1b4",
                    Text: unifiedgosdk.Pointer("comparo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("55f15382-9fb0-470c-988f-3e032c0deb4a"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Metadata: []shared.MessagingMetadata{},
            Reactions: []shared.MessagingReaction{},
            Subject: unifiedgosdk.Pointer("Summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-29T12:16:51.455Z"),
            WebURL: unifiedgosdk.Pointer("https://grumpy-kit.net"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingMessage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchMessagingMessageRequest](../../pkg/models/operations/patchmessagingmessagerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchMessagingMessageResponse](../../pkg/models/operations/patchmessagingmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMessagingMessage

Deletes the MessagingMessage identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMessagingMessage" method="delete" path="/messaging/{connection_id}/message/{id}" -->
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

    res, err := s.Message.RemoveMessagingMessage(ctx, operations.RemoveMessagingMessageRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveMessagingMessageRequest](../../pkg/models/operations/removemessagingmessagerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveMessagingMessageResponse](../../pkg/models/operations/removemessagingmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMessagingMessage

Updates the MessagingMessage identified by `id`. Provide the unified MessagingMessage fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMessagingMessage" method="put" path="/messaging/{connection_id}/message/{id}" example="messaging_message" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Message.UpdateMessagingMessage(ctx, operations.UpdateMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("192b9469-be67-4bf9-a511-53c478ff22d8"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("55f15382-9fb0-470c-988f-3e032c0deb4a"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "18bbaf51-6a05-4d7b-9ec1-128b1c9cb1b4",
                    Text: unifiedgosdk.Pointer("comparo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("55f15382-9fb0-470c-988f-3e032c0deb4a"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Metadata: []shared.MessagingMetadata{},
            Reactions: []shared.MessagingReaction{},
            Subject: unifiedgosdk.Pointer("Summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-29T12:16:51.455Z"),
            WebURL: unifiedgosdk.Pointer("https://grumpy-kit.net"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingMessage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateMessagingMessageRequest](../../pkg/models/operations/updatemessagingmessagerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateMessagingMessageResponse](../../pkg/models/operations/updatemessagingmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |