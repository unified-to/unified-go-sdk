# Job

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
        ConnectionID: "adipisci",
        ID: "2e3b49db-e0f2-43b7-b6d9-948d6eded477",
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
        ConnectionID: "voluptas",
        Limit: unifiedgosdk.Float64(5378.51),
        Offset: unifiedgosdk.Float64(504.05),
        Order: unifiedgosdk.String("reiciendis"),
        Query: unifiedgosdk.String("minus"),
        Sort: unifiedgosdk.String("iure"),
        UpdatedGte: types.MustTimeFromString("2022-11-11T00:39:35.207Z"),
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
        ConnectionID: "iure",
        ID: "a82e5e82-fd28-4d10-80a7-e91392ab44cb",
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
                    Address1: unifiedgosdk.String("architecto"),
                    Address2: unifiedgosdk.String("totam"),
                    City: unifiedgosdk.String("West Aaron"),
                    Country: unifiedgosdk.String("Aruba"),
                    CountryCode: unifiedgosdk.String("LA"),
                    PostalCode: unifiedgosdk.String("23078-3185"),
                    Region: unifiedgosdk.String("illo"),
                    RegionCode: unifiedgosdk.String("tempora"),
                },
            },
            ClosedAt: types.MustTimeFromString("2022-05-27T21:54:13.074Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("rem"),
                    Frequency: shared.AtsCompensationFrequencyHour.ToPointer(),
                    Max: unifiedgosdk.Float64(6129.79),
                    Min: unifiedgosdk.Float64(7255.92),
                    Type: shared.AtsCompensationTypeEquity,
                },
            },
            CreatedAt: types.MustTimeFromString("2022-08-10T22:57:33.587Z"),
            Departments: []string{
                "aperiam",
            },
            Description: unifiedgosdk.String("similique"),
            EmploymentType: shared.AtsJobEmploymentTypeFreelance.ToPointer(),
            HiringManagerIds: []string{
                "pariatur",
            },
            ID: unifiedgosdk.String("fde410c3-7daa-4918-aa49-d9625d3caffc"),
            LanguageLocale: unifiedgosdk.String("inventore"),
            Name: unifiedgosdk.String("Guy Von"),
            PublicJobUrls: []string{
                "modi",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "quaerat",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusPending.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-07-04T21:18:33.130Z"),
        },
        ConnectionID: "unde",
        ID: "2bcd440e-a98b-4ecc-a048-6de0d56d73b0",
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
                    Address1: unifiedgosdk.String("quae"),
                    Address2: unifiedgosdk.String("quis"),
                    City: unifiedgosdk.String("North Consuelo"),
                    Country: unifiedgosdk.String("Turkmenistan"),
                    CountryCode: unifiedgosdk.String("LK"),
                    PostalCode: unifiedgosdk.String("74149-9447"),
                    Region: unifiedgosdk.String("iure"),
                    RegionCode: unifiedgosdk.String("ullam"),
                },
            },
            ClosedAt: types.MustTimeFromString("2022-07-06T19:55:18.068Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("enim"),
                    Frequency: shared.AtsCompensationFrequencyWeek.ToPointer(),
                    Max: unifiedgosdk.Float64(3432.31),
                    Min: unifiedgosdk.Float64(6902.62),
                    Type: shared.AtsCompensationTypeStockOptions,
                },
            },
            CreatedAt: types.MustTimeFromString("2022-01-28T21:56:06.312Z"),
            Departments: []string{
                "consectetur",
            },
            Description: unifiedgosdk.String("vero"),
            EmploymentType: shared.AtsJobEmploymentTypeContractor.ToPointer(),
            HiringManagerIds: []string{
                "optio",
            },
            ID: unifiedgosdk.String("fcc6a91e-c526-424d-8001-4ef45cea11ac"),
            LanguageLocale: unifiedgosdk.String("minima"),
            Name: unifiedgosdk.String("Henrietta Powlowski"),
            PublicJobUrls: []string{
                "nostrum",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "corrupti",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusDraft.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-04-19T01:37:46.134Z"),
        },
        ConnectionID: "eius",
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
                    Address1: unifiedgosdk.String("voluptatem"),
                    Address2: unifiedgosdk.String("magnam"),
                    City: unifiedgosdk.String("West Randibury"),
                    Country: unifiedgosdk.String("Rwanda"),
                    CountryCode: unifiedgosdk.String("MK"),
                    PostalCode: unifiedgosdk.String("89820-0695"),
                    Region: unifiedgosdk.String("doloribus"),
                    RegionCode: unifiedgosdk.String("unde"),
                },
            },
            ClosedAt: types.MustTimeFromString("2022-03-17T10:24:00.538Z"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedgosdk.String("id"),
                    Frequency: shared.AtsCompensationFrequencyWeek.ToPointer(),
                    Max: unifiedgosdk.Float64(1012.53),
                    Min: unifiedgosdk.Float64(7482.56),
                    Type: shared.AtsCompensationTypeSalary,
                },
            },
            CreatedAt: types.MustTimeFromString("2022-08-21T15:07:46.264Z"),
            Departments: []string{
                "a",
            },
            Description: unifiedgosdk.String("architecto"),
            EmploymentType: shared.AtsJobEmploymentTypeFreelance.ToPointer(),
            HiringManagerIds: []string{
                "vitae",
            },
            ID: unifiedgosdk.String("4718c6fa-2fad-40c0-ac5d-95472cdd14fc"),
            LanguageLocale: unifiedgosdk.String("eius"),
            Name: unifiedgosdk.String("Miss Bridget King"),
            PublicJobUrls: []string{
                "fuga",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "laudantium",
            },
            Remote: unifiedgosdk.Bool(false),
            Status: shared.AtsJobStatusDraft.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2020-12-28T11:43:36.436Z"),
        },
        ConnectionID: "dignissimos",
        ID: "0c43351c-3dd1-4eb8-b7f7-5f4f23f1c0a5",
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

