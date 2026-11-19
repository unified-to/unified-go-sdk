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
            AppliedAt: types.MustNewTimeFromString("2025-10-18T12:36:16.023Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-06T10:17:46.353Z"),
            ID: unifiedgosdk.Pointer("5f1164eb-1f59-4af2-9c43-936d878ee00c"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c3e9bf7a-e048-4cd1-ae2b-20d24a42abd4"),
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
            RejectedAt: types.MustNewTimeFromString("2026-11-09T04:36:45.030Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-16T05:09:56.347Z"),
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
            AppliedAt: types.MustNewTimeFromString("2025-10-18T12:36:16.042Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-06T10:17:46.379Z"),
            ID: unifiedgosdk.Pointer("e05728c3-b2ca-475a-90ba-3f56096ae28a"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("4195a12d-5db7-45c4-824e-a881a3bb3e6a"),
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
            RejectedAt: types.MustNewTimeFromString("2026-11-09T04:36:45.060Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-16T05:09:56.377Z"),
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
            AppliedAt: types.MustNewTimeFromString("2025-10-18T12:36:16.042Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-06T10:17:46.379Z"),
            ID: unifiedgosdk.Pointer("e05728c3-b2ca-475a-90ba-3f56096ae28a"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("4195a12d-5db7-45c4-824e-a881a3bb3e6a"),
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
            RejectedAt: types.MustNewTimeFromString("2026-11-09T04:36:45.060Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-16T05:09:56.377Z"),
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