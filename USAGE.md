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
	res, err := s.Accounting.CreateAccountingCustomer(ctx, operations.CreateAccountingCustomerRequest{
		AccountingCustomer: &shared.AccountingCustomer{
			BillingAddress: &shared.PropertyAccountingCustomerBillingAddress{},
			Emails: []shared.AccountingEmail{
				shared.AccountingEmail{
					Email: "Kevon_Schultz42@gmail.com",
				},
			},
			Raw:             &shared.PropertyAccountingCustomerRaw{},
			ShippingAddress: &shared.PropertyAccountingCustomerShippingAddress{},
			Telephones: []shared.AccountingTelephone{
				shared.AccountingTelephone{
					Telephone: "string",
				},
			},
		},
		ConnectionID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountingCustomer != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->