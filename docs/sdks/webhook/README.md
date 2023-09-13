# Webhook

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteUnifiedWebhookIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Webhook.DeleteUnifiedWebhookID(ctx, operations.DeleteUnifiedWebhookIDRequest{
        ID: "d73809a0-2f06-4e92-a8b5-6065a5074bef",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteUnifiedWebhookIDRequest](../../models/operations/deleteunifiedwebhookidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.DeleteUnifiedWebhookIDSecurity](../../models/operations/deleteunifiedwebhookidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedWebhookSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Webhook.GetUnifiedWebhook(ctx, operations.GetUnifiedWebhookRequest{
        Env: unifiedto.String("cum"),
        Limit: unifiedto.Float64(5185.71),
        Object: unifiedto.String("laborum"),
        Offset: unifiedto.Float64(9427.54),
        Order: unifiedto.String("eum"),
        Sort: unifiedto.String("rem"),
        UpdatedGte: types.MustDateFromString("2022-05-29"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetUnifiedWebhookRequest](../../models/operations/getunifiedwebhookrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetUnifiedWebhookSecurity](../../models/operations/getunifiedwebhooksecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedWebhookIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Webhook.GetUnifiedWebhookID(ctx, operations.GetUnifiedWebhookIDRequest{
        ID: "d2b99404-363a-4096-8c05-3876e39def9c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetUnifiedWebhookIDRequest](../../models/operations/getunifiedwebhookidrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetUnifiedWebhookIDSecurity](../../models/operations/getunifiedwebhookidsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostUnifiedWebhookConnectionIDObjectSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Webhook.PostUnifiedWebhookConnectionIDObject(ctx, operations.PostUnifiedWebhookConnectionIDObjectRequest{
        Webhook: &shared.Webhook{
            CheckedAt: types.MustDateFromString("2022-08-13"),
            ConnectionID: "minima",
            CreatedAt: types.MustDateFromString("2020-02-03"),
            Environment: unifiedto.String("fugiat"),
            Events: []shared.PropertyWebhookEvents{
                shared.PropertyWebhookEventsUpdated,
            },
            HookURL: "ipsum",
            ID: unifiedto.String("54e5cb94-9770-417a-a620-4bb26ca4e999"),
            IncludeRaw: unifiedto.Bool(false),
            IntegrationType: "quos",
            Interval: 1768.7,
            ObjectType: shared.WebhookObjectTypeEnrichCompany,
            Subscriptions: []string{
                "iure",
            },
            UpdatedAt: types.MustDateFromString("2021-05-13"),
            WorkspaceID: "debitis",
        },
        ConnectionID: "reiciendis",
        Events: []PostUnifiedWebhookConnectionIDObjectEvents{
            operations.PostUnifiedWebhookConnectionIDObjectEventsCreated,
        },
        Object: "perferendis",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PostUnifiedWebhookConnectionIDObjectRequest](../../models/operations/postunifiedwebhookconnectionidobjectrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PostUnifiedWebhookConnectionIDObjectSecurity](../../models/operations/postunifiedwebhookconnectionidobjectsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


### Response

**[*operations.PostUnifiedWebhookConnectionIDObjectResponse](../../models/operations/postunifiedwebhookconnectionidobjectresponse.md), error**

