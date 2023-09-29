# User
(*User*)

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
    res, err := s.User.DeleteCrmConnectionIDUserID(ctx, operations.DeleteCrmConnectionIDUserIDRequest{
        ConnectionID: "Intranet Data",
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
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDUserIDRequest](../../models/operations/deletecrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.DeleteUnifiedUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


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
    res, err := s.User.GetCrmConnectionIDUser(ctx, operations.GetCrmConnectionIDUserRequest{
        ConnectionID: "suit Electronic Tampa",
        Limit: unifiedgosdk.Float64(2883.34),
        Offset: unifiedgosdk.Float64(8886.55),
        Order: unifiedgosdk.String("despite"),
        Query: unifiedgosdk.String("frightfully Fitness"),
        Sort: unifiedgosdk.String("success servant"),
        UpdatedGte: types.MustTimeFromString("2023-02-23T05:53:04.259Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDUserRequest](../../models/operations/getcrmconnectioniduserrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


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
    res, err := s.User.GetCrmConnectionIDUserID(ctx, operations.GetCrmConnectionIDUserIDRequest{
        ConnectionID: "connecting Program",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDUserIDRequest](../../models/operations/getcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetUnifiedUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetUnifiedUserToken(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedUserToken200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


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
    res, err := s.User.PatchCrmConnectionIDUserID(ctx, operations.PatchCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("Customer"),
                Address2: unifiedgosdk.String("violet groupware blanditiis"),
                City: unifiedgosdk.String("South Phoebeshire"),
                Country: unifiedgosdk.String("Thailand"),
                CountryCode: unifiedgosdk.String("NO"),
                PostalCode: unifiedgosdk.String("30801-4594"),
                Region: unifiedgosdk.String("portals Vanadium"),
                RegionCode: unifiedgosdk.String("Future"),
            },
            CreatedAt: types.MustTimeFromString("2023-01-04T02:42:28.788Z"),
            Currency: unifiedgosdk.String("Guinea Franc"),
            Department: unifiedgosdk.String("Gloves global rosin"),
            Division: unifiedgosdk.String("Berkshire Europium"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Wade.Dach@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Checking"),
            LanguageLocale: unifiedgosdk.String("Sedan Porsche matrix"),
            Name: unifiedgosdk.String("superstructure Nissan sedately"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "unto ubiquitous input",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("America/Tijuana"),
            Title: unifiedgosdk.String("Computer Bicycle"),
            UpdatedAt: types.MustTimeFromString("2021-12-13T16:36:33.886Z"),
        },
        ConnectionID: "gold generating",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDUserIDRequest](../../models/operations/patchcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.PatchUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2022-04-24T15:25:24.483Z"),
        Email: unifiedgosdk.String("Emmalee.Quitzon@yahoo.com"),
        Environment: unifiedgosdk.String("Bicycle"),
        ID: unifiedgosdk.String("<ID>"),
        Meta: &shared.PropertyUserMeta{},
        Name: "vice compressing",
        UpdatedAt: types.MustTimeFromString("2023-05-05T16:52:20.023Z"),
        WorkspaceID: "Hybrid methodologies",
        WorkspaceIds: []string{
            "Potassium",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


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
    res, err := s.User.PostCrmConnectionIDUser(ctx, operations.PostCrmConnectionIDUserRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("driver East"),
                Address2: unifiedgosdk.String("relationships Computer navigate"),
                City: unifiedgosdk.String("Homestead"),
                Country: unifiedgosdk.String("Cayman Islands"),
                CountryCode: unifiedgosdk.String("BW"),
                PostalCode: unifiedgosdk.String("34958"),
                Region: unifiedgosdk.String("South"),
                RegionCode: unifiedgosdk.String("morph an colossal"),
            },
            CreatedAt: types.MustTimeFromString("2021-02-02T08:27:21.693Z"),
            Currency: unifiedgosdk.String("Cayman Islands Dollar"),
            Department: unifiedgosdk.String("um"),
            Division: unifiedgosdk.String("West sievert generating"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jadon_Schumm45@gmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("radian Borders Southeast"),
            LanguageLocale: unifiedgosdk.String("Silicon Awesome Industrial"),
            Name: unifiedgosdk.String("Mouse Accounts"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "ohm Southeast ROI",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("Europe/Bratislava"),
            Title: unifiedgosdk.String("Money"),
            UpdatedAt: types.MustTimeFromString("2023-12-09T10:24:50.054Z"),
        },
        ConnectionID: "competent calculate",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDUserRequest](../../models/operations/postcrmconnectioniduserrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


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
    res, err := s.User.PutCrmConnectionIDUserID(ctx, operations.PutCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("Honduras"),
                Address2: unifiedgosdk.String("Oxygen Libyan Burundi"),
                City: unifiedgosdk.String("North Gertrudefield"),
                Country: unifiedgosdk.String("Belgium"),
                CountryCode: unifiedgosdk.String("DK"),
                PostalCode: unifiedgosdk.String("00781"),
                Region: unifiedgosdk.String("Wagon"),
                RegionCode: unifiedgosdk.String("how overriding"),
            },
            CreatedAt: types.MustTimeFromString("2023-03-13T00:47:15.649Z"),
            Currency: unifiedgosdk.String("Pakistan Rupee"),
            Department: unifiedgosdk.String("Tricycle vaguely"),
            Division: unifiedgosdk.String("Severn bluetooth Argon"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Karl_Stehr4@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("AGP romance didactic"),
            LanguageLocale: unifiedgosdk.String("ROI Polarised"),
            Name: unifiedgosdk.String("olive synthesizing"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Honda Indiana",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("Asia/Novosibirsk"),
            Title: unifiedgosdk.String("compelling red compressing"),
            UpdatedAt: types.MustTimeFromString("2022-09-03T15:59:05.095Z"),
        },
        ConnectionID: "relationships",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDUserIDRequest](../../models/operations/putcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.PutUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2023-07-31T04:46:29.769Z"),
        Email: unifiedgosdk.String("Selena59@yahoo.com"),
        Environment: unifiedgosdk.String("Bedfordshire Lucia"),
        ID: unifiedgosdk.String("<ID>"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Bicycle hacking South",
        UpdatedAt: types.MustTimeFromString("2023-03-15T15:08:26.238Z"),
        WorkspaceID: "Card defect",
        WorkspaceIds: []string{
            "repudiandae",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PutUnifiedUserResponse](../../models/operations/putunifieduserresponse.md), error**

