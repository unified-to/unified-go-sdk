# Member

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteMartechConnectionIDListIDMemberIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.DeleteMartechConnectionIDListIDMemberID(ctx, operations.DeleteMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "alias",
        ID: "d2743fd6-c2a1-40e6-8297-8ec256a5b092",
        ListID: "magni",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.DeleteMartechConnectionIDListIDMemberIDRequest](../../models/operations/deletemartechconnectionidlistidmemberidrequest.md)   | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.DeleteMartechConnectionIDListIDMemberIDSecurity](../../models/operations/deletemartechconnectionidlistidmemberidsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetMartechConnectionIDListIDMemberSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.GetMartechConnectionIDListIDMember(ctx, operations.GetMartechConnectionIDListIDMemberRequest{
        ConnectionID: "iure",
        Limit: unifiedto.Float64(9859.05),
        ListID: "quod",
        Offset: unifiedto.Float64(8111.96),
        Order: unifiedto.String("numquam"),
        Query: unifiedto.String("dignissimos"),
        Sort: unifiedto.String("natus"),
        UpdatedGte: types.MustTimeFromString("2022-03-19T03:29:08.234Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetMartechConnectionIDListIDMemberRequest](../../models/operations/getmartechconnectionidlistidmemberrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.GetMartechConnectionIDListIDMemberSecurity](../../models/operations/getmartechconnectionidlistidmembersecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetMartechConnectionIDListIDMemberIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.GetMartechConnectionIDListIDMemberID(ctx, operations.GetMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "optio",
        ID: "977bbc57-f389-428a-8600-c58d67d63e4a",
        ListID: "officia",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetMartechConnectionIDListIDMemberIDRequest](../../models/operations/getmartechconnectionidlistidmemberidrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.GetMartechConnectionIDListIDMemberIDSecurity](../../models/operations/getmartechconnectionidlistidmemberidsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchMartechConnectionIDListIDMemberIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.PatchMartechConnectionIDListIDMemberID(ctx, operations.PatchMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2022-08-11T15:05:29.161Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Ella32@yahoo.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("9cfc6c0e-503f-4568-b1f1-d8ed87b28e8a"),
            ListIds: []string{
                "a",
            },
            Name: unifiedto.String("Felipe Schmeler"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "nisi",
            },
            UpdatedAt: types.MustTimeFromString("2022-07-05T01:28:10.122Z"),
        },
        ConnectionID: "aliquam",
        ID: "1e43b234-2417-4d13-a3f6-2aa9ae4ae8ab",
        ListID: "numquam",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PatchMartechConnectionIDListIDMemberIDRequest](../../models/operations/patchmartechconnectionidlistidmemberidrequest.md)   | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.PatchMartechConnectionIDListIDMemberIDSecurity](../../models/operations/patchmartechconnectionidlistidmemberidsecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostMartechConnectionIDListIDMemberSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.PostMartechConnectionIDListIDMember(ctx, operations.PostMartechConnectionIDListIDMemberRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2021-10-26T01:34:29.397Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Elena80@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("e8ba5d4a-a4a5-408b-9380-c29aa8dd71bd"),
            ListIds: []string{
                "repellendus",
            },
            Name: unifiedto.String("Miss Hubert Emard"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "cum",
            },
            UpdatedAt: types.MustTimeFromString("2022-10-14T03:48:43.146Z"),
        },
        ConnectionID: "labore",
        ListID: "modi",
    }, operationSecurity)
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
| `request`                                                                                                                        | [operations.PostMartechConnectionIDListIDMemberRequest](../../models/operations/postmartechconnectionidlistidmemberrequest.md)   | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `security`                                                                                                                       | [operations.PostMartechConnectionIDListIDMemberSecurity](../../models/operations/postmartechconnectionidlistidmembersecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutMartechConnectionIDListIDMemberIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Member.PutMartechConnectionIDListIDMemberID(ctx, operations.PutMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustTimeFromString("2021-09-21T15:05:12.667Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Geoffrey.Mitchell53@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("d418bb71-804f-4423-9543-935f377ac5c9"),
            ListIds: []string{
                "nam",
            },
            Name: unifiedto.String("Gretchen Moore"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "suscipit",
            },
            UpdatedAt: types.MustTimeFromString("2022-08-12T05:09:42.944Z"),
        },
        ConnectionID: "optio",
        ID: "523105e7-c34c-4ab0-acb8-12a66148944a",
        ListID: "voluptatum",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PutMartechConnectionIDListIDMemberIDRequest](../../models/operations/putmartechconnectionidlistidmemberidrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PutMartechConnectionIDListIDMemberIDSecurity](../../models/operations/putmartechconnectionidlistidmemberidsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


### Response

**[*operations.PutMartechConnectionIDListIDMemberIDResponse](../../models/operations/putmartechconnectionidlistidmemberidresponse.md), error**

