# Event

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDEventIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.DeleteCrmConnectionIDEventID(ctx, operations.DeleteCrmConnectionIDEventIDRequest{
        ConnectionID: "unde",
        ID: "7e152297-510d-4a80-b122-92cc61c2a702",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteCrmConnectionIDEventIDRequest](../../models/operations/deletecrmconnectionideventidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.DeleteCrmConnectionIDEventIDSecurity](../../models/operations/deletecrmconnectionideventidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.DeleteCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "distinctio",
        ConnectionID: "soluta",
        ID: "97ee102d-a2de-435f-8e01-bf33eaab4540",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/deletecrmconnectionideventidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `security`                                                                                                                                         | [operations.DeleteCrmConnectionIDEventIDCompanyCompanyIDSecurity](../../models/operations/deletecrmconnectionideventidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                                 | The security requirements to use for the request.                                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDEventIDContactContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.DeleteCrmConnectionIDEventIDContactContactID(ctx, operations.DeleteCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "dolores",
        ContactID: "dolorum",
        ID: "c1704bf1-cc9f-4c61-aae5-eb5f0c492b57",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/deletecrmconnectionideventidcontactcontactidrequest.md)   | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `security`                                                                                                                                         | [operations.DeleteCrmConnectionIDEventIDContactContactIDSecurity](../../models/operations/deletecrmconnectionideventidcontactcontactidsecurity.md) | :heavy_check_mark:                                                                                                                                 | The security requirements to use for the request.                                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDEventIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.DeleteCrmConnectionIDEventIDDealDealID(ctx, operations.DeleteCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "ut",
        DealID: "incidunt",
        ID: "d08a2267-aaee-479e-bc71-ad31becb83d2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.DeleteCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/deletecrmconnectionideventiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.DeleteCrmConnectionIDEventIDDealDealIDSecurity](../../models/operations/deletecrmconnectionideventiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDEventSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.GetCrmConnectionIDEvent(ctx, operations.GetCrmConnectionIDEventRequest{
        CompanyID: unifiedto.String("dolor"),
        ConnectionID: "esse",
        ContactID: unifiedto.String("deleniti"),
        DealID: unifiedto.String("mollitia"),
        Limit: unifiedto.Float64(8941.65),
        Offset: unifiedto.Float64(2035.85),
        Order: unifiedto.String("facilis"),
        Query: unifiedto.String("sapiente"),
        Sort: unifiedto.String("maxime"),
        UpdatedGte: types.MustDateFromString("2022-10-11"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDEventRequest](../../models/operations/getcrmconnectionideventrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetCrmConnectionIDEventSecurity](../../models/operations/getcrmconnectionideventsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDEventIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.GetCrmConnectionIDEventID(ctx, operations.GetCrmConnectionIDEventIDRequest{
        ConnectionID: "nulla",
        ID: "9450a986-a495-4bac-b07f-06b28ecc8649",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCrmConnectionIDEventIDRequest](../../models/operations/getcrmconnectionideventidrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.GetCrmConnectionIDEventIDSecurity](../../models/operations/getcrmconnectionideventidsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDEventIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PatchCrmConnectionIDEventID(ctx, operations.PatchCrmConnectionIDEventIDRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedto.String("sunt"),
                Duration: unifiedto.Float64(2261.97),
            },
            CompanyIds: []string{
                "laudantium",
            },
            ContactIds: []string{
                "commodi",
            },
            CreatedAt: types.MustDateFromString("2021-10-18"),
            DealIds: []string{
                "qui",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedto.String("eligendi"),
                Cc: []string{
                    "perspiciatis",
                },
                From: unifiedto.String("eum"),
                Subject: unifiedto.String("sint"),
                To: []string{
                    "eligendi",
                },
            },
            ID: unifiedto.String("4cc6b788-90a3-4fd3-881d-a10f8c23df93"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedto.String("quae"),
                EndAt: types.MustDateFromString("2021-01-28"),
                StartAt: types.MustDateFromString("2022-02-09"),
                Title: unifiedto.String("Dr."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedto.String("tempore"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedto.String("minima"),
                Name: unifiedto.String("Winifred O'Reilly"),
                Status: unifiedto.String("numquam"),
            },
            Type: shared.CrmEventTypeMeeting.ToPointer(),
            UpdatedAt: types.MustDateFromString("2020-08-13"),
        },
        ConnectionID: "sint",
        ID: "43513772-6d15-4321-b832-a56d69180ff6",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchCrmConnectionIDEventIDRequest](../../models/operations/patchcrmconnectionideventidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PatchCrmConnectionIDEventIDSecurity](../../models/operations/patchcrmconnectionideventidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDEventIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PatchCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.PatchCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "consequatur",
        ConnectionID: "accusamus",
        ID: "b9a6658e-69a4-4b84-bd38-2dbec75c68c6",
    }, operationSecurity)
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
| `request`                                                                                                                                        | [operations.PatchCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/patchcrmconnectionideventidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `security`                                                                                                                                       | [operations.PatchCrmConnectionIDEventIDCompanyCompanyIDSecurity](../../models/operations/patchcrmconnectionideventidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDEventIDContactContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PatchCrmConnectionIDEventIDContactContactID(ctx, operations.PatchCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "aut",
        ContactID: "nisi",
        ID: "59468ce3-04d8-4849-bf82-14c337f96bb0",
    }, operationSecurity)
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
| `request`                                                                                                                                        | [operations.PatchCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/patchcrmconnectionideventidcontactcontactidrequest.md)   | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `security`                                                                                                                                       | [operations.PatchCrmConnectionIDEventIDContactContactIDSecurity](../../models/operations/patchcrmconnectionideventidcontactcontactidsecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDEventIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PatchCrmConnectionIDEventIDDealDealID(ctx, operations.PatchCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "porro",
        DealID: "vel",
        ID: "9e372db1-344b-4a9f-b8a5-c0ed7aab62e9",
    }, operationSecurity)
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
| `request`                                                                                                                            | [operations.PatchCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/patchcrmconnectionideventiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.PatchCrmConnectionIDEventIDDealDealIDSecurity](../../models/operations/patchcrmconnectionideventiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDEventSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PostCrmConnectionIDEvent(ctx, operations.PostCrmConnectionIDEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedto.String("iusto"),
                Duration: unifiedto.Float64(1323.66),
            },
            CompanyIds: []string{
                "ea",
            },
            ContactIds: []string{
                "architecto",
            },
            CreatedAt: types.MustDateFromString("2020-10-13"),
            DealIds: []string{
                "alias",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedto.String("quod"),
                Cc: []string{
                    "veniam",
                },
                From: unifiedto.String("corrupti"),
                Subject: unifiedto.String("temporibus"),
                To: []string{
                    "odit",
                },
            },
            ID: unifiedto.String("7b51996b-5b4b-450e-af71-2b7a7ab0344b"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedto.String("inventore"),
                EndAt: types.MustDateFromString("2022-11-20"),
                StartAt: types.MustDateFromString("2022-08-01"),
                Title: unifiedto.String("Ms."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedto.String("deleniti"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedto.String("illum"),
                Name: unifiedto.String("Caleb Purdy"),
                Status: unifiedto.String("rem"),
            },
            Type: shared.CrmEventTypeTask.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-01-23"),
        },
        ConnectionID: "velit",
    }, operationSecurity)
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
| `request`                                                                                                  | [operations.PostCrmConnectionIDEventRequest](../../models/operations/postcrmconnectionideventrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PostCrmConnectionIDEventSecurity](../../models/operations/postcrmconnectionideventsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDEventIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PutCrmConnectionIDEventID(ctx, operations.PutCrmConnectionIDEventIDRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedto.String("fugiat"),
                Duration: unifiedto.Float64(8660.78),
            },
            CompanyIds: []string{
                "voluptatem",
            },
            ContactIds: []string{
                "quod",
            },
            CreatedAt: types.MustDateFromString("2020-05-30"),
            DealIds: []string{
                "dolor",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedto.String("amet"),
                Cc: []string{
                    "tenetur",
                },
                From: unifiedto.String("quasi"),
                Subject: unifiedto.String("dicta"),
                To: []string{
                    "rerum",
                },
            },
            ID: unifiedto.String("3e4e080a-a104-4186-ac75-9e02f3702c5c"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedto.String("laudantium"),
                EndAt: types.MustDateFromString("2022-08-02"),
                StartAt: types.MustDateFromString("2022-05-30"),
                Title: unifiedto.String("Mr."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedto.String("voluptates"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedto.String("culpa"),
                Name: unifiedto.String("Mrs. Leonard Cartwright"),
                Status: unifiedto.String("culpa"),
            },
            Type: shared.CrmEventTypeEmail.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-07-25"),
        },
        ConnectionID: "alias",
        ID: "7bf375b4-4282-4821-bdb2-f69e59267c71",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutCrmConnectionIDEventIDRequest](../../models/operations/putcrmconnectionideventidrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PutCrmConnectionIDEventIDSecurity](../../models/operations/putcrmconnectionideventidsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDEventIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PutCrmConnectionIDEventIDCompanyCompanyID(ctx, operations.PutCrmConnectionIDEventIDCompanyCompanyIDRequest{
        CompanyID: "quo",
        ConnectionID: "optio",
        ID: "8d3cd425-8d03-458a-82c8-08fe2751a204",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutCrmConnectionIDEventIDCompanyCompanyIDRequest](../../models/operations/putcrmconnectionideventidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.PutCrmConnectionIDEventIDCompanyCompanyIDSecurity](../../models/operations/putcrmconnectionideventidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDEventIDContactContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PutCrmConnectionIDEventIDContactContactID(ctx, operations.PutCrmConnectionIDEventIDContactContactIDRequest{
        ConnectionID: "ducimus",
        ContactID: "quod",
        ID: "0449e143-f961-49bb-bd40-d5a11fa436e6",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutCrmConnectionIDEventIDContactContactIDRequest](../../models/operations/putcrmconnectionideventidcontactcontactidrequest.md)   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.PutCrmConnectionIDEventIDContactContactIDSecurity](../../models/operations/putcrmconnectionideventidcontactcontactidsecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDEventIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Event.PutCrmConnectionIDEventIDDealDealID(ctx, operations.PutCrmConnectionIDEventIDDealDealIDRequest{
        ConnectionID: "sed",
        DealID: "exercitationem",
        ID: "9233f95c-9d23-4739-bc78-5b5db4f50018",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PutCrmConnectionIDEventIDDealDealIDRequest](../../models/operations/putcrmconnectionideventiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `security`                                                                                                                       | [operations.PutCrmConnectionIDEventIDDealDealIDSecurity](../../models/operations/putcrmconnectionideventiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |


### Response

**[*operations.PutCrmConnectionIDEventIDDealDealIDResponse](../../models/operations/putcrmconnectionideventiddealdealidresponse.md), error**

