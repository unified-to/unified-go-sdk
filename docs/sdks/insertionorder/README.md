# Insertionorder

## Overview

### Available Operations

* [CreateAdsInsertionorder](#createadsinsertionorder) - Create an insertionorder
* [GetAdsInsertionorder](#getadsinsertionorder) - Retrieve an insertionorder
* [ListAdsInsertionorders](#listadsinsertionorders) - List all insertionorders
* [PatchAdsInsertionorder](#patchadsinsertionorder) - Update an insertionorder
* [RemoveAdsInsertionorder](#removeadsinsertionorder) - Remove an insertionorder
* [UpdateAdsInsertionorder](#updateadsinsertionorder) - Update an insertionorder

## CreateAdsInsertionorder

Create an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsInsertionorder" method="post" path="/ads/{connection_id}/insertionorder" example="ads_insertionorder" -->
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

    res, err := s.Insertionorder.CreateAdsInsertionorder(ctx, operations.CreateAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("0025631a-c198-4691-9215-01f54de73cae"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAdsInsertionorderRequest](../../pkg/models/operations/createadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAdsInsertionorderResponse](../../pkg/models/operations/createadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsInsertionorder

Retrieve an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsInsertionorder" method="get" path="/ads/{connection_id}/insertionorder/{id}" -->
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

    res, err := s.Insertionorder.GetAdsInsertionorder(ctx, operations.GetAdsInsertionorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAdsInsertionorderRequest](../../pkg/models/operations/getadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAdsInsertionorderResponse](../../pkg/models/operations/getadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsInsertionorders

List all insertionorders

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsInsertionorders" method="get" path="/ads/{connection_id}/insertionorder" -->
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

    res, err := s.Insertionorder.ListAdsInsertionorders(ctx, operations.ListAdsInsertionordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAdsInsertionordersRequest](../../pkg/models/operations/listadsinsertionordersrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAdsInsertionordersResponse](../../pkg/models/operations/listadsinsertionordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsInsertionorder

Update an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsInsertionorder" method="patch" path="/ads/{connection_id}/insertionorder/{id}" example="ads_insertionorder" -->
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

    res, err := s.Insertionorder.PatchAdsInsertionorder(ctx, operations.PatchAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("658ad6f8-cf35-434b-9652-34e030c35b64"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAdsInsertionorderRequest](../../pkg/models/operations/patchadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAdsInsertionorderResponse](../../pkg/models/operations/patchadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsInsertionorder

Remove an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsInsertionorder" method="delete" path="/ads/{connection_id}/insertionorder/{id}" -->
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

    res, err := s.Insertionorder.RemoveAdsInsertionorder(ctx, operations.RemoveAdsInsertionorderRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAdsInsertionorderRequest](../../pkg/models/operations/removeadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAdsInsertionorderResponse](../../pkg/models/operations/removeadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsInsertionorder

Update an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsInsertionorder" method="put" path="/ads/{connection_id}/insertionorder/{id}" example="ads_insertionorder" -->
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

    res, err := s.Insertionorder.UpdateAdsInsertionorder(ctx, operations.UpdateAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("658ad6f8-cf35-434b-9652-34e030c35b64"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAdsInsertionorderRequest](../../pkg/models/operations/updateadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAdsInsertionorderResponse](../../pkg/models/operations/updateadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |