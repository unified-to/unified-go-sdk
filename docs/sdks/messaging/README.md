# Messaging

## Overview

### Available Operations

* [CreateMessagingChannel](#createmessagingchannel) - Create a channel
* [CreateMessagingMessage](#createmessagingmessage) - Create a message
* [GetMessagingChannel](#getmessagingchannel) - Retrieve a channel
* [GetMessagingMessage](#getmessagingmessage) - Retrieve a message
* [ListMessagingChannels](#listmessagingchannels) - List all channels
* [ListMessagingMessages](#listmessagingmessages) - List all messages
* [PatchMessagingChannel](#patchmessagingchannel) - Update a channel
* [PatchMessagingEvent](#patchmessagingevent) - Update an event
* [PatchMessagingMessage](#patchmessagingmessage) - Update a message
* [RemoveMessagingChannel](#removemessagingchannel) - Remove a channel
* [RemoveMessagingMessage](#removemessagingmessage) - Remove a message
* [UpdateMessagingChannel](#updatemessagingchannel) - Update a channel
* [UpdateMessagingEvent](#updatemessagingevent) - Update an event
* [UpdateMessagingMessage](#updatemessagingmessage) - Update a message

## CreateMessagingChannel

Create a channel

### Example Usage

<!-- UsageSnippet language="go" operationID="createMessagingChannel" method="post" path="/messaging/{connection_id}/channel" example="messaging_channel" -->
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

    res, err := s.Messaging.CreateMessagingChannel(ctx, operations.CreateMessagingChannelRequest{
        MessagingChannel: shared.MessagingChannel{
            CreatedAt: types.MustNewTimeFromString("2023-10-05T02:09:22.795Z"),
            Description: unifiedgosdk.Pointer("Dolores tutis."),
            HasSubchannels: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("2c35d6e5-7053-4dba-8f8d-8f4a74fd60d9"),
            IsActive: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(true),
            Members: []shared.MessagingMember{},
            Name: unifiedgosdk.Pointer("tego"),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T22:52:13.426Z"),
            WebURL: unifiedgosdk.Pointer("https://svelte-rule.name/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateMessagingChannelRequest](../../pkg/models/operations/createmessagingchannelrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateMessagingChannelResponse](../../pkg/models/operations/createmessagingchannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Messaging.CreateMessagingMessage(ctx, operations.CreateMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("d9c0649c-94e5-4618-a748-5c8bb89e96cd"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("5eaed8e7-f457-484b-bcaa-548d7bacfb7b"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "00154f6c-7a48-4768-bde1-dcd68f3b9f80",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("5eaed8e7-f457-484b-bcaa-548d7bacfb7b"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T03:45:47.465Z"),
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

## GetMessagingChannel

Retrieve a channel

### Example Usage

<!-- UsageSnippet language="go" operationID="getMessagingChannel" method="get" path="/messaging/{connection_id}/channel/{id}" -->
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

    res, err := s.Messaging.GetMessagingChannel(ctx, operations.GetMessagingChannelRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetMessagingChannelRequest](../../pkg/models/operations/getmessagingchannelrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetMessagingChannelResponse](../../pkg/models/operations/getmessagingchannelresponse.md), error**

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

    res, err := s.Messaging.GetMessagingMessage(ctx, operations.GetMessagingMessageRequest{
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

## ListMessagingChannels

List all channels

### Example Usage

<!-- UsageSnippet language="go" operationID="listMessagingChannels" method="get" path="/messaging/{connection_id}/channel" -->
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

    res, err := s.Messaging.ListMessagingChannels(ctx, operations.ListMessagingChannelsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingChannels != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListMessagingChannelsRequest](../../pkg/models/operations/listmessagingchannelsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListMessagingChannelsResponse](../../pkg/models/operations/listmessagingchannelsresponse.md), error**

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

    res, err := s.Messaging.ListMessagingMessages(ctx, operations.ListMessagingMessagesRequest{
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

## PatchMessagingChannel

Update a channel

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMessagingChannel" method="patch" path="/messaging/{connection_id}/channel/{id}" example="messaging_channel" -->
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

    res, err := s.Messaging.PatchMessagingChannel(ctx, operations.PatchMessagingChannelRequest{
        MessagingChannel: shared.MessagingChannel{
            CreatedAt: types.MustNewTimeFromString("2023-10-05T02:09:22.795Z"),
            Description: unifiedgosdk.Pointer("Dolores tutis."),
            HasSubchannels: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("350b99e1-1ae1-4558-ab58-7ea39f375699"),
            IsActive: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(true),
            Members: []shared.MessagingMember{},
            Name: unifiedgosdk.Pointer("tego"),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T22:52:13.437Z"),
            WebURL: unifiedgosdk.Pointer("https://svelte-rule.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchMessagingChannelRequest](../../pkg/models/operations/patchmessagingchannelrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchMessagingChannelResponse](../../pkg/models/operations/patchmessagingchannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMessagingEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMessagingEvent" method="patch" path="/messaging/{connection_id}/event/{id}" example="messaging_event" -->
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

    res, err := s.Messaging.PatchMessagingEvent(ctx, operations.PatchMessagingEventRequest{
        MessagingEvent: shared.MessagingEvent{
            Channel: &shared.PropertyMessagingEventChannel{
                ID: unifiedgosdk.Pointer(""),
                Name: unifiedgosdk.Pointer(""),
            },
            CreatedAt: types.MustNewTimeFromString("2019-05-30T19:44:46.461Z"),
            ID: unifiedgosdk.Pointer("5553590c-968c-4bb1-85df-0fc2d583a312"),
            IsReplacingOriginal: unifiedgosdk.Pointer(false),
            Type: shared.MessagingEventTypeButtonClick.ToPointer(),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchMessagingEventRequest](../../pkg/models/operations/patchmessagingeventrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchMessagingEventResponse](../../pkg/models/operations/patchmessagingeventresponse.md), error**

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

    res, err := s.Messaging.PatchMessagingMessage(ctx, operations.PatchMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("8eb8e17e-bcfd-4a0f-90f8-5aaf7ff9985c"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("b6384027-5900-4d19-9acf-e41631db16a0"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "2d29b051-25a1-4051-bcd6-d5377ca8ed24",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("b6384027-5900-4d19-9acf-e41631db16a0"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T03:45:47.481Z"),
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

## RemoveMessagingChannel

Remove a channel

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMessagingChannel" method="delete" path="/messaging/{connection_id}/channel/{id}" -->
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

    res, err := s.Messaging.RemoveMessagingChannel(ctx, operations.RemoveMessagingChannelRequest{
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
| `request`                                                                                                | [operations.RemoveMessagingChannelRequest](../../pkg/models/operations/removemessagingchannelrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveMessagingChannelResponse](../../pkg/models/operations/removemessagingchannelresponse.md), error**

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

    res, err := s.Messaging.RemoveMessagingMessage(ctx, operations.RemoveMessagingMessageRequest{
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

## UpdateMessagingChannel

Update a channel

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMessagingChannel" method="put" path="/messaging/{connection_id}/channel/{id}" example="messaging_channel" -->
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

    res, err := s.Messaging.UpdateMessagingChannel(ctx, operations.UpdateMessagingChannelRequest{
        MessagingChannel: shared.MessagingChannel{
            CreatedAt: types.MustNewTimeFromString("2023-10-05T02:09:22.795Z"),
            Description: unifiedgosdk.Pointer("Dolores tutis."),
            HasSubchannels: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("350b99e1-1ae1-4558-ab58-7ea39f375699"),
            IsActive: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(true),
            Members: []shared.MessagingMember{},
            Name: unifiedgosdk.Pointer("tego"),
            UpdatedAt: types.MustNewTimeFromString("2026-05-12T22:52:13.437Z"),
            WebURL: unifiedgosdk.Pointer("https://svelte-rule.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateMessagingChannelRequest](../../pkg/models/operations/updatemessagingchannelrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateMessagingChannelResponse](../../pkg/models/operations/updatemessagingchannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMessagingEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMessagingEvent" method="put" path="/messaging/{connection_id}/event/{id}" example="messaging_event" -->
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

    res, err := s.Messaging.UpdateMessagingEvent(ctx, operations.UpdateMessagingEventRequest{
        MessagingEvent: shared.MessagingEvent{
            Channel: &shared.PropertyMessagingEventChannel{
                ID: unifiedgosdk.Pointer(""),
                Name: unifiedgosdk.Pointer(""),
            },
            CreatedAt: types.MustNewTimeFromString("2019-05-30T19:44:46.461Z"),
            ID: unifiedgosdk.Pointer("5553590c-968c-4bb1-85df-0fc2d583a312"),
            IsReplacingOriginal: unifiedgosdk.Pointer(false),
            Type: shared.MessagingEventTypeButtonClick.ToPointer(),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagingEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateMessagingEventRequest](../../pkg/models/operations/updatemessagingeventrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateMessagingEventResponse](../../pkg/models/operations/updatemessagingeventresponse.md), error**

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

    res, err := s.Messaging.UpdateMessagingMessage(ctx, operations.UpdateMessagingMessageRequest{
        MessagingMessage: shared.MessagingMessage{
            Attachments: []shared.MessagingAttachment{
                shared.MessagingAttachment{
                    ContentIdentifier: unifiedgosdk.Pointer("8eb8e17e-bcfd-4a0f-90f8-5aaf7ff9985c"),
                    ContentType: unifiedgosdk.Pointer("coaegresco"),
                    DownloadURL: unifiedgosdk.Pointer("https://rotating-advertisement.org"),
                    Filename: unifiedgosdk.Pointer("super"),
                    MessageID: unifiedgosdk.Pointer("b6384027-5900-4d19-9acf-e41631db16a0"),
                    Size: unifiedgosdk.Pointer[float64](327.0),
                },
            },
            Buttons: []shared.MessagingButton{
                shared.MessagingButton{
                    ID: "2d29b051-25a1-4051-bcd6-d5377ca8ed24",
                    Text: unifiedgosdk.Pointer("denuo"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-11-26T09:26:33.973Z"),
            DestinationMembers: []shared.MessagingMember{},
            HasChildren: unifiedgosdk.Pointer(true),
            HiddenMembers: []shared.MessagingMember{},
            ID: unifiedgosdk.Pointer("b6384027-5900-4d19-9acf-e41631db16a0"),
            IsUnread: unifiedgosdk.Pointer(false),
            MentionedMembers: []shared.MessagingMember{},
            Message: unifiedgosdk.Pointer("Sum utique aliquid."),
            MessageHTML: unifiedgosdk.Pointer("Articulus tardus tergiversatio."),
            MessageMarkdown: unifiedgosdk.Pointer("Territo uterque tergo curiositas."),
            Reactions: []shared.MessagingReaction{},
            Reference: unifiedgosdk.Pointer("571483f2-d95b-4f06-8b78-d35e7046bb74"),
            Subject: unifiedgosdk.Pointer("Cernuus optio cohaero summisse in."),
            UpdatedAt: types.MustNewTimeFromString("2023-07-14T03:45:47.481Z"),
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