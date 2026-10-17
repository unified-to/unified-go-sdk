# Observation

## Overview

### Available Operations

* [CreateHealthcareObservation](#createhealthcareobservation) - Create an observation
* [GetHealthcareObservation](#gethealthcareobservation) - Retrieve an observation
* [ListHealthcareObservations](#listhealthcareobservations) - List all observations

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

    res, err := s.Observation.CreateHealthcareObservation(ctx, operations.CreateHealthcareObservationRequest{
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
            ID: unifiedgosdk.Pointer("35a10042-ca80-4f7b-befc-260f8f57027b"),
            Interpretation: shared.InterpretationCriticalHigh.ToPointer(),
            IssuedAt: types.MustNewTimeFromString("2025-05-08T22:30:56.954Z"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Blood pressure panel"),
            ObservedAt: types.MustNewTimeFromString("2025-01-17T23:33:24.346Z"),
            Status: shared.HealthcareObservationStatusCorrected.ToPointer(),
            Type: shared.HealthcareObservationTypeVitalSigns.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-20T21:00:04.210Z"),
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

    res, err := s.Observation.GetHealthcareObservation(ctx, operations.GetHealthcareObservationRequest{
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

    res, err := s.Observation.ListHealthcareObservations(ctx, operations.ListHealthcareObservationsRequest{
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