# Label

## Overview

### Available Operations

* [CreateShippingLabel](#createshippinglabel) - Create a label
* [GetShippingLabel](#getshippinglabel) - Retrieve a label
* [ListShippingLabels](#listshippinglabels) - List all labels
* [PatchShippingLabel](#patchshippinglabel) - Update a label
* [RemoveShippingLabel](#removeshippinglabel) - Remove a label
* [UpdateShippingLabel](#updateshippinglabel) - Update a label

## CreateShippingLabel

Creates a ShippingLabel. Provide the unified ShippingLabel fields in the request body; they are mapped to the provider's API. Returns the created ShippingLabel.

### Example Usage

<!-- UsageSnippet language="go" operationID="createShippingLabel" method="post" path="/shipping/{connection_id}/label" example="shipping_label" -->
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

    res, err := s.Label.CreateShippingLabel(ctx, operations.CreateShippingLabelRequest{
        ShippingLabel: shared.ShippingLabel{
            CreatedAt: types.MustNewTimeFromString("2022-11-18T16:45:38.067Z"),
            ID: unifiedgosdk.Pointer("bee475f4-181b-44ae-a791-d2e0348600e2"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://labels.example.com/bee475f4-181b-44ae-a791-d2e0348600e2.pdf"),
            Metadata: []shared.ShippingMetadata{},
            ServiceCode: unifiedgosdk.Pointer("Dzy"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("ZYV60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-05-01T01:50:29.453Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingLabel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateShippingLabelRequest](../../pkg/models/operations/createshippinglabelrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateShippingLabelResponse](../../pkg/models/operations/createshippinglabelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetShippingLabel

Retrieves a single ShippingLabel by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getShippingLabel" method="get" path="/shipping/{connection_id}/label/{id}" -->
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

    res, err := s.Label.GetShippingLabel(ctx, operations.GetShippingLabelRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingLabel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetShippingLabelRequest](../../pkg/models/operations/getshippinglabelrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetShippingLabelResponse](../../pkg/models/operations/getshippinglabelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListShippingLabels

Returns a list of ShippingLabels. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listShippingLabels" method="get" path="/shipping/{connection_id}/label" -->
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

    res, err := s.Label.ListShippingLabels(ctx, operations.ListShippingLabelsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingLabels != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListShippingLabelsRequest](../../pkg/models/operations/listshippinglabelsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListShippingLabelsResponse](../../pkg/models/operations/listshippinglabelsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchShippingLabel

Updates the ShippingLabel identified by `id`. Provide the unified ShippingLabel fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchShippingLabel" method="patch" path="/shipping/{connection_id}/label/{id}" example="shipping_label" -->
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

    res, err := s.Label.PatchShippingLabel(ctx, operations.PatchShippingLabelRequest{
        ShippingLabel: shared.ShippingLabel{
            CreatedAt: types.MustNewTimeFromString("2022-11-18T16:45:38.067Z"),
            ID: unifiedgosdk.Pointer("c8826c5f-eba4-43ec-9674-66b4d523b64f"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://labels.example.com/c8826c5f-eba4-43ec-9674-66b4d523b64f.pdf"),
            Metadata: []shared.ShippingMetadata{},
            ServiceCode: unifiedgosdk.Pointer("Dzy"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("ZYV60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-05-01T01:50:29.458Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingLabel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchShippingLabelRequest](../../pkg/models/operations/patchshippinglabelrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchShippingLabelResponse](../../pkg/models/operations/patchshippinglabelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveShippingLabel

Deletes the ShippingLabel identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeShippingLabel" method="delete" path="/shipping/{connection_id}/label/{id}" -->
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

    res, err := s.Label.RemoveShippingLabel(ctx, operations.RemoveShippingLabelRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveShippingLabelRequest](../../pkg/models/operations/removeshippinglabelrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveShippingLabelResponse](../../pkg/models/operations/removeshippinglabelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateShippingLabel

Updates the ShippingLabel identified by `id`. Provide the unified ShippingLabel fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateShippingLabel" method="put" path="/shipping/{connection_id}/label/{id}" example="shipping_label" -->
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

    res, err := s.Label.UpdateShippingLabel(ctx, operations.UpdateShippingLabelRequest{
        ShippingLabel: shared.ShippingLabel{
            CreatedAt: types.MustNewTimeFromString("2022-11-18T16:45:38.067Z"),
            ID: unifiedgosdk.Pointer("c8826c5f-eba4-43ec-9674-66b4d523b64f"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://labels.example.com/c8826c5f-eba4-43ec-9674-66b4d523b64f.pdf"),
            Metadata: []shared.ShippingMetadata{},
            ServiceCode: unifiedgosdk.Pointer("Dzy"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("ZYV60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-05-01T01:50:29.458Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingLabel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateShippingLabelRequest](../../pkg/models/operations/updateshippinglabelrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateShippingLabelResponse](../../pkg/models/operations/updateshippinglabelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |