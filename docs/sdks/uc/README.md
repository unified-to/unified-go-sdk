# Uc

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.DeleteUcConnectionIDContactID(ctx, operations.DeleteUcConnectionIDContactIDRequest{
        ConnectionID: "incidunt",
        ID: "52a9f01f-3442-4c61-be13-3bacde532b65",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteUcConnectionIDContactIDRequest](../../models/operations/deleteucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.DeleteUcConnectionIDContactIDSecurity](../../models/operations/deleteucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUcConnectionIDAgentSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDAgent(ctx, operations.GetUcConnectionIDAgentRequest{
        ConnectionID: "eos",
        ContactID: unifiedto.String("laboriosam"),
        Limit: unifiedto.Float64(9714.32),
        Offset: unifiedto.Float64(5083.12),
        Order: unifiedto.String("suscipit"),
        Query: unifiedto.String("explicabo"),
        Sort: unifiedto.String("quos"),
        UpdatedGte: types.MustDateFromString("2022-10-16"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcAgents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetUcConnectionIDAgentRequest](../../models/operations/getucconnectionidagentrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetUcConnectionIDAgentSecurity](../../models/operations/getucconnectionidagentsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUcConnectionIDCallSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDCall(ctx, operations.GetUcConnectionIDCallRequest{
        AgentID: unifiedto.String("hic"),
        ConnectionID: "eveniet",
        ContactID: unifiedto.String("eos"),
        Limit: unifiedto.Float64(5126.45),
        Offset: unifiedto.Float64(3151.64),
        Order: unifiedto.String("provident"),
        Query: unifiedto.String("maxime"),
        Sort: unifiedto.String("officiis"),
        UpdatedGte: types.MustDateFromString("2022-11-01"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUcConnectionIDCallRequest](../../models/operations/getucconnectionidcallrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetUcConnectionIDCallSecurity](../../models/operations/getucconnectionidcallsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUcConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDContact(ctx, operations.GetUcConnectionIDContactRequest{
        AgentID: unifiedto.String("consequuntur"),
        ConnectionID: "quia",
        Limit: unifiedto.Float64(1905.14),
        Offset: unifiedto.Float64(718.84),
        Order: unifiedto.String("doloribus"),
        Query: unifiedto.String("earum"),
        Sort: unifiedto.String("commodi"),
        UpdatedGte: types.MustDateFromString("2022-08-09"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetUcConnectionIDContactRequest](../../models/operations/getucconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetUcConnectionIDContactSecurity](../../models/operations/getucconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.GetUcConnectionIDContactID(ctx, operations.GetUcConnectionIDContactIDRequest{
        ConnectionID: "dolore",
        ID: "c41d2fba-5cba-4069-b8d2-91beb810a2aa",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetUcConnectionIDContactIDRequest](../../models/operations/getucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetUcConnectionIDContactIDSecurity](../../models/operations/getucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.PatchUcConnectionIDContactID(ctx, operations.PatchUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Kreiger - Gutmann"),
            CreatedAt: types.MustDateFromString("2022-06-03"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Leanna_Walsh26@yahoo.com",
                    Type: shared.UcEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("cf7b50cf-87f0-48f3-9271-076a24b40c8f"),
            Name: unifiedto.String("Terry Rau"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "quae",
                    Type: shared.UcTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedto.String("Ms."),
            UpdatedAt: types.MustDateFromString("2022-02-04"),
        },
        ConnectionID: "laudantium",
        ID: "8f86996c-8e22-4be0-a3cf-47893bd23f86",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchUcConnectionIDContactIDRequest](../../models/operations/patchucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PatchUcConnectionIDContactIDSecurity](../../models/operations/patchucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostUcConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.PostUcConnectionIDContact(ctx, operations.PostUcConnectionIDContactRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Abbott - Beatty"),
            CreatedAt: types.MustDateFromString("2021-10-06"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Ova.Kovacek@gmail.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("273caa91-18b3-48f1-b61a-331a54dc1029"),
            Name: unifiedto.String("Johanna Muller"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "debitis",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedto.String("Miss"),
            UpdatedAt: types.MustDateFromString("2022-05-29"),
        },
        ConnectionID: "expedita",
    }, operationSecurity)
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
| `request`                                                                                                    | [operations.PostUcConnectionIDContactRequest](../../models/operations/postucconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PostUcConnectionIDContactSecurity](../../models/operations/postucconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Uc.PutUcConnectionIDContactID(ctx, operations.PutUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Lynch - Zemlak"),
            CreatedAt: types.MustDateFromString("2022-11-29"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Cary.McKenzie@hotmail.com",
                    Type: shared.UcEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("20ee1228-ac3a-4dc6-87d2-40bc11ea4828"),
            Name: unifiedto.String("Danielle Schamberger"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "aliquid",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustDateFromString("2022-10-31"),
        },
        ConnectionID: "reiciendis",
        ID: "5b9d3cb1-1a76-487d-b100-e8e2b9b0d746",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutUcConnectionIDContactIDRequest](../../models/operations/putucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PutUcConnectionIDContactIDSecurity](../../models/operations/putucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PutUcConnectionIDContactIDResponse](../../models/operations/putucconnectionidcontactidresponse.md), error**

