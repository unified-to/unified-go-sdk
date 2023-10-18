<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Apicall.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
		ID: "<ID>",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.APICall != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->