# Customer

### Available Operations

* [DeleteTicketingConnectionIDCustomerID](#deleteticketingconnectionidcustomerid) - Remove a customer
* [GetTicketingConnectionIDCustomer](#getticketingconnectionidcustomer) - List all customers
* [GetTicketingConnectionIDCustomerID](#getticketingconnectionidcustomerid) - Retrieve a customer
* [PatchTicketingConnectionIDCustomerID](#patchticketingconnectionidcustomerid) - Update a customer
* [PostTicketingConnectionIDCustomer](#postticketingconnectionidcustomer) - Create a customer
* [PutTicketingConnectionIDCustomerID](#putticketingconnectionidcustomerid) - Update a customer

## DeleteTicketingConnectionIDCustomerID

Remove a customer

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
    operationSecurity := operations.DeleteTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.DeleteTicketingConnectionIDCustomerID(ctx, operations.DeleteTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "nulla",
        ID: "acd38ed0-dc67-41dc-bf1e-3af15920c90d",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.DeleteTicketingConnectionIDCustomerIDRequest](../../models/operations/deleteticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.DeleteTicketingConnectionIDCustomerIDSecurity](../../models/operations/deleteticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.DeleteTicketingConnectionIDCustomerIDResponse](../../models/operations/deleteticketingconnectionidcustomeridresponse.md), error**


## GetTicketingConnectionIDCustomer

List all customers

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
    operationSecurity := operations.GetTicketingConnectionIDCustomerSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.GetTicketingConnectionIDCustomer(ctx, operations.GetTicketingConnectionIDCustomerRequest{
        ConnectionID: "ab",
        Limit: unifiedto.Float64(7414.54),
        Offset: unifiedto.Float64(2975.19),
        Order: unifiedto.String("natus"),
        Query: unifiedto.String("aperiam"),
        Sort: unifiedto.String("dicta"),
        UpdatedGte: types.MustDateFromString("2022-06-25"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetTicketingConnectionIDCustomerRequest](../../models/operations/getticketingconnectionidcustomerrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.GetTicketingConnectionIDCustomerSecurity](../../models/operations/getticketingconnectionidcustomersecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.GetTicketingConnectionIDCustomerResponse](../../models/operations/getticketingconnectionidcustomerresponse.md), error**


## GetTicketingConnectionIDCustomerID

Retrieve a customer

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
    operationSecurity := operations.GetTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.GetTicketingConnectionIDCustomerID(ctx, operations.GetTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "harum",
        ID: "d89c8a32-639d-4a5b-bb69-02b881a94f64",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetTicketingConnectionIDCustomerIDRequest](../../models/operations/getticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.GetTicketingConnectionIDCustomerIDSecurity](../../models/operations/getticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.GetTicketingConnectionIDCustomerIDResponse](../../models/operations/getticketingconnectionidcustomeridresponse.md), error**


## PatchTicketingConnectionIDCustomerID

Update a customer

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
    operationSecurity := operations.PatchTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.PatchTicketingConnectionIDCustomerID(ctx, operations.PatchTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustDateFromString("2022-08-10"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Donny_OKeefe@yahoo.com",
                    Type: shared.TicketingEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("af8c691d-732d-49fb-af94-76a2ae8dcc50"),
            Name: unifiedto.String("Clifton Nicolas"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "dicta",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "eos",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2022-10-15"),
        },
        ConnectionID: "voluptate",
        ID: "84893075-0a00-4e96-aec7-36d43194398c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PatchTicketingConnectionIDCustomerIDRequest](../../models/operations/patchticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PatchTicketingConnectionIDCustomerIDSecurity](../../models/operations/patchticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


### Response

**[*operations.PatchTicketingConnectionIDCustomerIDResponse](../../models/operations/patchticketingconnectionidcustomeridresponse.md), error**


## PostTicketingConnectionIDCustomer

Create a customer

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
    operationSecurity := operations.PostTicketingConnectionIDCustomerSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.PostTicketingConnectionIDCustomer(ctx, operations.PostTicketingConnectionIDCustomerRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustDateFromString("2022-06-17"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Nils22@gmail.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("8ed3d3ab-7ca3-4c5c-a864-9a70cfd5d698"),
            Name: unifiedto.String("Rudy Kirlin III"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "magnam",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "ad",
                    Type: shared.TicketingTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2022-07-25"),
        },
        ConnectionID: "voluptate",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PostTicketingConnectionIDCustomerRequest](../../models/operations/postticketingconnectionidcustomerrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.PostTicketingConnectionIDCustomerSecurity](../../models/operations/postticketingconnectionidcustomersecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.PostTicketingConnectionIDCustomerResponse](../../models/operations/postticketingconnectionidcustomerresponse.md), error**


## PutTicketingConnectionIDCustomerID

Update a customer

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
    operationSecurity := operations.PutTicketingConnectionIDCustomerIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Customer.PutTicketingConnectionIDCustomerID(ctx, operations.PutTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            CreatedAt: types.MustDateFromString("2022-09-26"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Shanny.Padberg@yahoo.com",
                    Type: shared.TicketingEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("492ed14b-8a2c-4195-8545-e955dcc185ea"),
            Name: unifiedto.String("Miss Lindsay Adams"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "cumque",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "quaerat",
                    Type: shared.TicketingTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustDateFromString("2021-05-15"),
        },
        ConnectionID: "explicabo",
        ID: "daa784ab-a3d2-430e-9f73-811a115382bd",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PutTicketingConnectionIDCustomerIDRequest](../../models/operations/putticketingconnectionidcustomeridrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.PutTicketingConnectionIDCustomerIDSecurity](../../models/operations/putticketingconnectionidcustomeridsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |


### Response

**[*operations.PutTicketingConnectionIDCustomerIDResponse](../../models/operations/putticketingconnectionidcustomeridresponse.md), error**

