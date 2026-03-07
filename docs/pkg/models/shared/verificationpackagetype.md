# VerificationPackageType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.VerificationPackageTypeIdentityVerification

// Open enum: custom values can be created with a direct type cast
custom := shared.VerificationPackageType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `VerificationPackageTypeIdentityVerification`   | IDENTITY_VERIFICATION                           |
| `VerificationPackageTypeScreening`              | SCREENING                                       |
| `VerificationPackageTypeBackgroundCheck`        | BACKGROUND_CHECK                                |
| `VerificationPackageTypeEmploymentVerification` | EMPLOYMENT_VERIFICATION                         |
| `VerificationPackageTypeEducationVerification`  | EDUCATION_VERIFICATION                          |
| `VerificationPackageTypeCreditCheck`            | CREDIT_CHECK                                    |
| `VerificationPackageTypeFraudPrevention`        | FRAUD_PREVENTION                                |
| `VerificationPackageTypeOther`                  | OTHER                                           |