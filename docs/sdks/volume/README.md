# Volume

## Overview

### Available Operations

* [CreateComputeVolume](#createcomputevolume) - Create a volume
* [GetComputeVolume](#getcomputevolume) - Retrieve a volume
* [ListComputeVolumes](#listcomputevolumes) - List all volumes
* [PatchComputeVolume](#patchcomputevolume) - Update a volume
* [RemoveComputeVolume](#removecomputevolume) - Remove a volume
* [UpdateComputeVolume](#updatecomputevolume) - Update a volume

## CreateComputeVolume

Create a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeVolume" method="post" path="/compute/{connection_id}/volume" example="compute_volume" -->
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

    res, err := s.Volume.CreateComputeVolume(ctx, operations.CreateComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("85caf030-6b1a-4799-b44c-0422dc7582b7"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:46:33.349Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateComputeVolumeRequest](../../pkg/models/operations/createcomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateComputeVolumeResponse](../../pkg/models/operations/createcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeVolume

Retrieve a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeVolume" method="get" path="/compute/{connection_id}/volume/{id}" -->
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

    res, err := s.Volume.GetComputeVolume(ctx, operations.GetComputeVolumeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeVolumeRequest](../../pkg/models/operations/getcomputevolumerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeVolumeResponse](../../pkg/models/operations/getcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeVolumes

List all volumes

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeVolumes" method="get" path="/compute/{connection_id}/volume" -->
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

    res, err := s.Volume.ListComputeVolumes(ctx, operations.ListComputeVolumesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolumes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeVolumesRequest](../../pkg/models/operations/listcomputevolumesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeVolumesResponse](../../pkg/models/operations/listcomputevolumesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeVolume

Update a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeVolume" method="patch" path="/compute/{connection_id}/volume/{id}" example="compute_volume" -->
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

    res, err := s.Volume.PatchComputeVolume(ctx, operations.PatchComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("aa0459b7-a5b0-43ac-b9ed-fc85a3bb23b1"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:46:33.355Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchComputeVolumeRequest](../../pkg/models/operations/patchcomputevolumerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchComputeVolumeResponse](../../pkg/models/operations/patchcomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeVolume

Remove a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeVolume" method="delete" path="/compute/{connection_id}/volume/{id}" -->
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

    res, err := s.Volume.RemoveComputeVolume(ctx, operations.RemoveComputeVolumeRequest{
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
| `request`                                                                                          | [operations.RemoveComputeVolumeRequest](../../pkg/models/operations/removecomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveComputeVolumeResponse](../../pkg/models/operations/removecomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeVolume

Update a volume

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeVolume" method="put" path="/compute/{connection_id}/volume/{id}" example="compute_volume" -->
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

    res, err := s.Volume.UpdateComputeVolume(ctx, operations.UpdateComputeVolumeRequest{
        ComputeVolume: shared.ComputeVolume{
            CreatedAt: types.MustNewTimeFromString("2022-11-15T15:12:48.055Z"),
            ID: unifiedgosdk.Pointer("aa0459b7-a5b0-43ac-b9ed-fc85a3bb23b1"),
            InstanceIds: []string{},
            Metadata: []shared.ComputeMetadata{},
            Name: unifiedgosdk.Pointer("vol-solum"),
            SizeMb: unifiedgosdk.Pointer[float64](1024000.0),
            Status: shared.ComputeVolumeStatusOther.ToPointer(),
            Type: shared.ComputeVolumeTypeHighPerformance.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:46:33.355Z"),
            UsedMb: unifiedgosdk.Pointer[float64](524986.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeVolume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateComputeVolumeRequest](../../pkg/models/operations/updatecomputevolumerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateComputeVolumeResponse](../../pkg/models/operations/updatecomputevolumeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |