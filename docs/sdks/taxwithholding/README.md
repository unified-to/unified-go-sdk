# Taxwithholding

## Overview

### Available Operations

* [CreateHrisTaxwithholding](#createhristaxwithholding) - Create a taxwithholding
* [GetHrisTaxwithholding](#gethristaxwithholding) - Retrieve a taxwithholding
* [ListHrisTaxwithholdings](#listhristaxwithholdings) - List all taxwithholdings
* [PatchHrisTaxwithholding](#patchhristaxwithholding) - Update a taxwithholding
* [RemoveHrisTaxwithholding](#removehristaxwithholding) - Remove a taxwithholding
* [UpdateHrisTaxwithholding](#updatehristaxwithholding) - Update a taxwithholding

## CreateHrisTaxwithholding

Creates a HrisTaxwithholding. Provide the unified HrisTaxwithholding fields in the request body; they are mapped to the provider's API. Returns the created HrisTaxwithholding.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTaxwithholding" method="post" path="/hris/{connection_id}/taxwithholding" example="hris_taxwithholding" -->
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

    res, err := s.Taxwithholding.CreateHrisTaxwithholding(ctx, operations.CreateHrisTaxwithholdingRequest{
        HrisTaxwithholding: shared.HrisTaxwithholding{
            Allowances: unifiedgosdk.Pointer[float64](1.0),
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            ExtraWithholdingAmount: unifiedgosdk.Pointer[float64](10.0),
            FilingStatus: shared.FilingStatusQualifyingSurvivor.ToPointer(),
            Form: shared.FormStateForm.ToPointer(),
            ID: unifiedgosdk.Pointer("e662e54a-99a8-4a50-8c8d-f5b72fb3d30a"),
            IsExempt: unifiedgosdk.Pointer(false),
            IsResident: unifiedgosdk.Pointer(true),
            Jurisdiction: unifiedgosdk.Pointer("MI"),
            JurisdictionType: shared.JurisdictionTypeState.ToPointer(),
            Metadata: []shared.HrisMetadata{},
            OverrideAmount: unifiedgosdk.Pointer[float64](19.0),
            OverrideType: shared.OverrideTypeFixedAmount.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-09T00:26:24.515Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxwithholding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateHrisTaxwithholdingRequest](../../pkg/models/operations/createhristaxwithholdingrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateHrisTaxwithholdingResponse](../../pkg/models/operations/createhristaxwithholdingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTaxwithholding

Retrieves a single HrisTaxwithholding by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTaxwithholding" method="get" path="/hris/{connection_id}/taxwithholding/{id}" -->
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

    res, err := s.Taxwithholding.GetHrisTaxwithholding(ctx, operations.GetHrisTaxwithholdingRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxwithholding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetHrisTaxwithholdingRequest](../../pkg/models/operations/gethristaxwithholdingrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetHrisTaxwithholdingResponse](../../pkg/models/operations/gethristaxwithholdingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTaxwithholdings

Returns a list of HrisTaxwithholdings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTaxwithholdings" method="get" path="/hris/{connection_id}/taxwithholding" -->
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

    res, err := s.Taxwithholding.ListHrisTaxwithholdings(ctx, operations.ListHrisTaxwithholdingsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxwithholdings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListHrisTaxwithholdingsRequest](../../pkg/models/operations/listhristaxwithholdingsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListHrisTaxwithholdingsResponse](../../pkg/models/operations/listhristaxwithholdingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTaxwithholding

Updates the HrisTaxwithholding identified by `id`. Provide the unified HrisTaxwithholding fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTaxwithholding" method="patch" path="/hris/{connection_id}/taxwithholding/{id}" example="hris_taxwithholding" -->
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

    res, err := s.Taxwithholding.PatchHrisTaxwithholding(ctx, operations.PatchHrisTaxwithholdingRequest{
        HrisTaxwithholding: shared.HrisTaxwithholding{
            Allowances: unifiedgosdk.Pointer[float64](1.0),
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            ExtraWithholdingAmount: unifiedgosdk.Pointer[float64](10.0),
            FilingStatus: shared.FilingStatusQualifyingSurvivor.ToPointer(),
            Form: shared.FormStateForm.ToPointer(),
            ID: unifiedgosdk.Pointer("281f870c-2096-4385-8883-5eb465e92086"),
            IsExempt: unifiedgosdk.Pointer(false),
            IsResident: unifiedgosdk.Pointer(true),
            Jurisdiction: unifiedgosdk.Pointer("MI"),
            JurisdictionType: shared.JurisdictionTypeState.ToPointer(),
            Metadata: []shared.HrisMetadata{},
            OverrideAmount: unifiedgosdk.Pointer[float64](19.0),
            OverrideType: shared.OverrideTypeFixedAmount.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-09T00:26:24.527Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxwithholding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchHrisTaxwithholdingRequest](../../pkg/models/operations/patchhristaxwithholdingrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PatchHrisTaxwithholdingResponse](../../pkg/models/operations/patchhristaxwithholdingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTaxwithholding

Deletes the HrisTaxwithholding identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTaxwithholding" method="delete" path="/hris/{connection_id}/taxwithholding/{id}" -->
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

    res, err := s.Taxwithholding.RemoveHrisTaxwithholding(ctx, operations.RemoveHrisTaxwithholdingRequest{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.RemoveHrisTaxwithholdingRequest](../../pkg/models/operations/removehristaxwithholdingrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.RemoveHrisTaxwithholdingResponse](../../pkg/models/operations/removehristaxwithholdingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTaxwithholding

Updates the HrisTaxwithholding identified by `id`. Provide the unified HrisTaxwithholding fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTaxwithholding" method="put" path="/hris/{connection_id}/taxwithholding/{id}" example="hris_taxwithholding" -->
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

    res, err := s.Taxwithholding.UpdateHrisTaxwithholding(ctx, operations.UpdateHrisTaxwithholdingRequest{
        HrisTaxwithholding: shared.HrisTaxwithholding{
            Allowances: unifiedgosdk.Pointer[float64](1.0),
            Country: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-08-14T05:49:35.141Z"),
            ExtraWithholdingAmount: unifiedgosdk.Pointer[float64](10.0),
            FilingStatus: shared.FilingStatusQualifyingSurvivor.ToPointer(),
            Form: shared.FormStateForm.ToPointer(),
            ID: unifiedgosdk.Pointer("281f870c-2096-4385-8883-5eb465e92086"),
            IsExempt: unifiedgosdk.Pointer(false),
            IsResident: unifiedgosdk.Pointer(true),
            Jurisdiction: unifiedgosdk.Pointer("MI"),
            JurisdictionType: shared.JurisdictionTypeState.ToPointer(),
            Metadata: []shared.HrisMetadata{},
            OverrideAmount: unifiedgosdk.Pointer[float64](19.0),
            OverrideType: shared.OverrideTypeFixedAmount.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-09T00:26:24.527Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxwithholding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateHrisTaxwithholdingRequest](../../pkg/models/operations/updatehristaxwithholdingrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateHrisTaxwithholdingResponse](../../pkg/models/operations/updatehristaxwithholdingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |