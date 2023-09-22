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
        ConnectionID: "unde",
        ID: "7e152297-510d-4a80-b122-92cc61c2a702",
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
        CompanyID: "distinctio",
        ConnectionID: "soluta",
        ID: "97ee102d-a2de-435f-8e01-bf33eaab4540",
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
        ConnectionID: "dolores",
        ContactID: "dolorum",
        ID: "c1704bf1-cc9f-4c61-aae5-eb5f0c492b57",
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
        ConnectionID: "ut",
        DealID: "incidunt",
        ID: "d08a2267-aaee-479e-bc71-ad31becb83d2",
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
        CompanyID: unifiedgosdk.String("dolor"),
        ConnectionID: "esse",
        ContactID: unifiedgosdk.String("deleniti"),
        DealID: unifiedgosdk.String("mollitia"),
        Limit: unifiedgosdk.Float64(8941.65),
        Offset: unifiedgosdk.Float64(2035.85),
        Order: unifiedgosdk.String("facilis"),
        Query: unifiedgosdk.String("sapiente"),
        Sort: unifiedgosdk.String("maxime"),
        UpdatedGte: types.MustTimeFromString("2022-10-11T22:27:53.755Z"),
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
        ConnectionID: "nulla",
        ID: "9450a986-a495-4bac-b07f-06b28ecc8649",
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
                Description: unifiedgosdk.String("sunt"),
                Duration: unifiedgosdk.Float64(2261.97),
            },
            CompanyIds: []string{
                "laudantium",
            },
            ContactIds: []string{
                "commodi",
            },
            CreatedAt: types.MustTimeFromString("2021-10-18T23:17:50.935Z"),
            DealIds: []string{
                "qui",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("eligendi"),
                Cc: []string{
                    "perspiciatis",
                },
                From: unifiedgosdk.String("eum"),
                Subject: unifiedgosdk.String("sint"),
                To: []string{
                    "eligendi",
                },
            },
            ID: unifiedgosdk.String("4cc6b788-90a3-4fd3-881d-a10f8c23df93"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("quae"),
                EndAt: types.MustTimeFromString("2021-01-28T19:50:54.435Z"),
                StartAt: types.MustTimeFromString("2022-02-09T08:20:37.262Z"),
                Title: unifiedgosdk.String("Dr."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("tempore"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("minima"),
                Name: unifiedgosdk.String("Winifred O'Reilly"),
                Status: unifiedgosdk.String("numquam"),
            },
            Type: shared.CrmEventTypeMeeting.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2020-08-13T07:08:20.678Z"),
        },
        ConnectionID: "sint",
        ID: "43513772-6d15-4321-b832-a56d69180ff6",
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
        CompanyID: "consequatur",
        ConnectionID: "accusamus",
        ID: "b9a6658e-69a4-4b84-bd38-2dbec75c68c6",
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
        ConnectionID: "aut",
        ContactID: "nisi",
        ID: "59468ce3-04d8-4849-bf82-14c337f96bb0",
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
        ConnectionID: "porro",
        DealID: "vel",
        ID: "9e372db1-344b-4a9f-b8a5-c0ed7aab62e9",
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
                Description: unifiedgosdk.String("iusto"),
                Duration: unifiedgosdk.Float64(1323.66),
            },
            CompanyIds: []string{
                "ea",
            },
            ContactIds: []string{
                "architecto",
            },
            CreatedAt: types.MustTimeFromString("2020-10-13T11:08:25.283Z"),
            DealIds: []string{
                "alias",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("quod"),
                Cc: []string{
                    "veniam",
                },
                From: unifiedgosdk.String("corrupti"),
                Subject: unifiedgosdk.String("temporibus"),
                To: []string{
                    "odit",
                },
            },
            ID: unifiedgosdk.String("7b51996b-5b4b-450e-af71-2b7a7ab0344b"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("inventore"),
                EndAt: types.MustTimeFromString("2022-11-20T13:22:34.695Z"),
                StartAt: types.MustTimeFromString("2022-08-01T07:59:36.823Z"),
                Title: unifiedgosdk.String("Ms."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("deleniti"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("illum"),
                Name: unifiedgosdk.String("Caleb Purdy"),
                Status: unifiedgosdk.String("rem"),
            },
            Type: shared.CrmEventTypeTask.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-01-23T07:05:36.436Z"),
        },
        ConnectionID: "velit",
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
                Description: unifiedgosdk.String("fugiat"),
                Duration: unifiedgosdk.Float64(8660.78),
            },
            CompanyIds: []string{
                "voluptatem",
            },
            ContactIds: []string{
                "quod",
            },
            CreatedAt: types.MustTimeFromString("2020-05-30T09:15:09.119Z"),
            DealIds: []string{
                "dolor",
            },
            Email: &shared.PropertyCrmEventEmail{
                Body: unifiedgosdk.String("amet"),
                Cc: []string{
                    "tenetur",
                },
                From: unifiedgosdk.String("quasi"),
                Subject: unifiedgosdk.String("dicta"),
                To: []string{
                    "rerum",
                },
            },
            ID: unifiedgosdk.String("3e4e080a-a104-4186-ac75-9e02f3702c5c"),
            Meeting: &shared.PropertyCrmEventMeeting{
                Description: unifiedgosdk.String("laudantium"),
                EndAt: types.MustTimeFromString("2022-08-02T07:50:37.926Z"),
                StartAt: types.MustTimeFromString("2022-05-30T16:07:16.405Z"),
                Title: unifiedgosdk.String("Mr."),
            },
            Note: &shared.PropertyCrmEventNote{
                Description: unifiedgosdk.String("voluptates"),
            },
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{
                Description: unifiedgosdk.String("culpa"),
                Name: unifiedgosdk.String("Mrs. Leonard Cartwright"),
                Status: unifiedgosdk.String("culpa"),
            },
            Type: shared.CrmEventTypeEmail.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-07-25T01:13:49.601Z"),
        },
        ConnectionID: "alias",
        ID: "7bf375b4-4282-4821-bdb2-f69e59267c71",
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
        CompanyID: "quo",
        ConnectionID: "optio",
        ID: "8d3cd425-8d03-458a-82c8-08fe2751a204",
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
        ConnectionID: "ducimus",
        ContactID: "quod",
        ID: "0449e143-f961-49bb-bd40-d5a11fa436e6",
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
        ConnectionID: "sed",
        DealID: "exercitationem",
        ID: "9233f95c-9d23-4739-bc78-5b5db4f50018",
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

