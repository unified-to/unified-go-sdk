# Relationship

Relationship of the patient to the policy holder

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.RelationshipSelf

// Open enum: custom values can be created with a direct type cast
custom := shared.Relationship("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `RelationshipSelf`   | SELF                 |
| `RelationshipSpouse` | SPOUSE               |
| `RelationshipChild`  | CHILD                |
| `RelationshipOther`  | OTHER                |