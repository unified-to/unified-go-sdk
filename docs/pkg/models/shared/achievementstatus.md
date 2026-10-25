# AchievementStatus

Progress towards the goal

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AchievementStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := shared.AchievementStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `AchievementStatusInProgress`    | IN_PROGRESS                      |
| `AchievementStatusImproving`     | IMPROVING                        |
| `AchievementStatusWorsening`     | WORSENING                        |
| `AchievementStatusNoChange`      | NO_CHANGE                        |
| `AchievementStatusAchieved`      | ACHIEVED                         |
| `AchievementStatusSustaining`    | SUSTAINING                       |
| `AchievementStatusNotAchieved`   | NOT_ACHIEVED                     |
| `AchievementStatusNoProgress`    | NO_PROGRESS                      |
| `AchievementStatusNotAttainable` | NOT_ATTAINABLE                   |
| `AchievementStatusOther`         | OTHER                            |