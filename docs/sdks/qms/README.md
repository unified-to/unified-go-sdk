# Qms

## Overview

### Available Operations

* [CreateQmsAction](#createqmsaction) - Create an action
* [CreateQmsAudit](#createqmsaudit) - Create an audit
* [CreateQmsInspection](#createqmsinspection) - Create an inspection
* [CreateQmsIssue](#createqmsissue) - Create an issue
* [GetQmsAction](#getqmsaction) - Retrieve an action
* [GetQmsAudit](#getqmsaudit) - Retrieve an audit
* [GetQmsInspection](#getqmsinspection) - Retrieve an inspection
* [GetQmsIssue](#getqmsissue) - Retrieve an issue
* [ListQmsActions](#listqmsactions) - List all actions
* [ListQmsAudits](#listqmsaudits) - List all audits
* [ListQmsInspections](#listqmsinspections) - List all inspections
* [ListQmsIssues](#listqmsissues) - List all issues
* [PatchQmsAction](#patchqmsaction) - Update an action
* [PatchQmsAudit](#patchqmsaudit) - Update an audit
* [PatchQmsInspection](#patchqmsinspection) - Update an inspection
* [PatchQmsIssue](#patchqmsissue) - Update an issue
* [UpdateQmsAction](#updateqmsaction) - Update an action
* [UpdateQmsAudit](#updateqmsaudit) - Update an audit
* [UpdateQmsInspection](#updateqmsinspection) - Update an inspection
* [UpdateQmsIssue](#updateqmsissue) - Update an issue

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

    res, err := s.Qms.CreateQmsAction(ctx, operations.CreateQmsActionRequest{
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

    res, err := s.Qms.CreateQmsAudit(ctx, operations.CreateQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-03-27T12:26:43.637Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("ba563e1e-69cb-4591-a0fe-89ee18200ecf"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("soggy certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-25T01:49:46.392Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:08:39.795Z"),
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

## CreateQmsInspection

Creates a QmsInspection. Provide the unified QmsInspection fields in the request body; they are mapped to the provider's API. Returns the created QmsInspection.

### Example Usage

<!-- UsageSnippet language="go" operationID="createQmsInspection" method="post" path="/qms/{connection_id}/inspection" example="qms_inspection" -->
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

    res, err := s.Qms.CreateQmsInspection(ctx, operations.CreateQmsInspectionRequest{
        QmsInspection: shared.QmsInspection{
            Characteristics: []shared.QmsCharacteristic{
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](152.9),
                    MeasuredValue: unifiedgosdk.Pointer("152.94"),
                    Name: unifiedgosdk.Pointer("Hole position"),
                    NominalValue: unifiedgosdk.Pointer[float64](153.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](153.1),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](24.95),
                    MeasuredValue: unifiedgosdk.Pointer("25.03"),
                    Name: unifiedgosdk.Pointer("Flatness"),
                    NominalValue: unifiedgosdk.Pointer[float64](25.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](25.05),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](53.99),
                    MeasuredValue: unifiedgosdk.Pointer("53.99"),
                    Name: unifiedgosdk.Pointer("Bore diameter"),
                    NominalValue: unifiedgosdk.Pointer[float64](54.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](54.01),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2020-04-18T21:23:54.000Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("a2ab7393-8840-4428-aa12-b0a300c19d46"),
            InspectedAt: types.MustNewTimeFromString("2021-02-07T09:05:13.853Z"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Complete"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Source inspection"),
            Number: unifiedgosdk.Pointer("INSP-96996"),
            PartNumber: unifiedgosdk.Pointer("PRT-048483"),
            QuantityAccepted: unifiedgosdk.Pointer[float64](41.0),
            QuantityInspected: unifiedgosdk.Pointer[float64](41.0),
            QuantityRejected: unifiedgosdk.Pointer[float64](0.0),
            Result: shared.QmsInspectionResultPass.ToPointer(),
            Revision: unifiedgosdk.Pointer("A"),
            Status: shared.QmsInspectionStatusClosed.ToPointer(),
            SupplierName: unifiedgosdk.Pointer("Gerlach LLC"),
            Type: shared.QmsInspectionTypeSource.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-08T16:33:47.202Z"),
            URL: unifiedgosdk.Pointer("https://purple-newsprint.name/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsInspection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateQmsInspectionRequest](../../pkg/models/operations/createqmsinspectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateQmsInspectionResponse](../../pkg/models/operations/createqmsinspectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Qms.CreateQmsIssue(ctx, operations.CreateQmsIssueRequest{
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

    res, err := s.Qms.GetQmsAction(ctx, operations.GetQmsActionRequest{
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

    res, err := s.Qms.GetQmsAudit(ctx, operations.GetQmsAuditRequest{
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

## GetQmsInspection

Retrieves a single QmsInspection by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQmsInspection" method="get" path="/qms/{connection_id}/inspection/{id}" -->
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

    res, err := s.Qms.GetQmsInspection(ctx, operations.GetQmsInspectionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsInspection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetQmsInspectionRequest](../../pkg/models/operations/getqmsinspectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetQmsInspectionResponse](../../pkg/models/operations/getqmsinspectionresponse.md), error**

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

    res, err := s.Qms.GetQmsIssue(ctx, operations.GetQmsIssueRequest{
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

    res, err := s.Qms.ListQmsActions(ctx, operations.ListQmsActionsRequest{
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

    res, err := s.Qms.ListQmsAudits(ctx, operations.ListQmsAuditsRequest{
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

## ListQmsInspections

Returns a list of QmsInspections. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listQmsInspections" method="get" path="/qms/{connection_id}/inspection" -->
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

    res, err := s.Qms.ListQmsInspections(ctx, operations.ListQmsInspectionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsInspections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListQmsInspectionsRequest](../../pkg/models/operations/listqmsinspectionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListQmsInspectionsResponse](../../pkg/models/operations/listqmsinspectionsresponse.md), error**

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

    res, err := s.Qms.ListQmsIssues(ctx, operations.ListQmsIssuesRequest{
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

    res, err := s.Qms.PatchQmsAction(ctx, operations.PatchQmsActionRequest{
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

    res, err := s.Qms.PatchQmsAudit(ctx, operations.PatchQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-03-27T12:26:43.653Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("f4536930-3cb4-4856-bb0b-0a6d3daf0716"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("soggy certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-25T01:49:46.398Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:08:39.804Z"),
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

## PatchQmsInspection

Updates the QmsInspection identified by `id`. Provide the unified QmsInspection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchQmsInspection" method="patch" path="/qms/{connection_id}/inspection/{id}" example="qms_inspection" -->
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

    res, err := s.Qms.PatchQmsInspection(ctx, operations.PatchQmsInspectionRequest{
        QmsInspection: shared.QmsInspection{
            Characteristics: []shared.QmsCharacteristic{
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](152.9),
                    MeasuredValue: unifiedgosdk.Pointer("152.94"),
                    Name: unifiedgosdk.Pointer("Hole position"),
                    NominalValue: unifiedgosdk.Pointer[float64](153.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](153.1),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](24.95),
                    MeasuredValue: unifiedgosdk.Pointer("25.03"),
                    Name: unifiedgosdk.Pointer("Flatness"),
                    NominalValue: unifiedgosdk.Pointer[float64](25.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](25.05),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](53.99),
                    MeasuredValue: unifiedgosdk.Pointer("53.99"),
                    Name: unifiedgosdk.Pointer("Bore diameter"),
                    NominalValue: unifiedgosdk.Pointer[float64](54.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](54.01),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2020-04-18T21:23:54.000Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("83a13528-fddb-40da-ab38-ba2368fd532f"),
            InspectedAt: types.MustNewTimeFromString("2021-02-07T09:05:13.856Z"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Complete"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Source inspection"),
            Number: unifiedgosdk.Pointer("INSP-96996"),
            PartNumber: unifiedgosdk.Pointer("PRT-048483"),
            QuantityAccepted: unifiedgosdk.Pointer[float64](41.0),
            QuantityInspected: unifiedgosdk.Pointer[float64](41.0),
            QuantityRejected: unifiedgosdk.Pointer[float64](0.0),
            Result: shared.QmsInspectionResultPass.ToPointer(),
            Revision: unifiedgosdk.Pointer("A"),
            Status: shared.QmsInspectionStatusClosed.ToPointer(),
            SupplierName: unifiedgosdk.Pointer("Gerlach LLC"),
            Type: shared.QmsInspectionTypeSource.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-08T16:33:47.223Z"),
            URL: unifiedgosdk.Pointer("https://purple-newsprint.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsInspection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchQmsInspectionRequest](../../pkg/models/operations/patchqmsinspectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchQmsInspectionResponse](../../pkg/models/operations/patchqmsinspectionresponse.md), error**

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

    res, err := s.Qms.PatchQmsIssue(ctx, operations.PatchQmsIssueRequest{
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

    res, err := s.Qms.UpdateQmsAction(ctx, operations.UpdateQmsActionRequest{
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

    res, err := s.Qms.UpdateQmsAudit(ctx, operations.UpdateQmsAuditRequest{
        QmsAudit: shared.QmsAudit{
            CreatedAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            EndAt: types.MustNewTimeFromString("2026-03-27T12:26:43.653Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("f4536930-3cb4-4856-bb0b-0a6d3daf0716"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("In Review"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("soggy certification audit"),
            Number: unifiedgosdk.Pointer("AUD-215"),
            ScheduledAt: types.MustNewTimeFromString("2021-02-20T23:02:37.667Z"),
            Scope: unifiedgosdk.Pointer("Calibration"),
            Standard: unifiedgosdk.Pointer("ISO 9001:2015"),
            StartAt: types.MustNewTimeFromString("2023-01-25T01:49:46.398Z"),
            Status: shared.QmsAuditStatusInReview.ToPointer(),
            Type: shared.QmsAuditTypeCertification.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-02-08T15:08:39.804Z"),
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

## UpdateQmsInspection

Updates the QmsInspection identified by `id`. Provide the unified QmsInspection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateQmsInspection" method="put" path="/qms/{connection_id}/inspection/{id}" example="qms_inspection" -->
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

    res, err := s.Qms.UpdateQmsInspection(ctx, operations.UpdateQmsInspectionRequest{
        QmsInspection: shared.QmsInspection{
            Characteristics: []shared.QmsCharacteristic{
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](152.9),
                    MeasuredValue: unifiedgosdk.Pointer("152.94"),
                    Name: unifiedgosdk.Pointer("Hole position"),
                    NominalValue: unifiedgosdk.Pointer[float64](153.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](153.1),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](24.95),
                    MeasuredValue: unifiedgosdk.Pointer("25.03"),
                    Name: unifiedgosdk.Pointer("Flatness"),
                    NominalValue: unifiedgosdk.Pointer[float64](25.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](25.05),
                },
                shared.QmsCharacteristic{
                    LowerLimit: unifiedgosdk.Pointer[float64](53.99),
                    MeasuredValue: unifiedgosdk.Pointer("53.99"),
                    Name: unifiedgosdk.Pointer("Bore diameter"),
                    NominalValue: unifiedgosdk.Pointer[float64](54.0),
                    Result: shared.ResultPass.ToPointer(),
                    UnitOfMeasure: unifiedgosdk.Pointer("mm"),
                    UpperLimit: unifiedgosdk.Pointer[float64](54.01),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2020-04-18T21:23:54.000Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("83a13528-fddb-40da-ab38-ba2368fd532f"),
            InspectedAt: types.MustNewTimeFromString("2021-02-07T09:05:13.856Z"),
            IssueIds: []string{},
            LifecycleState: unifiedgosdk.Pointer("Complete"),
            Metadata: []shared.QmsMetadata{},
            Name: unifiedgosdk.Pointer("Source inspection"),
            Number: unifiedgosdk.Pointer("INSP-96996"),
            PartNumber: unifiedgosdk.Pointer("PRT-048483"),
            QuantityAccepted: unifiedgosdk.Pointer[float64](41.0),
            QuantityInspected: unifiedgosdk.Pointer[float64](41.0),
            QuantityRejected: unifiedgosdk.Pointer[float64](0.0),
            Result: shared.QmsInspectionResultPass.ToPointer(),
            Revision: unifiedgosdk.Pointer("A"),
            Status: shared.QmsInspectionStatusClosed.ToPointer(),
            SupplierName: unifiedgosdk.Pointer("Gerlach LLC"),
            Type: shared.QmsInspectionTypeSource.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-08T16:33:47.223Z"),
            URL: unifiedgosdk.Pointer("https://purple-newsprint.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.QmsInspection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateQmsInspectionRequest](../../pkg/models/operations/updateqmsinspectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateQmsInspectionResponse](../../pkg/models/operations/updateqmsinspectionresponse.md), error**

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

    res, err := s.Qms.UpdateQmsIssue(ctx, operations.UpdateQmsIssueRequest{
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