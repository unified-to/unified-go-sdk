# Customer
(*Customer*)

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
    res, err := s.Customer.DeleteTicketingConnectionIDCustomerID(ctx, operations.DeleteTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "Electric Gloves pish",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.DeleteTicketingConnectionIDCustomerIDRequest](../../models/operations/deleteticketingconnectionidcustomeridrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


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
    res, err := s.Customer.GetTicketingConnectionIDCustomer(ctx, operations.GetTicketingConnectionIDCustomerRequest{
        ConnectionID: "SDD because Salad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetTicketingConnectionIDCustomerRequest](../../models/operations/getticketingconnectionidcustomerrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


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
    res, err := s.Customer.GetTicketingConnectionIDCustomerID(ctx, operations.GetTicketingConnectionIDCustomerIDRequest{
        ConnectionID: "further Ebert",
        ID: "<ID>",
    })
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
| `request`                                                                                                                    | [operations.GetTicketingConnectionIDCustomerIDRequest](../../models/operations/getticketingconnectionidcustomeridrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


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
    res, err := s.Customer.PatchTicketingConnectionIDCustomerID(ctx, operations.PatchTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Jaren_Ryan@hotmail.com",
                },
            },
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "Waco",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "youthfully orange",
                },
            },
        },
        ConnectionID: "Smyrna Hialeah auxiliary",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PatchTicketingConnectionIDCustomerIDRequest](../../models/operations/patchticketingconnectionidcustomeridrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


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
    res, err := s.Customer.PostTicketingConnectionIDCustomer(ctx, operations.PostTicketingConnectionIDCustomerRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Jaquelin.Boyer@yahoo.com",
                },
            },
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "withdrawal",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "Configuration neural",
                },
            },
        },
        ConnectionID: "Product Hybrid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TicketingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PostTicketingConnectionIDCustomerRequest](../../models/operations/postticketingconnectionidcustomerrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


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
    res, err := s.Customer.PutTicketingConnectionIDCustomerID(ctx, operations.PutTicketingConnectionIDCustomerIDRequest{
        TicketingCustomer: &shared.TicketingCustomer{
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Raleigh.Torp42@gmail.com",
                },
            },
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "copy",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "Developer Buckinghamshire Sausages",
                },
            },
        },
        ConnectionID: "kilogram",
        ID: "<ID>",
    })
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
| `request`                                                                                                                    | [operations.PutTicketingConnectionIDCustomerIDRequest](../../models/operations/putticketingconnectionidcustomeridrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PutTicketingConnectionIDCustomerIDResponse](../../models/operations/putticketingconnectionidcustomeridresponse.md), error**

