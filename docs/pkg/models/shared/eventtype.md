# EventType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EventTypePageView

// Open enum: custom values can be created with a direct type cast
custom := shared.EventType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `EventTypePageView`      | PAGE_VIEW                |
| `EventTypeScreenView`    | SCREEN_VIEW              |
| `EventTypeClick`         | CLICK                    |
| `EventTypeFormSubmit`    | FORM_SUBMIT              |
| `EventTypePurchase`      | PURCHASE                 |
| `EventTypeSignUp`        | SIGN_UP                  |
| `EventTypeLogin`         | LOGIN                    |
| `EventTypeLogout`        | LOGOUT                   |
| `EventTypeSearch`        | SEARCH                   |
| `EventTypeVideoPlay`     | VIDEO_PLAY               |
| `EventTypeVideoComplete` | VIDEO_COMPLETE           |
| `EventTypeFileDownload`  | FILE_DOWNLOAD            |
| `EventTypeScroll`        | SCROLL                   |
| `EventTypeSessionStart`  | SESSION_START            |
| `EventTypeFirstVisit`    | FIRST_VISIT              |
| `EventTypeCustom`        | CUSTOM                   |