# Person
(*Person*)

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
    res, err := s.Person.GetEnrichConnectionIDPerson(ctx, operations.GetEnrichConnectionIDPersonRequest{
        ConnectionID: "Iowa Account",
        Email: unifiedgosdk.String("Jaiden_Weimann24@gmail.com"),
        LinkedinURL: unifiedgosdk.String("paradigms integrate Creative"),
        Name: unifiedgosdk.String("Investment"),
        Twitter: unifiedgosdk.String("Hills"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichPerson != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetEnrichConnectionIDPersonRequest](../../models/operations/getenrichconnectionidpersonrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetEnrichConnectionIDPersonResponse](../../models/operations/getenrichconnectionidpersonresponse.md), error**
