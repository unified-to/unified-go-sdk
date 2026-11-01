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
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T19:38:21.275Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-19T20:54:22.886Z"),
                    ID: unifiedgosdk.Pointer("8d527382-1830-47eb-844d-c1e72138968b"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-03T15:43:31.913Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-29T01:00:28.093Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("d314bcc6-a518-4cae-bd60-828dc19c2220"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T13:45:34.410Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T19:36:55.482Z"),
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
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T19:38:21.276Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-19T20:54:22.907Z"),
                    ID: unifiedgosdk.Pointer("2900009d-ca09-46d4-a7b7-fb16df53f90d"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-03T15:43:31.931Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-29T01:00:28.115Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("ebefa7b9-7a7e-47ab-a28f-72a48f4f2f89"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T13:45:34.412Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T19:36:55.486Z"),
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
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T19:38:21.276Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-19T20:54:22.907Z"),
                    ID: unifiedgosdk.Pointer("2900009d-ca09-46d4-a7b7-fb16df53f90d"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-03T15:43:31.931Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-29T01:00:28.115Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("ebefa7b9-7a7e-47ab-a28f-72a48f4f2f89"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T13:45:34.412Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T19:36:55.486Z"),
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