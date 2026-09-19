# Taxonomy

## Overview

### Available Operations

* [CreateHrisTaxonomy](#createhristaxonomy) - Create a taxonomy
* [GetHrisTaxonomy](#gethristaxonomy) - Retrieve a taxonomy
* [ListCrmTaxonomies](#listcrmtaxonomies) - List all taxonomies
* [ListHrisTaxonomies](#listhristaxonomies) - List all taxonomies

## CreateHrisTaxonomy

Create a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTaxonomy" method="post" path="/hris/{connection_id}/taxonomy" example="hris_taxonomy" -->
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

    res, err := s.Taxonomy.CreateHrisTaxonomy(ctx, operations.CreateHrisTaxonomyRequest{
        HrisTaxonomy: shared.HrisTaxonomy{
            CreatedAt: types.MustNewTimeFromString("2022-06-23T02:10:00.789Z"),
            Description: unifiedgosdk.Pointer("Apto demonstro audacia adstringo cursim tristis solio careo."),
            Domain: unifiedgosdk.Pointer("Electronics"),
            ID: unifiedgosdk.Pointer("ede085db-5709-4d53-a490-746f3de5be17"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("International Functionality Architect"),
            ParentID: unifiedgosdk.Pointer("6524b2a7-6520-4e15-8c4e-1aa6793db837"),
            RoleIds: []string{
                "2b1ef757-eb4c-4207-8af1-929afe49cd65",
            },
            Subcategory: unifiedgosdk.Pointer("Bamboo"),
            Type: shared.HrisTaxonomyTypeKnowledge.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-22T09:46:16.697Z"),
            URL: unifiedgosdk.Pointer("https://our-polarisation.name"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisTaxonomyRequest](../../pkg/models/operations/createhristaxonomyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisTaxonomyResponse](../../pkg/models/operations/createhristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTaxonomy

Retrieve a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTaxonomy" method="get" path="/hris/{connection_id}/taxonomy/{id}" -->
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

    res, err := s.Taxonomy.GetHrisTaxonomy(ctx, operations.GetHrisTaxonomyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisTaxonomyRequest](../../pkg/models/operations/gethristaxonomyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisTaxonomyResponse](../../pkg/models/operations/gethristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmTaxonomies

List all taxonomies

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmTaxonomies" method="get" path="/crm/{connection_id}/taxonomy" -->
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

    res, err := s.Taxonomy.ListCrmTaxonomies(ctx, operations.ListCrmTaxonomiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmTaxonomies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCrmTaxonomiesRequest](../../pkg/models/operations/listcrmtaxonomiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCrmTaxonomiesResponse](../../pkg/models/operations/listcrmtaxonomiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTaxonomies

List all taxonomies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTaxonomies" method="get" path="/hris/{connection_id}/taxonomy" -->
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

    res, err := s.Taxonomy.ListHrisTaxonomies(ctx, operations.ListHrisTaxonomiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTaxonomiesRequest](../../pkg/models/operations/listhristaxonomiesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTaxonomiesResponse](../../pkg/models/operations/listhristaxonomiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |