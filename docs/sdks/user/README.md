# User
(*User*)

### Available Operations

* [CreateCrmUser](#createcrmuser) - Create a user
* [GetCrmUser](#getcrmuser) - Retrieve a user
* [ListCrmUsers](#listcrmusers) - List all users
* [PatchCrmUser](#patchcrmuser) - Update a user
* [RemoveCrmUser](#removecrmuser) - Remove a user
* [UpdateCrmUser](#updatecrmuser) - Update a user

## CreateCrmUser

Create a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateCrmUser(ctx, operations.CreateCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "zero",
                },
            },
        },
        ConnectionID: "male",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmUserRequest](../../models/operations/createcrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmUserResponse](../../models/operations/createcrmuserresponse.md), error**


## GetCrmUser

Retrieve a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.GetCrmUser(ctx, operations.GetCrmUserRequest{
        ConnectionID: "adipisci",
        Fields: []string{
            "connect",
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmUserRequest](../../models/operations/getcrmuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmUserResponse](../../models/operations/getcrmuserresponse.md), error**


## ListCrmUsers

List all users

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.ListCrmUsers(ctx, operations.ListCrmUsersRequest{
        ConnectionID: "CFA",
        Fields: []string{
            "Costa",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUsers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmUsersRequest](../../models/operations/listcrmusersrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmUsersResponse](../../models/operations/listcrmusersresponse.md), error**


## PatchCrmUser

Update a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.PatchCrmUser(ctx, operations.PatchCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Demiflux",
                },
            },
        },
        ConnectionID: "grey",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmUserRequest](../../models/operations/patchcrmuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmUserResponse](../../models/operations/patchcrmuserresponse.md), error**


## RemoveCrmUser

Remove a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.RemoveCrmUser(ctx, operations.RemoveCrmUserRequest{
        ConnectionID: "yellow",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmUserRequest](../../models/operations/removecrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmUserResponse](../../models/operations/removecrmuserresponse.md), error**


## UpdateCrmUser

Update a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.UpdateCrmUser(ctx, operations.UpdateCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "benchmark",
                },
            },
        },
        ConnectionID: "Rustic",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmUserRequest](../../models/operations/updatecrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmUserResponse](../../models/operations/updatecrmuserresponse.md), error**

