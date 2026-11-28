# Audit

## Overview

### Available Operations

* [CreateQmsAudit](#createqmsaudit) - Create an audit
* [GetQmsAudit](#getqmsaudit) - Retrieve an audit
* [ListQmsAudits](#listqmsaudits) - List all audits
* [PatchQmsAudit](#patchqmsaudit) - Update an audit
* [UpdateQmsAudit](#updateqmsaudit) - Update an audit

## CreateQmsAudit

Creates a QmsAudit. Provide the unified QmsAudit fields in the request body; they are mapped to the provider's API. Returns the created QmsAudit.

### Example Usage

<!-- UsageSnippet language="go" operationID="createQmsAudit" method="post" path="/qms/{connection_id}/audit" example="qms_audit" -->
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

    res, err := s.Audit.CreateQmsAudit(ctx, operations.CreateQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-04-08T22:48:16.677Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("0dd6b493-73a6-4f55-a87c-b447ca1e3c64"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("afraid certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-29T18:34:52.533Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-15T20:45:41.418Z"),
            URL: unifiedgosdk.Pointer("https://mushy-obesity.name/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAudit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateQmsAuditRequest](../../pkg/models/operations/createqmsauditrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateQmsAuditResponse](../../pkg/models/operations/createqmsauditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetQmsAudit

Retrieves a single QmsAudit by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQmsAudit" method="get" path="/qms/{connection_id}/audit/{id}" -->
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

    res, err := s.Audit.GetQmsAudit(ctx, operations.GetQmsAuditRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAudit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetQmsAuditRequest](../../pkg/models/operations/getqmsauditrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetQmsAuditResponse](../../pkg/models/operations/getqmsauditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListQmsAudits

Returns a list of QmsAudits. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listQmsAudits" method="get" path="/qms/{connection_id}/audit" -->
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

    res, err := s.Audit.ListQmsAudits(ctx, operations.ListQmsAuditsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAudits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListQmsAuditsRequest](../../pkg/models/operations/listqmsauditsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListQmsAuditsResponse](../../pkg/models/operations/listqmsauditsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchQmsAudit

Updates the QmsAudit identified by `id`. Provide the unified QmsAudit fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchQmsAudit" method="patch" path="/qms/{connection_id}/audit/{id}" example="qms_audit" -->
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

    res, err := s.Audit.PatchQmsAudit(ctx, operations.PatchQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-04-08T22:48:16.687Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("f52f64b9-28d2-4cae-8459-fe4a3bcf9cd7"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("afraid certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-29T18:34:52.537Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-15T20:45:41.424Z"),
            URL: unifiedgosdk.Pointer("https://mushy-obesity.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAudit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchQmsAuditRequest](../../pkg/models/operations/patchqmsauditrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchQmsAuditResponse](../../pkg/models/operations/patchqmsauditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateQmsAudit

Updates the QmsAudit identified by `id`. Provide the unified QmsAudit fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateQmsAudit" method="put" path="/qms/{connection_id}/audit/{id}" example="qms_audit" -->
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

    res, err := s.Audit.UpdateQmsAudit(ctx, operations.UpdateQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-04-08T22:48:16.687Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("f52f64b9-28d2-4cae-8459-fe4a3bcf9cd7"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("afraid certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-29T18:34:52.537Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-15T20:45:41.424Z"),
            URL: unifiedgosdk.Pointer("https://mushy-obesity.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsAudit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateQmsAuditRequest](../../pkg/models/operations/updateqmsauditrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateQmsAuditResponse](../../pkg/models/operations/updateqmsauditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |