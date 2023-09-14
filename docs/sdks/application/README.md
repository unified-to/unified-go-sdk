# Application

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDApplicationIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Application.DeleteAtsConnectionIDApplicationID(ctx, operations.DeleteAtsConnectionIDApplicationIDRequest{
        ConnectionID: "rerum",
        ID: "3fe49a8d-9cbf-4486-b332-3f9b77f3a410",
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
    res, err := s.Application.GetAtsConnectionIDApplication(ctx, operations.GetAtsConnectionIDApplicationRequest{
        CandidateID: unifiedto.String("ipsa"),
        ConnectionID: "iure",
        JobID: unifiedto.String("odio"),
        Limit: unifiedto.Float64(3117.96),
        Offset: unifiedto.Float64(8810.05),
        Order: unifiedto.String("quidem"),
        Query: unifiedto.String("voluptatibus"),
        Sort: unifiedto.String("voluptas"),
        UpdatedGte: types.MustTimeFromString("2022-08-22T21:20:36.034Z"),
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
    res, err := s.Application.GetAtsConnectionIDApplicationID(ctx, operations.GetAtsConnectionIDApplicationIDRequest{
        ConnectionID: "atque",
        ID: "0d1ba77a-89eb-4f73-bae4-203ce5e6a95d",
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
    res, err := s.Application.PatchAtsConnectionIDApplicationID(ctx, operations.PatchAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2021-09-28T20:14:16.452Z"),
            CandidateID: unifiedto.String("alias"),
            CreatedAt: types.MustTimeFromString("2022-01-24T12:18:47.570Z"),
            ID: unifiedto.String("46ce2af7-a73c-4f3b-a453-f870b326b5a7"),
            JobID: unifiedto.String("ipsum"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2022-10-24T22:37:32.805Z"),
            RejectedReason: unifiedto.String("cupiditate"),
            Source: unifiedto.String("maxime"),
            Status: shared.AtsApplicationStatusHired.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-10-07T04:50:29.805Z"),
        },
        ConnectionID: "laborum",
        ID: "8422bb67-9d23-4227-95bf-0cbb1e31b8b9",
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
    res, err := s.Application.PostAtsConnectionIDApplication(ctx, operations.PostAtsConnectionIDApplicationRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2022-01-14T21:26:14.171Z"),
            CandidateID: unifiedto.String("dolorem"),
            CreatedAt: types.MustTimeFromString("2022-09-18T06:37:26.413Z"),
            ID: unifiedto.String("3a1108e0-adcf-44b9-a187-9fce953f73ef"),
            JobID: unifiedto.String("dignissimos"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2020-11-08T11:03:10.206Z"),
            RejectedReason: unifiedto.String("quod"),
            Source: unifiedto.String("odio"),
            Status: shared.AtsApplicationStatusBackgroundCheck.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-04-02T18:28:29.036Z"),
        },
        ConnectionID: "ducimus",
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
    res, err := s.Application.PutAtsConnectionIDApplicationID(ctx, operations.PutAtsConnectionIDApplicationIDRequest{
        AtsApplication: &shared.AtsApplication{
            AppliedAt: types.MustTimeFromString("2022-02-26T17:44:28.591Z"),
            CandidateID: unifiedto.String("illum"),
            CreatedAt: types.MustTimeFromString("2022-05-20T11:24:00.260Z"),
            ID: unifiedto.String("c0f5d2cf-f7c7-40a4-9626-d436813f16d9"),
            JobID: unifiedto.String("voluptatibus"),
            Raw: &shared.PropertyAtsApplicationRaw{},
            RejectedAt: types.MustTimeFromString("2022-01-15T07:05:18.296Z"),
            RejectedReason: unifiedto.String("quisquam"),
            Source: unifiedto.String("saepe"),
            Status: shared.AtsApplicationStatusFirstInterview.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-12-03T14:20:49.127Z"),
        },
        ConnectionID: "veniam",
        ID: "6146c3e2-50fb-4008-842e-141aac366c8d",
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

