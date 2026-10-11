# Deduction

## Overview

### Available Operations

* [CreateHrisDeduction](#createhrisdeduction) - Create a deduction
* [GetHrisDeduction](#gethrisdeduction) - Retrieve a deduction
* [ListHrisDeductions](#listhrisdeductions) - List all deductions
* [PatchHrisDeduction](#patchhrisdeduction) - Update a deduction
* [RemoveHrisDeduction](#removehrisdeduction) - Remove a deduction
* [UpdateHrisDeduction](#updatehrisdeduction) - Update a deduction

## CreateHrisDeduction

Create a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDeduction" method="post" path="/hris/{connection_id}/deduction" example="hris_deduction" -->
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

    res, err := s.Deduction.CreateHrisDeduction(ctx, operations.CreateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-13T20:04:08.022Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("05d38118-0f78-4663-824a-01818995c24c"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-07T17:01:33.631Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-16T03:23:54.844Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisDeductionRequest](../../pkg/models/operations/createhrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisDeductionResponse](../../pkg/models/operations/createhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDeduction

Retrieve a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDeduction" method="get" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Deduction.GetHrisDeduction(ctx, operations.GetHrisDeductionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisDeductionRequest](../../pkg/models/operations/gethrisdeductionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisDeductionResponse](../../pkg/models/operations/gethrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDeductions

List all deductions

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDeductions" method="get" path="/hris/{connection_id}/deduction" -->
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

    res, err := s.Deduction.ListHrisDeductions(ctx, operations.ListHrisDeductionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeductions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisDeductionsRequest](../../pkg/models/operations/listhrisdeductionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisDeductionsResponse](../../pkg/models/operations/listhrisdeductionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDeduction" method="patch" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Deduction.PatchHrisDeduction(ctx, operations.PatchHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-13T20:04:08.043Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("694f5bac-c102-49b9-a80c-f01052d74fc7"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-07T17:01:33.648Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-16T03:23:54.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisDeductionRequest](../../pkg/models/operations/patchhrisdeductionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisDeductionResponse](../../pkg/models/operations/patchhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDeduction

Remove a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDeduction" method="delete" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Deduction.RemoveHrisDeduction(ctx, operations.RemoveHrisDeductionRequest{
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
| `request`                                                                                          | [operations.RemoveHrisDeductionRequest](../../pkg/models/operations/removehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisDeductionResponse](../../pkg/models/operations/removehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDeduction" method="put" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Deduction.UpdateHrisDeduction(ctx, operations.UpdateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-13T20:04:08.043Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("694f5bac-c102-49b9-a80c-f01052d74fc7"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-07T17:01:33.648Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-16T03:23:54.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisDeductionRequest](../../pkg/models/operations/updatehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisDeductionResponse](../../pkg/models/operations/updatehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |