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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
        Limit: unifiedgosdk.Float64(8049.62),
        Offset: unifiedgosdk.Float64(4323.42),
        Order: unifiedgosdk.String("override"),
        Query: unifiedgosdk.String("Rolls 1080p"),
        Sort: unifiedgosdk.String("quantifying Southeast Kansas"),
        UpdatedGte: types.MustTimeFromString("2023-12-20T19:18:39.254Z"),
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            CreatedAt: types.MustTimeFromString("2023-01-22T19:33:25.134Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Ora.Labadie94@yahoo.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("sensitise whiteboard Smyrna"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "Hialeah",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "connect",
                    Type: shared.TicketingTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2023-12-28T17:48:45.929Z"),
        },
        ConnectionID: "Tennessine",
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            CreatedAt: types.MustTimeFromString("2022-05-23T15:06:12.012Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Austin44@yahoo.com",
                    Type: shared.TicketingEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Configuration neural"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "engineer",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "Gasoline North gorgeous",
                    Type: shared.TicketingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-10-09T07:25:23.111Z"),
        },
        ConnectionID: "mole purple",
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
	"github.com/unified-to/unified-go-sdk/pkg/types"
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
            CreatedAt: types.MustTimeFromString("2021-04-21T09:25:32.395Z"),
            Emails: []shared.TicketingEmail{
                shared.TicketingEmail{
                    Email: "Shawna42@hotmail.com",
                    Type: shared.TicketingEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("gray"),
            Raw: shared.PropertyTicketingCustomerRaw{},
            Tags: []string{
                "Associate",
            },
            Telephones: []shared.TicketingTelephone{
                shared.TicketingTelephone{
                    Telephone: "Sausages ivory Small",
                    Type: shared.TicketingTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-09-01T05:56:15.314Z"),
        },
        ConnectionID: "mobile Cotton",
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

