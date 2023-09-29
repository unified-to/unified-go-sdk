# Interview
(*Interview*)

### Available Operations

* [DeleteAtsConnectionIDInterviewID](#deleteatsconnectionidinterviewid) - Remove a interview
* [GetAtsConnectionIDInterview](#getatsconnectionidinterview) - List all interviews
* [GetAtsConnectionIDInterviewID](#getatsconnectionidinterviewid) - Retrieve a interview
* [PatchAtsConnectionIDInterviewID](#patchatsconnectionidinterviewid) - Update a interview
* [PostAtsConnectionIDInterview](#postatsconnectionidinterview) - Create a interview
* [PutAtsConnectionIDInterviewID](#putatsconnectionidinterviewid) - Update a interview

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
    res, err := s.Interview.DeleteAtsConnectionIDInterviewID(ctx, operations.DeleteAtsConnectionIDInterviewIDRequest{
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
    res, err := s.Interview.GetAtsConnectionIDInterview(ctx, operations.GetAtsConnectionIDInterviewRequest{
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
    res, err := s.Interview.GetAtsConnectionIDInterviewID(ctx, operations.GetAtsConnectionIDInterviewIDRequest{
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
    res, err := s.Interview.PatchAtsConnectionIDInterviewID(ctx, operations.PatchAtsConnectionIDInterviewIDRequest{
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
    res, err := s.Interview.PostAtsConnectionIDInterview(ctx, operations.PostAtsConnectionIDInterviewRequest{
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
    res, err := s.Interview.PutAtsConnectionIDInterviewID(ctx, operations.PutAtsConnectionIDInterviewIDRequest{
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

