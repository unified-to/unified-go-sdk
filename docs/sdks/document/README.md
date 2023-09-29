# Document
(*Document*)

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
    res, err := s.Document.DeleteAtsConnectionIDScorecardID(ctx, operations.DeleteAtsConnectionIDScorecardIDRequest{
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
    res, err := s.Document.GetAtsConnectionIDScorecard(ctx, operations.GetAtsConnectionIDScorecardRequest{
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
    res, err := s.Document.GetAtsConnectionIDScorecardID(ctx, operations.GetAtsConnectionIDScorecardIDRequest{
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
    res, err := s.Document.PatchAtsConnectionIDScorecardID(ctx, operations.PatchAtsConnectionIDScorecardIDRequest{
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
    res, err := s.Document.PostAtsConnectionIDScorecard(ctx, operations.PostAtsConnectionIDScorecardRequest{
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
    res, err := s.Document.PutAtsConnectionIDScorecardID(ctx, operations.PutAtsConnectionIDScorecardIDRequest{
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

