# Branch

## Overview

### Available Operations

* [CreateRepoBranch](#createrepobranch) - Create a branch
* [GetRepoBranch](#getrepobranch) - Retrieve a branch
* [ListRepoBranches](#listrepobranches) - List all branches
* [PatchRepoBranch](#patchrepobranch) - Update a branch
* [RemoveRepoBranch](#removerepobranch) - Remove a branch
* [UpdateRepoBranch](#updaterepobranch) - Update a branch

## CreateRepoBranch

Creates a RepoBranch. Provide the unified RepoBranch fields in the request body; they are mapped to the provider's API. Returns the created RepoBranch.

### Example Usage

<!-- UsageSnippet language="go" operationID="createRepoBranch" method="post" path="/repo/{connection_id}/branch" example="repo_branch" -->
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

    res, err := s.Branch.CreateRepoBranch(ctx, operations.CreateRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("2e0abdc1-3bc6-4576-86b2-498ceec9b4b3"),
            Metadata: []shared.RepoMetadata{},
            Name: "transmitter-calculate",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-07T10:44:37.774Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoBranch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateRepoBranchRequest](../../pkg/models/operations/createrepobranchrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateRepoBranchResponse](../../pkg/models/operations/createrepobranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRepoBranch

Retrieves a single RepoBranch by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRepoBranch" method="get" path="/repo/{connection_id}/branch/{id}" -->
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

    res, err := s.Branch.GetRepoBranch(ctx, operations.GetRepoBranchRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoBranch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetRepoBranchRequest](../../pkg/models/operations/getrepobranchrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetRepoBranchResponse](../../pkg/models/operations/getrepobranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRepoBranches

Returns a list of RepoBranches. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listRepoBranches" method="get" path="/repo/{connection_id}/branch" -->
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

    res, err := s.Branch.ListRepoBranches(ctx, operations.ListRepoBranchesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoBranches != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListRepoBranchesRequest](../../pkg/models/operations/listrepobranchesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListRepoBranchesResponse](../../pkg/models/operations/listrepobranchesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchRepoBranch

Updates the RepoBranch identified by `id`. Provide the unified RepoBranch fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRepoBranch" method="patch" path="/repo/{connection_id}/branch/{id}" example="repo_branch" -->
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

    res, err := s.Branch.PatchRepoBranch(ctx, operations.PatchRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("61096c53-c8e7-4fe5-b5b6-a2ff57e6c228"),
            Metadata: []shared.RepoMetadata{},
            Name: "transmitter-calculate",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-07T10:44:37.776Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoBranch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchRepoBranchRequest](../../pkg/models/operations/patchrepobranchrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchRepoBranchResponse](../../pkg/models/operations/patchrepobranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveRepoBranch

Deletes the RepoBranch identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeRepoBranch" method="delete" path="/repo/{connection_id}/branch/{id}" -->
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

    res, err := s.Branch.RemoveRepoBranch(ctx, operations.RemoveRepoBranchRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveRepoBranchRequest](../../pkg/models/operations/removerepobranchrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveRepoBranchResponse](../../pkg/models/operations/removerepobranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRepoBranch

Updates the RepoBranch identified by `id`. Provide the unified RepoBranch fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateRepoBranch" method="put" path="/repo/{connection_id}/branch/{id}" example="repo_branch" -->
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

    res, err := s.Branch.UpdateRepoBranch(ctx, operations.UpdateRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("61096c53-c8e7-4fe5-b5b6-a2ff57e6c228"),
            Metadata: []shared.RepoMetadata{},
            Name: "transmitter-calculate",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-07T10:44:37.776Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoBranch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateRepoBranchRequest](../../pkg/models/operations/updaterepobranchrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateRepoBranchResponse](../../pkg/models/operations/updaterepobranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |