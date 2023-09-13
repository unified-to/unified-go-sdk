# Ats

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDApplicationIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.DeleteAtsConnectionIDApplicationID(ctx, operations.DeleteAtsConnectionIDApplicationIDRequest{
        ConnectionID: "nulla",
        ID: "6b144290-7474-4778-a7bd-466d28c10ab3",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteAtsConnectionIDApplicationIDRequest](../../models/operations/deleteatsconnectionidapplicationidrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.DeleteAtsConnectionIDApplicationIDSecurity](../../models/operations/deleteatsconnectionidapplicationidsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.DeleteAtsConnectionIDCandidateID(ctx, operations.DeleteAtsConnectionIDCandidateIDRequest{
        ConnectionID: "quo",
        ID: "dca42519-04e5-423c-be0b-c7178e4796f2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAtsConnectionIDCandidateIDRequest](../../models/operations/deleteatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeleteAtsConnectionIDCandidateIDSecurity](../../models/operations/deleteatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDInterviewIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.DeleteAtsConnectionIDInterviewID(ctx, operations.DeleteAtsConnectionIDInterviewIDRequest{
        ConnectionID: "deserunt",
        ID: "70c68828-2aa4-4825-a2f2-22e9817ee17c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAtsConnectionIDInterviewIDRequest](../../models/operations/deleteatsconnectionidinterviewidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeleteAtsConnectionIDInterviewIDSecurity](../../models/operations/deleteatsconnectionidinterviewidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDJobIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.DeleteAtsConnectionIDJobID(ctx, operations.DeleteAtsConnectionIDJobIDRequest{
        ConnectionID: "nam",
        ID: "e61e6b7b-95bc-40ab-bc20-c4f3789fd871",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteAtsConnectionIDJobIDRequest](../../models/operations/deleteatsconnectionidjobidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.DeleteAtsConnectionIDJobIDSecurity](../../models/operations/deleteatsconnectionidjobidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDScorecardIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.DeleteAtsConnectionIDScorecardID(ctx, operations.DeleteAtsConnectionIDScorecardIDRequest{
        ConnectionID: "a",
        ID: "99dd2efd-121a-4a6f-9e67-4bdb04f15756",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAtsConnectionIDScorecardIDRequest](../../models/operations/deleteatsconnectionidscorecardidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeleteAtsConnectionIDScorecardIDSecurity](../../models/operations/deleteatsconnectionidscorecardidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDApplicationSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDApplication(ctx, operations.GetAtsConnectionIDApplicationRequest{
        CandidateID: unifiedto.String("aut"),
        ConnectionID: "voluptatum",
        JobID: unifiedto.String("qui"),
        Limit: unifiedto.Float64(8453.58),
        Offset: unifiedto.Float64(4012.59),
        Order: unifiedto.String("deleniti"),
        Query: unifiedto.String("itaque"),
        Sort: unifiedto.String("dolorum"),
        UpdatedGte: types.MustDateFromString("2022-05-23"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsApplications != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAtsConnectionIDApplicationRequest](../../models/operations/getatsconnectionidapplicationrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetAtsConnectionIDApplicationSecurity](../../models/operations/getatsconnectionidapplicationsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDApplicationIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDApplicationID(ctx, operations.GetAtsConnectionIDApplicationIDRequest{
        ConnectionID: "tenetur",
        ID: "1d170513-39d0-4808-aa18-40394c26071f",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetAtsConnectionIDApplicationIDRequest](../../models/operations/getatsconnectionidapplicationidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.GetAtsConnectionIDApplicationIDSecurity](../../models/operations/getatsconnectionidapplicationidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDCandidateSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDCandidate(ctx, operations.GetAtsConnectionIDCandidateRequest{
        ConnectionID: "natus",
        Limit: unifiedto.Float64(2446.51),
        Offset: unifiedto.Float64(9742.57),
        Order: unifiedto.String("voluptas"),
        Query: unifiedto.String("asperiores"),
        Sort: unifiedto.String("aperiam"),
        UpdatedGte: types.MustDateFromString("2022-09-09"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAtsConnectionIDCandidateRequest](../../models/operations/getatsconnectionidcandidaterequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetAtsConnectionIDCandidateSecurity](../../models/operations/getatsconnectionidcandidatesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDCandidateID(ctx, operations.GetAtsConnectionIDCandidateIDRequest{
        ConnectionID: "consequuntur",
        ID: "dac7af51-5cc4-413a-a63a-ae8d67864dbb",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAtsConnectionIDCandidateIDRequest](../../models/operations/getatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetAtsConnectionIDCandidateIDSecurity](../../models/operations/getatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDInterviewSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDInterview(ctx, operations.GetAtsConnectionIDInterviewRequest{
        ApplicationID: unifiedto.String("commodi"),
        ConnectionID: "in",
        Limit: unifiedto.Float64(3605.45),
        Offset: unifiedto.Float64(9689.04),
        Order: unifiedto.String("assumenda"),
        Query: unifiedto.String("nemo"),
        Sort: unifiedto.String("recusandae"),
        UpdatedGte: types.MustDateFromString("2022-12-15"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAtsConnectionIDInterviewRequest](../../models/operations/getatsconnectionidinterviewrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetAtsConnectionIDInterviewSecurity](../../models/operations/getatsconnectionidinterviewsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDInterviewIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDInterviewID(ctx, operations.GetAtsConnectionIDInterviewIDRequest{
        ConnectionID: "cum",
        ID: "375ed4f6-fbee-441f-b331-7fe35b60eb1e",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAtsConnectionIDInterviewIDRequest](../../models/operations/getatsconnectionidinterviewidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetAtsConnectionIDInterviewIDSecurity](../../models/operations/getatsconnectionidinterviewidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDJobSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDJob(ctx, operations.GetAtsConnectionIDJobRequest{
        ConnectionID: "similique",
        Limit: unifiedto.Float64(2724.37),
        Offset: unifiedto.Float64(1328.15),
        Order: unifiedto.String("voluptas"),
        Query: unifiedto.String("voluptas"),
        Sort: unifiedto.String("voluptas"),
        UpdatedGte: types.MustDateFromString("2022-04-02"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAtsConnectionIDJobRequest](../../models/operations/getatsconnectionidjobrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetAtsConnectionIDJobSecurity](../../models/operations/getatsconnectionidjobsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDJobIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDJobID(ctx, operations.GetAtsConnectionIDJobIDRequest{
        ConnectionID: "dolorum",
        ID: "3c28744e-d53b-488f-ba8d-8f5c0b2f2fb7",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetAtsConnectionIDJobIDRequest](../../models/operations/getatsconnectionidjobidrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetAtsConnectionIDJobIDSecurity](../../models/operations/getatsconnectionidjobidsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDScorecardSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDScorecard(ctx, operations.GetAtsConnectionIDScorecardRequest{
        ApplicationID: unifiedto.String("expedita"),
        CandidateID: unifiedto.String("ab"),
        ConnectionID: "iste",
        InterviewID: unifiedto.String("dolore"),
        Limit: unifiedto.Float64(6719.07),
        Offset: unifiedto.Float64(1523.54),
        Order: unifiedto.String("in"),
        Query: unifiedto.String("commodi"),
        Sort: unifiedto.String("quidem"),
        UpdatedGte: types.MustDateFromString("2022-08-15"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAtsConnectionIDScorecardRequest](../../models/operations/getatsconnectionidscorecardrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetAtsConnectionIDScorecardSecurity](../../models/operations/getatsconnectionidscorecardsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDScorecardIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.GetAtsConnectionIDScorecardID(ctx, operations.GetAtsConnectionIDScorecardIDRequest{
        ConnectionID: "unde",
        ID: "16fe1f08-f429-44e3-a98f-447f603e8b44",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAtsConnectionIDScorecardIDRequest](../../models/operations/getatsconnectionidscorecardidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetAtsConnectionIDScorecardIDSecurity](../../models/operations/getatsconnectionidscorecardidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDApplicationIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PatchAtsConnectionIDApplicationID(ctx, operations.PatchAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustDateFromString("2022-02-08"),
            CandidateID: unifiedto.String("rem"),
            CreatedAt: types.MustDateFromString("2022-04-02"),
            ID: unifiedto.String("a55efd20-e457-4e18-98b6-a89fbe3a5aa8"),
            JobID: unifiedto.String("accusamus"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustDateFromString("2022-06-16"),
            RejectedReason: unifiedto.String("fugit"),
            Source: unifiedto.String("ut"),
            Status: shared.AtsApplicationStatusHired.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-05-14"),
        },
        ConnectionID: "expedita",
        ID: "4075088e-5186-4206-9e90-4f3b1194b8ab",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PatchAtsConnectionIDApplicationIDRequest](../../models/operations/patchatsconnectionidapplicationidrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.PatchAtsConnectionIDApplicationIDSecurity](../../models/operations/patchatsconnectionidapplicationidsecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PatchAtsConnectionIDCandidateID(ctx, operations.PatchAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("tenetur"),
                Address2: unifiedto.String("laboriosam"),
                City: unifiedto.String("East Lutherport"),
                Country: unifiedto.String("Morocco"),
                CountryCode: unifiedto.String("WF"),
                PostalCode: unifiedto.String("89906-6486"),
                Region: unifiedto.String("praesentium"),
                RegionCode: unifiedto.String("mollitia"),
            },
            CompanyName: unifiedto.String("veniam"),
            CreatedAt: types.MustDateFromString("2022-03-18"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Avis_Littel@hotmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("suscipit"),
            ID: unifiedto.String("bc173d68-9eee-4952-af8d-986e881ead4f"),
            ImageURL: unifiedto.String("doloremque"),
            Name: unifiedto.String("Mr. Keith Bashirian"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "laboriosam",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "velit",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedto.String("Ms."),
            UpdatedAt: types.MustDateFromString("2022-02-04"),
        },
        ConnectionID: "consequuntur",
        ID: "9e973e92-2a57-4a15-be3e-060807e2b6e3",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchAtsConnectionIDCandidateIDRequest](../../models/operations/patchatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PatchAtsConnectionIDCandidateIDSecurity](../../models/operations/patchatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDInterviewIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PatchAtsConnectionIDInterviewID(ctx, operations.PatchAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("laborum"),
            CandidateID: unifiedto.String("distinctio"),
            CreatedAt: types.MustDateFromString("2021-12-15"),
            EndAt: types.MustDateFromString("2022-09-05"),
            ExternalEventXref: unifiedto.String("repellat"),
            ID: unifiedto.String("0597a60f-f2a5-44a3-9e94-764a3e865e79"),
            JobID: unifiedto.String("quis"),
            Location: unifiedto.String("eum"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2021-03-22"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-11-26"),
            UserIds: []string{
                "animi",
            },
        },
        ConnectionID: "nostrum",
        ID: "a9da660f-f57b-4faa-94f9-efc1b4512c10",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchAtsConnectionIDInterviewIDRequest](../../models/operations/patchatsconnectionidinterviewidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PatchAtsConnectionIDInterviewIDSecurity](../../models/operations/patchatsconnectionidinterviewidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDJobIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PatchAtsConnectionIDJobID(ctx, operations.PatchAtsConnectionIDJobIDRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedto.String("velit"),
                    Address2: unifiedto.String("aspernatur"),
                    City: unifiedto.String("Devenstad"),
                    Country: unifiedto.String("Taiwan"),
                    CountryCode: unifiedto.String("SA"),
                    PostalCode: unifiedto.String("94131"),
                    Region: unifiedto.String("cupiditate"),
                    RegionCode: unifiedto.String("provident"),
                },
            },
            ClosedAt: types.MustDateFromString("2020-10-06"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedto.String("hic"),
                    Frequency: shared.AtsCompensationFrequencyMonth.ToPointer(),
                    Max: unifiedto.Float64(525.08),
                    Min: unifiedto.Float64(9358.33),
                    Type: shared.AtsCompensationTypeStockOptions,
                },
            },
            CreatedAt: types.MustDateFromString("2020-04-29"),
            Departments: []string{
                "aliquid",
            },
            Description: unifiedto.String("porro"),
            EmploymentType: shared.AtsJobEmploymentTypeIntern.ToPointer(),
            HiringManagerIds: []string{
                "dolorem",
            },
            ID: unifiedto.String("2ca3aed0-1179-4963-92fd-e04771778ff6"),
            LanguageLocale: unifiedto.String("architecto"),
            Name: unifiedto.String("Brian Carroll"),
            PublicJobUrls: []string{
                "esse",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "ex",
            },
            Remote: unifiedto.Bool(false),
            Status: shared.AtsJobStatusPending.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-12-10"),
        },
        ConnectionID: "laborum",
        ID: "15db6a66-0659-4a1a-9eaa-b5851d6c645b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchAtsConnectionIDJobIDRequest](../../models/operations/patchatsconnectionidjobidrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PatchAtsConnectionIDJobIDSecurity](../../models/operations/patchatsconnectionidjobidsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDScorecardIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PatchAtsConnectionIDScorecardID(ctx, operations.PatchAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("voluptatem"),
            CandidateID: unifiedto.String("molestias"),
            CreatedAt: types.MustDateFromString("2022-03-15"),
            ID: unifiedto.String("1891baa0-fe1a-4de0-88e6-f8c5f350d8cd"),
            InterviewID: unifiedto.String("nam"),
            InterviewerID: unifiedto.String("ipsam"),
            JobID: unifiedto.String("culpa"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationDefinitelyNo.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-12-02"),
        },
        ConnectionID: "deleniti",
        ID: "14301042-1813-4d52-88ec-e7e253b66845",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchAtsConnectionIDScorecardIDRequest](../../models/operations/patchatsconnectionidscorecardidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PatchAtsConnectionIDScorecardIDSecurity](../../models/operations/patchatsconnectionidscorecardidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDApplicationSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PostAtsConnectionIDApplication(ctx, operations.PostAtsConnectionIDApplicationRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustDateFromString("2022-03-21"),
            CandidateID: unifiedto.String("autem"),
            CreatedAt: types.MustDateFromString("2021-11-01"),
            ID: unifiedto.String("e205e16d-eab3-4fec-9578-a64584273a84"),
            JobID: unifiedto.String("quasi"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustDateFromString("2021-04-19"),
            RejectedReason: unifiedto.String("dicta"),
            Source: unifiedto.String("nisi"),
            Status: shared.AtsApplicationStatusReviewing.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-12-14"),
        },
        ConnectionID: "cupiditate",
    }, operationSecurity)
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
| `request`                                                                                                              | [operations.PostAtsConnectionIDApplicationRequest](../../models/operations/postatsconnectionidapplicationrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PostAtsConnectionIDApplicationSecurity](../../models/operations/postatsconnectionidapplicationsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDCandidateSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PostAtsConnectionIDCandidate(ctx, operations.PostAtsConnectionIDCandidateRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("reiciendis"),
                Address2: unifiedto.String("soluta"),
                City: unifiedto.String("New Chaunceystead"),
                Country: unifiedto.String("Namibia"),
                CountryCode: unifiedto.String("CH"),
                PostalCode: unifiedto.String("68976"),
                Region: unifiedto.String("delectus"),
                RegionCode: unifiedto.String("minima"),
            },
            CompanyName: unifiedto.String("praesentium"),
            CreatedAt: types.MustDateFromString("2022-02-06"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Katheryn.Johns52@hotmail.com",
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("modi"),
            ID: unifiedto.String("be056013-f59d-4a75-ba59-ecfef66ef1ca"),
            ImageURL: unifiedto.String("laborum"),
            Name: unifiedto.String("Shannon Lind"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "magni",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "soluta",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedto.String("Miss"),
            UpdatedAt: types.MustDateFromString("2022-07-08"),
        },
        ConnectionID: "voluptate",
    }, operationSecurity)
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
| `request`                                                                                                          | [operations.PostAtsConnectionIDCandidateRequest](../../models/operations/postatsconnectionidcandidaterequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PostAtsConnectionIDCandidateSecurity](../../models/operations/postatsconnectionidcandidatesecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDInterviewSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PostAtsConnectionIDInterview(ctx, operations.PostAtsConnectionIDInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("sequi"),
            CandidateID: unifiedto.String("dignissimos"),
            CreatedAt: types.MustDateFromString("2022-03-22"),
            EndAt: types.MustDateFromString("2021-04-24"),
            ExternalEventXref: unifiedto.String("iure"),
            ID: unifiedto.String("2f64d1db-1f2c-4431-8661-e96349e1cf9e"),
            JobID: unifiedto.String("alias"),
            Location: unifiedto.String("nisi"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2022-04-10"),
            Status: shared.AtsInterviewStatusComplete.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-10-11"),
            UserIds: []string{
                "iusto",
            },
        },
        ConnectionID: "sit",
    }, operationSecurity)
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
| `request`                                                                                                          | [operations.PostAtsConnectionIDInterviewRequest](../../models/operations/postatsconnectionidinterviewrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PostAtsConnectionIDInterviewSecurity](../../models/operations/postatsconnectionidinterviewsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDJobSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PostAtsConnectionIDJob(ctx, operations.PostAtsConnectionIDJobRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedto.String("doloremque"),
                    Address2: unifiedto.String("consequatur"),
                    City: unifiedto.String("Walterview"),
                    Country: unifiedto.String("Paraguay"),
                    CountryCode: unifiedto.String("GY"),
                    PostalCode: unifiedto.String("85759-4368"),
                    Region: unifiedto.String("animi"),
                    RegionCode: unifiedto.String("impedit"),
                },
            },
            ClosedAt: types.MustDateFromString("2022-08-23"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedto.String("est"),
                    Frequency: shared.AtsCompensationFrequencyHour.ToPointer(),
                    Max: unifiedto.Float64(4568.85),
                    Min: unifiedto.Float64(2885.7),
                    Type: shared.AtsCompensationTypeSalary,
                },
            },
            CreatedAt: types.MustDateFromString("2022-04-20"),
            Departments: []string{
                "vitae",
            },
            Description: unifiedto.String("inventore"),
            EmploymentType: shared.AtsJobEmploymentTypeContractor.ToPointer(),
            HiringManagerIds: []string{
                "ad",
            },
            ID: unifiedto.String("2965bb8a-7202-4611-835e-139dbc2259b1"),
            LanguageLocale: unifiedto.String("id"),
            Name: unifiedto.String("Laurence Nienow"),
            PublicJobUrls: []string{
                "sit",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "iusto",
            },
            Remote: unifiedto.Bool(false),
            Status: shared.AtsJobStatusArchived.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-10-05"),
        },
        ConnectionID: "aperiam",
    }, operationSecurity)
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
| `request`                                                                                              | [operations.PostAtsConnectionIDJobRequest](../../models/operations/postatsconnectionidjobrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.PostAtsConnectionIDJobSecurity](../../models/operations/postatsconnectionidjobsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDScorecardSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PostAtsConnectionIDScorecard(ctx, operations.PostAtsConnectionIDScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("totam"),
            CandidateID: unifiedto.String("dolore"),
            CreatedAt: types.MustDateFromString("2020-11-09"),
            ID: unifiedto.String("0672d1ad-879e-4eb9-a65b-85efbd02bae0"),
            InterviewID: unifiedto.String("expedita"),
            InterviewerID: unifiedto.String("officiis"),
            JobID: unifiedto.String("eos"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationStrongYes.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-06-28"),
        },
        ConnectionID: "odit",
    }, operationSecurity)
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
| `request`                                                                                                          | [operations.PostAtsConnectionIDScorecardRequest](../../models/operations/postatsconnectionidscorecardrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PostAtsConnectionIDScorecardSecurity](../../models/operations/postatsconnectionidscorecardsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDApplicationIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PutAtsConnectionIDApplicationID(ctx, operations.PutAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustDateFromString("2022-08-22"),
            CandidateID: unifiedto.String("error"),
            CreatedAt: types.MustDateFromString("2022-04-13"),
            ID: unifiedto.String("ea4b5197-f924-443d-a7ce-52b895c537c6"),
            JobID: unifiedto.String("modi"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustDateFromString("2022-09-14"),
            RejectedReason: unifiedto.String("voluptates"),
            Source: unifiedto.String("maiores"),
            Status: shared.AtsApplicationStatusOffered.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-04-10"),
        },
        ConnectionID: "ratione",
        ID: "4896c3ca-5acf-4be2-bd57-07577929177d",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsApplication != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutAtsConnectionIDApplicationIDRequest](../../models/operations/putatsconnectionidapplicationidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PutAtsConnectionIDApplicationIDSecurity](../../models/operations/putatsconnectionidapplicationidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PutAtsConnectionIDCandidateID(ctx, operations.PutAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("itaque"),
                Address2: unifiedto.String("similique"),
                City: unifiedto.String("Jacobsborough"),
                Country: unifiedto.String("Iraq"),
                CountryCode: unifiedto.String("TM"),
                PostalCode: unifiedto.String("63422-0581"),
                Region: unifiedto.String("repudiandae"),
                RegionCode: unifiedto.String("cum"),
            },
            CompanyName: unifiedto.String("dicta"),
            CreatedAt: types.MustDateFromString("2021-12-30"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Chanel16@yahoo.com",
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("nobis"),
            ID: unifiedto.String("07f116db-9954-45fc-95fa-88970e189dbb"),
            ImageURL: unifiedto.String("velit"),
            Name: unifiedto.String("Lucia Rutherford"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "adipisci",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "saepe",
                    Type: shared.AtsTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustDateFromString("2022-09-01"),
        },
        ConnectionID: "libero",
        ID: "197cd44e-2f52-4d82-9351-3bb6f48b656b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutAtsConnectionIDCandidateIDRequest](../../models/operations/putatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PutAtsConnectionIDCandidateIDSecurity](../../models/operations/putatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDInterviewIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PutAtsConnectionIDInterviewID(ctx, operations.PutAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("minus"),
            CandidateID: unifiedto.String("facere"),
            CreatedAt: types.MustDateFromString("2022-07-26"),
            EndAt: types.MustDateFromString("2022-01-10"),
            ExternalEventXref: unifiedto.String("voluptatibus"),
            ID: unifiedto.String("2e4b2753-7a8c-4d9e-b319-c177d525f77b"),
            JobID: unifiedto.String("illo"),
            Location: unifiedto.String("ab"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2022-02-14"),
            Status: shared.AtsInterviewStatusComplete.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-05-02"),
            UserIds: []string{
                "eos",
            },
        },
        ConnectionID: "reiciendis",
        ID: "f785fc37-814d-44c9-8e0c-2bb89eb75dad",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutAtsConnectionIDInterviewIDRequest](../../models/operations/putatsconnectionidinterviewidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PutAtsConnectionIDInterviewIDSecurity](../../models/operations/putatsconnectionidinterviewidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDJobIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PutAtsConnectionIDJobID(ctx, operations.PutAtsConnectionIDJobIDRequest{
        AtsJob: &shared.AtsJob{
            Addresses: []shared.AtsAddress{
                shared.AtsAddress{
                    Address1: unifiedto.String("iure"),
                    Address2: unifiedto.String("dolorem"),
                    City: unifiedto.String("Olenview"),
                    Country: unifiedto.String("Algeria"),
                    CountryCode: unifiedto.String("AO"),
                    PostalCode: unifiedto.String("02856"),
                    Region: unifiedto.String("cum"),
                    RegionCode: unifiedto.String("amet"),
                },
            },
            ClosedAt: types.MustDateFromString("2022-11-18"),
            Compensation: []shared.AtsCompensation{
                shared.AtsCompensation{
                    Currency: unifiedto.String("laudantium"),
                    Frequency: shared.AtsCompensationFrequencyOneTime.ToPointer(),
                    Max: unifiedto.Float64(9384.12),
                    Min: unifiedto.Float64(4797.07),
                    Type: shared.AtsCompensationTypeBonus,
                },
            },
            CreatedAt: types.MustDateFromString("2021-08-24"),
            Departments: []string{
                "necessitatibus",
            },
            Description: unifiedto.String("provident"),
            EmploymentType: shared.AtsJobEmploymentTypeOther.ToPointer(),
            HiringManagerIds: []string{
                "consequatur",
            },
            ID: unifiedto.String("57eb809e-2810-4331-b398-1d4c700b607f"),
            LanguageLocale: unifiedto.String("velit"),
            Name: unifiedto.String("Rene Feeney"),
            PublicJobUrls: []string{
                "consectetur",
            },
            Raw: &shared.PropertyAtsJobRaw{},
            RecruiterIds: []string{
                "soluta",
            },
            Remote: unifiedto.Bool(false),
            Status: shared.AtsJobStatusOpen.ToPointer(),
            UpdatedAt: types.MustDateFromString("2021-01-30"),
        },
        ConnectionID: "amet",
        ID: "f2ceda7e-23f2-4257-811f-af4b7544e472",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutAtsConnectionIDJobIDRequest](../../models/operations/putatsconnectionidjobidrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PutAtsConnectionIDJobIDSecurity](../../models/operations/putatsconnectionidjobidsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDScorecardIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Ats.PutAtsConnectionIDScorecardID(ctx, operations.PutAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("accusamus"),
            CandidateID: unifiedto.String("rem"),
            CreatedAt: types.MustDateFromString("2022-11-10"),
            ID: unifiedto.String("857a5b40-463a-47d5-b5f1-400e764ad733"),
            InterviewID: unifiedto.String("quaerat"),
            InterviewerID: unifiedto.String("itaque"),
            JobID: unifiedto.String("minus"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationDefinitelyNo.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-01-14"),
        },
        ConnectionID: "quas",
        ID: "1b36a080-88d1-400e-bada-200ef0422eb2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutAtsConnectionIDScorecardIDRequest](../../models/operations/putatsconnectionidscorecardidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PutAtsConnectionIDScorecardIDSecurity](../../models/operations/putatsconnectionidscorecardidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.PutAtsConnectionIDScorecardIDResponse](../../models/operations/putatsconnectionidscorecardidresponse.md), error**

