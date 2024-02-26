<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
	s := unifiedgosdk.New()

	operationSecurity := operations.CreateAccountingAccountSecurity{
		Jwt: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
		ConnectionID: "<value>",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->