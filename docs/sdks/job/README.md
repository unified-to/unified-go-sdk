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
            ID: unifiedgosdk.Pointer("0582e242-8bc9-455a-ab81-5ade3ac6e01b"),
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
                    ID: unifiedgosdk.Pointer("ff28c930-4a34-474e-a550-86e22d4953f8"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-06-28T08:25:02.867Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-23T12:19:43.965Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("b3351d54-c3f0-4153-9131-e42d97366f93"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-19T07:35:24.450Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("0228ea40-3e3d-4d20-b1f5-b4fec98dbf4b"),
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
                    ID: unifiedgosdk.Pointer("689d2876-f31e-4da9-9f65-eba1b2b7b8a6"),
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
                    ID: unifiedgosdk.Pointer("dcd0f84d-0c17-4b1e-9cef-61338178704d"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T23:21:59.980Z"),
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
            ID: unifiedgosdk.Pointer("da37ffab-9383-433d-8e8c-c6f821cba56c"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T21:36:10.260Z"),
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
            ID: unifiedgosdk.Pointer("625c4b28-e51e-446e-a50d-b644b9643118"),
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
                    ID: unifiedgosdk.Pointer("d3ad4988-4dad-43c4-b624-12920e95f1ea"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-06-28T08:25:02.905Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-23T12:19:44.003Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("5c1cbb30-6887-4fe1-bb4a-891b616c57ef"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-19T07:35:24.489Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("1bafabf0-9cff-45d5-bb76-12db74124041"),
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
                    ID: unifiedgosdk.Pointer("facabaf1-238a-424a-83f0-7fcc05e77d48"),
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
                    ID: unifiedgosdk.Pointer("45fdfe48-e2e3-4a69-8f13-cc19c5fac4a0"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T23:22:00.014Z"),
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
            ID: unifiedgosdk.Pointer("b21ddaad-ae98-4445-972d-5b57e3a54e14"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T21:36:10.297Z"),
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
            ID: unifiedgosdk.Pointer("625c4b28-e51e-446e-a50d-b644b9643118"),
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
                    ID: unifiedgosdk.Pointer("d3ad4988-4dad-43c4-b624-12920e95f1ea"),
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
                    OpenedAt: types.MustNewTimeFromString("2026-06-28T08:25:02.905Z"),
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
                    CreatedAt: types.MustNewTimeFromString("2026-08-23T12:19:44.003Z"),
                    Description: unifiedgosdk.Pointer("Beatus eum volutabrum balbus victus ipsa tubineus claustrum."),
                    ID: unifiedgosdk.Pointer("5c1cbb30-6887-4fe1-bb4a-891b616c57ef"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("1725 Hessel Keys"),
                    Name: unifiedgosdk.Pointer("International Identity Manager"),
                    PostingURL: unifiedgosdk.Pointer("https://obedient-yak.com"),
                    UpdatedAt: types.MustNewTimeFromString("2026-09-19T07:35:24.489Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Auctus deporto tolero administratio capitulus animi agnosco tendo aetas vado."),
                    ID: unifiedgosdk.Pointer("1bafabf0-9cff-45d5-bb76-12db74124041"),
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
                    ID: unifiedgosdk.Pointer("facabaf1-238a-424a-83f0-7fcc05e77d48"),
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
                    ID: unifiedgosdk.Pointer("45fdfe48-e2e3-4a69-8f13-cc19c5fac4a0"),
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
            UpdatedAt: types.MustNewTimeFromString("2026-03-17T23:22:00.014Z"),
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
            ID: unifiedgosdk.Pointer("b21ddaad-ae98-4445-972d-5b57e3a54e14"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-07T21:36:10.297Z"),
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