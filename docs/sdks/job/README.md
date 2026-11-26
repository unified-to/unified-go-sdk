# Job

## Overview

### Available Operations

* [CreateAtsJob](#createatsjob) - Create a job
* [CreateHrisJob](#createhrisjob) - Create a job
* [GetAtsJob](#getatsjob) - Retrieve a job
* [GetHrisJob](#gethrisjob) - Retrieve a job
* [ListAtsJobs](#listatsjobs) - List all jobs
* [ListHrisJobs](#listhrisjobs) - List all jobs
* [PatchAtsJob](#patchatsjob) - Update a job
* [PatchHrisJob](#patchhrisjob) - Update a job
* [RemoveAtsJob](#removeatsjob) - Remove a job
* [RemoveHrisJob](#removehrisjob) - Remove a job
* [UpdateAtsJob](#updateatsjob) - Update a job
* [UpdateHrisJob](#updatehrisjob) - Update a job

## CreateAtsJob

Creates a AtsJob. Provide the unified AtsJob fields in the request body; they are mapped to the provider's API. Returns the created AtsJob.

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

    res, err := s.Job.CreateAtsJob(ctx, operations.CreateAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("458 E 8th Street"),
                    City: unifiedgosdk.Pointer("Visalia"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("60669-0678"),
                    Region: unifiedgosdk.Pointer("Oregon"),
                    RegionCode: unifiedgosdk.Pointer("DE"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("KMF"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Senior"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            Fees: []shared.AtsFee{
                shared.AtsFee{
                    Amount: unifiedgosdk.Pointer[float64](35146.0),
                    Currency: unifiedgosdk.Pointer("AFN"),
                    Type: shared.AtsFeeTypeContingency.ToPointer(),
                },
            },
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("51b2e44b-bc3d-41b6-8a2b-801dc779bba1"),
            Industry: unifiedgosdk.Pointer("Refined Ceramic Keyboard"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("21504bf0-8901-4c2f-b580-c140dd7ed16c"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cursus",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Senior Interactions Developer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Tolero aequus tamquam inflammatio asporto canis vinitor utrum."),
                    OpenedAt: types.MustNewTimeFromString("2026-07-09T20:52:00.154Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("295 First Street"),
                        City: unifiedgosdk.Pointer("East Agustinton"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("57059"),
                        Region: unifiedgosdk.Pointer("Utah"),
                        RegionCode: unifiedgosdk.Pointer("NE"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-09-04T14:47:35.863Z"),
                    Description: unifiedgosdk.Pointer("Sublime."),
                    ID: unifiedgosdk.Pointer("baf58a0f-9733-491d-a1c0-466e05848254"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("5699 Purdy Motorway"),
                    Name: unifiedgosdk.Pointer("Legacy Implementation Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://elementary-privilege.com/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-10-01T16:44:34.707Z"),
                },
            },
            PublicJobUrls: []string{
                "https://parched-dash.info",
                "https://mild-midwife.org/",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Capitulus."),
                    ID: unifiedgosdk.Pointer("f9c87bce-2f83-43a5-97a2-990f49711e96"),
                    MaxLength: unifiedgosdk.Pointer[float64](214.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "aetas",
                        "vado",
                    },
                    Prompt: unifiedgosdk.Pointer("Agnosco tendo."),
                    Question: "Deporto tolero.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("05f3c542-54e8-4be5-b1f4-7a355cb58925"),
                    MaxLength: unifiedgosdk.Pointer[float64](309.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "sodalitas",
                    },
                    Prompt: unifiedgosdk.Pointer("Supra veritatis."),
                    Question: "Maxime tres tergiversatio tactus vehemens architecto ab nam defetiscor.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Confero trepide."),
                    ID: unifiedgosdk.Pointer("98fc3f1a-6987-4708-8242-819b85317545"),
                    MaxLength: unifiedgosdk.Pointer[float64](428.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "censura",
                    },
                    Prompt: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    Question: "Attonbitus reprehenderit voro depono constans vehemens.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "similique",
                "absque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Contra damno desolo suadeo utrimque cresco aetas turpis."),
            UpdatedAt: types.MustNewTimeFromString("2026-03-28T10:16:05.688Z"),
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

## CreateHrisJob

Creates a HrisJob. Provide the unified HrisJob fields in the request body; they are mapped to the provider's API. Returns the created HrisJob.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisJob" method="post" path="/hris/{connection_id}/job" example="hris_job" -->
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

    res, err := s.Job.CreateHrisJob(ctx, operations.CreateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-54753"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Bonus vobis creator tabella callide cavus triumphus caveo abbas patria."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("48f68109-bccc-45e3-ae33-c629ba7c189d"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Customer Security Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-12T18:53:24.575Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateHrisJobRequest](../../pkg/models/operations/createhrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateHrisJobResponse](../../pkg/models/operations/createhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsJob

Retrieves a single AtsJob by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

    res, err := s.Job.GetAtsJob(ctx, operations.GetAtsJobRequest{
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

## GetHrisJob

Retrieves a single HrisJob by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisJob" method="get" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Job.GetHrisJob(ctx, operations.GetHrisJobRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetHrisJobRequest](../../pkg/models/operations/gethrisjobrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetHrisJobResponse](../../pkg/models/operations/gethrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsJobs

Returns a list of AtsJobs. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Job.ListAtsJobs(ctx, operations.ListAtsJobsRequest{
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

## ListHrisJobs

Returns a list of HrisJobs. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisJobs" method="get" path="/hris/{connection_id}/job" -->
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

    res, err := s.Job.ListHrisJobs(ctx, operations.ListHrisJobsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListHrisJobsRequest](../../pkg/models/operations/listhrisjobsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListHrisJobsResponse](../../pkg/models/operations/listhrisjobsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsJob

Updates the AtsJob identified by `id`. Provide the unified AtsJob fields to change in the request body.

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

    res, err := s.Job.PatchAtsJob(ctx, operations.PatchAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("458 E 8th Street"),
                    City: unifiedgosdk.Pointer("Visalia"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("60669-0678"),
                    Region: unifiedgosdk.Pointer("Oregon"),
                    RegionCode: unifiedgosdk.Pointer("DE"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("KMF"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Senior"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            Fees: []shared.AtsFee{
                shared.AtsFee{
                    Amount: unifiedgosdk.Pointer[float64](35146.0),
                    Currency: unifiedgosdk.Pointer("AFN"),
                    Type: shared.AtsFeeTypeContingency.ToPointer(),
                },
            },
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("adeb58d5-a0fb-481b-855c-6616f087ccd9"),
            Industry: unifiedgosdk.Pointer("Refined Ceramic Keyboard"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("177d39f4-0857-4739-b21a-5566f26c685a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cursus",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Senior Interactions Developer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Tolero aequus tamquam inflammatio asporto canis vinitor utrum."),
                    OpenedAt: types.MustNewTimeFromString("2026-07-09T20:52:00.189Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("295 First Street"),
                        City: unifiedgosdk.Pointer("East Agustinton"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("57059"),
                        Region: unifiedgosdk.Pointer("Utah"),
                        RegionCode: unifiedgosdk.Pointer("NE"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-09-04T14:47:35.898Z"),
                    Description: unifiedgosdk.Pointer("Sublime."),
                    ID: unifiedgosdk.Pointer("26e7ff5b-c0f1-4abe-976f-49e77f5a67b4"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("5699 Purdy Motorway"),
                    Name: unifiedgosdk.Pointer("Legacy Implementation Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://elementary-privilege.com/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-10-01T16:44:34.744Z"),
                },
            },
            PublicJobUrls: []string{
                "https://parched-dash.info",
                "https://mild-midwife.org/",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Capitulus."),
                    ID: unifiedgosdk.Pointer("b6e92350-665c-4108-a972-8397a78c870a"),
                    MaxLength: unifiedgosdk.Pointer[float64](214.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "aetas",
                        "vado",
                    },
                    Prompt: unifiedgosdk.Pointer("Agnosco tendo."),
                    Question: "Deporto tolero.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("e08216a8-df46-4801-84df-9ad4db1091d6"),
                    MaxLength: unifiedgosdk.Pointer[float64](309.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "sodalitas",
                    },
                    Prompt: unifiedgosdk.Pointer("Supra veritatis."),
                    Question: "Maxime tres tergiversatio tactus vehemens architecto ab nam defetiscor.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Confero trepide."),
                    ID: unifiedgosdk.Pointer("d9bc518d-ba03-4606-ba5c-f1b33781a5a4"),
                    MaxLength: unifiedgosdk.Pointer[float64](428.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "censura",
                    },
                    Prompt: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    Question: "Attonbitus reprehenderit voro depono constans vehemens.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "similique",
                "absque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Contra damno desolo suadeo utrimque cresco aetas turpis."),
            UpdatedAt: types.MustNewTimeFromString("2026-03-28T10:16:05.719Z"),
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

## PatchHrisJob

Updates the HrisJob identified by `id`. Provide the unified HrisJob fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisJob" method="patch" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
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

    res, err := s.Job.PatchHrisJob(ctx, operations.PatchHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-54753"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Bonus vobis creator tabella callide cavus triumphus caveo abbas patria."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("7ca26670-3428-4d9f-8826-2ff233ba5ffb"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Customer Security Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-12T18:53:24.609Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchHrisJobRequest](../../pkg/models/operations/patchhrisjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchHrisJobResponse](../../pkg/models/operations/patchhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsJob

Deletes the AtsJob identified by `id`.

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

    res, err := s.Job.RemoveAtsJob(ctx, operations.RemoveAtsJobRequest{
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

## RemoveHrisJob

Deletes the HrisJob identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisJob" method="delete" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Job.RemoveHrisJob(ctx, operations.RemoveHrisJobRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveHrisJobRequest](../../pkg/models/operations/removehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveHrisJobResponse](../../pkg/models/operations/removehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsJob

Updates the AtsJob identified by `id`. Provide the unified AtsJob fields to change in the request body.

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

    res, err := s.Job.UpdateAtsJob(ctx, operations.UpdateAtsJobRequest{
        AtsJob: shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.Pointer("458 E 8th Street"),
                    City: unifiedgosdk.Pointer("Visalia"),
                    CountryCode: unifiedgosdk.Pointer("US"),
                    PostalCode: unifiedgosdk.Pointer("60669-0678"),
                    Region: unifiedgosdk.Pointer("Oregon"),
                    RegionCode: unifiedgosdk.Pointer("DE"),
                },
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.FrequencyDay.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](174303.0),
                    Min: unifiedgosdk.Pointer[float64](174042.0),
                    Type: shared.AtsCompensationTypeBonus.ToPointer(),
                },
                shared.AtsCompensation{
                    Currency: unifiedgosdk.Pointer("KMF"),
                    Frequency: shared.FrequencyMonth.ToPointer(),
                    Max: unifiedgosdk.Pointer[float64](171171.0),
                    Min: unifiedgosdk.Pointer[float64](151975.0),
                    Type: shared.AtsCompensationTypeSalary.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2023-06-16T12:51:44.518Z"),
            Description: unifiedgosdk.Pointer("Senior"),
            EmploymentType: shared.EmploymentTypeFreelance.ToPointer(),
            Fees: []shared.AtsFee{
                shared.AtsFee{
                    Amount: unifiedgosdk.Pointer[float64](35146.0),
                    Currency: unifiedgosdk.Pointer("AFN"),
                    Type: shared.AtsFeeTypeContingency.ToPointer(),
                },
            },
            HiringManagers: []shared.AtsReference{},
            ID: unifiedgosdk.Pointer("adeb58d5-a0fb-481b-855c-6616f087ccd9"),
            Industry: unifiedgosdk.Pointer("Refined Ceramic Keyboard"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("177d39f4-0857-4739-b21a-5566f26c685a"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "cursus",
                    )),
                },
            },
            MinimumDegree: unifiedgosdk.Pointer("Bachelor"),
            MinimumExperienceYears: unifiedgosdk.Pointer[float64](3.0),
            Name: unifiedgosdk.Pointer("Senior Interactions Developer"),
            NumberOfOpenings: unifiedgosdk.Pointer[float64](1.0),
            Openings: []shared.AtsJobOpening{
                shared.AtsJobOpening{
                    CloseReason: unifiedgosdk.Pointer("Tolero aequus tamquam inflammatio asporto canis vinitor utrum."),
                    OpenedAt: types.MustNewTimeFromString("2026-07-09T20:52:00.189Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("295 First Street"),
                        City: unifiedgosdk.Pointer("East Agustinton"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("57059"),
                        Region: unifiedgosdk.Pointer("Utah"),
                        RegionCode: unifiedgosdk.Pointer("NE"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-09-04T14:47:35.898Z"),
                    Description: unifiedgosdk.Pointer("Sublime."),
                    ID: unifiedgosdk.Pointer("26e7ff5b-c0f1-4abe-976f-49e77f5a67b4"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("5699 Purdy Motorway"),
                    Name: unifiedgosdk.Pointer("Legacy Implementation Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://elementary-privilege.com/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-10-01T16:44:34.744Z"),
                },
            },
            PublicJobUrls: []string{
                "https://parched-dash.info",
                "https://mild-midwife.org/",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Capitulus."),
                    ID: unifiedgosdk.Pointer("b6e92350-665c-4108-a972-8397a78c870a"),
                    MaxLength: unifiedgosdk.Pointer[float64](214.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "aetas",
                        "vado",
                    },
                    Prompt: unifiedgosdk.Pointer("Agnosco tendo."),
                    Question: "Deporto tolero.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("e08216a8-df46-4801-84df-9ad4db1091d6"),
                    MaxLength: unifiedgosdk.Pointer[float64](309.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "sodalitas",
                    },
                    Prompt: unifiedgosdk.Pointer("Supra veritatis."),
                    Question: "Maxime tres tergiversatio tactus vehemens architecto ab nam defetiscor.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Confero trepide."),
                    ID: unifiedgosdk.Pointer("d9bc518d-ba03-4606-ba5c-f1b33781a5a4"),
                    MaxLength: unifiedgosdk.Pointer[float64](428.0),
                    MinLength: unifiedgosdk.Pointer[float64](1.0),
                    Options: []string{
                        "censura",
                    },
                    Prompt: unifiedgosdk.Pointer("Taceo rem compello spectaculum mollitia arcus compello."),
                    Question: "Attonbitus reprehenderit voro depono constans vehemens.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "similique",
                "absque",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            Summary: unifiedgosdk.Pointer("Contra damno desolo suadeo utrimque cresco aetas turpis."),
            UpdatedAt: types.MustNewTimeFromString("2026-03-28T10:16:05.719Z"),
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

## UpdateHrisJob

Updates the HrisJob identified by `id`. Provide the unified HrisJob fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisJob" method="put" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
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

    res, err := s.Job.UpdateHrisJob(ctx, operations.UpdateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-54753"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Bonus vobis creator tabella callide cavus triumphus caveo abbas patria."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("7ca26670-3428-4d9f-8826-2ff233ba5ffb"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Customer Security Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-12T18:53:24.609Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateHrisJobRequest](../../pkg/models/operations/updatehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateHrisJobResponse](../../pkg/models/operations/updatehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |