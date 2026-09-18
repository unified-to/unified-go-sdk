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

Create a label

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
            ID: unifiedgosdk.Pointer("d64a026b-8f62-4b07-8bce-ac44e8999e22"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://optimal-meadow.net"),
            ServiceCode: unifiedgosdk.Pointer("GIz"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("zYv60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-16T18:42:43.800Z"),
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

Retrieve a label

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

List all labels

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

Update a label

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
            ID: unifiedgosdk.Pointer("4e75db1f-3d4a-4b8f-aafd-71d770747559"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://optimal-meadow.net"),
            ServiceCode: unifiedgosdk.Pointer("GIz"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("zYv60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-16T18:42:43.803Z"),
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

Remove a label

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

Update a label

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
            ID: unifiedgosdk.Pointer("4e75db1f-3d4a-4b8f-aafd-71d770747559"),
            IsVoided: unifiedgosdk.Pointer(false),
            LabelCost: unifiedgosdk.Pointer[float64](40.83653403213248),
            LabelCostCurrency: unifiedgosdk.Pointer("USD"),
            LabelFormat: shared.LabelFormatPng.ToPointer(),
            LabelURL: unifiedgosdk.Pointer("https://optimal-meadow.net"),
            ServiceCode: unifiedgosdk.Pointer("GIz"),
            Status: shared.ShippingLabelStatusException.ToPointer(),
            TrackingNumber: unifiedgosdk.Pointer("zYv60FOIBUJ6"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-16T18:42:43.803Z"),
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