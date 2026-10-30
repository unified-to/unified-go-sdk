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
            DueAt: types.MustNewTimeFromString("2026-08-07T06:59:31.809Z"),
            ID: unifiedgosdk.Pointer("b53aa5dc-9113-409b-8523-8c01da58c896"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-12T04:54:41.306Z"),
                    ID: unifiedgosdk.Pointer("5ade1310-d572-4476-a528-2537600d8331"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-17T00:40:06.117Z"),
                    ID: unifiedgosdk.Pointer("4ab21197-88e7-41c2-ae54-2f3d2e4d3983"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-19T04:53:01.386Z"),
                    ID: unifiedgosdk.Pointer("bc0ed91f-8542-4047-b696-71d1e353b63a"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-07-30T14:01:34.582Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-16T13:01:46.428Z"),
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
            DueAt: types.MustNewTimeFromString("2026-08-07T06:59:31.834Z"),
            ID: unifiedgosdk.Pointer("7105cec8-bb7c-4d56-b55a-4a89a4ac10ec"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-12T04:54:41.331Z"),
                    ID: unifiedgosdk.Pointer("947906c3-1da3-42a9-aa69-3a9031c72ee8"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-17T00:40:06.142Z"),
                    ID: unifiedgosdk.Pointer("96e06d46-177d-4360-aa87-a934c2d0dce6"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-19T04:53:01.411Z"),
                    ID: unifiedgosdk.Pointer("e5ee4363-39d9-4742-96cb-cb565fcb3951"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-07-30T14:01:34.604Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-16T13:01:46.438Z"),
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
            DueAt: types.MustNewTimeFromString("2026-08-07T06:59:31.834Z"),
            ID: unifiedgosdk.Pointer("7105cec8-bb7c-4d56-b55a-4a89a4ac10ec"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-06-12T04:54:41.331Z"),
                    ID: unifiedgosdk.Pointer("947906c3-1da3-42a9-aa69-3a9031c72ee8"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Team-oriented secondary matrix",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-08-17T00:40:06.142Z"),
                    ID: unifiedgosdk.Pointer("96e06d46-177d-4360-aa87-a934c2d0dce6"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Exclusive client-server conglomeration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Hic explicabo odit."),
                    DueAt: types.MustNewTimeFromString("2026-10-19T04:53:01.411Z"),
                    ID: unifiedgosdk.Pointer("e5ee4363-39d9-4742-96cb-cb565fcb3951"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Synchronised immersive success",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("Proactive national protocol"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-07-30T14:01:34.604Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-16T13:01:46.438Z"),
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