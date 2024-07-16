<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"os"
)

func main() {
	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity(os.Getenv("JWT")),
	)
	request := operations.CreateAccountingAccountRequest{
		ConnectionID: "<value>",
	}
	ctx := context.Background()
	res, err := s.Accounting.CreateAccountingAccount(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->