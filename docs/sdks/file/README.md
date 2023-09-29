# File
(*File*)

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
        ConnectionID: "Bicycle",
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
        CompanyID: unifiedgosdk.String("reboot"),
        ConnectionID: "customise far",
        ContactID: unifiedgosdk.String("Electronic proactive"),
        DealID: unifiedgosdk.String("withdrawal deposit Gloves"),
        Limit: unifiedgosdk.Float64(1588.79),
        Offset: unifiedgosdk.Float64(3754.81),
        Order: unifiedgosdk.String("Implemented fairly meh"),
        Query: unifiedgosdk.String("FTP Producer"),
        Sort: unifiedgosdk.String("soprano deliverables"),
        UpdatedGte: types.MustTimeFromString("2022-03-02T03:00:09.711Z"),
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
        ConnectionID: "Bicycle",
        ID: "<ID>",
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
            ActivityID: unifiedgosdk.String("duh Handmade harness"),
            CompanyID: unifiedgosdk.String("CFP"),
            ContactID: unifiedgosdk.String("unaware yellow generating"),
            CreatedAt: types.MustTimeFromString("2021-05-04T04:54:33.785Z"),
            DealID: unifiedgosdk.String("channels SUV"),
            Description: unifiedgosdk.String("De-engineered didactic hardware"),
            FileName: unifiedgosdk.String("metical_silver_yellow.html"),
            FileSize: unifiedgosdk.Float64(6861.53),
            FileType: unifiedgosdk.String("video"),
            FileURL: unifiedgosdk.String("navigate Funk"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("internal"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-02-21T13:46:42.012Z"),
            UserID: unifiedgosdk.String("Interactions"),
        },
        ConnectionID: "Handcrafted",
        ID: "<ID>",
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
            ActivityID: unifiedgosdk.String("tan impedit Pickup"),
            CompanyID: unifiedgosdk.String("Manager"),
            ContactID: unifiedgosdk.String("Florida Shoes East"),
            CreatedAt: types.MustTimeFromString("2023-01-08T11:37:24.708Z"),
            DealID: unifiedgosdk.String("Agent"),
            Description: unifiedgosdk.String("Multi-lateral well-modulated portal"),
            FileName: unifiedgosdk.String("panel_city.wav"),
            FileSize: unifiedgosdk.Float64(1401.73),
            FileType: unifiedgosdk.String("application"),
            FileURL: unifiedgosdk.String("for Chips under"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("abaft Checking"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-02-25T09:46:59.608Z"),
            UserID: unifiedgosdk.String("Mexico withdrawal"),
        },
        ConnectionID: "national Lead",
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
            ActivityID: unifiedgosdk.String("Cotton"),
            CompanyID: unifiedgosdk.String("Northeast"),
            ContactID: unifiedgosdk.String("Computer"),
            CreatedAt: types.MustTimeFromString("2021-04-09T13:10:01.367Z"),
            DealID: unifiedgosdk.String("toward confiscate East"),
            Description: unifiedgosdk.String("Devolved upward-trending matrices"),
            FileName: unifiedgosdk.String("generation_tactics.wav"),
            FileSize: unifiedgosdk.Float64(4770.09),
            FileType: unifiedgosdk.String("audio"),
            FileURL: unifiedgosdk.String("framework azure Metal"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("ampere costume"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-05-15T05:04:24.130Z"),
            UserID: unifiedgosdk.String("Research payment"),
        },
        ConnectionID: "East Associate Mazda",
        ID: "<ID>",
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

