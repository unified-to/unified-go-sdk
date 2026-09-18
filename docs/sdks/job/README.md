# Job

## Overview

### Available Operations

* [CreateAtsJob](#createatsjob) - Create a job
* [GetAtsJob](#getatsjob) - Retrieve a job
* [ListAtsJobs](#listatsjobs) - List all jobs
* [PatchAtsJob](#patchatsjob) - Update a job
* [RemoveAtsJob](#removeatsjob) - Remove a job
* [UpdateAtsJob](#updateatsjob) - Update a job

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
            HiringManagers: []shared.AtsReference{
                shared.AtsReference{
                    ID: unifiedgosdk.Pointer("fd9852e3-9035-4f42-beb3-bbf4e4022122"),
                    Name: unifiedgosdk.Pointer("Eloise Mueller PhD"),
                },
            },
            ID: unifiedgosdk.Pointer("046560a0-e320-4a61-854e-dd156bb78313"),
            Industry: unifiedgosdk.Pointer("Gorgeous Plastic Computer"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("fa2b29a7-1cd0-4c23-a530-c6c51fc9e075"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "acceptus",
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
                    OpenedAt: types.MustNewTimeFromString("2026-05-09T13:49:23.819Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("8460 Nils Trace"),
                        City: unifiedgosdk.Pointer("West Mervinburgh"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("14162"),
                        Region: unifiedgosdk.Pointer("Maine"),
                        RegionCode: unifiedgosdk.Pointer("MO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-07-02T05:10:20.142Z"),
                    Description: unifiedgosdk.Pointer("Deduco cultellus alii terebro depono thesaurus."),
                    ID: unifiedgosdk.Pointer("f6101769-deb3-4721-978c-d205638870ee"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("6788 Oxford Road"),
                    Name: unifiedgosdk.Pointer("Forward Security Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://ajar-metabolite.net/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-07-27T19:31:53.073Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Trepide provident taceo rem."),
                    ID: unifiedgosdk.Pointer("289f27c0-311c-41e5-ad9d-cbe2097332c2"),
                    Options: []string{
                        "censura",
                        "tum",
                    },
                    Prompt: unifiedgosdk.Pointer("Spectaculum mollitia arcus compello."),
                    Question: "Sodalitas nemo natus attonbitus reprehenderit voro depono constans vehemens ante.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("b3a0b53b-38f3-4e8d-84b9-f413a900d79b"),
                    Options: []string{
                        "odit",
                    },
                    Prompt: unifiedgosdk.Pointer("Similique absque temeritas celebrer enim."),
                    Question: "Vinitor sodalitas desino sollers viduo volo.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Abstergo possimus quibusdam deinde amoveo."),
                    ID: unifiedgosdk.Pointer("568be61d-060e-4d8c-a8ab-8a17cb25edf3"),
                    Options: []string{
                        "vallum",
                    },
                    Prompt: unifiedgosdk.Pointer("Ara thermae aetas vivo constans victoria volo carbo vehemens praesentium."),
                    Question: "Subiungo ambitus neque talis amitto terreo alienus quae vulticulus.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "amiculum",
                "crux",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-31T19:10:09.085Z"),
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
            HiringManagers: []shared.AtsReference{
                shared.AtsReference{
                    ID: unifiedgosdk.Pointer("fd9852e3-9035-4f42-beb3-bbf4e4022122"),
                    Name: unifiedgosdk.Pointer("Eloise Mueller PhD"),
                },
            },
            ID: unifiedgosdk.Pointer("11f2ef01-3e7f-48d1-9af6-0bf5864b5009"),
            Industry: unifiedgosdk.Pointer("Gorgeous Plastic Computer"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("3bb20593-8ef5-4c8d-b429-67e9261bbe23"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "acceptus",
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
                    OpenedAt: types.MustNewTimeFromString("2026-05-09T13:49:23.852Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("8460 Nils Trace"),
                        City: unifiedgosdk.Pointer("West Mervinburgh"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("14162"),
                        Region: unifiedgosdk.Pointer("Maine"),
                        RegionCode: unifiedgosdk.Pointer("MO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-07-02T05:10:20.178Z"),
                    Description: unifiedgosdk.Pointer("Deduco cultellus alii terebro depono thesaurus."),
                    ID: unifiedgosdk.Pointer("f6101769-deb3-4721-978c-d205638870ee"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("6788 Oxford Road"),
                    Name: unifiedgosdk.Pointer("Forward Security Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://ajar-metabolite.net/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-07-27T19:31:53.109Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Trepide provident taceo rem."),
                    ID: unifiedgosdk.Pointer("289f27c0-311c-41e5-ad9d-cbe2097332c2"),
                    Options: []string{
                        "censura",
                        "tum",
                    },
                    Prompt: unifiedgosdk.Pointer("Spectaculum mollitia arcus compello."),
                    Question: "Sodalitas nemo natus attonbitus reprehenderit voro depono constans vehemens ante.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("b3a0b53b-38f3-4e8d-84b9-f413a900d79b"),
                    Options: []string{
                        "odit",
                    },
                    Prompt: unifiedgosdk.Pointer("Similique absque temeritas celebrer enim."),
                    Question: "Vinitor sodalitas desino sollers viduo volo.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Abstergo possimus quibusdam deinde amoveo."),
                    ID: unifiedgosdk.Pointer("568be61d-060e-4d8c-a8ab-8a17cb25edf3"),
                    Options: []string{
                        "vallum",
                    },
                    Prompt: unifiedgosdk.Pointer("Ara thermae aetas vivo constans victoria volo carbo vehemens praesentium."),
                    Question: "Subiungo ambitus neque talis amitto terreo alienus quae vulticulus.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "amiculum",
                "crux",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-31T19:10:09.115Z"),
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
            HiringManagers: []shared.AtsReference{
                shared.AtsReference{
                    ID: unifiedgosdk.Pointer("fd9852e3-9035-4f42-beb3-bbf4e4022122"),
                    Name: unifiedgosdk.Pointer("Eloise Mueller PhD"),
                },
            },
            ID: unifiedgosdk.Pointer("11f2ef01-3e7f-48d1-9af6-0bf5864b5009"),
            Industry: unifiedgosdk.Pointer("Gorgeous Plastic Computer"),
            LanguageLocale: unifiedgosdk.Pointer("en"),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("3bb20593-8ef5-4c8d-b429-67e9261bbe23"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "acceptus",
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
                    OpenedAt: types.MustNewTimeFromString("2026-05-09T13:49:23.852Z"),
                    Status: shared.AtsJobOpeningStatusOpen.ToPointer(),
                },
            },
            Postings: []shared.AtsJobPosting{
                shared.AtsJobPosting{
                    Address: &shared.PropertyAtsJobPostingAddress{
                        Address1: unifiedgosdk.Pointer("8460 Nils Trace"),
                        City: unifiedgosdk.Pointer("West Mervinburgh"),
                        CountryCode: unifiedgosdk.Pointer("US"),
                        PostalCode: unifiedgosdk.Pointer("14162"),
                        Region: unifiedgosdk.Pointer("Maine"),
                        RegionCode: unifiedgosdk.Pointer("MO"),
                    },
                    CreatedAt: types.MustNewTimeFromString("2026-07-02T05:10:20.178Z"),
                    Description: unifiedgosdk.Pointer("Deduco cultellus alii terebro depono thesaurus."),
                    ID: unifiedgosdk.Pointer("f6101769-deb3-4721-978c-d205638870ee"),
                    IsActive: unifiedgosdk.Pointer(false),
                    Location: unifiedgosdk.Pointer("6788 Oxford Road"),
                    Name: unifiedgosdk.Pointer("Forward Security Orchestrator"),
                    PostingURL: unifiedgosdk.Pointer("https://ajar-metabolite.net/"),
                    UpdatedAt: types.MustNewTimeFromString("2026-07-27T19:31:53.109Z"),
                },
            },
            PublicJobUrls: []string{
                "https://trustworthy-elver.info",
                "https://parched-dash.info",
            },
            Questions: []shared.AtsJobQuestion{
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Trepide provident taceo rem."),
                    ID: unifiedgosdk.Pointer("289f27c0-311c-41e5-ad9d-cbe2097332c2"),
                    Options: []string{
                        "censura",
                        "tum",
                    },
                    Prompt: unifiedgosdk.Pointer("Spectaculum mollitia arcus compello."),
                    Question: "Sodalitas nemo natus attonbitus reprehenderit voro depono constans vehemens ante.",
                    Required: unifiedgosdk.Pointer(true),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    ID: unifiedgosdk.Pointer("b3a0b53b-38f3-4e8d-84b9-f413a900d79b"),
                    Options: []string{
                        "odit",
                    },
                    Prompt: unifiedgosdk.Pointer("Similique absque temeritas celebrer enim."),
                    Question: "Vinitor sodalitas desino sollers viduo volo.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
                shared.AtsJobQuestion{
                    Description: unifiedgosdk.Pointer("Abstergo possimus quibusdam deinde amoveo."),
                    ID: unifiedgosdk.Pointer("568be61d-060e-4d8c-a8ab-8a17cb25edf3"),
                    Options: []string{
                        "vallum",
                    },
                    Prompt: unifiedgosdk.Pointer("Ara thermae aetas vivo constans victoria volo carbo vehemens praesentium."),
                    Question: "Subiungo ambitus neque talis amitto terreo alienus quae vulticulus.",
                    Required: unifiedgosdk.Pointer(false),
                    Type: shared.AtsJobQuestionTypeText,
                },
            },
            Skills: []string{
                "amiculum",
                "crux",
            },
            Status: shared.AtsJobStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-01-31T19:10:09.115Z"),
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