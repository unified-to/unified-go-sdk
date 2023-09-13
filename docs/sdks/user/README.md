# User

### Available Operations

* [DeleteCrmConnectionIDUserID](#deletecrmconnectioniduserid) - Remove a user
* [DeleteUnifiedUser](#deleteunifieduser) - Delete your user object
* [GetCrmConnectionIDUser](#getcrmconnectioniduser) - List all users
* [GetCrmConnectionIDUserID](#getcrmconnectioniduserid) - Retrieve a user
* [GetUnifiedUser](#getunifieduser) - Retrieve your user object
* [GetUnifiedUserToken](#getunifiedusertoken) - Retrieve your user token
* [PatchCrmConnectionIDUserID](#patchcrmconnectioniduserid) - Update a user
* [PatchUnifiedUser](#patchunifieduser) - Updates your user object
* [PostCrmConnectionIDUser](#postcrmconnectioniduser) - Create a user
* [PutCrmConnectionIDUserID](#putcrmconnectioniduserid) - Update a user
* [PutUnifiedUser](#putunifieduser) - Updates your user object

## DeleteCrmConnectionIDUserID

Remove a user

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
    operationSecurity := operations.DeleteCrmConnectionIDUserIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.DeleteCrmConnectionIDUserID(ctx, operations.DeleteCrmConnectionIDUserIDRequest{
        ConnectionID: "deleniti",
        ID: "b1d187b5-1eb5-4fd3-8bfe-03490cf20254",
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
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDUserIDRequest](../../models/operations/deletecrmconnectioniduseridrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.DeleteCrmConnectionIDUserIDSecurity](../../models/operations/deletecrmconnectioniduseridsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.DeleteCrmConnectionIDUserIDResponse](../../models/operations/deletecrmconnectioniduseridresponse.md), error**


## DeleteUnifiedUser

Delete your user object

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
    operationSecurity := operations.DeleteUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.DeleteUnifiedUser(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `security`                                                                                   | [operations.DeleteUnifiedUserSecurity](../../models/operations/deleteunifiedusersecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.DeleteUnifiedUserResponse](../../models/operations/deleteunifieduserresponse.md), error**


## GetCrmConnectionIDUser

List all users

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
    operationSecurity := operations.GetCrmConnectionIDUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.GetCrmConnectionIDUser(ctx, operations.GetCrmConnectionIDUserRequest{
        ConnectionID: "id",
        Limit: unifiedto.Float64(6080.08),
        Offset: unifiedto.Float64(3220.16),
        Order: unifiedto.String("unde"),
        Query: unifiedto.String("consequatur"),
        Sort: unifiedto.String("quaerat"),
        UpdatedGte: types.MustDateFromString("2022-03-13"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUsers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDUserRequest](../../models/operations/getcrmconnectioniduserrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCrmConnectionIDUserSecurity](../../models/operations/getcrmconnectionidusersecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetCrmConnectionIDUserResponse](../../models/operations/getcrmconnectioniduserresponse.md), error**


## GetCrmConnectionIDUserID

Retrieve a user

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
    operationSecurity := operations.GetCrmConnectionIDUserIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.GetCrmConnectionIDUserID(ctx, operations.GetCrmConnectionIDUserIDRequest{
        ConnectionID: "distinctio",
        ID: "462d6bc9-917f-498e-8792-b979a413d6a8",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDUserIDRequest](../../models/operations/getcrmconnectioniduseridrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetCrmConnectionIDUserIDSecurity](../../models/operations/getcrmconnectioniduseridsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.GetCrmConnectionIDUserIDResponse](../../models/operations/getcrmconnectioniduseridresponse.md), error**


## GetUnifiedUser

Retrieve your user object

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
    operationSecurity := operations.GetUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.GetUnifiedUser(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.GetUnifiedUserSecurity](../../models/operations/getunifiedusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetUnifiedUserResponse](../../models/operations/getunifieduserresponse.md), error**


## GetUnifiedUserToken

Retrieve your user token

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
    operationSecurity := operations.GetUnifiedUserTokenSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.GetUnifiedUserToken(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedUserToken200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `security`                                                                                       | [operations.GetUnifiedUserTokenSecurity](../../models/operations/getunifiedusertokensecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetUnifiedUserTokenResponse](../../models/operations/getunifiedusertokenresponse.md), error**


## PatchCrmConnectionIDUserID

Update a user

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
    operationSecurity := operations.PatchCrmConnectionIDUserIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.PatchCrmConnectionIDUserID(ctx, operations.PatchCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedto.String("impedit"),
                Address2: unifiedto.String("perspiciatis"),
                City: unifiedto.String("South Josianeberg"),
                Country: unifiedto.String("Papua New Guinea"),
                CountryCode: unifiedto.String("TG"),
                PostalCode: unifiedto.String("40808-6577"),
                Region: unifiedto.String("repellat"),
                RegionCode: unifiedto.String("voluptatum"),
            },
            CreatedAt: types.MustDateFromString("2021-04-22"),
            Currency: unifiedto.String("amet"),
            Department: unifiedto.String("totam"),
            Division: unifiedto.String("ex"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Bennie_Langosh@gmail.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("76c002fa-cb13-4ac2-8c81-43b866c575a1"),
            ImageURL: unifiedto.String("recusandae"),
            LanguageLocale: unifiedto.String("quia"),
            Name: unifiedto.String("Carla Lubowitz"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "accusantium",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedto.String("accusamus"),
            Title: unifiedto.String("Mrs."),
            UpdatedAt: types.MustDateFromString("2022-04-24"),
        },
        ConnectionID: "sit",
        ID: "e8fbc48d-dc7e-469b-9351-0505014dca10",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDUserIDRequest](../../models/operations/patchcrmconnectioniduseridrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PatchCrmConnectionIDUserIDSecurity](../../models/operations/patchcrmconnectioniduseridsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PatchCrmConnectionIDUserIDResponse](../../models/operations/patchcrmconnectioniduseridresponse.md), error**


## PatchUnifiedUser

Only the name field is updated.

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
    operationSecurity := operations.PatchUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.PatchUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustDateFromString("2022-06-17"),
        Email: unifiedto.String("Candido.Hahn@yahoo.com"),
        Environment: unifiedto.String("impedit"),
        ID: unifiedto.String("36e94889-2782-4d34-a0b8-fc0d59f57b9f"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Miss Ian Connelly",
        UpdatedAt: types.MustDateFromString("2022-07-09"),
        WorkspaceID: "deleniti",
        WorkspaceIds: []string{
            "aperiam",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.User](../../models/shared/user.md)                                                 | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.PatchUnifiedUserSecurity](../../models/operations/patchunifiedusersecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.PatchUnifiedUserResponse](../../models/operations/patchunifieduserresponse.md), error**


## PostCrmConnectionIDUser

Create a user

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
    operationSecurity := operations.PostCrmConnectionIDUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.PostCrmConnectionIDUser(ctx, operations.PostCrmConnectionIDUserRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedto.String("quos"),
                Address2: unifiedto.String("maxime"),
                City: unifiedto.String("South Nestor"),
                Country: unifiedto.String("Mongolia"),
                CountryCode: unifiedto.String("TO"),
                PostalCode: unifiedto.String("94023"),
                Region: unifiedto.String("magnam"),
                RegionCode: unifiedto.String("recusandae"),
            },
            CreatedAt: types.MustDateFromString("2022-12-13"),
            Currency: unifiedto.String("maiores"),
            Department: unifiedto.String("tempora"),
            Division: unifiedto.String("reprehenderit"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Stuart22@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("483f748e-efcc-4b69-9541-b4b393f35666"),
            ImageURL: unifiedto.String("consequuntur"),
            LanguageLocale: unifiedto.String("quis"),
            Name: unifiedto.String("Tomas Pacocha"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "consequuntur",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Timezone: unifiedto.String("illo"),
            Title: unifiedto.String("Dr."),
            UpdatedAt: types.MustDateFromString("2020-08-23"),
        },
        ConnectionID: "sequi",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostCrmConnectionIDUserRequest](../../models/operations/postcrmconnectioniduserrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PostCrmConnectionIDUserSecurity](../../models/operations/postcrmconnectionidusersecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.PostCrmConnectionIDUserResponse](../../models/operations/postcrmconnectioniduserresponse.md), error**


## PutCrmConnectionIDUserID

Update a user

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
    operationSecurity := operations.PutCrmConnectionIDUserIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.PutCrmConnectionIDUserID(ctx, operations.PutCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedto.String("reprehenderit"),
                Address2: unifiedto.String("sint"),
                City: unifiedto.String("Hintzfurt"),
                Country: unifiedto.String("Martinique"),
                CountryCode: unifiedto.String("TR"),
                PostalCode: unifiedto.String("08257-3819"),
                Region: unifiedto.String("omnis"),
                RegionCode: unifiedto.String("itaque"),
            },
            CreatedAt: types.MustDateFromString("2022-11-29"),
            Currency: unifiedto.String("fugiat"),
            Department: unifiedto.String("provident"),
            Division: unifiedto.String("voluptatem"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Rosanna45@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("6bfc7677-f0f5-404a-ae48-28fb6daee19c"),
            ImageURL: unifiedto.String("explicabo"),
            LanguageLocale: unifiedto.String("nisi"),
            Name: unifiedto.String("Frank Ryan"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "quasi",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Timezone: unifiedto.String("maxime"),
            Title: unifiedto.String("Ms."),
            UpdatedAt: types.MustDateFromString("2022-10-11"),
        },
        ConnectionID: "vitae",
        ID: "cabdab76-7a44-44dd-8da0-abe58eb3d54b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDUserIDRequest](../../models/operations/putcrmconnectioniduseridrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PutCrmConnectionIDUserIDSecurity](../../models/operations/putcrmconnectioniduseridsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PutCrmConnectionIDUserIDResponse](../../models/operations/putcrmconnectioniduseridresponse.md), error**


## PutUnifiedUser

Only the name field is updated.

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
    operationSecurity := operations.PutUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.User.PutUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustDateFromString("2022-11-07"),
        Email: unifiedto.String("Melissa_Dooley30@hotmail.com"),
        Environment: unifiedto.String("sint"),
        ID: unifiedto.String("b8e5c18b-25e8-47f6-8823-255be95c0cbc"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Billy Schmeler",
        UpdatedAt: types.MustDateFromString("2022-06-05"),
        WorkspaceID: "quae",
        WorkspaceIds: []string{
            "quos",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.User](../../models/shared/user.md)                                             | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.PutUnifiedUserSecurity](../../models/operations/putunifiedusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.PutUnifiedUserResponse](../../models/operations/putunifieduserresponse.md), error**

