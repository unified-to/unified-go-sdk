# Document

### Available Operations

* [DeleteAtsConnectionIDScorecardID](#deleteatsconnectionidscorecardid) - Remove a scorecard
* [GetAtsConnectionIDScorecard](#getatsconnectionidscorecard) - List all scorecards
* [GetAtsConnectionIDScorecardID](#getatsconnectionidscorecardid) - Retrieve a scorecard
* [PatchAtsConnectionIDScorecardID](#patchatsconnectionidscorecardid) - Update a scorecard
* [PostAtsConnectionIDScorecard](#postatsconnectionidscorecard) - Create a scorecard
* [PutAtsConnectionIDScorecardID](#putatsconnectionidscorecardid) - Update a scorecard

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
    res, err := s.Document.DeleteAtsConnectionIDScorecardID(ctx, operations.DeleteAtsConnectionIDScorecardIDRequest{
        ConnectionID: "facere",
        ID: "9c337473-082b-494f-aab1-fd5671e9c326",
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
    res, err := s.Document.GetAtsConnectionIDScorecard(ctx, operations.GetAtsConnectionIDScorecardRequest{
        ApplicationID: unifiedto.String("neque"),
        CandidateID: unifiedto.String("enim"),
        ConnectionID: "eaque",
        InterviewID: unifiedto.String("officia"),
        Limit: unifiedto.Float64(2702.53),
        Offset: unifiedto.Float64(4310.35),
        Order: unifiedto.String("molestiae"),
        Query: unifiedto.String("architecto"),
        Sort: unifiedto.String("aliquam"),
        UpdatedGte: types.MustTimeFromString("2022-07-12T22:54:11.511Z"),
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
    res, err := s.Document.GetAtsConnectionIDScorecardID(ctx, operations.GetAtsConnectionIDScorecardIDRequest{
        ConnectionID: "blanditiis",
        ID: "9ce0e991-594d-493a-b4c0-252fe3b4b4db",
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
    res, err := s.Document.PatchAtsConnectionIDScorecardID(ctx, operations.PatchAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("atque"),
            CandidateID: unifiedto.String("tempore"),
            CreatedAt: types.MustTimeFromString("2022-07-22T19:46:16.313Z"),
            ID: unifiedto.String("8ebb6e1d-2cf5-402b-afb2-cbc4635d5e65"),
            InterviewID: unifiedto.String("at"),
            InterviewerID: unifiedto.String("culpa"),
            JobID: unifiedto.String("alias"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationDefinitelyNo.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-06-03T11:24:52.324Z"),
        },
        ConnectionID: "dolor",
        ID: "e951a1e3-0fda-4966-889d-7b78673e13a1",
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
    res, err := s.Document.PostAtsConnectionIDScorecard(ctx, operations.PostAtsConnectionIDScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("eos"),
            CandidateID: unifiedto.String("dolorum"),
            CreatedAt: types.MustTimeFromString("2022-04-19T16:37:31.203Z"),
            ID: unifiedto.String("99249459-4487-4f5c-8438-36b86b3cdf64"),
            InterviewID: unifiedto.String("dicta"),
            InterviewerID: unifiedto.String("minima"),
            JobID: unifiedto.String("facilis"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationDefinitelyNo.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-09-13T14:57:39.091Z"),
        },
        ConnectionID: "molestias",
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
    res, err := s.Document.PutAtsConnectionIDScorecardID(ctx, operations.PutAtsConnectionIDScorecardIDRequest{
        AtsScorecard: &shared.AtsScorecard{
            ApplicationID: unifiedto.String("hic"),
            CandidateID: unifiedto.String("error"),
            CreatedAt: types.MustTimeFromString("2020-02-11T23:32:43.703Z"),
            ID: unifiedto.String("13f4eedb-e78b-4f60-a825-894ea763d5c7"),
            InterviewID: unifiedto.String("fugit"),
            InterviewerID: unifiedto.String("voluptate"),
            JobID: unifiedto.String("provident"),
            Raw: shared.PropertyAtsScorecardRaw{},
            Recommendation: shared.AtsScorecardRecommendationNo.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-01-28T06:25:58.770Z"),
        },
        ConnectionID: "laudantium",
        ID: "5148d6d5-49e5-4635-b33b-c0f970c42fc9",
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

