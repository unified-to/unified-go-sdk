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
    res, err := s.File.DeleteCrmConnectionIDFileID(ctx, operations.DeleteCrmConnectionIDFileIDRequest{
        ConnectionID: "adipisci",
        ID: "febdf676-b720-46da-b750-052a5647edc4",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDFileIDRequest](../../models/operations/deletecrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


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
    res, err := s.File.GetCrmConnectionIDFile(ctx, operations.GetCrmConnectionIDFileRequest{
        CompanyID: unifiedgosdk.String("sequi"),
        ConnectionID: "natus",
        ContactID: unifiedgosdk.String("saepe"),
        DealID: unifiedgosdk.String("quibusdam"),
        Limit: unifiedgosdk.Float64(5481.43),
        Offset: unifiedgosdk.Float64(8071.51),
        Order: unifiedgosdk.String("aliquam"),
        Query: unifiedgosdk.String("adipisci"),
        Sort: unifiedgosdk.String("explicabo"),
        UpdatedGte: types.MustTimeFromString("2022-01-22T06:38:09.253Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDFileRequest](../../models/operations/getcrmconnectionidfilerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


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
    res, err := s.File.GetCrmConnectionIDFileID(ctx, operations.GetCrmConnectionIDFileIDRequest{
        ConnectionID: "incidunt",
        ID: "1240d448-7ac6-493b-94c3-b9d2488d795a",
    })
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
| `request`                                                                                                | [operations.GetCrmConnectionIDFileIDRequest](../../models/operations/getcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


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
    res, err := s.File.PatchCrmConnectionIDFileID(ctx, operations.PatchCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("fuga"),
            CompanyID: unifiedgosdk.String("incidunt"),
            ContactID: unifiedgosdk.String("aspernatur"),
            CreatedAt: types.MustTimeFromString("2020-08-02T08:40:50.776Z"),
            DealID: unifiedgosdk.String("dolore"),
            Description: unifiedgosdk.String("accusantium"),
            FileName: unifiedgosdk.String("corporis"),
            FileSize: unifiedgosdk.Float64(3881.8),
            FileType: unifiedgosdk.String("laboriosam"),
            FileURL: unifiedgosdk.String("omnis"),
            ID: unifiedgosdk.String("f69a006d-2124-4945-8819-d7c3b1b41844"),
            LeadID: unifiedgosdk.String("consequatur"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2022-12-09T21:50:08.252Z"),
            UserID: unifiedgosdk.String("saepe"),
        },
        ConnectionID: "accusantium",
        ID: "0310d023-dc90-41f5-afd2-a6c44846ae9d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDFileIDRequest](../../models/operations/patchcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


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
    res, err := s.File.PostCrmConnectionIDFile(ctx, operations.PostCrmConnectionIDFileRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("praesentium"),
            CompanyID: unifiedgosdk.String("occaecati"),
            ContactID: unifiedgosdk.String("eos"),
            CreatedAt: types.MustTimeFromString("2022-10-08T19:31:07.425Z"),
            DealID: unifiedgosdk.String("nobis"),
            Description: unifiedgosdk.String("quos"),
            FileName: unifiedgosdk.String("provident"),
            FileSize: unifiedgosdk.Float64(4099.18),
            FileType: unifiedgosdk.String("consequuntur"),
            FileURL: unifiedgosdk.String("delectus"),
            ID: unifiedgosdk.String("4896bf51-e465-42d3-8343-d61778af4912"),
            LeadID: unifiedgosdk.String("numquam"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2022-07-08T10:09:32.871Z"),
            UserID: unifiedgosdk.String("magni"),
        },
        ConnectionID: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDFileRequest](../../models/operations/postcrmconnectionidfilerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


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
    res, err := s.File.PutCrmConnectionIDFileID(ctx, operations.PutCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("eveniet"),
            CompanyID: unifiedgosdk.String("commodi"),
            ContactID: unifiedgosdk.String("magni"),
            CreatedAt: types.MustTimeFromString("2022-05-23T03:31:28.636Z"),
            DealID: unifiedgosdk.String("aut"),
            Description: unifiedgosdk.String("occaecati"),
            FileName: unifiedgosdk.String("vero"),
            FileSize: unifiedgosdk.Float64(6231.5),
            FileType: unifiedgosdk.String("inventore"),
            FileURL: unifiedgosdk.String("ipsa"),
            ID: unifiedgosdk.String("44a5de59-ac77-4066-b0cf-1cf593260525"),
            LeadID: unifiedgosdk.String("beatae"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2021-10-17T02:29:52.101Z"),
            UserID: unifiedgosdk.String("ex"),
        },
        ConnectionID: "harum",
        ID: "b426897d-99a2-4d33-9670-e93ee6cf59f3",
    })
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
| `request`                                                                                                | [operations.PutCrmConnectionIDFileIDRequest](../../models/operations/putcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDFileIDResponse](../../models/operations/putcrmconnectionidfileidresponse.md), error**

