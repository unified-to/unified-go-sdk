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
			Currency:            unifiedgosdk.Pointer("BOB"),
			CustomerDefinedCode: unifiedgosdk.Pointer("quo"),
			Description:         unifiedgosdk.Pointer("Spoliatio comedo vilitas harum cupiditate."),
			ID:                  unifiedgosdk.Pointer("b096adb6-85c6-43bd-8b65-9908676ad8f8"),
			IsPayable:           unifiedgosdk.Pointer(true),
			Name:                unifiedgosdk.Pointer("Electronic Aluminum Tuna"),
			Status:              shared.StatusArchived.ToPointer(),
			Taxonomy: []shared.AccountingAccountTaxonomy{
				shared.AccountingAccountTaxonomy{
					OriginalType: unifiedgosdk.Pointer("vesper"),
					Type:         shared.AccountingAccountTaxonomyTypeSubgroup,
					Value:        "iste",
				},
				shared.AccountingAccountTaxonomy{
					OriginalType: unifiedgosdk.Pointer("adamo"),
					Type:         shared.AccountingAccountTaxonomyTypeSubgroup,
					Value:        "peccatus",
				},
			},
			Type:      shared.TypeBank.ToPointer(),
			UpdatedAt: types.MustNewTimeFromString("2023-01-05T23:34:54.694Z"),
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