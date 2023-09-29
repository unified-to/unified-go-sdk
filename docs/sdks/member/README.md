# Member
(*Member*)

### Available Operations

* [DeleteMartechConnectionIDListIDMemberID](#deletemartechconnectionidlistidmemberid) - Remove member from a list
* [GetMartechConnectionIDListIDMember](#getmartechconnectionidlistidmember) - List all members in a list
* [GetMartechConnectionIDListIDMemberID](#getmartechconnectionidlistidmemberid) - Retrieve a member from a list
* [PatchMartechConnectionIDListIDMemberID](#patchmartechconnectionidlistidmemberid) - Update a member in a list
* [PostMartechConnectionIDListIDMember](#postmartechconnectionidlistidmember) - Create a member in a list
* [PutMartechConnectionIDListIDMemberID](#putmartechconnectionidlistidmemberid) - Update a member in a list

## DeleteMartechConnectionIDListIDMemberID

Remove member from a list

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
    res, err := s.Member.DeleteMartechConnectionIDListIDMemberID(ctx, operations.DeleteMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "Southwest fib",
        ID: "<ID>",
        ListID: "pascal",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.DeleteMartechConnectionIDListIDMemberIDRequest](../../models/operations/deletemartechconnectionidlistidmemberidrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.DeleteMartechConnectionIDListIDMemberIDResponse](../../models/operations/deletemartechconnectionidlistidmemberidresponse.md), error**


## GetMartechConnectionIDListIDMember

List all members in a list

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
    res, err := s.Member.GetMartechConnectionIDListIDMember(ctx, operations.GetMartechConnectionIDListIDMemberRequest{
        ConnectionID: "fuchsia economics",
        Limit: unifiedgosdk.Float64(3725.92),
        ListID: "Southwest",
        Offset: unifiedgosdk.Float64(1114.27),
        Order: unifiedgosdk.String("emulation"),
        Query: unifiedgosdk.String("male male"),
        Sort: unifiedgosdk.String("Arizona Oklahoma Land"),
        UpdatedGte: types.MustTimeFromString("2021-03-19T13:12:48.332Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetMartechConnectionIDListIDMemberRequest](../../models/operations/getmartechconnectionidlistidmemberrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetMartechConnectionIDListIDMemberResponse](../../models/operations/getmartechconnectionidlistidmemberresponse.md), error**


## GetMartechConnectionIDListIDMemberID

Retrieve a member from a list

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
    res, err := s.Member.GetMartechConnectionIDListIDMemberID(ctx, operations.GetMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "male",
        ID: "<ID>",
        ListID: "Gasoline Home allot",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetMartechConnectionIDListIDMemberIDRequest](../../models/operations/getmartechconnectionidlistidmemberidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetMartechConnectionIDListIDMemberIDResponse](../../models/operations/getmartechconnectionidlistidmemberidresponse.md), error**


## PatchMartechConnectionIDListIDMemberID

Update a member in a list

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
    res, err := s.Member.PatchMartechConnectionIDListIDMemberID(ctx, operations.PatchMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2022-06-21T07:15:04.418Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Zula_Bogan76@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ListIds: []string{
                "gadzooks",
            },
            Name: unifiedgosdk.String("Haven Hatchback"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "mutiny",
            },
            UpdatedAt: types.MustTimeFromString("2021-08-13T23:04:27.910Z"),
        },
        ConnectionID: "female",
        ID: "<ID>",
        ListID: "yellow overriding Rock",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PatchMartechConnectionIDListIDMemberIDRequest](../../models/operations/patchmartechconnectionidlistidmemberidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PatchMartechConnectionIDListIDMemberIDResponse](../../models/operations/patchmartechconnectionidlistidmemberidresponse.md), error**


## PostMartechConnectionIDListIDMember

Create a member in a list

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
    res, err := s.Member.PostMartechConnectionIDListIDMember(ctx, operations.PostMartechConnectionIDListIDMemberRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2022-12-27T02:49:51.488Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Esta.Dach@hotmail.com",
                    Type: shared.MarketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ListIds: []string{
                "virtual",
            },
            Name: unifiedgosdk.String("dolorum Wooden Granite"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "Road",
            },
            UpdatedAt: types.MustTimeFromString("2021-06-09T19:47:01.060Z"),
        },
        ConnectionID: "Pennsylvania leverage sheath",
        ListID: "parse exercitationem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PostMartechConnectionIDListIDMemberRequest](../../models/operations/postmartechconnectionidlistidmemberrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.PostMartechConnectionIDListIDMemberResponse](../../models/operations/postmartechconnectionidlistidmemberresponse.md), error**


## PutMartechConnectionIDListIDMemberID

Update a member in a list

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
    res, err := s.Member.PutMartechConnectionIDListIDMemberID(ctx, operations.PutMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2023-09-24T05:00:50.743Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Lorenz_Kautzer@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ListIds: []string{
                "SMS",
            },
            Name: unifiedgosdk.String("East platforms"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "Estonia",
            },
            UpdatedAt: types.MustTimeFromString("2023-01-20T05:09:32.775Z"),
        },
        ConnectionID: "following quia Intelligent",
        ID: "<ID>",
        ListID: "Cab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PutMartechConnectionIDListIDMemberIDRequest](../../models/operations/putmartechconnectionidlistidmemberidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PutMartechConnectionIDListIDMemberIDResponse](../../models/operations/putmartechconnectionidlistidmemberidresponse.md), error**

