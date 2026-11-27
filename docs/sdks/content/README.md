# Content

## Overview

### Available Operations

* [CreateLmsContent](#createlmscontent) - Create a content
* [GetLmsContent](#getlmscontent) - Retrieve a content
* [ListLmsContents](#listlmscontents) - List all contents
* [PatchLmsContent](#patchlmscontent) - Update a content
* [RemoveLmsContent](#removelmscontent) - Remove a content
* [UpdateLmsContent](#updatelmscontent) - Update a content

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

    res, err := s.Content.CreateLmsContent(ctx, operations.CreateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("a782415e-f468-435f-8205-27cc2c980857"),
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
            PublishedAt: types.MustNewTimeFromString("2022-09-14T18:24:28.167Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T05:13:53.055Z"),
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

    res, err := s.Content.GetLmsContent(ctx, operations.GetLmsContentRequest{
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

    res, err := s.Content.ListLmsContents(ctx, operations.ListLmsContentsRequest{
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

    res, err := s.Content.PatchLmsContent(ctx, operations.PatchLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("24424297-71fd-4051-880a-d90346131054"),
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
            PublishedAt: types.MustNewTimeFromString("2022-09-14T18:24:28.176Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T05:13:53.064Z"),
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

    res, err := s.Content.RemoveLmsContent(ctx, operations.RemoveLmsContentRequest{
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

    res, err := s.Content.UpdateLmsContent(ctx, operations.UpdateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "tergum",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Adflicto despecto suppellex socius cresco."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ID: unifiedgosdk.Pointer("24424297-71fd-4051-880a-d90346131054"),
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
            PublishedAt: types.MustNewTimeFromString("2022-09-14T18:24:28.176Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-10-16T05:13:53.064Z"),
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