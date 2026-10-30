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

Creates a HrisEarning. Provide the unified HrisEarning fields in the request body; they are mapped to the provider's API. Returns the created HrisEarning.

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
            ID: unifiedgosdk.Pointer("79de15af-e2f5-46f7-8e5f-dc7117ff9507"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-29T11:26:48.701Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-10T00:54:12.139Z"),
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

Retrieves a single HrisEarning by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of HrisEarnings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the HrisEarning identified by `id`. Provide the unified HrisEarning fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("9c7f0e04-f6ef-4354-b9a7-904ef5e7dc45"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-29T11:26:48.710Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-10T00:54:12.141Z"),
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

Deletes the HrisEarning identified by `id`.

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

Updates the HrisEarning identified by `id`. Provide the unified HrisEarning fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("9c7f0e04-f6ef-4354-b9a7-904ef5e7dc45"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-29T11:26:48.710Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-10T00:54:12.141Z"),
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