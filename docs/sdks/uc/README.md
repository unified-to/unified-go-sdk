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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDCall(ctx, operations.GetUcConnectionIDCallRequest{
        ConnectionID: "Directives",
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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDContact(ctx, operations.GetUcConnectionIDContactRequest{
        ConnectionID: "Refined Practical",
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
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Van84@hotmail.com",
                },
            },
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "Denar",
                },
            },
        },
        ConnectionID: "strategy Synergized",
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
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Gilda_Jacobs@gmail.com",
                },
            },
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "Account Orchestrator",
                },
            },
        },
        ConnectionID: "extend grey Avon",
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
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Darien12@hotmail.com",
                },
            },
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "reboot",
                },
            },
        },
        ConnectionID: "payment hem",
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

