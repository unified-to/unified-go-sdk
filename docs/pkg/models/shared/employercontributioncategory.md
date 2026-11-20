# EmployerContributionCategory

kind of employer retirement contribution (eg. MATCH, SAFE_HARBOR_MATCH, PROFIT_SHARING)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EmployerContributionCategoryMatch

// Open enum: custom values can be created with a direct type cast
custom := shared.EmployerContributionCategory("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `EmployerContributionCategoryMatch`                 | MATCH                                               |
| `EmployerContributionCategorySafeHarborMatch`       | SAFE_HARBOR_MATCH                                   |
| `EmployerContributionCategorySafeHarborNonelective` | SAFE_HARBOR_NONELECTIVE                             |
| `EmployerContributionCategoryQacaMatch`             | QACA_MATCH                                          |
| `EmployerContributionCategoryNonelective`           | NONELECTIVE                                         |
| `EmployerContributionCategoryProfitSharing`         | PROFIT_SHARING                                      |
| `EmployerContributionCategoryOther`                 | OTHER                                               |