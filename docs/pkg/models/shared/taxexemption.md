# TaxExemption

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TaxExemptionFederalGov

// Open enum: custom values can be created with a direct type cast
custom := shared.TaxExemption("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `TaxExemptionFederalGov`     | FEDERAL_GOV                  |
| `TaxExemptionRegionGov`      | REGION_GOV                   |
| `TaxExemptionLocalGov`       | LOCAL_GOV                    |
| `TaxExemptionTribalGov`      | TRIBAL_GOV                   |
| `TaxExemptionCharitableOrg`  | CHARITABLE_ORG               |
| `TaxExemptionReligiousOrg`   | RELIGIOUS_ORG                |
| `TaxExemptionEducationalOrg` | EDUCATIONAL_ORG              |
| `TaxExemptionMedicalOrg`     | MEDICAL_ORG                  |
| `TaxExemptionResale`         | RESALE                       |
| `TaxExemptionForeign`        | FOREIGN                      |
| `TaxExemptionOther`          | OTHER                        |