# Webhook
(*Webhook*)

### Available Operations

* [DeleteUnifiedWebhookID](#deleteunifiedwebhookid) - Remove webhook subscription
* [GetUnifiedWebhook](#getunifiedwebhook) - Returns all registered webhooks
* [GetUnifiedWebhookID](#getunifiedwebhookid) - Retrieve webhook by its ID
* [PostUnifiedWebhookConnectionIDObject](#postunifiedwebhookconnectionidobject) - Create webhook subscription

## DeleteUnifiedWebhookID

Remove webhook subscription

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
    res, err := s.Webhook.DeleteUnifiedWebhookID(ctx, operations.DeleteUnifiedWebhookIDRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteUnifiedWebhookIDRequest](../../models/operations/deleteunifiedwebhookidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteUnifiedWebhookIDResponse](../../models/operations/deleteunifiedwebhookidresponse.md), error**


## GetUnifiedWebhook

Returns all registered webhooks

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.GetUnifiedWebhook(ctx, operations.GetUnifiedWebhookRequest{
        Env: unifiedgosdk.String("Investor methodical Fitness"),
        Limit: unifiedgosdk.Float64(8087.22),
        Object: unifiedgosdk.String("Franc past salmon"),
        Offset: unifiedgosdk.Float64(5240.75),
        Order: unifiedgosdk.String("program"),
        Sort: unifiedgosdk.String("below JSON"),
        UpdatedGte: types.MustTimeFromString("2022-05-29T13:22:55.562Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetUnifiedWebhookRequest](../../models/operations/getunifiedwebhookrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetUnifiedWebhookResponse](../../models/operations/getunifiedwebhookresponse.md), error**


## GetUnifiedWebhookID

Retrieve webhook by its ID

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
    res, err := s.Webhook.GetUnifiedWebhookID(ctx, operations.GetUnifiedWebhookIDRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetUnifiedWebhookIDRequest](../../models/operations/getunifiedwebhookidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedWebhookIDResponse](../../models/operations/getunifiedwebhookidresponse.md), error**


## PostUnifiedWebhookConnectionIDObject

To maintain compatibility with the webhooks specification and Zapier webhooks, only the hook_url field is required in the payload when creating a Webhook.  When updated/new objects are found, the array of objects will be POSTed to the hook_url with the paramater hookId=<hookId>.

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.PostUnifiedWebhookConnectionIDObject(ctx, operations.PostUnifiedWebhookConnectionIDObjectRequest{
        Webhook: &shared.Webhook{
            CheckedAt: types.MustTimeFromString("2021-02-25T07:12:08.980Z"),
            ConnectionID: "deposit 1080p Passenger",
            CreatedAt: types.MustTimeFromString("2023-02-21T14:58:56.193Z"),
            Environment: unifiedgosdk.String("Minnesota Soap"),
            Events: []shared.PropertyWebhookEvents{
                shared.PropertyWebhookEventsUpdated,
            },
            HookURL: "Table female ken",
            ID: unifiedgosdk.String("<ID>"),
            IncludeRaw: unifiedgosdk.Bool(false),
            IntegrationType: "chocolate",
            Interval: 1710.16,
            ObjectType: shared.WebhookObjectTypeEnrichCompany,
            Subscriptions: []string{
                "female",
            },
            UpdatedAt: types.MustTimeFromString("2022-08-02T17:13:06.397Z"),
            WorkspaceID: "hertz",
        },
        ConnectionID: "Borders",
        Events: []PostUnifiedWebhookConnectionIDObjectEvents{
            operations.PostUnifiedWebhookConnectionIDObjectEventsCreated,
        },
        Object: "scalable",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PostUnifiedWebhookConnectionIDObjectRequest](../../models/operations/postunifiedwebhookconnectionidobjectrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PostUnifiedWebhookConnectionIDObjectResponse](../../models/operations/postunifiedwebhookconnectionidobjectresponse.md), error**

