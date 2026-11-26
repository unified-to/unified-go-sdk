# Performance

## Overview

### Available Operations

* [CreatePerformanceFeedback](#createperformancefeedback) - Create a feedback
* [CreatePerformanceGoal](#createperformancegoal) - Create a goal
* [GetPerformanceCycle](#getperformancecycle) - Retrieve a cycle
* [GetPerformanceFeedback](#getperformancefeedback) - Retrieve a feedback
* [GetPerformanceGoal](#getperformancegoal) - Retrieve a goal
* [GetPerformanceReview](#getperformancereview) - Retrieve a review
* [ListPerformanceCycles](#listperformancecycles) - List all cycles
* [ListPerformanceFeedbacks](#listperformancefeedbacks) - List all feedbacks
* [ListPerformanceGoals](#listperformancegoals) - List all goals
* [ListPerformanceReviews](#listperformancereviews) - List all reviews
* [PatchPerformanceGoal](#patchperformancegoal) - Update a goal
* [RemovePerformanceGoal](#removeperformancegoal) - Remove a goal
* [UpdatePerformanceGoal](#updateperformancegoal) - Update a goal

## CreatePerformanceFeedback

Creates a PerformanceFeedback. Provide the unified PerformanceFeedback fields in the request body; they are mapped to the provider's API. Returns the created PerformanceFeedback.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPerformanceFeedback" method="post" path="/performance/{connection_id}/feedback" example="performance_feedback" -->
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

    res, err := s.Performance.CreatePerformanceFeedback(ctx, operations.CreatePerformanceFeedbackRequest{
        PerformanceFeedback: shared.PerformanceFeedback{
            CreatedAt: types.MustNewTimeFromString("2023-04-11T16:21:53.862Z"),
            ID: unifiedgosdk.Pointer("0ca1f8c9-f791-4cca-8899-2ff8aa8d6e48"),
            IsVisible: unifiedgosdk.Pointer(true),
            Message: unifiedgosdk.Pointer("Aestus viscus curo iure cernuus contigo damno decor carpo succedo est tergiversatio."),
            Metadata: []shared.PerformanceMetadata{},
            Tags: []string{
                "blind",
                "mathematics",
            },
            Type: shared.PerformanceFeedbackTypePraise.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-09-28T13:22:47.302Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceFeedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreatePerformanceFeedbackRequest](../../pkg/models/operations/createperformancefeedbackrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreatePerformanceFeedbackResponse](../../pkg/models/operations/createperformancefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Performance.CreatePerformanceGoal(ctx, operations.CreatePerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Vel facere valeo vallum degero vado thymum patruus."),
            DueAt: types.MustNewTimeFromString("2026-09-02T12:09:20.341Z"),
            ID: unifiedgosdk.Pointer("0eaa5e2a-c0a3-4d9d-a478-7d9b6c986071"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-07-07T19:22:49.764Z"),
                    ID: unifiedgosdk.Pointer("4b569de0-ab02-427c-b7c8-f1cc60922f4e"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Balanced system-worthy collaboration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-09-12T08:22:57.968Z"),
                    ID: unifiedgosdk.Pointer("bcb44dc4-52d4-4902-ae42-4814ffe1b165"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Persevering zero tolerance forecast",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Adiuvo ullus magnam eos cohors."),
                    DueAt: types.MustNewTimeFromString("2026-11-15T05:08:59.489Z"),
                    ID: unifiedgosdk.Pointer("10b063d5-6fd3-45ea-acc7-74c2f2c05a33"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Networked logistical toolset",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("User-centric radical parallelism"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-21T17:32:31.917Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-27T05:57:47.757Z"),
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

## GetPerformanceCycle

Retrieves a single PerformanceCycle by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceCycle" method="get" path="/performance/{connection_id}/cycle/{id}" -->
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

    res, err := s.Performance.GetPerformanceCycle(ctx, operations.GetPerformanceCycleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceCycle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPerformanceCycleRequest](../../pkg/models/operations/getperformancecyclerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPerformanceCycleResponse](../../pkg/models/operations/getperformancecycleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPerformanceFeedback

Retrieves a single PerformanceFeedback by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceFeedback" method="get" path="/performance/{connection_id}/feedback/{id}" -->
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

    res, err := s.Performance.GetPerformanceFeedback(ctx, operations.GetPerformanceFeedbackRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceFeedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetPerformanceFeedbackRequest](../../pkg/models/operations/getperformancefeedbackrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetPerformanceFeedbackResponse](../../pkg/models/operations/getperformancefeedbackresponse.md), error**

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

    res, err := s.Performance.GetPerformanceGoal(ctx, operations.GetPerformanceGoalRequest{
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

## GetPerformanceReview

Retrieves a single PerformanceReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceReview" method="get" path="/performance/{connection_id}/review/{id}" -->
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

    res, err := s.Performance.GetPerformanceReview(ctx, operations.GetPerformanceReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPerformanceReviewRequest](../../pkg/models/operations/getperformancereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetPerformanceReviewResponse](../../pkg/models/operations/getperformancereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPerformanceCycles

Returns a list of PerformanceCycles. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPerformanceCycles" method="get" path="/performance/{connection_id}/cycle" -->
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

    res, err := s.Performance.ListPerformanceCycles(ctx, operations.ListPerformanceCyclesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceCycles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListPerformanceCyclesRequest](../../pkg/models/operations/listperformancecyclesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListPerformanceCyclesResponse](../../pkg/models/operations/listperformancecyclesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPerformanceFeedbacks

Returns a list of PerformanceFeedbacks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPerformanceFeedbacks" method="get" path="/performance/{connection_id}/feedback" -->
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

    res, err := s.Performance.ListPerformanceFeedbacks(ctx, operations.ListPerformanceFeedbacksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceFeedbacks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListPerformanceFeedbacksRequest](../../pkg/models/operations/listperformancefeedbacksrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListPerformanceFeedbacksResponse](../../pkg/models/operations/listperformancefeedbacksresponse.md), error**

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

    res, err := s.Performance.ListPerformanceGoals(ctx, operations.ListPerformanceGoalsRequest{
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

## ListPerformanceReviews

Returns a list of PerformanceReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPerformanceReviews" method="get" path="/performance/{connection_id}/review" -->
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

    res, err := s.Performance.ListPerformanceReviews(ctx, operations.ListPerformanceReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListPerformanceReviewsRequest](../../pkg/models/operations/listperformancereviewsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListPerformanceReviewsResponse](../../pkg/models/operations/listperformancereviewsresponse.md), error**

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

    res, err := s.Performance.PatchPerformanceGoal(ctx, operations.PatchPerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Vel facere valeo vallum degero vado thymum patruus."),
            DueAt: types.MustNewTimeFromString("2026-09-02T12:09:20.364Z"),
            ID: unifiedgosdk.Pointer("1d66cb95-141a-4437-8888-a5f639668328"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-07-07T19:22:49.786Z"),
                    ID: unifiedgosdk.Pointer("bf49823c-76cb-414e-93a6-5dbb629d2e9d"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Balanced system-worthy collaboration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-09-12T08:22:57.991Z"),
                    ID: unifiedgosdk.Pointer("b61488e1-abd3-47eb-8d0c-41756e1c2a7f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Persevering zero tolerance forecast",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Adiuvo ullus magnam eos cohors."),
                    DueAt: types.MustNewTimeFromString("2026-11-15T05:08:59.512Z"),
                    ID: unifiedgosdk.Pointer("8ea37be4-dbc3-42e9-9d46-7f4b166505c7"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Networked logistical toolset",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("User-centric radical parallelism"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-21T17:32:31.936Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-27T05:57:47.766Z"),
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

    res, err := s.Performance.RemovePerformanceGoal(ctx, operations.RemovePerformanceGoalRequest{
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

    res, err := s.Performance.UpdatePerformanceGoal(ctx, operations.UpdatePerformanceGoalRequest{
        PerformanceGoal: shared.PerformanceGoal{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T20:43:07.380Z"),
            Description: unifiedgosdk.Pointer("Vel facere valeo vallum degero vado thymum patruus."),
            DueAt: types.MustNewTimeFromString("2026-09-02T12:09:20.364Z"),
            ID: unifiedgosdk.Pointer("1d66cb95-141a-4437-8888-a5f639668328"),
            Metadata: []shared.PerformanceMetadata{},
            Milestones: []shared.PerformanceGoalMilestone{
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](10.0),
                    DueAt: types.MustNewTimeFromString("2026-07-07T19:22:49.786Z"),
                    ID: unifiedgosdk.Pointer("bf49823c-76cb-414e-93a6-5dbb629d2e9d"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Balanced system-worthy collaboration",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Unit: unifiedgosdk.Pointer("%"),
                    Weight: unifiedgosdk.Pointer[float64](7.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](0.0),
                    DueAt: types.MustNewTimeFromString("2026-09-12T08:22:57.991Z"),
                    ID: unifiedgosdk.Pointer("b61488e1-abd3-47eb-8d0c-41756e1c2a7f"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Persevering zero tolerance forecast",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
                shared.PerformanceGoalMilestone{
                    CurrentValue: unifiedgosdk.Pointer[float64](31.0),
                    Description: unifiedgosdk.Pointer("Adiuvo ullus magnam eos cohors."),
                    DueAt: types.MustNewTimeFromString("2026-11-15T05:08:59.512Z"),
                    ID: unifiedgosdk.Pointer("8ea37be4-dbc3-42e9-9d46-7f4b166505c7"),
                    IsCompleted: unifiedgosdk.Pointer(true),
                    Name: "Networked logistical toolset",
                    TargetValue: unifiedgosdk.Pointer[float64](32.0),
                    Weight: unifiedgosdk.Pointer[float64](5.0),
                },
            },
            Name: unifiedgosdk.Pointer("User-centric radical parallelism"),
            Progress: unifiedgosdk.Pointer[float64](3.0),
            StartAt: types.MustNewTimeFromString("2025-08-21T17:32:31.936Z"),
            Status: shared.PerformanceGoalStatusClosed.ToPointer(),
            Type: shared.PerformanceGoalSchemasTypeCompany.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-09-27T05:57:47.766Z"),
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