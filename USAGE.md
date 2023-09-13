<!-- Start SDK Example Usage -->


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
    operationSecurity := operations.DeleteTicketingConnectionIDAgentIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Agent.DeleteTicketingConnectionIDAgentID(ctx, operations.DeleteTicketingConnectionIDAgentIDRequest{
        ConnectionID: "corrupti",
        ID: "9bd9d8d6-9a67-44e0-b467-cc8796ed151a",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->