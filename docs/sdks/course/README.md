# Course

## Overview

### Available Operations

* [CreateLmsCourse](#createlmscourse) - Create a course
* [GetLmsCourse](#getlmscourse) - Retrieve a course
* [ListLmsCourses](#listlmscourses) - List all courses
* [PatchLmsCourse](#patchlmscourse) - Update a course
* [RemoveLmsCourse](#removelmscourse) - Remove a course
* [UpdateLmsCourse](#updatelmscourse) - Update a course

## CreateLmsCourse

Creates a LmsCourse. Provide the unified LmsCourse fields in the request body; they are mapped to the provider's API. Returns the created LmsCourse.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsCourse" method="post" path="/lms/{connection_id}/course" example="lms_course" -->
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

    res, err := s.Course.CreateLmsCourse(ctx, operations.CreateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("ded9a63c-2d03-486c-8f8e-0c28e0c86b65"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "desparatus",
                "earum",
                "deripio",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adeptio crudelis ipsum utrimque quae architecto. Cum eius conitor anser abutor error adsuesco abeo. Denego nihil caries aveho."),
                    Description: unifiedgosdk.Pointer("Adipiscor."),
                    Languages: []string{
                        "ms",
                        "te",
                    },
                    Name: unifiedgosdk.Pointer("tandem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/syTatRhK03/928/273"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/fQAbsk/2472/1671"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2022-11-05T19:11:32.281Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T22:31:48.463Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCourse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLmsCourseRequest](../../pkg/models/operations/createlmscourserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateLmsCourseResponse](../../pkg/models/operations/createlmscourseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsCourse

Retrieves a single LmsCourse by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsCourse" method="get" path="/lms/{connection_id}/course/{id}" -->
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

    res, err := s.Course.GetLmsCourse(ctx, operations.GetLmsCourseRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCourse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLmsCourseRequest](../../pkg/models/operations/getlmscourserequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetLmsCourseResponse](../../pkg/models/operations/getlmscourseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsCourses

Returns a list of LmsCourses. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsCourses" method="get" path="/lms/{connection_id}/course" -->
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

    res, err := s.Course.ListLmsCourses(ctx, operations.ListLmsCoursesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCourses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLmsCoursesRequest](../../pkg/models/operations/listlmscoursesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListLmsCoursesResponse](../../pkg/models/operations/listlmscoursesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsCourse

Updates the LmsCourse identified by `id`. Provide the unified LmsCourse fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsCourse" method="patch" path="/lms/{connection_id}/course/{id}" example="lms_course" -->
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

    res, err := s.Course.PatchLmsCourse(ctx, operations.PatchLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("af97a2a9-d3c9-4797-896f-33179563d5c4"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "desparatus",
                "earum",
                "deripio",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adeptio crudelis ipsum utrimque quae architecto. Cum eius conitor anser abutor error adsuesco abeo. Denego nihil caries aveho."),
                    Description: unifiedgosdk.Pointer("Adipiscor."),
                    Languages: []string{
                        "ms",
                        "te",
                    },
                    Name: unifiedgosdk.Pointer("tandem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/syTatRhK03/928/273"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/fQAbsk/2472/1671"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2022-11-05T19:11:32.281Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T22:31:48.466Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCourse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchLmsCourseRequest](../../pkg/models/operations/patchlmscourserequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchLmsCourseResponse](../../pkg/models/operations/patchlmscourseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsCourse

Deletes the LmsCourse identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsCourse" method="delete" path="/lms/{connection_id}/course/{id}" -->
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

    res, err := s.Course.RemoveLmsCourse(ctx, operations.RemoveLmsCourseRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveLmsCourseRequest](../../pkg/models/operations/removelmscourserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveLmsCourseResponse](../../pkg/models/operations/removelmscourseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsCourse

Updates the LmsCourse identified by `id`. Provide the unified LmsCourse fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsCourse" method="put" path="/lms/{connection_id}/course/{id}" example="lms_course" -->
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

    res, err := s.Course.UpdateLmsCourse(ctx, operations.UpdateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("af97a2a9-d3c9-4797-896f-33179563d5c4"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "desparatus",
                "earum",
                "deripio",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adeptio crudelis ipsum utrimque quae architecto. Cum eius conitor anser abutor error adsuesco abeo. Denego nihil caries aveho."),
                    Description: unifiedgosdk.Pointer("Adipiscor."),
                    Languages: []string{
                        "ms",
                        "te",
                    },
                    Name: unifiedgosdk.Pointer("tandem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/syTatRhK03/928/273"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/fQAbsk/2472/1671"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2022-11-05T19:11:32.281Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T22:31:48.466Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCourse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateLmsCourseRequest](../../pkg/models/operations/updatelmscourserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateLmsCourseResponse](../../pkg/models/operations/updatelmscourseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |