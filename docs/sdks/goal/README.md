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
            DueAt: types.MustNewTimeFromString("2026-08-14T05:15:46.767Z"),
            ID: unifiedgosdk.Pointer("1c2e46c0-1f79-4173-a117-37027791dc07"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-18T23:17:56.211Z"),
                    ID: unifiedgosdk.Pointer("8c5acb55-a02f-49a9-9229-6359331c99d7"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-23T23:36:47.977Z"),
                    ID: unifiedgosdk.Pointer("c7174651-3d74-4871-a49b-2b8fa727f443"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-26T08:12:10.273Z"),
                    ID: unifiedgosdk.Pointer("9946adbf-f6af-4fe6-a814-c1a071919010"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-05T10:29:29.392Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-19T08:55:48.588Z"),
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
            DueAt: types.MustNewTimeFromString("2026-08-14T05:15:46.785Z"),
            ID: unifiedgosdk.Pointer("ef4633f8-dec2-4910-aab8-b1a98b7eccc3"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-18T23:17:56.228Z"),
                    ID: unifiedgosdk.Pointer("d0b471b9-e1c3-446c-88e4-b07acca3498f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-23T23:36:47.994Z"),
                    ID: unifiedgosdk.Pointer("865d09c4-fcc2-49c2-9762-b75acb81122f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-26T08:12:10.291Z"),
                    ID: unifiedgosdk.Pointer("bd29c50f-73f6-4082-9f2c-c4865e75699d"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-05T10:29:29.407Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-19T08:55:48.595Z"),
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
            DueAt: types.MustNewTimeFromString("2026-08-14T05:15:46.785Z"),
            ID: unifiedgosdk.Pointer("ef4633f8-dec2-4910-aab8-b1a98b7eccc3"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-18T23:17:56.228Z"),
                    ID: unifiedgosdk.Pointer("d0b471b9-e1c3-446c-88e4-b07acca3498f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-23T23:36:47.994Z"),
                    ID: unifiedgosdk.Pointer("865d09c4-fcc2-49c2-9762-b75acb81122f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-26T08:12:10.291Z"),
                    ID: unifiedgosdk.Pointer("bd29c50f-73f6-4082-9f2c-c4865e75699d"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-05T10:29:29.407Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-19T08:55:48.595Z"),
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