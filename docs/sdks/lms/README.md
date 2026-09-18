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

Create an activity

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
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-12T21:36:58.402Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("346f5399-7fa5-4643-bcd3-c9cedc1207c5"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T21:57:49.703Z"),
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

Create a class

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
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("6216bb21-e805-4860-8073-3f3432c321e1"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.088Z"),
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

Create a collection

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
            Description: unifiedgosdk.Pointer("Ab."),
            ID: unifiedgosdk.Pointer("80c93993-eef0-4ab6-830e-da2ab05153e6"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Accusamus earum sulum libero adficio testimonium vitae. Calcar contra vergo curis sollers. Caste brevis denuo. Tam amita ducimus capillus. Vulgaris temporibus arbustum solium id. Suppono commodo fuga surculus tripudio doloribus."),
                    Description: unifiedgosdk.Pointer("Aliquam tardus careo hic umbra."),
                    Languages: []string{
                        "gl",
                    },
                    Name: unifiedgosdk.Pointer("thymum"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/15O5EfV/2982/752"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2679/70?lock=6078357625960554"),
                },
            },
            Name: unifiedgosdk.Pointer("ara"),
            UpdatedAt: types.MustNewTimeFromString("2026-06-28T08:59:23.953Z"),
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

Create a content

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
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("66cad5f0-7b3f-4edb-b8ce-40e7f385b6fa"),
            Instructors: []shared.LmsReference{
                shared.LmsReference{
                    ID: unifiedgosdk.Pointer("91a23b20-a7a3-4323-9548-0897c09eb49e"),
                    Name: unifiedgosdk.Pointer("Winston Ferry"),
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "despecto",
                "suppellex",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Numquam."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("validus"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Callide."),
                    Language: unifiedgosdk.Pointer("fr"),
                    Name: unifiedgosdk.Pointer("crux"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adnuo antepono vulticulus incidunt. Commemoro voro ocer arca velut vigor venustas una audeo. Consectetur totidem amor amo vigor. Patior ulciscor cohors necessitatibus beatae. Copiose cedo sub decens acer."),
                    Description: unifiedgosdk.Pointer("Venia aeternus tandem spargo."),
                    Languages: []string{
                        "zu",
                        "ba",
                    },
                    Name: unifiedgosdk.Pointer("subiungo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/2056/3712?lock=5644845642923518"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2593/1553?lock=8591263400111785"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum. Absque adversus desino aut tabernus tollo vigor. Cur agnitio totidem consuasor bene doloribus. Vetus abundans curriculum curatio usitas."),
                    Description: unifiedgosdk.Pointer("Comedo valde caste combibo."),
                    Languages: []string{
                        "it",
                        "hu",
                    },
                    Name: unifiedgosdk.Pointer("beneficium"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/pNFr1/2597/885"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3597/239?lock=7142808124990633"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Coepi demo adversus capillus aro unus templum. Aspicio alius vehemens terminatio varietas vomica. Apud thorax defero decet impedit verbera pauci laboriosam molestiae urbanus. Aveho vergo crux certus nulla caute sophismata. Caute voluptatibus patrocinor demo verumtamen adeo canis sapiente depraedor verus."),
                    Description: unifiedgosdk.Pointer("Tunc barba decens."),
                    Languages: []string{
                        "bn",
                        "yo",
                    },
                    Name: unifiedgosdk.Pointer("qui"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/1375/3377?lock=6601832177607674"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3927/2086?lock=5199784913821481"),
                },
            },
            Name: unifiedgosdk.Pointer("ut"),
            ProviderName: unifiedgosdk.Pointer("Berge LLC"),
            PublishedAt: types.MustNewTimeFromString("2023-11-08T11:32:09.080Z"),
            ShortDescription: unifiedgosdk.Pointer("Commemoro."),
            Skills: []string{
                "trucido",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("tibi"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "dens",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T11:33:10.875Z"),
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

    res, err := s.Lms.CreateLmsCourse(ctx, operations.CreateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("6766c05c-539e-4557-baf7-0987631b8b0f"),
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
            UpdatedAt: types.MustNewTimeFromString("2023-02-06T22:35:58.033Z"),
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

Create an instructor

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
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("2abdf9c3-7c43-490d-8aa8-9f34d0631e20"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/20232618"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(352) 551-7989",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Product Solutions Engineer"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T14:52:02.059Z"),
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

Create a student

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
                Address1: unifiedgosdk.Pointer("94082 Kassandra Camp"),
                Address2: unifiedgosdk.Pointer("Apt. 461"),
                City: unifiedgosdk.Pointer("New Ibrahimmouth"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("52851"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("d4188163-3b82-495e-af69-58af5d991f39"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/36301374"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(482) 469-8067",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-06-19T14:16:48.184Z"),
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

Retrieve an activity

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

Retrieve a class

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

Retrieve a collection

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

Retrieve a content

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

Retrieve an instructor

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

Retrieve a student

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

List all activities

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

List all classes

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

List all collections

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

List all contents

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

List all instructors

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

List all students

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

Update an activity

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
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-12T21:36:58.404Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("45739c5e-b76c-4232-93a4-f9424dded91a"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T21:57:49.704Z"),
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

Update a class

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
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("fcbd51ed-fc99-4bce-9a89-92b0c5dc6a74"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.098Z"),
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

Update a collection

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
            Description: unifiedgosdk.Pointer("Ab."),
            ID: unifiedgosdk.Pointer("ce0158a5-5bf7-4d3c-b207-94ac9e643014"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Accusamus earum sulum libero adficio testimonium vitae. Calcar contra vergo curis sollers. Caste brevis denuo. Tam amita ducimus capillus. Vulgaris temporibus arbustum solium id. Suppono commodo fuga surculus tripudio doloribus."),
                    Description: unifiedgosdk.Pointer("Aliquam tardus careo hic umbra."),
                    Languages: []string{
                        "gl",
                    },
                    Name: unifiedgosdk.Pointer("thymum"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/15O5EfV/2982/752"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2679/70?lock=6078357625960554"),
                },
            },
            Name: unifiedgosdk.Pointer("ara"),
            UpdatedAt: types.MustNewTimeFromString("2026-06-28T08:59:23.963Z"),
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

Update a content

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
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("85cb2e02-5568-4bc8-b95f-ec6ae95ef317"),
            Instructors: []shared.LmsReference{
                shared.LmsReference{
                    ID: unifiedgosdk.Pointer("91a23b20-a7a3-4323-9548-0897c09eb49e"),
                    Name: unifiedgosdk.Pointer("Winston Ferry"),
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "despecto",
                "suppellex",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Numquam."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("validus"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Callide."),
                    Language: unifiedgosdk.Pointer("fr"),
                    Name: unifiedgosdk.Pointer("crux"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adnuo antepono vulticulus incidunt. Commemoro voro ocer arca velut vigor venustas una audeo. Consectetur totidem amor amo vigor. Patior ulciscor cohors necessitatibus beatae. Copiose cedo sub decens acer."),
                    Description: unifiedgosdk.Pointer("Venia aeternus tandem spargo."),
                    Languages: []string{
                        "zu",
                        "ba",
                    },
                    Name: unifiedgosdk.Pointer("subiungo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/2056/3712?lock=5644845642923518"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2593/1553?lock=8591263400111785"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum. Absque adversus desino aut tabernus tollo vigor. Cur agnitio totidem consuasor bene doloribus. Vetus abundans curriculum curatio usitas."),
                    Description: unifiedgosdk.Pointer("Comedo valde caste combibo."),
                    Languages: []string{
                        "it",
                        "hu",
                    },
                    Name: unifiedgosdk.Pointer("beneficium"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/pNFr1/2597/885"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3597/239?lock=7142808124990633"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Coepi demo adversus capillus aro unus templum. Aspicio alius vehemens terminatio varietas vomica. Apud thorax defero decet impedit verbera pauci laboriosam molestiae urbanus. Aveho vergo crux certus nulla caute sophismata. Caute voluptatibus patrocinor demo verumtamen adeo canis sapiente depraedor verus."),
                    Description: unifiedgosdk.Pointer("Tunc barba decens."),
                    Languages: []string{
                        "bn",
                        "yo",
                    },
                    Name: unifiedgosdk.Pointer("qui"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/1375/3377?lock=6601832177607674"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3927/2086?lock=5199784913821481"),
                },
            },
            Name: unifiedgosdk.Pointer("ut"),
            ProviderName: unifiedgosdk.Pointer("Berge LLC"),
            PublishedAt: types.MustNewTimeFromString("2023-11-08T11:32:09.080Z"),
            ShortDescription: unifiedgosdk.Pointer("Commemoro."),
            Skills: []string{
                "trucido",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("tibi"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "dens",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T11:33:10.882Z"),
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

    res, err := s.Lms.PatchLmsCourse(ctx, operations.PatchLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("03631626-cd93-49ba-b513-beb4ffc43937"),
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
            UpdatedAt: types.MustNewTimeFromString("2023-02-06T22:35:58.035Z"),
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

Update an instructor

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
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("d05a1560-35c6-418a-81e7-6746e39c822e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/20232618"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(352) 551-7989",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Product Solutions Engineer"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T14:52:02.064Z"),
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

Update a student

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
                Address1: unifiedgosdk.Pointer("94082 Kassandra Camp"),
                Address2: unifiedgosdk.Pointer("Apt. 461"),
                City: unifiedgosdk.Pointer("New Ibrahimmouth"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("52851"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("9cd03af6-3555-4f93-92fb-b2b9423568d8"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/36301374"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(482) 469-8067",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-06-19T14:16:48.187Z"),
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

Remove an activity

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

Remove a class

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

Remove a collection

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

Remove a content

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

Remove an instructor

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

Remove a student

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

Update an activity

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
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-12T21:36:58.404Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("45739c5e-b76c-4232-93a4-f9424dded91a"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-23T21:57:49.704Z"),
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

Update a class

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
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("fcbd51ed-fc99-4bce-9a89-92b0c5dc6a74"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.098Z"),
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

Update a collection

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
            Description: unifiedgosdk.Pointer("Ab."),
            ID: unifiedgosdk.Pointer("ce0158a5-5bf7-4d3c-b207-94ac9e643014"),
            IsActive: unifiedgosdk.Pointer(true),
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Accusamus earum sulum libero adficio testimonium vitae. Calcar contra vergo curis sollers. Caste brevis denuo. Tam amita ducimus capillus. Vulgaris temporibus arbustum solium id. Suppono commodo fuga surculus tripudio doloribus."),
                    Description: unifiedgosdk.Pointer("Aliquam tardus careo hic umbra."),
                    Languages: []string{
                        "gl",
                    },
                    Name: unifiedgosdk.Pointer("thymum"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/15O5EfV/2982/752"),
                    Type: shared.LmsMediaTypeHeadshot.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2679/70?lock=6078357625960554"),
                },
            },
            Name: unifiedgosdk.Pointer("ara"),
            UpdatedAt: types.MustNewTimeFromString("2026-06-28T08:59:23.963Z"),
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

Update a content

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
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("85cb2e02-5568-4bc8-b95f-ec6ae95ef317"),
            Instructors: []shared.LmsReference{
                shared.LmsReference{
                    ID: unifiedgosdk.Pointer("91a23b20-a7a3-4323-9548-0897c09eb49e"),
                    Name: unifiedgosdk.Pointer("Winston Ferry"),
                },
            },
            IsActive: unifiedgosdk.Pointer(true),
            Languages: []string{
                "despecto",
                "suppellex",
            },
            Localizations: []shared.LmsContentLocalization{
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Numquam."),
                    Language: unifiedgosdk.Pointer("es"),
                    Name: unifiedgosdk.Pointer("validus"),
                },
                shared.LmsContentLocalization{
                    Description: unifiedgosdk.Pointer("Callide."),
                    Language: unifiedgosdk.Pointer("fr"),
                    Name: unifiedgosdk.Pointer("crux"),
                },
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Adnuo antepono vulticulus incidunt. Commemoro voro ocer arca velut vigor venustas una audeo. Consectetur totidem amor amo vigor. Patior ulciscor cohors necessitatibus beatae. Copiose cedo sub decens acer."),
                    Description: unifiedgosdk.Pointer("Venia aeternus tandem spargo."),
                    Languages: []string{
                        "zu",
                        "ba",
                    },
                    Name: unifiedgosdk.Pointer("subiungo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/2056/3712?lock=5644845642923518"),
                    Type: shared.LmsMediaTypeOther.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/2593/1553?lock=8591263400111785"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Terminatio acidus tripudio teres. Compono demum aut aestivus ambitus pecus tergum verumtamen vestrum. Absque adversus desino aut tabernus tollo vigor. Cur agnitio totidem consuasor bene doloribus. Vetus abundans curriculum curatio usitas."),
                    Description: unifiedgosdk.Pointer("Comedo valde caste combibo."),
                    Languages: []string{
                        "it",
                        "hu",
                    },
                    Name: unifiedgosdk.Pointer("beneficium"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://picsum.photos/seed/pNFr1/2597/885"),
                    Type: shared.LmsMediaTypeWeb.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3597/239?lock=7142808124990633"),
                },
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Coepi demo adversus capillus aro unus templum. Aspicio alius vehemens terminatio varietas vomica. Apud thorax defero decet impedit verbera pauci laboriosam molestiae urbanus. Aveho vergo crux certus nulla caute sophismata. Caute voluptatibus patrocinor demo verumtamen adeo canis sapiente depraedor verus."),
                    Description: unifiedgosdk.Pointer("Tunc barba decens."),
                    Languages: []string{
                        "bn",
                        "yo",
                    },
                    Name: unifiedgosdk.Pointer("qui"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/1375/3377?lock=6601832177607674"),
                    Type: shared.LmsMediaTypeImage.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/3927/2086?lock=5199784913821481"),
                },
            },
            Name: unifiedgosdk.Pointer("ut"),
            ProviderName: unifiedgosdk.Pointer("Berge LLC"),
            PublishedAt: types.MustNewTimeFromString("2023-11-08T11:32:09.080Z"),
            ShortDescription: unifiedgosdk.Pointer("Commemoro."),
            Skills: []string{
                "trucido",
            },
            SortOrder: unifiedgosdk.Pointer[float64](3.0),
            Subjects: []shared.LmsSubject{
                shared.LmsSubject{
                    Name: unifiedgosdk.Pointer("tibi"),
                    Rank: unifiedgosdk.Pointer[float64](1.0),
                },
            },
            Tags: []string{
                "dens",
            },
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T11:33:10.882Z"),
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

    res, err := s.Lms.UpdateLmsCourse(ctx, operations.UpdateLmsCourseRequest{
        LmsCourse: shared.LmsCourse{
            Categories: []string{
                "tergiversatio",
                "tumultus",
            },
            CreatedAt: types.MustNewTimeFromString("2022-10-06T09:58:53.559Z"),
            Currency: unifiedgosdk.Pointer("FJD"),
            Description: unifiedgosdk.Pointer("Vinco alias aut capitulus."),
            DurationMinutes: unifiedgosdk.Pointer[float64](148.0),
            ID: unifiedgosdk.Pointer("03631626-cd93-49ba-b513-beb4ffc43937"),
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
            UpdatedAt: types.MustNewTimeFromString("2023-02-06T22:35:58.035Z"),
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

Update an instructor

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
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("d05a1560-35c6-418a-81e7-6746e39c822e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/20232618"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(352) 551-7989",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Product Solutions Engineer"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-29T14:52:02.064Z"),
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

Update a student

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
                Address1: unifiedgosdk.Pointer("94082 Kassandra Camp"),
                Address2: unifiedgosdk.Pointer("Apt. 461"),
                City: unifiedgosdk.Pointer("New Ibrahimmouth"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("52851"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("9cd03af6-3555-4f93-92fb-b2b9423568d8"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/36301374"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(482) 469-8067",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-06-19T14:16:48.187Z"),
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