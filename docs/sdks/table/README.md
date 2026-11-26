# Table

## Overview

### Available Operations

* [CreateDatastoreTable](#createdatastoretable) - Create a table
* [GetCommerceTable](#getcommercetable) - Retrieve a table
* [GetDatastoreTable](#getdatastoretable) - Retrieve a table
* [ListCommerceTables](#listcommercetables) - List all tables
* [ListDatastoreTables](#listdatastoretables) - List all tables
* [PatchDatastoreTable](#patchdatastoretable) - Update a table
* [RemoveDatastoreTable](#removedatastoretable) - Remove a table
* [UpdateDatastoreTable](#updatedatastoretable) - Update a table

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

    res, err := s.Table.CreateDatastoreTable(ctx, operations.CreateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Carpo vomito magni tantum tubineus corona deputo."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Basium accendo conservo curto thesis amet supellex."),
                    ID: unifiedgosdk.Pointer("5f356920-8d57-4cf3-9c32-3042c999c1b8"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Cassin - Greenfelder",
                    OriginalType: "tenus",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Alo admiratio vallum decet tamquam utilis callide adhuc verumtamen."),
                    ID: unifiedgosdk.Pointer("c1d380ef-add8-44f1-bf93-1b0b12364eb2"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Gerhold, Hermiston and Schmidt",
                    OriginalType: "doloremque",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Celer charisma coniuratio usus audax articulus triumphus tametsi cura."),
                    ID: unifiedgosdk.Pointer("ac2c0251-9b38-4dc8-b372-acaed346b938"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "McDermott LLC",
                    OriginalType: "tenetur",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Adficio voluptatibus ciminatio."),
                    ID: unifiedgosdk.Pointer("20335e1f-53f3-47f5-a1c7-2ad36f3347de"),
                    Name: "Koch - Krajcik",
                    OriginalType: "umbra",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Demitto ullam solium bellicus urbs concedo."),
                    ID: unifiedgosdk.Pointer("48c44851-3561-4c1a-98ca-4cb13d871e3c"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Kling - Bosco",
                    OriginalType: "venustas",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("e9ee7aa4-1a15-432c-91d0-18e315fecc87"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Bauch LLC"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Nikolaus - Kub",
                    Name: "Brakus and Sons",
                    TableID: "e9ee7aa4-1a15-432c-91d0-18e315fecc87",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-14T16:38:36.038Z"),
            WebURL: unifiedgosdk.Pointer("https://torn-bench.biz/"),
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

## GetCommerceTable

Retrieves a single CommerceTable by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceTable" method="get" path="/commerce/{connection_id}/table/{id}" -->
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

    res, err := s.Table.GetCommerceTable(ctx, operations.GetCommerceTableRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceTable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCommerceTableRequest](../../pkg/models/operations/getcommercetablerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetCommerceTableResponse](../../pkg/models/operations/getcommercetableresponse.md), error**

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

## ListCommerceTables

Returns a list of CommerceTables. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceTables" method="get" path="/commerce/{connection_id}/table" -->
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

    res, err := s.Table.ListCommerceTables(ctx, operations.ListCommerceTablesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceTables != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCommerceTablesRequest](../../pkg/models/operations/listcommercetablesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListCommerceTablesResponse](../../pkg/models/operations/listcommercetablesresponse.md), error**

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

    res, err := s.Table.PatchDatastoreTable(ctx, operations.PatchDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Carpo vomito magni tantum tubineus corona deputo."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Basium accendo conservo curto thesis amet supellex."),
                    ID: unifiedgosdk.Pointer("0c22e0df-531a-49f4-8744-3aa146c9fd85"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Cassin - Greenfelder",
                    OriginalType: "tenus",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Alo admiratio vallum decet tamquam utilis callide adhuc verumtamen."),
                    ID: unifiedgosdk.Pointer("fb66dcd7-319f-4c14-a3ff-71e874b28cf3"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Gerhold, Hermiston and Schmidt",
                    OriginalType: "doloremque",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Celer charisma coniuratio usus audax articulus triumphus tametsi cura."),
                    ID: unifiedgosdk.Pointer("2260a2fa-9e84-4909-97bd-8d2212731364"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "McDermott LLC",
                    OriginalType: "tenetur",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Adficio voluptatibus ciminatio."),
                    ID: unifiedgosdk.Pointer("c3be7703-5845-4731-879a-0ada3cb8e3ea"),
                    Name: "Koch - Krajcik",
                    OriginalType: "umbra",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Demitto ullam solium bellicus urbs concedo."),
                    ID: unifiedgosdk.Pointer("c6fa7111-ddd0-41bd-b114-c5adb39f62c2"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Kling - Bosco",
                    OriginalType: "venustas",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("bd6693ee-a7d0-4ca1-89d5-8055745e13e3"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Bauch LLC"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Nikolaus - Kub",
                    Name: "Brakus and Sons",
                    TableID: "bd6693ee-a7d0-4ca1-89d5-8055745e13e3",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-14T16:38:36.040Z"),
            WebURL: unifiedgosdk.Pointer("https://torn-bench.biz/"),
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

    res, err := s.Table.UpdateDatastoreTable(ctx, operations.UpdateDatastoreTableRequest{
        DatastoreTable: shared.DatastoreTable{
            CreatedAt: types.MustNewTimeFromString("2022-05-22T04:49:00.762Z"),
            Description: unifiedgosdk.Pointer("Carpo vomito magni tantum tubineus corona deputo."),
            Fields: []shared.DatastoreField{
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Basium accendo conservo curto thesis amet supellex."),
                    ID: unifiedgosdk.Pointer("0c22e0df-531a-49f4-8744-3aa146c9fd85"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Cassin - Greenfelder",
                    OriginalType: "tenus",
                    Scale: unifiedgosdk.Pointer[float64](64.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Alo admiratio vallum decet tamquam utilis callide adhuc verumtamen."),
                    ID: unifiedgosdk.Pointer("fb66dcd7-319f-4c14-a3ff-71e874b28cf3"),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    Name: "Gerhold, Hermiston and Schmidt",
                    OriginalType: "doloremque",
                    Precision: unifiedgosdk.Pointer[float64](57.0),
                    Type: shared.DatastoreFieldTypeSingleSelect,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Celer charisma coniuratio usus audax articulus triumphus tametsi cura."),
                    ID: unifiedgosdk.Pointer("2260a2fa-9e84-4909-97bd-8d2212731364"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsNullable: unifiedgosdk.Pointer(true),
                    IsUnique: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](62.0),
                    Name: "McDermott LLC",
                    OriginalType: "tenetur",
                    Precision: unifiedgosdk.Pointer[float64](47.0),
                    Type: shared.DatastoreFieldTypeNumber,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Adficio voluptatibus ciminatio."),
                    ID: unifiedgosdk.Pointer("c3be7703-5845-4731-879a-0ada3cb8e3ea"),
                    Name: "Koch - Krajcik",
                    OriginalType: "umbra",
                    Precision: unifiedgosdk.Pointer[float64](94.0),
                    Scale: unifiedgosdk.Pointer[float64](24.0),
                    Type: shared.DatastoreFieldTypeRelation,
                },
                shared.DatastoreField{
                    Description: unifiedgosdk.Pointer("Demitto ullam solium bellicus urbs concedo."),
                    ID: unifiedgosdk.Pointer("c6fa7111-ddd0-41bd-b114-c5adb39f62c2"),
                    IsIndexed: unifiedgosdk.Pointer(true),
                    IsPrimaryKey: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    Name: "Kling - Bosco",
                    OriginalType: "venustas",
                    Scale: unifiedgosdk.Pointer[float64](68.0),
                    Type: shared.DatastoreFieldTypeEmail,
                },
            },
            ID: unifiedgosdk.Pointer("bd6693ee-a7d0-4ca1-89d5-8055745e13e3"),
            Metadata: []shared.DatastoreMetadata{},
            Name: unifiedgosdk.Pointer("Bauch LLC"),
            Relationships: []shared.DatastoreRelationship{
                shared.DatastoreRelationship{
                    Field: "Nikolaus - Kub",
                    Name: "Brakus and Sons",
                    TableID: "bd6693ee-a7d0-4ca1-89d5-8055745e13e3",
                    Type: shared.DatastoreRelationshipTypeMany,
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-08-14T16:38:36.040Z"),
            WebURL: unifiedgosdk.Pointer("https://torn-bench.biz/"),
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