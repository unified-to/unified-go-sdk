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
            ID: unifiedgosdk.Pointer("74f2d019-91ed-4d70-8267-f6ff4321a2ab"),
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
                    ID: unifiedgosdk.Pointer("8b135ac1-4e14-47b7-b066-5920dedd5447"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-07-01T20:54:37.792Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-27T05:06:19.670Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("b3d280da-2a6b-40bc-b8da-ee7db2de7fd3"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-23T02:24:39.385Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("70f509c1-9f2a-4448-9b86-0de13035cc18"),
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
                    ID: unifiedgosdk.Pointer("f5bf7ad7-659f-4583-8be7-b54b3eac2173"),
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
                    ID: unifiedgosdk.Pointer("e18857d9-ad26-4e59-b0d5-8c276d4ccc6a"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-21T04:03:05.014Z"),
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
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("ab32e9f6-f034-48ba-a454-e0b1b6b3e77f"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T09:27:00.625Z"),
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
            ID: unifiedgosdk.Pointer("5d4ba0cd-220d-4535-a040-25d7818312f7"),
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
                    ID: unifiedgosdk.Pointer("b07f77d4-76a5-41c0-9f14-f7279a2ca8b9"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-07-01T20:54:37.837Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-27T05:06:19.717Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("2b069de3-6bdf-4a01-b984-a2133fd83a36"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-23T02:24:39.434Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("adb6d1e3-b205-492a-bb7a-1ca8dcfe4ca5"),
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
                    ID: unifiedgosdk.Pointer("8c2e776f-fdb4-43e8-b35f-c83f5464c9e7"),
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
                    ID: unifiedgosdk.Pointer("5fcb686f-0921-4aa3-8dc0-c15152cc2578"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-21T04:03:05.056Z"),
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
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("2e14cebb-d024-4bba-b77e-89d7ad13a6b4"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T09:27:00.667Z"),
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
            ID: unifiedgosdk.Pointer("5d4ba0cd-220d-4535-a040-25d7818312f7"),
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
                    ID: unifiedgosdk.Pointer("b07f77d4-76a5-41c0-9f14-f7279a2ca8b9"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-07-01T20:54:37.837Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-27T05:06:19.717Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("2b069de3-6bdf-4a01-b984-a2133fd83a36"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-23T02:24:39.434Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("adb6d1e3-b205-492a-bb7a-1ca8dcfe4ca5"),
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
                    ID: unifiedgosdk.Pointer("8c2e776f-fdb4-43e8-b35f-c83f5464c9e7"),
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
                    ID: unifiedgosdk.Pointer("5fcb686f-0921-4aa3-8dc0-c15152cc2578"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-21T04:03:05.056Z"),
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
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("2e14cebb-d024-4bba-b77e-89d7ad13a6b4"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-09T09:27:00.667Z"),
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