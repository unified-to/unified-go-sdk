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

    res, err := s.Content.CreateLmsContent(ctx, operations.CreateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("c45188d6-708d-4b44-ac9f-f62ee7022393"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T18:29:31.319Z"),
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

    res, err := s.Content.PatchLmsContent(ctx, operations.PatchLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("5aad389e-c82c-40bb-a897-30377dd1596c"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T18:29:31.327Z"),
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

    res, err := s.Content.UpdateLmsContent(ctx, operations.UpdateLmsContentRequest{
        LmsContent: shared.LmsContent{
            Categories: []string{
                "territo",
            },
            CreatedAt: types.MustNewTimeFromString("2020-10-22T22:30:50.963Z"),
            Description: unifiedgosdk.Pointer("Usque laboriosam ventosus adflicto."),
            Difficulty: unifiedgosdk.Pointer("Beginner"),
            DurationMinutes: unifiedgosdk.Pointer[float64](19.0),
            ExternalReference: unifiedgosdk.Pointer("0d230e31-a9c4-4a35-a5b9-9168e91ffff5"),
            ID: unifiedgosdk.Pointer("5aad389e-c82c-40bb-a897-30377dd1596c"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-09-23T18:29:31.327Z"),
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