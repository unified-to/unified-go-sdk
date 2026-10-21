# Earning

## Overview

### Available Operations

* [CreateHrisEarning](#createhrisearning) - Create an earning
* [GetHrisEarning](#gethrisearning) - Retrieve an earning
* [ListHrisEarnings](#listhrisearnings) - List all earnings
* [PatchHrisEarning](#patchhrisearning) - Update an earning
* [RemoveHrisEarning](#removehrisearning) - Remove an earning
* [UpdateHrisEarning](#updatehrisearning) - Update an earning

## CreateHrisEarning

Create an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisEarning" method="post" path="/hris/{connection_id}/earning" example="hris_earning" -->
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

    res, err := s.Earning.CreateHrisEarning(ctx, operations.CreateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("e217da60-71ee-4135-afac-bd2548e026c2"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-25T11:48:48.385Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-08T20:32:52.338Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisEarningRequest](../../pkg/models/operations/createhrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisEarningResponse](../../pkg/models/operations/createhrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisEarning

Retrieve an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisEarning" method="get" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Earning.GetHrisEarning(ctx, operations.GetHrisEarningRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisEarningRequest](../../pkg/models/operations/gethrisearningrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisEarningResponse](../../pkg/models/operations/gethrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisEarnings

List all earnings

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisEarnings" method="get" path="/hris/{connection_id}/earning" -->
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

    res, err := s.Earning.ListHrisEarnings(ctx, operations.ListHrisEarningsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarnings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisEarningsRequest](../../pkg/models/operations/listhrisearningsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisEarningsResponse](../../pkg/models/operations/listhrisearningsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisEarning

Update an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisEarning" method="patch" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
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

    res, err := s.Earning.PatchHrisEarning(ctx, operations.PatchHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("ae74f559-3b04-412c-92c7-6b4867daa099"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-25T11:48:48.394Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-08T20:32:52.341Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisEarningRequest](../../pkg/models/operations/patchhrisearningrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisEarningResponse](../../pkg/models/operations/patchhrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisEarning

Remove an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisEarning" method="delete" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Earning.RemoveHrisEarning(ctx, operations.RemoveHrisEarningRequest{
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
| `request`                                                                                      | [operations.RemoveHrisEarningRequest](../../pkg/models/operations/removehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisEarningResponse](../../pkg/models/operations/removehrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisEarning

Update an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisEarning" method="put" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
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

    res, err := s.Earning.UpdateHrisEarning(ctx, operations.UpdateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("ae74f559-3b04-412c-92c7-6b4867daa099"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-25T11:48:48.394Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-08T20:32:52.341Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisEarningRequest](../../pkg/models/operations/updatehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisEarningResponse](../../pkg/models/operations/updatehrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |