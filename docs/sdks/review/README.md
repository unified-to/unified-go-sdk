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

Creates a CommerceReview. Provide the unified CommerceReview fields in the request body; they are mapped to the provider's API. Returns the created CommerceReview.

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
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("5511e647-c88b-49ec-a719-ff9102b72c8a"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("655ee5a2-586f-49e6-b928-af0c396e311b"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("f431de97-875f-45bb-b368-5a2124680506"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("8ab75126-17bf-45c6-b025-20689efa94a0"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("37f21345-4a56-42ee-bdba-e3331ac2b3a7"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T04:26:27.338Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
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

Retrieves a single CommerceReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single PerformanceReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single SocialReview by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of CommerceReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of PerformanceReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of SocialReviews. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the CommerceReview identified by `id`. Provide the unified CommerceReview fields to change in the request body.

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
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("68f33a18-c78b-4069-a9b5-44e7f13e5a51"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("34369488-7bac-478b-9dcd-1a33de100275"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("54ecfb50-7895-4f9d-87d5-db6f71068c76"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("1e8a1d76-307a-42fa-aa87-d954fe49d115"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("0cbc7a59-778a-4ad9-9e52-3788fb32cc25"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T04:26:27.383Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
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

Updates the SocialReview identified by `id`. Provide the unified SocialReview fields to change in the request body.

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

Deletes the CommerceReview identified by `id`.

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

Updates the CommerceReview identified by `id`. Provide the unified CommerceReview fields to change in the request body.

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
            AuthorAvatarURL: unifiedgosdk.Pointer("https://loremflickr.com/7/2660?lock=4279109054086693"),
            AuthorEmail: unifiedgosdk.Pointer("nyasia.macgyver@hotmail.com"),
            AuthorLocation: unifiedgosdk.Pointer("denique"),
            AuthorName: unifiedgosdk.Pointer("Nyasia MacGyver"),
            Comments: []shared.CommerceReviewComment{},
            Content: unifiedgosdk.Pointer("Taedium thymum adipiscor amicitia cui."),
            CreatedAt: types.MustNewTimeFromString("2019-12-12T18:10:22.988Z"),
            HelpfulVotes: unifiedgosdk.Pointer[float64](66.0),
            ID: unifiedgosdk.Pointer("68f33a18-c78b-4069-a9b5-44e7f13e5a51"),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsPublic: unifiedgosdk.Pointer(true),
            IsVerified: unifiedgosdk.Pointer(false),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Desipio carpo sol laudantium colo."),
                    Height: unifiedgosdk.Pointer[float64](749.0),
                    ID: unifiedgosdk.Pointer("34369488-7bac-478b-9dcd-1a33de100275"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("54ecfb50-7895-4f9d-87d5-db6f71068c76"),
                            Slug: unifiedgosdk.Pointer("textor"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quos",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](29.0),
                    Type: shared.CommerceItemMediaTypeVideo.ToPointer(),
                    URL: "https://loremflickr.com/196/2041?lock=5845524773114807",
                    Width: unifiedgosdk.Pointer[float64](72.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Adopto teneo numquam."),
                    Height: unifiedgosdk.Pointer[float64](957.0),
                    ID: unifiedgosdk.Pointer("1e8a1d76-307a-42fa-aa87-d954fe49d115"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("0cbc7a59-778a-4ad9-9e52-3788fb32cc25"),
                            Slug: unifiedgosdk.Pointer("bos"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "casso",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](69.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/1178/2175?lock=5466640101053206",
                    Width: unifiedgosdk.Pointer[float64](61.0),
                },
            },
            Metadata: []shared.CommerceMetadata{},
            Rating: unifiedgosdk.Pointer[float64](3.0),
            Status: shared.CommerceReviewStatusApproved.ToPointer(),
            Title: unifiedgosdk.Pointer("Coepi adamo amicitia auxilium toties."),
            UnhelpfulVotes: unifiedgosdk.Pointer[float64](88.0),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T04:26:27.383Z"),
            URL: unifiedgosdk.Pointer("https://diligent-bandwidth.biz"),
            VerifiedPurchase: unifiedgosdk.Pointer(true),
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

Updates the SocialReview identified by `id`. Provide the unified SocialReview fields to change in the request body.

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