# Task

## Overview

### Available Operations

* [CreateTaskComment](#createtaskcomment) - Create a comment
* [CreateTaskProject](#createtaskproject) - Create a project
* [CreateTaskTask](#createtasktask) - Create a task
* [GetTaskChange](#gettaskchange) - Retrieve a change
* [GetTaskComment](#gettaskcomment) - Retrieve a comment
* [GetTaskProject](#gettaskproject) - Retrieve a project
* [GetTaskTask](#gettasktask) - Retrieve a task
* [ListTaskChanges](#listtaskchanges) - List all changes
* [ListTaskComments](#listtaskcomments) - List all comments
* [ListTaskProjects](#listtaskprojects) - List all projects
* [ListTaskTasks](#listtasktasks) - List all tasks
* [PatchTaskComment](#patchtaskcomment) - Update a comment
* [PatchTaskProject](#patchtaskproject) - Update a project
* [PatchTaskTask](#patchtasktask) - Update a task
* [RemoveTaskComment](#removetaskcomment) - Remove a comment
* [RemoveTaskProject](#removetaskproject) - Remove a project
* [RemoveTaskTask](#removetasktask) - Remove a task
* [UpdateTaskComment](#updatetaskcomment) - Update a comment
* [UpdateTaskProject](#updatetaskproject) - Update a project
* [UpdateTaskTask](#updatetasktask) - Update a task

## CreateTaskComment

Create a comment

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
            ID: unifiedgosdk.Pointer("d6055da6-657e-416c-8565-856e90354101"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-24T01:29:44.004Z"),
            UserName: unifiedgosdk.Pointer("Santina Abbott"),
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

Create a project

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
            ID: unifiedgosdk.Pointer("3687e4c3-b243-4270-99b2-c844cfc26462"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("97e50e71-e1ed-4d6c-90b6-094016d24768"),
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
                    ID: unifiedgosdk.Pointer("6010bbb0-510c-4b9c-8b1d-27fc0418ecd4"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-08T14:57:56.362Z"),
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

Create a task

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
            CompletedAt: types.MustNewTimeFromString("2022-03-24T12:16:02.229Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-04-23T09:38:15.654Z"),
            EndAt: types.MustNewTimeFromString("2022-10-13T17:51:18.132Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("59e40094-d497-4608-93d5-1f6d1c0f4c3c"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-19T11:46:59.440Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-13T10:52:39.955Z"),
            URL: unifiedgosdk.Pointer("https://dismal-silk.net/"),
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

## GetTaskChange

Retrieve a change

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

Retrieve a comment

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

Retrieve a project

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

Retrieve a task

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

## ListTaskChanges

List all changes

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

List all comments

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

List all projects

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

List all tasks

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

## PatchTaskComment

Update a comment

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
            ID: unifiedgosdk.Pointer("4f9bb369-67f3-478f-b7f9-73df24044738"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-24T01:29:44.005Z"),
            UserName: unifiedgosdk.Pointer("Santina Abbott"),
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

Update a project

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
            ID: unifiedgosdk.Pointer("a8597f9e-fdb3-4244-bcd3-c62331ca5ba9"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("20b84b41-2336-4083-bba1-4867b5ccbb3b"),
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
                    ID: unifiedgosdk.Pointer("417b98cd-43b3-462e-b109-2362cc1aacf9"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-08T14:57:56.362Z"),
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

Update a task

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
            CompletedAt: types.MustNewTimeFromString("2022-03-24T12:16:02.235Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-04-23T09:38:15.670Z"),
            EndAt: types.MustNewTimeFromString("2022-10-13T17:51:18.140Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("eb02963c-f722-4ea6-ba76-a9b8519874cc"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-19T11:46:59.446Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-13T10:52:39.955Z"),
            URL: unifiedgosdk.Pointer("https://dismal-silk.net/"),
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

## RemoveTaskComment

Remove a comment

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

Remove a project

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

Remove a task

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

## UpdateTaskComment

Update a comment

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
            ID: unifiedgosdk.Pointer("4f9bb369-67f3-478f-b7f9-73df24044738"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-24T01:29:44.005Z"),
            UserName: unifiedgosdk.Pointer("Santina Abbott"),
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

Update a project

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
            ID: unifiedgosdk.Pointer("a8597f9e-fdb3-4244-bcd3-c62331ca5ba9"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("20b84b41-2336-4083-bba1-4867b5ccbb3b"),
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
                    ID: unifiedgosdk.Pointer("417b98cd-43b3-462e-b109-2362cc1aacf9"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-08T14:57:56.362Z"),
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

Update a task

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
            CompletedAt: types.MustNewTimeFromString("2022-03-24T12:16:02.235Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-04-23T09:38:15.670Z"),
            EndAt: types.MustNewTimeFromString("2022-10-13T17:51:18.140Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("eb02963c-f722-4ea6-ba76-a9b8519874cc"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-19T11:46:59.446Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-13T10:52:39.955Z"),
            URL: unifiedgosdk.Pointer("https://dismal-silk.net/"),
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