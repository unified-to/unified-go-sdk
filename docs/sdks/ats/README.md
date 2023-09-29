# Ats
(*Ats*)

### Available Operations

* [DeleteAtsConnectionIDApplicationID](#deleteatsconnectionidapplicationid) - Remove an application
* [DeleteAtsConnectionIDCandidateID](#deleteatsconnectionidcandidateid) - Remove a candidate
* [DeleteAtsConnectionIDInterviewID](#deleteatsconnectionidinterviewid) - Remove a interview
* [DeleteAtsConnectionIDJobID](#deleteatsconnectionidjobid) - Remove a job
* [DeleteAtsConnectionIDScorecardID](#deleteatsconnectionidscorecardid) - Remove a scorecard
* [GetAtsConnectionIDApplication](#getatsconnectionidapplication) - List all applications
* [GetAtsConnectionIDApplicationID](#getatsconnectionidapplicationid) - Retrieve an application
* [GetAtsConnectionIDCandidate](#getatsconnectionidcandidate) - List all candidates
* [GetAtsConnectionIDCandidateID](#getatsconnectionidcandidateid) - Retrieve a candidate
* [GetAtsConnectionIDInterview](#getatsconnectionidinterview) - List all interviews
* [GetAtsConnectionIDInterviewID](#getatsconnectionidinterviewid) - Retrieve a interview
* [GetAtsConnectionIDJob](#getatsconnectionidjob) - List all jobs
* [GetAtsConnectionIDJobID](#getatsconnectionidjobid) - Retrieve a job
* [GetAtsConnectionIDScorecard](#getatsconnectionidscorecard) - List all scorecards
* [GetAtsConnectionIDScorecardID](#getatsconnectionidscorecardid) - Retrieve a scorecard
* [PatchAtsConnectionIDApplicationID](#patchatsconnectionidapplicationid) - Update an application
* [PatchAtsConnectionIDCandidateID](#patchatsconnectionidcandidateid) - Update a candidate
* [PatchAtsConnectionIDInterviewID](#patchatsconnectionidinterviewid) - Update a interview
* [PatchAtsConnectionIDJobID](#patchatsconnectionidjobid) - Update a job
* [PatchAtsConnectionIDScorecardID](#patchatsconnectionidscorecardid) - Update a scorecard
* [PostAtsConnectionIDApplication](#postatsconnectionidapplication) - Create an application
* [PostAtsConnectionIDCandidate](#postatsconnectionidcandidate) - Create a candidate
* [PostAtsConnectionIDInterview](#postatsconnectionidinterview) - Create a interview
* [PostAtsConnectionIDJob](#postatsconnectionidjob) - Create a job
* [PostAtsConnectionIDScorecard](#postatsconnectionidscorecard) - Create a scorecard
* [PutAtsConnectionIDApplicationID](#putatsconnectionidapplicationid) - Update an application
* [PutAtsConnectionIDCandidateID](#putatsconnectionidcandidateid) - Update a candidate
* [PutAtsConnectionIDInterviewID](#putatsconnectionidinterviewid) - Update a interview
* [PutAtsConnectionIDJobID](#putatsconnectionidjobid) - Update a job
* [PutAtsConnectionIDScorecardID](#putatsconnectionidscorecardid) - Update a scorecard

## DeleteAtsConnectionIDApplicationID

Remove an application

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
    res, err := s.Ats.DeleteAtsConnectionIDApplicationID(ctx, operations.DeleteAtsConnectionIDApplicationIDRequest{
        ConnectionID: "markets sievert meh",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.DeleteAtsConnectionIDApplicationIDRequest](../../models/operations/deleteatsconnectionidapplicationidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.DeleteAtsConnectionIDApplicationIDResponse](../../models/operations/deleteatsconnectionidapplicationidresponse.md), error**


## DeleteAtsConnectionIDCandidateID

Remove a candidate

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
    res, err := s.Ats.DeleteAtsConnectionIDCandidateID(ctx, operations.DeleteAtsConnectionIDCandidateIDRequest{
        ConnectionID: "multimedia",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteAtsConnectionIDCandidateIDRequest](../../models/operations/deleteatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeleteAtsConnectionIDCandidateIDResponse](../../models/operations/deleteatsconnectionidcandidateidresponse.md), error**


## DeleteAtsConnectionIDInterviewID

Remove a interview

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
    res, err := s.Ats.DeleteAtsConnectionIDInterviewID(ctx, operations.DeleteAtsConnectionIDInterviewIDRequest{
        ConnectionID: "redundant Health Hayes",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteAtsConnectionIDInterviewIDRequest](../../models/operations/deleteatsconnectionidinterviewidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeleteAtsConnectionIDInterviewIDResponse](../../models/operations/deleteatsconnectionidinterviewidresponse.md), error**


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
    res, err := s.Ats.DeleteAtsConnectionIDJobID(ctx, operations.DeleteAtsConnectionIDJobIDRequest{
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


## DeleteAtsConnectionIDScorecardID

Remove a scorecard

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
    res, err := s.Ats.DeleteAtsConnectionIDScorecardID(ctx, operations.DeleteAtsConnectionIDScorecardIDRequest{
        ConnectionID: "Agent intrepid",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteAtsConnectionIDScorecardIDRequest](../../models/operations/deleteatsconnectionidscorecardidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeleteAtsConnectionIDScorecardIDResponse](../../models/operations/deleteatsconnectionidscorecardidresponse.md), error**


## GetAtsConnectionIDApplication

List all applications

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
    res, err := s.Ats.GetAtsConnectionIDApplication(ctx, operations.GetAtsConnectionIDApplicationRequest{
        CandidateID: unifiedgosdk.String("turquoise"),
        ConnectionID: "Regional Bedfordshire",
        JobID: unifiedgosdk.String("Northwest portal Electric"),
        Limit: unifiedgosdk.Float64(576.8),
        Offset: unifiedgosdk.Float64(7467.13),
        Order: unifiedgosdk.String("Architect"),
        Query: unifiedgosdk.String("loosely contingency"),
        Sort: unifiedgosdk.String("female"),
        UpdatedGte: types.MustTimeFromString("2023-09-05T13:59:23.348Z"),
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAtsConnectionIDApplicationRequest](../../models/operations/getatsconnectionidapplicationrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAtsConnectionIDApplicationResponse](../../models/operations/getatsconnectionidapplicationresponse.md), error**


## GetAtsConnectionIDApplicationID

Retrieve an application

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
    res, err := s.Ats.GetAtsConnectionIDApplicationID(ctx, operations.GetAtsConnectionIDApplicationIDRequest{
        ConnectionID: "Buckinghamshire functionalities",
        ID: "<ID>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetAtsConnectionIDApplicationIDRequest](../../models/operations/getatsconnectionidapplicationidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetAtsConnectionIDApplicationIDResponse](../../models/operations/getatsconnectionidapplicationidresponse.md), error**


## GetAtsConnectionIDCandidate

List all candidates

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
    res, err := s.Ats.GetAtsConnectionIDCandidate(ctx, operations.GetAtsConnectionIDCandidateRequest{
        ConnectionID: "Northwest forceful Moore",
        Limit: unifiedgosdk.Float64(2623.89),
        Offset: unifiedgosdk.Float64(7811.91),
        Order: unifiedgosdk.String("Mouse whether deploy"),
        Query: unifiedgosdk.String("pink"),
        Sort: unifiedgosdk.String("huzzah thistle"),
        UpdatedGte: types.MustTimeFromString("2022-03-13T15:14:03.645Z"),
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAtsConnectionIDCandidateRequest](../../models/operations/getatsconnectionidcandidaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetAtsConnectionIDCandidateResponse](../../models/operations/getatsconnectionidcandidateresponse.md), error**


## GetAtsConnectionIDCandidateID

Retrieve a candidate

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
    res, err := s.Ats.GetAtsConnectionIDCandidateID(ctx, operations.GetAtsConnectionIDCandidateIDRequest{
        ConnectionID: "ha Loan",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAtsConnectionIDCandidateIDRequest](../../models/operations/getatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAtsConnectionIDCandidateIDResponse](../../models/operations/getatsconnectionidcandidateidresponse.md), error**


## GetAtsConnectionIDInterview

List all interviews

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
    res, err := s.Ats.GetAtsConnectionIDInterview(ctx, operations.GetAtsConnectionIDInterviewRequest{
        ApplicationID: unifiedgosdk.String("Fresh Pickup converse"),
        ConnectionID: "vortals",
        Limit: unifiedgosdk.Float64(5167.08),
        Offset: unifiedgosdk.Float64(6488.61),
        Order: unifiedgosdk.String("Oregon Metal"),
        Query: unifiedgosdk.String("Account"),
        Sort: unifiedgosdk.String("haptic"),
        UpdatedGte: types.MustTimeFromString("2021-09-23T19:46:35.825Z"),
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAtsConnectionIDInterviewRequest](../../models/operations/getatsconnectionidinterviewrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetAtsConnectionIDInterviewResponse](../../models/operations/getatsconnectionidinterviewresponse.md), error**


## GetAtsConnectionIDInterviewID

Retrieve a interview

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
    res, err := s.Ats.GetAtsConnectionIDInterviewID(ctx, operations.GetAtsConnectionIDInterviewIDRequest{
        ConnectionID: "Loan Gorgeous lux",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAtsConnectionIDInterviewIDRequest](../../models/operations/getatsconnectionidinterviewidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAtsConnectionIDInterviewIDResponse](../../models/operations/getatsconnectionidinterviewidresponse.md), error**


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
    res, err := s.Ats.GetAtsConnectionIDJob(ctx, operations.GetAtsConnectionIDJobRequest{
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
    res, err := s.Ats.GetAtsConnectionIDJobID(ctx, operations.GetAtsConnectionIDJobIDRequest{
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


## GetAtsConnectionIDScorecard

List all scorecards

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
    res, err := s.Ats.GetAtsConnectionIDScorecard(ctx, operations.GetAtsConnectionIDScorecardRequest{
        ApplicationID: unifiedgosdk.String("Licensed deep"),
        CandidateID: unifiedgosdk.String("happily"),
        ConnectionID: "lunch accusamus",
        InterviewID: unifiedgosdk.String("for famously Southwest"),
        Limit: unifiedgosdk.Float64(950.05),
        Offset: unifiedgosdk.Float64(6133.23),
        Order: unifiedgosdk.String("withdrawal"),
        Query: unifiedgosdk.String("Bicycle copy Bronze"),
        Sort: unifiedgosdk.String("ouch non ut"),
        UpdatedGte: types.MustTimeFromString("2021-06-01T09:53:52.927Z"),
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAtsConnectionIDScorecardRequest](../../models/operations/getatsconnectionidscorecardrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetAtsConnectionIDScorecardResponse](../../models/operations/getatsconnectionidscorecardresponse.md), error**


## GetAtsConnectionIDScorecardID

Retrieve a scorecard

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
    res, err := s.Ats.GetAtsConnectionIDScorecardID(ctx, operations.GetAtsConnectionIDScorecardIDRequest{
        ConnectionID: "East mobile Mini",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAtsConnectionIDScorecardIDRequest](../../models/operations/getatsconnectionidscorecardidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAtsConnectionIDScorecardIDResponse](../../models/operations/getatsconnectionidscorecardidresponse.md), error**


## PatchAtsConnectionIDApplicationID

Update an application

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
    res, err := s.Ats.PatchAtsConnectionIDApplicationID(ctx, operations.PatchAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2023-10-17T09:51:42.165Z"),
            CandidateID: unifiedgosdk.String("North et beyond"),
            CreatedAt: types.MustTimeFromString("2023-01-08T08:26:22.845Z"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("ick Sausages Bronze"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2023-10-18T00:47:25.469Z"),
            RejectedReason: unifiedgosdk.String("Avon Sum quis"),
            Source: unifiedgosdk.String("Carolina Wooden Pop"),
            Status: shared.AtsApplicationStatusHired.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-07-20T22:05:46.009Z"),
        },
        ConnectionID: "Baby Paucek",
        ID: "<ID>",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PatchAtsConnectionIDApplicationIDRequest](../../models/operations/patchatsconnectionidapplicationidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PatchAtsConnectionIDApplicationIDResponse](../../models/operations/patchatsconnectionidapplicationidresponse.md), error**


## PatchAtsConnectionIDCandidateID

Update a candidate

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
    res, err := s.Ats.PatchAtsConnectionIDCandidateID(ctx, operations.PatchAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("closely Goyette plus"),
                Address2: unifiedgosdk.String("culpa"),
                City: unifiedgosdk.String("Darrinshire"),
                Country: unifiedgosdk.String("Mongolia"),
                CountryCode: unifiedgosdk.String("GW"),
                PostalCode: unifiedgosdk.String("05275"),
                Region: unifiedgosdk.String("TLS calculating"),
                RegionCode: unifiedgosdk.String("up Argon Internal"),
            },
            CompanyName: unifiedgosdk.String("Fadel, Schulist and Koss"),
            CreatedAt: types.MustTimeFromString("2022-12-09T07:16:54.728Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Gregory63@gmail.com",
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("Elegant"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Tricycle Yttrium Hybrid"),
            Name: unifiedgosdk.String("ornery whether"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "Cadillac",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "Marketing Cotton",
                    Type: shared.AtsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("East"),
            UpdatedAt: types.MustTimeFromString("2023-10-31T11:53:36.953Z"),
        },
        ConnectionID: "redundant Tricycle unloose",
        ID: "<ID>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchAtsConnectionIDCandidateIDRequest](../../models/operations/patchatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PatchAtsConnectionIDCandidateIDResponse](../../models/operations/patchatsconnectionidcandidateidresponse.md), error**


## PatchAtsConnectionIDInterviewID

Update a interview

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
    res, err := s.Ats.PatchAtsConnectionIDInterviewID(ctx, operations.PatchAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedgosdk.String("SSD green pascal"),
            CandidateID: unifiedgosdk.String("Buckinghamshire example"),
            CreatedAt: types.MustTimeFromString("2021-08-24T08:30:07.073Z"),
            EndAt: types.MustTimeFromString("2021-06-27T04:06:46.373Z"),
            ExternalEventXref: unifiedgosdk.String("apropos Gadolinium"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("transgender transmitting"),
            Location: unifiedgosdk.String("Investor synthesizing"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustTimeFromString("2021-01-19T01:51:02.213Z"),
            Status: shared.AtsInterviewStatusAwaitingFeedback.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-01-21T17:38:09.113Z"),
            UserIds: []string{
                "Honda",
            },
        },
        ConnectionID: "Myrl Dram Trail",
        ID: "<ID>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchAtsConnectionIDInterviewIDRequest](../../models/operations/patchatsconnectionidinterviewidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PatchAtsConnectionIDInterviewIDResponse](../../models/operations/patchatsconnectionidinterviewidresponse.md), error**


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
    res, err := s.Ats.PatchAtsConnectionIDJobID(ctx, operations.PatchAtsConnectionIDJobIDRequest{
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


## PatchAtsConnectionIDScorecardID

Update a scorecard

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
    res, err := s.Ats.PatchAtsConnectionIDScorecardID(ctx, operations.PatchAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedgosdk.String("Carter Hatchback functionalities"),
            CandidateID: unifiedgosdk.String("disagree gold New"),
            CreatedAt: types.MustTimeFromString("2023-05-08T15:11:07.692Z"),
            ID: unifiedgosdk.String("<ID>"),
            InterviewID: unifiedgosdk.String("blue"),
            InterviewerID: unifiedgosdk.String("North Buckinghamshire blur"),
            JobID: unifiedgosdk.String("kelvin hack Fantastic"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationDefinitelyNo.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-11-18T04:49:38.005Z"),
        },
        ConnectionID: "hacking meter",
        ID: "<ID>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchAtsConnectionIDScorecardIDRequest](../../models/operations/patchatsconnectionidscorecardidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PatchAtsConnectionIDScorecardIDResponse](../../models/operations/patchatsconnectionidscorecardidresponse.md), error**


## PostAtsConnectionIDApplication

Create an application

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
    res, err := s.Ats.PostAtsConnectionIDApplication(ctx, operations.PostAtsConnectionIDApplicationRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2021-10-26T15:24:28.979Z"),
            CandidateID: unifiedgosdk.String("solid"),
            CreatedAt: types.MustTimeFromString("2022-09-13T17:17:33.049Z"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("Gloves Pizza virtual"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2023-12-27T18:41:56.821Z"),
            RejectedReason: unifiedgosdk.String("Northwest Kids"),
            Source: unifiedgosdk.String("Human Tasty Loan"),
            Status: shared.AtsApplicationStatusNew.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-11-01T21:08:50.319Z"),
        },
        ConnectionID: "Jazz",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostAtsConnectionIDApplicationRequest](../../models/operations/postatsconnectionidapplicationrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostAtsConnectionIDApplicationResponse](../../models/operations/postatsconnectionidapplicationresponse.md), error**


## PostAtsConnectionIDCandidate

Create a candidate

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
    res, err := s.Ats.PostAtsConnectionIDCandidate(ctx, operations.PostAtsConnectionIDCandidateRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("incubate"),
                Address2: unifiedgosdk.String("azure Trans"),
                City: unifiedgosdk.String("Port Rory"),
                Country: unifiedgosdk.String("El Salvador"),
                CountryCode: unifiedgosdk.String("CX"),
                PostalCode: unifiedgosdk.String("54222-0235"),
                Region: unifiedgosdk.String("modi fooey"),
                RegionCode: unifiedgosdk.String("Metal TCP incidunt"),
            },
            CompanyName: unifiedgosdk.String("McCullough, Rosenbaum and Daugherty"),
            CreatedAt: types.MustTimeFromString("2023-02-07T05:55:59.357Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Eleanora.Rogahn44@hotmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("South though"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Pants"),
            Name: unifiedgosdk.String("Raleigh"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "morph",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "lavender Sedan Folk",
                    Type: shared.AtsTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Savings panel"),
            UpdatedAt: types.MustTimeFromString("2022-02-09T15:32:35.578Z"),
        },
        ConnectionID: "Ngultrum red glean",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAtsConnectionIDCandidateRequest](../../models/operations/postatsconnectionidcandidaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostAtsConnectionIDCandidateResponse](../../models/operations/postatsconnectionidcandidateresponse.md), error**


## PostAtsConnectionIDInterview

Create a interview

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
    res, err := s.Ats.PostAtsConnectionIDInterview(ctx, operations.PostAtsConnectionIDInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedgosdk.String("round Hat Savings"),
            CandidateID: unifiedgosdk.String("Northeast"),
            CreatedAt: types.MustTimeFromString("2022-12-27T10:33:09.160Z"),
            EndAt: types.MustTimeFromString("2021-11-12T23:57:19.974Z"),
            ExternalEventXref: unifiedgosdk.String("platforms"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("payment panel Identity"),
            Location: unifiedgosdk.String("Northwest Buckinghamshire"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustTimeFromString("2022-11-02T05:07:18.592Z"),
            Status: shared.AtsInterviewStatusComplete.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-07-13T16:35:04.177Z"),
            UserIds: []string{
                "Chevrolet",
            },
        },
        ConnectionID: "Shoes Northeast SMTP",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAtsConnectionIDInterviewRequest](../../models/operations/postatsconnectionidinterviewrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostAtsConnectionIDInterviewResponse](../../models/operations/postatsconnectionidinterviewresponse.md), error**


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
    res, err := s.Ats.PostAtsConnectionIDJob(ctx, operations.PostAtsConnectionIDJobRequest{
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


## PostAtsConnectionIDScorecard

Create a scorecard

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
    res, err := s.Ats.PostAtsConnectionIDScorecard(ctx, operations.PostAtsConnectionIDScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedgosdk.String("female bah"),
            CandidateID: unifiedgosdk.String("if since"),
            CreatedAt: types.MustTimeFromString("2022-02-26T00:06:29.981Z"),
            ID: unifiedgosdk.String("<ID>"),
            InterviewID: unifiedgosdk.String("invoice"),
            InterviewerID: unifiedgosdk.String("male"),
            JobID: unifiedgosdk.String("Accountability"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationStrongYes.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-10-04T17:15:51.015Z"),
        },
        ConnectionID: "Legacy tan",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAtsConnectionIDScorecardRequest](../../models/operations/postatsconnectionidscorecardrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostAtsConnectionIDScorecardResponse](../../models/operations/postatsconnectionidscorecardresponse.md), error**


## PutAtsConnectionIDApplicationID

Update an application

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
    res, err := s.Ats.PutAtsConnectionIDApplicationID(ctx, operations.PutAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2022-06-15T22:25:51.833Z"),
            CandidateID: unifiedgosdk.String("farad Indianapolis"),
            CreatedAt: types.MustTimeFromString("2022-04-01T21:03:58.880Z"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("enable foreground"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2022-11-06T11:41:30.414Z"),
            RejectedReason: unifiedgosdk.String("virtual North plum"),
            Source: unifiedgosdk.String("Fort solid"),
            Status: shared.AtsApplicationStatusSubmitted.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-02-21T04:47:57.079Z"),
        },
        ConnectionID: "Southeast",
        ID: "<ID>",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PutAtsConnectionIDApplicationIDRequest](../../models/operations/putatsconnectionidapplicationidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PutAtsConnectionIDApplicationIDResponse](../../models/operations/putatsconnectionidapplicationidresponse.md), error**


## PutAtsConnectionIDCandidateID

Update a candidate

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
    res, err := s.Ats.PutAtsConnectionIDCandidateID(ctx, operations.PutAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("archive"),
                Address2: unifiedgosdk.String("Specialist Kyat"),
                City: unifiedgosdk.String("New Dennis"),
                Country: unifiedgosdk.String("Mauritius"),
                CountryCode: unifiedgosdk.String("TL"),
                PostalCode: unifiedgosdk.String("49105-9909"),
                Region: unifiedgosdk.String("copy olive"),
                RegionCode: unifiedgosdk.String("withdrawal cumque person"),
            },
            CompanyName: unifiedgosdk.String("Kuhn and Sons"),
            CreatedAt: types.MustTimeFromString("2022-01-28T10:51:00.922Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Hester.Jenkins@gmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("Loan EXE"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("deliver executive RSS"),
            Name: unifiedgosdk.String("because aha black"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "program",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "empower exit Pangender",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Corporate anenst Electronic"),
            UpdatedAt: types.MustTimeFromString("2022-03-30T08:00:53.284Z"),
        },
        ConnectionID: "Flerovium azure",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutAtsConnectionIDCandidateIDRequest](../../models/operations/putatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutAtsConnectionIDCandidateIDResponse](../../models/operations/putatsconnectionidcandidateidresponse.md), error**


## PutAtsConnectionIDInterviewID

Update a interview

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
    res, err := s.Ats.PutAtsConnectionIDInterviewID(ctx, operations.PutAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedgosdk.String("Generic capacitor"),
            CandidateID: unifiedgosdk.String("Road disbelieve"),
            CreatedAt: types.MustTimeFromString("2022-06-22T01:57:06.573Z"),
            EndAt: types.MustTimeFromString("2022-05-28T02:29:32.144Z"),
            ExternalEventXref: unifiedgosdk.String("architectures"),
            ID: unifiedgosdk.String("<ID>"),
            JobID: unifiedgosdk.String("Casper 1080p South"),
            Location: unifiedgosdk.String("program siemens Cis"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustTimeFromString("2021-03-14T15:20:41.084Z"),
            Status: shared.AtsInterviewStatusAwaitingFeedback.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-07-14T19:59:39.905Z"),
            UserIds: []string{
                "East",
            },
        },
        ConnectionID: "ASCII yet Hybrid",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutAtsConnectionIDInterviewIDRequest](../../models/operations/putatsconnectionidinterviewidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutAtsConnectionIDInterviewIDResponse](../../models/operations/putatsconnectionidinterviewidresponse.md), error**


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
    res, err := s.Ats.PutAtsConnectionIDJobID(ctx, operations.PutAtsConnectionIDJobIDRequest{
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


## PutAtsConnectionIDScorecardID

Update a scorecard

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
    res, err := s.Ats.PutAtsConnectionIDScorecardID(ctx, operations.PutAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedgosdk.String("East Granite"),
            CandidateID: unifiedgosdk.String("South"),
            CreatedAt: types.MustTimeFromString("2022-03-02T12:33:41.490Z"),
            ID: unifiedgosdk.String("<ID>"),
            InterviewID: unifiedgosdk.String("Texas Technetium hack"),
            InterviewerID: unifiedgosdk.String("Adventure Kyrgyz Organic"),
            JobID: unifiedgosdk.String("Home Dynamic Integration"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationNo.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-11-13T03:01:57.066Z"),
        },
        ConnectionID: "Transexual Manager Rap",
        ID: "<ID>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutAtsConnectionIDScorecardIDRequest](../../models/operations/putatsconnectionidscorecardidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutAtsConnectionIDScorecardIDResponse](../../models/operations/putatsconnectionidscorecardidresponse.md), error**

