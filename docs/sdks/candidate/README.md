# Candidate

## Overview

### Available Operations

* [CreateAtsCandidate](#createatscandidate) - Create a candidate
* [GetAtsCandidate](#getatscandidate) - Retrieve a candidate
* [ListAtsCandidates](#listatscandidates) - List all candidates
* [PatchAtsCandidate](#patchatscandidate) - Update a candidate
* [RemoveAtsCandidate](#removeatscandidate) - Remove a candidate
* [UpdateAtsCandidate](#updateatscandidate) - Update a candidate

## CreateAtsCandidate

Create a candidate

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsCandidate" method="post" path="/ats/{connection_id}/candidate" example="ats_candidate" -->
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

    res, err := s.Candidate.CreateAtsCandidate(ctx, operations.CreateAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("802 Roberts Squares"),
                Address2: unifiedgosdk.Pointer("Suite 550"),
                City: unifiedgosdk.Pointer("Lake Raeganside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("44530-0054"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Ferry, Legros and Feest"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("mouser throughout"),
                    EndAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("solutio"),
                    Institution: unifiedgosdk.Pointer("Heller - Lubowitz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("2001-03-26T08:12:11.510Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Ardith.Beatty@hotmail.com",
                    Name: unifiedgosdk.Pointer("Opal Lindgren"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Ardith_Beatty@gmail.com",
                    Name: unifiedgosdk.Pointer("Kristi Nader"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Donnelly, Buckridge and Steuber"),
                    EndAt: types.MustNewTimeFromString("1978-06-20T02:53:48.383Z"),
                    StartAt: types.MustNewTimeFromString("1980-02-06T17:16:53.798Z"),
                    Title: unifiedgosdk.Pointer("Principal Brand Strategist"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("0af97bf1-0a02-4d1d-87e9-4ee436986be2"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/40/3693?lock=5634712403880328"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://sizzling-legislature.com",
                "https://soupy-interchange.net",
                "https://troubled-substitution.info",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("1d612abc-ff4d-4002-9ac2-c413c6f92f8f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cariosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginSourced.ToPointer(),
            Skills: []string{
                "vita",
                "cohors",
            },
            Sources: []string{
                "tactus",
            },
            Summary: unifiedgosdk.Pointer("Denego barba rerum similique via templum totam suus voluptatem. Depraedor virgo cui comminor commodi curvo. Chirographum pax spero nostrum damnatio averto pecus cervus aspicio absens."),
            Tags: []string{
                "aliquid",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(779) 296-5994",
                    Type: shared.AtsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Principal Implementation Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T00:48:03.794Z"),
            WebURL: unifiedgosdk.Pointer("https://expert-lender.name/"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAtsCandidateRequest](../../pkg/models/operations/createatscandidaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateAtsCandidateResponse](../../pkg/models/operations/createatscandidateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsCandidate

Retrieve a candidate

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsCandidate" method="get" path="/ats/{connection_id}/candidate/{id}" -->
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

    res, err := s.Candidate.GetAtsCandidate(ctx, operations.GetAtsCandidateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAtsCandidateRequest](../../pkg/models/operations/getatscandidaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetAtsCandidateResponse](../../pkg/models/operations/getatscandidateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsCandidates

List all candidates

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsCandidates" method="get" path="/ats/{connection_id}/candidate" -->
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

    res, err := s.Candidate.ListAtsCandidates(ctx, operations.ListAtsCandidatesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCandidates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsCandidatesRequest](../../pkg/models/operations/listatscandidatesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsCandidatesResponse](../../pkg/models/operations/listatscandidatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsCandidate

Update a candidate

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsCandidate" method="patch" path="/ats/{connection_id}/candidate/{id}" example="ats_candidate" -->
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

    res, err := s.Candidate.PatchAtsCandidate(ctx, operations.PatchAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("802 Roberts Squares"),
                Address2: unifiedgosdk.Pointer("Suite 550"),
                City: unifiedgosdk.Pointer("Lake Raeganside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("44530-0054"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Ferry, Legros and Feest"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("mouser throughout"),
                    EndAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("solutio"),
                    Institution: unifiedgosdk.Pointer("Heller - Lubowitz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("2001-03-26T08:12:11.510Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Ardith.Beatty@hotmail.com",
                    Name: unifiedgosdk.Pointer("Opal Lindgren"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Ardith_Beatty@gmail.com",
                    Name: unifiedgosdk.Pointer("Kristi Nader"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Donnelly, Buckridge and Steuber"),
                    EndAt: types.MustNewTimeFromString("1978-06-20T02:53:48.383Z"),
                    StartAt: types.MustNewTimeFromString("1980-02-06T17:16:53.798Z"),
                    Title: unifiedgosdk.Pointer("Principal Brand Strategist"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("0f939edd-bb5f-4dbd-adf8-41f2873ff823"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/40/3693?lock=5634712403880328"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://sizzling-legislature.com",
                "https://soupy-interchange.net",
                "https://troubled-substitution.info",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b0a32bf6-b055-480b-84ef-077fc83d48f6"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cariosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginSourced.ToPointer(),
            Skills: []string{
                "vita",
                "cohors",
            },
            Sources: []string{
                "tactus",
            },
            Summary: unifiedgosdk.Pointer("Denego barba rerum similique via templum totam suus voluptatem. Depraedor virgo cui comminor commodi curvo. Chirographum pax spero nostrum damnatio averto pecus cervus aspicio absens."),
            Tags: []string{
                "aliquid",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(779) 296-5994",
                    Type: shared.AtsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Principal Implementation Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T00:48:03.805Z"),
            WebURL: unifiedgosdk.Pointer("https://expert-lender.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchAtsCandidateRequest](../../pkg/models/operations/patchatscandidaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchAtsCandidateResponse](../../pkg/models/operations/patchatscandidateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsCandidate

Remove a candidate

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsCandidate" method="delete" path="/ats/{connection_id}/candidate/{id}" -->
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

    res, err := s.Candidate.RemoveAtsCandidate(ctx, operations.RemoveAtsCandidateRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveAtsCandidateRequest](../../pkg/models/operations/removeatscandidaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveAtsCandidateResponse](../../pkg/models/operations/removeatscandidateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsCandidate

Update a candidate

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsCandidate" method="put" path="/ats/{connection_id}/candidate/{id}" example="ats_candidate" -->
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

    res, err := s.Candidate.UpdateAtsCandidate(ctx, operations.UpdateAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("802 Roberts Squares"),
                Address2: unifiedgosdk.Pointer("Suite 550"),
                City: unifiedgosdk.Pointer("Lake Raeganside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("44530-0054"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Ferry, Legros and Feest"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("mouser throughout"),
                    EndAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("solutio"),
                    Institution: unifiedgosdk.Pointer("Heller - Lubowitz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("2001-03-26T08:12:11.510Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Ardith.Beatty@hotmail.com",
                    Name: unifiedgosdk.Pointer("Opal Lindgren"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Ardith_Beatty@gmail.com",
                    Name: unifiedgosdk.Pointer("Kristi Nader"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Donnelly, Buckridge and Steuber"),
                    EndAt: types.MustNewTimeFromString("1978-06-20T02:53:48.383Z"),
                    StartAt: types.MustNewTimeFromString("1980-02-06T17:16:53.798Z"),
                    Title: unifiedgosdk.Pointer("Principal Brand Strategist"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("0f939edd-bb5f-4dbd-adf8-41f2873ff823"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/40/3693?lock=5634712403880328"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://sizzling-legislature.com",
                "https://soupy-interchange.net",
                "https://troubled-substitution.info",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b0a32bf6-b055-480b-84ef-077fc83d48f6"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cariosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginSourced.ToPointer(),
            Skills: []string{
                "vita",
                "cohors",
            },
            Sources: []string{
                "tactus",
            },
            Summary: unifiedgosdk.Pointer("Denego barba rerum similique via templum totam suus voluptatem. Depraedor virgo cui comminor commodi curvo. Chirographum pax spero nostrum damnatio averto pecus cervus aspicio absens."),
            Tags: []string{
                "aliquid",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(779) 296-5994",
                    Type: shared.AtsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Principal Implementation Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T00:48:03.805Z"),
            WebURL: unifiedgosdk.Pointer("https://expert-lender.name/"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAtsCandidateRequest](../../pkg/models/operations/updateatscandidaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAtsCandidateResponse](../../pkg/models/operations/updateatscandidateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |