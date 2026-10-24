# Call

## Overview

### Available Operations

* [CreateUcCall](#createuccall) - Create a call
* [GetUcCall](#getuccall) - Retrieve a call
* [ListUcCalls](#listuccalls) - List all calls
* [PatchUcCall](#patchuccall) - Update a call
* [UpdateUcCall](#updateuccall) - Update a call

## CreateUcCall

Creates a UcCall. Provide the unified UcCall fields in the request body; they are mapped to the provider's API. Returns the created UcCall.

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcCall" method="post" path="/uc/{connection_id}/call" example="uc_call" -->
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

    res, err := s.Call.CreateUcCall(ctx, operations.CreateUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-15T07:47:08.492Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("6c877acb-8697-44b1-944a-12fd65567f21"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-05-30T23:16:24.982Z"),
            ID: unifiedgosdk.Pointer("6c877acb-8697-44b1-944a-12fd65567f21"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("069faba1-04db-4721-a34a-648578987742"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-17T22:51:38.201Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-03T14:05:38.715Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateUcCallRequest](../../pkg/models/operations/createuccallrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateUcCallResponse](../../pkg/models/operations/createuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcCall

Retrieves a single UcCall by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcCall" method="get" path="/uc/{connection_id}/call/{id}" -->
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

    res, err := s.Call.GetUcCall(ctx, operations.GetUcCallRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetUcCallRequest](../../pkg/models/operations/getuccallrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetUcCallResponse](../../pkg/models/operations/getuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcCalls

Returns a list of UcCalls. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcCalls" method="get" path="/uc/{connection_id}/call" -->
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

    res, err := s.Call.ListUcCalls(ctx, operations.ListUcCallsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListUcCallsRequest](../../pkg/models/operations/listuccallsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListUcCallsResponse](../../pkg/models/operations/listuccallsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchUcCall

Updates the UcCall identified by `id`. Provide the unified UcCall fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcCall" method="patch" path="/uc/{connection_id}/call/{id}" example="uc_call" -->
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

    res, err := s.Call.PatchUcCall(ctx, operations.PatchUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-15T07:47:08.514Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("4b9e5c65-1b4c-4273-8851-a2ed95f0f15f"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-05-30T23:16:24.994Z"),
            ID: unifiedgosdk.Pointer("4b9e5c65-1b4c-4273-8851-a2ed95f0f15f"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("714962a5-bc60-4607-844f-ed609f8a3710"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-17T22:51:38.208Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-03T14:05:38.718Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchUcCallRequest](../../pkg/models/operations/patchuccallrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PatchUcCallResponse](../../pkg/models/operations/patchuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUcCall

Updates the UcCall identified by `id`. Provide the unified UcCall fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcCall" method="put" path="/uc/{connection_id}/call/{id}" example="uc_call" -->
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

    res, err := s.Call.UpdateUcCall(ctx, operations.UpdateUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-15T07:47:08.514Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("4b9e5c65-1b4c-4273-8851-a2ed95f0f15f"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-05-30T23:16:24.994Z"),
            ID: unifiedgosdk.Pointer("4b9e5c65-1b4c-4273-8851-a2ed95f0f15f"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("714962a5-bc60-4607-844f-ed609f8a3710"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-17T22:51:38.208Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-03T14:05:38.718Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateUcCallRequest](../../pkg/models/operations/updateuccallrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateUcCallResponse](../../pkg/models/operations/updateuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |