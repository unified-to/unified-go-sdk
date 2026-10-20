# Table

## Overview

### Available Operations

* [CreateDatastoreTable](#createdatastoretable) - Create a table
* [GetDatastoreTable](#getdatastoretable) - Retrieve a table
* [ListDatastoreTables](#listdatastoretables) - List all tables
* [PatchDatastoreTable](#patchdatastoretable) - Update a table
* [RemoveDatastoreTable](#removedatastoretable) - Remove a table
* [UpdateDatastoreTable](#updatedatastoretable) - Update a table

## CreateDatastoreTable

Create a table

### Example Usage

<!-- UsageSnippet language="go" operationID="createDatastoreTable" method="post" path="/datastore/{connection_id}/table" example="datastore_table" -->
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

    res, err := s.Table.CreateDatastoreTable(ctx, operations.CreateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("fd3020f0-8daa-4d01-b1ce-fce949adfbfd"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("29916ab9-e706-4a13-b0a6-27b1a0f5f82c"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("1098677c-552d-4b40-be30-853bafdcdaa9"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "Toy - Kiehn",
                    OriginalType: "sursum",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Audax articulus triumphus tametsi cura tenetur colligo arcesso bestia."),
                    ID: unifiedgosdk.Pointer("4764cf99-0f51-433f-8f92-3aa91f53e698"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("1ae1e0d1-4107-44e1-9661-a49384ae1940"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("2b49a787-8ff2-4719-aaa7-4334ec0fc6f1"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "2b49a787-8ff2-4719-aaa7-4334ec0fc6f1",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-12T19:11:23.441Z"),
            WebURL: unifiedgosdk.Pointer("https://nervous-deck.com"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreTable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDatastoreTableRequest](../../pkg/models/operations/createdatastoretablerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDatastoreTableResponse](../../pkg/models/operations/createdatastoretableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDatastoreTable

Retrieve a table

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatastoreTable" method="get" path="/datastore/{connection_id}/table/{id}" -->
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

    res, err := s.Table.GetDatastoreTable(ctx, operations.GetDatastoreTableRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreTable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDatastoreTableRequest](../../pkg/models/operations/getdatastoretablerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDatastoreTableResponse](../../pkg/models/operations/getdatastoretableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDatastoreTables

List all tables

### Example Usage

<!-- UsageSnippet language="go" operationID="listDatastoreTables" method="get" path="/datastore/{connection_id}/table" -->
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

    res, err := s.Table.ListDatastoreTables(ctx, operations.ListDatastoreTablesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreTables != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDatastoreTablesRequest](../../pkg/models/operations/listdatastoretablesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDatastoreTablesResponse](../../pkg/models/operations/listdatastoretablesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDatastoreTable

Update a table

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDatastoreTable" method="patch" path="/datastore/{connection_id}/table/{id}" example="datastore_table" -->
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

    res, err := s.Table.PatchDatastoreTable(ctx, operations.PatchDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("7f4c87fc-0bfe-4866-a6d4-d38955c97483"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("7eb13a26-f2ac-4a25-acfa-c84b428434e7"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("76f7cee1-2b36-4a0d-b321-03d9f3b1fb7e"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "Toy - Kiehn",
                    OriginalType: "sursum",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Audax articulus triumphus tametsi cura tenetur colligo arcesso bestia."),
                    ID: unifiedgosdk.Pointer("4a4fccf3-5cf4-4ead-b0f2-a295c31ff464"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("8246ba7f-caaf-45fa-9b21-26697903e383"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("fef048c6-c86f-40d6-84f5-45c3c7a80317"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "fef048c6-c86f-40d6-84f5-45c3c7a80317",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-12T19:11:23.443Z"),
            WebURL: unifiedgosdk.Pointer("https://nervous-deck.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreTable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDatastoreTableRequest](../../pkg/models/operations/patchdatastoretablerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDatastoreTableResponse](../../pkg/models/operations/patchdatastoretableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveDatastoreTable

Remove a table

### Example Usage

<!-- UsageSnippet language="go" operationID="removeDatastoreTable" method="delete" path="/datastore/{connection_id}/table/{id}" -->
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

    res, err := s.Table.RemoveDatastoreTable(ctx, operations.RemoveDatastoreTableRequest{
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
| `request`                                                                                            | [operations.RemoveDatastoreTableRequest](../../pkg/models/operations/removedatastoretablerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveDatastoreTableResponse](../../pkg/models/operations/removedatastoretableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDatastoreTable

Update a table

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatastoreTable" method="put" path="/datastore/{connection_id}/table/{id}" example="datastore_table" -->
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

    res, err := s.Table.UpdateDatastoreTable(ctx, operations.UpdateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("7f4c87fc-0bfe-4866-a6d4-d38955c97483"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("7eb13a26-f2ac-4a25-acfa-c84b428434e7"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("76f7cee1-2b36-4a0d-b321-03d9f3b1fb7e"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "Toy - Kiehn",
                    OriginalType: "sursum",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Audax articulus triumphus tametsi cura tenetur colligo arcesso bestia."),
                    ID: unifiedgosdk.Pointer("4a4fccf3-5cf4-4ead-b0f2-a295c31ff464"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("8246ba7f-caaf-45fa-9b21-26697903e383"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("fef048c6-c86f-40d6-84f5-45c3c7a80317"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "fef048c6-c86f-40d6-84f5-45c3c7a80317",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-12T19:11:23.443Z"),
            WebURL: unifiedgosdk.Pointer("https://nervous-deck.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreTable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDatastoreTableRequest](../../pkg/models/operations/updatedatastoretablerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDatastoreTableResponse](../../pkg/models/operations/updatedatastoretableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |