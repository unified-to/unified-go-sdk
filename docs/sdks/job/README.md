# Job
(*Job*)

### Available Operations

* [DeleteAtsConnectionIDJobID](#deleteatsconnectionidjobid) - Remove a job
* [GetAtsConnectionIDJob](#getatsconnectionidjob) - List all jobs
* [GetAtsConnectionIDJobID](#getatsconnectionidjobid) - Retrieve a job
* [PatchAtsConnectionIDJobID](#patchatsconnectionidjobid) - Update a job
* [PostAtsConnectionIDJob](#postatsconnectionidjob) - Create a job
* [PutAtsConnectionIDJobID](#putatsconnectionidjobid) - Update a job

## DeleteAtsConnectionIDJobID

Remove a job

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.DeleteAtsConnectionIDJobID(ctx, operations.DeleteAtsConnectionIDJobIDRequest{
        ConnectionID: "Sedan Bedfordshire Hybrid",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteAtsConnectionIDJobIDRequest](../../models/operations/deleteatsconnectionidjobidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.DeleteAtsConnectionIDJobIDResponse](../../models/operations/deleteatsconnectionidjobidresponse.md), error**


## GetAtsConnectionIDJob

List all jobs

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.GetAtsConnectionIDJob(ctx, operations.GetAtsConnectionIDJobRequest{
        ConnectionID: "City katal",
        Limit: unifiedgosdk.Float64(3542.62),
        Offset: unifiedgosdk.Float64(5417.97),
        Order: unifiedgosdk.String("publisher"),
        Query: unifiedgosdk.String("Folding"),
        Sort: unifiedgosdk.String("Kip gross recontextualize"),
        UpdatedGte: types.MustTimeFromString("2022-10-12T03:36:20.050Z"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAtsConnectionIDJobRequest](../../models/operations/getatsconnectionidjobrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetAtsConnectionIDJobResponse](../../models/operations/getatsconnectionidjobresponse.md), error**


## GetAtsConnectionIDJobID

Retrieve a job

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.GetAtsConnectionIDJobID(ctx, operations.GetAtsConnectionIDJobIDRequest{
        ConnectionID: "Jazz",
        ID: "<ID>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetAtsConnectionIDJobIDRequest](../../models/operations/getatsconnectionidjobidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetAtsConnectionIDJobIDResponse](../../models/operations/getatsconnectionidjobidresponse.md), error**


## PatchAtsConnectionIDJobID

Update a job

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.PatchAtsConnectionIDJobID(ctx, operations.PatchAtsConnectionIDJobIDRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.String("Transexual"),
                    Address2: unifiedgosdk.String("Planner redundant Towels"),
                    City: unifiedgosdk.String("Starkboro"),
                    Country: unifiedgosdk.String("Chad"),
                    CountryCode: unifiedgosdk.String("NU"),
                    PostalCode: unifiedgosdk.String("22603"),
                    Region: unifiedgosdk.String("Cambridgeshire"),
                    RegionCode: unifiedgosdk.String("Account Copernicium at"),
                },
            },
            ClosedAt: types.MustTimeFromString("2023-01-02T09:14:26.844Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("Metical"),
                    Frequency: shared.AtsCompensationFrequencyHour.ToPointer(),
                    Max: unifiedgosdk.Float64(1424.24),
                    Min: unifiedgosdk.Float64(3626.17),
                    Type: shared.AtsCompensationTypeSalary,
                },
            },
            CreatedAt: types.MustTimeFromString("2022-03-16T15:29:37.822Z"),
            Departments: []string{
                "Sports",
            },
            Description: unifiedgosdk.String("Operative bi-directional capability"),
            EmploymentType: shared.AtsJobEmploymentTypeIntern.ToPointer(),
            HiringManagerIds: []string{
                "Hop",
            },
            ID: unifiedgosdk.String("<ID>"),
            LanguageLocale: unifiedgosdk.String("hence gracefully invoice"),
            Name: unifiedgosdk.String("Southeast vacantly Uranium"),
            PublicJobUrls: []string{
                "Keith",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "happily",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusOpen.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-08-04T07:33:03.088Z"),
        },
        ConnectionID: "Cis benchmark",
        ID: "<ID>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchAtsConnectionIDJobIDRequest](../../models/operations/patchatsconnectionidjobidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PatchAtsConnectionIDJobIDResponse](../../models/operations/patchatsconnectionidjobidresponse.md), error**


## PostAtsConnectionIDJob

Create a job

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.PostAtsConnectionIDJob(ctx, operations.PostAtsConnectionIDJobRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.String("Forward"),
                    Address2: unifiedgosdk.String("Electric fuchsia kelvin"),
                    City: unifiedgosdk.String("Fort Sibylmouth"),
                    Country: unifiedgosdk.String("Solomon Islands"),
                    CountryCode: unifiedgosdk.String("DO"),
                    PostalCode: unifiedgosdk.String("39037"),
                    Region: unifiedgosdk.String("Rockford"),
                    RegionCode: unifiedgosdk.String("Trafficway eaque athwart"),
                },
            },
            ClosedAt: types.MustTimeFromString("2022-04-22T19:01:40.265Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("Gourde"),
                    Frequency: shared.AtsCompensationFrequencyYear.ToPointer(),
                    Max: unifiedgosdk.Float64(5349.62),
                    Min: unifiedgosdk.Float64(526.63),
                    Type: shared.AtsCompensationTypeEquity,
                },
            },
            CreatedAt: types.MustTimeFromString("2021-04-02T18:44:02.642Z"),
            Departments: []string{
                "Polonium",
            },
            Description: unifiedgosdk.String("Progressive disintermediate matrix"),
            EmploymentType: shared.AtsJobEmploymentTypeIntern.ToPointer(),
            HiringManagerIds: []string{
                "itaque",
            },
            ID: unifiedgosdk.String("<ID>"),
            LanguageLocale: unifiedgosdk.String("the joyfully"),
            Name: unifiedgosdk.String("Other because harbor"),
            PublicJobUrls: []string{
                "coil",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "JSON",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusArchived.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-05-24T04:21:24.567Z"),
        },
        ConnectionID: "Coordinator applications",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PostAtsConnectionIDJobRequest](../../models/operations/postatsconnectionidjobrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PostAtsConnectionIDJobResponse](../../models/operations/postatsconnectionidjobresponse.md), error**


## PutAtsConnectionIDJobID

Update a job

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.PutAtsConnectionIDJobID(ctx, operations.PutAtsConnectionIDJobIDRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedgosdk.String("cotton Washington"),
                    Address2: unifiedgosdk.String("rosin meanwhile male"),
                    City: unifiedgosdk.String("East Sierra"),
                    Country: unifiedgosdk.String("Somalia"),
                    CountryCode: unifiedgosdk.String("BQ"),
                    PostalCode: unifiedgosdk.String("63475-6123"),
                    Region: unifiedgosdk.String("lighthearted online Bicycle"),
                    RegionCode: unifiedgosdk.String("robust"),
                },
            },
            ClosedAt: types.MustTimeFromString("2021-05-06T11:53:52.940Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("Iranian Rial"),
                    Frequency: shared.AtsCompensationFrequencyYear.ToPointer(),
                    Max: unifiedgosdk.Float64(5965.42),
                    Min: unifiedgosdk.Float64(5273.81),
                    Type: shared.AtsCompensationTypeBonus,
                },
            },
            CreatedAt: types.MustTimeFromString("2023-07-19T02:36:00.215Z"),
            Departments: []string{
                "embrace",
            },
            Description: unifiedgosdk.String("Programmable tertiary benchmark"),
            EmploymentType: shared.AtsJobEmploymentTypeContractor.ToPointer(),
            HiringManagerIds: []string{
                "New",
            },
            ID: unifiedgosdk.String("<ID>"),
            LanguageLocale: unifiedgosdk.String("Facilitator Gasoline application"),
            Name: unifiedgosdk.String("North impractical"),
            PublicJobUrls: []string{
                "clamber",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "West",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusClosed.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-02-01T04:41:47.094Z"),
        },
        ConnectionID: "North",
        ID: "<ID>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PutAtsConnectionIDJobIDRequest](../../models/operations/putatsconnectionidjobidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PutAtsConnectionIDJobIDResponse](../../models/operations/putatsconnectionidjobidresponse.md), error**

