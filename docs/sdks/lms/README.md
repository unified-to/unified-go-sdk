# Lms

## Overview

### Available Operations

* [CreateLmsActivity](#createlmsactivity) - Create an activity
* [CreateLmsClass](#createlmsclass) - Create a class
* [CreateLmsCollection](#createlmscollection) - Create a collection
* [CreateLmsContent](#createlmscontent) - Create a content
* [CreateLmsCourse](#createlmscourse) - Create a course
* [CreateLmsInstructor](#createlmsinstructor) - Create an instructor
* [CreateLmsStudent](#createlmsstudent) - Create a student
* [GetLmsActivity](#getlmsactivity) - Retrieve an activity
* [GetLmsClass](#getlmsclass) - Retrieve a class
* [GetLmsCollection](#getlmscollection) - Retrieve a collection
* [GetLmsContent](#getlmscontent) - Retrieve a content
* [GetLmsCourse](#getlmscourse) - Retrieve a course
* [GetLmsInstructor](#getlmsinstructor) - Retrieve an instructor
* [GetLmsStudent](#getlmsstudent) - Retrieve a student
* [ListLmsActivities](#listlmsactivities) - List all activities
* [ListLmsClasses](#listlmsclasses) - List all classes
* [ListLmsCollections](#listlmscollections) - List all collections
* [ListLmsContents](#listlmscontents) - List all contents
* [ListLmsCourses](#listlmscourses) - List all courses
* [ListLmsInstructors](#listlmsinstructors) - List all instructors
* [ListLmsStudents](#listlmsstudents) - List all students
* [PatchLmsActivity](#patchlmsactivity) - Update an activity
* [PatchLmsClass](#patchlmsclass) - Update a class
* [PatchLmsCollection](#patchlmscollection) - Update a collection
* [PatchLmsContent](#patchlmscontent) - Update a content
* [PatchLmsCourse](#patchlmscourse) - Update a course
* [PatchLmsInstructor](#patchlmsinstructor) - Update an instructor
* [PatchLmsStudent](#patchlmsstudent) - Update a student
* [RemoveLmsActivity](#removelmsactivity) - Remove an activity
* [RemoveLmsClass](#removelmsclass) - Remove a class
* [RemoveLmsCollection](#removelmscollection) - Remove a collection
* [RemoveLmsContent](#removelmscontent) - Remove a content
* [RemoveLmsCourse](#removelmscourse) - Remove a course
* [RemoveLmsInstructor](#removelmsinstructor) - Remove an instructor
* [RemoveLmsStudent](#removelmsstudent) - Remove a student
* [UpdateLmsActivity](#updatelmsactivity) - Update an activity
* [UpdateLmsClass](#updatelmsclass) - Update a class
* [UpdateLmsCollection](#updatelmscollection) - Update a collection
* [UpdateLmsContent](#updatelmscontent) - Update a content
* [UpdateLmsCourse](#updatelmscourse) - Update a course
* [UpdateLmsInstructor](#updatelmsinstructor) - Update an instructor
* [UpdateLmsStudent](#updatelmsstudent) - Update a student

## CreateLmsActivity

Creates a LmsActivity. Provide the unified LmsActivity fields in the request body; they are mapped to the provider's API. Returns the created LmsActivity.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsActivity" method="post" path="/lms/{connection_id}/activity" example="lms_activity" -->
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

    res, err := s.Lms.CreateLmsActivity(ctx, operations.CreateLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("deserunt"),
            CompletedAt: types.MustNewTimeFromString("2023-12-25T19:30:13.916Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("c2def3e7-4d8b-4a32-b9b8-a87778bb1a65"),
            IsCompleted: unifiedgosdk.Pointer(true),
            Metadata: []shared.LmsMetadata{},
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2021-04-28T11:24:24.796Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-08T03:24:15.850Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLmsActivityRequest](../../pkg/models/operations/createlmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateLmsActivityResponse](../../pkg/models/operations/createlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsClass

Creates a LmsClass. Provide the unified LmsClass fields in the request body; they are mapped to the provider's API. Returns the created LmsClass.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsClass" method="post" path="/lms/{connection_id}/class" example="lms_class" -->
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

    res, err := s.Lms.CreateLmsClass(ctx, operations.CreateLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Animus callide corpus."),
            ID: unifiedgosdk.Pointer("b323a6e3-baf4-4c87-b679-b303bf590bbd"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ceno earum culpa adhaero tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana. Solvo uter adficio."),
                    Description: unifiedgosdk.Pointer("Defetiscor aetas."),
                    Languages: []string{
                        "ko",
                        "lv",
                    },
                    Name: unifiedgosdk.Pointer("suspendo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/coddeQdTZn/105/852"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/wl3tnt/2277/2170"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("sperno"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-09-03T23:40:45.249Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLmsClassRequest](../../pkg/models/operations/createlmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateLmsClassResponse](../../pkg/models/operations/createlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsCollection

Creates a LmsCollection. Provide the unified LmsCollection fields in the request body; they are mapped to the provider's API. Returns the created LmsCollection.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsCollection" method="post" path="/lms/{connection_id}/collection" example="lms_collection" -->
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

    res, err := s.Lms.CreateLmsCollection(ctx, operations.CreateLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("439785fe-6475-4a32-9d13-d51268ed7f41"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.029Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateLmsCollectionRequest](../../pkg/models/operations/createlmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateLmsCollectionResponse](../../pkg/models/operations/createlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsContent

Creates a LmsContent. Provide the unified LmsContent fields in the request body; they are mapped to the provider's API. Returns the created LmsContent.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsContent" method="post" path="/lms/{connection_id}/content" example="lms_content" -->
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

    res, err := s.Lms.CreateLmsContent(ctx, operations.CreateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("c04dc5d6-355d-45bf-b1e1-3965e5ba2747"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "campana",
                "viriliter",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Vindico vorago."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("aeger"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Sulum."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("voluptate"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Cohors necessitatibus beatae copiose copiose cedo sub decens acer. Spiritus velum adulatio commodi textus beneficium termes comedo valde caste. Sodalitas bos timor abeo cur. Sperno decet acquiro tabgo canto terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum."),
                    Description: unifiedgosdk.Pointer("Incidunt tepidus commemoro voro ocer."),
                    Languages: []string{
                        "ro",
                        "af",
                    },
                    Name: unifiedgosdk.Pointer("antepono"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/URbmlO66Wk/587/3603"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/4KFvVbCx83/2788/2850"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ago neque qui cursim tunc barba decens. Tutamen amplus tabella tabgo infit. Demo adversus capillus aro unus. Dedico aspicio alius vehemens terminatio varietas vomica voluptatibus apud. Defero decet impedit verbera pauci laboriosam molestiae urbanus mollitia. Vergo crux certus nulla."),
                    Description: unifiedgosdk.Pointer("Bene doloribus."),
                    Languages: []string{
                        "ru",
                        "sw",
                    },
                    Name: unifiedgosdk.Pointer("totidem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/qpRev/1286/3718"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/JOWrp5/254/2004"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Laborum adiuvo labore officia demulceo. Abduco mollitia ventosus synagoga certus magni vado cilicium stella. Velut capitulus attero odio advenio similique acsi totam atavus aufero. Ventosus campana animi."),
                    Description: unifiedgosdk.Pointer("Trucido dens pecus aer stella."),
                    Languages: []string{
                        "is",
                        "sk",
                    },
                    Name: unifiedgosdk.Pointer("territo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/81EuDedaf/557/986"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/BtV2eDuVm7/3644/1033"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("laboriosam"),
            ProviderName: unifiedgosdk.Pointer("Mraz - Keeling"),
            PublishedAt: types.MustNewTimeFromString("2022-09-15T01:48:30.262Z"),
            ShortDescription: unifiedgosdk.Pointer("Videlicet pauci."),
            Skills: []string{
                "soleo",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("clam"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "audio",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T12:58:06.317Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLmsContentRequest](../../pkg/models/operations/createlmscontentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateLmsContentResponse](../../pkg/models/operations/createlmscontentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Lms.CreateLmsCourse(ctx, operations.CreateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "deripio",
                "tergiversatio",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("DZD"),
            Description: unifiedgosdk.Pointer("Aut."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("75268521-c6d2-4c34-8e8b-6b5635c65c0b"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "capitulus",
                "desparatus",
                "earum",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Carbo dapifer adeptio. Ipsum utrimque quae architecto repellendus cum. Conitor anser abutor error adsuesco abeo caelestis."),
                    Description: unifiedgosdk.Pointer("Accommodo adipiscor campana aegrotatio."),
                    Languages: []string{
                        "el",
                        "hr",
                    },
                    Name: unifiedgosdk.Pointer("conduco"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/yTatRhK/3873/3840"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/vcfQAbs/3357/2472"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("vinco"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Roberts Group"),
            PublishedAt: types.MustNewTimeFromString("2022-11-06T12:56:32.674Z"),
            Skills: []string{
                "aveho",
                "via",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-13T00:26:39.780Z"),
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

## CreateLmsInstructor

Creates a LmsInstructor. Provide the unified LmsInstructor fields in the request body; they are mapped to the provider's API. Returns the created LmsInstructor.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsInstructor" method="post" path="/lms/{connection_id}/instructor" example="lms_instructor" -->
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

    res, err := s.Lms.CreateLmsInstructor(ctx, operations.CreateLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Miranda"),
            ID: unifiedgosdk.Pointer("6fa6db00-20d6-49da-867a-ae305b627fb9"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/96902788"),
            LastName: unifiedgosdk.Pointer("Turcotte"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Miranda Turcotte"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(752) 816-3330",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(415) 345-1798",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Brand Planner"),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T00:58:33.332Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateLmsInstructorRequest](../../pkg/models/operations/createlmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateLmsInstructorResponse](../../pkg/models/operations/createlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsStudent

Creates a LmsStudent. Provide the unified LmsStudent fields in the request body; they are mapped to the provider's API. Returns the created LmsStudent.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsStudent" method="post" path="/lms/{connection_id}/student" example="lms_student" -->
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

    res, err := s.Lms.CreateLmsStudent(ctx, operations.CreateLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("4610 Smith Cove"),
                City: unifiedgosdk.Pointer("Lake Milanchester"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("58634-4787"),
                Region: unifiedgosdk.Pointer("Iowa"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Lazaro"),
            ID: unifiedgosdk.Pointer("f3da5ff0-b8df-48f5-b92c-e66cd074e9b8"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/3245424"),
            LastName: unifiedgosdk.Pointer("Zulauf-Schowalter"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Lazaro Zulauf-Schowalter"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(476) 500-2870",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(338) 436-9806",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(852) 940-8205",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-14T02:26:33.985Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLmsStudentRequest](../../pkg/models/operations/createlmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateLmsStudentResponse](../../pkg/models/operations/createlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsActivity

Retrieves a single LmsActivity by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsActivity" method="get" path="/lms/{connection_id}/activity/{id}" -->
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

    res, err := s.Lms.GetLmsActivity(ctx, operations.GetLmsActivityRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLmsActivityRequest](../../pkg/models/operations/getlmsactivityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetLmsActivityResponse](../../pkg/models/operations/getlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsClass

Retrieves a single LmsClass by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsClass" method="get" path="/lms/{connection_id}/class/{id}" -->
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

    res, err := s.Lms.GetLmsClass(ctx, operations.GetLmsClassRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLmsClassRequest](../../pkg/models/operations/getlmsclassrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetLmsClassResponse](../../pkg/models/operations/getlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsCollection

Retrieves a single LmsCollection by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsCollection" method="get" path="/lms/{connection_id}/collection/{id}" -->
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

    res, err := s.Lms.GetLmsCollection(ctx, operations.GetLmsCollectionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetLmsCollectionRequest](../../pkg/models/operations/getlmscollectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetLmsCollectionResponse](../../pkg/models/operations/getlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsContent

Retrieves a single LmsContent by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsContent" method="get" path="/lms/{connection_id}/content/{id}" -->
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

    res, err := s.Lms.GetLmsContent(ctx, operations.GetLmsContentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLmsContentRequest](../../pkg/models/operations/getlmscontentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetLmsContentResponse](../../pkg/models/operations/getlmscontentresponse.md), error**

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

    res, err := s.Lms.GetLmsCourse(ctx, operations.GetLmsCourseRequest{
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

## GetLmsInstructor

Retrieves a single LmsInstructor by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsInstructor" method="get" path="/lms/{connection_id}/instructor/{id}" -->
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

    res, err := s.Lms.GetLmsInstructor(ctx, operations.GetLmsInstructorRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetLmsInstructorRequest](../../pkg/models/operations/getlmsinstructorrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetLmsInstructorResponse](../../pkg/models/operations/getlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsStudent

Retrieves a single LmsStudent by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsStudent" method="get" path="/lms/{connection_id}/student/{id}" -->
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

    res, err := s.Lms.GetLmsStudent(ctx, operations.GetLmsStudentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLmsStudentRequest](../../pkg/models/operations/getlmsstudentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetLmsStudentResponse](../../pkg/models/operations/getlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsActivities

Returns a list of LmsActivities. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsActivities" method="get" path="/lms/{connection_id}/activity" -->
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

    res, err := s.Lms.ListLmsActivities(ctx, operations.ListLmsActivitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLmsActivitiesRequest](../../pkg/models/operations/listlmsactivitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListLmsActivitiesResponse](../../pkg/models/operations/listlmsactivitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsClasses

Returns a list of LmsClasses. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsClasses" method="get" path="/lms/{connection_id}/class" -->
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

    res, err := s.Lms.ListLmsClasses(ctx, operations.ListLmsClassesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClasses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLmsClassesRequest](../../pkg/models/operations/listlmsclassesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListLmsClassesResponse](../../pkg/models/operations/listlmsclassesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsCollections

Returns a list of LmsCollections. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsCollections" method="get" path="/lms/{connection_id}/collection" -->
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

    res, err := s.Lms.ListLmsCollections(ctx, operations.ListLmsCollectionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLmsCollectionsRequest](../../pkg/models/operations/listlmscollectionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListLmsCollectionsResponse](../../pkg/models/operations/listlmscollectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsContents

Returns a list of LmsContents. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsContents" method="get" path="/lms/{connection_id}/content" -->
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

    res, err := s.Lms.ListLmsContents(ctx, operations.ListLmsContentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLmsContentsRequest](../../pkg/models/operations/listlmscontentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListLmsContentsResponse](../../pkg/models/operations/listlmscontentsresponse.md), error**

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

    res, err := s.Lms.ListLmsCourses(ctx, operations.ListLmsCoursesRequest{
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

## ListLmsInstructors

Returns a list of LmsInstructors. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsInstructors" method="get" path="/lms/{connection_id}/instructor" -->
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

    res, err := s.Lms.ListLmsInstructors(ctx, operations.ListLmsInstructorsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLmsInstructorsRequest](../../pkg/models/operations/listlmsinstructorsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListLmsInstructorsResponse](../../pkg/models/operations/listlmsinstructorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsStudents

Returns a list of LmsStudents. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsStudents" method="get" path="/lms/{connection_id}/student" -->
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

    res, err := s.Lms.ListLmsStudents(ctx, operations.ListLmsStudentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLmsStudentsRequest](../../pkg/models/operations/listlmsstudentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListLmsStudentsResponse](../../pkg/models/operations/listlmsstudentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsActivity

Updates the LmsActivity identified by `id`. Provide the unified LmsActivity fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsActivity" method="patch" path="/lms/{connection_id}/activity/{id}" example="lms_activity" -->
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

    res, err := s.Lms.PatchLmsActivity(ctx, operations.PatchLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("deserunt"),
            CompletedAt: types.MustNewTimeFromString("2023-12-25T19:30:13.923Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("0078f713-45cc-4197-8b1f-a274178faf2b"),
            IsCompleted: unifiedgosdk.Pointer(true),
            Metadata: []shared.LmsMetadata{},
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2021-04-28T11:24:24.797Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-08T03:24:15.853Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchLmsActivityRequest](../../pkg/models/operations/patchlmsactivityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchLmsActivityResponse](../../pkg/models/operations/patchlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsClass

Updates the LmsClass identified by `id`. Provide the unified LmsClass fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsClass" method="patch" path="/lms/{connection_id}/class/{id}" example="lms_class" -->
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

    res, err := s.Lms.PatchLmsClass(ctx, operations.PatchLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Animus callide corpus."),
            ID: unifiedgosdk.Pointer("f852b6ce-2791-460e-9ac5-dfc9f72be28c"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ceno earum culpa adhaero tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana. Solvo uter adficio."),
                    Description: unifiedgosdk.Pointer("Defetiscor aetas."),
                    Languages: []string{
                        "ko",
                        "lv",
                    },
                    Name: unifiedgosdk.Pointer("suspendo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/coddeQdTZn/105/852"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/wl3tnt/2277/2170"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("sperno"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-09-03T23:40:45.258Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchLmsClassRequest](../../pkg/models/operations/patchlmsclassrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchLmsClassResponse](../../pkg/models/operations/patchlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsCollection

Updates the LmsCollection identified by `id`. Provide the unified LmsCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsCollection" method="patch" path="/lms/{connection_id}/collection/{id}" example="lms_collection" -->
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

    res, err := s.Lms.PatchLmsCollection(ctx, operations.PatchLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("b3ba5ccf-3f87-4ae8-86a4-6adb7d2b652a"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.038Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchLmsCollectionRequest](../../pkg/models/operations/patchlmscollectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchLmsCollectionResponse](../../pkg/models/operations/patchlmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsContent

Updates the LmsContent identified by `id`. Provide the unified LmsContent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsContent" method="patch" path="/lms/{connection_id}/content/{id}" example="lms_content" -->
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

    res, err := s.Lms.PatchLmsContent(ctx, operations.PatchLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("3f7fe3ef-01ea-41f6-9e1a-99302ff334e5"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "campana",
                "viriliter",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Vindico vorago."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("aeger"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Sulum."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("voluptate"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Cohors necessitatibus beatae copiose copiose cedo sub decens acer. Spiritus velum adulatio commodi textus beneficium termes comedo valde caste. Sodalitas bos timor abeo cur. Sperno decet acquiro tabgo canto terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum."),
                    Description: unifiedgosdk.Pointer("Incidunt tepidus commemoro voro ocer."),
                    Languages: []string{
                        "ro",
                        "af",
                    },
                    Name: unifiedgosdk.Pointer("antepono"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/URbmlO66Wk/587/3603"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/4KFvVbCx83/2788/2850"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ago neque qui cursim tunc barba decens. Tutamen amplus tabella tabgo infit. Demo adversus capillus aro unus. Dedico aspicio alius vehemens terminatio varietas vomica voluptatibus apud. Defero decet impedit verbera pauci laboriosam molestiae urbanus mollitia. Vergo crux certus nulla."),
                    Description: unifiedgosdk.Pointer("Bene doloribus."),
                    Languages: []string{
                        "ru",
                        "sw",
                    },
                    Name: unifiedgosdk.Pointer("totidem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/qpRev/1286/3718"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/JOWrp5/254/2004"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Laborum adiuvo labore officia demulceo. Abduco mollitia ventosus synagoga certus magni vado cilicium stella. Velut capitulus attero odio advenio similique acsi totam atavus aufero. Ventosus campana animi."),
                    Description: unifiedgosdk.Pointer("Trucido dens pecus aer stella."),
                    Languages: []string{
                        "is",
                        "sk",
                    },
                    Name: unifiedgosdk.Pointer("territo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/81EuDedaf/557/986"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/BtV2eDuVm7/3644/1033"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("laboriosam"),
            ProviderName: unifiedgosdk.Pointer("Mraz - Keeling"),
            PublishedAt: types.MustNewTimeFromString("2022-09-15T01:48:30.270Z"),
            ShortDescription: unifiedgosdk.Pointer("Videlicet pauci."),
            Skills: []string{
                "soleo",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("clam"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "audio",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T12:58:06.325Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchLmsContentRequest](../../pkg/models/operations/patchlmscontentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchLmsContentResponse](../../pkg/models/operations/patchlmscontentresponse.md), error**

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

    res, err := s.Lms.PatchLmsCourse(ctx, operations.PatchLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "deripio",
                "tergiversatio",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("DZD"),
            Description: unifiedgosdk.Pointer("Aut."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("5b382565-417b-4497-8643-fed1c43b86b7"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "capitulus",
                "desparatus",
                "earum",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Carbo dapifer adeptio. Ipsum utrimque quae architecto repellendus cum. Conitor anser abutor error adsuesco abeo caelestis."),
                    Description: unifiedgosdk.Pointer("Accommodo adipiscor campana aegrotatio."),
                    Languages: []string{
                        "el",
                        "hr",
                    },
                    Name: unifiedgosdk.Pointer("conduco"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/yTatRhK/3873/3840"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/vcfQAbs/3357/2472"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("vinco"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Roberts Group"),
            PublishedAt: types.MustNewTimeFromString("2022-11-06T12:56:32.674Z"),
            Skills: []string{
                "aveho",
                "via",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-13T00:26:39.781Z"),
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

## PatchLmsInstructor

Updates the LmsInstructor identified by `id`. Provide the unified LmsInstructor fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsInstructor" method="patch" path="/lms/{connection_id}/instructor/{id}" example="lms_instructor" -->
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

    res, err := s.Lms.PatchLmsInstructor(ctx, operations.PatchLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Miranda"),
            ID: unifiedgosdk.Pointer("0267df64-7a37-426a-adf5-74b4fd83d28f"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/96902788"),
            LastName: unifiedgosdk.Pointer("Turcotte"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Miranda Turcotte"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(752) 816-3330",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(415) 345-1798",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Brand Planner"),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T00:58:33.339Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchLmsInstructorRequest](../../pkg/models/operations/patchlmsinstructorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchLmsInstructorResponse](../../pkg/models/operations/patchlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsStudent

Updates the LmsStudent identified by `id`. Provide the unified LmsStudent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsStudent" method="patch" path="/lms/{connection_id}/student/{id}" example="lms_student" -->
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

    res, err := s.Lms.PatchLmsStudent(ctx, operations.PatchLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("4610 Smith Cove"),
                City: unifiedgosdk.Pointer("Lake Milanchester"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("58634-4787"),
                Region: unifiedgosdk.Pointer("Iowa"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Lazaro"),
            ID: unifiedgosdk.Pointer("5f3e807b-a3cd-41fb-a9f5-78a4f2905d63"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/3245424"),
            LastName: unifiedgosdk.Pointer("Zulauf-Schowalter"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Lazaro Zulauf-Schowalter"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(476) 500-2870",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(338) 436-9806",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(852) 940-8205",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-14T02:26:33.989Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchLmsStudentRequest](../../pkg/models/operations/patchlmsstudentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchLmsStudentResponse](../../pkg/models/operations/patchlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsActivity

Deletes the LmsActivity identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsActivity" method="delete" path="/lms/{connection_id}/activity/{id}" -->
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

    res, err := s.Lms.RemoveLmsActivity(ctx, operations.RemoveLmsActivityRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveLmsActivityRequest](../../pkg/models/operations/removelmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveLmsActivityResponse](../../pkg/models/operations/removelmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsClass

Deletes the LmsClass identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsClass" method="delete" path="/lms/{connection_id}/class/{id}" -->
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

    res, err := s.Lms.RemoveLmsClass(ctx, operations.RemoveLmsClassRequest{
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
| `request`                                                                                | [operations.RemoveLmsClassRequest](../../pkg/models/operations/removelmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveLmsClassResponse](../../pkg/models/operations/removelmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsCollection

Deletes the LmsCollection identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsCollection" method="delete" path="/lms/{connection_id}/collection/{id}" -->
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

    res, err := s.Lms.RemoveLmsCollection(ctx, operations.RemoveLmsCollectionRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveLmsCollectionRequest](../../pkg/models/operations/removelmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveLmsCollectionResponse](../../pkg/models/operations/removelmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsContent

Deletes the LmsContent identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsContent" method="delete" path="/lms/{connection_id}/content/{id}" -->
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

    res, err := s.Lms.RemoveLmsContent(ctx, operations.RemoveLmsContentRequest{
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
| `request`                                                                                    | [operations.RemoveLmsContentRequest](../../pkg/models/operations/removelmscontentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveLmsContentResponse](../../pkg/models/operations/removelmscontentresponse.md), error**

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

    res, err := s.Lms.RemoveLmsCourse(ctx, operations.RemoveLmsCourseRequest{
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

## RemoveLmsInstructor

Deletes the LmsInstructor identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsInstructor" method="delete" path="/lms/{connection_id}/instructor/{id}" -->
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

    res, err := s.Lms.RemoveLmsInstructor(ctx, operations.RemoveLmsInstructorRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveLmsInstructorRequest](../../pkg/models/operations/removelmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveLmsInstructorResponse](../../pkg/models/operations/removelmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsStudent

Deletes the LmsStudent identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsStudent" method="delete" path="/lms/{connection_id}/student/{id}" -->
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

    res, err := s.Lms.RemoveLmsStudent(ctx, operations.RemoveLmsStudentRequest{
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
| `request`                                                                                    | [operations.RemoveLmsStudentRequest](../../pkg/models/operations/removelmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveLmsStudentResponse](../../pkg/models/operations/removelmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsActivity

Updates the LmsActivity identified by `id`. Provide the unified LmsActivity fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsActivity" method="put" path="/lms/{connection_id}/activity/{id}" example="lms_activity" -->
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

    res, err := s.Lms.UpdateLmsActivity(ctx, operations.UpdateLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("deserunt"),
            CompletedAt: types.MustNewTimeFromString("2023-12-25T19:30:13.923Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("0078f713-45cc-4197-8b1f-a274178faf2b"),
            IsCompleted: unifiedgosdk.Pointer(true),
            Metadata: []shared.LmsMetadata{},
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2021-04-28T11:24:24.797Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-08T03:24:15.853Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLmsActivityRequest](../../pkg/models/operations/updatelmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateLmsActivityResponse](../../pkg/models/operations/updatelmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsClass

Updates the LmsClass identified by `id`. Provide the unified LmsClass fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsClass" method="put" path="/lms/{connection_id}/class/{id}" example="lms_class" -->
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

    res, err := s.Lms.UpdateLmsClass(ctx, operations.UpdateLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Animus callide corpus."),
            ID: unifiedgosdk.Pointer("f852b6ce-2791-460e-9ac5-dfc9f72be28c"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ceno earum culpa adhaero tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana. Solvo uter adficio."),
                    Description: unifiedgosdk.Pointer("Defetiscor aetas."),
                    Languages: []string{
                        "ko",
                        "lv",
                    },
                    Name: unifiedgosdk.Pointer("suspendo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/coddeQdTZn/105/852"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/wl3tnt/2277/2170"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("sperno"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-09-03T23:40:45.258Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLmsClassRequest](../../pkg/models/operations/updatelmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateLmsClassResponse](../../pkg/models/operations/updatelmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsCollection

Updates the LmsCollection identified by `id`. Provide the unified LmsCollection fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsCollection" method="put" path="/lms/{connection_id}/collection/{id}" example="lms_collection" -->
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

    res, err := s.Lms.UpdateLmsCollection(ctx, operations.UpdateLmsCollectionRequest{
        LmsCollection: shared.LmsCollection{
            CreatedAt: types.MustNewTimeFromString("2019-08-19T14:40:29.227Z"),
            Description: unifiedgosdk.Pointer("Defluo."),
            ID: unifiedgosdk.Pointer("b3ba5ccf-3f87-4ae8-86a4-6adb7d2b652a"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Sollers adeo caste brevis denuo attonbitus. Amita ducimus capillus claudeo vulgaris temporibus arbustum solium. Degusto suppono commodo fuga surculus tripudio doloribus. Quaerat cauda civitas carcer alveus alienus sperno vetus. Adulescens conspergo appositus arguo. Fugiat cimentarius depono angustus somniculosus tubineus terra beatae sui."),
                    Description: unifiedgosdk.Pointer("Alias sit careo vesco tollo."),
                    Languages: []string{
                        "eo",
                    },
                    Name: unifiedgosdk.Pointer("aiunt"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/Iy2MXlem/2398/73"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/5KfxRKbzo1/2699/3216"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("spargo"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-05T02:29:09.038Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateLmsCollectionRequest](../../pkg/models/operations/updatelmscollectionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateLmsCollectionResponse](../../pkg/models/operations/updatelmscollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsContent

Updates the LmsContent identified by `id`. Provide the unified LmsContent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsContent" method="put" path="/lms/{connection_id}/content/{id}" example="lms_content" -->
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

    res, err := s.Lms.UpdateLmsContent(ctx, operations.UpdateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("3f7fe3ef-01ea-41f6-9e1a-99302ff334e5"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "campana",
                "viriliter",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Vindico vorago."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("aeger"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Sulum."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("voluptate"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Cohors necessitatibus beatae copiose copiose cedo sub decens acer. Spiritus velum adulatio commodi textus beneficium termes comedo valde caste. Sodalitas bos timor abeo cur. Sperno decet acquiro tabgo canto terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum."),
                    Description: unifiedgosdk.Pointer("Incidunt tepidus commemoro voro ocer."),
                    Languages: []string{
                        "ro",
                        "af",
                    },
                    Name: unifiedgosdk.Pointer("antepono"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/URbmlO66Wk/587/3603"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/4KFvVbCx83/2788/2850"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Ago neque qui cursim tunc barba decens. Tutamen amplus tabella tabgo infit. Demo adversus capillus aro unus. Dedico aspicio alius vehemens terminatio varietas vomica voluptatibus apud. Defero decet impedit verbera pauci laboriosam molestiae urbanus mollitia. Vergo crux certus nulla."),
                    Description: unifiedgosdk.Pointer("Bene doloribus."),
                    Languages: []string{
                        "ru",
                        "sw",
                    },
                    Name: unifiedgosdk.Pointer("totidem"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/qpRev/1286/3718"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/JOWrp5/254/2004"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Laborum adiuvo labore officia demulceo. Abduco mollitia ventosus synagoga certus magni vado cilicium stella. Velut capitulus attero odio advenio similique acsi totam atavus aufero. Ventosus campana animi."),
                    Description: unifiedgosdk.Pointer("Trucido dens pecus aer stella."),
                    Languages: []string{
                        "is",
                        "sk",
                    },
                    Name: unifiedgosdk.Pointer("territo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/81EuDedaf/557/986"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/BtV2eDuVm7/3644/1033"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("laboriosam"),
            ProviderName: unifiedgosdk.Pointer("Mraz - Keeling"),
            PublishedAt: types.MustNewTimeFromString("2022-09-15T01:48:30.270Z"),
            ShortDescription: unifiedgosdk.Pointer("Videlicet pauci."),
            Skills: []string{
                "soleo",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("clam"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "audio",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T12:58:06.325Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLmsContentRequest](../../pkg/models/operations/updatelmscontentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateLmsContentResponse](../../pkg/models/operations/updatelmscontentresponse.md), error**

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

    res, err := s.Lms.UpdateLmsCourse(ctx, operations.UpdateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "deripio",
                "tergiversatio",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("DZD"),
            Description: unifiedgosdk.Pointer("Aut."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("5b382565-417b-4497-8643-fed1c43b86b7"),
            Instructors: []shared.LmsReference{},
            IsActive: unifiedgosdk.Pointer(true),
            IsPrivate: unifiedgosdk.Pointer(false),
            Languages: []string{
                "capitulus",
                "desparatus",
                "earum",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Carbo dapifer adeptio. Ipsum utrimque quae architecto repellendus cum. Conitor anser abutor error adsuesco abeo caelestis."),
                    Description: unifiedgosdk.Pointer("Accommodo adipiscor campana aegrotatio."),
                    Languages: []string{
                        "el",
                        "hr",
                    },
                    Name: unifiedgosdk.Pointer("conduco"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/yTatRhK/3873/3840"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://picsum.photos/seed/vcfQAbs/3357/2472"),
                },
            },
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("vinco"),
            PriceAmount: unifiedgosdk.Pointer[float64](84.0),
            ProviderName: unifiedgosdk.Pointer("Roberts Group"),
            PublishedAt: types.MustNewTimeFromString("2022-11-06T12:56:32.674Z"),
            Skills: []string{
                "aveho",
                "via",
            },
            Students: []shared.LmsReference{},
            TimeEstimateMinutes: unifiedgosdk.Pointer[float64](100.0),
            UpdatedAt: types.MustNewTimeFromString("2023-02-13T00:26:39.781Z"),
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

## UpdateLmsInstructor

Updates the LmsInstructor identified by `id`. Provide the unified LmsInstructor fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsInstructor" method="put" path="/lms/{connection_id}/instructor/{id}" example="lms_instructor" -->
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

    res, err := s.Lms.UpdateLmsInstructor(ctx, operations.UpdateLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Miranda"),
            ID: unifiedgosdk.Pointer("0267df64-7a37-426a-adf5-74b4fd83d28f"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/96902788"),
            LastName: unifiedgosdk.Pointer("Turcotte"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Miranda Turcotte"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(752) 816-3330",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(415) 345-1798",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Brand Planner"),
            UpdatedAt: types.MustNewTimeFromString("2025-08-22T00:58:33.339Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateLmsInstructorRequest](../../pkg/models/operations/updatelmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateLmsInstructorResponse](../../pkg/models/operations/updatelmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsStudent

Updates the LmsStudent identified by `id`. Provide the unified LmsStudent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsStudent" method="put" path="/lms/{connection_id}/student/{id}" example="lms_student" -->
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

    res, err := s.Lms.UpdateLmsStudent(ctx, operations.UpdateLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("4610 Smith Cove"),
                City: unifiedgosdk.Pointer("Lake Milanchester"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("58634-4787"),
                Region: unifiedgosdk.Pointer("Iowa"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Lazaro"),
            ID: unifiedgosdk.Pointer("5f3e807b-a3cd-41fb-a9f5-78a4f2905d63"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/3245424"),
            LastName: unifiedgosdk.Pointer("Zulauf-Schowalter"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Lazaro Zulauf-Schowalter"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(476) 500-2870",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(338) 436-9806",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(852) 940-8205",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-14T02:26:33.989Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLmsStudentRequest](../../pkg/models/operations/updatelmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateLmsStudentResponse](../../pkg/models/operations/updatelmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |