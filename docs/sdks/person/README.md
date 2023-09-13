# Person

### Available Operations

* [GetEnrichConnectionIDPerson](#getenrichconnectionidperson) - Retrieve enrichment information for a person

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
    res, err := s.Person.GetEnrichConnectionIDPerson(ctx, operations.GetEnrichConnectionIDPersonRequest{
        ConnectionID: "numquam",
        Email: unifiedto.String("Donna44@yahoo.com"),
        LinkedinURL: unifiedto.String("laboriosam"),
        Name: unifiedto.String("Phillip Waelchi"),
        Twitter: unifiedto.String("totam"),
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

