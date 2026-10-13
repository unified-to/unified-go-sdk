# Healthcare

## Overview

### Available Operations

* [CreateHealthcareAppointment](#createhealthcareappointment) - Create an appointment
* [CreateHealthcareObservation](#createhealthcareobservation) - Create an observation
* [CreateHealthcarePatient](#createhealthcarepatient) - Create a patient
* [GetHealthcareAllergy](#gethealthcareallergy) - Retrieve an allergy
* [GetHealthcareAppointment](#gethealthcareappointment) - Retrieve an appointment
* [GetHealthcareCondition](#gethealthcarecondition) - Retrieve a condition
* [GetHealthcareEncounter](#gethealthcareencounter) - Retrieve an encounter
* [GetHealthcareImmunization](#gethealthcareimmunization) - Retrieve an immunization
* [GetHealthcareLocation](#gethealthcarelocation) - Retrieve a location
* [GetHealthcareMedication](#gethealthcaremedication) - Retrieve a medication
* [GetHealthcareObservation](#gethealthcareobservation) - Retrieve an observation
* [GetHealthcarePatient](#gethealthcarepatient) - Retrieve a patient
* [GetHealthcarePractitioner](#gethealthcarepractitioner) - Retrieve a practitioner
* [GetHealthcareSlot](#gethealthcareslot) - Retrieve a slot
* [ListHealthcareAllergies](#listhealthcareallergies) - List all allergies
* [ListHealthcareAppointments](#listhealthcareappointments) - List all appointments
* [ListHealthcareConditions](#listhealthcareconditions) - List all conditions
* [ListHealthcareEncounters](#listhealthcareencounters) - List all encounters
* [ListHealthcareImmunizations](#listhealthcareimmunizations) - List all immunizations
* [ListHealthcareLocations](#listhealthcarelocations) - List all locations
* [ListHealthcareMedications](#listhealthcaremedications) - List all medications
* [ListHealthcareObservations](#listhealthcareobservations) - List all observations
* [ListHealthcarePatients](#listhealthcarepatients) - List all patients
* [ListHealthcarePractitioners](#listhealthcarepractitioners) - List all practitioners
* [ListHealthcareSlots](#listhealthcareslots) - List all slots
* [PatchHealthcareAppointment](#patchhealthcareappointment) - Update an appointment
* [PatchHealthcarePatient](#patchhealthcarepatient) - Update a patient
* [RemoveHealthcareAppointment](#removehealthcareappointment) - Remove an appointment
* [UpdateHealthcareAppointment](#updatehealthcareappointment) - Update an appointment
* [UpdateHealthcarePatient](#updatehealthcarepatient) - Update a patient

## CreateHealthcareAppointment

Create an appointment

### Example Usage

<!-- UsageSnippet language="go" operationID="createHealthcareAppointment" method="post" path="/healthcare/{connection_id}/appointment" example="healthcare_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.CreateHealthcareAppointment(ctx, operations.CreateHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("clarus aptus cetera"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("7c319059-dde7-4949-873e-c16ed53353b9"),
            Metadata: []shared.HealthcareMetadata{},
            PractitionerIds: []string{},
            Reasons: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("cubicularis caput"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("61128108"),
                    Display: unifiedgosdk.Pointer("Annual physical"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-13T06:51:07.920Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateHealthcareAppointmentRequest](../../pkg/models/operations/createhealthcareappointmentrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateHealthcareAppointmentResponse](../../pkg/models/operations/createhealthcareappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHealthcareObservation

Create an observation

### Example Usage

<!-- UsageSnippet language="go" operationID="createHealthcareObservation" method="post" path="/healthcare/{connection_id}/observation" example="healthcare_observation" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.CreateHealthcareObservation(ctx, operations.CreateHealthcareObservationRequest{
        HealthcareObservation: shared.HealthcareObservation{
            Codes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("85354-9"),
                    Display: unifiedgosdk.Pointer("Blood pressure panel"),
                    System: unifiedgosdk.Pointer("http://loinc.org"),
                },
            },
            Components: []shared.HealthcareObservationComponent{
                shared.HealthcareObservationComponent{
                    Codes: []shared.HealthcareCode{
                        shared.HealthcareCode{
                            Code: unifiedgosdk.Pointer("8480-6"),
                            System: unifiedgosdk.Pointer("http://loinc.org"),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Systolic blood pressure"),
                    ValueNumber: unifiedgosdk.Pointer[float64](103.0),
                    ValueUnit: unifiedgosdk.Pointer("mm[Hg]"),
                },
                shared.HealthcareObservationComponent{
                    Codes: []shared.HealthcareCode{
                        shared.HealthcareCode{
                            Code: unifiedgosdk.Pointer("8462-4"),
                            System: unifiedgosdk.Pointer("http://loinc.org"),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Diastolic blood pressure"),
                    ValueNumber: unifiedgosdk.Pointer[float64](60.0),
                    ValueUnit: unifiedgosdk.Pointer("mm[Hg]"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2022-12-06T21:34:04.293Z"),
            ID: unifiedgosdk.Pointer("318b9d52-873e-449c-adaf-8a5a4d160b0a"),
            Interpretation: shared.InterpretationCriticalHigh.ToPointer(),
            IssuedAt: types.MustNewTimeFromString("2025-05-06T05:10:44.706Z"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Blood pressure panel"),
            ObservedAt: types.MustNewTimeFromString("2025-01-15T14:25:13.808Z"),
            Status: shared.HealthcareObservationStatusCorrected.ToPointer(),
            Type: shared.HealthcareObservationTypeVitalSigns.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T04:18:52.114Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareObservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateHealthcareObservationRequest](../../pkg/models/operations/createhealthcareobservationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateHealthcareObservationResponse](../../pkg/models/operations/createhealthcareobservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHealthcarePatient

Create a patient

### Example Usage

<!-- UsageSnippet language="go" operationID="createHealthcarePatient" method="post" path="/healthcare/{connection_id}/patient" example="healthcare_patient" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.CreateHealthcarePatient(ctx, operations.CreateHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("3645 Zoila Isle"),
                Address2: unifiedgosdk.Pointer("Suite 613"),
                City: unifiedgosdk.Pointer("Jacintostead"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("31579"),
                Region: unifiedgosdk.Pointer("Minnesota"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-03"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@yahoo.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("1bfecddf-8427-41d0-8e5e-37ffcd4dca34"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "92636004",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("en"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "659.223.0829 x590",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-15T00:39:33.004Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePatient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateHealthcarePatientRequest](../../pkg/models/operations/createhealthcarepatientrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateHealthcarePatientResponse](../../pkg/models/operations/createhealthcarepatientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareAllergy

Retrieve an allergy

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareAllergy" method="get" path="/healthcare/{connection_id}/allergy/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareAllergy(ctx, operations.GetHealthcareAllergyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAllergy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetHealthcareAllergyRequest](../../pkg/models/operations/gethealthcareallergyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetHealthcareAllergyResponse](../../pkg/models/operations/gethealthcareallergyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareAppointment

Retrieve an appointment

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareAppointment" method="get" path="/healthcare/{connection_id}/appointment/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareAppointment(ctx, operations.GetHealthcareAppointmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetHealthcareAppointmentRequest](../../pkg/models/operations/gethealthcareappointmentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetHealthcareAppointmentResponse](../../pkg/models/operations/gethealthcareappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareCondition

Retrieve a condition

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareCondition" method="get" path="/healthcare/{connection_id}/condition/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareCondition(ctx, operations.GetHealthcareConditionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareCondition != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetHealthcareConditionRequest](../../pkg/models/operations/gethealthcareconditionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetHealthcareConditionResponse](../../pkg/models/operations/gethealthcareconditionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareEncounter

Retrieve an encounter

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareEncounter" method="get" path="/healthcare/{connection_id}/encounter/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareEncounter(ctx, operations.GetHealthcareEncounterRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareEncounter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetHealthcareEncounterRequest](../../pkg/models/operations/gethealthcareencounterrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetHealthcareEncounterResponse](../../pkg/models/operations/gethealthcareencounterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareImmunization

Retrieve an immunization

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareImmunization" method="get" path="/healthcare/{connection_id}/immunization/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareImmunization(ctx, operations.GetHealthcareImmunizationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareImmunization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetHealthcareImmunizationRequest](../../pkg/models/operations/gethealthcareimmunizationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetHealthcareImmunizationResponse](../../pkg/models/operations/gethealthcareimmunizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareLocation" method="get" path="/healthcare/{connection_id}/location/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareLocation(ctx, operations.GetHealthcareLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetHealthcareLocationRequest](../../pkg/models/operations/gethealthcarelocationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetHealthcareLocationResponse](../../pkg/models/operations/gethealthcarelocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareMedication

Retrieve a medication

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareMedication" method="get" path="/healthcare/{connection_id}/medication/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareMedication(ctx, operations.GetHealthcareMedicationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareMedication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetHealthcareMedicationRequest](../../pkg/models/operations/gethealthcaremedicationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetHealthcareMedicationResponse](../../pkg/models/operations/gethealthcaremedicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareObservation

Retrieve an observation

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareObservation" method="get" path="/healthcare/{connection_id}/observation/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareObservation(ctx, operations.GetHealthcareObservationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareObservation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetHealthcareObservationRequest](../../pkg/models/operations/gethealthcareobservationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetHealthcareObservationResponse](../../pkg/models/operations/gethealthcareobservationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcarePatient

Retrieve a patient

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcarePatient" method="get" path="/healthcare/{connection_id}/patient/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcarePatient(ctx, operations.GetHealthcarePatientRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePatient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetHealthcarePatientRequest](../../pkg/models/operations/gethealthcarepatientrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetHealthcarePatientResponse](../../pkg/models/operations/gethealthcarepatientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcarePractitioner

Retrieve a practitioner

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcarePractitioner" method="get" path="/healthcare/{connection_id}/practitioner/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcarePractitioner(ctx, operations.GetHealthcarePractitionerRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePractitioner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetHealthcarePractitionerRequest](../../pkg/models/operations/gethealthcarepractitionerrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetHealthcarePractitionerResponse](../../pkg/models/operations/gethealthcarepractitionerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareSlot

Retrieve a slot

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareSlot" method="get" path="/healthcare/{connection_id}/slot/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.GetHealthcareSlot(ctx, operations.GetHealthcareSlotRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareSlot != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHealthcareSlotRequest](../../pkg/models/operations/gethealthcareslotrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetHealthcareSlotResponse](../../pkg/models/operations/gethealthcareslotresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareAllergies

List all allergies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareAllergies" method="get" path="/healthcare/{connection_id}/allergy" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareAllergies(ctx, operations.ListHealthcareAllergiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAllergies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListHealthcareAllergiesRequest](../../pkg/models/operations/listhealthcareallergiesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListHealthcareAllergiesResponse](../../pkg/models/operations/listhealthcareallergiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareAppointments

List all appointments

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareAppointments" method="get" path="/healthcare/{connection_id}/appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareAppointments(ctx, operations.ListHealthcareAppointmentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAppointments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListHealthcareAppointmentsRequest](../../pkg/models/operations/listhealthcareappointmentsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListHealthcareAppointmentsResponse](../../pkg/models/operations/listhealthcareappointmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareConditions

List all conditions

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareConditions" method="get" path="/healthcare/{connection_id}/condition" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareConditions(ctx, operations.ListHealthcareConditionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareConditions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListHealthcareConditionsRequest](../../pkg/models/operations/listhealthcareconditionsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListHealthcareConditionsResponse](../../pkg/models/operations/listhealthcareconditionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareEncounters

List all encounters

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareEncounters" method="get" path="/healthcare/{connection_id}/encounter" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareEncounters(ctx, operations.ListHealthcareEncountersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareEncounters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListHealthcareEncountersRequest](../../pkg/models/operations/listhealthcareencountersrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListHealthcareEncountersResponse](../../pkg/models/operations/listhealthcareencountersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareImmunizations

List all immunizations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareImmunizations" method="get" path="/healthcare/{connection_id}/immunization" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareImmunizations(ctx, operations.ListHealthcareImmunizationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareImmunizations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListHealthcareImmunizationsRequest](../../pkg/models/operations/listhealthcareimmunizationsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListHealthcareImmunizationsResponse](../../pkg/models/operations/listhealthcareimmunizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareLocations" method="get" path="/healthcare/{connection_id}/location" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareLocations(ctx, operations.ListHealthcareLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListHealthcareLocationsRequest](../../pkg/models/operations/listhealthcarelocationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListHealthcareLocationsResponse](../../pkg/models/operations/listhealthcarelocationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareMedications

List all medications

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareMedications" method="get" path="/healthcare/{connection_id}/medication" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareMedications(ctx, operations.ListHealthcareMedicationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareMedications != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListHealthcareMedicationsRequest](../../pkg/models/operations/listhealthcaremedicationsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListHealthcareMedicationsResponse](../../pkg/models/operations/listhealthcaremedicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareObservations

List all observations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareObservations" method="get" path="/healthcare/{connection_id}/observation" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareObservations(ctx, operations.ListHealthcareObservationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareObservations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListHealthcareObservationsRequest](../../pkg/models/operations/listhealthcareobservationsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListHealthcareObservationsResponse](../../pkg/models/operations/listhealthcareobservationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcarePatients

List all patients

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcarePatients" method="get" path="/healthcare/{connection_id}/patient" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcarePatients(ctx, operations.ListHealthcarePatientsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePatients != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListHealthcarePatientsRequest](../../pkg/models/operations/listhealthcarepatientsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListHealthcarePatientsResponse](../../pkg/models/operations/listhealthcarepatientsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcarePractitioners

List all practitioners

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcarePractitioners" method="get" path="/healthcare/{connection_id}/practitioner" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcarePractitioners(ctx, operations.ListHealthcarePractitionersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePractitioners != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListHealthcarePractitionersRequest](../../pkg/models/operations/listhealthcarepractitionersrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListHealthcarePractitionersResponse](../../pkg/models/operations/listhealthcarepractitionersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareSlots

List all slots

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareSlots" method="get" path="/healthcare/{connection_id}/slot" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.ListHealthcareSlots(ctx, operations.ListHealthcareSlotsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareSlots != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListHealthcareSlotsRequest](../../pkg/models/operations/listhealthcareslotsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListHealthcareSlotsResponse](../../pkg/models/operations/listhealthcareslotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHealthcareAppointment

Update an appointment

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHealthcareAppointment" method="patch" path="/healthcare/{connection_id}/appointment/{id}" example="healthcare_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.PatchHealthcareAppointment(ctx, operations.PatchHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("clarus aptus cetera"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("d96cb461-044a-453e-bc7e-4e71a39ab8ec"),
            Metadata: []shared.HealthcareMetadata{},
            PractitionerIds: []string{},
            Reasons: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("cubicularis caput"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("61128108"),
                    Display: unifiedgosdk.Pointer("Annual physical"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-13T06:51:07.937Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchHealthcareAppointmentRequest](../../pkg/models/operations/patchhealthcareappointmentrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PatchHealthcareAppointmentResponse](../../pkg/models/operations/patchhealthcareappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHealthcarePatient

Update a patient

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHealthcarePatient" method="patch" path="/healthcare/{connection_id}/patient/{id}" example="healthcare_patient" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.PatchHealthcarePatient(ctx, operations.PatchHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("3645 Zoila Isle"),
                Address2: unifiedgosdk.Pointer("Suite 613"),
                City: unifiedgosdk.Pointer("Jacintostead"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("31579"),
                Region: unifiedgosdk.Pointer("Minnesota"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-03"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@yahoo.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("f9866fe1-60b4-4b05-a7ae-1b5e70803f43"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "92636004",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("en"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "659.223.0829 x590",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-15T00:39:33.032Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePatient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchHealthcarePatientRequest](../../pkg/models/operations/patchhealthcarepatientrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchHealthcarePatientResponse](../../pkg/models/operations/patchhealthcarepatientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHealthcareAppointment

Remove an appointment

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHealthcareAppointment" method="delete" path="/healthcare/{connection_id}/appointment/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.RemoveHealthcareAppointment(ctx, operations.RemoveHealthcareAppointmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.RemoveHealthcareAppointmentRequest](../../pkg/models/operations/removehealthcareappointmentrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.RemoveHealthcareAppointmentResponse](../../pkg/models/operations/removehealthcareappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHealthcareAppointment

Update an appointment

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHealthcareAppointment" method="put" path="/healthcare/{connection_id}/appointment/{id}" example="healthcare_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.UpdateHealthcareAppointment(ctx, operations.UpdateHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("clarus aptus cetera"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("d96cb461-044a-453e-bc7e-4e71a39ab8ec"),
            Metadata: []shared.HealthcareMetadata{},
            PractitionerIds: []string{},
            Reasons: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("cubicularis caput"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("61128108"),
                    Display: unifiedgosdk.Pointer("Annual physical"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-13T06:51:07.937Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateHealthcareAppointmentRequest](../../pkg/models/operations/updatehealthcareappointmentrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateHealthcareAppointmentResponse](../../pkg/models/operations/updatehealthcareappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHealthcarePatient

Update a patient

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHealthcarePatient" method="put" path="/healthcare/{connection_id}/patient/{id}" example="healthcare_patient" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Healthcare.UpdateHealthcarePatient(ctx, operations.UpdateHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("3645 Zoila Isle"),
                Address2: unifiedgosdk.Pointer("Suite 613"),
                City: unifiedgosdk.Pointer("Jacintostead"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("31579"),
                Region: unifiedgosdk.Pointer("Minnesota"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-03"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@yahoo.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("f9866fe1-60b4-4b05-a7ae-1b5e70803f43"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "92636004",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("en"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "659.223.0829 x590",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-15T00:39:33.032Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcarePatient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateHealthcarePatientRequest](../../pkg/models/operations/updatehealthcarepatientrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateHealthcarePatientResponse](../../pkg/models/operations/updatehealthcarepatientresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |