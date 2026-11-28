# Attendance

## Overview

### Available Operations

* [CreateHrisAttendance](#createhrisattendance) - Create an attendance
* [GetHrisAttendance](#gethrisattendance) - Retrieve an attendance
* [ListHrisAttendances](#listhrisattendances) - List all attendances
* [PatchHrisAttendance](#patchhrisattendance) - Update an attendance
* [RemoveHrisAttendance](#removehrisattendance) - Remove an attendance
* [UpdateHrisAttendance](#updatehrisattendance) - Update an attendance

## CreateHrisAttendance

Creates a HrisAttendance. Provide the unified HrisAttendance fields in the request body; they are mapped to the provider's API. Returns the created HrisAttendance.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisAttendance" method="post" path="/hris/{connection_id}/attendance" example="hris_attendance" -->
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

    res, err := s.Attendance.CreateHrisAttendance(ctx, operations.CreateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("18002 Elm Street"),
                City: unifiedgosdk.Pointer("Port Gaetano"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("42344"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("MI"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-20T09:13:44.314Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-03-03T22:31:59.460Z"),
                    ID: unifiedgosdk.Pointer("301bc13d-1a2f-43aa-8dbb-1d2895feae93"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-15T04:35:00.877Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-05-13T02:08:13.200Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("6b2eeb8c-f3c6-4fef-a5df-ff4e3d49df97"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-12T21:16:26.386Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Africa/Ouagadougou"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T02:59:48.588Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateHrisAttendanceRequest](../../pkg/models/operations/createhrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateHrisAttendanceResponse](../../pkg/models/operations/createhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisAttendance

Retrieves a single HrisAttendance by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisAttendance" method="get" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Attendance.GetHrisAttendance(ctx, operations.GetHrisAttendanceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHrisAttendanceRequest](../../pkg/models/operations/gethrisattendancerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetHrisAttendanceResponse](../../pkg/models/operations/gethrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisAttendances

Returns a list of HrisAttendances. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisAttendances" method="get" path="/hris/{connection_id}/attendance" -->
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

    res, err := s.Attendance.ListHrisAttendances(ctx, operations.ListHrisAttendancesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListHrisAttendancesRequest](../../pkg/models/operations/listhrisattendancesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListHrisAttendancesResponse](../../pkg/models/operations/listhrisattendancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisAttendance

Updates the HrisAttendance identified by `id`. Provide the unified HrisAttendance fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisAttendance" method="patch" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Attendance.PatchHrisAttendance(ctx, operations.PatchHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("18002 Elm Street"),
                City: unifiedgosdk.Pointer("Port Gaetano"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("42344"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("MI"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-20T09:13:44.314Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-03-03T22:31:59.470Z"),
                    ID: unifiedgosdk.Pointer("a040b96b-218d-4b4b-8388-766e184e69bd"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-15T04:35:00.886Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-05-13T02:08:13.211Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("fc6eaf8c-256c-4ae5-96b3-9bc3417808f9"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-12T21:16:26.387Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Africa/Ouagadougou"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T02:59:48.590Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchHrisAttendanceRequest](../../pkg/models/operations/patchhrisattendancerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchHrisAttendanceResponse](../../pkg/models/operations/patchhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisAttendance

Deletes the HrisAttendance identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisAttendance" method="delete" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Attendance.RemoveHrisAttendance(ctx, operations.RemoveHrisAttendanceRequest{
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
| `request`                                                                                            | [operations.RemoveHrisAttendanceRequest](../../pkg/models/operations/removehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveHrisAttendanceResponse](../../pkg/models/operations/removehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisAttendance

Updates the HrisAttendance identified by `id`. Provide the unified HrisAttendance fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisAttendance" method="put" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Attendance.UpdateHrisAttendance(ctx, operations.UpdateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("18002 Elm Street"),
                City: unifiedgosdk.Pointer("Port Gaetano"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("42344"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("MI"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-20T09:13:44.314Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-03-03T22:31:59.470Z"),
                    ID: unifiedgosdk.Pointer("a040b96b-218d-4b4b-8388-766e184e69bd"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-15T04:35:00.886Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-05-13T02:08:13.211Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("fc6eaf8c-256c-4ae5-96b3-9bc3417808f9"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-12T21:16:26.387Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Africa/Ouagadougou"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T02:59:48.590Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateHrisAttendanceRequest](../../pkg/models/operations/updatehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateHrisAttendanceResponse](../../pkg/models/operations/updatehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |