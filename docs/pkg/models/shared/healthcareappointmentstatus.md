# HealthcareAppointmentStatus

Status of the appointment

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareAppointmentStatusProposed

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareAppointmentStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `HealthcareAppointmentStatusProposed`       | PROPOSED                                    |
| `HealthcareAppointmentStatusPending`        | PENDING                                     |
| `HealthcareAppointmentStatusBooked`         | BOOKED                                      |
| `HealthcareAppointmentStatusArrived`        | ARRIVED                                     |
| `HealthcareAppointmentStatusCheckedIn`      | CHECKED_IN                                  |
| `HealthcareAppointmentStatusFulfilled`      | FULFILLED                                   |
| `HealthcareAppointmentStatusCancelled`      | CANCELLED                                   |
| `HealthcareAppointmentStatusNoshow`         | NOSHOW                                      |
| `HealthcareAppointmentStatusWaitlist`       | WAITLIST                                    |
| `HealthcareAppointmentStatusEnteredInError` | ENTERED_IN_ERROR                            |
| `HealthcareAppointmentStatusOther`          | OTHER                                       |