# Interview

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDInterviewIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Interview.DeleteAtsConnectionIDInterviewID(ctx, operations.DeleteAtsConnectionIDInterviewIDRequest{
        ConnectionID: "minus",
        ID: "7dda70ec-60e6-4075-894d-61c14cd90227",
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
    res, err := s.Interview.GetAtsConnectionIDInterview(ctx, operations.GetAtsConnectionIDInterviewRequest{
        ApplicationID: unifiedto.String("recusandae"),
        ConnectionID: "neque",
        Limit: unifiedto.Float64(4704),
        Offset: unifiedto.Float64(7513.92),
        Order: unifiedto.String("eaque"),
        Query: unifiedto.String("facere"),
        Sort: unifiedto.String("iste"),
        UpdatedGte: types.MustDateFromString("2022-07-18"),
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
    res, err := s.Interview.GetAtsConnectionIDInterviewID(ctx, operations.GetAtsConnectionIDInterviewIDRequest{
        ConnectionID: "reiciendis",
        ID: "1a5491ab-e975-41b1-86d2-3e03e69815aa",
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
    res, err := s.Interview.PatchAtsConnectionIDInterviewID(ctx, operations.PatchAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("officiis"),
            CandidateID: unifiedto.String("omnis"),
            CreatedAt: types.MustDateFromString("2021-02-14"),
            EndAt: types.MustDateFromString("2020-07-04"),
            ExternalEventXref: unifiedto.String("necessitatibus"),
            ID: unifiedto.String("9e729c9d-4f2d-48a4-8640-ca60db73a2f9"),
            JobID: unifiedto.String("dolorem"),
            Location: unifiedto.String("maiores"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2022-07-30"),
            Status: shared.AtsInterviewStatusAwaitingFeedback.ToPointer(),
            UpdatedAt: types.MustDateFromString("2020-08-23"),
            UserIds: []string{
                "quae",
            },
        },
        ConnectionID: "possimus",
        ID: "8da56122-026a-4b8f-a774-85c1976af980",
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
    res, err := s.Interview.PostAtsConnectionIDInterview(ctx, operations.PostAtsConnectionIDInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("nulla"),
            CandidateID: unifiedto.String("culpa"),
            CreatedAt: types.MustDateFromString("2022-05-03"),
            EndAt: types.MustDateFromString("2022-07-01"),
            ExternalEventXref: unifiedto.String("unde"),
            ID: unifiedto.String("fc44db27-4530-4e5c-87c6-d0cbcfdcd334"),
            JobID: unifiedto.String("facilis"),
            Location: unifiedto.String("autem"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2021-10-11"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-04-07"),
            UserIds: []string{
                "minus",
            },
        },
        ConnectionID: "repudiandae",
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
    res, err := s.Interview.PutAtsConnectionIDInterviewID(ctx, operations.PutAtsConnectionIDInterviewIDRequest{
        AtsInterview: &shared.AtsInterview{
            ApplicationID: unifiedto.String("quisquam"),
            CandidateID: unifiedto.String("mollitia"),
            CreatedAt: types.MustDateFromString("2022-04-26"),
            EndAt: types.MustDateFromString("2022-04-30"),
            ExternalEventXref: unifiedto.String("voluptates"),
            ID: unifiedto.String("e5e0da8b-9af6-4ad0-9486-e7b413cbe2d1"),
            JobID: unifiedto.String("molestiae"),
            Location: unifiedto.String("ea"),
            Raw: &shared.PropertyAtsInterviewRaw{},
            StartAt: types.MustDateFromString("2020-07-27"),
            Status: shared.AtsInterviewStatusScheduled.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-03-28"),
            UserIds: []string{
                "adipisci",
            },
        },
        ConnectionID: "pariatur",
        ID: "40f61d17-1157-4cbe-9ee4-f7211840772f",
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

