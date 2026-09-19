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

Create a course

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
            ID: unifiedgosdk.Pointer("07366ae4-bfb5-4e49-8037-c19297201498"),
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
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2023-12-30T03:35:03.902Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T00:25:41.280Z"),
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

Retrieve a course

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

List all courses

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

Update a course

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
            ID: unifiedgosdk.Pointer("6ed64692-685d-4ec9-98f9-1c42ef617a8e"),
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
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2023-12-30T03:35:03.902Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T00:25:41.282Z"),
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

Remove a course

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

Update a course

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
            ID: unifiedgosdk.Pointer("6ed64692-685d-4ec9-98f9-1c42ef617a8e"),
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
            Name: unifiedgosdk.Pointer("comptus"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Homenick - Wunsch"),
            PublishedAt: types.MustNewTimeFromString("2023-12-30T03:35:03.902Z"),
            Skills: []string{
                "adiuvo",
                "tam",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T00:25:41.282Z"),
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