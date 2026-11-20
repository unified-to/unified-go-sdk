# RetirementPlanType

the retirement plan this benefit belongs to (for RETIREMENT* types)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.RetirementPlanTypePlan401K

// Open enum: custom values can be created with a direct type cast
custom := shared.RetirementPlanType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RetirementPlanTypePlan401K`  | PLAN_401K                     |
| `RetirementPlanTypePlan403B`  | PLAN_403B                     |
| `RetirementPlanTypePlan457B`  | PLAN_457B                     |
| `RetirementPlanTypeSimpleIra` | SIMPLE_IRA                    |
| `RetirementPlanTypeSepIra`    | SEP_IRA                       |
| `RetirementPlanTypeNqdc`      | NQDC                          |
| `RetirementPlanTypePension`   | PENSION                       |
| `RetirementPlanTypeOther`     | OTHER                         |