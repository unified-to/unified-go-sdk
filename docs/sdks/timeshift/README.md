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

Create a timeshift

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T01:51:06.464Z"),
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
            EndAt: types.MustNewTimeFromString("2026-09-27T01:26:47.942Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("6b0e3410-8eef-4221-b88e-e4d4be208730"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T00:30:36.291Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-01T21:53:01.115Z"),
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

Retrieve a timeshift

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

List all timeshifts

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

Update a timeshift

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T01:51:06.472Z"),
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
            EndAt: types.MustNewTimeFromString("2026-09-27T01:26:47.962Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("860cce82-b177-4974-b6a6-dda841cb1baa"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T00:30:36.302Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-01T21:53:01.120Z"),
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

Remove a timeshift

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

Update a timeshift

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-18T01:51:06.472Z"),
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
            EndAt: types.MustNewTimeFromString("2026-09-27T01:26:47.962Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("860cce82-b177-4974-b6a6-dda841cb1baa"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-13T00:30:36.302Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-01T21:53:01.120Z"),
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