# Commit

## Overview

### Available Operations

* [CreateRepoCommit](#createrepocommit) - Create a commit
* [GetRepoCommit](#getrepocommit) - Retrieve a commit
* [ListRepoCommits](#listrepocommits) - List all commits
* [PatchRepoCommit](#patchrepocommit) - Update a commit
* [RemoveRepoCommit](#removerepocommit) - Remove a commit
* [UpdateRepoCommit](#updaterepocommit) - Update a commit

## CreateRepoCommit

Create a commit

### Example Usage

<!-- UsageSnippet language="go" operationID="createRepoCommit" method="post" path="/repo/{connection_id}/commit" example="repo_commit" -->
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

    res, err := s.Commit.CreateRepoCommit(ctx, operations.CreateRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("4576cc80-721f-45a5-b425-dab15445761f"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-16T13:15:36.723Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoCommit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateRepoCommitRequest](../../pkg/models/operations/createrepocommitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateRepoCommitResponse](../../pkg/models/operations/createrepocommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRepoCommit

Retrieve a commit

### Example Usage

<!-- UsageSnippet language="go" operationID="getRepoCommit" method="get" path="/repo/{connection_id}/commit/{id}" -->
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

    res, err := s.Commit.GetRepoCommit(ctx, operations.GetRepoCommitRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoCommit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetRepoCommitRequest](../../pkg/models/operations/getrepocommitrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetRepoCommitResponse](../../pkg/models/operations/getrepocommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRepoCommits

List all commits

### Example Usage

<!-- UsageSnippet language="go" operationID="listRepoCommits" method="get" path="/repo/{connection_id}/commit" -->
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

    res, err := s.Commit.ListRepoCommits(ctx, operations.ListRepoCommitsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoCommits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListRepoCommitsRequest](../../pkg/models/operations/listrepocommitsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListRepoCommitsResponse](../../pkg/models/operations/listrepocommitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchRepoCommit

Update a commit

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRepoCommit" method="patch" path="/repo/{connection_id}/commit/{id}" example="repo_commit" -->
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

    res, err := s.Commit.PatchRepoCommit(ctx, operations.PatchRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("667d61f3-a0d4-4910-9d1d-c80d44c629dd"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-16T13:15:36.726Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoCommit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchRepoCommitRequest](../../pkg/models/operations/patchrepocommitrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchRepoCommitResponse](../../pkg/models/operations/patchrepocommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveRepoCommit

Remove a commit

### Example Usage

<!-- UsageSnippet language="go" operationID="removeRepoCommit" method="delete" path="/repo/{connection_id}/commit/{id}" -->
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

    res, err := s.Commit.RemoveRepoCommit(ctx, operations.RemoveRepoCommitRequest{
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
| `request`                                                                                    | [operations.RemoveRepoCommitRequest](../../pkg/models/operations/removerepocommitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveRepoCommitResponse](../../pkg/models/operations/removerepocommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRepoCommit

Update a commit

### Example Usage

<!-- UsageSnippet language="go" operationID="updateRepoCommit" method="put" path="/repo/{connection_id}/commit/{id}" example="repo_commit" -->
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

    res, err := s.Commit.UpdateRepoCommit(ctx, operations.UpdateRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("667d61f3-a0d4-4910-9d1d-c80d44c629dd"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-16T13:15:36.726Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoCommit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateRepoCommitRequest](../../pkg/models/operations/updaterepocommitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateRepoCommitResponse](../../pkg/models/operations/updaterepocommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |