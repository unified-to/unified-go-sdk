# Uc
(*Uc*)

### Available Operations

* [DeleteUcConnectionIDContactID](#deleteucconnectionidcontactid) - Remove a contact
* [GetUcConnectionIDAgent](#getucconnectionidagent) - List all agents
* [GetUcConnectionIDCall](#getucconnectionidcall) - List all calls
* [GetUcConnectionIDContact](#getucconnectionidcontact) - List all contacts
* [GetUcConnectionIDContactID](#getucconnectionidcontactid) - Retrieve a contact
* [PatchUcConnectionIDContactID](#patchucconnectionidcontactid) - Update a contact
* [PostUcConnectionIDContact](#postucconnectionidcontact) - Create a contact
* [PutUcConnectionIDContactID](#putucconnectionidcontactid) - Update a contact

## DeleteUcConnectionIDContactID

Remove a contact

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
    res, err := s.Uc.DeleteUcConnectionIDContactID(ctx, operations.DeleteUcConnectionIDContactIDRequest{
        ConnectionID: "Southeast Modern commonly",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteUcConnectionIDContactIDRequest](../../models/operations/deleteucconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteUcConnectionIDContactIDResponse](../../models/operations/deleteucconnectionidcontactidresponse.md), error**


## GetUcConnectionIDAgent

List all agents

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
    res, err := s.Uc.GetUcConnectionIDAgent(ctx, operations.GetUcConnectionIDAgentRequest{
        ConnectionID: "Regional East Sedan",
        ContactID: unifiedgosdk.String("blue"),
        Limit: unifiedgosdk.Float64(7827.68),
        Offset: unifiedgosdk.Float64(2116.69),
        Order: unifiedgosdk.String("Bicycle"),
        Query: unifiedgosdk.String("Bacon officia iterate"),
        Sort: unifiedgosdk.String("sticky vote lumen"),
        UpdatedGte: types.MustTimeFromString("2021-07-05T19:53:29.041Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUcConnectionIDAgentRequest](../../models/operations/getucconnectionidagentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetUcConnectionIDAgentResponse](../../models/operations/getucconnectionidagentresponse.md), error**


## GetUcConnectionIDCall

List all calls

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
    res, err := s.Uc.GetUcConnectionIDCall(ctx, operations.GetUcConnectionIDCallRequest{
        AgentID: unifiedgosdk.String("Directives"),
        ConnectionID: "female than",
        ContactID: unifiedgosdk.String("reintermediate Enid Applications"),
        Limit: unifiedgosdk.Float64(1980.39),
        Offset: unifiedgosdk.Float64(3478),
        Order: unifiedgosdk.String("white Oklahoma Functionality"),
        Query: unifiedgosdk.String("pricing whether Hillsboro"),
        Sort: unifiedgosdk.String("Wooden desensitize SCSI"),
        UpdatedGte: types.MustTimeFromString("2021-11-03T12:40:46.997Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUcConnectionIDCallRequest](../../models/operations/getucconnectionidcallrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetUcConnectionIDCallResponse](../../models/operations/getucconnectionidcallresponse.md), error**


## GetUcConnectionIDContact

List all contacts

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
    res, err := s.Uc.GetUcConnectionIDContact(ctx, operations.GetUcConnectionIDContactRequest{
        AgentID: unifiedgosdk.String("Refined Practical"),
        ConnectionID: "inasmuch Dodge",
        Limit: unifiedgosdk.Float64(7215.14),
        Offset: unifiedgosdk.Float64(2910.48),
        Order: unifiedgosdk.String("Vermont"),
        Query: unifiedgosdk.String("maroon JBOD"),
        Sort: unifiedgosdk.String("hertz"),
        UpdatedGte: types.MustTimeFromString("2023-01-29T17:06:35.136Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetUcConnectionIDContactRequest](../../models/operations/getucconnectionidcontactrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetUcConnectionIDContactResponse](../../models/operations/getucconnectionidcontactresponse.md), error**


## GetUcConnectionIDContactID

Retrieve a contact

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
    res, err := s.Uc.GetUcConnectionIDContactID(ctx, operations.GetUcConnectionIDContactIDRequest{
        ConnectionID: "Land",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetUcConnectionIDContactIDRequest](../../models/operations/getucconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetUcConnectionIDContactIDResponse](../../models/operations/getucconnectionidcontactidresponse.md), error**


## PatchUcConnectionIDContactID

Update a contact

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
    res, err := s.Uc.PatchUcConnectionIDContactID(ctx, operations.PatchUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedgosdk.String("Wilderman, Cremin and Gislason"),
            CreatedAt: types.MustTimeFromString("2023-07-18T06:13:06.229Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Henry.Leannon@gmail.com",
                    Type: shared.UcEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("quirky digital"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "Lead 24/7 overriding",
                    Type: shared.UcTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Small Legacy"),
            UpdatedAt: types.MustTimeFromString("2022-07-11T16:02:41.922Z"),
        },
        ConnectionID: "Bohrium",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchUcConnectionIDContactIDRequest](../../models/operations/patchucconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PatchUcConnectionIDContactIDResponse](../../models/operations/patchucconnectionidcontactidresponse.md), error**


## PostUcConnectionIDContact

Create a contact

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
    res, err := s.Uc.PostUcConnectionIDContact(ctx, operations.PostUcConnectionIDContactRequest{
        UcContact: &shared.UcContact{
            Company: unifiedgosdk.String("Howell and Sons"),
            CreatedAt: types.MustTimeFromString("2022-12-18T04:56:44.573Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Garret81@hotmail.com",
                    Type: shared.UcEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Southeast Gasoline extend"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "Togo Division Human",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("COM that"),
            UpdatedAt: types.MustTimeFromString("2023-02-07T16:19:58.439Z"),
        },
        ConnectionID: "Tennessee",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostUcConnectionIDContactRequest](../../models/operations/postucconnectionidcontactrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PostUcConnectionIDContactResponse](../../models/operations/postucconnectionidcontactresponse.md), error**


## PutUcConnectionIDContactID

Update a contact

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
    res, err := s.Uc.PutUcConnectionIDContactID(ctx, operations.PutUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedgosdk.String("Feeney, Gusikowski and Douglas"),
            CreatedAt: types.MustTimeFromString("2021-05-15T18:36:56.888Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Katrina.Walker@gmail.com",
                    Type: shared.UcEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Investment Hip Southwest"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "powerless Shirt",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Wooden Buckinghamshire"),
            UpdatedAt: types.MustTimeFromString("2022-10-29T19:58:07.810Z"),
        },
        ConnectionID: "doubtfully",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutUcConnectionIDContactIDRequest](../../models/operations/putucconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutUcConnectionIDContactIDResponse](../../models/operations/putucconnectionidcontactidresponse.md), error**

