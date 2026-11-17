# Form

withholding form the values come from (eg. 2020+ W-4)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FormW42020

// Open enum: custom values can be created with a direct type cast
custom := shared.Form("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `FormW42020`    | W4_2020         |
| `FormW4Legacy`  | W4_LEGACY       |
| `FormStateForm` | STATE_FORM      |
| `FormTd1`       | TD1             |
| `FormOther`     | OTHER           |