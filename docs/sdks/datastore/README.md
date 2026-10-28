# Datastore

## Overview

### Available Operations

* [CreateDatastoreDatabase](#createdatastoredatabase) - Create a database
* [CreateDatastoreQuery](#createdatastorequery) - Create a query
* [CreateDatastoreRecord](#createdatastorerecord) - Create a record
* [CreateDatastoreTable](#createdatastoretable) - Create a table
* [GetDatastoreDatabase](#getdatastoredatabase) - Retrieve a database
* [GetDatastoreRecord](#getdatastorerecord) - Retrieve a record
* [GetDatastoreTable](#getdatastoretable) - Retrieve a table
* [ListDatastoreDatabases](#listdatastoredatabases) - List all databases
* [ListDatastoreRecords](#listdatastorerecords) - List all records
* [ListDatastoreTables](#listdatastoretables) - List all tables
* [PatchDatastoreDatabase](#patchdatastoredatabase) - Update a database
* [PatchDatastoreRecord](#patchdatastorerecord) - Update a record
* [PatchDatastoreTable](#patchdatastoretable) - Update a table
* [RemoveDatastoreDatabase](#removedatastoredatabase) - Remove a database
* [RemoveDatastoreRecord](#removedatastorerecord) - Remove a record
* [RemoveDatastoreTable](#removedatastoretable) - Remove a table
* [UpdateDatastoreDatabase](#updatedatastoredatabase) - Update a database
* [UpdateDatastoreRecord](#updatedatastorerecord) - Update a record
* [UpdateDatastoreTable](#updatedatastoretable) - Update a table

## CreateDatastoreDatabase

Creates a DatastoreDatabase. Provide the unified DatastoreDatabase fields in the request body; they are mapped to the provider's API. Returns the created DatastoreDatabase.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDatastoreDatabase" method="post" path="/datastore/{connection_id}/database" example="datastore_database" -->
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

    res, err := s.Datastore.CreateDatastoreDatabase(ctx, operations.CreateDatastoreDatabaseRequest{
        DatastoreDatabase: shared.DatastoreDatabase{
            CreatedAt: types.MustNewTimeFromString("2019-03-05T14:31:00.223Z"),
            Description: unifiedgosdk.Pointer("Substantia tantillus coniuratio repellat coerceo amet vito vorago."),
            ID: unifiedgosdk.Pointer("6de8910c-48b2-4e9a-b4db-7b4552a13bfb"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-04T19:26:18.791Z"),
            WebURL: unifiedgosdk.Pointer("https://optimal-signature.name"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateDatastoreDatabaseRequest](../../pkg/models/operations/createdatastoredatabaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateDatastoreDatabaseResponse](../../pkg/models/operations/createdatastoredatabaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDatastoreQuery

Creates a DatastoreQuery. Provide the unified DatastoreQuery fields in the request body; they are mapped to the provider's API. Returns the created DatastoreQuery.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDatastoreQuery" method="post" path="/datastore/{connection_id}/query" example="datastore_query" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Datastore.CreateDatastoreQuery(ctx, operations.CreateDatastoreQueryRequest{
        DatastoreQuery: shared.DatastoreQuery{
            Metadata: []shared.DatastoreMetadata{},
            Query: &shared.PropertyDatastoreQueryQuery{
                Filter: &shared.PropertyDatastoreQueryQueryFilter{
                    Type: shared.PropertyDatastoreQueryQueryFilterTypeOr,
                },
                Select: []string{
                    "*",
                },
                SQL: unifiedgosdk.Pointer(""),
            },
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDatastoreQueryRequest](../../pkg/models/operations/createdatastorequeryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDatastoreQueryResponse](../../pkg/models/operations/createdatastorequeryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDatastoreRecord

Creates a DatastoreRecord. Provide the unified DatastoreRecord fields in the request body; they are mapped to the provider's API. Returns the created DatastoreRecord.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDatastoreRecord" method="post" path="/datastore/{connection_id}/record" example="datastore_record" -->
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

    res, err := s.Datastore.CreateDatastoreRecord(ctx, operations.CreateDatastoreRecordRequest{
        DatastoreRecord: shared.DatastoreRecord{
            CreatedAt: types.MustNewTimeFromString("2022-07-10T04:42:03.076Z"),
            Fields: map[string]shared.DatastoreFieldValue{

            },
            ID: unifiedgosdk.Pointer("ea88cc46-f4f3-4266-b887-b32bf35aa8a7"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-09T11:46:56.070Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateDatastoreRecordRequest](../../pkg/models/operations/createdatastorerecordrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateDatastoreRecordResponse](../../pkg/models/operations/createdatastorerecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDatastoreTable

Creates a DatastoreTable. Provide the unified DatastoreTable fields in the request body; they are mapped to the provider's API. Returns the created DatastoreTable.

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

    res, err := s.Datastore.CreateDatastoreTable(ctx, operations.CreateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("69e1c552-b2f9-4b5d-a9d4-7071fbde343f"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("bda5b71b-5c13-4a31-ae50-880ffaf629ec"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("89a28ee6-ea48-426b-b1d4-4e4749cc7b97"),
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
                    ID: unifiedgosdk.Pointer("2cc8b8f6-9d97-4d0e-a327-30587673ac70"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("f0f9aec2-be40-4fee-aca8-1d72dd85f4b9"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("fb487132-c19b-4923-ad8b-99cd79d85954"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "fb487132-c19b-4923-ad8b-99cd79d85954",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-13T05:04:57.248Z"),
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

## GetDatastoreDatabase

Retrieves a single DatastoreDatabase by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatastoreDatabase" method="get" path="/datastore/{connection_id}/database/{id}" -->
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

    res, err := s.Datastore.GetDatastoreDatabase(ctx, operations.GetDatastoreDatabaseRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetDatastoreDatabaseRequest](../../pkg/models/operations/getdatastoredatabaserequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetDatastoreDatabaseResponse](../../pkg/models/operations/getdatastoredatabaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDatastoreRecord

Retrieves a single DatastoreRecord by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDatastoreRecord" method="get" path="/datastore/{connection_id}/record/{id}" -->
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

    res, err := s.Datastore.GetDatastoreRecord(ctx, operations.GetDatastoreRecordRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetDatastoreRecordRequest](../../pkg/models/operations/getdatastorerecordrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetDatastoreRecordResponse](../../pkg/models/operations/getdatastorerecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDatastoreTable

Retrieves a single DatastoreTable by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

    res, err := s.Datastore.GetDatastoreTable(ctx, operations.GetDatastoreTableRequest{
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

## ListDatastoreDatabases

Returns a list of DatastoreDatabases. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDatastoreDatabases" method="get" path="/datastore/{connection_id}/database" -->
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

    res, err := s.Datastore.ListDatastoreDatabases(ctx, operations.ListDatastoreDatabasesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreDatabases != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListDatastoreDatabasesRequest](../../pkg/models/operations/listdatastoredatabasesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListDatastoreDatabasesResponse](../../pkg/models/operations/listdatastoredatabasesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDatastoreRecords

Returns a list of DatastoreRecords. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDatastoreRecords" method="get" path="/datastore/{connection_id}/record" -->
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

    res, err := s.Datastore.ListDatastoreRecords(ctx, operations.ListDatastoreRecordsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreRecords != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListDatastoreRecordsRequest](../../pkg/models/operations/listdatastorerecordsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListDatastoreRecordsResponse](../../pkg/models/operations/listdatastorerecordsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDatastoreTables

Returns a list of DatastoreTables. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Datastore.ListDatastoreTables(ctx, operations.ListDatastoreTablesRequest{
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

## PatchDatastoreDatabase

Updates the DatastoreDatabase identified by `id`. Provide the unified DatastoreDatabase fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDatastoreDatabase" method="patch" path="/datastore/{connection_id}/database/{id}" example="datastore_database" -->
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

    res, err := s.Datastore.PatchDatastoreDatabase(ctx, operations.PatchDatastoreDatabaseRequest{
        DatastoreDatabase: shared.DatastoreDatabase{
            CreatedAt: types.MustNewTimeFromString("2019-03-05T14:31:00.223Z"),
            Description: unifiedgosdk.Pointer("Substantia tantillus coniuratio repellat coerceo amet vito vorago."),
            ID: unifiedgosdk.Pointer("04c17a34-c526-4614-81ba-2f6d159a5ca6"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-04T19:26:18.795Z"),
            WebURL: unifiedgosdk.Pointer("https://optimal-signature.name"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchDatastoreDatabaseRequest](../../pkg/models/operations/patchdatastoredatabaserequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchDatastoreDatabaseResponse](../../pkg/models/operations/patchdatastoredatabaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDatastoreRecord

Updates the DatastoreRecord identified by `id`. Provide the unified DatastoreRecord fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDatastoreRecord" method="patch" path="/datastore/{connection_id}/record/{id}" example="datastore_record" -->
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

    res, err := s.Datastore.PatchDatastoreRecord(ctx, operations.PatchDatastoreRecordRequest{
        DatastoreRecord: shared.DatastoreRecord{
            CreatedAt: types.MustNewTimeFromString("2022-07-10T04:42:03.076Z"),
            Fields: map[string]shared.DatastoreFieldValue{

            },
            ID: unifiedgosdk.Pointer("6ec47a3e-be05-4464-8e3d-5fd0dde22ea5"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-09T11:46:56.078Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchDatastoreRecordRequest](../../pkg/models/operations/patchdatastorerecordrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchDatastoreRecordResponse](../../pkg/models/operations/patchdatastorerecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDatastoreTable

Updates the DatastoreTable identified by `id`. Provide the unified DatastoreTable fields to change in the request body.

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

    res, err := s.Datastore.PatchDatastoreTable(ctx, operations.PatchDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("22f8b807-8951-4664-b911-32041b098649"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("fa1cc0a1-7ae7-4964-b73d-e85b4ccb0be5"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("7d578bd9-814c-4781-b759-d50eef63021b"),
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
                    ID: unifiedgosdk.Pointer("00291bcc-b2ae-421d-a3d5-df6a66cb46aa"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("8b4d5903-19b9-494a-9a21-d1a80bd75bba"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("ca1c7dce-36cd-4ba5-9606-9e9b31552abb"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "ca1c7dce-36cd-4ba5-9606-9e9b31552abb",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-13T05:04:57.250Z"),
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

## RemoveDatastoreDatabase

Deletes the DatastoreDatabase identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeDatastoreDatabase" method="delete" path="/datastore/{connection_id}/database/{id}" -->
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

    res, err := s.Datastore.RemoveDatastoreDatabase(ctx, operations.RemoveDatastoreDatabaseRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveDatastoreDatabaseRequest](../../pkg/models/operations/removedatastoredatabaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveDatastoreDatabaseResponse](../../pkg/models/operations/removedatastoredatabaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveDatastoreRecord

Deletes the DatastoreRecord identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeDatastoreRecord" method="delete" path="/datastore/{connection_id}/record/{id}" -->
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

    res, err := s.Datastore.RemoveDatastoreRecord(ctx, operations.RemoveDatastoreRecordRequest{
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
| `request`                                                                                              | [operations.RemoveDatastoreRecordRequest](../../pkg/models/operations/removedatastorerecordrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveDatastoreRecordResponse](../../pkg/models/operations/removedatastorerecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveDatastoreTable

Deletes the DatastoreTable identified by `id`.

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

    res, err := s.Datastore.RemoveDatastoreTable(ctx, operations.RemoveDatastoreTableRequest{
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

## UpdateDatastoreDatabase

Updates the DatastoreDatabase identified by `id`. Provide the unified DatastoreDatabase fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatastoreDatabase" method="put" path="/datastore/{connection_id}/database/{id}" example="datastore_database" -->
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

    res, err := s.Datastore.UpdateDatastoreDatabase(ctx, operations.UpdateDatastoreDatabaseRequest{
        DatastoreDatabase: shared.DatastoreDatabase{
            CreatedAt: types.MustNewTimeFromString("2019-03-05T14:31:00.223Z"),
            Description: unifiedgosdk.Pointer("Substantia tantillus coniuratio repellat coerceo amet vito vorago."),
            ID: unifiedgosdk.Pointer("04c17a34-c526-4614-81ba-2f6d159a5ca6"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-04T19:26:18.795Z"),
            WebURL: unifiedgosdk.Pointer("https://optimal-signature.name"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateDatastoreDatabaseRequest](../../pkg/models/operations/updatedatastoredatabaserequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateDatastoreDatabaseResponse](../../pkg/models/operations/updatedatastoredatabaseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDatastoreRecord

Updates the DatastoreRecord identified by `id`. Provide the unified DatastoreRecord fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatastoreRecord" method="put" path="/datastore/{connection_id}/record/{id}" example="datastore_record" -->
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

    res, err := s.Datastore.UpdateDatastoreRecord(ctx, operations.UpdateDatastoreRecordRequest{
        DatastoreRecord: shared.DatastoreRecord{
            CreatedAt: types.MustNewTimeFromString("2022-07-10T04:42:03.076Z"),
            Fields: map[string]shared.DatastoreFieldValue{

            },
            ID: unifiedgosdk.Pointer("6ec47a3e-be05-4464-8e3d-5fd0dde22ea5"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-09T11:46:56.078Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatastoreRecord != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateDatastoreRecordRequest](../../pkg/models/operations/updatedatastorerecordrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateDatastoreRecordResponse](../../pkg/models/operations/updatedatastorerecordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDatastoreTable

Updates the DatastoreTable identified by `id`. Provide the unified DatastoreTable fields to change in the request body.

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

    res, err := s.Datastore.UpdateDatastoreTable(ctx, operations.UpdateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Quasi adhuc paulatim ducimus carpo vomito magni tantum tubineus corona."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Antea soluta itaque clamo despecto basium accendo."),
                    ID: unifiedgosdk.Pointer("22f8b807-8951-4664-b911-32041b098649"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Ferry LLC",
                    OriginalType: "conservo",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Nulla cattus aequus sumo neque alter cauda subseco."),
                    ID: unifiedgosdk.Pointer("fa1cc0a1-7ae7-4964-b73d-e85b4ccb0be5"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Hand - Nader",
                    OriginalType: "conservo",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Utilis callide adhuc verumtamen doloremque ceno sui sint."),
                    ID: unifiedgosdk.Pointer("7d578bd9-814c-4781-b759-d50eef63021b"),
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
                    ID: unifiedgosdk.Pointer("00291bcc-b2ae-421d-a3d5-df6a66cb46aa"),
                    Name: "Gislason, Gorczany and Hegmann",
                    OriginalType: "tego",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Communis coma cura audio aduro acer canis expedita."),
                    ID: unifiedgosdk.Pointer("8b4d5903-19b9-494a-9a21-d1a80bd75bba"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Graham, Steuber and Kshlerin",
                    OriginalType: "defessus",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("ca1c7dce-36cd-4ba5-9606-9e9b31552abb"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rogahn, Lowe and Heller"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Macejkovic - Shields",
                    Name: "Hyatt, Ryan and Nienow",
                    TableID: "ca1c7dce-36cd-4ba5-9606-9e9b31552abb",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-13T05:04:57.250Z"),
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