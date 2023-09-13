# Connection

### Available Operations

* [DeleteUnifiedConnectionID](#deleteunifiedconnectionid) - Remove connection
* [GetUnifiedConnection](#getunifiedconnection) - List all connections
* [GetUnifiedConnectionID](#getunifiedconnectionid) - Retrieve connection
* [PatchUnifiedConnectionID](#patchunifiedconnectionid) - Update connection
* [PostUnifiedConnection](#postunifiedconnection) - Create connection
* [PutUnifiedConnectionID](#putunifiedconnectionid) - Update connection

## DeleteUnifiedConnectionID

Remove connection

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
    operationSecurity := operations.DeleteUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.DeleteUnifiedConnectionID(ctx, operations.DeleteUnifiedConnectionIDRequest{
        ID: "bf60c321-f023-4b75-9236-7fe1a0cc8df7",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteUnifiedConnectionIDRequest](../../models/operations/deleteunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.DeleteUnifiedConnectionIDSecurity](../../models/operations/deleteunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.DeleteUnifiedConnectionIDResponse](../../models/operations/deleteunifiedconnectionidresponse.md), error**


## GetUnifiedConnection

List all connections

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
    operationSecurity := operations.GetUnifiedConnectionSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.GetUnifiedConnection(ctx, operations.GetUnifiedConnectionRequest{
        Categories: []GetUnifiedConnectionCategories{
            operations.GetUnifiedConnectionCategoriesEnrich,
        },
        Env: unifiedto.String("sapiente"),
        ExternalXref: unifiedto.String("aperiam"),
        Limit: unifiedto.Float64(6277.17),
        Offset: unifiedto.Float64(1979.82),
        Order: unifiedto.String("provident"),
        Sort: unifiedto.String("ex"),
        UpdatedGte: types.MustDateFromString("2021-03-13"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUnifiedConnectionRequest](../../models/operations/getunifiedconnectionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.GetUnifiedConnectionSecurity](../../models/operations/getunifiedconnectionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.GetUnifiedConnectionResponse](../../models/operations/getunifiedconnectionresponse.md), error**


## GetUnifiedConnectionID

Retrieve connection

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
    operationSecurity := operations.GetUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.GetUnifiedConnectionID(ctx, operations.GetUnifiedConnectionIDRequest{
        ID: "0c364b7c-15df-4bac-a188-b1c4ee2c8c6c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetUnifiedConnectionIDRequest](../../models/operations/getunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetUnifiedConnectionIDSecurity](../../models/operations/getunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetUnifiedConnectionIDResponse](../../models/operations/getunifiedconnectionidresponse.md), error**


## PatchUnifiedConnectionID

Update connection

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
    operationSecurity := operations.PatchUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.PatchUnifiedConnectionID(ctx, operations.PatchUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedto.String("recusandae"),
                APIURL: unifiedto.String("ex"),
                AppID: unifiedto.String("beatae"),
                AuthorizeURL: unifiedto.String("veritatis"),
                ClientID: unifiedto.String("maiores"),
                ClientSecret: unifiedto.String("itaque"),
                ConsumerKey: unifiedto.String("vero"),
                ConsumerSecret: unifiedto.String("quidem"),
                Emails: []string{
                    "illo",
                },
                ExpiresIn: unifiedto.Float64(7782.42),
                ExpiryDate: types.MustDateFromString("2022-03-18"),
                Key: unifiedto.String("distinctio"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedto.String("Felipe Homenick"),
                OtherAuthInfo: []string{
                    "quod",
                },
                Pem: unifiedto.String("nihil"),
                RefreshToken: unifiedto.String("quaerat"),
                State: unifiedto.String("ipsum"),
                Token: unifiedto.String("ducimus"),
                TokenURL: unifiedto.String("laudantium"),
            },
            AuthAwsArn: unifiedto.String("rerum"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesEnrich,
            },
            CreatedAt: types.MustDateFromString("2022-09-06"),
            Environment: unifiedto.String("sequi"),
            ExternalXref: unifiedto.String("beatae"),
            ID: unifiedto.String("7747dc91-5ad2-4caf-9dd6-723dc0f5ae2f"),
            IntegrationType: "neque",
            IsPaused: unifiedto.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsTicketingCustomerRead,
            },
            UpdatedAt: types.MustDateFromString("2022-04-24"),
            WorkspaceID: unifiedto.String("ducimus"),
        },
        ID: "00878756-143f-45a6-898b-55554080d40b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchUnifiedConnectionIDRequest](../../models/operations/patchunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PatchUnifiedConnectionIDSecurity](../../models/operations/patchunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PatchUnifiedConnectionIDResponse](../../models/operations/patchunifiedconnectionidresponse.md), error**


## PostUnifiedConnection

Create connection

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
    operationSecurity := operations.PostUnifiedConnectionSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.PostUnifiedConnection(ctx, shared.Connection{
        Auth: &shared.PropertyConnectionAuth{
            AccessToken: unifiedto.String("nobis"),
            APIURL: unifiedto.String("similique"),
            AppID: unifiedto.String("porro"),
            AuthorizeURL: unifiedto.String("impedit"),
            ClientID: unifiedto.String("nisi"),
            ClientSecret: unifiedto.String("cumque"),
            ConsumerKey: unifiedto.String("soluta"),
            ConsumerSecret: unifiedto.String("fugiat"),
            Emails: []string{
                "laboriosam",
            },
            ExpiresIn: unifiedto.Float64(7203.19),
            ExpiryDate: types.MustDateFromString("2022-01-08"),
            Key: unifiedto.String("consectetur"),
            Meta: &shared.PropertyPropertyConnectionAuthMeta{},
            Name: unifiedto.String("Ms. Wilbert McGlynn"),
            OtherAuthInfo: []string{
                "accusantium",
            },
            Pem: unifiedto.String("magnam"),
            RefreshToken: unifiedto.String("repellat"),
            State: unifiedto.String("omnis"),
            Token: unifiedto.String("explicabo"),
            TokenURL: unifiedto.String("vel"),
        },
        AuthAwsArn: unifiedto.String("cum"),
        Categories: []shared.PropertyConnectionCategories{
            shared.PropertyConnectionCategoriesEnrich,
        },
        CreatedAt: types.MustDateFromString("2022-07-28"),
        Environment: unifiedto.String("ipsam"),
        ExternalXref: unifiedto.String("nostrum"),
        ID: unifiedto.String("3819b474-b0ed-420e-9624-8fff639a910a"),
        IntegrationType: "nam",
        IsPaused: unifiedto.Bool(false),
        Permissions: []shared.PropertyConnectionPermissions{
            shared.PropertyConnectionPermissionsHrisGroupWrite,
        },
        UpdatedAt: types.MustDateFromString("2020-12-12"),
        WorkspaceID: unifiedto.String("tempore"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.Connection](../../models/shared/connection.md)                                               | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.PostUnifiedConnectionSecurity](../../models/operations/postunifiedconnectionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.PostUnifiedConnectionResponse](../../models/operations/postunifiedconnectionresponse.md), error**


## PutUnifiedConnectionID

Update connection

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
    operationSecurity := operations.PutUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Connection.PutUnifiedConnectionID(ctx, operations.PutUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedto.String("commodi"),
                APIURL: unifiedto.String("fugit"),
                AppID: unifiedto.String("suscipit"),
                AuthorizeURL: unifiedto.String("voluptate"),
                ClientID: unifiedto.String("nisi"),
                ClientSecret: unifiedto.String("aliquid"),
                ConsumerKey: unifiedto.String("provident"),
                ConsumerSecret: unifiedto.String("laboriosam"),
                Emails: []string{
                    "accusamus",
                },
                ExpiresIn: unifiedto.Float64(682.92),
                ExpiryDate: types.MustDateFromString("2020-08-18"),
                Key: unifiedto.String("eaque"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedto.String("Miss Lois Cruickshank"),
                OtherAuthInfo: []string{
                    "amet",
                },
                Pem: unifiedto.String("exercitationem"),
                RefreshToken: unifiedto.String("illum"),
                State: unifiedto.String("praesentium"),
                Token: unifiedto.String("unde"),
                TokenURL: unifiedto.String("similique"),
            },
            AuthAwsArn: unifiedto.String("eligendi"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesMartech,
            },
            CreatedAt: types.MustDateFromString("2022-02-10"),
            Environment: unifiedto.String("nobis"),
            ExternalXref: unifiedto.String("asperiores"),
            ID: unifiedto.String("da8d0c54-9ef0-4300-8978-a61fa1cf2068"),
            IntegrationType: "totam",
            IsPaused: unifiedto.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsAtsJobRead,
            },
            UpdatedAt: types.MustDateFromString("2022-07-16"),
            WorkspaceID: unifiedto.String("quod"),
        },
        ID: "1ffc71dc-a163-4f2a-bc80-a97ff334cddf",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PutUnifiedConnectionIDRequest](../../models/operations/putunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.PutUnifiedConnectionIDSecurity](../../models/operations/putunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.PutUnifiedConnectionIDResponse](../../models/operations/putunifiedconnectionidresponse.md), error**

