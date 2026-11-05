# CalendarBotStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CalendarBotStatusReady

// Open enum: custom values can be created with a direct type cast
custom := shared.CalendarBotStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CalendarBotStatusReady`                      | READY                                         |
| `CalendarBotStatusJoiningCall`                | JOINING_CALL                                  |
| `CalendarBotStatusInWaitingRoom`              | IN_WAITING_ROOM                               |
| `CalendarBotStatusInCallNotRecording`         | IN_CALL_NOT_RECORDING                         |
| `CalendarBotStatusRecordingPermissionAllowed` | RECORDING_PERMISSION_ALLOWED                  |
| `CalendarBotStatusRecordingPermissionDenied`  | RECORDING_PERMISSION_DENIED                   |
| `CalendarBotStatusInCallRecording`            | IN_CALL_RECORDING                             |
| `CalendarBotStatusCallEnded`                  | CALL_ENDED                                    |
| `CalendarBotStatusDone`                       | DONE                                          |
| `CalendarBotStatusFatal`                      | FATAL                                         |