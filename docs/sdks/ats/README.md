# Ats

## Overview

### Available Operations

* [CreateAtsActivity](#createatsactivity) - Create an activity
* [CreateAtsApplication](#createatsapplication) - Create an application
* [CreateAtsCandidate](#createatscandidate) - Create a candidate
* [CreateAtsCompany](#createatscompany) - Create a company
* [CreateAtsDocument](#createatsdocument) - Create a document
* [CreateAtsInterview](#createatsinterview) - Create an interview
* [CreateAtsJob](#createatsjob) - Create a job
* [CreateAtsScorecard](#createatsscorecard) - Create a scorecard
* [GetAtsActivity](#getatsactivity) - Retrieve an activity
* [GetAtsApplication](#getatsapplication) - Retrieve an application
* [GetAtsApplicationstatus](#getatsapplicationstatus) - Retrieve an applicationstatus
* [GetAtsCandidate](#getatscandidate) - Retrieve a candidate
* [GetAtsCompany](#getatscompany) - Retrieve a company
* [GetAtsDocument](#getatsdocument) - Retrieve a document
* [GetAtsInterview](#getatsinterview) - Retrieve an interview
* [GetAtsJob](#getatsjob) - Retrieve a job
* [GetAtsScorecard](#getatsscorecard) - Retrieve a scorecard
* [ListAtsActivities](#listatsactivities) - List all activities
* [ListAtsApplications](#listatsapplications) - List all applications
* [ListAtsApplicationstatuses](#listatsapplicationstatuses) - List all applicationstatuses
* [ListAtsCandidates](#listatscandidates) - List all candidates
* [ListAtsCompanies](#listatscompanies) - List all companies
* [ListAtsDocuments](#listatsdocuments) - List all documents
* [ListAtsInterviews](#listatsinterviews) - List all interviews
* [ListAtsJobs](#listatsjobs) - List all jobs
* [ListAtsScorecards](#listatsscorecards) - List all scorecards
* [PatchAtsActivity](#patchatsactivity) - Update an activity
* [PatchAtsApplication](#patchatsapplication) - Update an application
* [PatchAtsCandidate](#patchatscandidate) - Update a candidate
* [PatchAtsCompany](#patchatscompany) - Update a company
* [PatchAtsDocument](#patchatsdocument) - Update a document
* [PatchAtsInterview](#patchatsinterview) - Update an interview
* [PatchAtsJob](#patchatsjob) - Update a job
* [PatchAtsScorecard](#patchatsscorecard) - Update a scorecard
* [RemoveAtsActivity](#removeatsactivity) - Remove an activity
* [RemoveAtsApplication](#removeatsapplication) - Remove an application
* [RemoveAtsCandidate](#removeatscandidate) - Remove a candidate
* [RemoveAtsCompany](#removeatscompany) - Remove a company
* [RemoveAtsDocument](#removeatsdocument) - Remove a document
* [RemoveAtsInterview](#removeatsinterview) - Remove an interview
* [RemoveAtsJob](#removeatsjob) - Remove a job
* [RemoveAtsScorecard](#removeatsscorecard) - Remove a scorecard
* [UpdateAtsActivity](#updateatsactivity) - Update an activity
* [UpdateAtsApplication](#updateatsapplication) - Update an application
* [UpdateAtsCandidate](#updateatscandidate) - Update a candidate
* [UpdateAtsCompany](#updateatscompany) - Update a company
* [UpdateAtsDocument](#updateatsdocument) - Update a document
* [UpdateAtsInterview](#updateatsinterview) - Update an interview
* [UpdateAtsJob](#updateatsjob) - Update a job
* [UpdateAtsScorecard](#updateatsscorecard) - Update a scorecard

## CreateAtsActivity

Create an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsActivity" method="post" path="/ats/{connection_id}/activity" example="ats_activity" -->
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

    res, err := s.Ats.CreateAtsActivity(ctx, operations.CreateAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Sapiente."),
            ID: unifiedgosdk.Pointer("e8a2f1e7-5ef1-44d6-b9b0-dd4ae05406d4"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("e90754bb-3fc5-46ca-87c0-2ef0049e309a"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("cimentarius"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Reprehenderit verto.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Venio blandior aureus."),
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-04T06:38:15.355Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAtsActivityRequest](../../pkg/models/operations/createatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAtsActivityResponse](../../pkg/models/operations/createatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAtsApplication

Create an application

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsApplication" method="post" path="/ats/{connection_id}/application" example="ats_application" -->
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

    res, err := s.Ats.CreateAtsApplication(ctx, operations.CreateAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-09-29T18:41:44.737Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-05-12T18:30:35.439Z"),
            ID: unifiedgosdk.Pointer("26ec1638-c52c-492b-9fd7-09de41cb00e7"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("c5228933-7b14-40ad-9f1b-19a455f0df4a"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-10-11T13:07:05.713Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-10-18T09:21:22.037Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateAtsApplicationRequest](../../pkg/models/operations/createatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateAtsApplicationResponse](../../pkg/models/operations/createatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Ats.CreateAtsCandidate(ctx, operations.CreateAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("6790 S 1st Street"),
                City: unifiedgosdk.Pointer("Ferrystad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("91605"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("HI"),
            },
            CompanyName: unifiedgosdk.Pointer("Lehner and Sons"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            DateOfBirth: types.MustNewTimeFromString("1994-07-05T19:48:57.631Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("emphasize tricky"),
                    EndAt: types.MustNewTimeFromString("1978-04-14T07:58:36.326Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("curtus"),
                    Institution: unifiedgosdk.Pointer("Abbott - Batz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("1994-06-09T13:31:39.059Z"),
                },
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("blah cappelletti"),
                    EndAt: types.MustNewTimeFromString("1989-09-06T03:32:39.901Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("crustulum"),
                    Institution: unifiedgosdk.Pointer("Herman, Franey and Lebsack"),
                    Level: unifiedgosdk.Pointer("highschool"),
                    StartAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "ardith.beatty@outlook.com",
                    Name: unifiedgosdk.Pointer("Ardith Beatty"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Kutch, Lind and Marks"),
                    EndAt: types.MustNewTimeFromString("1974-03-01T17:02:58.824Z"),
                    StartAt: types.MustNewTimeFromString("1999-09-21T22:12:51.663Z"),
                    Title: unifiedgosdk.Pointer("Lead Security Technician"),
                },
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Trantow, Sipes and Witting"),
                    EndAt: types.MustNewTimeFromString("2004-10-14T14:01:00.717Z"),
                    StartAt: types.MustNewTimeFromString("1996-12-15T04:12:04.582Z"),
                    Title: unifiedgosdk.Pointer("Regional Applications Planner"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("664f6e73-f63a-457c-8af7-87c2ab8ec9c2"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/27464445"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/ardith-beatty",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7971fe19-2a15-4b5f-9240-5511e39c68cd"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "officia",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginAgency.ToPointer(),
            Skills: []string{
                "trepide",
                "voveo",
            },
            Sources: []string{
                "vesper",
                "utroque",
            },
            Summary: unifiedgosdk.Pointer("Laborum trado carcer delibero patruus totam. Vita cohors cariosus trans denego barba rerum similique. Templum totam suus voluptatem cui depraedor virgo cui comminor commodi."),
            Tags: []string{
                "voro",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(837) 251-3896",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Division Producer"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T17:47:32.446Z"),
            WebURL: unifiedgosdk.Pointer("https://fragrant-husband.net/"),
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

## CreateAtsCompany

Create a company

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsCompany" method="post" path="/ats/{connection_id}/company" example="ats_company" -->
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

    res, err := s.Ats.CreateAtsCompany(ctx, operations.CreateAtsCompanyRequest{
        AtsCompany: shared.AtsCompany{
            CreatedAt: types.MustNewTimeFromString("2019-04-22T03:50:02.920Z"),
            ID: unifiedgosdk.Pointer("bc88c053-d149-4889-9144-a45bbbbb011e"),
            Name: unifiedgosdk.Pointer("Gulgowski, Dibbert and Wilderman"),
            Phone: unifiedgosdk.Pointer("1-882-360-2010"),
            UpdatedAt: types.MustNewTimeFromString("2020-10-01T04:12:23.715Z"),
            WebsiteURL: unifiedgosdk.Pointer("https://www.gulgowskidibbert.com"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateAtsCompanyRequest](../../pkg/models/operations/createatscompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateAtsCompanyResponse](../../pkg/models/operations/createatscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAtsDocument

Create a document

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsDocument" method="post" path="/ats/{connection_id}/document" example="ats_document" -->
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

    res, err := s.Ats.CreateAtsDocument(ctx, operations.CreateAtsDocumentRequest{
        AtsDocument: shared.AtsDocument{
            CreatedAt: types.MustNewTimeFromString("2021-08-20T08:00:27.437Z"),
            DocumentURL: unifiedgosdk.Pointer("https://vengeful-lashes.biz"),
            Filename: unifiedgosdk.Pointer("lonzo_schowalter_resume.pdf"),
            ID: unifiedgosdk.Pointer("29c939d9-2af8-44c4-a7d5-67e03b8e6897"),
            Metadata: []shared.AtsMetadata{},
            Type: shared.AtsDocumentTypeCoverLetter.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-12-07T04:59:55.140Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAtsDocumentRequest](../../pkg/models/operations/createatsdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAtsDocumentResponse](../../pkg/models/operations/createatsdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAtsInterview

Create an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsInterview" method="post" path="/ats/{connection_id}/interview" example="ats_interview" -->
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

    res, err := s.Ats.CreateAtsInterview(ctx, operations.CreateAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2021-12-08T06:33:19.583Z"),
            ID: unifiedgosdk.Pointer("75d45222-2444-450b-8db6-d716949cdd12"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            MeetingURL: unifiedgosdk.Pointer("https://immaculate-foodstuffs.biz"),
            Metadata: []shared.AtsMetadata{},
            StartAt: types.MustNewTimeFromString("2021-12-08T06:03:19.583Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-09-23T16:39:11.717Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAtsInterviewRequest](../../pkg/models/operations/createatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateAtsInterviewResponse](../../pkg/models/operations/createatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAtsJob

Create a job

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsJob" method="post" path="/ats/{connection_id}/job" example="ats_job" -->
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

    res, err := s.Ats.CreateAtsJob(ctx, operations.CreateAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("98097 Carlo Trail"),
                    City: unifiedgosdk.Pointer("South Judd"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("89776-0669"),
                    Region: unifiedgosdk.Pointer("Mississippi"),
                    RegionCode: unifiedgosdk.Pointer("FL"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("AUD"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("MZN"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Global"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("f1ed3a25-8406-4efd-afcf-0a385a1fadb2"),
            Industry: unifiedgosdk.Pointer("Handcrafted Bamboo Soap"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6413a6cc-6a8e-4306-8822-db3f3900f798"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "desolo",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Forward Brand Producer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Admoveo trado textilis."),
                    OpenedAt: types.MustNewTimeFromString("2026-06-07T21:28:07.896Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("75008 Medhurst Road"),
                        City: unifiedgosdk.Pointer("Madelynnshire"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("98754"),
                        Region: unifiedgosdk.Pointer("Wisconsin"),
                        RegionCode: unifiedgosdk.Pointer("CO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-08-02T00:29:26.215Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("b6f8c91a-7f05-4e3b-ac0a-c79968a8a9af"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-08-28T07:52:25.733Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("81736c36-1b8a-4700-a981-c384da1e68f5"),
                    MaxLength: unifiedgosdk.Pointer[float64](117.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "tactus",
                        "vehemens",
                    },
                    Prompt: unifiedgosdk.Pointer("Totam maxime tres tergiversatio."),
                    Question: "Doloribus mollitia.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("fcd02da9-266c-4de9-a918-9a98f813e617"),
                    MaxLength: unifiedgosdk.Pointer[float64](453.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "constans",
                    },
                    Prompt: unifiedgosdk.Pointer("Veritatis sodalitas nemo natus attonbitus reprehenderit voro depono."),
                    Question: "Nam.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    ID: unifiedgosdk.Pointer("e29c1433-3d51-4fa1-8492-3354e0f8e77d"),
                    MaxLength: unifiedgosdk.Pointer[float64](66.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "repellendus",
                    },
                    Prompt: unifiedgosdk.Pointer("Tum tamisium autus."),
                    Question: "Confero trepide.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "suadeo",
                "utrimque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crepusculum adsuesco vere."),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T09:47:18.130Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateAtsJobRequest](../../pkg/models/operations/createatsjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateAtsJobResponse](../../pkg/models/operations/createatsjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAtsScorecard

Create a scorecard

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsScorecard" method="post" path="/ats/{connection_id}/scorecard" example="ats_scorecard" -->
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

    res, err := s.Ats.CreateAtsScorecard(ctx, operations.CreateAtsScorecardRequest{
        AtsScorecard: shared.AtsScorecard{
            Comment: unifiedgosdk.Pointer("Maiores enim."),
            CreatedAt: types.MustNewTimeFromString("2022-02-20T17:09:45.498Z"),
            ID: unifiedgosdk.Pointer("2dea7d23-46fb-48d0-a30a-b909d9a83d91"),
            Metadata: []shared.AtsMetadata{},
            Questions: []shared.AtsScorecardQuestion{
                shared.AtsScorecardQuestion{
                    Description: unifiedgosdk.Pointer("Sulum textor eveniet facere vita."),
                    Text: "Aliquam.",
                },
                shared.AtsScorecardQuestion{
                    Answer: unifiedgosdk.Pointer("Decretum."),
                    Description: unifiedgosdk.Pointer("Conatus cicuta doloremque statua bonus."),
                    Text: "Pecto vulpes libero vomer comburo.",
                },
            },
            Recommendation: shared.RecommendationStrongYes.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-06-05T13:40:04.404Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAtsScorecardRequest](../../pkg/models/operations/createatsscorecardrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateAtsScorecardResponse](../../pkg/models/operations/createatsscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsActivity

Retrieve an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsActivity" method="get" path="/ats/{connection_id}/activity/{id}" -->
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

    res, err := s.Ats.GetAtsActivity(ctx, operations.GetAtsActivityRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAtsActivityRequest](../../pkg/models/operations/getatsactivityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAtsActivityResponse](../../pkg/models/operations/getatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsApplication

Retrieve an application

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsApplication" method="get" path="/ats/{connection_id}/application/{id}" -->
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

    res, err := s.Ats.GetAtsApplication(ctx, operations.GetAtsApplicationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAtsApplicationRequest](../../pkg/models/operations/getatsapplicationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetAtsApplicationResponse](../../pkg/models/operations/getatsapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsApplicationstatus

Retrieve an applicationstatus

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsApplicationstatus" method="get" path="/ats/{connection_id}/applicationstatus/{id}" -->
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

    res, err := s.Ats.GetAtsApplicationstatus(ctx, operations.GetAtsApplicationstatusRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetAtsApplicationstatusRequest](../../pkg/models/operations/getatsapplicationstatusrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAtsApplicationstatusResponse](../../pkg/models/operations/getatsapplicationstatusresponse.md), error**

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

    res, err := s.Ats.GetAtsCandidate(ctx, operations.GetAtsCandidateRequest{
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

## GetAtsCompany

Retrieve a company

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsCompany" method="get" path="/ats/{connection_id}/company/{id}" -->
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

    res, err := s.Ats.GetAtsCompany(ctx, operations.GetAtsCompanyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAtsCompanyRequest](../../pkg/models/operations/getatscompanyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetAtsCompanyResponse](../../pkg/models/operations/getatscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsDocument

Retrieve a document

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsDocument" method="get" path="/ats/{connection_id}/document/{id}" -->
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

    res, err := s.Ats.GetAtsDocument(ctx, operations.GetAtsDocumentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAtsDocumentRequest](../../pkg/models/operations/getatsdocumentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAtsDocumentResponse](../../pkg/models/operations/getatsdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsInterview

Retrieve an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsInterview" method="get" path="/ats/{connection_id}/interview/{id}" -->
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

    res, err := s.Ats.GetAtsInterview(ctx, operations.GetAtsInterviewRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAtsInterviewRequest](../../pkg/models/operations/getatsinterviewrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetAtsInterviewResponse](../../pkg/models/operations/getatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsJob

Retrieve a job

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsJob" method="get" path="/ats/{connection_id}/job/{id}" -->
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

    res, err := s.Ats.GetAtsJob(ctx, operations.GetAtsJobRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetAtsJobRequest](../../pkg/models/operations/getatsjobrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetAtsJobResponse](../../pkg/models/operations/getatsjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsScorecard

Retrieve a scorecard

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsScorecard" method="get" path="/ats/{connection_id}/scorecard/{id}" -->
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

    res, err := s.Ats.GetAtsScorecard(ctx, operations.GetAtsScorecardRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAtsScorecardRequest](../../pkg/models/operations/getatsscorecardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetAtsScorecardResponse](../../pkg/models/operations/getatsscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsActivities

List all activities

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsActivities" method="get" path="/ats/{connection_id}/activity" -->
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

    res, err := s.Ats.ListAtsActivities(ctx, operations.ListAtsActivitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsActivitiesRequest](../../pkg/models/operations/listatsactivitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsActivitiesResponse](../../pkg/models/operations/listatsactivitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsApplications

List all applications

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsApplications" method="get" path="/ats/{connection_id}/application" -->
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

    res, err := s.Ats.ListAtsApplications(ctx, operations.ListAtsApplicationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplications != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAtsApplicationsRequest](../../pkg/models/operations/listatsapplicationsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAtsApplicationsResponse](../../pkg/models/operations/listatsapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsApplicationstatuses

List all applicationstatuses

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsApplicationstatuses" method="get" path="/ats/{connection_id}/applicationstatus" -->
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

    res, err := s.Ats.ListAtsApplicationstatuses(ctx, operations.ListAtsApplicationstatusesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListAtsApplicationstatusesRequest](../../pkg/models/operations/listatsapplicationstatusesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListAtsApplicationstatusesResponse](../../pkg/models/operations/listatsapplicationstatusesresponse.md), error**

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

    res, err := s.Ats.ListAtsCandidates(ctx, operations.ListAtsCandidatesRequest{
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

## ListAtsCompanies

List all companies

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsCompanies" method="get" path="/ats/{connection_id}/company" -->
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

    res, err := s.Ats.ListAtsCompanies(ctx, operations.ListAtsCompaniesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAtsCompaniesRequest](../../pkg/models/operations/listatscompaniesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAtsCompaniesResponse](../../pkg/models/operations/listatscompaniesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsDocuments

List all documents

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsDocuments" method="get" path="/ats/{connection_id}/document" -->
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

    res, err := s.Ats.ListAtsDocuments(ctx, operations.ListAtsDocumentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsDocuments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAtsDocumentsRequest](../../pkg/models/operations/listatsdocumentsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAtsDocumentsResponse](../../pkg/models/operations/listatsdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsInterviews

List all interviews

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsInterviews" method="get" path="/ats/{connection_id}/interview" -->
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

    res, err := s.Ats.ListAtsInterviews(ctx, operations.ListAtsInterviewsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsInterviewsRequest](../../pkg/models/operations/listatsinterviewsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsInterviewsResponse](../../pkg/models/operations/listatsinterviewsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsJobs

List all jobs

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsJobs" method="get" path="/ats/{connection_id}/job" -->
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

    res, err := s.Ats.ListAtsJobs(ctx, operations.ListAtsJobsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListAtsJobsRequest](../../pkg/models/operations/listatsjobsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListAtsJobsResponse](../../pkg/models/operations/listatsjobsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsScorecards

List all scorecards

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsScorecards" method="get" path="/ats/{connection_id}/scorecard" -->
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

    res, err := s.Ats.ListAtsScorecards(ctx, operations.ListAtsScorecardsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsScorecards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsScorecardsRequest](../../pkg/models/operations/listatsscorecardsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsScorecardsResponse](../../pkg/models/operations/listatsscorecardsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsActivity" method="patch" path="/ats/{connection_id}/activity/{id}" example="ats_activity" -->
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

    res, err := s.Ats.PatchAtsActivity(ctx, operations.PatchAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Sapiente."),
            ID: unifiedgosdk.Pointer("1dd88965-536d-4641-a508-0b97b622bef5"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("26853e59-8a39-416d-88e8-757699c82c57"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("cimentarius"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Reprehenderit verto.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Venio blandior aureus."),
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-04T06:38:15.375Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAtsActivityRequest](../../pkg/models/operations/patchatsactivityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAtsActivityResponse](../../pkg/models/operations/patchatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsApplication

Update an application

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsApplication" method="patch" path="/ats/{connection_id}/application/{id}" example="ats_application" -->
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

    res, err := s.Ats.PatchAtsApplication(ctx, operations.PatchAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-09-29T18:41:44.750Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-05-12T18:30:35.456Z"),
            ID: unifiedgosdk.Pointer("95d97fad-3b9a-4527-bb22-1953ed79ef05"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("382ffa38-ae9e-4b90-9c6f-f3dacf302b7c"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-10-11T13:07:05.733Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-10-18T09:21:22.057Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchAtsApplicationRequest](../../pkg/models/operations/patchatsapplicationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchAtsApplicationResponse](../../pkg/models/operations/patchatsapplicationresponse.md), error**

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

    res, err := s.Ats.PatchAtsCandidate(ctx, operations.PatchAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("6790 S 1st Street"),
                City: unifiedgosdk.Pointer("Ferrystad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("91605"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("HI"),
            },
            CompanyName: unifiedgosdk.Pointer("Lehner and Sons"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            DateOfBirth: types.MustNewTimeFromString("1994-07-05T19:48:57.631Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("emphasize tricky"),
                    EndAt: types.MustNewTimeFromString("1978-04-14T07:58:36.326Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("curtus"),
                    Institution: unifiedgosdk.Pointer("Abbott - Batz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("1994-06-09T13:31:39.059Z"),
                },
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("blah cappelletti"),
                    EndAt: types.MustNewTimeFromString("1989-09-06T03:32:39.901Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("crustulum"),
                    Institution: unifiedgosdk.Pointer("Herman, Franey and Lebsack"),
                    Level: unifiedgosdk.Pointer("highschool"),
                    StartAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "ardith.beatty@outlook.com",
                    Name: unifiedgosdk.Pointer("Ardith Beatty"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Kutch, Lind and Marks"),
                    EndAt: types.MustNewTimeFromString("1974-03-01T17:02:58.824Z"),
                    StartAt: types.MustNewTimeFromString("1999-09-21T22:12:51.663Z"),
                    Title: unifiedgosdk.Pointer("Lead Security Technician"),
                },
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Trantow, Sipes and Witting"),
                    EndAt: types.MustNewTimeFromString("2004-10-14T14:01:00.717Z"),
                    StartAt: types.MustNewTimeFromString("1996-12-15T04:12:04.582Z"),
                    Title: unifiedgosdk.Pointer("Regional Applications Planner"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("e2d5ef47-da9e-4500-8493-4ce44bec7b6e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/27464445"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/ardith-beatty",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("cca4fa50-af15-422e-8c16-8ed5906d98de"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "officia",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginAgency.ToPointer(),
            Skills: []string{
                "trepide",
                "voveo",
            },
            Sources: []string{
                "vesper",
                "utroque",
            },
            Summary: unifiedgosdk.Pointer("Laborum trado carcer delibero patruus totam. Vita cohors cariosus trans denego barba rerum similique. Templum totam suus voluptatem cui depraedor virgo cui comminor commodi."),
            Tags: []string{
                "voro",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(837) 251-3896",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Division Producer"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T17:47:32.451Z"),
            WebURL: unifiedgosdk.Pointer("https://fragrant-husband.net/"),
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

## PatchAtsCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsCompany" method="patch" path="/ats/{connection_id}/company/{id}" example="ats_company" -->
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

    res, err := s.Ats.PatchAtsCompany(ctx, operations.PatchAtsCompanyRequest{
        AtsCompany: shared.AtsCompany{
            CreatedAt: types.MustNewTimeFromString("2019-04-22T03:50:02.920Z"),
            ID: unifiedgosdk.Pointer("b636b5e8-021c-497e-9250-478a87440e00"),
            Name: unifiedgosdk.Pointer("Gulgowski, Dibbert and Wilderman"),
            Phone: unifiedgosdk.Pointer("1-882-360-2010"),
            UpdatedAt: types.MustNewTimeFromString("2020-10-01T04:12:23.719Z"),
            WebsiteURL: unifiedgosdk.Pointer("https://www.gulgowskidibbert.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchAtsCompanyRequest](../../pkg/models/operations/patchatscompanyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchAtsCompanyResponse](../../pkg/models/operations/patchatscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsDocument" method="patch" path="/ats/{connection_id}/document/{id}" example="ats_document" -->
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

    res, err := s.Ats.PatchAtsDocument(ctx, operations.PatchAtsDocumentRequest{
        AtsDocument: shared.AtsDocument{
            CreatedAt: types.MustNewTimeFromString("2021-08-20T08:00:27.437Z"),
            DocumentURL: unifiedgosdk.Pointer("https://vengeful-lashes.biz"),
            Filename: unifiedgosdk.Pointer("lonzo_schowalter_resume.pdf"),
            ID: unifiedgosdk.Pointer("907ef7b6-235a-44f3-8ec7-d96d3a868fc2"),
            Metadata: []shared.AtsMetadata{},
            Type: shared.AtsDocumentTypeCoverLetter.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-12-07T04:59:55.146Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAtsDocumentRequest](../../pkg/models/operations/patchatsdocumentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAtsDocumentResponse](../../pkg/models/operations/patchatsdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsInterview

Update an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsInterview" method="patch" path="/ats/{connection_id}/interview/{id}" example="ats_interview" -->
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

    res, err := s.Ats.PatchAtsInterview(ctx, operations.PatchAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2021-12-08T06:33:19.583Z"),
            ID: unifiedgosdk.Pointer("e4f6fa22-60cf-4f9b-98c5-bf66341af9c5"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            MeetingURL: unifiedgosdk.Pointer("https://immaculate-foodstuffs.biz"),
            Metadata: []shared.AtsMetadata{},
            StartAt: types.MustNewTimeFromString("2021-12-08T06:03:19.583Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-09-23T16:39:11.723Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchAtsInterviewRequest](../../pkg/models/operations/patchatsinterviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchAtsInterviewResponse](../../pkg/models/operations/patchatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsJob

Update a job

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsJob" method="patch" path="/ats/{connection_id}/job/{id}" example="ats_job" -->
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

    res, err := s.Ats.PatchAtsJob(ctx, operations.PatchAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("98097 Carlo Trail"),
                    City: unifiedgosdk.Pointer("South Judd"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("89776-0669"),
                    Region: unifiedgosdk.Pointer("Mississippi"),
                    RegionCode: unifiedgosdk.Pointer("FL"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("AUD"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("MZN"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Global"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("0b07a34e-39f9-404a-af24-ab50ae45bbd3"),
            Industry: unifiedgosdk.Pointer("Handcrafted Bamboo Soap"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5ae4056e-ffe0-4bd5-bc78-8b4c4dae688f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "desolo",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Forward Brand Producer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Admoveo trado textilis."),
                    OpenedAt: types.MustNewTimeFromString("2026-06-07T21:28:07.934Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("75008 Medhurst Road"),
                        City: unifiedgosdk.Pointer("Madelynnshire"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("98754"),
                        Region: unifiedgosdk.Pointer("Wisconsin"),
                        RegionCode: unifiedgosdk.Pointer("CO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-08-02T00:29:26.254Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("0711f01a-d966-4692-856e-ebefe2c72287"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-08-28T07:52:25.774Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("8a96c284-ea18-4779-bda9-ab1be23a235d"),
                    MaxLength: unifiedgosdk.Pointer[float64](117.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "tactus",
                        "vehemens",
                    },
                    Prompt: unifiedgosdk.Pointer("Totam maxime tres tergiversatio."),
                    Question: "Doloribus mollitia.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("9ed04038-4f62-4955-ba05-9ccf7856a92c"),
                    MaxLength: unifiedgosdk.Pointer[float64](453.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "constans",
                    },
                    Prompt: unifiedgosdk.Pointer("Veritatis sodalitas nemo natus attonbitus reprehenderit voro depono."),
                    Question: "Nam.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    ID: unifiedgosdk.Pointer("6306a3f2-d536-4273-9170-bad83ea4b5ae"),
                    MaxLength: unifiedgosdk.Pointer[float64](66.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "repellendus",
                    },
                    Prompt: unifiedgosdk.Pointer("Tum tamisium autus."),
                    Question: "Confero trepide.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "suadeo",
                "utrimque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crepusculum adsuesco vere."),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T09:47:18.165Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchAtsJobRequest](../../pkg/models/operations/patchatsjobrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PatchAtsJobResponse](../../pkg/models/operations/patchatsjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsScorecard

Update a scorecard

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsScorecard" method="patch" path="/ats/{connection_id}/scorecard/{id}" example="ats_scorecard" -->
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

    res, err := s.Ats.PatchAtsScorecard(ctx, operations.PatchAtsScorecardRequest{
        AtsScorecard: shared.AtsScorecard{
            Comment: unifiedgosdk.Pointer("Maiores enim."),
            CreatedAt: types.MustNewTimeFromString("2022-02-20T17:09:45.498Z"),
            ID: unifiedgosdk.Pointer("b9b9b7a0-dae3-4133-a9d1-fa9a7b5843ed"),
            Metadata: []shared.AtsMetadata{},
            Questions: []shared.AtsScorecardQuestion{
                shared.AtsScorecardQuestion{
                    Description: unifiedgosdk.Pointer("Sulum textor eveniet facere vita."),
                    Text: "Aliquam.",
                },
                shared.AtsScorecardQuestion{
                    Answer: unifiedgosdk.Pointer("Decretum."),
                    Description: unifiedgosdk.Pointer("Conatus cicuta doloremque statua bonus."),
                    Text: "Pecto vulpes libero vomer comburo.",
                },
            },
            Recommendation: shared.RecommendationStrongYes.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-06-05T13:40:04.409Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchAtsScorecardRequest](../../pkg/models/operations/patchatsscorecardrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchAtsScorecardResponse](../../pkg/models/operations/patchatsscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsActivity

Remove an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsActivity" method="delete" path="/ats/{connection_id}/activity/{id}" -->
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

    res, err := s.Ats.RemoveAtsActivity(ctx, operations.RemoveAtsActivityRequest{
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
| `request`                                                                                      | [operations.RemoveAtsActivityRequest](../../pkg/models/operations/removeatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAtsActivityResponse](../../pkg/models/operations/removeatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsApplication

Remove an application

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsApplication" method="delete" path="/ats/{connection_id}/application/{id}" -->
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

    res, err := s.Ats.RemoveAtsApplication(ctx, operations.RemoveAtsApplicationRequest{
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
| `request`                                                                                            | [operations.RemoveAtsApplicationRequest](../../pkg/models/operations/removeatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveAtsApplicationResponse](../../pkg/models/operations/removeatsapplicationresponse.md), error**

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

    res, err := s.Ats.RemoveAtsCandidate(ctx, operations.RemoveAtsCandidateRequest{
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

## RemoveAtsCompany

Remove a company

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsCompany" method="delete" path="/ats/{connection_id}/company/{id}" -->
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

    res, err := s.Ats.RemoveAtsCompany(ctx, operations.RemoveAtsCompanyRequest{
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
| `request`                                                                                    | [operations.RemoveAtsCompanyRequest](../../pkg/models/operations/removeatscompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveAtsCompanyResponse](../../pkg/models/operations/removeatscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsDocument

Remove a document

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsDocument" method="delete" path="/ats/{connection_id}/document/{id}" -->
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

    res, err := s.Ats.RemoveAtsDocument(ctx, operations.RemoveAtsDocumentRequest{
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
| `request`                                                                                      | [operations.RemoveAtsDocumentRequest](../../pkg/models/operations/removeatsdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAtsDocumentResponse](../../pkg/models/operations/removeatsdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsInterview

Remove an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsInterview" method="delete" path="/ats/{connection_id}/interview/{id}" -->
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

    res, err := s.Ats.RemoveAtsInterview(ctx, operations.RemoveAtsInterviewRequest{
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
| `request`                                                                                        | [operations.RemoveAtsInterviewRequest](../../pkg/models/operations/removeatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveAtsInterviewResponse](../../pkg/models/operations/removeatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsJob

Remove a job

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsJob" method="delete" path="/ats/{connection_id}/job/{id}" -->
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

    res, err := s.Ats.RemoveAtsJob(ctx, operations.RemoveAtsJobRequest{
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RemoveAtsJobRequest](../../pkg/models/operations/removeatsjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.RemoveAtsJobResponse](../../pkg/models/operations/removeatsjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsScorecard

Remove a scorecard

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsScorecard" method="delete" path="/ats/{connection_id}/scorecard/{id}" -->
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

    res, err := s.Ats.RemoveAtsScorecard(ctx, operations.RemoveAtsScorecardRequest{
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
| `request`                                                                                        | [operations.RemoveAtsScorecardRequest](../../pkg/models/operations/removeatsscorecardrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveAtsScorecardResponse](../../pkg/models/operations/removeatsscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsActivity" method="put" path="/ats/{connection_id}/activity/{id}" example="ats_activity" -->
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

    res, err := s.Ats.UpdateAtsActivity(ctx, operations.UpdateAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Sapiente."),
            ID: unifiedgosdk.Pointer("1dd88965-536d-4641-a508-0b97b622bef5"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("26853e59-8a39-416d-88e8-757699c82c57"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("cimentarius"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Reprehenderit verto.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Venio blandior aureus."),
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-04T06:38:15.375Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAtsActivityRequest](../../pkg/models/operations/updateatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAtsActivityResponse](../../pkg/models/operations/updateatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsApplication

Update an application

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsApplication" method="put" path="/ats/{connection_id}/application/{id}" example="ats_application" -->
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

    res, err := s.Ats.UpdateAtsApplication(ctx, operations.UpdateAtsApplicationRequest{
        AtsApplication: shared.AtsApplication{
            Answers: []shared.AtsApplicationAnswer{},
            AppliedAt: types.MustNewTimeFromString("2025-09-29T18:41:44.750Z"),
            CreatedAt: types.MustNewTimeFromString("2023-10-17T07:19:48.787Z"),
            HiredAt: types.MustNewTimeFromString("2026-05-12T18:30:35.456Z"),
            ID: unifiedgosdk.Pointer("95d97fad-3b9a-4527-bb22-1953ed79ef05"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("382ffa38-ae9e-4b90-9c6f-f3dacf302b7c"),
                    Namespace: unifiedgosdk.Pointer("application"),
                    Slug: unifiedgosdk.Pointer("stella"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Caput.",
                    )),
                },
            },
            Offers: []shared.AtsOffer{},
            OriginalStatus: unifiedgosdk.Pointer("vomica"),
            OriginalSubstatus: unifiedgosdk.Pointer("allatus"),
            RejectedAt: types.MustNewTimeFromString("2026-10-11T13:07:05.733Z"),
            RejectedReason: unifiedgosdk.Pointer("Cometes amplitudo videlicet talio."),
            Source: unifiedgosdk.Pointer("credo"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            Summary: unifiedgosdk.Pointer("Arguo repudiandae defessus. Adflicto caecus claustrum annus addo vere crebro. Alias sponte turbo sopor vergo surculus nam."),
            UpdatedAt: types.MustNewTimeFromString("2026-10-18T09:21:22.057Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateAtsApplicationRequest](../../pkg/models/operations/updateatsapplicationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateAtsApplicationResponse](../../pkg/models/operations/updateatsapplicationresponse.md), error**

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

    res, err := s.Ats.UpdateAtsCandidate(ctx, operations.UpdateAtsCandidateRequest{
        AtsCandidate: shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.Pointer("6790 S 1st Street"),
                City: unifiedgosdk.Pointer("Ferrystad"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("91605"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("HI"),
            },
            CompanyName: unifiedgosdk.Pointer("Lehner and Sons"),
            CreatedAt: types.MustNewTimeFromString("2023-10-16T05:42:56.049Z"),
            DateOfBirth: types.MustNewTimeFromString("1994-07-05T19:48:57.631Z"),
            Education: []shared.AtsCandidateEducation{
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("emphasize tricky"),
                    EndAt: types.MustNewTimeFromString("1978-04-14T07:58:36.326Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("curtus"),
                    Institution: unifiedgosdk.Pointer("Abbott - Batz"),
                    Level: unifiedgosdk.Pointer("phd"),
                    StartAt: types.MustNewTimeFromString("1994-06-09T13:31:39.059Z"),
                },
                shared.AtsCandidateEducation{
                    Degree: unifiedgosdk.Pointer("blah cappelletti"),
                    EndAt: types.MustNewTimeFromString("1989-09-06T03:32:39.901Z"),
                    FieldOfStudy: unifiedgosdk.Pointer("crustulum"),
                    Institution: unifiedgosdk.Pointer("Herman, Franey and Lebsack"),
                    Level: unifiedgosdk.Pointer("highschool"),
                    StartAt: types.MustNewTimeFromString("1992-11-28T20:23:20.311Z"),
                },
            },
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "ardith.beatty@outlook.com",
                    Name: unifiedgosdk.Pointer("Ardith Beatty"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Experiences: []shared.AtsCandidateExperience{
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Kutch, Lind and Marks"),
                    EndAt: types.MustNewTimeFromString("1974-03-01T17:02:58.824Z"),
                    StartAt: types.MustNewTimeFromString("1999-09-21T22:12:51.663Z"),
                    Title: unifiedgosdk.Pointer("Lead Security Technician"),
                },
                shared.AtsCandidateExperience{
                    CompanyName: unifiedgosdk.Pointer("Trantow, Sipes and Witting"),
                    EndAt: types.MustNewTimeFromString("2004-10-14T14:01:00.717Z"),
                    StartAt: types.MustNewTimeFromString("1996-12-15T04:12:04.582Z"),
                    Title: unifiedgosdk.Pointer("Regional Applications Planner"),
                },
            },
            FirstName: unifiedgosdk.Pointer("Ardith"),
            ID: unifiedgosdk.Pointer("e2d5ef47-da9e-4500-8493-4ce44bec7b6e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/27464445"),
            JobIds: []string{},
            LastName: unifiedgosdk.Pointer("Beatty"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/ardith-beatty",
            },
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("cca4fa50-af15-422e-8c16-8ed5906d98de"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "officia",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Ardith Beatty"),
            Origin: shared.OriginAgency.ToPointer(),
            Skills: []string{
                "trepide",
                "voveo",
            },
            Sources: []string{
                "vesper",
                "utroque",
            },
            Summary: unifiedgosdk.Pointer("Laborum trado carcer delibero patruus totam. Vita cohors cariosus trans denego barba rerum similique. Templum totam suus voluptatem cui depraedor virgo cui comminor commodi."),
            Tags: []string{
                "voro",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "(837) 251-3896",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Customer Division Producer"),
            UpdatedAt: types.MustNewTimeFromString("2024-04-28T17:47:32.451Z"),
            WebURL: unifiedgosdk.Pointer("https://fragrant-husband.net/"),
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

## UpdateAtsCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsCompany" method="put" path="/ats/{connection_id}/company/{id}" example="ats_company" -->
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

    res, err := s.Ats.UpdateAtsCompany(ctx, operations.UpdateAtsCompanyRequest{
        AtsCompany: shared.AtsCompany{
            CreatedAt: types.MustNewTimeFromString("2019-04-22T03:50:02.920Z"),
            ID: unifiedgosdk.Pointer("b636b5e8-021c-497e-9250-478a87440e00"),
            Name: unifiedgosdk.Pointer("Gulgowski, Dibbert and Wilderman"),
            Phone: unifiedgosdk.Pointer("1-882-360-2010"),
            UpdatedAt: types.MustNewTimeFromString("2020-10-01T04:12:23.719Z"),
            WebsiteURL: unifiedgosdk.Pointer("https://www.gulgowskidibbert.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateAtsCompanyRequest](../../pkg/models/operations/updateatscompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateAtsCompanyResponse](../../pkg/models/operations/updateatscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsDocument" method="put" path="/ats/{connection_id}/document/{id}" example="ats_document" -->
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

    res, err := s.Ats.UpdateAtsDocument(ctx, operations.UpdateAtsDocumentRequest{
        AtsDocument: shared.AtsDocument{
            CreatedAt: types.MustNewTimeFromString("2021-08-20T08:00:27.437Z"),
            DocumentURL: unifiedgosdk.Pointer("https://vengeful-lashes.biz"),
            Filename: unifiedgosdk.Pointer("lonzo_schowalter_resume.pdf"),
            ID: unifiedgosdk.Pointer("907ef7b6-235a-44f3-8ec7-d96d3a868fc2"),
            Metadata: []shared.AtsMetadata{},
            Type: shared.AtsDocumentTypeCoverLetter.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-12-07T04:59:55.146Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAtsDocumentRequest](../../pkg/models/operations/updateatsdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAtsDocumentResponse](../../pkg/models/operations/updateatsdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsInterview

Update an interview

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsInterview" method="put" path="/ats/{connection_id}/interview/{id}" example="ats_interview" -->
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

    res, err := s.Ats.UpdateAtsInterview(ctx, operations.UpdateAtsInterviewRequest{
        AtsInterview: shared.AtsInterview{
            CreatedAt: types.MustNewTimeFromString("2021-11-28T03:14:47.774Z"),
            EndAt: types.MustNewTimeFromString("2021-12-08T06:33:19.583Z"),
            ID: unifiedgosdk.Pointer("e4f6fa22-60cf-4f9b-98c5-bf66341af9c5"),
            Location: unifiedgosdk.Pointer("26596 Halle Trafficway"),
            MeetingURL: unifiedgosdk.Pointer("https://immaculate-foodstuffs.biz"),
            Metadata: []shared.AtsMetadata{},
            StartAt: types.MustNewTimeFromString("2021-12-08T06:03:19.583Z"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-09-23T16:39:11.723Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAtsInterviewRequest](../../pkg/models/operations/updateatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAtsInterviewResponse](../../pkg/models/operations/updateatsinterviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsJob

Update a job

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsJob" method="put" path="/ats/{connection_id}/job/{id}" example="ats_job" -->
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

    res, err := s.Ats.UpdateAtsJob(ctx, operations.UpdateAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("98097 Carlo Trail"),
                    City: unifiedgosdk.Pointer("South Judd"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("89776-0669"),
                    Region: unifiedgosdk.Pointer("Mississippi"),
                    RegionCode: unifiedgosdk.Pointer("FL"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("AUD"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("MZN"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Global"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("0b07a34e-39f9-404a-af24-ab50ae45bbd3"),
            Industry: unifiedgosdk.Pointer("Handcrafted Bamboo Soap"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5ae4056e-ffe0-4bd5-bc78-8b4c4dae688f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "desolo",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Forward Brand Producer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Admoveo trado textilis."),
                    OpenedAt: types.MustNewTimeFromString("2026-06-07T21:28:07.934Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("75008 Medhurst Road"),
                        City: unifiedgosdk.Pointer("Madelynnshire"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("98754"),
                        Region: unifiedgosdk.Pointer("Wisconsin"),
                        RegionCode: unifiedgosdk.Pointer("CO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-08-02T00:29:26.254Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("0711f01a-d966-4692-856e-ebefe2c72287"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-08-28T07:52:25.774Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("8a96c284-ea18-4779-bda9-ab1be23a235d"),
                    MaxLength: unifiedgosdk.Pointer[float64](117.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "tactus",
                        "vehemens",
                    },
                    Prompt: unifiedgosdk.Pointer("Totam maxime tres tergiversatio."),
                    Question: "Doloribus mollitia.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("9ed04038-4f62-4955-ba05-9ccf7856a92c"),
                    MaxLength: unifiedgosdk.Pointer[float64](453.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "constans",
                    },
                    Prompt: unifiedgosdk.Pointer("Veritatis sodalitas nemo natus attonbitus reprehenderit voro depono."),
                    Question: "Nam.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    ID: unifiedgosdk.Pointer("6306a3f2-d536-4273-9170-bad83ea4b5ae"),
                    MaxLength: unifiedgosdk.Pointer[float64](66.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "repellendus",
                    },
                    Prompt: unifiedgosdk.Pointer("Tum tamisium autus."),
                    Question: "Confero trepide.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "suadeo",
                "utrimque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Crepusculum adsuesco vere."),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T09:47:18.165Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateAtsJobRequest](../../pkg/models/operations/updateatsjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateAtsJobResponse](../../pkg/models/operations/updateatsjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsScorecard

Update a scorecard

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsScorecard" method="put" path="/ats/{connection_id}/scorecard/{id}" example="ats_scorecard" -->
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

    res, err := s.Ats.UpdateAtsScorecard(ctx, operations.UpdateAtsScorecardRequest{
        AtsScorecard: shared.AtsScorecard{
            Comment: unifiedgosdk.Pointer("Maiores enim."),
            CreatedAt: types.MustNewTimeFromString("2022-02-20T17:09:45.498Z"),
            ID: unifiedgosdk.Pointer("b9b9b7a0-dae3-4133-a9d1-fa9a7b5843ed"),
            Metadata: []shared.AtsMetadata{},
            Questions: []shared.AtsScorecardQuestion{
                shared.AtsScorecardQuestion{
                    Description: unifiedgosdk.Pointer("Sulum textor eveniet facere vita."),
                    Text: "Aliquam.",
                },
                shared.AtsScorecardQuestion{
                    Answer: unifiedgosdk.Pointer("Decretum."),
                    Description: unifiedgosdk.Pointer("Conatus cicuta doloremque statua bonus."),
                    Text: "Pecto vulpes libero vomer comburo.",
                },
            },
            Recommendation: shared.RecommendationStrongYes.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-06-05T13:40:04.409Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAtsScorecardRequest](../../pkg/models/operations/updateatsscorecardrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAtsScorecardResponse](../../pkg/models/operations/updateatsscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |