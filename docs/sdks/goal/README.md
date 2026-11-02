# Goal

## Overview

### Available Operations

* [CreatePerformanceGoal](#createperformancegoal) - Create a goal
* [GetHealthcareGoal](#gethealthcaregoal) - Retrieve a goal
* [GetPerformanceGoal](#getperformancegoal) - Retrieve a goal
* [ListHealthcareGoals](#listhealthcaregoals) - List all goals
* [ListPerformanceGoals](#listperformancegoals) - List all goals
* [PatchPerformanceGoal](#patchperformancegoal) - Update a goal
* [RemovePerformanceGoal](#removeperformancegoal) - Remove a goal
* [UpdatePerformanceGoal](#updateperformancegoal) - Update a goal

## CreatePerformanceGoal

Creates a PerformanceGoal. Provide the unified PerformanceGoal fields in the request body; they are mapped to the provider's API. Returns the created PerformanceGoal.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPerformanceGoal" method="post" path="/performance/{connection_id}/goal" example="performance_goal" -->
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

    res, err := s.Goal.CreatePerformanceGoal(ctx, operations.CreatePerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Suscipit suspendo vulnero vel facere valeo vallum degero."),
            DueAt: types.MustNewTimeFromString("2026-08-10T06:05:04.381Z"),
            ID: unifiedgosdk.Pointer("3e3986b7-9604-4c27-b527-0a37bbbe06a1"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-15T02:20:36.429Z"),
                    ID: unifiedgosdk.Pointer("b282ba8f-1753-44e2-9310-22b5efb8d287"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-20T00:02:56.358Z"),
                    ID: unifiedgosdk.Pointer("6ab4a85f-8d50-426e-a873-138ea54d1183"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-22T06:08:04.580Z"),
                    ID: unifiedgosdk.Pointer("c7aa3a0c-3eb8-4a42-9362-accd80821242"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-02T02:05:05.857Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-17T18:03:42.512Z"),
            Weight: unifiedgosdk.Pointer[float64](5.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceGoal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreatePerformanceGoalRequest](../../pkg/models/operations/createperformancegoalrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreatePerformanceGoalResponse](../../pkg/models/operations/createperformancegoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHealthcareGoal

Retrieves a single HealthcareGoal by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthcareGoal" method="get" path="/healthcare/{connection_id}/goal/{id}" -->
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

    res, err := s.Goal.GetHealthcareGoal(ctx, operations.GetHealthcareGoalRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareGoal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHealthcareGoalRequest](../../pkg/models/operations/gethealthcaregoalrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetHealthcareGoalResponse](../../pkg/models/operations/gethealthcaregoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPerformanceGoal

Retrieves a single PerformanceGoal by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceGoal" method="get" path="/performance/{connection_id}/goal/{id}" -->
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

    res, err := s.Goal.GetPerformanceGoal(ctx, operations.GetPerformanceGoalRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceGoal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPerformanceGoalRequest](../../pkg/models/operations/getperformancegoalrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetPerformanceGoalResponse](../../pkg/models/operations/getperformancegoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHealthcareGoals

Returns a list of HealthcareGoals. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHealthcareGoals" method="get" path="/healthcare/{connection_id}/goal" -->
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

    res, err := s.Goal.ListHealthcareGoals(ctx, operations.ListHealthcareGoalsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthcareGoals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListHealthcareGoalsRequest](../../pkg/models/operations/listhealthcaregoalsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListHealthcareGoalsResponse](../../pkg/models/operations/listhealthcaregoalsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPerformanceGoals

Returns a list of PerformanceGoals. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPerformanceGoals" method="get" path="/performance/{connection_id}/goal" -->
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

    res, err := s.Goal.ListPerformanceGoals(ctx, operations.ListPerformanceGoalsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceGoals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPerformanceGoalsRequest](../../pkg/models/operations/listperformancegoalsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListPerformanceGoalsResponse](../../pkg/models/operations/listperformancegoalsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPerformanceGoal

Updates the PerformanceGoal identified by `id`. Provide the unified PerformanceGoal fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPerformanceGoal" method="patch" path="/performance/{connection_id}/goal/{id}" example="performance_goal" -->
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

    res, err := s.Goal.PatchPerformanceGoal(ctx, operations.PatchPerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Suscipit suspendo vulnero vel facere valeo vallum degero."),
            DueAt: types.MustNewTimeFromString("2026-08-10T06:05:04.395Z"),
            ID: unifiedgosdk.Pointer("9fe66387-88b0-4545-b146-2cc1208ddebe"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-15T02:20:36.443Z"),
                    ID: unifiedgosdk.Pointer("f800ced8-c3f9-44fb-9e07-00eb6e7c3871"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-20T00:02:56.374Z"),
                    ID: unifiedgosdk.Pointer("5157fa90-7843-4ff4-8ef8-2ac0dbbf3351"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-22T06:08:04.596Z"),
                    ID: unifiedgosdk.Pointer("11990b2b-05ba-4015-8e0f-7e8baf5d3e42"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-02T02:05:05.869Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-17T18:03:42.518Z"),
            Weight: unifiedgosdk.Pointer[float64](5.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceGoal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchPerformanceGoalRequest](../../pkg/models/operations/patchperformancegoalrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchPerformanceGoalResponse](../../pkg/models/operations/patchperformancegoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePerformanceGoal

Deletes the PerformanceGoal identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePerformanceGoal" method="delete" path="/performance/{connection_id}/goal/{id}" -->
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

    res, err := s.Goal.RemovePerformanceGoal(ctx, operations.RemovePerformanceGoalRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemovePerformanceGoalRequest](../../pkg/models/operations/removeperformancegoalrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemovePerformanceGoalResponse](../../pkg/models/operations/removeperformancegoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePerformanceGoal

Updates the PerformanceGoal identified by `id`. Provide the unified PerformanceGoal fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePerformanceGoal" method="put" path="/performance/{connection_id}/goal/{id}" example="performance_goal" -->
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

    res, err := s.Goal.UpdatePerformanceGoal(ctx, operations.UpdatePerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Suscipit suspendo vulnero vel facere valeo vallum degero."),
            DueAt: types.MustNewTimeFromString("2026-08-10T06:05:04.395Z"),
            ID: unifiedgosdk.Pointer("9fe66387-88b0-4545-b146-2cc1208ddebe"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-15T02:20:36.443Z"),
                    ID: unifiedgosdk.Pointer("f800ced8-c3f9-44fb-9e07-00eb6e7c3871"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-20T00:02:56.374Z"),
                    ID: unifiedgosdk.Pointer("5157fa90-7843-4ff4-8ef8-2ac0dbbf3351"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-22T06:08:04.596Z"),
                    ID: unifiedgosdk.Pointer("11990b2b-05ba-4015-8e0f-7e8baf5d3e42"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-02T02:05:05.869Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-17T18:03:42.518Z"),
            Weight: unifiedgosdk.Pointer[float64](5.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceGoal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdatePerformanceGoalRequest](../../pkg/models/operations/updateperformancegoalrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdatePerformanceGoalResponse](../../pkg/models/operations/updateperformancegoalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |