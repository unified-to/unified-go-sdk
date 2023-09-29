# Event
(*Event*)

### Available Operations

* [DeleteCrmConnectionIDEventID](#deletecrmconnectionideventid) - Remove a event
* [DeleteCrmConnectionIDEventIDCompanyCompanyID](#deletecrmconnectionideventidcompanycompanyid) - Remove company association from an event
* [DeleteCrmConnectionIDEventIDContactContactID](#deletecrmconnectionideventidcontactcontactid) - Remove contact association from an event
* [DeleteCrmConnectionIDEventIDDealDealID](#deletecrmconnectionideventiddealdealid) - Remove deal association from an event
* [GetCrmConnectionIDEvent](#getcrmconnectionidevent) - List all events
* [GetCrmConnectionIDEventID](#getcrmconnectionideventid) - Retrieve a event
* [PatchCrmConnectionIDEventID](#patchcrmconnectionideventid) - Update a event
* [PatchCrmConnectionIDEventIDCompanyCompanyID](#patchcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PatchCrmConnectionIDEventIDContactContactID](#patchcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PatchCrmConnectionIDEventIDDealDealID](#patchcrmconnectionideventiddealdealid) - Associate a deal with an event
* [PostCrmConnectionIDEvent](#postcrmconnectionidevent) - Create a event
* [PutCrmConnectionIDEventID](#putcrmconnectionideventid) - Update a event
* [PutCrmConnectionIDEventIDCompanyCompanyID](#putcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PutCrmConnectionIDEventIDContactContactID](#putcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PutCrmConnectionIDEventIDDealDealID](#putcrmconnectionideventiddealdealid) - Associate a deal with an event

## DeleteCrmConnectionIDEventID

Remove a event

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
    res, err := s.Event.DeleteCrmConnectionIDEventID(ctx, operations.DeleteCrmConnectionIDEventIDRequest{
        ConnectionID: "Wooden Latin",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDEventIDRequest](../../models/operations/deletecrmconnectionideventidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.DeleteCrmConnectionIDEventIDResponse](../../models/operations/deletecrmconnectionideventidresponse.md), error**


## DeleteCrmConnectionIDEventIDCompanyCompanyID

Remove company association from an event

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
    res, err := s.Event.DeleteCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "Gasoline gentle Japan",
        ConnectionID: "Outdoors",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/deletecrmconnectionideventidcompanycompanyidrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDResponse](../../models/operations/deletecrmconnectionideventidcompanycompanyidresponse.md), error**


## DeleteCrmConnectionIDEventIDContactContactID

Remove contact association from an event

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
    res, err := s.Event.DeleteCrmConnectionIDEventIDContactContactID(ctx, operations.DeleteCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "lecture Funk",
        ContactID: "Bedfordshire Genderqueer",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.DeleteCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/deletecrmconnectionideventidcontactcontactidrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.DeleteCrmConnectionIDEventIDContactContactIDResponse](../../models/operations/deletecrmconnectionideventidcontactcontactidresponse.md), error**


## DeleteCrmConnectionIDEventIDDealDealID

Remove deal association from an event

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
    res, err := s.Event.DeleteCrmConnectionIDEventIDDealDealID(ctx, operations.DeleteCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "Southeast",
        DealID: "Granite gah Dysprosium",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.DeleteCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/deletecrmconnectionideventiddealdealidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.DeleteCrmConnectionIDEventIDDealDealIDResponse](../../models/operations/deletecrmconnectionideventiddealdealidresponse.md), error**


## GetCrmConnectionIDEvent

List all events

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
    res, err := s.Event.GetCrmConnectionIDEvent(ctx, operations.GetCrmConnectionIDEventRequest{
        CompanyID: unifiedgosdk.String("Zirconium Avon Bedfordshire"),
        ConnectionID: "Hybrid grey Ferrari",
        ContactID: unifiedgosdk.String("Checking Southeast"),
        DealID: unifiedgosdk.String("Graham till Caesium"),
        Limit: unifiedgosdk.Float64(2928.84),
        Offset: unifiedgosdk.Float64(5904.77),
        Order: unifiedgosdk.String("furthermore Tricycle Hop"),
        Query: unifiedgosdk.String("auxiliary"),
        Sort: unifiedgosdk.String("Southeast Bicycle Gorgeous"),
        UpdatedGte: types.MustTimeFromString("2023-01-15T23:49:53.643Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDEventRequest](../../models/operations/getcrmconnectionideventrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetCrmConnectionIDEventResponse](../../models/operations/getcrmconnectionideventresponse.md), error**


## GetCrmConnectionIDEventID

Retrieve a event

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
    res, err := s.Event.GetCrmConnectionIDEventID(ctx, operations.GetCrmConnectionIDEventIDRequest{
        ConnectionID: "Future equalise",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDEventIDRequest](../../models/operations/getcrmconnectionideventidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetCrmConnectionIDEventIDResponse](../../models/operations/getcrmconnectionideventidresponse.md), error**


## PatchCrmConnectionIDEventID

Update a event

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
    res, err := s.Event.PatchCrmConnectionIDEventID(ctx, operations.PatchCrmConnectionIDEventIDRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.String("Optional zero defect function"),
                Duration: unifiedgosdk.Float64(5434.61),
            },
            CompanyIds: []string{
                "silver",
            },
            ContactIds: []string{
                "redefine",
            },
            CreatedAt: types.MustTimeFromString("2021-07-21T06:46:42.528Z"),
            DealIds: []string{
                "Solutions",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("French"),
                Cc: []string{
                    "Checking",
                },
                From: unifiedgosdk.String("SDD Toyota Northeast"),
                Subject: unifiedgosdk.String("Convertible"),
                To: []string{
                    "Electronics",
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("Monitored mission-critical customer loyalty"),
                EndAt: types.MustTimeFromString("2022-09-22T17:43:00.863Z"),
                StartAt: types.MustTimeFromString("2023-04-24T06:40:04.926Z"),
                Title: unifiedgosdk.String("Kip Switchable Chicken"),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("Cross-group high-level functionalities"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("Horizontal empowering forecast"),
                Name: unifiedgosdk.String("Principal extremely Jast"),
                Status: unifiedgosdk.String("striped Concrete Bronze"),
            },
            Type: shared.CrmEventTypeNote.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-02-18T21:34:24.992Z"),
        },
        ConnectionID: "Dinar benchmark till",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDEventIDRequest](../../models/operations/patchcrmconnectionideventidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PatchCrmConnectionIDEventIDResponse](../../models/operations/patchcrmconnectionideventidresponse.md), error**


## PatchCrmConnectionIDEventIDCompanyCompanyID

Associate a company with an event

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
    res, err := s.Event.PatchCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.PatchCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "quis Operations",
        ConnectionID: "convergence programming",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PatchCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/patchcrmconnectionideventidcompanycompanyidrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PatchCrmConnectionIDEventIDCompanyCompanyIDResponse](../../models/operations/patchcrmconnectionideventidcompanycompanyidresponse.md), error**


## PatchCrmConnectionIDEventIDContactContactID

Associate a contact with an event

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
    res, err := s.Event.PatchCrmConnectionIDEventIDContactContactID(ctx, operations.PatchCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "gosh",
        ContactID: "North Guilder",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PatchCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/patchcrmconnectionideventidcontactcontactidrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PatchCrmConnectionIDEventIDContactContactIDResponse](../../models/operations/patchcrmconnectionideventidcontactcontactidresponse.md), error**


## PatchCrmConnectionIDEventIDDealDealID

Associate a deal with an event

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
    res, err := s.Event.PatchCrmConnectionIDEventIDDealDealID(ctx, operations.PatchCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "nobis Developer withdrawal",
        DealID: "male connect",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PatchCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/patchcrmconnectionideventiddealdealidrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.PatchCrmConnectionIDEventIDDealDealIDResponse](../../models/operations/patchcrmconnectionideventiddealdealidresponse.md), error**


## PostCrmConnectionIDEvent

Create a event

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
    res, err := s.Event.PostCrmConnectionIDEvent(ctx, operations.PostCrmConnectionIDEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.String("Visionary bandwidth-monitored hardware"),
                Duration: unifiedgosdk.Float64(9256.02),
            },
            CompanyIds: []string{
                "Kentucky",
            },
            ContactIds: []string{
                "Rustic",
            },
            CreatedAt: types.MustTimeFromString("2023-02-12T10:03:55.861Z"),
            DealIds: []string{
                "agonizing",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("protocol"),
                Cc: []string{
                    "Ratke",
                },
                From: unifiedgosdk.String("woman"),
                Subject: unifiedgosdk.String("East Soft"),
                To: []string{
                    "Southeast",
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("Streamlined intangible time-frame"),
                EndAt: types.MustTimeFromString("2022-04-18T21:50:55.608Z"),
                StartAt: types.MustTimeFromString("2021-08-24T14:06:25.626Z"),
                Title: unifiedgosdk.String("violet Synergized blah"),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("Mandatory eco-centric toolset"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("Team-oriented dynamic forecast"),
                Name: unifiedgosdk.String("Grocery"),
                Status: unifiedgosdk.String("excitedly Bacon"),
            },
            Type: shared.CrmEventTypeEmail.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-09-09T20:12:06.214Z"),
        },
        ConnectionID: "Progressive",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostCrmConnectionIDEventRequest](../../models/operations/postcrmconnectionideventrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PostCrmConnectionIDEventResponse](../../models/operations/postcrmconnectionideventresponse.md), error**


## PutCrmConnectionIDEventID

Update a event

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
    res, err := s.Event.PutCrmConnectionIDEventID(ctx, operations.PutCrmConnectionIDEventIDRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.String("Re-engineered composite time-frame"),
                Duration: unifiedgosdk.Float64(5444.58),
            },
            CompanyIds: []string{
                "DNS",
            },
            ContactIds: []string{
                "Skokie",
            },
            CreatedAt: types.MustTimeFromString("2022-07-05T01:37:36.877Z"),
            DealIds: []string{
                "lux",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("Hatchback card"),
                Cc: []string{
                    "Gasoline",
                },
                From: unifiedgosdk.String("Caribbean"),
                Subject: unifiedgosdk.String("Account medium"),
                To: []string{
                    "copy",
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("Inverse optimizing model"),
                EndAt: types.MustTimeFromString("2022-03-21T17:32:41.888Z"),
                StartAt: types.MustTimeFromString("2022-10-17T10:31:48.119Z"),
                Title: unifiedgosdk.String("male henry Hat"),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("Self-enabling asymmetric functionalities"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("Reduced 4th generation analyzer"),
                Name: unifiedgosdk.String("Savings Female nor"),
                Status: unifiedgosdk.String("Customer sky"),
            },
            Type: shared.CrmEventTypeNote.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2023-07-27T14:02:37.510Z"),
        },
        ConnectionID: "transparent",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDEventIDRequest](../../models/operations/putcrmconnectionideventidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PutCrmConnectionIDEventIDResponse](../../models/operations/putcrmconnectionideventidresponse.md), error**


## PutCrmConnectionIDEventIDCompanyCompanyID

Associate a company with an event

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
    res, err := s.Event.PutCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.PutCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "Administrator",
        ConnectionID: "Computers",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PutCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/putcrmconnectionideventidcompanycompanyidrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.PutCrmConnectionIDEventIDCompanyCompanyIDResponse](../../models/operations/putcrmconnectionideventidcompanycompanyidresponse.md), error**


## PutCrmConnectionIDEventIDContactContactID

Associate a contact with an event

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
    res, err := s.Event.PutCrmConnectionIDEventIDContactContactID(ctx, operations.PutCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "olive",
        ContactID: "deposit Bicycle",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PutCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/putcrmconnectionideventidcontactcontactidrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.PutCrmConnectionIDEventIDContactContactIDResponse](../../models/operations/putcrmconnectionideventidcontactcontactidresponse.md), error**


## PutCrmConnectionIDEventIDDealDealID

Associate a deal with an event

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
    res, err := s.Event.PutCrmConnectionIDEventIDDealDealID(ctx, operations.PutCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "Bloomington",
        DealID: "provided apud Southeast",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PutCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/putcrmconnectionideventiddealdealidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.PutCrmConnectionIDEventIDDealDealIDResponse](../../models/operations/putcrmconnectionideventiddealdealidresponse.md), error**

