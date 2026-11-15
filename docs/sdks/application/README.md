# Application

## Overview

### Available Operations

* [CreateAtsApplication](#createatsapplication) - Create an application
* [GetAtsApplication](#getatsapplication) - Retrieve an application
* [ListAtsApplications](#listatsapplications) - List all applications
* [PatchAtsApplication](#patchatsapplication) - Update an application
* [RemoveAtsApplication](#removeatsapplication) - Remove an application
* [UpdateAtsApplication](#updateatsapplication) - Update an application

## CreateAtsApplication

Creates a AtsApplication. Provide the unified AtsApplication fields in the request body; they are mapped to the provider's API. Returns the created AtsApplication.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsApplication" method="post" path="/ats/{connection_id}/application" example="ats_application" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Application.CreateAtsApplication(ctx, operations.CreateAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-10-16T00:29:02.676Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-03T03:13:01.246Z"),
            ID: unifiedgosdk.Pointer("2fb513c4-b5a3-4b35-a1d9-80bab574e614"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("39dcfc07-6c76-4896-85c5-2c976de4f793"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-11-05T08:44:38.628Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-12T08:43:14.035Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateAtsApplicationRequest](../../pkg/models/operations/createatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateAtsApplicationResponse](../../pkg/models/operations/createatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsApplication

Retrieves a single AtsApplication by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsApplication" method="get" path="/ats/{connection_id}/application/{id}" -->
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

    res, err := s.Application.GetAtsApplication(ctx, operations.GetAtsApplicationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAtsApplicationRequest](../../pkg/models/operations/getatsapplicationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetAtsApplicationResponse](../../pkg/models/operations/getatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsApplications

Returns a list of AtsApplications. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsApplications" method="get" path="/ats/{connection_id}/application" -->
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

    res, err := s.Application.ListAtsApplications(ctx, operations.ListAtsApplicationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplications != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAtsApplicationsRequest](../../pkg/models/operations/listatsapplicationsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAtsApplicationsResponse](../../pkg/models/operations/listatsapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsApplication

Updates the AtsApplication identified by `id`. Provide the unified AtsApplication fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsApplication" method="patch" path="/ats/{connection_id}/application/{id}" example="ats_application" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Application.PatchAtsApplication(ctx, operations.PatchAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-10-16T00:29:02.692Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-03T03:13:01.268Z"),
            ID: unifiedgosdk.Pointer("337db61f-124f-4b49-8c0a-0ce0187c565a"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("dbca9030-6710-41d8-92b0-d66911da9ea6"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-11-05T08:44:38.653Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-12T08:43:14.060Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchAtsApplicationRequest](../../pkg/models/operations/patchatsapplicationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchAtsApplicationResponse](../../pkg/models/operations/patchatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsApplication

Deletes the AtsApplication identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsApplication" method="delete" path="/ats/{connection_id}/application/{id}" -->
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

    res, err := s.Application.RemoveAtsApplication(ctx, operations.RemoveAtsApplicationRequest{
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
| `request`                                                                                            | [operations.RemoveAtsApplicationRequest](../../pkg/models/operations/removeatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveAtsApplicationResponse](../../pkg/models/operations/removeatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsApplication

Updates the AtsApplication identified by `id`. Provide the unified AtsApplication fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsApplication" method="put" path="/ats/{connection_id}/application/{id}" example="ats_application" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Application.UpdateAtsApplication(ctx, operations.UpdateAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-10-16T00:29:02.692Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-03T03:13:01.268Z"),
            ID: unifiedgosdk.Pointer("337db61f-124f-4b49-8c0a-0ce0187c565a"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("dbca9030-6710-41d8-92b0-d66911da9ea6"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-11-05T08:44:38.653Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-12T08:43:14.060Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateAtsApplicationRequest](../../pkg/models/operations/updateatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateAtsApplicationResponse](../../pkg/models/operations/updateatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |