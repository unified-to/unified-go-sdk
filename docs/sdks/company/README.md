# Company
(*Company*)

### Available Operations

* [DeleteCrmConnectionIDCompanyID](#deletecrmconnectionidcompanyid) - Remove a company
* [GetCrmConnectionIDCompany](#getcrmconnectionidcompany) - List all companies
* [GetCrmConnectionIDCompanyID](#getcrmconnectionidcompanyid) - Retrieve a company
* [GetEnrichConnectionIDCompany](#getenrichconnectionidcompany) - Retrieve enrichment information for a company
* [PatchCrmConnectionIDCompanyID](#patchcrmconnectionidcompanyid) - Update a company
* [PostCrmConnectionIDCompany](#postcrmconnectionidcompany) - Create a company
* [PutCrmConnectionIDCompanyID](#putcrmconnectionidcompanyid) - Update a company

## DeleteCrmConnectionIDCompanyID

Remove a company

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
    res, err := s.Company.DeleteCrmConnectionIDCompanyID(ctx, operations.DeleteCrmConnectionIDCompanyIDRequest{
        ConnectionID: "hertz morph",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteCrmConnectionIDCompanyIDRequest](../../models/operations/deletecrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteCrmConnectionIDCompanyIDResponse](../../models/operations/deletecrmconnectionidcompanyidresponse.md), error**


## GetCrmConnectionIDCompany

List all companies

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
    res, err := s.Company.GetCrmConnectionIDCompany(ctx, operations.GetCrmConnectionIDCompanyRequest{
        ConnectionID: "indexing",
        ContactID: unifiedgosdk.String("Porsche firewall"),
        DealID: unifiedgosdk.String("Hafnium Computers"),
        Limit: unifiedgosdk.Float64(902.85),
        Offset: unifiedgosdk.Float64(2893.88),
        Order: unifiedgosdk.String("Interactions relationships juxtapose"),
        Query: unifiedgosdk.String("newton Luxembourg"),
        Sort: unifiedgosdk.String("Dakota quantifying Actinium"),
        UpdatedGte: types.MustTimeFromString("2022-09-27T07:42:48.074Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDCompanyRequest](../../models/operations/getcrmconnectionidcompanyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetCrmConnectionIDCompanyResponse](../../models/operations/getcrmconnectionidcompanyresponse.md), error**


## GetCrmConnectionIDCompanyID

Retrieve a company

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
    res, err := s.Company.GetCrmConnectionIDCompanyID(ctx, operations.GetCrmConnectionIDCompanyIDRequest{
        ConnectionID: "Netherlands",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCrmConnectionIDCompanyIDRequest](../../models/operations/getcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCrmConnectionIDCompanyIDResponse](../../models/operations/getcrmconnectionidcompanyidresponse.md), error**


## GetEnrichConnectionIDCompany

Retrieve enrichment information for a company

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
    res, err := s.Company.GetEnrichConnectionIDCompany(ctx, operations.GetEnrichConnectionIDCompanyRequest{
        ConnectionID: "female Computers Central",
        Domain: unifiedgosdk.String("scientific-facet.biz"),
        Name: unifiedgosdk.String("Outdoors embrace interface"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetEnrichConnectionIDCompanyRequest](../../models/operations/getenrichconnectionidcompanyrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetEnrichConnectionIDCompanyResponse](../../models/operations/getenrichconnectionidcompanyresponse.md), error**


## PatchCrmConnectionIDCompanyID

Update a company

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
    res, err := s.Company.PatchCrmConnectionIDCompanyID(ctx, operations.PatchCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("invoice"),
                Address2: unifiedgosdk.String("indexing Ford"),
                City: unifiedgosdk.String("McAllen"),
                Country: unifiedgosdk.String("Netherlands"),
                CountryCode: unifiedgosdk.String("PF"),
                PostalCode: unifiedgosdk.String("93486"),
                Region: unifiedgosdk.String("Steel impactful"),
                RegionCode: unifiedgosdk.String("Dong"),
            },
            CreatedAt: types.MustTimeFromString("2023-07-25T08:43:38.995Z"),
            DealIds: []string{
                "usefully",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Annabel31@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Toys Vermont Astatine"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "Trigender",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "female",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-12-22T11:39:56.432Z"),
            Websites: []string{
                "Latin",
            },
        },
        ConnectionID: "North kilogram connecting",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchCrmConnectionIDCompanyIDRequest](../../models/operations/patchcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PatchCrmConnectionIDCompanyIDResponse](../../models/operations/patchcrmconnectionidcompanyidresponse.md), error**


## PostCrmConnectionIDCompany

Create a company

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
    res, err := s.Company.PostCrmConnectionIDCompany(ctx, operations.PostCrmConnectionIDCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("consequently gosh"),
                Address2: unifiedgosdk.String("phooey"),
                City: unifiedgosdk.String("Antonettaville"),
                Country: unifiedgosdk.String("Lebanon"),
                CountryCode: unifiedgosdk.String("SI"),
                PostalCode: unifiedgosdk.String("79462"),
                Region: unifiedgosdk.String("orchid Oxygen Kids"),
                RegionCode: unifiedgosdk.String("Electric utilisation"),
            },
            CreatedAt: types.MustTimeFromString("2021-10-01T08:46:18.197Z"),
            DealIds: []string{
                "Tennessee",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jaida.McDermott26@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Hydrogen Wooden"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "CSS",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Account invoice",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-09-12T17:31:24.634Z"),
            Websites: []string{
                "Intuitive",
            },
        },
        ConnectionID: "Gasoline",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostCrmConnectionIDCompanyRequest](../../models/operations/postcrmconnectionidcompanyrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostCrmConnectionIDCompanyResponse](../../models/operations/postcrmconnectionidcompanyresponse.md), error**


## PutCrmConnectionIDCompanyID

Update a company

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
    res, err := s.Company.PutCrmConnectionIDCompanyID(ctx, operations.PutCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("Northwest Northwest"),
                Address2: unifiedgosdk.String("portals Diesel"),
                City: unifiedgosdk.String("Azusa"),
                Country: unifiedgosdk.String("Qatar"),
                CountryCode: unifiedgosdk.String("CG"),
                PostalCode: unifiedgosdk.String("52396"),
                Region: unifiedgosdk.String("Tuna sticky lest"),
                RegionCode: unifiedgosdk.String("Soft boo Missoula"),
            },
            CreatedAt: types.MustTimeFromString("2022-05-14T19:17:30.970Z"),
            DealIds: []string{
                "Hybrid",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Vance_Cruickshank93@gmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("lest Northwest"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "East",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Bronze round",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-02-21T09:41:36.696Z"),
            Websites: []string{
                "Keyboard",
            },
        },
        ConnectionID: "orange Bespoke",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutCrmConnectionIDCompanyIDRequest](../../models/operations/putcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutCrmConnectionIDCompanyIDResponse](../../models/operations/putcrmconnectionidcompanyidresponse.md), error**

