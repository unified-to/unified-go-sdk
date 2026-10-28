# Benefit

## Overview

### Available Operations

* [CreateHrisBenefit](#createhrisbenefit) - Create a benefit
* [GetHrisBenefit](#gethrisbenefit) - Retrieve a benefit
* [ListHrisBenefits](#listhrisbenefits) - List all benefits
* [PatchHrisBenefit](#patchhrisbenefit) - Update a benefit
* [RemoveHrisBenefit](#removehrisbenefit) - Remove a benefit
* [UpdateHrisBenefit](#updatehrisbenefit) - Update a benefit

## CreateHrisBenefit

Creates a HrisBenefit. Provide the unified HrisBenefit fields in the request body; they are mapped to the provider's API. Returns the created HrisBenefit.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBenefit" method="post" path="/hris/{connection_id}/benefit" example="hris_benefit" -->
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

    res, err := s.Benefit.CreateHrisBenefit(ctx, operations.CreateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("0cd1b177-5f3d-4867-b7c7-5896e7e84399"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T21:57:34.549Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisBenefitRequest](../../pkg/models/operations/createhrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisBenefitResponse](../../pkg/models/operations/createhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBenefit

Retrieves a single HrisBenefit by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBenefit" method="get" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Benefit.GetHrisBenefit(ctx, operations.GetHrisBenefitRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisBenefitRequest](../../pkg/models/operations/gethrisbenefitrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisBenefitResponse](../../pkg/models/operations/gethrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBenefits

Returns a list of HrisBenefits. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBenefits" method="get" path="/hris/{connection_id}/benefit" -->
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

    res, err := s.Benefit.ListHrisBenefits(ctx, operations.ListHrisBenefitsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisBenefitsRequest](../../pkg/models/operations/listhrisbenefitsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisBenefitsResponse](../../pkg/models/operations/listhrisbenefitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBenefit

Updates the HrisBenefit identified by `id`. Provide the unified HrisBenefit fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBenefit" method="patch" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Benefit.PatchHrisBenefit(ctx, operations.PatchHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("5e3897eb-3bfe-4de7-9445-fa2e0ff4ea97"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T21:57:34.556Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisBenefitRequest](../../pkg/models/operations/patchhrisbenefitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisBenefitResponse](../../pkg/models/operations/patchhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBenefit

Deletes the HrisBenefit identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBenefit" method="delete" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Benefit.RemoveHrisBenefit(ctx, operations.RemoveHrisBenefitRequest{
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
| `request`                                                                                      | [operations.RemoveHrisBenefitRequest](../../pkg/models/operations/removehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisBenefitResponse](../../pkg/models/operations/removehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBenefit

Updates the HrisBenefit identified by `id`. Provide the unified HrisBenefit fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBenefit" method="put" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Benefit.UpdateHrisBenefit(ctx, operations.UpdateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("5e3897eb-3bfe-4de7-9445-fa2e0ff4ea97"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T21:57:34.556Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisBenefitRequest](../../pkg/models/operations/updatehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisBenefitResponse](../../pkg/models/operations/updatehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |