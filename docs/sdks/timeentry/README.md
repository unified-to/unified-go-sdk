# Timeentry

## Overview

### Available Operations

* [CreateLegalTimeentry](#createlegaltimeentry) - Create a timeentry
* [GetLegalTimeentry](#getlegaltimeentry) - Retrieve a timeentry
* [ListLegalTimeentries](#listlegaltimeentries) - List all timeentries
* [PatchLegalTimeentry](#patchlegaltimeentry) - Update a timeentry
* [RemoveLegalTimeentry](#removelegaltimeentry) - Remove a timeentry
* [UpdateLegalTimeentry](#updatelegaltimeentry) - Update a timeentry

## CreateLegalTimeentry

Create a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalTimeentry" method="post" path="/legal/{connection_id}/timeentry" example="legal_timeentry" -->
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

    res, err := s.Timeentry.CreateLegalTimeentry(ctx, operations.CreateLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("68fc0fa4-3720-4d3e-8348-aaf9a6f5f953"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T19:04:58.728Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateLegalTimeentryRequest](../../pkg/models/operations/createlegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateLegalTimeentryResponse](../../pkg/models/operations/createlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalTimeentry

Retrieve a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalTimeentry" method="get" path="/legal/{connection_id}/timeentry/{id}" -->
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

    res, err := s.Timeentry.GetLegalTimeentry(ctx, operations.GetLegalTimeentryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetLegalTimeentryRequest](../../pkg/models/operations/getlegaltimeentryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetLegalTimeentryResponse](../../pkg/models/operations/getlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalTimeentries

List all timeentries

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalTimeentries" method="get" path="/legal/{connection_id}/timeentry" -->
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

    res, err := s.Timeentry.ListLegalTimeentries(ctx, operations.ListLegalTimeentriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListLegalTimeentriesRequest](../../pkg/models/operations/listlegaltimeentriesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListLegalTimeentriesResponse](../../pkg/models/operations/listlegaltimeentriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLegalTimeentry

Update a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalTimeentry" method="patch" path="/legal/{connection_id}/timeentry/{id}" example="legal_timeentry" -->
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

    res, err := s.Timeentry.PatchLegalTimeentry(ctx, operations.PatchLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("2ae7d5c4-9da1-4880-9d35-5d8406fb0a4e"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T19:04:58.729Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchLegalTimeentryRequest](../../pkg/models/operations/patchlegaltimeentryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchLegalTimeentryResponse](../../pkg/models/operations/patchlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLegalTimeentry

Remove a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalTimeentry" method="delete" path="/legal/{connection_id}/timeentry/{id}" -->
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

    res, err := s.Timeentry.RemoveLegalTimeentry(ctx, operations.RemoveLegalTimeentryRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveLegalTimeentryRequest](../../pkg/models/operations/removelegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveLegalTimeentryResponse](../../pkg/models/operations/removelegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLegalTimeentry

Update a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalTimeentry" method="put" path="/legal/{connection_id}/timeentry/{id}" example="legal_timeentry" -->
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

    res, err := s.Timeentry.UpdateLegalTimeentry(ctx, operations.UpdateLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("2ae7d5c4-9da1-4880-9d35-5d8406fb0a4e"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T19:04:58.729Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateLegalTimeentryRequest](../../pkg/models/operations/updatelegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateLegalTimeentryResponse](../../pkg/models/operations/updatelegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |