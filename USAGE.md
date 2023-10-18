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
	res, err := s.Agent.CreateTicketingAgent(ctx, operations.CreateTicketingAgentRequest{
		TicketingAgent: &shared.TicketingAgent{
			Emails: []shared.TicketingEmail{
				shared.TicketingEmail{
					Email: "Paolo.Cole8@yahoo.com",
				},
			},
			Raw: shared.PropertyTicketingAgentRaw{},
			Telephones: []shared.TicketingTelephone{
				shared.TicketingTelephone{
					Telephone: "Seaborgium",
				},
			},
		},
		ConnectionID: "Manager",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.TicketingAgent != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->