# Envvar

## Overview

### Available Operations

* [CreateComputeEnvvar](#createcomputeenvvar) - Create an envvar
* [GetComputeEnvvar](#getcomputeenvvar) - Retrieve an envvar
* [ListComputeEnvvars](#listcomputeenvvars) - List all envvars
* [PatchComputeEnvvar](#patchcomputeenvvar) - Update an envvar
* [RemoveComputeEnvvar](#removecomputeenvvar) - Remove an envvar
* [UpdateComputeEnvvar](#updatecomputeenvvar) - Update an envvar

## CreateComputeEnvvar

Create an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="createComputeEnvvar" method="post" path="/compute/{connection_id}/envvar" example="compute_envvar" -->
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

    res, err := s.Envvar.CreateComputeEnvvar(ctx, operations.CreateComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("98f895be-6740-4b15-8020-bac2e65f6b95"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T18:05:49.045Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateComputeEnvvarRequest](../../pkg/models/operations/createcomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateComputeEnvvarResponse](../../pkg/models/operations/createcomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComputeEnvvar

Retrieve an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="getComputeEnvvar" method="get" path="/compute/{connection_id}/envvar/{id}" -->
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

    res, err := s.Envvar.GetComputeEnvvar(ctx, operations.GetComputeEnvvarRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetComputeEnvvarRequest](../../pkg/models/operations/getcomputeenvvarrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetComputeEnvvarResponse](../../pkg/models/operations/getcomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComputeEnvvars

List all envvars

### Example Usage

<!-- UsageSnippet language="go" operationID="listComputeEnvvars" method="get" path="/compute/{connection_id}/envvar" -->
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

    res, err := s.Envvar.ListComputeEnvvars(ctx, operations.ListComputeEnvvarsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListComputeEnvvarsRequest](../../pkg/models/operations/listcomputeenvvarsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListComputeEnvvarsResponse](../../pkg/models/operations/listcomputeenvvarsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchComputeEnvvar

Update an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="patchComputeEnvvar" method="patch" path="/compute/{connection_id}/envvar/{id}" example="compute_envvar" -->
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

    res, err := s.Envvar.PatchComputeEnvvar(ctx, operations.PatchComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("1dc7e191-68bd-4478-bacf-ad4bee42fcbb"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T18:05:49.049Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchComputeEnvvarRequest](../../pkg/models/operations/patchcomputeenvvarrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchComputeEnvvarResponse](../../pkg/models/operations/patchcomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveComputeEnvvar

Remove an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="removeComputeEnvvar" method="delete" path="/compute/{connection_id}/envvar/{id}" -->
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

    res, err := s.Envvar.RemoveComputeEnvvar(ctx, operations.RemoveComputeEnvvarRequest{
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
| `request`                                                                                          | [operations.RemoveComputeEnvvarRequest](../../pkg/models/operations/removecomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveComputeEnvvarResponse](../../pkg/models/operations/removecomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComputeEnvvar

Update an envvar

### Example Usage

<!-- UsageSnippet language="go" operationID="updateComputeEnvvar" method="put" path="/compute/{connection_id}/envvar/{id}" example="compute_envvar" -->
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

    res, err := s.Envvar.UpdateComputeEnvvar(ctx, operations.UpdateComputeEnvvarRequest{
        ComputeEnvvar: shared.ComputeEnvvar{
            CreatedAt: types.MustNewTimeFromString("2021-08-16T07:52:27.306Z"),
            Environments: []shared.PropertyComputeEnvvarEnvironments{
                shared.PropertyComputeEnvvarEnvironmentsPreview,
            },
            ID: unifiedgosdk.Pointer("1dc7e191-68bd-4478-bacf-ad4bee42fcbb"),
            IsSecret: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("AEGER_SUBSTANTIA"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T18:05:49.049Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ComputeEnvvar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateComputeEnvvarRequest](../../pkg/models/operations/updatecomputeenvvarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateComputeEnvvarResponse](../../pkg/models/operations/updatecomputeenvvarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |