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

Create a message

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
                    ContentIdentifier: unifiedgosdk.Pointer("ef3f64ac-140c-41e3-94c9-1c6604f60a5f"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("e13c69ad-a43f-4b9c-a7cd-0631dacd0498"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "8d6c57d3-f091-4711-8d8e-e59039084998",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("e13c69ad-a43f-4b9c-a7cd-0631dacd0498"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T17:08:23.000Z"),
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

Retrieve a message

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

List all messages

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

Update a message

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
                    ContentIdentifier: unifiedgosdk.Pointer("a7457040-f645-4e83-92c4-ce4e57fd6015"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("6288c40e-281e-4f58-9dcb-d0a9757b1347"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "2a2499ea-2cc7-4609-89fa-1c34c2c8f841",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("6288c40e-281e-4f58-9dcb-d0a9757b1347"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T17:08:23.015Z"),
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

Remove a message

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

Update a message

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
                    ContentIdentifier: unifiedgosdk.Pointer("a7457040-f645-4e83-92c4-ce4e57fd6015"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("6288c40e-281e-4f58-9dcb-d0a9757b1347"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "2a2499ea-2cc7-4609-89fa-1c34c2c8f841",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("6288c40e-281e-4f58-9dcb-d0a9757b1347"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T17:08:23.015Z"),
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