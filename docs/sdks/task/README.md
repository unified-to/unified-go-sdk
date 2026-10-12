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

Create a task

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
            CompletedAt: types.MustNewTimeFromString("2025-09-23T18:24:22.065Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("bd88ff56-3515-41de-aff3-fa358528a7a8"),
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
            StartedAt: types.MustNewTimeFromString("2024-05-16T00:03:01.183Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-22T17:50:57.202Z"),
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
            ID: unifiedgosdk.Pointer("7ed0cefd-76ae-431a-8f4c-22665b2072ba"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-30T18:27:46.985Z"),
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
            ID: unifiedgosdk.Pointer("ffd408c6-9b5a-4530-aab0-31d3762b20d6"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("2e630dc8-318c-400f-b0ed-eb190efe123a"),
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
                    ID: unifiedgosdk.Pointer("95767a4e-b80a-417f-bc51-cbf7b14bfe14"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-10T18:43:27.641Z"),
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
            CompletedAt: types.MustNewTimeFromString("2022-04-03T08:03:55.170Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-05-15T23:37:19.757Z"),
            EndAt: types.MustNewTimeFromString("2022-10-25T07:23:28.714Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("634aafe0-4e79-4ce2-863b-ae90c673d1ce"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-28T18:25:59.869Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-14T20:22:22.681Z"),
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

## GetGenaiTask

Retrieve a task

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

## ListGenaiTasks

List all tasks

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

## PatchGenaiTask

Update a task

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
            CompletedAt: types.MustNewTimeFromString("2025-09-23T18:24:22.082Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("f52e5187-2209-44b0-b515-11d2c80d85b1"),
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
            StartedAt: types.MustNewTimeFromString("2024-05-16T00:03:01.195Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-22T17:50:57.209Z"),
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
            ID: unifiedgosdk.Pointer("59da4442-d55b-4309-aeb2-823d2f161777"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-30T18:27:46.986Z"),
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
            ID: unifiedgosdk.Pointer("6b43199d-8b61-47b7-b455-bc0ffbbd4ea7"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c7f6460f-d4b1-414b-9051-7d80a673fae4"),
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
                    ID: unifiedgosdk.Pointer("b1419f74-2b89-4771-9ecb-dfa72a05866e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-10T18:43:27.642Z"),
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
            CompletedAt: types.MustNewTimeFromString("2022-04-03T08:03:55.177Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-05-15T23:37:19.775Z"),
            EndAt: types.MustNewTimeFromString("2022-10-25T07:23:28.723Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("38690bb9-6898-4a16-83fc-b67434b8c7f5"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-28T18:25:59.876Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-14T20:22:22.682Z"),
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

## RemoveGenaiTask

Remove a task

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

## UpdateGenaiTask

Update a task

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
            CompletedAt: types.MustNewTimeFromString("2025-09-23T18:24:22.082Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-25T20:19:33.247Z"),
            FilesChanged: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("f52e5187-2209-44b0-b515-11d2c80d85b1"),
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
            StartedAt: types.MustNewTimeFromString("2024-05-16T00:03:01.195Z"),
            Status: shared.GenaiTaskStatusBlocked.ToPointer(),
            Summary: unifiedgosdk.Pointer("Cur aeternus cogito vesper."),
            TargetBranchIdentifier: unifiedgosdk.Pointer("agent/joshingly-ignorance"),
            TokensUsed: unifiedgosdk.Pointer[float64](2165.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-22T17:50:57.209Z"),
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
            ID: unifiedgosdk.Pointer("59da4442-d55b-4309-aeb2-823d2f161777"),
            Text: unifiedgosdk.Pointer("Colo ulciscor sublime tabernus."),
            UpdatedAt: types.MustNewTimeFromString("2021-09-30T18:27:46.986Z"),
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
            ID: unifiedgosdk.Pointer("6b43199d-8b61-47b7-b455-bc0ffbbd4ea7"),
            Metadata: []shared.TaskMetadata{
                shared.TaskMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateTaskMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.TaskMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c7f6460f-d4b1-414b-9051-7d80a673fae4"),
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
                    ID: unifiedgosdk.Pointer("b1419f74-2b89-4771-9ecb-dfa72a05866e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("benevolentia"),
                    Value: unifiedgosdk.Pointer(shared.CreateTaskMetadataValueStr(
                        "pariatur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Garden"),
            UpdatedAt: types.MustNewTimeFromString("2023-10-10T18:43:27.642Z"),
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
            CompletedAt: types.MustNewTimeFromString("2022-04-03T08:03:55.177Z"),
            CreatedAt: types.MustNewTimeFromString("2019-01-31T08:34:55.626Z"),
            DueAt: types.MustNewTimeFromString("2026-05-15T23:37:19.775Z"),
            EndAt: types.MustNewTimeFromString("2022-10-25T07:23:28.723Z"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("38690bb9-6898-4a16-83fc-b67434b8c7f5"),
            Metadata: []shared.TaskMetadata{},
            Name: unifiedgosdk.Pointer("Direct Markets Architect"),
            Notes: unifiedgosdk.Pointer("Calcar vilicus audacia ut cultura argentum ventosus. Talis neque thymbra titulus absconditus peccatus crustulum tollo. Volva vacuus eos cedo spero. Utpote coadunatio denuncio adopto autus sono atrocitas vulnero."),
            Priority: unifiedgosdk.Pointer("LOW"),
            Progress: unifiedgosdk.Pointer[float64](2.0),
            StartAt: types.MustNewTimeFromString("2022-01-28T18:25:59.876Z"),
            Status: shared.TaskTaskStatusInProgress.ToPointer(),
            StoryPoints: unifiedgosdk.Pointer[float64](0.0),
            Tags: []string{
                "concido",
                "rerum",
            },
            TimeSpent: unifiedgosdk.Pointer[float64](957.0),
            TimeSpentUnit: unifiedgosdk.Pointer("SECONDS"),
            Type: unifiedgosdk.Pointer("tubineus"),
            UpdatedAt: types.MustNewTimeFromString("2019-07-14T20:22:22.682Z"),
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