# Application
(*Application*)

### Available Operations

* [DeleteAtsConnectionIDApplicationID](#deleteatsconnectionidapplicationid) - Remove an application
* [GetAtsConnectionIDApplication](#getatsconnectionidapplication) - List all applications
* [GetAtsConnectionIDApplicationID](#getatsconnectionidapplicationid) - Retrieve an application
* [PatchAtsConnectionIDApplicationID](#patchatsconnectionidapplicationid) - Update an application
* [PostAtsConnectionIDApplication](#postatsconnectionidapplication) - Create an application
* [PutAtsConnectionIDApplicationID](#putatsconnectionidapplicationid) - Update an application

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
    res, err := s.Application.DeleteAtsConnectionIDApplicationID(ctx, operations.DeleteAtsConnectionIDApplicationIDRequest{
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
    res, err := s.Application.GetAtsConnectionIDApplication(ctx, operations.GetAtsConnectionIDApplicationRequest{
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
    res, err := s.Application.GetAtsConnectionIDApplicationID(ctx, operations.GetAtsConnectionIDApplicationIDRequest{
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
    res, err := s.Application.PatchAtsConnectionIDApplicationID(ctx, operations.PatchAtsConnectionIDApplicationIDRequest{
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
    res, err := s.Application.PostAtsConnectionIDApplication(ctx, operations.PostAtsConnectionIDApplicationRequest{
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
    res, err := s.Application.PutAtsConnectionIDApplicationID(ctx, operations.PutAtsConnectionIDApplicationIDRequest{
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

