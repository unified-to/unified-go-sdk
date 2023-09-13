# Enrich

### Available Operations

* [GetEnrichConnectionIDCompany](#getenrichconnectionidcompany) - Retrieve enrichment information for a company
* [GetEnrichConnectionIDPerson](#getenrichconnectionidperson) - Retrieve enrichment information for a person

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
    res, err := s.Enrich.GetEnrichConnectionIDCompany(ctx, operations.GetEnrichConnectionIDCompanyRequest{
        ConnectionID: "mollitia",
        Domain: unifiedto.String("beatae"),
        Name: unifiedto.String("Irma Bayer"),
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


## GetEnrichConnectionIDPerson

Retrieve enrichment information for a person

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
    operationSecurity := operations.GetEnrichConnectionIDPersonSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Enrich.GetEnrichConnectionIDPerson(ctx, operations.GetEnrichConnectionIDPersonRequest{
        ConnectionID: "at",
        Email: unifiedto.String("Curtis.Barrows40@hotmail.com"),
        LinkedinURL: unifiedto.String("cupiditate"),
        Name: unifiedto.String("Ms. Ivan Breitenberg IV"),
        Twitter: unifiedto.String("temporibus"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichPerson != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetEnrichConnectionIDPersonRequest](../../models/operations/getenrichconnectionidpersonrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetEnrichConnectionIDPersonSecurity](../../models/operations/getenrichconnectionidpersonsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.GetEnrichConnectionIDPersonResponse](../../models/operations/getenrichconnectionidpersonresponse.md), error**

