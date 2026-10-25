# HealthcareDocumentType

Unified, normalized kind of clinical document (OTHER when unmappable)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareDocumentTypeProgressNote

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareDocumentType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `HealthcareDocumentTypeProgressNote`       | PROGRESS_NOTE                              |
| `HealthcareDocumentTypeConsultationNote`   | CONSULTATION_NOTE                          |
| `HealthcareDocumentTypeHistoryAndPhysical` | HISTORY_AND_PHYSICAL                       |
| `HealthcareDocumentTypeDischargeSummary`   | DISCHARGE_SUMMARY                          |
| `HealthcareDocumentTypeProcedureNote`      | PROCEDURE_NOTE                             |
| `HealthcareDocumentTypeReferralNote`       | REFERRAL_NOTE                              |
| `HealthcareDocumentTypeSummary`            | SUMMARY                                    |
| `HealthcareDocumentTypeLaboratoryReport`   | LABORATORY_REPORT                          |
| `HealthcareDocumentTypeImagingReport`      | IMAGING_REPORT                             |
| `HealthcareDocumentTypeConsent`            | CONSENT                                    |
| `HealthcareDocumentTypeOther`              | OTHER                                      |