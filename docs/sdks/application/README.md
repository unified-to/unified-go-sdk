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
            AppliedAt: types.MustNewTimeFromString("2025-10-24T10:31:55.978Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-14T04:58:50.053Z"),
            ID: unifiedgosdk.Pointer("0a8b932c-7639-41c6-adad-3df2fa2ee6f4"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("0270e7ac-e401-4258-b875-29adf64e198f"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("arguo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Defessus occaecati adflicto caecus claustrum annus addo.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("caput"),
            OriginalSubstatus: unifiedgosdk.Pointer("admiratio"),
            RejectedAt: types.MustNewTimeFromString("2026-11-18T05:29:19.640Z"),
            RejectedReason: unifiedgosdk.Pointer("Talio credo vomica allatus stella."),
            Source: unifiedgosdk.Pointer("alii"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crebro nostrum alias sponte turbo sopor vergo surculus nam copiose. Enim casso cursus aestus deprecator. Despecto tabella argentum decretum cultellus aveho distinctio verecundia stella."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-25T07:24:11.610Z"),
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
            AppliedAt: types.MustNewTimeFromString("2025-10-24T10:31:55.990Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-14T04:58:50.068Z"),
            ID: unifiedgosdk.Pointer("aa6756ef-7c09-43fd-9393-35d55c912e7b"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c7b1ada9-e108-42cf-af64-dfe31a9b8eb1"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("arguo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Defessus occaecati adflicto caecus claustrum annus addo.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("caput"),
            OriginalSubstatus: unifiedgosdk.Pointer("admiratio"),
            RejectedAt: types.MustNewTimeFromString("2026-11-18T05:29:19.657Z"),
            RejectedReason: unifiedgosdk.Pointer("Talio credo vomica allatus stella."),
            Source: unifiedgosdk.Pointer("alii"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crebro nostrum alias sponte turbo sopor vergo surculus nam copiose. Enim casso cursus aestus deprecator. Despecto tabella argentum decretum cultellus aveho distinctio verecundia stella."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-25T07:24:11.628Z"),
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
            AppliedAt: types.MustNewTimeFromString("2025-10-24T10:31:55.990Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-06-14T04:58:50.068Z"),
            ID: unifiedgosdk.Pointer("aa6756ef-7c09-43fd-9393-35d55c912e7b"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c7b1ada9-e108-42cf-af64-dfe31a9b8eb1"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("arguo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Defessus occaecati adflicto caecus claustrum annus addo.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("caput"),
            OriginalSubstatus: unifiedgosdk.Pointer("admiratio"),
            RejectedAt: types.MustNewTimeFromString("2026-11-18T05:29:19.657Z"),
            RejectedReason: unifiedgosdk.Pointer("Talio credo vomica allatus stella."),
            Source: unifiedgosdk.Pointer("alii"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crebro nostrum alias sponte turbo sopor vergo surculus nam copiose. Enim casso cursus aestus deprecator. Despecto tabella argentum decretum cultellus aveho distinctio verecundia stella."),
            UpdatedAt: types.MustNewTimeFromString("2026-11-25T07:24:11.628Z"),
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