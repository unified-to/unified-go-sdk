# Page

## Overview

### Available Operations

* [CreateKmsPage](#createkmspage) - Create a page
* [GetKmsPage](#getkmspage) - Retrieve a page
* [ListKmsPages](#listkmspages) - List all pages
* [PatchKmsPage](#patchkmspage) - Update a page
* [RemoveKmsPage](#removekmspage) - Remove a page
* [UpdateKmsPage](#updatekmspage) - Update a page

## CreateKmsPage

Create a page

### Example Usage

<!-- UsageSnippet language="go" operationID="createKmsPage" method="post" path="/kms/{connection_id}/page" example="kms_page" -->
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

    res, err := s.Page.CreateKmsPage(ctx, operations.CreateKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("fb7ec6e4-c3a4-4255-8bf2-36c69840a206"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-09-11T13:40:48.688Z"),
            WebURL: unifiedgosdk.Pointer("https://another-petticoat.info"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateKmsPageRequest](../../pkg/models/operations/createkmspagerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateKmsPageResponse](../../pkg/models/operations/createkmspageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetKmsPage

Retrieve a page

### Example Usage

<!-- UsageSnippet language="go" operationID="getKmsPage" method="get" path="/kms/{connection_id}/page/{id}" -->
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

    res, err := s.Page.GetKmsPage(ctx, operations.GetKmsPageRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetKmsPageRequest](../../pkg/models/operations/getkmspagerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetKmsPageResponse](../../pkg/models/operations/getkmspageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListKmsPages

List all pages

### Example Usage

<!-- UsageSnippet language="go" operationID="listKmsPages" method="get" path="/kms/{connection_id}/page" -->
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

    res, err := s.Page.ListKmsPages(ctx, operations.ListKmsPagesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsPages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListKmsPagesRequest](../../pkg/models/operations/listkmspagesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListKmsPagesResponse](../../pkg/models/operations/listkmspagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchKmsPage

Update a page

### Example Usage

<!-- UsageSnippet language="go" operationID="patchKmsPage" method="patch" path="/kms/{connection_id}/page/{id}" example="kms_page" -->
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

    res, err := s.Page.PatchKmsPage(ctx, operations.PatchKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f8857f1f-711a-4d62-bee8-b016804a9b73"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-09-11T13:40:48.701Z"),
            WebURL: unifiedgosdk.Pointer("https://another-petticoat.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchKmsPageRequest](../../pkg/models/operations/patchkmspagerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchKmsPageResponse](../../pkg/models/operations/patchkmspageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveKmsPage

Remove a page

### Example Usage

<!-- UsageSnippet language="go" operationID="removeKmsPage" method="delete" path="/kms/{connection_id}/page/{id}" -->
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

    res, err := s.Page.RemoveKmsPage(ctx, operations.RemoveKmsPageRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveKmsPageRequest](../../pkg/models/operations/removekmspagerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveKmsPageResponse](../../pkg/models/operations/removekmspageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateKmsPage

Update a page

### Example Usage

<!-- UsageSnippet language="go" operationID="updateKmsPage" method="put" path="/kms/{connection_id}/page/{id}" example="kms_page" -->
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

    res, err := s.Page.UpdateKmsPage(ctx, operations.UpdateKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f8857f1f-711a-4d62-bee8-b016804a9b73"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-09-11T13:40:48.701Z"),
            WebURL: unifiedgosdk.Pointer("https://another-petticoat.info"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateKmsPageRequest](../../pkg/models/operations/updatekmspagerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateKmsPageResponse](../../pkg/models/operations/updatekmspageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |