# Kms

## Overview

### Available Operations

* [CreateKmsComment](#createkmscomment) - Create a comment
* [CreateKmsPage](#createkmspage) - Create a page
* [CreateKmsSpace](#createkmsspace) - Create a space
* [GetKmsComment](#getkmscomment) - Retrieve a comment
* [GetKmsPage](#getkmspage) - Retrieve a page
* [GetKmsSpace](#getkmsspace) - Retrieve a space
* [ListKmsComments](#listkmscomments) - List all comments
* [ListKmsPages](#listkmspages) - List all pages
* [ListKmsSpaces](#listkmsspaces) - List all spaces
* [PatchKmsComment](#patchkmscomment) - Update a comment
* [PatchKmsPage](#patchkmspage) - Update a page
* [PatchKmsSpace](#patchkmsspace) - Update a space
* [RemoveKmsComment](#removekmscomment) - Remove a comment
* [RemoveKmsPage](#removekmspage) - Remove a page
* [RemoveKmsSpace](#removekmsspace) - Remove a space
* [UpdateKmsComment](#updatekmscomment) - Update a comment
* [UpdateKmsPage](#updatekmspage) - Update a page
* [UpdateKmsSpace](#updatekmsspace) - Update a space

## CreateKmsComment

Create a comment

### Example Usage

<!-- UsageSnippet language="go" operationID="createKmsComment" method="post" path="/kms/{connection_id}/comment" example="kms_comment" -->
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

    res, err := s.Kms.CreateKmsComment(ctx, operations.CreateKmsCommentRequest{
        KmsComment: shared.KmsComment{
            Content: unifiedgosdk.Pointer("Decimus tolero viriliter usque."),
            ContentType: shared.ContentTypeHTML.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2022-08-26T14:40:49.732Z"),
            ID: unifiedgosdk.Pointer("c11056cf-81eb-4884-986e-079ad8990b0e"),
            Type: shared.KmsCommentTypePage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-24T17:42:16.112Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateKmsCommentRequest](../../pkg/models/operations/createkmscommentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateKmsCommentResponse](../../pkg/models/operations/createkmscommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Kms.CreateKmsPage(ctx, operations.CreateKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("78e850ce-2693-442d-bbf1-5d74130445d0"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-10-04T18:20:33.040Z"),
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

## CreateKmsSpace

Create a space

### Example Usage

<!-- UsageSnippet language="go" operationID="createKmsSpace" method="post" path="/kms/{connection_id}/space" example="kms_space" -->
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

    res, err := s.Kms.CreateKmsSpace(ctx, operations.CreateKmsSpaceRequest{
        KmsSpace: shared.KmsSpace{
            CreatedAt: types.MustNewTimeFromString("2022-10-31T00:56:54.246Z"),
            Description: unifiedgosdk.Pointer("Acer."),
            ID: unifiedgosdk.Pointer("f929bab0-6f18-41a9-9a1e-055e2b2fc2c8"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("into nor afore"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-25T20:45:41.197Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateKmsSpaceRequest](../../pkg/models/operations/createkmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateKmsSpaceResponse](../../pkg/models/operations/createkmsspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetKmsComment

Retrieve a comment

### Example Usage

<!-- UsageSnippet language="go" operationID="getKmsComment" method="get" path="/kms/{connection_id}/comment/{id}" -->
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

    res, err := s.Kms.GetKmsComment(ctx, operations.GetKmsCommentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetKmsCommentRequest](../../pkg/models/operations/getkmscommentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetKmsCommentResponse](../../pkg/models/operations/getkmscommentresponse.md), error**

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

    res, err := s.Kms.GetKmsPage(ctx, operations.GetKmsPageRequest{
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

## GetKmsSpace

Retrieve a space

### Example Usage

<!-- UsageSnippet language="go" operationID="getKmsSpace" method="get" path="/kms/{connection_id}/space/{id}" -->
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

    res, err := s.Kms.GetKmsSpace(ctx, operations.GetKmsSpaceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetKmsSpaceRequest](../../pkg/models/operations/getkmsspacerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetKmsSpaceResponse](../../pkg/models/operations/getkmsspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListKmsComments

List all comments

### Example Usage

<!-- UsageSnippet language="go" operationID="listKmsComments" method="get" path="/kms/{connection_id}/comment" -->
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

    res, err := s.Kms.ListKmsComments(ctx, operations.ListKmsCommentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsComments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListKmsCommentsRequest](../../pkg/models/operations/listkmscommentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListKmsCommentsResponse](../../pkg/models/operations/listkmscommentsresponse.md), error**

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

    res, err := s.Kms.ListKmsPages(ctx, operations.ListKmsPagesRequest{
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

## ListKmsSpaces

List all spaces

### Example Usage

<!-- UsageSnippet language="go" operationID="listKmsSpaces" method="get" path="/kms/{connection_id}/space" -->
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

    res, err := s.Kms.ListKmsSpaces(ctx, operations.ListKmsSpacesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListKmsSpacesRequest](../../pkg/models/operations/listkmsspacesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListKmsSpacesResponse](../../pkg/models/operations/listkmsspacesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchKmsComment

Update a comment

### Example Usage

<!-- UsageSnippet language="go" operationID="patchKmsComment" method="patch" path="/kms/{connection_id}/comment/{id}" example="kms_comment" -->
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

    res, err := s.Kms.PatchKmsComment(ctx, operations.PatchKmsCommentRequest{
        KmsComment: shared.KmsComment{
            Content: unifiedgosdk.Pointer("Decimus tolero viriliter usque."),
            ContentType: shared.ContentTypeHTML.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2022-08-26T14:40:49.732Z"),
            ID: unifiedgosdk.Pointer("58e3a7b7-7cd4-4dc4-8df1-6b8fad804d7d"),
            Type: shared.KmsCommentTypePage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-24T17:42:16.115Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchKmsCommentRequest](../../pkg/models/operations/patchkmscommentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchKmsCommentResponse](../../pkg/models/operations/patchkmscommentresponse.md), error**

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

    res, err := s.Kms.PatchKmsPage(ctx, operations.PatchKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f75906db-f894-4709-b14c-7918c2f8b3e4"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-10-04T18:20:33.057Z"),
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

## PatchKmsSpace

Update a space

### Example Usage

<!-- UsageSnippet language="go" operationID="patchKmsSpace" method="patch" path="/kms/{connection_id}/space/{id}" example="kms_space" -->
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

    res, err := s.Kms.PatchKmsSpace(ctx, operations.PatchKmsSpaceRequest{
        KmsSpace: shared.KmsSpace{
            CreatedAt: types.MustNewTimeFromString("2022-10-31T00:56:54.246Z"),
            Description: unifiedgosdk.Pointer("Acer."),
            ID: unifiedgosdk.Pointer("cefd70b5-89f6-454d-94fc-0b15f61d73a1"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("into nor afore"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-25T20:45:41.205Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchKmsSpaceRequest](../../pkg/models/operations/patchkmsspacerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchKmsSpaceResponse](../../pkg/models/operations/patchkmsspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveKmsComment

Remove a comment

### Example Usage

<!-- UsageSnippet language="go" operationID="removeKmsComment" method="delete" path="/kms/{connection_id}/comment/{id}" -->
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

    res, err := s.Kms.RemoveKmsComment(ctx, operations.RemoveKmsCommentRequest{
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
| `request`                                                                                    | [operations.RemoveKmsCommentRequest](../../pkg/models/operations/removekmscommentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveKmsCommentResponse](../../pkg/models/operations/removekmscommentresponse.md), error**

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

    res, err := s.Kms.RemoveKmsPage(ctx, operations.RemoveKmsPageRequest{
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

## RemoveKmsSpace

Remove a space

### Example Usage

<!-- UsageSnippet language="go" operationID="removeKmsSpace" method="delete" path="/kms/{connection_id}/space/{id}" -->
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

    res, err := s.Kms.RemoveKmsSpace(ctx, operations.RemoveKmsSpaceRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveKmsSpaceRequest](../../pkg/models/operations/removekmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveKmsSpaceResponse](../../pkg/models/operations/removekmsspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateKmsComment

Update a comment

### Example Usage

<!-- UsageSnippet language="go" operationID="updateKmsComment" method="put" path="/kms/{connection_id}/comment/{id}" example="kms_comment" -->
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

    res, err := s.Kms.UpdateKmsComment(ctx, operations.UpdateKmsCommentRequest{
        KmsComment: shared.KmsComment{
            Content: unifiedgosdk.Pointer("Decimus tolero viriliter usque."),
            ContentType: shared.ContentTypeHTML.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2022-08-26T14:40:49.732Z"),
            ID: unifiedgosdk.Pointer("58e3a7b7-7cd4-4dc4-8df1-6b8fad804d7d"),
            Type: shared.KmsCommentTypePage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-24T17:42:16.115Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateKmsCommentRequest](../../pkg/models/operations/updatekmscommentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateKmsCommentResponse](../../pkg/models/operations/updatekmscommentresponse.md), error**

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

    res, err := s.Kms.UpdateKmsPage(ctx, operations.UpdateKmsPageRequest{
        KmsPage: shared.KmsPage{
            CreatedAt: types.MustNewTimeFromString("2019-05-20T18:06:50.749Z"),
            DownloadURL: unifiedgosdk.Pointer("https://agitated-validity.info"),
            HasChildren: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("f75906db-f894-4709-b14c-7918c2f8b3e4"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.KmsPageMetadata{},
            Title: unifiedgosdk.Pointer("even minister extract"),
            Type: shared.KmsPageTypeHTML,
            UpdatedAt: types.MustNewTimeFromString("2025-10-04T18:20:33.057Z"),
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

## UpdateKmsSpace

Update a space

### Example Usage

<!-- UsageSnippet language="go" operationID="updateKmsSpace" method="put" path="/kms/{connection_id}/space/{id}" example="kms_space" -->
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

    res, err := s.Kms.UpdateKmsSpace(ctx, operations.UpdateKmsSpaceRequest{
        KmsSpace: shared.KmsSpace{
            CreatedAt: types.MustNewTimeFromString("2022-10-31T00:56:54.246Z"),
            Description: unifiedgosdk.Pointer("Acer."),
            ID: unifiedgosdk.Pointer("cefd70b5-89f6-454d-94fc-0b15f61d73a1"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("into nor afore"),
            UpdatedAt: types.MustNewTimeFromString("2025-12-25T20:45:41.205Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateKmsSpaceRequest](../../pkg/models/operations/updatekmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateKmsSpaceResponse](../../pkg/models/operations/updatekmsspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |