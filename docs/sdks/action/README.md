# Action

## Overview

### Available Operations

* [CreateQmsAction](#createqmsaction) - Create an action
* [GetQmsAction](#getqmsaction) - Retrieve an action
* [ListQmsActions](#listqmsactions) - List all actions
* [PatchQmsAction](#patchqmsaction) - Update an action
* [UpdateQmsAction](#updateqmsaction) - Update an action

## CreateQmsAction

Creates a QmsAction. Provide the unified QmsAction fields in the request body; they are mapped to the provider's API. Returns the created QmsAction.

### Example Usage

<!-- UsageSnippet language="go" operationID="createQmsAction" method="post" path="/qms/{connection_id}/action" example="qms_action" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Action.CreateQmsAction(ctx, operations.CreateQmsActionRequest{
        QmsAction: shared.QmsAction{
            AffectedItems: []shared.QmsAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2022-03-07T03:01:29.933Z"),
            Description: unifiedgosdk.Pointer("Auditor sunt tego apparatus aegre ascisco."),
            DueAt: types.MustNewTimeFromString("2022-04-04T03:42:01.290Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("1fc96ebc-7b1b-44d8-8dd9-7358a63bbfba"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Cancelled"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Add inspection step"),
            Number: unifiedgosdk.Pointer("CAPA-6268"),
            OriginalType: unifiedgosdk.Pointer("Corrective Action"),
            Priority: unifiedgosdk.Pointer("Medium"),
            RootCause: unifiedgosdk.Pointer("Worn tooling"),
            Status: shared.QmsActionStatusCancelled.ToPointer(),
            Type: shared.QmsActionTypeCorrective.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-17T01:04:22.672Z"),
            URL: unifiedgosdk.Pointer("https://incomplete-conservation.org"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateQmsActionRequest](../../pkg/models/operations/createqmsactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateQmsActionResponse](../../pkg/models/operations/createqmsactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetQmsAction

Retrieves a single QmsAction by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQmsAction" method="get" path="/qms/{connection_id}/action/{id}" -->
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

    res, err := s.Action.GetQmsAction(ctx, operations.GetQmsActionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetQmsActionRequest](../../pkg/models/operations/getqmsactionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetQmsActionResponse](../../pkg/models/operations/getqmsactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListQmsActions

Returns a list of QmsActions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listQmsActions" method="get" path="/qms/{connection_id}/action" -->
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

    res, err := s.Action.ListQmsActions(ctx, operations.ListQmsActionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsActions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListQmsActionsRequest](../../pkg/models/operations/listqmsactionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListQmsActionsResponse](../../pkg/models/operations/listqmsactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchQmsAction

Updates the QmsAction identified by `id`. Provide the unified QmsAction fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchQmsAction" method="patch" path="/qms/{connection_id}/action/{id}" example="qms_action" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Action.PatchQmsAction(ctx, operations.PatchQmsActionRequest{
        QmsAction: shared.QmsAction{
            AffectedItems: []shared.QmsAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2022-03-07T03:01:29.933Z"),
            Description: unifiedgosdk.Pointer("Auditor sunt tego apparatus aegre ascisco."),
            DueAt: types.MustNewTimeFromString("2022-04-04T03:42:01.290Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("fefe74a8-a238-4567-bcf4-da9ac23d6791"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Cancelled"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Add inspection step"),
            Number: unifiedgosdk.Pointer("CAPA-6268"),
            OriginalType: unifiedgosdk.Pointer("Corrective Action"),
            Priority: unifiedgosdk.Pointer("Medium"),
            RootCause: unifiedgosdk.Pointer("Worn tooling"),
            Status: shared.QmsActionStatusCancelled.ToPointer(),
            Type: shared.QmsActionTypeCorrective.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-17T01:04:22.691Z"),
            URL: unifiedgosdk.Pointer("https://incomplete-conservation.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchQmsActionRequest](../../pkg/models/operations/patchqmsactionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchQmsActionResponse](../../pkg/models/operations/patchqmsactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateQmsAction

Updates the QmsAction identified by `id`. Provide the unified QmsAction fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateQmsAction" method="put" path="/qms/{connection_id}/action/{id}" example="qms_action" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Action.UpdateQmsAction(ctx, operations.UpdateQmsActionRequest{
        QmsAction: shared.QmsAction{
            AffectedItems: []shared.QmsAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2022-03-07T03:01:29.933Z"),
            Description: unifiedgosdk.Pointer("Auditor sunt tego apparatus aegre ascisco."),
            DueAt: types.MustNewTimeFromString("2022-04-04T03:42:01.290Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("fefe74a8-a238-4567-bcf4-da9ac23d6791"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Cancelled"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Add inspection step"),
            Number: unifiedgosdk.Pointer("CAPA-6268"),
            OriginalType: unifiedgosdk.Pointer("Corrective Action"),
            Priority: unifiedgosdk.Pointer("Medium"),
            RootCause: unifiedgosdk.Pointer("Worn tooling"),
            Status: shared.QmsActionStatusCancelled.ToPointer(),
            Type: shared.QmsActionTypeCorrective.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-02-17T01:04:22.691Z"),
            URL: unifiedgosdk.Pointer("https://incomplete-conservation.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateQmsActionRequest](../../pkg/models/operations/updateqmsactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateQmsActionResponse](../../pkg/models/operations/updateqmsactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |