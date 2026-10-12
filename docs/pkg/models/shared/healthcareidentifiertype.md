# HealthcareIdentifierType

The kind of identifier

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareIdentifierTypeMrn

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareIdentifierType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HealthcareIdentifierTypeMrn`            | MRN                                      |
| `HealthcareIdentifierTypeNpi`            | NPI                                      |
| `HealthcareIdentifierTypeSsn`            | SSN                                      |
| `HealthcareIdentifierTypeDriversLicense` | DRIVERS_LICENSE                          |
| `HealthcareIdentifierTypePassport`       | PASSPORT                                 |
| `HealthcareIdentifierTypeInsurance`      | INSURANCE                                |
| `HealthcareIdentifierTypeOther`          | OTHER                                    |