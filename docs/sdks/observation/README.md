# Observation

## Overview

### Available Operations

* [CreateHealthcareObservation](#createhealthcareobservation) - Create an observation
* [GetHealthcareObservation](#gethealthcareobservation) - Retrieve an observation
* [ListHealthcareObservations](#listhealthcareobservations) - List all observations

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
            ID: unifiedgosdk.Pointer("a456fb13-f2ba-40ef-b4f4-a317532b74e1"),
            Interpretation: shared.InterpretationCriticalHigh.ToPointer(),
            IssuedAt: types.MustNewTimeFromString("2025-05-17T17:21:23.086Z"),
            Metadata: []shared.HealthcareMetadata{},
            Name: unifiedgosdk.Pointer("Blood pressure panel"),
            ObservedAt: types.MustNewTimeFromString("2025-01-25T15:56:04.396Z"),
            Status: shared.HealthcareObservationStatusCorrected.ToPointer(),
            Type: shared.HealthcareObservationTypeVitalSigns.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-01T19:11:30.118Z"),
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