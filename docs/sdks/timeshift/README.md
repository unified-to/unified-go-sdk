# Timeshift

## Overview

### Available Operations

* [CreateHrisTimeshift](#createhristimeshift) - Create a timeshift
* [GetHrisTimeshift](#gethristimeshift) - Retrieve a timeshift
* [ListHrisTimeshifts](#listhristimeshifts) - List all timeshifts
* [PatchHrisTimeshift](#patchhristimeshift) - Update a timeshift
* [RemoveHrisTimeshift](#removehristimeshift) - Remove a timeshift
* [UpdateHrisTimeshift](#updatehristimeshift) - Update a timeshift

## CreateHrisTimeshift

Creates a HrisTimeshift. Provide the unified HrisTimeshift fields in the request body; they are mapped to the provider's API. Returns the created HrisTimeshift.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeshift" method="post" path="/hris/{connection_id}/timeshift" example="hris_timeshift" -->
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

    res, err := s.Timeshift.CreateHrisTimeshift(ctx, operations.CreateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T11:45:36.212Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-28T01:00:15.668Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("92b01fd8-ed4d-4643-95ba-0d8a8b541854"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T13:37:31.157Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-02T04:23:45.080Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisTimeshiftRequest](../../pkg/models/operations/createhristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisTimeshiftResponse](../../pkg/models/operations/createhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeshift

Retrieves a single HrisTimeshift by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeshift" method="get" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Timeshift.GetHrisTimeshift(ctx, operations.GetHrisTimeshiftRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisTimeshiftRequest](../../pkg/models/operations/gethristimeshiftrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisTimeshiftResponse](../../pkg/models/operations/gethristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeshifts

Returns a list of HrisTimeshifts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeshifts" method="get" path="/hris/{connection_id}/timeshift" -->
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

    res, err := s.Timeshift.ListHrisTimeshifts(ctx, operations.ListHrisTimeshiftsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshifts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTimeshiftsRequest](../../pkg/models/operations/listhristimeshiftsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTimeshiftsResponse](../../pkg/models/operations/listhristimeshiftsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeshift

Updates the HrisTimeshift identified by `id`. Provide the unified HrisTimeshift fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeshift" method="patch" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
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

    res, err := s.Timeshift.PatchHrisTimeshift(ctx, operations.PatchHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T11:45:36.221Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-28T01:00:15.689Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("be8b2098-4454-4b17-a20f-c3f23a1d0e71"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T13:37:31.169Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-02T04:23:45.085Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisTimeshiftRequest](../../pkg/models/operations/patchhristimeshiftrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisTimeshiftResponse](../../pkg/models/operations/patchhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeshift

Deletes the HrisTimeshift identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeshift" method="delete" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Timeshift.RemoveHrisTimeshift(ctx, operations.RemoveHrisTimeshiftRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveHrisTimeshiftRequest](../../pkg/models/operations/removehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisTimeshiftResponse](../../pkg/models/operations/removehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeshift

Updates the HrisTimeshift identified by `id`. Provide the unified HrisTimeshift fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeshift" method="put" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
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

    res, err := s.Timeshift.UpdateHrisTimeshift(ctx, operations.UpdateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T11:45:36.221Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-28T01:00:15.689Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("be8b2098-4454-4b17-a20f-c3f23a1d0e71"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T13:37:31.169Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-02T04:23:45.085Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisTimeshiftRequest](../../pkg/models/operations/updatehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisTimeshiftResponse](../../pkg/models/operations/updatehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |