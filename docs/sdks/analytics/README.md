# Analytics

## Overview

### Available Operations

* [CreateAnalyticsEvent](#createanalyticsevent) - Create an event
* [CreateAnalyticsProperty](#createanalyticsproperty) - Create a property
* [CreateAnalyticsVisitor](#createanalyticsvisitor) - Create a visitor
* [GetAnalyticsEvent](#getanalyticsevent) - Retrieve an event
* [GetAnalyticsProperty](#getanalyticsproperty) - Retrieve a property
* [GetAnalyticsSession](#getanalyticssession) - Retrieve a session
* [GetAnalyticsVisitor](#getanalyticsvisitor) - Retrieve a visitor
* [ListAnalyticsEvents](#listanalyticsevents) - List all events
* [ListAnalyticsProperties](#listanalyticsproperties) - List all properties
* [ListAnalyticsReports](#listanalyticsreports) - List all reports
* [ListAnalyticsSessions](#listanalyticssessions) - List all sessions
* [ListAnalyticsVisitors](#listanalyticsvisitors) - List all visitors
* [PatchAnalyticsProperty](#patchanalyticsproperty) - Update a property
* [PatchAnalyticsVisitor](#patchanalyticsvisitor) - Update a visitor
* [RemoveAnalyticsProperty](#removeanalyticsproperty) - Remove a property
* [RemoveAnalyticsVisitor](#removeanalyticsvisitor) - Remove a visitor
* [UpdateAnalyticsProperty](#updateanalyticsproperty) - Update a property
* [UpdateAnalyticsVisitor](#updateanalyticsvisitor) - Update a visitor

## CreateAnalyticsEvent

Create an event

### Example Usage

<!-- UsageSnippet language="go" operationID="createAnalyticsEvent" method="post" path="/analytics/{connection_id}/event" example="analytics_event" -->
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

    res, err := s.Analytics.CreateAnalyticsEvent(ctx, operations.CreateAnalyticsEventRequest{
        AnalyticsEvent: shared.AnalyticsEvent{
            CreatedAt: types.MustNewTimeFromString("2023-06-21T03:13:22.954Z"),
            EventType: shared.EventTypeScreenView.ToPointer(),
            ID: unifiedgosdk.Pointer("e18d88af-2f06-4a98-955b-d191b15011c5"),
            Metadata: map[string]shared.PropertyAnalyticsEventMetadata{
                "key": shared.PropertyAnalyticsEventMetadata{},
            },
            Name: unifiedgosdk.Pointer("Xk707ttsb51v"),
            UpdatedAt: types.MustNewTimeFromString("2023-09-24T10:48:15.319Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateAnalyticsEventRequest](../../pkg/models/operations/createanalyticseventrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateAnalyticsEventResponse](../../pkg/models/operations/createanalyticseventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAnalyticsProperty

Create a property

### Example Usage

<!-- UsageSnippet language="go" operationID="createAnalyticsProperty" method="post" path="/analytics/{connection_id}/property" example="analytics_property" -->
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

    res, err := s.Analytics.CreateAnalyticsProperty(ctx, operations.CreateAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("96f37af2-09c8-4762-9d7d-3758ad8bbdb0"),
            Name: unifiedgosdk.Pointer("Daniel, Goldner and Dickinson"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-14T20:05:39.058Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAnalyticsPropertyRequest](../../pkg/models/operations/createanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAnalyticsPropertyResponse](../../pkg/models/operations/createanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAnalyticsVisitor

Create a visitor

### Example Usage

<!-- UsageSnippet language="go" operationID="createAnalyticsVisitor" method="post" path="/analytics/{connection_id}/visitor" example="analytics_visitor" -->
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

    res, err := s.Analytics.CreateAnalyticsVisitor(ctx, operations.CreateAnalyticsVisitorRequest{
        AnalyticsVisitor: shared.AnalyticsVisitor{
            CreatedAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            Email: unifiedgosdk.Pointer("Dallas_Mitchell@yahoo.com"),
            FirstSeenAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            ID: unifiedgosdk.Pointer("9ee04127-35e6-4f5c-808a-a4daab459cab"),
            LastSeenAt: types.MustNewTimeFromString("2021-12-12T09:39:23.604Z"),
            Metadata: map[string]shared.PropertyAnalyticsVisitorMetadata{
                "segment": shared.PropertyAnalyticsVisitorMetadata{},
            },
            Name: unifiedgosdk.Pointer("Desiree O'Hara"),
            TotalEvents: unifiedgosdk.Pointer[float64](3639.0),
            UpdatedAt: types.MustNewTimeFromString("2025-06-27T08:25:40.361Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsVisitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateAnalyticsVisitorRequest](../../pkg/models/operations/createanalyticsvisitorrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateAnalyticsVisitorResponse](../../pkg/models/operations/createanalyticsvisitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAnalyticsEvent

Retrieve an event

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsEvent" method="get" path="/analytics/{connection_id}/event/{id}" -->
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

    res, err := s.Analytics.GetAnalyticsEvent(ctx, operations.GetAnalyticsEventRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAnalyticsEventRequest](../../pkg/models/operations/getanalyticseventrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetAnalyticsEventResponse](../../pkg/models/operations/getanalyticseventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAnalyticsProperty

Retrieve a property

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsProperty" method="get" path="/analytics/{connection_id}/property/{id}" -->
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

    res, err := s.Analytics.GetAnalyticsProperty(ctx, operations.GetAnalyticsPropertyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAnalyticsPropertyRequest](../../pkg/models/operations/getanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAnalyticsPropertyResponse](../../pkg/models/operations/getanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAnalyticsSession

Retrieve a session

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsSession" method="get" path="/analytics/{connection_id}/session/{id}" -->
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

    res, err := s.Analytics.GetAnalyticsSession(ctx, operations.GetAnalyticsSessionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAnalyticsSessionRequest](../../pkg/models/operations/getanalyticssessionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetAnalyticsSessionResponse](../../pkg/models/operations/getanalyticssessionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAnalyticsVisitor

Retrieve a visitor

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsVisitor" method="get" path="/analytics/{connection_id}/visitor/{id}" -->
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

    res, err := s.Analytics.GetAnalyticsVisitor(ctx, operations.GetAnalyticsVisitorRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsVisitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAnalyticsVisitorRequest](../../pkg/models/operations/getanalyticsvisitorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetAnalyticsVisitorResponse](../../pkg/models/operations/getanalyticsvisitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsEvents

List all events

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsEvents" method="get" path="/analytics/{connection_id}/event" -->
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

    res, err := s.Analytics.ListAnalyticsEvents(ctx, operations.ListAnalyticsEventsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsEvents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAnalyticsEventsRequest](../../pkg/models/operations/listanalyticseventsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAnalyticsEventsResponse](../../pkg/models/operations/listanalyticseventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsProperties

List all properties

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsProperties" method="get" path="/analytics/{connection_id}/property" -->
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

    res, err := s.Analytics.ListAnalyticsProperties(ctx, operations.ListAnalyticsPropertiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperties != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListAnalyticsPropertiesRequest](../../pkg/models/operations/listanalyticspropertiesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListAnalyticsPropertiesResponse](../../pkg/models/operations/listanalyticspropertiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsReports

List all reports

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsReports" method="get" path="/analytics/{connection_id}/report" -->
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

    res, err := s.Analytics.ListAnalyticsReports(ctx, operations.ListAnalyticsReportsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsReports != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAnalyticsReportsRequest](../../pkg/models/operations/listanalyticsreportsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAnalyticsReportsResponse](../../pkg/models/operations/listanalyticsreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsSessions

List all sessions

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsSessions" method="get" path="/analytics/{connection_id}/session" -->
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

    res, err := s.Analytics.ListAnalyticsSessions(ctx, operations.ListAnalyticsSessionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsSessions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListAnalyticsSessionsRequest](../../pkg/models/operations/listanalyticssessionsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListAnalyticsSessionsResponse](../../pkg/models/operations/listanalyticssessionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAnalyticsVisitors

List all visitors

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsVisitors" method="get" path="/analytics/{connection_id}/visitor" -->
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

    res, err := s.Analytics.ListAnalyticsVisitors(ctx, operations.ListAnalyticsVisitorsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsVisitors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListAnalyticsVisitorsRequest](../../pkg/models/operations/listanalyticsvisitorsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListAnalyticsVisitorsResponse](../../pkg/models/operations/listanalyticsvisitorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAnalyticsProperty

Update a property

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAnalyticsProperty" method="patch" path="/analytics/{connection_id}/property/{id}" example="analytics_property" -->
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

    res, err := s.Analytics.PatchAnalyticsProperty(ctx, operations.PatchAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("3295786e-46d1-4b17-9f6a-7dfb69a599cf"),
            Name: unifiedgosdk.Pointer("Daniel, Goldner and Dickinson"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-14T20:05:39.058Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAnalyticsPropertyRequest](../../pkg/models/operations/patchanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAnalyticsPropertyResponse](../../pkg/models/operations/patchanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAnalyticsVisitor

Update a visitor

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAnalyticsVisitor" method="patch" path="/analytics/{connection_id}/visitor/{id}" example="analytics_visitor" -->
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

    res, err := s.Analytics.PatchAnalyticsVisitor(ctx, operations.PatchAnalyticsVisitorRequest{
        AnalyticsVisitor: shared.AnalyticsVisitor{
            CreatedAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            Email: unifiedgosdk.Pointer("Dallas_Mitchell@yahoo.com"),
            FirstSeenAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            ID: unifiedgosdk.Pointer("b7fa8ffd-53ca-4848-9834-0ba67ad4b5a6"),
            LastSeenAt: types.MustNewTimeFromString("2021-12-12T09:39:23.607Z"),
            Metadata: map[string]shared.PropertyAnalyticsVisitorMetadata{
                "segment": shared.PropertyAnalyticsVisitorMetadata{},
            },
            Name: unifiedgosdk.Pointer("Desiree O'Hara"),
            TotalEvents: unifiedgosdk.Pointer[float64](3639.0),
            UpdatedAt: types.MustNewTimeFromString("2025-06-27T08:25:40.368Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsVisitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchAnalyticsVisitorRequest](../../pkg/models/operations/patchanalyticsvisitorrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchAnalyticsVisitorResponse](../../pkg/models/operations/patchanalyticsvisitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAnalyticsProperty

Remove a property

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAnalyticsProperty" method="delete" path="/analytics/{connection_id}/property/{id}" -->
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

    res, err := s.Analytics.RemoveAnalyticsProperty(ctx, operations.RemoveAnalyticsPropertyRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAnalyticsPropertyRequest](../../pkg/models/operations/removeanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAnalyticsPropertyResponse](../../pkg/models/operations/removeanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAnalyticsVisitor

Remove a visitor

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAnalyticsVisitor" method="delete" path="/analytics/{connection_id}/visitor/{id}" -->
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

    res, err := s.Analytics.RemoveAnalyticsVisitor(ctx, operations.RemoveAnalyticsVisitorRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveAnalyticsVisitorRequest](../../pkg/models/operations/removeanalyticsvisitorrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveAnalyticsVisitorResponse](../../pkg/models/operations/removeanalyticsvisitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAnalyticsProperty

Update a property

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAnalyticsProperty" method="put" path="/analytics/{connection_id}/property/{id}" example="analytics_property" -->
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

    res, err := s.Analytics.UpdateAnalyticsProperty(ctx, operations.UpdateAnalyticsPropertyRequest{
        AnalyticsProperty: shared.AnalyticsProperty{
            CreatedAt: types.MustNewTimeFromString("2021-09-05T19:04:58.430Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("3295786e-46d1-4b17-9f6a-7dfb69a599cf"),
            Name: unifiedgosdk.Pointer("Daniel, Goldner and Dickinson"),
            Timezone: unifiedgosdk.Pointer("UTC"),
            UpdatedAt: types.MustNewTimeFromString("2021-09-14T20:05:39.058Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsProperty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAnalyticsPropertyRequest](../../pkg/models/operations/updateanalyticspropertyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAnalyticsPropertyResponse](../../pkg/models/operations/updateanalyticspropertyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAnalyticsVisitor

Update a visitor

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAnalyticsVisitor" method="put" path="/analytics/{connection_id}/visitor/{id}" example="analytics_visitor" -->
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

    res, err := s.Analytics.UpdateAnalyticsVisitor(ctx, operations.UpdateAnalyticsVisitorRequest{
        AnalyticsVisitor: shared.AnalyticsVisitor{
            CreatedAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            Email: unifiedgosdk.Pointer("Dallas_Mitchell@yahoo.com"),
            FirstSeenAt: types.MustNewTimeFromString("2020-04-16T20:29:48.281Z"),
            ID: unifiedgosdk.Pointer("b7fa8ffd-53ca-4848-9834-0ba67ad4b5a6"),
            LastSeenAt: types.MustNewTimeFromString("2021-12-12T09:39:23.607Z"),
            Metadata: map[string]shared.PropertyAnalyticsVisitorMetadata{
                "segment": shared.PropertyAnalyticsVisitorMetadata{},
            },
            Name: unifiedgosdk.Pointer("Desiree O'Hara"),
            TotalEvents: unifiedgosdk.Pointer[float64](3639.0),
            UpdatedAt: types.MustNewTimeFromString("2025-06-27T08:25:40.368Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsVisitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAnalyticsVisitorRequest](../../pkg/models/operations/updateanalyticsvisitorrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateAnalyticsVisitorResponse](../../pkg/models/operations/updateanalyticsvisitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |