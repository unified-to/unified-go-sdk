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
                    ContentIdentifier: unifiedgosdk.Pointer("4ad3db36-7546-429c-acba-15c5eb1696c7"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("c1e9b8e5-7823-4034-b85b-2c908cf14d76"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "76272918-9117-4669-bb53-8be6fd85cd2a",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("c1e9b8e5-7823-4034-b85b-2c908cf14d76"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-07T09:39:46.353Z"),
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
                    ContentIdentifier: unifiedgosdk.Pointer("23d09ba2-b788-4f49-b161-2ff0d6c1a52b"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("e8b9a8f4-0b79-42b1-b70b-4a7b06421309"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "a5423149-9552-40da-9ce9-8b72eaf8d1e7",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("e8b9a8f4-0b79-42b1-b70b-4a7b06421309"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-07T09:39:46.366Z"),
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
                    ContentIdentifier: unifiedgosdk.Pointer("23d09ba2-b788-4f49-b161-2ff0d6c1a52b"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("e8b9a8f4-0b79-42b1-b70b-4a7b06421309"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "a5423149-9552-40da-9ce9-8b72eaf8d1e7",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("e8b9a8f4-0b79-42b1-b70b-4a7b06421309"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-07T09:39:46.366Z"),
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