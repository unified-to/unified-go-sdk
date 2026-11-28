# Appointment

## Overview

### Available Operations

* [CreateDmsAppointment](#createdmsappointment) - Create an appointment
* [CreateHealthcareAppointment](#createhealthcareappointment) - Create an appointment
* [GetDmsAppointment](#getdmsappointment) - Retrieve an appointment
* [GetHealthcareAppointment](#gethealthcareappointment) - Retrieve an appointment
* [ListDmsAppointments](#listdmsappointments) - List all appointments
* [ListHealthcareAppointments](#listhealthcareappointments) - List all appointments
* [PatchDmsAppointment](#patchdmsappointment) - Update an appointment
* [PatchHealthcareAppointment](#patchhealthcareappointment) - Update an appointment
* [RemoveDmsAppointment](#removedmsappointment) - Remove an appointment
* [RemoveHealthcareAppointment](#removehealthcareappointment) - Remove an appointment
* [UpdateDmsAppointment](#updatedmsappointment) - Update an appointment
* [UpdateHealthcareAppointment](#updatehealthcareappointment) - Update an appointment

## CreateDmsAppointment

Creates a DmsAppointment. Provide the unified DmsAppointment fields in the request body; they are mapped to the provider's API. Returns the created DmsAppointment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsAppointment" method="post" path="/dms/{connection_id}/appointment" example="dms_appointment" -->
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

    res, err := s.Appointment.CreateDmsAppointment(ctx, operations.CreateDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T22:12:38.146Z"),
            ID: unifiedgosdk.Pointer("54322fa7-868e-4452-9426-c457536633f9"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("184dc7d5-0460-49ad-9610-a8530dc54a7e"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("c5beac99-3cc1-4db3-b4a6-640f697e19f7"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("9d377581-f85d-45e6-b010-9f68d5e19785"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-08T15:12:38.146Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T03:46:05.215Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDmsAppointmentRequest](../../pkg/models/operations/createdmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDmsAppointmentResponse](../../pkg/models/operations/createdmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Appointment.CreateHealthcareAppointment(ctx, operations.CreateHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("cubicularis caput viriliter"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("347b259f-f1ca-4ba7-8de3-03df2f2207b2"),
            IsVirtual: unifiedgosdk.Pointer(true),
            MeetingURL: unifiedgosdk.Pointer("https://cautious-cannon.com/"),
            Metadata: []shared.HealthcareMetadata{},
            Notes: unifiedgosdk.Pointer("Casus anser ver alo beatus."),
            PractitionerIds: []string{},
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("New patient visit"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-29T09:22:09.184Z"),
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

## GetDmsAppointment

Retrieves a single DmsAppointment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsAppointment" method="get" path="/dms/{connection_id}/appointment/{id}" -->
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

    res, err := s.Appointment.GetDmsAppointment(ctx, operations.GetDmsAppointmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDmsAppointmentRequest](../../pkg/models/operations/getdmsappointmentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDmsAppointmentResponse](../../pkg/models/operations/getdmsappointmentresponse.md), error**

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

    res, err := s.Appointment.GetHealthcareAppointment(ctx, operations.GetHealthcareAppointmentRequest{
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

## ListDmsAppointments

Returns a list of DmsAppointments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsAppointments" method="get" path="/dms/{connection_id}/appointment" -->
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

    res, err := s.Appointment.ListDmsAppointments(ctx, operations.ListDmsAppointmentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDmsAppointmentsRequest](../../pkg/models/operations/listdmsappointmentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDmsAppointmentsResponse](../../pkg/models/operations/listdmsappointmentsresponse.md), error**

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

    res, err := s.Appointment.ListHealthcareAppointments(ctx, operations.ListHealthcareAppointmentsRequest{
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

## PatchDmsAppointment

Updates the DmsAppointment identified by `id`. Provide the unified DmsAppointment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsAppointment" method="patch" path="/dms/{connection_id}/appointment/{id}" example="dms_appointment" -->
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

    res, err := s.Appointment.PatchDmsAppointment(ctx, operations.PatchDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T22:12:38.157Z"),
            ID: unifiedgosdk.Pointer("d7ca4f59-9666-4ec4-8468-fcdd8dfb165b"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("71c984ba-6761-48cb-9944-25fa212c7acd"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("671ad0ac-a45c-4819-8c30-9e3072f1f514"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("1d41690b-17b4-4119-813b-fec6a9070ccd"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-08T15:12:38.157Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T03:46:05.221Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDmsAppointmentRequest](../../pkg/models/operations/patchdmsappointmentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDmsAppointmentResponse](../../pkg/models/operations/patchdmsappointmentresponse.md), error**

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

    res, err := s.Appointment.PatchHealthcareAppointment(ctx, operations.PatchHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("cubicularis caput viriliter"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("bbf91cbf-6e4b-43a5-bb65-0ca8b1519ccf"),
            IsVirtual: unifiedgosdk.Pointer(true),
            MeetingURL: unifiedgosdk.Pointer("https://cautious-cannon.com/"),
            Metadata: []shared.HealthcareMetadata{},
            Notes: unifiedgosdk.Pointer("Casus anser ver alo beatus."),
            PractitionerIds: []string{},
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("New patient visit"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-29T09:22:09.195Z"),
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

## RemoveDmsAppointment

Deletes the DmsAppointment identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeDmsAppointment" method="delete" path="/dms/{connection_id}/appointment/{id}" -->
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

    res, err := s.Appointment.RemoveDmsAppointment(ctx, operations.RemoveDmsAppointmentRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveDmsAppointmentRequest](../../pkg/models/operations/removedmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveDmsAppointmentResponse](../../pkg/models/operations/removedmsappointmentresponse.md), error**

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

    res, err := s.Appointment.RemoveHealthcareAppointment(ctx, operations.RemoveHealthcareAppointmentRequest{
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

## UpdateDmsAppointment

Updates the DmsAppointment identified by `id`. Provide the unified DmsAppointment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsAppointment" method="put" path="/dms/{connection_id}/appointment/{id}" example="dms_appointment" -->
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

    res, err := s.Appointment.UpdateDmsAppointment(ctx, operations.UpdateDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T22:12:38.157Z"),
            ID: unifiedgosdk.Pointer("d7ca4f59-9666-4ec4-8468-fcdd8dfb165b"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("71c984ba-6761-48cb-9944-25fa212c7acd"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("671ad0ac-a45c-4819-8c30-9e3072f1f514"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("1d41690b-17b4-4119-813b-fec6a9070ccd"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-08T15:12:38.157Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-07T03:46:05.221Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDmsAppointmentRequest](../../pkg/models/operations/updatedmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDmsAppointmentResponse](../../pkg/models/operations/updatedmsappointmentresponse.md), error**

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

    res, err := s.Appointment.UpdateHealthcareAppointment(ctx, operations.UpdateHealthcareAppointmentRequest{
        HealthcareAppointment: shared.HealthcareAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T08:16:20.874Z"),
            Description: unifiedgosdk.Pointer("cubicularis caput viriliter"),
            EndAt: types.MustNewTimeFromString("2023-12-12T13:04:41.451Z"),
            ID: unifiedgosdk.Pointer("bbf91cbf-6e4b-43a5-bb65-0ca8b1519ccf"),
            IsVirtual: unifiedgosdk.Pointer(true),
            MeetingURL: unifiedgosdk.Pointer("https://cautious-cannon.com/"),
            Metadata: []shared.HealthcareMetadata{},
            Notes: unifiedgosdk.Pointer("Casus anser ver alo beatus."),
            PractitionerIds: []string{},
            ServiceTypes: []shared.HealthcareCode{
                shared.HealthcareCode{
                    Code: unifiedgosdk.Pointer("12762740"),
                    Display: unifiedgosdk.Pointer("New patient visit"),
                    System: unifiedgosdk.Pointer("http://snomed.info/sct"),
                },
            },
            SlotIds: []string{},
            StartAt: types.MustNewTimeFromString("2023-12-12T12:34:41.451Z"),
            Status: shared.HealthcareAppointmentStatusBooked.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-29T09:22:09.195Z"),
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