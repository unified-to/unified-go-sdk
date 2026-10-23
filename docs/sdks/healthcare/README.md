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

Creates a HealthcareAppointment. Provide the unified HealthcareAppointment fields in the request body; they are mapped to the provider's API. Returns the created HealthcareAppointment.

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
            ID: unifiedgosdk.Pointer("d085955c-472b-4845-8b7f-04d33f72b132"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T22:55:07.068Z"),
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

Creates a HealthcareObservation. Provide the unified HealthcareObservation fields in the request body; they are mapped to the provider's API. Returns the created HealthcareObservation.

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
            ID: unifiedgosdk.Pointer("f3c7fe5f-abad-4692-ba01-766ff43d794e"),
            Interpretation: shared.InterpretationCriticalHigh.ToPointer(),
            IssuedAt: types.MustNewTimeFromString("2025-05-12T17:12:08.304Z"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Blood pressure panel"),
            ObservedAt: types.MustNewTimeFromString("2025-01-21T06:51:39.970Z"),
            Status: shared.HealthcareObservationStatusCorrected.ToPointer(),
            Type: shared.HealthcareObservationTypeVitalSigns.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-26T00:05:49.468Z"),
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

Creates a HealthcarePatient. Provide the unified HealthcarePatient fields in the request body; they are mapped to the provider's API. Returns the created HealthcarePatient.

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
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-14"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("8781bf0a-6ce1-49ba-8cb7-7dbbe674ce81"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "57992636",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("zh"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "1-219-555-9023 x08295",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-24T18:17:45.978Z"),
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

Retrieves a single HealthcareAllergy by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareAppointment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareCondition by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareEncounter by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareImmunization by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareLocation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareMedication by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareObservation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcarePatient by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcarePractitioner by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HealthcareSlot by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of HealthcareAllergies. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareAppointments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareConditions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareEncounters. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareImmunizations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareLocations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareMedications. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareObservations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcarePatients. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcarePractitioners. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HealthcareSlots. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the HealthcareAppointment identified by `id`. Provide the unified HealthcareAppointment fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("63ae924e-6d92-4bda-b2c9-2707869ae2c8"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T22:55:07.089Z"),
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

Updates the HealthcarePatient identified by `id`. Provide the unified HealthcarePatient fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-14"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("0d2a6f22-3113-4f70-b49f-1e1b55a73b0c"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "57992636",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("zh"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "1-219-555-9023 x08295",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-24T18:17:46.012Z"),
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

Deletes the HealthcareAppointment identified by `id`.

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

Updates the HealthcareAppointment identified by `id`. Provide the unified HealthcareAppointment fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("63ae924e-6d92-4bda-b2c9-2707869ae2c8"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-20T22:55:07.089Z"),
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

Updates the HealthcarePatient identified by `id`. Provide the unified HealthcarePatient fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-14"),
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("0d2a6f22-3113-4f70-b49f-1e1b55a73b0c"),
            Identifiers: []shared.HealthcareIdentifier{
                shared.HealthcareIdentifier{
                    Type: shared.HealthcareIdentifierTypeMrn.ToPointer(),
                    Value: "57992636",
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            IsDeceased: unifiedgosdk.Pointer(false),
            Language: unifiedgosdk.Pointer("zh"),
            LastName: unifiedgosdk.Pointer("Parker"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Alejandrin Parker"),
            PractitionerIds: []string{},
            Telephones: []shared.HealthcareTelephone{
                shared.HealthcareTelephone{
                    Telephone: "1-219-555-9023 x08295",
                    Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2026-06-24T18:17:46.012Z"),
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