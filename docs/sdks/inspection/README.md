# Inspection

## Overview

### Available Operations

* [CreateQmsInspection](#createqmsinspection) - Create an inspection
* [GetQmsInspection](#getqmsinspection) - Retrieve an inspection
* [ListQmsInspections](#listqmsinspections) - List all inspections
* [PatchQmsInspection](#patchqmsinspection) - Update an inspection
* [UpdateQmsInspection](#updateqmsinspection) - Update an inspection

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

    res, err := s.Inspection.CreateQmsInspection(ctx, operations.CreateQmsInspectionRequest{
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

    res, err := s.Inspection.GetQmsInspection(ctx, operations.GetQmsInspectionRequest{
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

    res, err := s.Inspection.ListQmsInspections(ctx, operations.ListQmsInspectionsRequest{
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

    res, err := s.Inspection.PatchQmsInspection(ctx, operations.PatchQmsInspectionRequest{
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

    res, err := s.Inspection.UpdateQmsInspection(ctx, operations.UpdateQmsInspectionRequest{
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