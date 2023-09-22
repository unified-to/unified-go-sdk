# Company

### Available Operations

* [DeleteCrmConnectionIDCompanyID](#deletecrmconnectionidcompanyid) - Remove a company
* [DeleteCrmConnectionIDCompanyIDDealDealID](#deletecrmconnectionidcompanyiddealdealid) - Remove deal association from a company
* [GetCrmConnectionIDCompany](#getcrmconnectionidcompany) - List all companies
* [GetCrmConnectionIDCompanyID](#getcrmconnectionidcompanyid) - Retrieve a company
* [GetEnrichConnectionIDCompany](#getenrichconnectionidcompany) - Retrieve enrichment information for a company
* [PatchCrmConnectionIDCompanyID](#patchcrmconnectionidcompanyid) - Update a company
* [PatchCrmConnectionIDCompanyIDDealDealID](#patchcrmconnectionidcompanyiddealdealid) - Associate a deal with a company
* [PostCrmConnectionIDCompany](#postcrmconnectionidcompany) - Create a company
* [PutCrmConnectionIDCompanyID](#putcrmconnectionidcompanyid) - Update a company
* [PutCrmConnectionIDCompanyIDDealDealID](#putcrmconnectionidcompanyiddealdealid) - Associate a deal with a company

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
        ConnectionID: "eos",
        ID: "b9c247c8-8373-4a40-a194-2f32e5505575",
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


## DeleteCrmConnectionIDCompanyIDDealDealID

Remove deal association from a company

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
    res, err := s.Company.DeleteCrmConnectionIDCompanyIDDealDealID(ctx, operations.DeleteCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "nisi",
        DealID: "tenetur",
        ID: "5d56d0bd-0af2-4dfe-93db-4f62cba3f894",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.DeleteCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/deletecrmconnectionidcompanyiddealdealidrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.DeleteCrmConnectionIDCompanyIDDealDealIDResponse](../../models/operations/deletecrmconnectionidcompanyiddealdealidresponse.md), error**


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
        ConnectionID: "quasi",
        ContactID: unifiedgosdk.String("mollitia"),
        DealID: unifiedgosdk.String("accusamus"),
        Limit: unifiedgosdk.Float64(6875.89),
        Offset: unifiedgosdk.Float64(7691.56),
        Order: unifiedgosdk.String("doloremque"),
        Query: unifiedgosdk.String("expedita"),
        Sort: unifiedgosdk.String("corrupti"),
        UpdatedGte: types.MustTimeFromString("2022-05-11T05:34:43.056Z"),
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
        ConnectionID: "aliquid",
        ID: "924d3b2e-cfcc-48f8-9501-0f5dd3d6fa18",
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
        ConnectionID: "aperiam",
        Domain: unifiedgosdk.String("non"),
        Name: unifiedgosdk.String("Derek Haag"),
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
                Address1: unifiedgosdk.String("consequuntur"),
                Address2: unifiedgosdk.String("maiores"),
                City: unifiedgosdk.String("South Jonathanchester"),
                Country: unifiedgosdk.String("Cyprus"),
                CountryCode: unifiedgosdk.String("HU"),
                PostalCode: unifiedgosdk.String("75542"),
                Region: unifiedgosdk.String("recusandae"),
                RegionCode: unifiedgosdk.String("tempora"),
            },
            CreatedAt: types.MustTimeFromString("2022-06-26T14:48:08.360Z"),
            DealIds: []string{
                "sequi",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Alejandrin94@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("b8ca275a-60a0-44c4-95cc-699171b51c1b"),
            Name: unifiedgosdk.String("Johnathan Braun"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "labore",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "expedita",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-11-30T08:27:41.789Z"),
            Websites: []string{
                "officiis",
            },
        },
        ConnectionID: "cum",
        ID: "dfc4ccca-99bc-47fc-8b2d-ce10873e42b0",
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


## PatchCrmConnectionIDCompanyIDDealDealID

Associate a deal with a company

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
    res, err := s.Company.PatchCrmConnectionIDCompanyIDDealDealID(ctx, operations.PatchCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "voluptatem",
        DealID: "eum",
        ID: "d678878b-a858-41a5-8208-c54fefa9c95f",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PatchCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/patchcrmconnectionidcompanyiddealdealidrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.PatchCrmConnectionIDCompanyIDDealDealIDResponse](../../models/operations/patchcrmconnectionidcompanyiddealdealidresponse.md), error**


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
                Address1: unifiedgosdk.String("quia"),
                Address2: unifiedgosdk.String("officiis"),
                City: unifiedgosdk.String("Runolfsdottirborough"),
                Country: unifiedgosdk.String("Gambia"),
                CountryCode: unifiedgosdk.String("IR"),
                PostalCode: unifiedgosdk.String("82047"),
                Region: unifiedgosdk.String("asperiores"),
                RegionCode: unifiedgosdk.String("recusandae"),
            },
            CreatedAt: types.MustTimeFromString("2021-06-25T15:44:39.144Z"),
            DealIds: []string{
                "dicta",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Albina.Hyatt53@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("3fa4a41c-480d-43f2-932a-f03102d514f4"),
            Name: unifiedgosdk.String("Loren Jakubowski IV"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "expedita",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "hic",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-11-06T17:51:41.197Z"),
            Websites: []string{
                "beatae",
            },
        },
        ConnectionID: "similique",
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
                Address1: unifiedgosdk.String("ea"),
                Address2: unifiedgosdk.String("animi"),
                City: unifiedgosdk.String("Fort Jeremie"),
                Country: unifiedgosdk.String("Kenya"),
                CountryCode: unifiedgosdk.String("NG"),
                PostalCode: unifiedgosdk.String("49928-2794"),
                Region: unifiedgosdk.String("ipsam"),
                RegionCode: unifiedgosdk.String("explicabo"),
            },
            CreatedAt: types.MustTimeFromString("2021-10-23T06:28:50.254Z"),
            DealIds: []string{
                "quis",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Cyril10@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("e3bb91c8-d975-4e0e-8419-d8f84f144f3e"),
            Name: unifiedgosdk.String("Joy Toy"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "cumque",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "aliquam",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-05-12T12:30:36.705Z"),
            Websites: []string{
                "reiciendis",
            },
        },
        ConnectionID: "sequi",
        ID: "cabd905a-972e-4056-b282-27b2d309470b",
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


## PutCrmConnectionIDCompanyIDDealDealID

Associate a deal with a company

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
    res, err := s.Company.PutCrmConnectionIDCompanyIDDealDealID(ctx, operations.PutCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "sapiente",
        DealID: "quam",
        ID: "a4fa87cf-535a-46fa-a54e-bf60c321f023",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PutCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/putcrmconnectionidcompanyiddealdealidrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.PutCrmConnectionIDCompanyIDDealDealIDResponse](../../models/operations/putcrmconnectionidcompanyiddealdealidresponse.md), error**

