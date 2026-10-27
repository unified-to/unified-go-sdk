<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
		AccountingAccount: shared.AccountingAccount{
			Balance:             unifiedgosdk.Pointer[float64](12092.0),
			CreatedAt:           types.MustNewTimeFromString("2022-07-03T17:57:07.391Z"),
			Currency:            unifiedgosdk.Pointer("CAD"),
			CustomerDefinedCode: unifiedgosdk.Pointer("6170"),
			Description:         unifiedgosdk.Pointer("Voluptatibus valde spoliatio."),
			ID:                  unifiedgosdk.Pointer("d20e9e79-d841-4e65-9472-9328ecda8c07"),
			IsPayable:           unifiedgosdk.Pointer(false),
			Metadata:            []shared.AccountingMetadata{},
			Name:                unifiedgosdk.Pointer("Auto Loan Account"),
			Status:              shared.StatusArchived.ToPointer(),
			Taxonomy: []shared.AccountingAccountTaxonomy{
				shared.AccountingAccountTaxonomy{
					OriginalType: unifiedgosdk.Pointer("vilitas"),
					Type:         shared.AccountingAccountTaxonomyTypeSubgroup,
					Value:        "comedo",
				},
			},
			Type:      shared.TypeAccountsReceivable.ToPointer(),
			UpdatedAt: types.MustNewTimeFromString("2023-08-08T09:33:13.693Z"),
		},
		ConnectionID: "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->