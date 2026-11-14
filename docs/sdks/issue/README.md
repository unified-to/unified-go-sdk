# Issue

## Overview

### Available Operations

* [CreateQmsIssue](#createqmsissue) - Create an issue
* [GetQmsIssue](#getqmsissue) - Retrieve an issue
* [GetUnifiedIssue](#getunifiedissue) - Retrieve support issue
* [ListQmsIssues](#listqmsissues) - List all issues
* [ListUnifiedIssues](#listunifiedissues) - List support issues
* [PatchQmsIssue](#patchqmsissue) - Update an issue
* [UpdateQmsIssue](#updateqmsissue) - Update an issue

## CreateQmsIssue

Creates a QmsIssue. Provide the unified QmsIssue fields in the request body; they are mapped to the provider's API. Returns the created QmsIssue.

### Example Usage

<!-- UsageSnippet language="go" operationID="createQmsIssue" method="post" path="/qms/{connection_id}/issue" example="qms_issue" -->
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

    res, err := s.Issue.CreateQmsIssue(ctx, operations.CreateQmsIssueRequest{
        QmsIssue: shared.QmsIssue{
            ActionIds: []string{},
            AffectedItems: []shared.QmsAffectedItem{
                shared.QmsAffectedItem{
                    PartNumber: unifiedgosdk.Pointer("PRT-051260"),
                    Quantity: unifiedgosdk.Pointer[float64](27.0),
                    Revision: unifiedgosdk.Pointer("B"),
                    SerialNumber: unifiedgosdk.Pointer("SN-419784"),
                },
            },
            ChangeIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-10T05:10:42.904Z"),
            Description: unifiedgosdk.Pointer("Adfero delectatio earum contabesco convoco uxor avaritia volaticus dens."),
            DueAt: types.MustNewTimeFromString("2024-07-24T01:11:27.132Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("6a736e92-e636-4355-8e69-91993b68e66c"),
            LifecycleState: unifiedgosdk.Pointer("In Progress"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Labeling error on tuxedo"),
            Number: unifiedgosdk.Pointer("CMP-5154"),
            OpenedAt: types.MustNewTimeFromString("2024-07-09T00:37:03.340Z"),
            OriginalType: unifiedgosdk.Pointer("Customer Complaint"),
            Severity: shared.SeverityMedium.ToPointer(),
            Source: shared.QmsIssueSourceCustomer.ToPointer(),
            Status: shared.QmsIssueStatusInProgress.ToPointer(),
            Type: shared.QmsIssueTypeComplaint.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-03-13T04:59:55.897Z"),
            URL: unifiedgosdk.Pointer("https://annual-cycle.org/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateQmsIssueRequest](../../pkg/models/operations/createqmsissuerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateQmsIssueResponse](../../pkg/models/operations/createqmsissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetQmsIssue

Retrieves a single QmsIssue by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQmsIssue" method="get" path="/qms/{connection_id}/issue/{id}" -->
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

    res, err := s.Issue.GetQmsIssue(ctx, operations.GetQmsIssueRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetQmsIssueRequest](../../pkg/models/operations/getqmsissuerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetQmsIssueResponse](../../pkg/models/operations/getqmsissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUnifiedIssue

Retrieve support issue

### Example Usage

<!-- UsageSnippet language="go" operationID="getUnifiedIssue" method="get" path="/unified/issue/{id}" -->
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

    res, err := s.Issue.GetUnifiedIssue(ctx, operations.GetUnifiedIssueRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetUnifiedIssueRequest](../../pkg/models/operations/getunifiedissuerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetUnifiedIssueResponse](../../pkg/models/operations/getunifiedissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListQmsIssues

Returns a list of QmsIssues. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listQmsIssues" method="get" path="/qms/{connection_id}/issue" -->
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

    res, err := s.Issue.ListQmsIssues(ctx, operations.ListQmsIssuesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsIssues != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListQmsIssuesRequest](../../pkg/models/operations/listqmsissuesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListQmsIssuesResponse](../../pkg/models/operations/listqmsissuesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedIssues

List support issues

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedIssues" method="get" path="/unified/issue" -->
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

    res, err := s.Issue.ListUnifiedIssues(ctx, operations.ListUnifiedIssuesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Issues != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListUnifiedIssuesRequest](../../pkg/models/operations/listunifiedissuesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListUnifiedIssuesResponse](../../pkg/models/operations/listunifiedissuesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchQmsIssue

Updates the QmsIssue identified by `id`. Provide the unified QmsIssue fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchQmsIssue" method="patch" path="/qms/{connection_id}/issue/{id}" example="qms_issue" -->
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

    res, err := s.Issue.PatchQmsIssue(ctx, operations.PatchQmsIssueRequest{
        QmsIssue: shared.QmsIssue{
            ActionIds: []string{},
            AffectedItems: []shared.QmsAffectedItem{
                shared.QmsAffectedItem{
                    PartNumber: unifiedgosdk.Pointer("PRT-051260"),
                    Quantity: unifiedgosdk.Pointer[float64](27.0),
                    Revision: unifiedgosdk.Pointer("B"),
                    SerialNumber: unifiedgosdk.Pointer("SN-419784"),
                },
            },
            ChangeIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-10T05:10:42.904Z"),
            Description: unifiedgosdk.Pointer("Adfero delectatio earum contabesco convoco uxor avaritia volaticus dens."),
            DueAt: types.MustNewTimeFromString("2024-07-24T01:11:27.151Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("346e2a73-787e-4b59-94ed-68a91185cd06"),
            LifecycleState: unifiedgosdk.Pointer("In Progress"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Labeling error on tuxedo"),
            Number: unifiedgosdk.Pointer("CMP-5154"),
            OpenedAt: types.MustNewTimeFromString("2024-07-09T00:37:03.359Z"),
            OriginalType: unifiedgosdk.Pointer("Customer Complaint"),
            Severity: shared.SeverityMedium.ToPointer(),
            Source: shared.QmsIssueSourceCustomer.ToPointer(),
            Status: shared.QmsIssueStatusInProgress.ToPointer(),
            Type: shared.QmsIssueTypeComplaint.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-03-13T04:59:55.900Z"),
            URL: unifiedgosdk.Pointer("https://annual-cycle.org/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchQmsIssueRequest](../../pkg/models/operations/patchqmsissuerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchQmsIssueResponse](../../pkg/models/operations/patchqmsissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateQmsIssue

Updates the QmsIssue identified by `id`. Provide the unified QmsIssue fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateQmsIssue" method="put" path="/qms/{connection_id}/issue/{id}" example="qms_issue" -->
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

    res, err := s.Issue.UpdateQmsIssue(ctx, operations.UpdateQmsIssueRequest{
        QmsIssue: shared.QmsIssue{
            ActionIds: []string{},
            AffectedItems: []shared.QmsAffectedItem{
                shared.QmsAffectedItem{
                    PartNumber: unifiedgosdk.Pointer("PRT-051260"),
                    Quantity: unifiedgosdk.Pointer[float64](27.0),
                    Revision: unifiedgosdk.Pointer("B"),
                    SerialNumber: unifiedgosdk.Pointer("SN-419784"),
                },
            },
            ChangeIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2020-06-10T05:10:42.904Z"),
            Description: unifiedgosdk.Pointer("Adfero delectatio earum contabesco convoco uxor avaritia volaticus dens."),
            DueAt: types.MustNewTimeFromString("2024-07-24T01:11:27.151Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("346e2a73-787e-4b59-94ed-68a91185cd06"),
            LifecycleState: unifiedgosdk.Pointer("In Progress"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Labeling error on tuxedo"),
            Number: unifiedgosdk.Pointer("CMP-5154"),
            OpenedAt: types.MustNewTimeFromString("2024-07-09T00:37:03.359Z"),
            OriginalType: unifiedgosdk.Pointer("Customer Complaint"),
            Severity: shared.SeverityMedium.ToPointer(),
            Source: shared.QmsIssueSourceCustomer.ToPointer(),
            Status: shared.QmsIssueStatusInProgress.ToPointer(),
            Type: shared.QmsIssueTypeComplaint.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-03-13T04:59:55.900Z"),
            URL: unifiedgosdk.Pointer("https://annual-cycle.org/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateQmsIssueRequest](../../pkg/models/operations/updateqmsissuerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateQmsIssueResponse](../../pkg/models/operations/updateqmsissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |