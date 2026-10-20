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

Create a database

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
            ID: unifiedgosdk.Pointer("301611d6-fba7-4eac-bd5f-e5bc75a0e2bb"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-30T06:05:59.540Z"),
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

Create a query

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

Create a record

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
            ID: unifiedgosdk.Pointer("bbe76082-d277-45d5-9790-7c910df8940c"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-04T12:10:24.138Z"),
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

    res, err := s.Datastore.CreateDatastoreTable(ctx, operations.CreateDatastoreTableRequest{
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

## GetDatastoreDatabase

Retrieve a database

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

Retrieve a record

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

List all databases

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

List all records

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

Update a database

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
            ID: unifiedgosdk.Pointer("9104b90f-f1ed-458d-bc5e-d563986c2a8e"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-30T06:05:59.548Z"),
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

Update a record

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
            ID: unifiedgosdk.Pointer("32245e8a-fc37-4310-980f-7ed5ffce5a56"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-04T12:10:24.146Z"),
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

    res, err := s.Datastore.PatchDatastoreTable(ctx, operations.PatchDatastoreTableRequest{
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

## RemoveDatastoreDatabase

Remove a database

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

Remove a record

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

Update a database

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
            ID: unifiedgosdk.Pointer("9104b90f-f1ed-458d-bc5e-d563986c2a8e"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Rau - O'Reilly"),
            Password: unifiedgosdk.Pointer("iYLhYwOHm0NLDvI"),
            Region: unifiedgosdk.Pointer("France"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-30T06:05:59.548Z"),
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

Update a record

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
            ID: unifiedgosdk.Pointer("32245e8a-fc37-4310-980f-7ed5ffce5a56"),
            Metadata: []shared.DatastoreMetadata{},
            RowNumber: unifiedgosdk.Pointer[float64](33.0),
            UpdatedAt: types.MustNewTimeFromString("2025-03-04T12:10:24.146Z"),
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

    res, err := s.Datastore.UpdateDatastoreTable(ctx, operations.UpdateDatastoreTableRequest{
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