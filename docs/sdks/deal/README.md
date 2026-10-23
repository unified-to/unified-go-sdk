# Deal

## Overview

### Available Operations

* [CreateCrmDeal](#createcrmdeal) - Create a deal
* [GetCrmDeal](#getcrmdeal) - Retrieve a deal
* [ListCrmDeals](#listcrmdeals) - List all deals
* [PatchCrmDeal](#patchcrmdeal) - Update a deal
* [RemoveCrmDeal](#removecrmdeal) - Remove a deal
* [UpdateCrmDeal](#updatecrmdeal) - Update a deal

## CreateCrmDeal

Creates a CrmDeal. Provide the unified CrmDeal fields in the request body; they are mapped to the provider's API. Returns the created CrmDeal.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmDeal" method="post" path="/crm/{connection_id}/deal" example="crm_deal" -->
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

    res, err := s.Deal.CreateCrmDeal(ctx, operations.CreateCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T20:16:21.884Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-01T20:43:10.436Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("66a2a92a-7ec6-4623-b1f6-2aabedba0c82"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7c5696af-7e3c-4bac-b66e-3ef5bff2bbb9"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-12T16:08:30.562Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateCrmDealRequest](../../pkg/models/operations/createcrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateCrmDealResponse](../../pkg/models/operations/createcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmDeal

Retrieves a single CrmDeal by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmDeal" method="get" path="/crm/{connection_id}/deal/{id}" -->
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

    res, err := s.Deal.GetCrmDeal(ctx, operations.GetCrmDealRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetCrmDealRequest](../../pkg/models/operations/getcrmdealrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetCrmDealResponse](../../pkg/models/operations/getcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmDeals

Returns a list of CrmDeals. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmDeals" method="get" path="/crm/{connection_id}/deal" -->
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

    res, err := s.Deal.ListCrmDeals(ctx, operations.ListCrmDealsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmDeals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCrmDealsRequest](../../pkg/models/operations/listcrmdealsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCrmDealsResponse](../../pkg/models/operations/listcrmdealsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmDeal

Updates the CrmDeal identified by `id`. Provide the unified CrmDeal fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmDeal" method="patch" path="/crm/{connection_id}/deal/{id}" example="crm_deal" -->
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

    res, err := s.Deal.PatchCrmDeal(ctx, operations.PatchCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T20:16:21.888Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-01T20:43:10.450Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("18a24c2c-2642-4b1c-8a0e-aec4bc016b46"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("79053c92-1e0e-4570-be7c-18d310f77d8f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-12T16:08:30.570Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmDealRequest](../../pkg/models/operations/patchcrmdealrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchCrmDealResponse](../../pkg/models/operations/patchcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmDeal

Deletes the CrmDeal identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmDeal" method="delete" path="/crm/{connection_id}/deal/{id}" -->
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

    res, err := s.Deal.RemoveCrmDeal(ctx, operations.RemoveCrmDealRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmDealRequest](../../pkg/models/operations/removecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveCrmDealResponse](../../pkg/models/operations/removecrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmDeal

Updates the CrmDeal identified by `id`. Provide the unified CrmDeal fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmDeal" method="put" path="/crm/{connection_id}/deal/{id}" example="crm_deal" -->
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

    res, err := s.Deal.UpdateCrmDeal(ctx, operations.UpdateCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T20:16:21.888Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-01T20:43:10.450Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("18a24c2c-2642-4b1c-8a0e-aec4bc016b46"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("79053c92-1e0e-4570-be7c-18d310f77d8f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-12T16:08:30.570Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmDealRequest](../../pkg/models/operations/updatecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateCrmDealResponse](../../pkg/models/operations/updatecrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |