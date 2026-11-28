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
			Description:         unifiedgosdk.Pointer("Comedo vilitas harum cupiditate animus quo iste."),
			ID:                  unifiedgosdk.Pointer("431d14b1-6945-465e-92c5-4d9d4c75f6f0"),
			IsPayable:           unifiedgosdk.Pointer(false),
			Metadata:            []shared.AccountingMetadata{},
			Name:                unifiedgosdk.Pointer("Savings Account"),
			Status:              shared.StatusArchived.ToPointer(),
			Taxonomy: []shared.AccountingAccountTaxonomy{
				shared.AccountingAccountTaxonomy{
					OriginalType: unifiedgosdk.Pointer("peccatus"),
					Type:         shared.AccountingAccountTaxonomyTypeSubgroup,
					Value:        "vesper",
				},
			},
			Type:      shared.TypeAccountsReceivable.ToPointer(),
			UpdatedAt: types.MustNewTimeFromString("2023-08-16T13:33:48.928Z"),
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