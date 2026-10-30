# Patient

## Overview

### Available Operations

* [CreateHealthcarePatient](#createhealthcarepatient) - Create a patient
* [GetHealthcarePatient](#gethealthcarepatient) - Retrieve a patient
* [ListHealthcarePatients](#listhealthcarepatients) - List all patients
* [PatchHealthcarePatient](#patchhealthcarepatient) - Update a patient
* [UpdateHealthcarePatient](#updatehealthcarepatient) - Update a patient

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

    res, err := s.Patient.CreateHealthcarePatient(ctx, operations.CreateHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-20"),
            Contacts: []shared.HealthcarePatientContact{
                shared.HealthcarePatientContact{
                    Emails: []shared.HealthcareEmail{},
                    Name: unifiedgosdk.Pointer("Antonia Lebsack"),
                    Relationship: unifiedgosdk.Pointer("Spouse"),
                    Telephones: []shared.HealthcareTelephone{
                        shared.HealthcareTelephone{
                            Telephone: "(695) 328-7029 x4412",
                            Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                        },
                    },
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("d0117b96-94bd-49d4-8874-7fd3c573da48"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-07-01T03:02:51.808Z"),
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

    res, err := s.Patient.PatchHealthcarePatient(ctx, operations.PatchHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-20"),
            Contacts: []shared.HealthcarePatientContact{
                shared.HealthcarePatientContact{
                    Emails: []shared.HealthcareEmail{},
                    Name: unifiedgosdk.Pointer("Antonia Lebsack"),
                    Relationship: unifiedgosdk.Pointer("Spouse"),
                    Telephones: []shared.HealthcareTelephone{
                        shared.HealthcareTelephone{
                            Telephone: "(695) 328-7029 x4412",
                            Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                        },
                    },
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("3b1e578b-deac-4202-ba1f-d42c4068a019"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-07-01T03:02:51.846Z"),
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

    res, err := s.Patient.UpdateHealthcarePatient(ctx, operations.UpdateHealthcarePatientRequest{
        HealthcarePatient: shared.HealthcarePatient{
            Address: &shared.PropertyHealthcarePatientAddress{
                Address1: unifiedgosdk.Pointer("162 Kirlin Mall"),
                City: unifiedgosdk.Pointer("Dickinsonborough"),
                Country: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54031"),
                Region: unifiedgosdk.Pointer("Nebraska"),
            },
            Birthdate: unifiedgosdk.Pointer("2025-02-20"),
            Contacts: []shared.HealthcarePatientContact{
                shared.HealthcarePatientContact{
                    Emails: []shared.HealthcareEmail{},
                    Name: unifiedgosdk.Pointer("Antonia Lebsack"),
                    Relationship: unifiedgosdk.Pointer("Spouse"),
                    Telephones: []shared.HealthcareTelephone{
                        shared.HealthcareTelephone{
                            Telephone: "(695) 328-7029 x4412",
                            Type: shared.HealthcareTelephoneTypeMobile.ToPointer(),
                        },
                    },
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-06-05T21:08:04.917Z"),
            Emails: []shared.HealthcareEmail{
                shared.HealthcareEmail{
                    Email: "alejandrin.parker@icloud.com",
                    Type: shared.HealthcareEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alejandrin"),
            Gender: shared.HealthcarePatientGenderFemale.ToPointer(),
            ID: unifiedgosdk.Pointer("3b1e578b-deac-4202-ba1f-d42c4068a019"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-07-01T03:02:51.846Z"),
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