# Plm

## Overview

### Available Operations

* [CreatePlmBomline](#createplmbomline) - Create a bomline
* [CreatePlmChange](#createplmchange) - Create a change
* [CreatePlmDocument](#createplmdocument) - Create a document
* [CreatePlmPart](#createplmpart) - Create a part
* [GetPlmBomline](#getplmbomline) - Retrieve a bomline
* [GetPlmCategory](#getplmcategory) - Retrieve a category
* [GetPlmChange](#getplmchange) - Retrieve a change
* [GetPlmDocument](#getplmdocument) - Retrieve a document
* [GetPlmPart](#getplmpart) - Retrieve a part
* [ListPlmBomlines](#listplmbomlines) - List all bomlines
* [ListPlmCategories](#listplmcategories) - List all categories
* [ListPlmChanges](#listplmchanges) - List all changes
* [ListPlmDocuments](#listplmdocuments) - List all documents
* [ListPlmParts](#listplmparts) - List all parts
* [PatchPlmBomline](#patchplmbomline) - Update a bomline
* [PatchPlmChange](#patchplmchange) - Update a change
* [PatchPlmDocument](#patchplmdocument) - Update a document
* [PatchPlmPart](#patchplmpart) - Update a part
* [RemovePlmBomline](#removeplmbomline) - Remove a bomline
* [UpdatePlmBomline](#updateplmbomline) - Update a bomline
* [UpdatePlmChange](#updateplmchange) - Update a change
* [UpdatePlmDocument](#updateplmdocument) - Update a document
* [UpdatePlmPart](#updateplmpart) - Update a part

## CreatePlmBomline

Creates a PlmBomline. Provide the unified PlmBomline fields in the request body; they are mapped to the provider's API. Returns the created PlmBomline.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlmBomline" method="post" path="/plm/{connection_id}/bomline" example="plm_bomline" -->
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

    res, err := s.Plm.CreatePlmBomline(ctx, operations.CreatePlmBomlineRequest{
        PlmBomline: shared.PlmBomline{
            CreatedAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            ID: unifiedgosdk.Pointer("08d64988-c488-4fb9-b151-406341e08276"),
            Quantity: unifiedgosdk.Pointer[float64](7.0),
            ReferenceDesignators: []string{},
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2023-04-19T23:28:58.887Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmBomline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreatePlmBomlineRequest](../../pkg/models/operations/createplmbomlinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreatePlmBomlineResponse](../../pkg/models/operations/createplmbomlineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePlmChange

Creates a PlmChange. Provide the unified PlmChange fields in the request body; they are mapped to the provider's API. Returns the created PlmChange.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlmChange" method="post" path="/plm/{connection_id}/change" example="plm_change" -->
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

    res, err := s.Plm.CreatePlmChange(ctx, operations.CreatePlmChangeRequest{
        PlmChange: shared.PlmChange{
            AffectedItems: []shared.PlmAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2021-08-02T05:58:32.287Z"),
            Description: unifiedgosdk.Pointer("Vorago spiculum avaritia adipiscor callide."),
            ID: unifiedgosdk.Pointer("dba9fa58-f5e6-4b05-939a-c5eece7dc300"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Update goodwill"),
            Number: unifiedgosdk.Pointer("ECO-6591"),
            Priority: unifiedgosdk.Pointer("Medium"),
            Status: shared.PlmChangeStatusInReview.ToPointer(),
            Type: shared.PlmChangeTypeOrder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-27T10:41:41.592Z"),
            URL: unifiedgosdk.Pointer("https://warmhearted-fraudster.biz/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmChange != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePlmChangeRequest](../../pkg/models/operations/createplmchangerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreatePlmChangeResponse](../../pkg/models/operations/createplmchangeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePlmDocument

Creates a PlmDocument. Provide the unified PlmDocument fields in the request body; they are mapped to the provider's API. Returns the created PlmDocument.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlmDocument" method="post" path="/plm/{connection_id}/document" example="plm_document" -->
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

    res, err := s.Plm.CreatePlmDocument(ctx, operations.CreatePlmDocumentRequest{
        PlmDocument: shared.PlmDocument{
            CreatedAt: types.MustNewTimeFromString("2020-10-16T01:44:30.293Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("abdb4294-7d62-43ea-abdc-0bf84c6ba1b8"),
            LifecycleState: unifiedgosdk.Pointer("Under Review"),
            MasterIdentifier: unifiedgosdk.Pointer("2f469643-9036-4654-86d8-f2a20c9f1ef9"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Procedure taxicab"),
            Number: unifiedgosdk.Pointer("DOC-086886"),
            OriginalType: unifiedgosdk.Pointer("Procedure"),
            Revision: unifiedgosdk.Pointer("C"),
            Status: shared.PlmDocumentStatusInReview.ToPointer(),
            Type: shared.PlmDocumentTypeProcedure.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-08T04:07:37.607Z"),
            URL: unifiedgosdk.Pointer("https://uncommon-guide.info/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreatePlmDocumentRequest](../../pkg/models/operations/createplmdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreatePlmDocumentResponse](../../pkg/models/operations/createplmdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePlmPart

Creates a PlmPart. Provide the unified PlmPart fields in the request body; they are mapped to the provider's API. Returns the created PlmPart.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlmPart" method="post" path="/plm/{connection_id}/part" example="plm_part" -->
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

    res, err := s.Plm.CreatePlmPart(ctx, operations.CreatePlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("619441a7-5f2b-40ba-8137-db9e630d546b"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("28ab7e05-ee8d-4943-87ac-76fd00ed173e"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-12T10:26:14.684Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreatePlmPartRequest](../../pkg/models/operations/createplmpartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreatePlmPartResponse](../../pkg/models/operations/createplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmBomline

Retrieves a single PlmBomline by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmBomline" method="get" path="/plm/{connection_id}/bomline/{id}" -->
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

    res, err := s.Plm.GetPlmBomline(ctx, operations.GetPlmBomlineRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmBomline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetPlmBomlineRequest](../../pkg/models/operations/getplmbomlinerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetPlmBomlineResponse](../../pkg/models/operations/getplmbomlineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmCategory

Retrieves a single PlmCategory by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmCategory" method="get" path="/plm/{connection_id}/category/{id}" -->
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

    res, err := s.Plm.GetPlmCategory(ctx, operations.GetPlmCategoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPlmCategoryRequest](../../pkg/models/operations/getplmcategoryrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetPlmCategoryResponse](../../pkg/models/operations/getplmcategoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmChange

Retrieves a single PlmChange by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmChange" method="get" path="/plm/{connection_id}/change/{id}" -->
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

    res, err := s.Plm.GetPlmChange(ctx, operations.GetPlmChangeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmChange != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPlmChangeRequest](../../pkg/models/operations/getplmchangerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetPlmChangeResponse](../../pkg/models/operations/getplmchangeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmDocument

Retrieves a single PlmDocument by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmDocument" method="get" path="/plm/{connection_id}/document/{id}" -->
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

    res, err := s.Plm.GetPlmDocument(ctx, operations.GetPlmDocumentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPlmDocumentRequest](../../pkg/models/operations/getplmdocumentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetPlmDocumentResponse](../../pkg/models/operations/getplmdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlmPart

Retrieves a single PlmPart by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlmPart" method="get" path="/plm/{connection_id}/part/{id}" -->
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

    res, err := s.Plm.GetPlmPart(ctx, operations.GetPlmPartRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetPlmPartRequest](../../pkg/models/operations/getplmpartrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetPlmPartResponse](../../pkg/models/operations/getplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmBomlines

Returns a list of PlmBomlines. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmBomlines" method="get" path="/plm/{connection_id}/bomline" -->
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

    res, err := s.Plm.ListPlmBomlines(ctx, operations.ListPlmBomlinesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmBomlines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListPlmBomlinesRequest](../../pkg/models/operations/listplmbomlinesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListPlmBomlinesResponse](../../pkg/models/operations/listplmbomlinesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmCategories

Returns a list of PlmCategories. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmCategories" method="get" path="/plm/{connection_id}/category" -->
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

    res, err := s.Plm.ListPlmCategories(ctx, operations.ListPlmCategoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmCategories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListPlmCategoriesRequest](../../pkg/models/operations/listplmcategoriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListPlmCategoriesResponse](../../pkg/models/operations/listplmcategoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmChanges

Returns a list of PlmChanges. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmChanges" method="get" path="/plm/{connection_id}/change" -->
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

    res, err := s.Plm.ListPlmChanges(ctx, operations.ListPlmChangesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmChanges != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListPlmChangesRequest](../../pkg/models/operations/listplmchangesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListPlmChangesResponse](../../pkg/models/operations/listplmchangesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmDocuments

Returns a list of PlmDocuments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmDocuments" method="get" path="/plm/{connection_id}/document" -->
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

    res, err := s.Plm.ListPlmDocuments(ctx, operations.ListPlmDocumentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmDocuments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPlmDocumentsRequest](../../pkg/models/operations/listplmdocumentsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListPlmDocumentsResponse](../../pkg/models/operations/listplmdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPlmParts

Returns a list of PlmParts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPlmParts" method="get" path="/plm/{connection_id}/part" -->
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

    res, err := s.Plm.ListPlmParts(ctx, operations.ListPlmPartsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmParts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPlmPartsRequest](../../pkg/models/operations/listplmpartsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPlmPartsResponse](../../pkg/models/operations/listplmpartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPlmBomline

Updates the PlmBomline identified by `id`. Provide the unified PlmBomline fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPlmBomline" method="patch" path="/plm/{connection_id}/bomline/{id}" example="plm_bomline" -->
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

    res, err := s.Plm.PatchPlmBomline(ctx, operations.PatchPlmBomlineRequest{
        PlmBomline: shared.PlmBomline{
            CreatedAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            ID: unifiedgosdk.Pointer("74f62a25-cd60-46bc-9f6f-0e34f205b352"),
            Quantity: unifiedgosdk.Pointer[float64](7.0),
            ReferenceDesignators: []string{},
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2023-04-19T23:28:58.889Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmBomline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchPlmBomlineRequest](../../pkg/models/operations/patchplmbomlinerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchPlmBomlineResponse](../../pkg/models/operations/patchplmbomlineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPlmChange

Updates the PlmChange identified by `id`. Provide the unified PlmChange fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPlmChange" method="patch" path="/plm/{connection_id}/change/{id}" example="plm_change" -->
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

    res, err := s.Plm.PatchPlmChange(ctx, operations.PatchPlmChangeRequest{
        PlmChange: shared.PlmChange{
            AffectedItems: []shared.PlmAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2021-08-02T05:58:32.287Z"),
            Description: unifiedgosdk.Pointer("Vorago spiculum avaritia adipiscor callide."),
            ID: unifiedgosdk.Pointer("5a328ef7-0eab-4dfb-92c6-f8ca8a6339ed"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Update goodwill"),
            Number: unifiedgosdk.Pointer("ECO-6591"),
            Priority: unifiedgosdk.Pointer("Medium"),
            Status: shared.PlmChangeStatusInReview.ToPointer(),
            Type: shared.PlmChangeTypeOrder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-27T10:41:41.609Z"),
            URL: unifiedgosdk.Pointer("https://warmhearted-fraudster.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmChange != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchPlmChangeRequest](../../pkg/models/operations/patchplmchangerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchPlmChangeResponse](../../pkg/models/operations/patchplmchangeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPlmDocument

Updates the PlmDocument identified by `id`. Provide the unified PlmDocument fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPlmDocument" method="patch" path="/plm/{connection_id}/document/{id}" example="plm_document" -->
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

    res, err := s.Plm.PatchPlmDocument(ctx, operations.PatchPlmDocumentRequest{
        PlmDocument: shared.PlmDocument{
            CreatedAt: types.MustNewTimeFromString("2020-10-16T01:44:30.293Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("84873777-7152-4e75-875b-0328ea6c2123"),
            LifecycleState: unifiedgosdk.Pointer("Under Review"),
            MasterIdentifier: unifiedgosdk.Pointer("51902057-d51a-495c-b146-0f3ddff647e4"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Procedure taxicab"),
            Number: unifiedgosdk.Pointer("DOC-086886"),
            OriginalType: unifiedgosdk.Pointer("Procedure"),
            Revision: unifiedgosdk.Pointer("C"),
            Status: shared.PlmDocumentStatusInReview.ToPointer(),
            Type: shared.PlmDocumentTypeProcedure.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-08T04:07:37.616Z"),
            URL: unifiedgosdk.Pointer("https://uncommon-guide.info/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchPlmDocumentRequest](../../pkg/models/operations/patchplmdocumentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchPlmDocumentResponse](../../pkg/models/operations/patchplmdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPlmPart

Updates the PlmPart identified by `id`. Provide the unified PlmPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPlmPart" method="patch" path="/plm/{connection_id}/part/{id}" example="plm_part" -->
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

    res, err := s.Plm.PatchPlmPart(ctx, operations.PatchPlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("4272fc6f-d650-4593-88de-8d62deee8aef"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("099afe83-64ef-4ce4-8dce-c5c0d4273d54"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-12T10:26:14.689Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchPlmPartRequest](../../pkg/models/operations/patchplmpartrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchPlmPartResponse](../../pkg/models/operations/patchplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePlmBomline

Deletes the PlmBomline identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePlmBomline" method="delete" path="/plm/{connection_id}/bomline/{id}" -->
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

    res, err := s.Plm.RemovePlmBomline(ctx, operations.RemovePlmBomlineRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemovePlmBomlineRequest](../../pkg/models/operations/removeplmbomlinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemovePlmBomlineResponse](../../pkg/models/operations/removeplmbomlineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlmBomline

Updates the PlmBomline identified by `id`. Provide the unified PlmBomline fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlmBomline" method="put" path="/plm/{connection_id}/bomline/{id}" example="plm_bomline" -->
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

    res, err := s.Plm.UpdatePlmBomline(ctx, operations.UpdatePlmBomlineRequest{
        PlmBomline: shared.PlmBomline{
            CreatedAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            EffectiveAt: types.MustNewTimeFromString("2022-04-02T07:09:29.598Z"),
            ID: unifiedgosdk.Pointer("74f62a25-cd60-46bc-9f6f-0e34f205b352"),
            Quantity: unifiedgosdk.Pointer[float64](7.0),
            ReferenceDesignators: []string{},
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2023-04-19T23:28:58.889Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmBomline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdatePlmBomlineRequest](../../pkg/models/operations/updateplmbomlinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdatePlmBomlineResponse](../../pkg/models/operations/updateplmbomlineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlmChange

Updates the PlmChange identified by `id`. Provide the unified PlmChange fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlmChange" method="put" path="/plm/{connection_id}/change/{id}" example="plm_change" -->
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

    res, err := s.Plm.UpdatePlmChange(ctx, operations.UpdatePlmChangeRequest{
        PlmChange: shared.PlmChange{
            AffectedItems: []shared.PlmAffectedItem{},
            CreatedAt: types.MustNewTimeFromString("2021-08-02T05:58:32.287Z"),
            Description: unifiedgosdk.Pointer("Vorago spiculum avaritia adipiscor callide."),
            ID: unifiedgosdk.Pointer("5a328ef7-0eab-4dfb-92c6-f8ca8a6339ed"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Update goodwill"),
            Number: unifiedgosdk.Pointer("ECO-6591"),
            Priority: unifiedgosdk.Pointer("Medium"),
            Status: shared.PlmChangeStatusInReview.ToPointer(),
            Type: shared.PlmChangeTypeOrder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-27T10:41:41.609Z"),
            URL: unifiedgosdk.Pointer("https://warmhearted-fraudster.biz/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmChange != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdatePlmChangeRequest](../../pkg/models/operations/updateplmchangerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdatePlmChangeResponse](../../pkg/models/operations/updateplmchangeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlmDocument

Updates the PlmDocument identified by `id`. Provide the unified PlmDocument fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlmDocument" method="put" path="/plm/{connection_id}/document/{id}" example="plm_document" -->
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

    res, err := s.Plm.UpdatePlmDocument(ctx, operations.UpdatePlmDocumentRequest{
        PlmDocument: shared.PlmDocument{
            CreatedAt: types.MustNewTimeFromString("2020-10-16T01:44:30.293Z"),
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("84873777-7152-4e75-875b-0328ea6c2123"),
            LifecycleState: unifiedgosdk.Pointer("Under Review"),
            MasterIdentifier: unifiedgosdk.Pointer("51902057-d51a-495c-b146-0f3ddff647e4"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("Procedure taxicab"),
            Number: unifiedgosdk.Pointer("DOC-086886"),
            OriginalType: unifiedgosdk.Pointer("Procedure"),
            Revision: unifiedgosdk.Pointer("C"),
            Status: shared.PlmDocumentStatusInReview.ToPointer(),
            Type: shared.PlmDocumentTypeProcedure.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-09-08T04:07:37.616Z"),
            URL: unifiedgosdk.Pointer("https://uncommon-guide.info/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePlmDocumentRequest](../../pkg/models/operations/updateplmdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdatePlmDocumentResponse](../../pkg/models/operations/updateplmdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlmPart

Updates the PlmPart identified by `id`. Provide the unified PlmPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlmPart" method="put" path="/plm/{connection_id}/part/{id}" example="plm_part" -->
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

    res, err := s.Plm.UpdatePlmPart(ctx, operations.UpdatePlmPartRequest{
        PlmPart: shared.PlmPart{
            CreatedAt: types.MustNewTimeFromString("2019-04-13T06:25:34.770Z"),
            Description: unifiedgosdk.Pointer("Pectus tero ulterius aestivus talis tero ullam cruciamentum."),
            DocumentIds: []string{},
            FileIds: []string{},
            ID: unifiedgosdk.Pointer("4272fc6f-d650-4593-88de-8d62deee8aef"),
            IsAssembly: unifiedgosdk.Pointer(false),
            LifecycleState: unifiedgosdk.Pointer("Design"),
            MasterIdentifier: unifiedgosdk.Pointer("099afe83-64ef-4ce4-8dce-c5c0d4273d54"),
            Metadata: []shared.PlmMetadata{},
            Name: unifiedgosdk.Pointer("grumpy PCB"),
            Number: unifiedgosdk.Pointer("PRT-009154"),
            OriginalType: unifiedgosdk.Pointer("Software"),
            Revision: unifiedgosdk.Pointer("B"),
            Source: shared.SourceMake.ToPointer(),
            Status: shared.PlmPartStatusInWork.ToPointer(),
            Type: shared.PlmPartTypeSoftware.ToPointer(),
            UnitOfMeasure: unifiedgosdk.Pointer("EA"),
            UpdatedAt: types.MustNewTimeFromString("2021-04-12T10:26:14.689Z"),
            URL: unifiedgosdk.Pointer("https://inborn-independence.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlmPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdatePlmPartRequest](../../pkg/models/operations/updateplmpartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdatePlmPartResponse](../../pkg/models/operations/updateplmpartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |