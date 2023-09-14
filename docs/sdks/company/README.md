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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.DeleteCrmConnectionIDCompanyID(ctx, operations.DeleteCrmConnectionIDCompanyIDRequest{
        ConnectionID: "maiores",
        ID: "eda53e5a-e6e0-4ac1-84c2-b9c247c88373",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteCrmConnectionIDCompanyIDRequest](../../models/operations/deletecrmconnectionidcompanyidrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.DeleteCrmConnectionIDCompanyIDSecurity](../../models/operations/deletecrmconnectionidcompanyidsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDCompanyIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.DeleteCrmConnectionIDCompanyIDDealDealID(ctx, operations.DeleteCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "laborum",
        DealID: "modi",
        ID: "0e1942f3-2e55-4055-b56f-5d56d0bd0af2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.DeleteCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/deletecrmconnectionidcompanyiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `security`                                                                                                                                 | [operations.DeleteCrmConnectionIDCompanyIDDealDealIDSecurity](../../models/operations/deletecrmconnectionidcompanyiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                         | The security requirements to use for the request.                                                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDCompanySecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.GetCrmConnectionIDCompany(ctx, operations.GetCrmConnectionIDCompanyRequest{
        ConnectionID: "possimus",
        ContactID: unifiedto.String("repellat"),
        DealID: unifiedto.String("repudiandae"),
        Limit: unifiedto.Float64(1002.33),
        Offset: unifiedto.Float64(2406.96),
        Order: unifiedto.String("pariatur"),
        Query: unifiedto.String("harum"),
        Sort: unifiedto.String("dolore"),
        UpdatedGte: types.MustTimeFromString("2021-09-11T06:54:38.386Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCrmConnectionIDCompanyRequest](../../models/operations/getcrmconnectionidcompanyrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.GetCrmConnectionIDCompanySecurity](../../models/operations/getcrmconnectionidcompanysecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.GetCrmConnectionIDCompanyID(ctx, operations.GetCrmConnectionIDCompanyIDRequest{
        ConnectionID: "explicabo",
        ID: "cba3f894-1aeb-4c0b-80a6-924d3b2ecfcc",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCrmConnectionIDCompanyIDRequest](../../models/operations/getcrmconnectionidcompanyidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetCrmConnectionIDCompanyIDSecurity](../../models/operations/getcrmconnectionidcompanyidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetEnrichConnectionIDCompanySecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.GetEnrichConnectionIDCompany(ctx, operations.GetEnrichConnectionIDCompanyRequest{
        ConnectionID: "quos",
        Domain: unifiedto.String("asperiores"),
        Name: unifiedto.String("Mr. Nick Hessel DVM"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetEnrichConnectionIDCompanyRequest](../../models/operations/getenrichconnectionidcompanyrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.GetEnrichConnectionIDCompanySecurity](../../models/operations/getenrichconnectionidcompanysecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.PatchCrmConnectionIDCompanyID(ctx, operations.PatchCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedto.String("nostrum"),
                Address2: unifiedto.String("at"),
                City: unifiedto.String("Coon Rapids"),
                Country: unifiedto.String("Syrian Arab Republic"),
                CountryCode: unifiedto.String("IO"),
                PostalCode: unifiedto.String("60502-9337"),
                Region: unifiedto.String("quas"),
                RegionCode: unifiedto.String("consequuntur"),
            },
            CreatedAt: types.MustTimeFromString("2022-10-03T16:02:53.877Z"),
            DealIds: []string{
                "aliquid",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Mandy.Ernser@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("8873e484-380b-41f6-b8ca-275a60a04c49"),
            Name: unifiedto.String("Lynne Schroeder"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "unde",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "illo",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-11T06:13:47.031Z"),
            Websites: []string{
                "ipsam",
            },
        },
        ConnectionID: "quasi",
        ID: "c1bdb1cf-4b88-48eb-9fc4-ccca99bc7fc0",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchCrmConnectionIDCompanyIDRequest](../../models/operations/patchcrmconnectionidcompanyidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PatchCrmConnectionIDCompanyIDSecurity](../../models/operations/patchcrmconnectionidcompanyidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDCompanyIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.PatchCrmConnectionIDCompanyIDDealDealID(ctx, operations.PatchCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "soluta",
        DealID: "fugit",
        ID: "dce10873-e42b-4006-9678-878ba8581a58",
    }, operationSecurity)
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
| `request`                                                                                                                                | [operations.PatchCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/patchcrmconnectionidcompanyiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.PatchCrmConnectionIDCompanyIDDealDealIDSecurity](../../models/operations/patchcrmconnectionidcompanyiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDCompanySecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.PostCrmConnectionIDCompany(ctx, operations.PostCrmConnectionIDCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedto.String("magni"),
                Address2: unifiedto.String("quae"),
                City: unifiedto.String("Randalborough"),
                Country: unifiedto.String("Fiji"),
                CountryCode: unifiedto.String("VG"),
                PostalCode: unifiedto.String("96676-3918"),
                Region: unifiedto.String("mollitia"),
                RegionCode: unifiedto.String("cumque"),
            },
            CreatedAt: types.MustTimeFromString("2022-09-08T12:56:45.780Z"),
            DealIds: []string{
                "eum",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Roslyn48@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("fee81206-e281-43fa-8a41-c480d3f2132a"),
            Name: unifiedto.String("Mr. Matthew Ebert"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "nulla",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "enim",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-01-14T22:16:14.005Z"),
            Websites: []string{
                "numquam",
            },
        },
        ConnectionID: "optio",
    }, operationSecurity)
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
| `request`                                                                                                      | [operations.PostCrmConnectionIDCompanyRequest](../../models/operations/postcrmconnectionidcompanyrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PostCrmConnectionIDCompanySecurity](../../models/operations/postcrmconnectionidcompanysecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.PutCrmConnectionIDCompanyID(ctx, operations.PutCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedto.String("nobis"),
                Address2: unifiedto.String("ex"),
                City: unifiedto.String("Beavercreek"),
                Country: unifiedto.String("Malaysia"),
                CountryCode: unifiedto.String("PK"),
                PostalCode: unifiedto.String("53116-4629"),
                Region: unifiedto.String("dignissimos"),
                RegionCode: unifiedto.String("esse"),
            },
            CreatedAt: types.MustTimeFromString("2021-12-22T21:12:12.654Z"),
            DealIds: []string{
                "esse",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Travon.Franecki70@hotmail.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("752c65b3-4418-4e3b-b91c-8d975e0e8419"),
            Name: unifiedto.String("Wallace Wiegand"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "earum",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "architecto",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-01-07T01:07:05.981Z"),
            Websites: []string{
                "sequi",
            },
        },
        ConnectionID: "saepe",
        ID: "07edcc4a-a5f3-4cab-9905-a972e0567282",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutCrmConnectionIDCompanyIDRequest](../../models/operations/putcrmconnectionidcompanyidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PutCrmConnectionIDCompanyIDSecurity](../../models/operations/putcrmconnectionidcompanyidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDCompanyIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Company.PutCrmConnectionIDCompanyIDDealDealID(ctx, operations.PutCrmConnectionIDCompanyIDDealDealIDRequest{
        ConnectionID: "odit",
        DealID: "iusto",
        ID: "b2d30947-0bf7-4a4f-a87c-f535a6fae54e",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PutCrmConnectionIDCompanyIDDealDealIDRequest](../../models/operations/putcrmconnectionidcompanyiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.PutCrmConnectionIDCompanyIDDealDealIDSecurity](../../models/operations/putcrmconnectionidcompanyiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.PutCrmConnectionIDCompanyIDDealDealIDResponse](../../models/operations/putcrmconnectionidcompanyiddealdealidresponse.md), error**

