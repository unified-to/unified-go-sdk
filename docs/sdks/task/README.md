# Task

## Overview

### Available Operations

* [CreateGenaiTask](#creategenaitask) - Create a task
* [CreateTaskComment](#createtaskcomment) - Create a comment
* [CreateTaskProject](#createtaskproject) - Create a project
* [CreateTaskTask](#createtasktask) - Create a task
* [GetGenaiTask](#getgenaitask) - Retrieve a task
* [GetTaskChange](#gettaskchange) - Retrieve a change
* [GetTaskComment](#gettaskcomment) - Retrieve a comment
* [GetTaskProject](#gettaskproject) - Retrieve a project
* [GetTaskTask](#gettasktask) - Retrieve a task
* [ListGenaiTasks](#listgenaitasks) - List all tasks
* [ListTaskChanges](#listtaskchanges) - List all changes
* [ListTaskComments](#listtaskcomments) - List all comments
* [ListTaskProjects](#listtaskprojects) - List all projects
* [ListTaskTasks](#listtasktasks) - List all tasks
* [PatchGenaiTask](#patchgenaitask) - Update a task
* [PatchTaskComment](#patchtaskcomment) - Update a comment
* [PatchTaskProject](#patchtaskproject) - Update a project
* [PatchTaskTask](#patchtasktask) - Update a task
* [RemoveGenaiTask](#removegenaitask) - Remove a task
* [RemoveTaskComment](#removetaskcomment) - Remove a comment
* [RemoveTaskProject](#removetaskproject) - Remove a project
* [RemoveTaskTask](#removetasktask) - Remove a task
* [UpdateGenaiTask](#updategenaitask) - Update a task
* [UpdateTaskComment](#updatetaskcomment) - Update a comment
* [UpdateTaskProject](#updatetaskproject) - Update a project
* [UpdateTaskTask](#updatetasktask) - Update a task

## CreateGenaiTask

Creates a GenaiTask. Provide the unified GenaiTask fields in the request body; they are mapped to the provider's API. Returns the created GenaiTask.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGenaiTask" method="post" path="/genai/{connection_id}/task" example="genai_task" -->
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

    res, err := s.Task.CreateGenaiTask(ctx, operations.CreateGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-10-02T02:15:37.668Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("442fcbae-d067-4f9c-9db3-a70e9c393a98"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-22T00:40:23.675Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-26T16:31:59.677Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateGenaiTaskRequest](../../pkg/models/operations/creategenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateGenaiTaskResponse](../../pkg/models/operations/creategenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTaskComment

Creates a TaskComment. Provide the unified TaskComment fields in the request body; they are mapped to the provider's API. Returns the created TaskComment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTaskComment" method="post" path="/task/{connection_id}/comment" example="task_comment" -->
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

    res, err := s.Task.CreateTaskComment(ctx, operations.CreateTaskCommentRequest{
        TaskComment: shared.TaskComment{
            CreatedAt: types.MustNewTimeFromString("2019-10-12T20:33:37.879Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("0443a83e-d435-44b8-887a-9cbb133ec389"),
            Metadata: []shared.TaskMetadata{},
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-10-03T14:40:44.794Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateTaskCommentRequest](../../pkg/models/operations/createtaskcommentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateTaskCommentResponse](../../pkg/models/operations/createtaskcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTaskProject

Creates a TaskProject. Provide the unified TaskProject fields in the request body; they are mapped to the provider's API. Returns the created TaskProject.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTaskProject" method="post" path="/task/{connection_id}/project" example="task_project" -->
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

    res, err := s.Task.CreateTaskProject(ctx, operations.CreateTaskProjectRequest{
        TaskProject: shared.TaskProject{
            CreatedAt: types.MustNewTimeFromString("2023-06-23T16:39:40.446Z"),
            Description: unifiedgosdk.Pointer("Valetudo aggredior accommodo curiositas vox."),
            HasChildren: unifiedgosdk.Pointer(false),
            HasTasks: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("586e2b04-721f-4fb4-b1c9-ff0295b8f983"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("ac6d7296-95fc-4c6c-b2b1-b41a06651184"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("decens"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "uterque",
                    )),
                },
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("9bfa2064-4758-497f-942c-b3cc0ea19dfd"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-11T16:39:32.559Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateTaskProjectRequest](../../pkg/models/operations/createtaskprojectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateTaskProjectResponse](../../pkg/models/operations/createtaskprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTaskTask

Creates a TaskTask. Provide the unified TaskTask fields in the request body; they are mapped to the provider's API. Returns the created TaskTask.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTaskTask" method="post" path="/task/{connection_id}/task" example="task_task" -->
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

    res, err := s.Task.CreateTaskTask(ctx, operations.CreateTaskTaskRequest{
        TaskTask: shared.TaskTask{
            AttachmentIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2025-09-12T14:15:13.461Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("6f667320-6d4d-44b6-b6db-61853b30438f"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("I'll compress the optical IB interface, that should capacitor the JBOD system!"),
            Notes: unifiedgosdk.Pointer("Argentum ventosus rem talis neque thymbra. Absconditus peccatus crustulum tollo cometes volva vacuus eos cedo. Somnus utpote coadunatio denuncio adopto autus sono atrocitas. Concido rerum cura aveho torqueo tempore tubineus vomer tempore suscipit. Clibanus valetudo vulnero pauper caput absorbeo correptius vir deduco."),
            Priority: unifiedgosdk.Pointer("MEDIUM"),
            Progress: unifiedgosdk.Pointer[float64](33.0),
            StartAt: types.MustNewTimeFromString("2025-08-15T04:23:44.499Z"),
            Status: shared.TaskTaskStatusOpened.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "curis",
                "crinis",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](5359.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("magni"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-07T11:59:37.144Z"),
            URL: unifiedgosdk.Pointer("https://elementary-utilization.net/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateTaskTaskRequest](../../pkg/models/operations/createtasktaskrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateTaskTaskResponse](../../pkg/models/operations/createtasktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGenaiTask

Retrieves a single GenaiTask by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGenaiTask" method="get" path="/genai/{connection_id}/task/{id}" -->
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

    res, err := s.Task.GetGenaiTask(ctx, operations.GetGenaiTaskRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetGenaiTaskRequest](../../pkg/models/operations/getgenaitaskrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetGenaiTaskResponse](../../pkg/models/operations/getgenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTaskChange

Retrieves a single TaskChange by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskChange" method="get" path="/task/{connection_id}/change/{id}" -->
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

    res, err := s.Task.GetTaskChange(ctx, operations.GetTaskChangeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskChange != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetTaskChangeRequest](../../pkg/models/operations/gettaskchangerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetTaskChangeResponse](../../pkg/models/operations/gettaskchangeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTaskComment

Retrieves a single TaskComment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskComment" method="get" path="/task/{connection_id}/comment/{id}" -->
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

    res, err := s.Task.GetTaskComment(ctx, operations.GetTaskCommentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTaskCommentRequest](../../pkg/models/operations/gettaskcommentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetTaskCommentResponse](../../pkg/models/operations/gettaskcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTaskProject

Retrieves a single TaskProject by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskProject" method="get" path="/task/{connection_id}/project/{id}" -->
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

    res, err := s.Task.GetTaskProject(ctx, operations.GetTaskProjectRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTaskProjectRequest](../../pkg/models/operations/gettaskprojectrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetTaskProjectResponse](../../pkg/models/operations/gettaskprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTaskTask

Retrieves a single TaskTask by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskTask" method="get" path="/task/{connection_id}/task/{id}" -->
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

    res, err := s.Task.GetTaskTask(ctx, operations.GetTaskTaskRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetTaskTaskRequest](../../pkg/models/operations/gettasktaskrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetTaskTaskResponse](../../pkg/models/operations/gettasktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGenaiTasks

Returns a list of GenaiTasks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenaiTasks" method="get" path="/genai/{connection_id}/task" -->
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

    res, err := s.Task.ListGenaiTasks(ctx, operations.ListGenaiTasksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListGenaiTasksRequest](../../pkg/models/operations/listgenaitasksrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListGenaiTasksResponse](../../pkg/models/operations/listgenaitasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTaskChanges

Returns a list of TaskChanges. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTaskChanges" method="get" path="/task/{connection_id}/change" -->
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

    res, err := s.Task.ListTaskChanges(ctx, operations.ListTaskChangesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskChanges != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListTaskChangesRequest](../../pkg/models/operations/listtaskchangesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListTaskChangesResponse](../../pkg/models/operations/listtaskchangesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTaskComments

Returns a list of TaskComments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTaskComments" method="get" path="/task/{connection_id}/comment" -->
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

    res, err := s.Task.ListTaskComments(ctx, operations.ListTaskCommentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskComments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTaskCommentsRequest](../../pkg/models/operations/listtaskcommentsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListTaskCommentsResponse](../../pkg/models/operations/listtaskcommentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTaskProjects

Returns a list of TaskProjects. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTaskProjects" method="get" path="/task/{connection_id}/project" -->
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

    res, err := s.Task.ListTaskProjects(ctx, operations.ListTaskProjectsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskProjects != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTaskProjectsRequest](../../pkg/models/operations/listtaskprojectsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListTaskProjectsResponse](../../pkg/models/operations/listtaskprojectsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTaskTasks

Returns a list of TaskTasks. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTaskTasks" method="get" path="/task/{connection_id}/task" -->
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

    res, err := s.Task.ListTaskTasks(ctx, operations.ListTaskTasksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskTasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListTaskTasksRequest](../../pkg/models/operations/listtasktasksrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListTaskTasksResponse](../../pkg/models/operations/listtasktasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchGenaiTask

Updates the GenaiTask identified by `id`. Provide the unified GenaiTask fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchGenaiTask" method="patch" path="/genai/{connection_id}/task/{id}" example="genai_task" -->
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

    res, err := s.Task.PatchGenaiTask(ctx, operations.PatchGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-10-02T02:15:37.686Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("4013dc78-f8cd-401f-8553-cc0548fc6a44"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-22T00:40:23.688Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-26T16:31:59.686Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchGenaiTaskRequest](../../pkg/models/operations/patchgenaitaskrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchGenaiTaskResponse](../../pkg/models/operations/patchgenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchTaskComment

Updates the TaskComment identified by `id`. Provide the unified TaskComment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchTaskComment" method="patch" path="/task/{connection_id}/comment/{id}" example="task_comment" -->
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

    res, err := s.Task.PatchTaskComment(ctx, operations.PatchTaskCommentRequest{
        TaskComment: shared.TaskComment{
            CreatedAt: types.MustNewTimeFromString("2019-10-12T20:33:37.879Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("7327dfbd-4c0c-4c6c-b2f8-f9a1b36b2677"),
            Metadata: []shared.TaskMetadata{},
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-10-03T14:40:44.798Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchTaskCommentRequest](../../pkg/models/operations/patchtaskcommentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchTaskCommentResponse](../../pkg/models/operations/patchtaskcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchTaskProject

Updates the TaskProject identified by `id`. Provide the unified TaskProject fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchTaskProject" method="patch" path="/task/{connection_id}/project/{id}" example="task_project" -->
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

    res, err := s.Task.PatchTaskProject(ctx, operations.PatchTaskProjectRequest{
        TaskProject: shared.TaskProject{
            CreatedAt: types.MustNewTimeFromString("2023-06-23T16:39:40.446Z"),
            Description: unifiedgosdk.Pointer("Valetudo aggredior accommodo curiositas vox."),
            HasChildren: unifiedgosdk.Pointer(false),
            HasTasks: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("4b217ca0-ecb9-4c4b-ae3d-75b658e307e7"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b7b73d5c-f1ff-48b2-8a4c-101025b18919"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("decens"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "uterque",
                    )),
                },
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("279643ea-8251-402f-bc4c-eeb7f28d150a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-11T16:39:32.560Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchTaskProjectRequest](../../pkg/models/operations/patchtaskprojectrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchTaskProjectResponse](../../pkg/models/operations/patchtaskprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchTaskTask

Updates the TaskTask identified by `id`. Provide the unified TaskTask fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchTaskTask" method="patch" path="/task/{connection_id}/task/{id}" example="task_task" -->
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

    res, err := s.Task.PatchTaskTask(ctx, operations.PatchTaskTaskRequest{
        TaskTask: shared.TaskTask{
            AttachmentIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2025-09-12T14:15:13.482Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("3ea49707-d19b-40c7-82b3-7a5159a1cfea"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("I'll compress the optical IB interface, that should capacitor the JBOD system!"),
            Notes: unifiedgosdk.Pointer("Argentum ventosus rem talis neque thymbra. Absconditus peccatus crustulum tollo cometes volva vacuus eos cedo. Somnus utpote coadunatio denuncio adopto autus sono atrocitas. Concido rerum cura aveho torqueo tempore tubineus vomer tempore suscipit. Clibanus valetudo vulnero pauper caput absorbeo correptius vir deduco."),
            Priority: unifiedgosdk.Pointer("MEDIUM"),
            Progress: unifiedgosdk.Pointer[float64](33.0),
            StartAt: types.MustNewTimeFromString("2025-08-15T04:23:44.520Z"),
            Status: shared.TaskTaskStatusOpened.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "curis",
                "crinis",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](5359.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("magni"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-07T11:59:37.154Z"),
            URL: unifiedgosdk.Pointer("https://elementary-utilization.net/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchTaskTaskRequest](../../pkg/models/operations/patchtasktaskrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchTaskTaskResponse](../../pkg/models/operations/patchtasktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveGenaiTask

Deletes the GenaiTask identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGenaiTask" method="delete" path="/genai/{connection_id}/task/{id}" -->
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

    res, err := s.Task.RemoveGenaiTask(ctx, operations.RemoveGenaiTaskRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveGenaiTaskRequest](../../pkg/models/operations/removegenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveGenaiTaskResponse](../../pkg/models/operations/removegenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveTaskComment

Deletes the TaskComment identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeTaskComment" method="delete" path="/task/{connection_id}/comment/{id}" -->
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

    res, err := s.Task.RemoveTaskComment(ctx, operations.RemoveTaskCommentRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveTaskCommentRequest](../../pkg/models/operations/removetaskcommentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveTaskCommentResponse](../../pkg/models/operations/removetaskcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveTaskProject

Deletes the TaskProject identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeTaskProject" method="delete" path="/task/{connection_id}/project/{id}" -->
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

    res, err := s.Task.RemoveTaskProject(ctx, operations.RemoveTaskProjectRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveTaskProjectRequest](../../pkg/models/operations/removetaskprojectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveTaskProjectResponse](../../pkg/models/operations/removetaskprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveTaskTask

Deletes the TaskTask identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeTaskTask" method="delete" path="/task/{connection_id}/task/{id}" -->
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

    res, err := s.Task.RemoveTaskTask(ctx, operations.RemoveTaskTaskRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveTaskTaskRequest](../../pkg/models/operations/removetasktaskrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveTaskTaskResponse](../../pkg/models/operations/removetasktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateGenaiTask

Updates the GenaiTask identified by `id`. Provide the unified GenaiTask fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGenaiTask" method="put" path="/genai/{connection_id}/task/{id}" example="genai_task" -->
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

    res, err := s.Task.UpdateGenaiTask(ctx, operations.UpdateGenaiTaskRequest{
        GenaiTask: shared.GenaiTask{
            CompletedAt: types.MustNewTimeFromString("2025-10-02T02:15:37.686Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("4013dc78-f8cd-401f-8553-cc0548fc6a44"),
            Instructions: unifiedgosdk.Pointer("Benigne canonicus officiis solvo adsidue deleo angustus."),
            LinesAdded: unifiedgosdk.Pointer[float64](244.0),
            LinesDeleted: unifiedgosdk.Pointer[float64](118.0),
            Messages: []shared.GenaiContent{
                shared.GenaiContent{
                    Content: "Stultus esse cursim stabilis tenetur amet contigo tristis.",
                    Images: []string{},
                    Role: shared.RoleAssistant.ToPointer(),
                },
            },
            Name: unifiedgosdk.Pointer("connect multi-byte port"),
            PullrequestURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance/pull/383"),
            RepoURL: unifiedgosdk.Pointer("https://github.com/berenice.satterfield/joshingly-ignorance"),
            SourceBranchIdentifier: unifiedgosdk.Pointer("main"),
            StartedAt: types.MustNewTimeFromString("2024-05-22T00:40:23.688Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-26T16:31:59.686Z"),
            WebURL: unifiedgosdk.Pointer("https://inexperienced-adrenalin.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GenaiTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateGenaiTaskRequest](../../pkg/models/operations/updategenaitaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateGenaiTaskResponse](../../pkg/models/operations/updategenaitaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTaskComment

Updates the TaskComment identified by `id`. Provide the unified TaskComment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaskComment" method="put" path="/task/{connection_id}/comment/{id}" example="task_comment" -->
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

    res, err := s.Task.UpdateTaskComment(ctx, operations.UpdateTaskCommentRequest{
        TaskComment: shared.TaskComment{
            CreatedAt: types.MustNewTimeFromString("2019-10-12T20:33:37.879Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("7327dfbd-4c0c-4c6c-b2f8-f9a1b36b2677"),
            Metadata: []shared.TaskMetadata{},
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-10-03T14:40:44.798Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateTaskCommentRequest](../../pkg/models/operations/updatetaskcommentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateTaskCommentResponse](../../pkg/models/operations/updatetaskcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTaskProject

Updates the TaskProject identified by `id`. Provide the unified TaskProject fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaskProject" method="put" path="/task/{connection_id}/project/{id}" example="task_project" -->
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

    res, err := s.Task.UpdateTaskProject(ctx, operations.UpdateTaskProjectRequest{
        TaskProject: shared.TaskProject{
            CreatedAt: types.MustNewTimeFromString("2023-06-23T16:39:40.446Z"),
            Description: unifiedgosdk.Pointer("Valetudo aggredior accommodo curiositas vox."),
            HasChildren: unifiedgosdk.Pointer(false),
            HasTasks: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("4b217ca0-ecb9-4c4b-ae3d-75b658e307e7"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b7b73d5c-f1ff-48b2-8a4c-101025b18919"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("decens"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "uterque",
                    )),
                },
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("279643ea-8251-402f-bc4c-eeb7f28d150a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-11T16:39:32.560Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateTaskProjectRequest](../../pkg/models/operations/updatetaskprojectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateTaskProjectResponse](../../pkg/models/operations/updatetaskprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTaskTask

Updates the TaskTask identified by `id`. Provide the unified TaskTask fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaskTask" method="put" path="/task/{connection_id}/task/{id}" example="task_task" -->
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

    res, err := s.Task.UpdateTaskTask(ctx, operations.UpdateTaskTaskRequest{
        TaskTask: shared.TaskTask{
            AttachmentIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2025-09-12T14:15:13.482Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("3ea49707-d19b-40c7-82b3-7a5159a1cfea"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("I'll compress the optical IB interface, that should capacitor the JBOD system!"),
            Notes: unifiedgosdk.Pointer("Argentum ventosus rem talis neque thymbra. Absconditus peccatus crustulum tollo cometes volva vacuus eos cedo. Somnus utpote coadunatio denuncio adopto autus sono atrocitas. Concido rerum cura aveho torqueo tempore tubineus vomer tempore suscipit. Clibanus valetudo vulnero pauper caput absorbeo correptius vir deduco."),
            Priority: unifiedgosdk.Pointer("MEDIUM"),
            Progress: unifiedgosdk.Pointer[float64](33.0),
            StartAt: types.MustNewTimeFromString("2025-08-15T04:23:44.520Z"),
            Status: shared.TaskTaskStatusOpened.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "curis",
                "crinis",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](5359.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("magni"),
            UpdatedAt: types.MustNewTimeFromString("2022-04-07T11:59:37.154Z"),
            URL: unifiedgosdk.Pointer("https://elementary-utilization.net/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskTask != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateTaskTaskRequest](../../pkg/models/operations/updatetasktaskrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateTaskTaskResponse](../../pkg/models/operations/updatetasktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |