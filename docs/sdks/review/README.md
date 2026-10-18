# Review

## Overview

### Available Operations

* [CreateCommerceReview](#createcommercereview) - Create a review
* [GetCommerceReview](#getcommercereview) - Retrieve a review
* [GetPerformanceReview](#getperformancereview) - Retrieve a review
* [GetSocialReview](#getsocialreview) - Retrieve a review
* [ListCommerceReviews](#listcommercereviews) - List all reviews
* [ListPerformanceReviews](#listperformancereviews) - List all reviews
* [ListSocialReviews](#listsocialreviews) - List all reviews
* [PatchCommerceReview](#patchcommercereview) - Update a review
* [PatchSocialReview](#patchsocialreview) - Update a review
* [RemoveCommerceReview](#removecommercereview) - Remove a review
* [UpdateCommerceReview](#updatecommercereview) - Update a review
* [UpdateSocialReview](#updatesocialreview) - Update a review

## CreateCommerceReview

Create a review

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceReview" method="post" path="/commerce/{connection_id}/review" example="commerce_review" -->
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

    res, err := s.Review.CreateCommerceReview(ctx, operations.CreateCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/ix4Br3LA/2245/1245"),
            AuthorEmail: unifiedgosdk.Pointer("Cleve_Yundt@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("ipsum"),
            AuthorName: unifiedgosdk.Pointer("Marsha Krajcik"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](26.0),
            ID: unifiedgosdk.Pointer("e4999363-74a7-4d55-9444-fe053499533a"),
            IsFeatured: unifiedgosdk.Pointer(true),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adulescens."),
                    Height: unifiedgosdk.Pointer[float64](519.0),
                    ID: unifiedgosdk.Pointer("47515e14-d0b7-4def-b8e6-325f06579fa9"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("072f9165-4ad1-459f-a5b0-e99fdd072af4"),
                            Slug: unifiedgosdk.Pointer("aggero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "tero",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/882/1004?lock=7448492654002422",
                    Width: unifiedgosdk.Pointer[float64](75.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Pauci timidus sol comburo thema."),
                    Height: unifiedgosdk.Pointer[float64](297.0),
                    ID: unifiedgosdk.Pointer("f208b000-48fe-4d06-bbd7-23019a3e6d0e"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("e611e244-0dcd-4050-b6ea-38b786e0fb5b"),
                            Slug: unifiedgosdk.Pointer("vito"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "cuppedia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](61.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/3QDZ8/1208/2171",
                    Width: unifiedgosdk.Pointer[float64](96.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Cuppedia vestrum patruus."),
                    Height: unifiedgosdk.Pointer[float64](6.0),
                    ID: unifiedgosdk.Pointer("3d953db2-853f-46a2-836e-76e4ed02fc31"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("ed7153b2-244d-428b-a71b-33a138f71de2"),
                            Slug: unifiedgosdk.Pointer("arbitro"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "villa",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/ytybC/2616/710",
                    Width: unifiedgosdk.Pointer[float64](74.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](49.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-18T21:40:29.740Z"),
            URL: unifiedgosdk.Pointer("https://excitable-underneath.com"),
            VerifiedPurchase: unifiedgosdk.Pointer(false),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateCommerceReviewRequest](../../pkg/models/operations/createcommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateCommerceReviewResponse](../../pkg/models/operations/createcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceReview

Retrieve a review

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceReview" method="get" path="/commerce/{connection_id}/review/{id}" -->
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

    res, err := s.Review.GetCommerceReview(ctx, operations.GetCommerceReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetCommerceReviewRequest](../../pkg/models/operations/getcommercereviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetCommerceReviewResponse](../../pkg/models/operations/getcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPerformanceReview

Retrieve a review

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceReview" method="get" path="/performance/{connection_id}/review/{id}" -->
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

    res, err := s.Review.GetPerformanceReview(ctx, operations.GetPerformanceReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPerformanceReviewRequest](../../pkg/models/operations/getperformancereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetPerformanceReviewResponse](../../pkg/models/operations/getperformancereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSocialReview

Retrieve a review

### Example Usage

<!-- UsageSnippet language="go" operationID="getSocialReview" method="get" path="/social/{connection_id}/review/{id}" -->
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

    res, err := s.Review.GetSocialReview(ctx, operations.GetSocialReviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SocialReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSocialReviewRequest](../../pkg/models/operations/getsocialreviewrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetSocialReviewResponse](../../pkg/models/operations/getsocialreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceReviews

List all reviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceReviews" method="get" path="/commerce/{connection_id}/review" -->
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

    res, err := s.Review.ListCommerceReviews(ctx, operations.ListCommerceReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListCommerceReviewsRequest](../../pkg/models/operations/listcommercereviewsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListCommerceReviewsResponse](../../pkg/models/operations/listcommercereviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPerformanceReviews

List all reviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listPerformanceReviews" method="get" path="/performance/{connection_id}/review" -->
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

    res, err := s.Review.ListPerformanceReviews(ctx, operations.ListPerformanceReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListPerformanceReviewsRequest](../../pkg/models/operations/listperformancereviewsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListPerformanceReviewsResponse](../../pkg/models/operations/listperformancereviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSocialReviews

List all reviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listSocialReviews" method="get" path="/social/{connection_id}/review" -->
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

    res, err := s.Review.ListSocialReviews(ctx, operations.ListSocialReviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SocialReviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListSocialReviewsRequest](../../pkg/models/operations/listsocialreviewsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListSocialReviewsResponse](../../pkg/models/operations/listsocialreviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceReview" method="patch" path="/commerce/{connection_id}/review/{id}" example="commerce_review" -->
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

    res, err := s.Review.PatchCommerceReview(ctx, operations.PatchCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/ix4Br3LA/2245/1245"),
            AuthorEmail: unifiedgosdk.Pointer("Cleve_Yundt@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("ipsum"),
            AuthorName: unifiedgosdk.Pointer("Marsha Krajcik"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](26.0),
            ID: unifiedgosdk.Pointer("7076c563-c8b3-4bb5-bc80-c0407cf62fcb"),
            IsFeatured: unifiedgosdk.Pointer(true),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adulescens."),
                    Height: unifiedgosdk.Pointer[float64](519.0),
                    ID: unifiedgosdk.Pointer("6893409b-6cc2-4251-9c57-bc799dc1c1c7"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a083cc1f-41a7-4e9b-ad9d-6631be990740"),
                            Slug: unifiedgosdk.Pointer("aggero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "tero",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/882/1004?lock=7448492654002422",
                    Width: unifiedgosdk.Pointer[float64](75.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Pauci timidus sol comburo thema."),
                    Height: unifiedgosdk.Pointer[float64](297.0),
                    ID: unifiedgosdk.Pointer("ba3c7b36-a906-4567-865e-7cdff0aa9a54"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("3ead0b9a-d365-4821-bb78-67f314bd275b"),
                            Slug: unifiedgosdk.Pointer("vito"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "cuppedia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](61.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/3QDZ8/1208/2171",
                    Width: unifiedgosdk.Pointer[float64](96.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Cuppedia vestrum patruus."),
                    Height: unifiedgosdk.Pointer[float64](6.0),
                    ID: unifiedgosdk.Pointer("958f61d8-b62e-492d-9d1e-47754df0b28d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("51b95390-8c86-4766-803f-1b3eacc6e0da"),
                            Slug: unifiedgosdk.Pointer("arbitro"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "villa",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/ytybC/2616/710",
                    Width: unifiedgosdk.Pointer[float64](74.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](49.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-18T21:40:29.763Z"),
            URL: unifiedgosdk.Pointer("https://excitable-underneath.com"),
            VerifiedPurchase: unifiedgosdk.Pointer(false),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchCommerceReviewRequest](../../pkg/models/operations/patchcommercereviewrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchCommerceReviewResponse](../../pkg/models/operations/patchcommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchSocialReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="patchSocialReview" method="patch" path="/social/{connection_id}/review/{id}" -->
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

    res, err := s.Review.PatchSocialReview(ctx, operations.PatchSocialReviewRequest{
        SocialReview: shared.SocialReview{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SocialReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchSocialReviewRequest](../../pkg/models/operations/patchsocialreviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchSocialReviewResponse](../../pkg/models/operations/patchsocialreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceReview

Remove a review

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceReview" method="delete" path="/commerce/{connection_id}/review/{id}" -->
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

    res, err := s.Review.RemoveCommerceReview(ctx, operations.RemoveCommerceReviewRequest{
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
| `request`                                                                                            | [operations.RemoveCommerceReviewRequest](../../pkg/models/operations/removecommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveCommerceReviewResponse](../../pkg/models/operations/removecommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceReview" method="put" path="/commerce/{connection_id}/review/{id}" example="commerce_review" -->
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

    res, err := s.Review.UpdateCommerceReview(ctx, operations.UpdateCommerceReviewRequest{
        CommerceReview: shared.CommerceReview{
            AuthorAvatarURL: unifiedgosdk.Pointer("https://picsum.photos/seed/ix4Br3LA/2245/1245"),
            AuthorEmail: unifiedgosdk.Pointer("Cleve_Yundt@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("ipsum"),
            AuthorName: unifiedgosdk.Pointer("Marsha Krajcik"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](26.0),
            ID: unifiedgosdk.Pointer("7076c563-c8b3-4bb5-bc80-c0407cf62fcb"),
            IsFeatured: unifiedgosdk.Pointer(true),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adulescens."),
                    Height: unifiedgosdk.Pointer[float64](519.0),
                    ID: unifiedgosdk.Pointer("6893409b-6cc2-4251-9c57-bc799dc1c1c7"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a083cc1f-41a7-4e9b-ad9d-6631be990740"),
                            Slug: unifiedgosdk.Pointer("aggero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "tero",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/882/1004?lock=7448492654002422",
                    Width: unifiedgosdk.Pointer[float64](75.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Pauci timidus sol comburo thema."),
                    Height: unifiedgosdk.Pointer[float64](297.0),
                    ID: unifiedgosdk.Pointer("ba3c7b36-a906-4567-865e-7cdff0aa9a54"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("3ead0b9a-d365-4821-bb78-67f314bd275b"),
                            Slug: unifiedgosdk.Pointer("vito"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "cuppedia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](61.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/3QDZ8/1208/2171",
                    Width: unifiedgosdk.Pointer[float64](96.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Cuppedia vestrum patruus."),
                    Height: unifiedgosdk.Pointer[float64](6.0),
                    ID: unifiedgosdk.Pointer("958f61d8-b62e-492d-9d1e-47754df0b28d"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("51b95390-8c86-4766-803f-1b3eacc6e0da"),
                            Slug: unifiedgosdk.Pointer("arbitro"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "villa",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://picsum.photos/seed/ytybC/2616/710",
                    Width: unifiedgosdk.Pointer[float64](74.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](49.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-18T21:40:29.763Z"),
            URL: unifiedgosdk.Pointer("https://excitable-underneath.com"),
            VerifiedPurchase: unifiedgosdk.Pointer(false),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateCommerceReviewRequest](../../pkg/models/operations/updatecommercereviewrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateCommerceReviewResponse](../../pkg/models/operations/updatecommercereviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateSocialReview

Update a review

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSocialReview" method="put" path="/social/{connection_id}/review/{id}" -->
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

    res, err := s.Review.UpdateSocialReview(ctx, operations.UpdateSocialReviewRequest{
        SocialReview: shared.SocialReview{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SocialReview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateSocialReviewRequest](../../pkg/models/operations/updatesocialreviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateSocialReviewResponse](../../pkg/models/operations/updatesocialreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |