# HealthcareReportStatus

Status of the report

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareReportStatusRegistered

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareReportStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `HealthcareReportStatusRegistered`     | REGISTERED                             |
| `HealthcareReportStatusPartial`        | PARTIAL                                |
| `HealthcareReportStatusPreliminary`    | PRELIMINARY                            |
| `HealthcareReportStatusFinal`          | FINAL                                  |
| `HealthcareReportStatusAmended`        | AMENDED                                |
| `HealthcareReportStatusCorrected`      | CORRECTED                              |
| `HealthcareReportStatusAppended`       | APPENDED                               |
| `HealthcareReportStatusCancelled`      | CANCELLED                              |
| `HealthcareReportStatusEnteredInError` | ENTERED_IN_ERROR                       |
| `HealthcareReportStatusUnknown`        | UNKNOWN                                |
| `HealthcareReportStatusOther`          | OTHER                                  |