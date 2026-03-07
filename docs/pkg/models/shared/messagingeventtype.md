# MessagingEventType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.MessagingEventTypeMessageReceived

// Open enum: custom values can be created with a direct type cast
custom := shared.MessagingEventType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `MessagingEventTypeMessageReceived` | MESSAGE_RECEIVED                    |
| `MessagingEventTypeReactionAdded`   | REACTION_ADDED                      |
| `MessagingEventTypeReactionRemoved` | REACTION_REMOVED                    |
| `MessagingEventTypeButtonClick`     | BUTTON_CLICK                        |
| `MessagingEventTypeAppMention`      | APP_MENTION                         |
| `MessagingEventTypeChannelJoined`   | CHANNEL_JOINED                      |
| `MessagingEventTypeChannelLeft`     | CHANNEL_LEFT                        |
| `MessagingEventTypeChannelCreated`  | CHANNEL_CREATED                     |
| `MessagingEventTypeChannelDeleted`  | CHANNEL_DELETED                     |
| `MessagingEventTypeUserCreated`     | USER_CREATED                        |
| `MessagingEventTypeUserDeleted`     | USER_DELETED                        |
| `MessagingEventTypeUserUpdated`     | USER_UPDATED                        |