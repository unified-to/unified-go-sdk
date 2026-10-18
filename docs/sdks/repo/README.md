# Repo

## Overview

### Available Operations

* [CreateRepoBranch](#createrepobranch) - Create a branch
* [CreateRepoCommit](#createrepocommit) - Create a commit
* [CreateRepoOrganization](#createrepoorganization) - Create an organization
* [CreateRepoPullrequest](#createrepopullrequest) - Create a pullrequest
* [CreateRepoRepository](#createreporepository) - Create a repository
* [GetRepoBranch](#getrepobranch) - Retrieve a branch
* [GetRepoCommit](#getrepocommit) - Retrieve a commit
* [GetRepoOrganization](#getrepoorganization) - Retrieve an organization
* [GetRepoPullrequest](#getrepopullrequest) - Retrieve a pullrequest
* [GetRepoRepository](#getreporepository) - Retrieve a repository
* [ListRepoBranches](#listrepobranches) - List all branches
* [ListRepoCommits](#listrepocommits) - List all commits
* [ListRepoOrganizations](#listrepoorganizations) - List all organizations
* [ListRepoPullrequests](#listrepopullrequests) - List all pullrequests
* [ListRepoRepositories](#listreporepositories) - List all repositories
* [PatchRepoBranch](#patchrepobranch) - Update a branch
* [PatchRepoCommit](#patchrepocommit) - Update a commit
* [PatchRepoOrganization](#patchrepoorganization) - Update an organization
* [PatchRepoPullrequest](#patchrepopullrequest) - Update a pullrequest
* [PatchRepoRepository](#patchreporepository) - Update a repository
* [RemoveRepoBranch](#removerepobranch) - Remove a branch
* [RemoveRepoCommit](#removerepocommit) - Remove a commit
* [RemoveRepoOrganization](#removerepoorganization) - Remove an organization
* [RemoveRepoPullrequest](#removerepopullrequest) - Remove a pullrequest
* [RemoveRepoRepository](#removereporepository) - Remove a repository
* [UpdateRepoBranch](#updaterepobranch) - Update a branch
* [UpdateRepoCommit](#updaterepocommit) - Update a commit
* [UpdateRepoOrganization](#updaterepoorganization) - Update an organization
* [UpdateRepoPullrequest](#updaterepopullrequest) - Update a pullrequest
* [UpdateRepoRepository](#updatereporepository) - Update a repository

## CreateRepoBranch

Create a branch

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

    res, err := s.Repo.CreateRepoBranch(ctx, operations.CreateRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("d2abb4b7-098b-44bb-9a35-23c49ef36839"),
            Name: "voluptas",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-06T23:59:08.053Z"),
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

    res, err := s.Repo.CreateRepoCommit(ctx, operations.CreateRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("fd767343-7ba7-44c8-b054-cc57cdccd546"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-29T22:11:03.089Z"),
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

## CreateRepoOrganization

Create an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createRepoOrganization" method="post" path="/repo/{connection_id}/organization" example="repo_organization" -->
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

    res, err := s.Repo.CreateRepoOrganization(ctx, operations.CreateRepoOrganizationRequest{
        RepoOrganization: shared.RepoOrganization{
            AvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/fGl6Lb/3157/3173"),
            CreatedAt: types.MustNewTimeFromString("2022-07-07T00:18:40.748Z"),
            Description: unifiedgosdk.Pointer("Trepide defendo supra testimonium ager."),
            ID: unifiedgosdk.Pointer("3096cffd-9859-40e6-b178-20903357edef"),
            Name: unifiedgosdk.Pointer("Denesik - Lemke"),
            UpdatedAt: types.MustNewTimeFromString("2023-08-20T20:30:03.755Z"),
            WebURL: unifiedgosdk.Pointer("https://turbulent-overheard.biz"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateRepoOrganizationRequest](../../pkg/models/operations/createrepoorganizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateRepoOrganizationResponse](../../pkg/models/operations/createrepoorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateRepoPullrequest

Create a pullrequest

### Example Usage

<!-- UsageSnippet language="go" operationID="createRepoPullrequest" method="post" path="/repo/{connection_id}/pullrequest" example="repo_pullrequest" -->
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

    res, err := s.Repo.CreateRepoPullrequest(ctx, operations.CreateRepoPullrequestRequest{
        RepoPullrequest: shared.RepoPullrequest{
            ClosedAt: types.MustNewTimeFromString("2025-04-29T19:27:47.405Z"),
            CreatedAt: types.MustNewTimeFromString("2023-02-27T09:37:13.663Z"),
            ID: unifiedgosdk.Pointer("0db18505-3d26-4159-96c7-ffb73cd99c3a"),
            Labels: []string{
                "adhuc",
                "quaerat",
            },
            Notes: unifiedgosdk.Pointer("Coadunatio turbo curtus ceno consuasor aggero. Suggero adeo creptio tutamen vulnus aqua delicate adopto derelinquo caritas. Maiores vulgivagus succurro temporibus."),
            SourceBranchID: unifiedgosdk.Pointer("microchip-navigate"),
            Status: shared.RepoPullrequestStatusRejected.ToPointer(),
            TargetBranchID: unifiedgosdk.Pointer("feed-reboot"),
            Title: unifiedgosdk.Pointer("Cunae aegrus averto texo advoco bibo amet asporto."),
            UpdatedAt: types.MustNewTimeFromString("2025-01-15T13:07:16.043Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoPullrequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateRepoPullrequestRequest](../../pkg/models/operations/createrepopullrequestrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateRepoPullrequestResponse](../../pkg/models/operations/createrepopullrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateRepoRepository

Create a repository

### Example Usage

<!-- UsageSnippet language="go" operationID="createRepoRepository" method="post" path="/repo/{connection_id}/repository" example="repo_repository" -->
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

    res, err := s.Repo.CreateRepoRepository(ctx, operations.CreateRepoRepositoryRequest{
        RepoRepository: shared.RepoRepository{
            CreatedAt: types.MustNewTimeFromString("2023-06-12T09:42:00.080Z"),
            Description: unifiedgosdk.Pointer("Tribuo torqueo aetas ustulo illum."),
            ID: unifiedgosdk.Pointer("9d1a0961-db27-4d60-92df-0cbec099c5d7"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Name: "suggero",
            Owner: unifiedgosdk.Pointer("Marcella Kuhic"),
            UpdatedAt: types.MustNewTimeFromString("2024-02-29T10:51:11.675Z"),
            WebURL: unifiedgosdk.Pointer("https://brown-phrase.info"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoRepository != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateRepoRepositoryRequest](../../pkg/models/operations/createreporepositoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateRepoRepositoryResponse](../../pkg/models/operations/createreporepositoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRepoBranch

Retrieve a branch

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

    res, err := s.Repo.GetRepoBranch(ctx, operations.GetRepoBranchRequest{
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

    res, err := s.Repo.GetRepoCommit(ctx, operations.GetRepoCommitRequest{
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

## GetRepoOrganization

Retrieve an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getRepoOrganization" method="get" path="/repo/{connection_id}/organization/{id}" -->
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

    res, err := s.Repo.GetRepoOrganization(ctx, operations.GetRepoOrganizationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetRepoOrganizationRequest](../../pkg/models/operations/getrepoorganizationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetRepoOrganizationResponse](../../pkg/models/operations/getrepoorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRepoPullrequest

Retrieve a pullrequest

### Example Usage

<!-- UsageSnippet language="go" operationID="getRepoPullrequest" method="get" path="/repo/{connection_id}/pullrequest/{id}" -->
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

    res, err := s.Repo.GetRepoPullrequest(ctx, operations.GetRepoPullrequestRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoPullrequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetRepoPullrequestRequest](../../pkg/models/operations/getrepopullrequestrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetRepoPullrequestResponse](../../pkg/models/operations/getrepopullrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRepoRepository

Retrieve a repository

### Example Usage

<!-- UsageSnippet language="go" operationID="getRepoRepository" method="get" path="/repo/{connection_id}/repository/{id}" -->
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

    res, err := s.Repo.GetRepoRepository(ctx, operations.GetRepoRepositoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoRepository != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetRepoRepositoryRequest](../../pkg/models/operations/getreporepositoryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetRepoRepositoryResponse](../../pkg/models/operations/getreporepositoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRepoBranches

List all branches

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

    res, err := s.Repo.ListRepoBranches(ctx, operations.ListRepoBranchesRequest{
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

    res, err := s.Repo.ListRepoCommits(ctx, operations.ListRepoCommitsRequest{
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

## ListRepoOrganizations

List all organizations

### Example Usage

<!-- UsageSnippet language="go" operationID="listRepoOrganizations" method="get" path="/repo/{connection_id}/organization" -->
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

    res, err := s.Repo.ListRepoOrganizations(ctx, operations.ListRepoOrganizationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoOrganizations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListRepoOrganizationsRequest](../../pkg/models/operations/listrepoorganizationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListRepoOrganizationsResponse](../../pkg/models/operations/listrepoorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRepoPullrequests

List all pullrequests

### Example Usage

<!-- UsageSnippet language="go" operationID="listRepoPullrequests" method="get" path="/repo/{connection_id}/pullrequest" -->
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

    res, err := s.Repo.ListRepoPullrequests(ctx, operations.ListRepoPullrequestsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoPullrequests != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListRepoPullrequestsRequest](../../pkg/models/operations/listrepopullrequestsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListRepoPullrequestsResponse](../../pkg/models/operations/listrepopullrequestsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRepoRepositories

List all repositories

### Example Usage

<!-- UsageSnippet language="go" operationID="listRepoRepositories" method="get" path="/repo/{connection_id}/repository" -->
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

    res, err := s.Repo.ListRepoRepositories(ctx, operations.ListRepoRepositoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoRepositories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListRepoRepositoriesRequest](../../pkg/models/operations/listreporepositoriesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListRepoRepositoriesResponse](../../pkg/models/operations/listreporepositoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchRepoBranch

Update a branch

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

    res, err := s.Repo.PatchRepoBranch(ctx, operations.PatchRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("f1936069-ce59-4eff-8f16-57e654e95600"),
            Name: "voluptas",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-06T23:59:08.054Z"),
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

    res, err := s.Repo.PatchRepoCommit(ctx, operations.PatchRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("617054fa-c131-43e7-b42c-b86f2605e660"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-29T22:11:03.091Z"),
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

## PatchRepoOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRepoOrganization" method="patch" path="/repo/{connection_id}/organization/{id}" example="repo_organization" -->
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

    res, err := s.Repo.PatchRepoOrganization(ctx, operations.PatchRepoOrganizationRequest{
        RepoOrganization: shared.RepoOrganization{
            AvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/fGl6Lb/3157/3173"),
            CreatedAt: types.MustNewTimeFromString("2022-07-07T00:18:40.748Z"),
            Description: unifiedgosdk.Pointer("Trepide defendo supra testimonium ager."),
            ID: unifiedgosdk.Pointer("2d9032b5-49ff-4043-b62c-2fa76fe27b94"),
            Name: unifiedgosdk.Pointer("Denesik - Lemke"),
            UpdatedAt: types.MustNewTimeFromString("2023-08-20T20:30:03.757Z"),
            WebURL: unifiedgosdk.Pointer("https://turbulent-overheard.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchRepoOrganizationRequest](../../pkg/models/operations/patchrepoorganizationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchRepoOrganizationResponse](../../pkg/models/operations/patchrepoorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchRepoPullrequest

Update a pullrequest

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRepoPullrequest" method="patch" path="/repo/{connection_id}/pullrequest/{id}" example="repo_pullrequest" -->
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

    res, err := s.Repo.PatchRepoPullrequest(ctx, operations.PatchRepoPullrequestRequest{
        RepoPullrequest: shared.RepoPullrequest{
            ClosedAt: types.MustNewTimeFromString("2025-04-29T19:27:47.412Z"),
            CreatedAt: types.MustNewTimeFromString("2023-02-27T09:37:13.663Z"),
            ID: unifiedgosdk.Pointer("d548b202-a11d-49c4-807b-33622b0d0b5b"),
            Labels: []string{
                "adhuc",
                "quaerat",
            },
            Notes: unifiedgosdk.Pointer("Coadunatio turbo curtus ceno consuasor aggero. Suggero adeo creptio tutamen vulnus aqua delicate adopto derelinquo caritas. Maiores vulgivagus succurro temporibus."),
            SourceBranchID: unifiedgosdk.Pointer("microchip-navigate"),
            Status: shared.RepoPullrequestStatusRejected.ToPointer(),
            TargetBranchID: unifiedgosdk.Pointer("feed-reboot"),
            Title: unifiedgosdk.Pointer("Cunae aegrus averto texo advoco bibo amet asporto."),
            UpdatedAt: types.MustNewTimeFromString("2025-01-15T13:07:16.050Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoPullrequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchRepoPullrequestRequest](../../pkg/models/operations/patchrepopullrequestrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchRepoPullrequestResponse](../../pkg/models/operations/patchrepopullrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchRepoRepository

Update a repository

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRepoRepository" method="patch" path="/repo/{connection_id}/repository/{id}" example="repo_repository" -->
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

    res, err := s.Repo.PatchRepoRepository(ctx, operations.PatchRepoRepositoryRequest{
        RepoRepository: shared.RepoRepository{
            CreatedAt: types.MustNewTimeFromString("2023-06-12T09:42:00.080Z"),
            Description: unifiedgosdk.Pointer("Tribuo torqueo aetas ustulo illum."),
            ID: unifiedgosdk.Pointer("ba14dfc9-e813-4580-b0fe-4e38cc7a51f4"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Name: "suggero",
            Owner: unifiedgosdk.Pointer("Marcella Kuhic"),
            UpdatedAt: types.MustNewTimeFromString("2024-02-29T10:51:11.677Z"),
            WebURL: unifiedgosdk.Pointer("https://brown-phrase.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoRepository != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchRepoRepositoryRequest](../../pkg/models/operations/patchreporepositoryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchRepoRepositoryResponse](../../pkg/models/operations/patchreporepositoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveRepoBranch

Remove a branch

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

    res, err := s.Repo.RemoveRepoBranch(ctx, operations.RemoveRepoBranchRequest{
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

    res, err := s.Repo.RemoveRepoCommit(ctx, operations.RemoveRepoCommitRequest{
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

## RemoveRepoOrganization

Remove an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="removeRepoOrganization" method="delete" path="/repo/{connection_id}/organization/{id}" -->
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

    res, err := s.Repo.RemoveRepoOrganization(ctx, operations.RemoveRepoOrganizationRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveRepoOrganizationRequest](../../pkg/models/operations/removerepoorganizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveRepoOrganizationResponse](../../pkg/models/operations/removerepoorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveRepoPullrequest

Remove a pullrequest

### Example Usage

<!-- UsageSnippet language="go" operationID="removeRepoPullrequest" method="delete" path="/repo/{connection_id}/pullrequest/{id}" -->
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

    res, err := s.Repo.RemoveRepoPullrequest(ctx, operations.RemoveRepoPullrequestRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveRepoPullrequestRequest](../../pkg/models/operations/removerepopullrequestrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveRepoPullrequestResponse](../../pkg/models/operations/removerepopullrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveRepoRepository

Remove a repository

### Example Usage

<!-- UsageSnippet language="go" operationID="removeRepoRepository" method="delete" path="/repo/{connection_id}/repository/{id}" -->
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

    res, err := s.Repo.RemoveRepoRepository(ctx, operations.RemoveRepoRepositoryRequest{
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
| `request`                                                                                            | [operations.RemoveRepoRepositoryRequest](../../pkg/models/operations/removereporepositoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveRepoRepositoryResponse](../../pkg/models/operations/removereporepositoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRepoBranch

Update a branch

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

    res, err := s.Repo.UpdateRepoBranch(ctx, operations.UpdateRepoBranchRequest{
        RepoBranch: shared.RepoBranch{
            CreatedAt: types.MustNewTimeFromString("2019-02-24T09:33:48.540Z"),
            ID: unifiedgosdk.Pointer("f1936069-ce59-4eff-8f16-57e654e95600"),
            Name: "voluptas",
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2020-01-06T23:59:08.054Z"),
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

    res, err := s.Repo.UpdateRepoCommit(ctx, operations.UpdateRepoCommitRequest{
        RepoCommit: shared.RepoCommit{
            CreatedAt: types.MustNewTimeFromString("2020-07-12T16:20:42.520Z"),
            ID: unifiedgosdk.Pointer("617054fa-c131-43e7-b42c-b86f2605e660"),
            LinesAdded: unifiedgosdk.Pointer[float64](313.0),
            LinesChanged: unifiedgosdk.Pointer[float64](659.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](482.0),
            Message: unifiedgosdk.Pointer("Auctus ascisco esse attollo clarus odio tum bis rerum."),
            RepoID: "<id>",
            UpdatedAt: types.MustNewTimeFromString("2023-05-29T22:11:03.091Z"),
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

## UpdateRepoOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="updateRepoOrganization" method="put" path="/repo/{connection_id}/organization/{id}" example="repo_organization" -->
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

    res, err := s.Repo.UpdateRepoOrganization(ctx, operations.UpdateRepoOrganizationRequest{
        RepoOrganization: shared.RepoOrganization{
            AvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/fGl6Lb/3157/3173"),
            CreatedAt: types.MustNewTimeFromString("2022-07-07T00:18:40.748Z"),
            Description: unifiedgosdk.Pointer("Trepide defendo supra testimonium ager."),
            ID: unifiedgosdk.Pointer("2d9032b5-49ff-4043-b62c-2fa76fe27b94"),
            Name: unifiedgosdk.Pointer("Denesik - Lemke"),
            UpdatedAt: types.MustNewTimeFromString("2023-08-20T20:30:03.757Z"),
            WebURL: unifiedgosdk.Pointer("https://turbulent-overheard.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateRepoOrganizationRequest](../../pkg/models/operations/updaterepoorganizationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateRepoOrganizationResponse](../../pkg/models/operations/updaterepoorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRepoPullrequest

Update a pullrequest

### Example Usage

<!-- UsageSnippet language="go" operationID="updateRepoPullrequest" method="put" path="/repo/{connection_id}/pullrequest/{id}" example="repo_pullrequest" -->
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

    res, err := s.Repo.UpdateRepoPullrequest(ctx, operations.UpdateRepoPullrequestRequest{
        RepoPullrequest: shared.RepoPullrequest{
            ClosedAt: types.MustNewTimeFromString("2025-04-29T19:27:47.412Z"),
            CreatedAt: types.MustNewTimeFromString("2023-02-27T09:37:13.663Z"),
            ID: unifiedgosdk.Pointer("d548b202-a11d-49c4-807b-33622b0d0b5b"),
            Labels: []string{
                "adhuc",
                "quaerat",
            },
            Notes: unifiedgosdk.Pointer("Coadunatio turbo curtus ceno consuasor aggero. Suggero adeo creptio tutamen vulnus aqua delicate adopto derelinquo caritas. Maiores vulgivagus succurro temporibus."),
            SourceBranchID: unifiedgosdk.Pointer("microchip-navigate"),
            Status: shared.RepoPullrequestStatusRejected.ToPointer(),
            TargetBranchID: unifiedgosdk.Pointer("feed-reboot"),
            Title: unifiedgosdk.Pointer("Cunae aegrus averto texo advoco bibo amet asporto."),
            UpdatedAt: types.MustNewTimeFromString("2025-01-15T13:07:16.050Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoPullrequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateRepoPullrequestRequest](../../pkg/models/operations/updaterepopullrequestrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateRepoPullrequestResponse](../../pkg/models/operations/updaterepopullrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRepoRepository

Update a repository

### Example Usage

<!-- UsageSnippet language="go" operationID="updateRepoRepository" method="put" path="/repo/{connection_id}/repository/{id}" example="repo_repository" -->
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

    res, err := s.Repo.UpdateRepoRepository(ctx, operations.UpdateRepoRepositoryRequest{
        RepoRepository: shared.RepoRepository{
            CreatedAt: types.MustNewTimeFromString("2023-06-12T09:42:00.080Z"),
            Description: unifiedgosdk.Pointer("Tribuo torqueo aetas ustulo illum."),
            ID: unifiedgosdk.Pointer("ba14dfc9-e813-4580-b0fe-4e38cc7a51f4"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Name: "suggero",
            Owner: unifiedgosdk.Pointer("Marcella Kuhic"),
            UpdatedAt: types.MustNewTimeFromString("2024-02-29T10:51:11.677Z"),
            WebURL: unifiedgosdk.Pointer("https://brown-phrase.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepoRepository != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateRepoRepositoryRequest](../../pkg/models/operations/updatereporepositoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateRepoRepositoryResponse](../../pkg/models/operations/updatereporepositoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |