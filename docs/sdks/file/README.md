# File

### Available Operations

* [DeleteCrmConnectionIDFileID](#deletecrmconnectionidfileid) - Remove a file
* [GetCrmConnectionIDFile](#getcrmconnectionidfile) - List all files
* [GetCrmConnectionIDFileID](#getcrmconnectionidfileid) - Retrieve a file
* [PatchCrmConnectionIDFileID](#patchcrmconnectionidfileid) - Update a file
* [PostCrmConnectionIDFile](#postcrmconnectionidfile) - Create a file
* [PutCrmConnectionIDFileID](#putcrmconnectionidfileid) - Update a file

## DeleteCrmConnectionIDFileID

Remove a file

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
    operationSecurity := operations.DeleteCrmConnectionIDFileIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.DeleteCrmConnectionIDFileID(ctx, operations.DeleteCrmConnectionIDFileIDRequest{
        ConnectionID: "adipisci",
        ID: "febdf676-b720-46da-b750-052a5647edc4",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDFileIDRequest](../../models/operations/deletecrmconnectionidfileidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.DeleteCrmConnectionIDFileIDSecurity](../../models/operations/deletecrmconnectionidfileidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.DeleteCrmConnectionIDFileIDResponse](../../models/operations/deletecrmconnectionidfileidresponse.md), error**


## GetCrmConnectionIDFile

List all files

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
    operationSecurity := operations.GetCrmConnectionIDFileSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.GetCrmConnectionIDFile(ctx, operations.GetCrmConnectionIDFileRequest{
        CompanyID: unifiedto.String("sequi"),
        ConnectionID: "natus",
        ContactID: unifiedto.String("saepe"),
        DealID: unifiedto.String("quibusdam"),
        Limit: unifiedto.Float64(5481.43),
        Offset: unifiedto.Float64(8071.51),
        Order: unifiedto.String("aliquam"),
        Query: unifiedto.String("adipisci"),
        Sort: unifiedto.String("explicabo"),
        UpdatedGte: types.MustTimeFromString("2022-01-22T06:38:09.253Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDFileRequest](../../models/operations/getcrmconnectionidfilerequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCrmConnectionIDFileSecurity](../../models/operations/getcrmconnectionidfilesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetCrmConnectionIDFileResponse](../../models/operations/getcrmconnectionidfileresponse.md), error**


## GetCrmConnectionIDFileID

Retrieve a file

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
    operationSecurity := operations.GetCrmConnectionIDFileIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.GetCrmConnectionIDFileID(ctx, operations.GetCrmConnectionIDFileIDRequest{
        ConnectionID: "incidunt",
        ID: "1240d448-7ac6-493b-94c3-b9d2488d795a",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDFileIDRequest](../../models/operations/getcrmconnectionidfileidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetCrmConnectionIDFileIDSecurity](../../models/operations/getcrmconnectionidfileidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.GetCrmConnectionIDFileIDResponse](../../models/operations/getcrmconnectionidfileidresponse.md), error**


## PatchCrmConnectionIDFileID

Update a file

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
    operationSecurity := operations.PatchCrmConnectionIDFileIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.PatchCrmConnectionIDFileID(ctx, operations.PatchCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedto.Bool(false),
            ActivityID: unifiedto.String("fuga"),
            CompanyID: unifiedto.String("incidunt"),
            ContactID: unifiedto.String("aspernatur"),
            CreatedAt: types.MustTimeFromString("2020-08-02T08:40:50.776Z"),
            DealID: unifiedto.String("dolore"),
            Description: unifiedto.String("accusantium"),
            FileName: unifiedto.String("corporis"),
            FileSize: unifiedto.Float64(3881.8),
            FileType: unifiedto.String("laboriosam"),
            FileURL: unifiedto.String("omnis"),
            ID: unifiedto.String("f69a006d-2124-4945-8819-d7c3b1b41844"),
            LeadID: unifiedto.String("consequatur"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2022-12-09T21:50:08.252Z"),
            UserID: unifiedto.String("saepe"),
        },
        ConnectionID: "accusantium",
        ID: "0310d023-dc90-41f5-afd2-a6c44846ae9d",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDFileIDRequest](../../models/operations/patchcrmconnectionidfileidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PatchCrmConnectionIDFileIDSecurity](../../models/operations/patchcrmconnectionidfileidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PatchCrmConnectionIDFileIDResponse](../../models/operations/patchcrmconnectionidfileidresponse.md), error**


## PostCrmConnectionIDFile

Create a file

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
    operationSecurity := operations.PostCrmConnectionIDFileSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.PostCrmConnectionIDFile(ctx, operations.PostCrmConnectionIDFileRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedto.Bool(false),
            ActivityID: unifiedto.String("praesentium"),
            CompanyID: unifiedto.String("occaecati"),
            ContactID: unifiedto.String("eos"),
            CreatedAt: types.MustTimeFromString("2022-10-08T19:31:07.425Z"),
            DealID: unifiedto.String("nobis"),
            Description: unifiedto.String("quos"),
            FileName: unifiedto.String("provident"),
            FileSize: unifiedto.Float64(4099.18),
            FileType: unifiedto.String("consequuntur"),
            FileURL: unifiedto.String("delectus"),
            ID: unifiedto.String("4896bf51-e465-42d3-8343-d61778af4912"),
            LeadID: unifiedto.String("numquam"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2022-07-08T10:09:32.871Z"),
            UserID: unifiedto.String("magni"),
        },
        ConnectionID: "enim",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostCrmConnectionIDFileRequest](../../models/operations/postcrmconnectionidfilerequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PostCrmConnectionIDFileSecurity](../../models/operations/postcrmconnectionidfilesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.PostCrmConnectionIDFileResponse](../../models/operations/postcrmconnectionidfileresponse.md), error**


## PutCrmConnectionIDFileID

Update a file

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
    operationSecurity := operations.PutCrmConnectionIDFileIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.File.PutCrmConnectionIDFileID(ctx, operations.PutCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedto.Bool(false),
            ActivityID: unifiedto.String("eveniet"),
            CompanyID: unifiedto.String("commodi"),
            ContactID: unifiedto.String("magni"),
            CreatedAt: types.MustTimeFromString("2022-05-23T03:31:28.636Z"),
            DealID: unifiedto.String("aut"),
            Description: unifiedto.String("occaecati"),
            FileName: unifiedto.String("vero"),
            FileSize: unifiedto.Float64(6231.5),
            FileType: unifiedto.String("inventore"),
            FileURL: unifiedto.String("ipsa"),
            ID: unifiedto.String("44a5de59-ac77-4066-b0cf-1cf593260525"),
            LeadID: unifiedto.String("beatae"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2021-10-17T02:29:52.101Z"),
            UserID: unifiedto.String("ex"),
        },
        ConnectionID: "harum",
        ID: "b426897d-99a2-4d33-9670-e93ee6cf59f3",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDFileIDRequest](../../models/operations/putcrmconnectionidfileidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PutCrmConnectionIDFileIDSecurity](../../models/operations/putcrmconnectionidfileidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PutCrmConnectionIDFileIDResponse](../../models/operations/putcrmconnectionidfileidresponse.md), error**

