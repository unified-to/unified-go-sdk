# SocialPostStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SocialPostStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.SocialPostStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `SocialPostStatusDraft`      | DRAFT                        |
| `SocialPostStatusScheduled`  | SCHEDULED                    |
| `SocialPostStatusPublished`  | PUBLISHED                    |
| `SocialPostStatusRejected`   | REJECTED                     |
| `SocialPostStatusProcessing` | PROCESSING                   |
| `SocialPostStatusDeleted`    | DELETED                      |
| `SocialPostStatusOther`      | OTHER                        |