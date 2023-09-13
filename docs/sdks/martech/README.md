# Martech

### Available Operations

* [DeleteMartechConnectionIDListID](#deletemartechconnectionidlistid) - Remove a list
* [DeleteMartechConnectionIDListIDMemberID](#deletemartechconnectionidlistidmemberid) - Remove member from a list
* [GetMartechConnectionIDList](#getmartechconnectionidlist) - List all lists
* [GetMartechConnectionIDListID](#getmartechconnectionidlistid) - Retrieve a list
* [GetMartechConnectionIDListIDMember](#getmartechconnectionidlistidmember) - List all members in a list
* [GetMartechConnectionIDListIDMemberID](#getmartechconnectionidlistidmemberid) - Retrieve a member from a list
* [PatchMartechConnectionIDListID](#patchmartechconnectionidlistid) - Update a list
* [PatchMartechConnectionIDListIDMemberID](#patchmartechconnectionidlistidmemberid) - Update a member in a list
* [PostMartechConnectionIDList](#postmartechconnectionidlist) - Create a list
* [PostMartechConnectionIDListIDMember](#postmartechconnectionidlistidmember) - Create a member in a list
* [PutMartechConnectionIDListID](#putmartechconnectionidlistid) - Update a list
* [PutMartechConnectionIDListIDMemberID](#putmartechconnectionidlistidmemberid) - Update a member in a list

## DeleteMartechConnectionIDListID

Remove a list

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
    operationSecurity := operations.DeleteMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.DeleteMartechConnectionIDListID(ctx, operations.DeleteMartechConnectionIDListIDRequest{
        ConnectionID: "expedita",
        ID: "855e889d-9ef9-432e-9000-a13ad8124208",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteMartechConnectionIDListIDRequest](../../models/operations/deletemartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.DeleteMartechConnectionIDListIDSecurity](../../models/operations/deletemartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.DeleteMartechConnectionIDListIDResponse](../../models/operations/deletemartechconnectionidlistidresponse.md), error**


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
    res, err := s.Martech.DeleteMartechConnectionIDListIDMemberID(ctx, operations.DeleteMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "voluptates",
        ID: "fd234118-98e7-4387-9efb-e8baebabb794",
        ListID: "voluptas",
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


## GetMartechConnectionIDList

List all lists

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
    operationSecurity := operations.GetMartechConnectionIDListSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.GetMartechConnectionIDList(ctx, operations.GetMartechConnectionIDListRequest{
        ConnectionID: "neque",
        Limit: unifiedto.Float64(4300.03),
        Offset: unifiedto.Float64(9186.14),
        Order: unifiedto.String("excepturi"),
        Query: unifiedto.String("ipsa"),
        Sort: unifiedto.String("velit"),
        UpdatedGte: types.MustDateFromString("2022-12-03"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetMartechConnectionIDListRequest](../../models/operations/getmartechconnectionidlistrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetMartechConnectionIDListSecurity](../../models/operations/getmartechconnectionidlistsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.GetMartechConnectionIDListResponse](../../models/operations/getmartechconnectionidlistresponse.md), error**


## GetMartechConnectionIDListID

Retrieve a list

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
    operationSecurity := operations.GetMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.GetMartechConnectionIDListID(ctx, operations.GetMartechConnectionIDListIDRequest{
        ConnectionID: "harum",
        ID: "b9763172-0b77-4a5a-9365-a79f15271f01",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetMartechConnectionIDListIDRequest](../../models/operations/getmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.GetMartechConnectionIDListIDSecurity](../../models/operations/getmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.GetMartechConnectionIDListIDResponse](../../models/operations/getmartechconnectionidlistidresponse.md), error**


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
    res, err := s.Martech.GetMartechConnectionIDListIDMember(ctx, operations.GetMartechConnectionIDListIDMemberRequest{
        ConnectionID: "quo",
        Limit: unifiedto.Float64(34.09),
        ListID: "illum",
        Offset: unifiedto.Float64(2078.87),
        Order: unifiedto.String("commodi"),
        Query: unifiedto.String("veritatis"),
        Sort: unifiedto.String("reiciendis"),
        UpdatedGte: types.MustDateFromString("2020-05-30"),
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
    res, err := s.Martech.GetMartechConnectionIDListIDMemberID(ctx, operations.GetMartechConnectionIDListIDMemberIDRequest{
        ConnectionID: "blanditiis",
        ID: "dc5effb4-53e9-4089-a871-fdb4d697bdd9",
        ListID: "quisquam",
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


## PatchMartechConnectionIDListID

Update a list

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
    operationSecurity := operations.PatchMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.PatchMartechConnectionIDListID(ctx, operations.PatchMartechConnectionIDListIDRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("molestias"),
            ID: unifiedto.String("85e43734-a5d7-42d9-add7-85be5e7afe55"),
            Name: unifiedto.String("Kristina Kozey"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustDateFromString("2022-11-15"),
        },
        ConnectionID: "laudantium",
        ID: "1f44e3a2-3394-4a68-8c80-d30ff72164d0",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchMartechConnectionIDListIDRequest](../../models/operations/patchmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PatchMartechConnectionIDListIDSecurity](../../models/operations/patchmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.PatchMartechConnectionIDListIDResponse](../../models/operations/patchmartechconnectionidlistidresponse.md), error**


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
    res, err := s.Martech.PatchMartechConnectionIDListIDMemberID(ctx, operations.PatchMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustDateFromString("2021-11-04"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Zora_Volkman@gmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("6553b89e-0009-4c66-92de-7b3562201a6a"),
            ListIds: []string{
                "error",
            },
            Name: unifiedto.String("Ray O'Connell"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "harum",
            },
            UpdatedAt: types.MustDateFromString("2022-04-25"),
        },
        ConnectionID: "ipsam",
        ID: "b908d4e3-0491-4aa3-9d4a-839f03bab77b",
        ListID: "occaecati",
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


## PostMartechConnectionIDList

Create a list

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
    operationSecurity := operations.PostMartechConnectionIDListSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.PostMartechConnectionIDList(ctx, operations.PostMartechConnectionIDListRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("quasi"),
            ID: unifiedto.String("8f031398-4507-4e0e-b9c7-e23ecb060465"),
            Name: unifiedto.String("Kate Cummerata"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustDateFromString("2022-02-21"),
        },
        ConnectionID: "autem",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostMartechConnectionIDListRequest](../../models/operations/postmartechconnectionidlistrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PostMartechConnectionIDListSecurity](../../models/operations/postmartechconnectionidlistsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.PostMartechConnectionIDListResponse](../../models/operations/postmartechconnectionidlistresponse.md), error**


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
    res, err := s.Martech.PostMartechConnectionIDListIDMember(ctx, operations.PostMartechConnectionIDListIDMemberRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustDateFromString("2021-11-13"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Ines_Swift@yahoo.com",
                    Type: shared.MarketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("8f7f002d-1986-4aa9-9d3a-1d32329e4583"),
            ListIds: []string{
                "voluptate",
            },
            Name: unifiedto.String("Alberto Wisozk"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "fugiat",
            },
            UpdatedAt: types.MustDateFromString("2022-04-07"),
        },
        ConnectionID: "harum",
        ListID: "inventore",
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


## PutMartechConnectionIDListID

Update a list

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
    operationSecurity := operations.PutMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Martech.PutMartechConnectionIDListID(ctx, operations.PutMartechConnectionIDListIDRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("aut"),
            ID: unifiedto.String("e255fdc4-80d6-4e33-8867-5cbf186856a7"),
            Name: unifiedto.String("Jaime Conroy"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustDateFromString("2021-03-04"),
        },
        ConnectionID: "assumenda",
        ID: "0fc282c6-66af-43c3-b558-9bea5d264e41",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutMartechConnectionIDListIDRequest](../../models/operations/putmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PutMartechConnectionIDListIDSecurity](../../models/operations/putmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PutMartechConnectionIDListIDResponse](../../models/operations/putmartechconnectionidlistidresponse.md), error**


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
    res, err := s.Martech.PutMartechConnectionIDListIDMemberID(ctx, operations.PutMartechConnectionIDListIDMemberIDRequest{
        MarketingMember: &shared.MarketingMember{
            CreatedAt: types.MustDateFromString("2022-06-10"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "Mark54@hotmail.com",
                    Type: shared.MarketingEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("2e513f6d-9d2a-4d37-8309-9077c10b6879"),
            ListIds: []string{
                "sed",
            },
            Name: unifiedto.String("Lucy Fahey"),
            Raw: &shared.PropertyMarketingMemberRaw{},
            Tags: []string{
                "nihil",
            },
            UpdatedAt: types.MustDateFromString("2022-03-31"),
        },
        ConnectionID: "rem",
        ID: "860543c0-a304-49c3-8f6c-0276e7b21bad",
        ListID: "provident",
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

