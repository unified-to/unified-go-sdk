# CalendarBotStatusChangeStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CalendarBotStatusChangeStatusReady

// Open enum: custom values can be created with a direct type cast
custom := shared.CalendarBotStatusChangeStatus("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CalendarBotStatusChangeStatusReady`                      | READY                                                     |
| `CalendarBotStatusChangeStatusJoiningCall`                | JOINING_CALL                                              |
| `CalendarBotStatusChangeStatusInWaitingRoom`              | IN_WAITING_ROOM                                           |
| `CalendarBotStatusChangeStatusInCallNotRecording`         | IN_CALL_NOT_RECORDING                                     |
| `CalendarBotStatusChangeStatusRecordingPermissionAllowed` | RECORDING_PERMISSION_ALLOWED                              |
| `CalendarBotStatusChangeStatusRecordingPermissionDenied`  | RECORDING_PERMISSION_DENIED                               |
| `CalendarBotStatusChangeStatusInCallRecording`            | IN_CALL_RECORDING                                         |
| `CalendarBotStatusChangeStatusCallEnded`                  | CALL_ENDED                                                |
| `CalendarBotStatusChangeStatusDone`                       | DONE                                                      |
| `CalendarBotStatusChangeStatusFatal`                      | FATAL                                                     |