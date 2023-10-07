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
                shared.AtsAddress{},
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Type: shared.AtsCompensationTypeBonus,
                },
            },
            Departments: []string{
                "Transexual",
            },
            HiringManagerIds: []string{
                "leach",
            },
            PublicJobUrls: []string{
                "national",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "Kia",
            },
        },
        ConnectionID: "Ferrari Facilitator",
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
                shared.AtsAddress{},
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Type: shared.AtsCompensationTypeSalary,
                },
            },
            Departments: []string{
                "Forward",
            },
            HiringManagerIds: []string{
                "Americium",
            },
            PublicJobUrls: []string{
                "shiny",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "yellow",
            },
        },
        ConnectionID: "neatly Diesel virtual",
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
                shared.AtsAddress{},
            },
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Type: shared.AtsCompensationTypeStockOptions,
                },
            },
            Departments: []string{
                "cotton",
            },
            HiringManagerIds: []string{
                "Washington",
            },
            PublicJobUrls: []string{
                "Hybrid",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "henry",
            },
        },
        ConnectionID: "male Intelligent",
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

