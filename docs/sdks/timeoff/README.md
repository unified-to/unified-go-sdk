# Timeoff

## Overview

### Available Operations

* [CreateHrisTimeoff](#createhristimeoff) - Create a timeoff
* [GetHrisTimeoff](#gethristimeoff) - Retrieve a timeoff
* [ListHrisTimeoffs](#listhristimeoffs) - List all timeoffs
* [PatchHrisTimeoff](#patchhristimeoff) - Update a timeoff
* [RemoveHrisTimeoff](#removehristimeoff) - Remove a timeoff
* [UpdateHrisTimeoff](#updatehristimeoff) - Update a timeoff

## CreateHrisTimeoff

Create a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeoff" method="post" path="/hris/{connection_id}/timeoff" example="hris_timeoff" -->
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

    res, err := s.Timeoff.CreateHrisTimeoff(ctx, operations.CreateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T23:47:57.715Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-25T10:24:09.882Z"),
            ID: unifiedgosdk.Pointer("d64b2498-657f-4ec9-be8e-b8dec44f4099"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-02T20:41:19.916Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-12T01:01:24.438Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisTimeoffRequest](../../pkg/models/operations/createhristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisTimeoffResponse](../../pkg/models/operations/createhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeoff

Retrieve a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeoff" method="get" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Timeoff.GetHrisTimeoff(ctx, operations.GetHrisTimeoffRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisTimeoffRequest](../../pkg/models/operations/gethristimeoffrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisTimeoffResponse](../../pkg/models/operations/gethristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeoffs

List all timeoffs

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeoffs" method="get" path="/hris/{connection_id}/timeoff" -->
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

    res, err := s.Timeoff.ListHrisTimeoffs(ctx, operations.ListHrisTimeoffsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoffs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisTimeoffsRequest](../../pkg/models/operations/listhristimeoffsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisTimeoffsResponse](../../pkg/models/operations/listhristimeoffsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeoff" method="patch" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
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

    res, err := s.Timeoff.PatchHrisTimeoff(ctx, operations.PatchHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T23:47:57.716Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-25T10:24:09.892Z"),
            ID: unifiedgosdk.Pointer("d5a69846-8928-4817-a555-32aee994342f"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-02T20:41:19.921Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-12T01:01:24.440Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisTimeoffRequest](../../pkg/models/operations/patchhristimeoffrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisTimeoffResponse](../../pkg/models/operations/patchhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeoff

Remove a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeoff" method="delete" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Timeoff.RemoveHrisTimeoff(ctx, operations.RemoveHrisTimeoffRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveHrisTimeoffRequest](../../pkg/models/operations/removehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisTimeoffResponse](../../pkg/models/operations/removehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeoff" method="put" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
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

    res, err := s.Timeoff.UpdateHrisTimeoff(ctx, operations.UpdateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T23:47:57.716Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-25T10:24:09.892Z"),
            ID: unifiedgosdk.Pointer("d5a69846-8928-4817-a555-32aee994342f"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-02T20:41:19.921Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-12T01:01:24.440Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisTimeoffRequest](../../pkg/models/operations/updatehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisTimeoffResponse](../../pkg/models/operations/updatehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |