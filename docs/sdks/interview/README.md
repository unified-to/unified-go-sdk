# Interview

## Overview

### Available Operations

* [CreateAtsInterview](#createatsinterview) - Create an interview
* [GetAtsInterview](#getatsinterview) - Retrieve an interview
* [ListAtsInterviews](#listatsinterviews) - List all interviews
* [PatchAtsInterview](#patchatsinterview) - Update an interview
* [RemoveAtsInterview](#removeatsinterview) - Remove an interview
* [UpdateAtsInterview](#updateatsinterview) - Update an interview

## CreateAtsInterview

Create an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsInterview" method="post" path="/ats/{connection_id}/interview" example="ats_interview" -->
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

    res, err := s.Interview.CreateAtsInterview(ctx, operations.CreateAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2025-09-24T02:04:33.958Z"),
            ExternalEventXref: unifiedgosdk.Pointer("90e21303-e7ee-4b6e-93bc-29e148e6657e"),
            ID: unifiedgosdk.Pointer("c075d815-1a0c-4c73-b327-686782706e21"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            StartAt: types.MustNewTimeFromString("2025-05-19T22:35:24.880Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-04T20:06:11.434Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAtsInterviewRequest](../../pkg/models/operations/createatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateAtsInterviewResponse](../../pkg/models/operations/createatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsInterview

Retrieve an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsInterview" method="get" path="/ats/{connection_id}/interview/{id}" -->
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

    res, err := s.Interview.GetAtsInterview(ctx, operations.GetAtsInterviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAtsInterviewRequest](../../pkg/models/operations/getatsinterviewrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetAtsInterviewResponse](../../pkg/models/operations/getatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsInterviews

List all interviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsInterviews" method="get" path="/ats/{connection_id}/interview" -->
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

    res, err := s.Interview.ListAtsInterviews(ctx, operations.ListAtsInterviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsInterviewsRequest](../../pkg/models/operations/listatsinterviewsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsInterviewsResponse](../../pkg/models/operations/listatsinterviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsInterview

Update an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsInterview" method="patch" path="/ats/{connection_id}/interview/{id}" example="ats_interview" -->
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

    res, err := s.Interview.PatchAtsInterview(ctx, operations.PatchAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2025-09-24T02:04:33.963Z"),
            ExternalEventXref: unifiedgosdk.Pointer("ae365a27-4969-4b9e-aded-6612321a55f8"),
            ID: unifiedgosdk.Pointer("e0d6206e-7b81-4cf6-8eac-5493466b8b65"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            StartAt: types.MustNewTimeFromString("2025-05-19T22:35:24.885Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-04T20:06:11.440Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchAtsInterviewRequest](../../pkg/models/operations/patchatsinterviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchAtsInterviewResponse](../../pkg/models/operations/patchatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsInterview

Remove an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsInterview" method="delete" path="/ats/{connection_id}/interview/{id}" -->
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

    res, err := s.Interview.RemoveAtsInterview(ctx, operations.RemoveAtsInterviewRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveAtsInterviewRequest](../../pkg/models/operations/removeatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveAtsInterviewResponse](../../pkg/models/operations/removeatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsInterview

Update an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsInterview" method="put" path="/ats/{connection_id}/interview/{id}" example="ats_interview" -->
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

    res, err := s.Interview.UpdateAtsInterview(ctx, operations.UpdateAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2025-09-24T02:04:33.963Z"),
            ExternalEventXref: unifiedgosdk.Pointer("ae365a27-4969-4b9e-aded-6612321a55f8"),
            ID: unifiedgosdk.Pointer("e0d6206e-7b81-4cf6-8eac-5493466b8b65"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            StartAt: types.MustNewTimeFromString("2025-05-19T22:35:24.885Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-02-04T20:06:11.440Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAtsInterviewRequest](../../pkg/models/operations/updateatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAtsInterviewResponse](../../pkg/models/operations/updateatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |