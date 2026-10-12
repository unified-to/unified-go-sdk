# Patient

## Overview

### Available Operations

* [CreateHealthcarePatient](#createhealthcarepatient) - Create a patient
* [GetHealthcarePatient](#gethealthcarepatient) - Retrieve a patient
* [ListHealthcarePatients](#listhealthcarepatients) - List all patients
* [PatchHealthcarePatient](#patchhealthcarepatient) - Update a patient
* [UpdateHealthcarePatient](#updatehealthcarepatient) - Update a patient

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

    res, err := s.Patient.CreateHealthcarePatient(ctx, operations.CreateHealthcarePatientRequest{
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
            ID: unifiedgosdk.Pointer("4816e8cc-1092-4a6c-81b5-6613b67d1536"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-06-14T06:56:01.220Z"),
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

    res, err := s.Patient.GetHealthcarePatient(ctx, operations.GetHealthcarePatientRequest{
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

    res, err := s.Patient.ListHealthcarePatients(ctx, operations.ListHealthcarePatientsRequest{
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

    res, err := s.Patient.PatchHealthcarePatient(ctx, operations.PatchHealthcarePatientRequest{
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
            ID: unifiedgosdk.Pointer("f1dbfa20-dd10-4a99-b2d9-30bb5d13bef2"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-06-14T06:56:01.249Z"),
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

    res, err := s.Patient.UpdateHealthcarePatient(ctx, operations.UpdateHealthcarePatientRequest{
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
            ID: unifiedgosdk.Pointer("f1dbfa20-dd10-4a99-b2d9-30bb5d13bef2"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-06-14T06:56:01.249Z"),
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