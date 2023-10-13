# Crm
(*Crm*)

### Available Operations

* [CreateCrmCompany](#createcrmcompany) - Create a company
* [CreateCrmContact](#createcrmcontact) - Create a contact
* [CreateCrmDeal](#createcrmdeal) - Create a deal
* [CreateCrmEvent](#createcrmevent) - Create a event
* [CreateCrmFile](#createcrmfile) - Create a file
* [CreateCrmLead](#createcrmlead) - Create a lead
* [CreateCrmPipeline](#createcrmpipeline) - Create a pipeline
* [CreateCrmTeam](#createcrmteam) - Create a team
* [CreateCrmUser](#createcrmuser) - Create a user
* [GetCrmCompany](#getcrmcompany) - Retrieve a company
* [GetCrmContact](#getcrmcontact) - Retrieve a contact
* [GetCrmDeal](#getcrmdeal) - Retrieve a deal
* [GetCrmEvent](#getcrmevent) - Retrieve a event
* [GetCrmFile](#getcrmfile) - Retrieve a file
* [GetCrmLead](#getcrmlead) - Retrieve a lead
* [GetCrmPipeline](#getcrmpipeline) - Retrieve a pipeline
* [GetCrmTeam](#getcrmteam) - Retrieve a team
* [GetCrmUser](#getcrmuser) - Retrieve a user
* [ListCrmCompanies](#listcrmcompanies) - List all companies
* [ListCrmContacts](#listcrmcontacts) - List all contacts
* [ListCrmDeals](#listcrmdeals) - List all deals
* [ListCrmEvents](#listcrmevents) - List all events
* [ListCrmFiles](#listcrmfiles) - List all files
* [ListCrmLeads](#listcrmleads) - List all leads
* [ListCrmPipelines](#listcrmpipelines) - List all pipelines
* [ListCrmTeams](#listcrmteams) - List all teams
* [ListCrmUsers](#listcrmusers) - List all users
* [PatchCrmCompany](#patchcrmcompany) - Update a company
* [PatchCrmContact](#patchcrmcontact) - Update a contact
* [PatchCrmDeal](#patchcrmdeal) - Update a deal
* [PatchCrmEvent](#patchcrmevent) - Update a event
* [PatchCrmFile](#patchcrmfile) - Update a file
* [PatchCrmLead](#patchcrmlead) - Update a lead
* [PatchCrmPipeline](#patchcrmpipeline) - Update a pipeline
* [PatchCrmTeam](#patchcrmteam) - Update a team
* [PatchCrmUser](#patchcrmuser) - Update a user
* [RemoveCrmCompany](#removecrmcompany) - Remove a company
* [RemoveCrmContact](#removecrmcontact) - Remove a contact
* [RemoveCrmDeal](#removecrmdeal) - Remove a deal
* [RemoveCrmEvent](#removecrmevent) - Remove a event
* [RemoveCrmFile](#removecrmfile) - Remove a file
* [RemoveCrmLead](#removecrmlead) - Remove a lead
* [RemoveCrmPipeline](#removecrmpipeline) - Remove a pipeline
* [RemoveCrmTeam](#removecrmteam) - Remove a team
* [RemoveCrmUser](#removecrmuser) - Remove a user
* [UpdateCrmCompany](#updatecrmcompany) - Update a company
* [UpdateCrmContact](#updatecrmcontact) - Update a contact
* [UpdateCrmDeal](#updatecrmdeal) - Update a deal
* [UpdateCrmEvent](#updatecrmevent) - Update a event
* [UpdateCrmFile](#updatecrmfile) - Update a file
* [UpdateCrmLead](#updatecrmlead) - Update a lead
* [UpdateCrmPipeline](#updatecrmpipeline) - Update a pipeline
* [UpdateCrmTeam](#updatecrmteam) - Update a team
* [UpdateCrmUser](#updatecrmuser) - Update a user

## CreateCrmCompany

Create a company

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmCompany(ctx, operations.CreateCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "connecting",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "carouse",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Germany",
                },
            },
            Websites: []string{
                "yippee",
            },
        },
        ConnectionID: "magenta Data woot",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCrmCompanyRequest](../../models/operations/createcrmcompanyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateCrmCompanyResponse](../../models/operations/createcrmcompanyresponse.md), error**


## CreateCrmContact

Create a contact

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmContact(ctx, operations.CreateCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "Mendelevium",
            },
            DealIds: []string{
                "Account",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "towards",
                },
            },
        },
        ConnectionID: "Cambridgeshire Passenger Producer",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCrmContactRequest](../../models/operations/createcrmcontactrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateCrmContactResponse](../../models/operations/createcrmcontactresponse.md), error**


## CreateCrmDeal

Create a deal

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmDeal(ctx, operations.CreateCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "Toys",
            },
        },
        ConnectionID: "Music Rap",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmDealRequest](../../models/operations/createcrmdealrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmDealResponse](../../models/operations/createcrmdealresponse.md), error**


## CreateCrmEvent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmEvent(ctx, operations.CreateCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "repeatedly",
            },
            ContactIds: []string{
                "Sedan",
            },
            DealIds: []string{
                "altruistic",
            },
            Email: &shared.PropertyCrmEventEmail{
                Cc: []string{
                    "Hills",
                },
                To: []string{
                    "Bronze",
                },
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "Savings",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateCrmEventRequest](../../models/operations/createcrmeventrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateCrmEventResponse](../../models/operations/createcrmeventresponse.md), error**


## CreateCrmFile

Create a file

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmFile(ctx, operations.CreateCrmFileRequest{
        CrmFile: &shared.CrmFile{
            Raw: &shared.PropertyCrmFileRaw{},
        },
        ConnectionID: "ASCII Wooden the",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmFileRequest](../../models/operations/createcrmfilerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmFileResponse](../../models/operations/createcrmfileresponse.md), error**


## CreateCrmLead

Create a lead

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmLead(ctx, operations.CreateCrmLeadRequest{
        CrmLead: &shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "port steradian prize",
                },
            },
        },
        ConnectionID: "ability Einsteinium Orchestrator",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmLeadRequest](../../models/operations/createcrmleadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmLeadResponse](../../models/operations/createcrmleadresponse.md), error**


## CreateCrmPipeline

Create a pipeline

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmPipeline(ctx, operations.CreateCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "Paradigm Vista fuchsia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateCrmPipelineRequest](../../models/operations/createcrmpipelinerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateCrmPipelineResponse](../../models/operations/createcrmpipelineresponse.md), error**


## CreateCrmTeam

Create a team

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmTeam(ctx, operations.CreateCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "exercitationem",
            },
        },
        ConnectionID: "as New Senior",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmTeamRequest](../../models/operations/createcrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmTeamResponse](../../models/operations/createcrmteamresponse.md), error**


## CreateCrmUser

Create a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmUser(ctx, operations.CreateCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "await male",
                },
            },
        },
        ConnectionID: "Incredible Virginia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmUserRequest](../../models/operations/createcrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmUserResponse](../../models/operations/createcrmuserresponse.md), error**


## GetCrmCompany

Retrieve a company

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmCompany(ctx, operations.GetCrmCompanyRequest{
        ConnectionID: "THX Strategist deposit",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCrmCompanyRequest](../../models/operations/getcrmcompanyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetCrmCompanyResponse](../../models/operations/getcrmcompanyresponse.md), error**


## GetCrmContact

Retrieve a contact

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmContact(ctx, operations.GetCrmContactRequest{
        ConnectionID: "Oregon",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCrmContactRequest](../../models/operations/getcrmcontactrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetCrmContactResponse](../../models/operations/getcrmcontactresponse.md), error**


## GetCrmDeal

Retrieve a deal

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmDeal(ctx, operations.GetCrmDealRequest{
        ConnectionID: "male orange",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmDealRequest](../../models/operations/getcrmdealrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmDealResponse](../../models/operations/getcrmdealresponse.md), error**


## GetCrmEvent

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmEvent(ctx, operations.GetCrmEventRequest{
        ConnectionID: "Metal South blockchains",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetCrmEventRequest](../../models/operations/getcrmeventrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetCrmEventResponse](../../models/operations/getcrmeventresponse.md), error**


## GetCrmFile

Retrieve a file

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmFile(ctx, operations.GetCrmFileRequest{
        ConnectionID: "ease",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmFileRequest](../../models/operations/getcrmfilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmFileResponse](../../models/operations/getcrmfileresponse.md), error**


## GetCrmLead

Retrieve a lead

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmLead(ctx, operations.GetCrmLeadRequest{
        ConnectionID: "Handmade Keyboard yum",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmLeadRequest](../../models/operations/getcrmleadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmLeadResponse](../../models/operations/getcrmleadresponse.md), error**


## GetCrmPipeline

Retrieve a pipeline

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmPipeline(ctx, operations.GetCrmPipelineRequest{
        ConnectionID: "withdrawal Southeast",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetCrmPipelineRequest](../../models/operations/getcrmpipelinerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetCrmPipelineResponse](../../models/operations/getcrmpipelineresponse.md), error**


## GetCrmTeam

Retrieve a team

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmTeam(ctx, operations.GetCrmTeamRequest{
        ConnectionID: "digital awful",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmTeamRequest](../../models/operations/getcrmteamrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmTeamResponse](../../models/operations/getcrmteamresponse.md), error**


## GetCrmUser

Retrieve a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmUser(ctx, operations.GetCrmUserRequest{
        ConnectionID: "Bespoke Dollar",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmUserRequest](../../models/operations/getcrmuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmUserResponse](../../models/operations/getcrmuserresponse.md), error**


## ListCrmCompanies

List all companies

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmCompanies(ctx, operations.ListCrmCompaniesRequest{
        ConnectionID: "Jazz solid Lamborghini",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCrmCompaniesRequest](../../models/operations/listcrmcompaniesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListCrmCompaniesResponse](../../models/operations/listcrmcompaniesresponse.md), error**


## ListCrmContacts

List all contacts

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmContacts(ctx, operations.ListCrmContactsRequest{
        ConnectionID: "Awesome index steradian",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCrmContactsRequest](../../models/operations/listcrmcontactsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListCrmContactsResponse](../../models/operations/listcrmcontactsresponse.md), error**


## ListCrmDeals

List all deals

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmDeals(ctx, operations.ListCrmDealsRequest{
        ConnectionID: "Lamborghini",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmDealsRequest](../../models/operations/listcrmdealsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmDealsResponse](../../models/operations/listcrmdealsresponse.md), error**


## ListCrmEvents

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmEvents(ctx, operations.ListCrmEventsRequest{
        ConnectionID: "invoice gratefully",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCrmEventsRequest](../../models/operations/listcrmeventsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListCrmEventsResponse](../../models/operations/listcrmeventsresponse.md), error**


## ListCrmFiles

List all files

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmFiles(ctx, operations.ListCrmFilesRequest{
        ConnectionID: "lavender Genderflux Southeast",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmFilesRequest](../../models/operations/listcrmfilesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmFilesResponse](../../models/operations/listcrmfilesresponse.md), error**


## ListCrmLeads

List all leads

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmLeads(ctx, operations.ListCrmLeadsRequest{
        ConnectionID: "International",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLeads != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmLeadsRequest](../../models/operations/listcrmleadsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmLeadsResponse](../../models/operations/listcrmleadsresponse.md), error**


## ListCrmPipelines

List all pipelines

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmPipelines(ctx, operations.ListCrmPipelinesRequest{
        ConnectionID: "primary",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCrmPipelinesRequest](../../models/operations/listcrmpipelinesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListCrmPipelinesResponse](../../models/operations/listcrmpipelinesresponse.md), error**


## ListCrmTeams

List all teams

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmTeams(ctx, operations.ListCrmTeamsRequest{
        ConnectionID: "Classical microchip Wooden",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeams != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmTeamsRequest](../../models/operations/listcrmteamsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmTeamsResponse](../../models/operations/listcrmteamsresponse.md), error**


## ListCrmUsers

List all users

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmUsers(ctx, operations.ListCrmUsersRequest{
        ConnectionID: "careless Costa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUsers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmUsersRequest](../../models/operations/listcrmusersrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmUsersResponse](../../models/operations/listcrmusersresponse.md), error**


## PatchCrmCompany

Update a company

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmCompany(ctx, operations.PatchCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "Producer",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "Corporate",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "haptic Phased",
                },
            },
            Websites: []string{
                "Investment",
            },
        },
        ConnectionID: "Trans",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchCrmCompanyRequest](../../models/operations/patchcrmcompanyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchCrmCompanyResponse](../../models/operations/patchcrmcompanyresponse.md), error**


## PatchCrmContact

Update a contact

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmContact(ctx, operations.PatchCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "architecture",
            },
            DealIds: []string{
                "Buckinghamshire",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Practical",
                },
            },
        },
        ConnectionID: "Future Diesel",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchCrmContactRequest](../../models/operations/patchcrmcontactrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchCrmContactResponse](../../models/operations/patchcrmcontactresponse.md), error**


## PatchCrmDeal

Update a deal

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmDeal(ctx, operations.PatchCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "consign",
            },
        },
        ConnectionID: "Platinum female",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmDealRequest](../../models/operations/patchcrmdealrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmDealResponse](../../models/operations/patchcrmdealresponse.md), error**


## PatchCrmEvent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmEvent(ctx, operations.PatchCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "XML",
            },
            ContactIds: []string{
                "Accountability",
            },
            DealIds: []string{
                "copying",
            },
            Email: &shared.PropertyCrmEventEmail{
                Cc: []string{
                    "after",
                },
                To: []string{
                    "Research",
                },
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "female",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchCrmEventRequest](../../models/operations/patchcrmeventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PatchCrmEventResponse](../../models/operations/patchcrmeventresponse.md), error**


## PatchCrmFile

Update a file

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmFile(ctx, operations.PatchCrmFileRequest{
        CrmFile: &shared.CrmFile{
            Raw: &shared.PropertyCrmFileRaw{},
        },
        ConnectionID: "bluetooth",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmFileRequest](../../models/operations/patchcrmfilerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmFileResponse](../../models/operations/patchcrmfileresponse.md), error**


## PatchCrmLead

Update a lead

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmLead(ctx, operations.PatchCrmLeadRequest{
        CrmLead: &shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "teal Hat",
                },
            },
        },
        ConnectionID: "Ball Chips",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmLeadRequest](../../models/operations/patchcrmleadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmLeadResponse](../../models/operations/patchcrmleadresponse.md), error**


## PatchCrmPipeline

Update a pipeline

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmPipeline(ctx, operations.PatchCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "imperfect Costa Southwest",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchCrmPipelineRequest](../../models/operations/patchcrmpipelinerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PatchCrmPipelineResponse](../../models/operations/patchcrmpipelineresponse.md), error**


## PatchCrmTeam

Update a team

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmTeam(ctx, operations.PatchCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "Account",
            },
        },
        ConnectionID: "Transexual compress redefine",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmTeamRequest](../../models/operations/patchcrmteamrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmTeamResponse](../../models/operations/patchcrmteamresponse.md), error**


## PatchCrmUser

Update a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmUser(ctx, operations.PatchCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Bronze composite",
                },
            },
        },
        ConnectionID: "katal Industrial Classical",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmUserRequest](../../models/operations/patchcrmuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmUserResponse](../../models/operations/patchcrmuserresponse.md), error**


## RemoveCrmCompany

Remove a company

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmCompany(ctx, operations.RemoveCrmCompanyRequest{
        ConnectionID: "Mayaguez index wireless",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveCrmCompanyRequest](../../models/operations/removecrmcompanyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RemoveCrmCompanyResponse](../../models/operations/removecrmcompanyresponse.md), error**


## RemoveCrmContact

Remove a contact

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmContact(ctx, operations.RemoveCrmContactRequest{
        ConnectionID: "Folk granular Concrete",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveCrmContactRequest](../../models/operations/removecrmcontactrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RemoveCrmContactResponse](../../models/operations/removecrmcontactresponse.md), error**


## RemoveCrmDeal

Remove a deal

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmDeal(ctx, operations.RemoveCrmDealRequest{
        ConnectionID: "Nihonium",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmDealRequest](../../models/operations/removecrmdealrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmDealResponse](../../models/operations/removecrmdealresponse.md), error**


## RemoveCrmEvent

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmEvent(ctx, operations.RemoveCrmEventRequest{
        ConnectionID: "card",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RemoveCrmEventRequest](../../models/operations/removecrmeventrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.RemoveCrmEventResponse](../../models/operations/removecrmeventresponse.md), error**


## RemoveCrmFile

Remove a file

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmFile(ctx, operations.RemoveCrmFileRequest{
        ConnectionID: "cash",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmFileRequest](../../models/operations/removecrmfilerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmFileResponse](../../models/operations/removecrmfileresponse.md), error**


## RemoveCrmLead

Remove a lead

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmLead(ctx, operations.RemoveCrmLeadRequest{
        ConnectionID: "Southeast",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmLeadRequest](../../models/operations/removecrmleadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmLeadResponse](../../models/operations/removecrmleadresponse.md), error**


## RemoveCrmPipeline

Remove a pipeline

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmPipeline(ctx, operations.RemoveCrmPipelineRequest{
        ConnectionID: "Hybrid merrily",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveCrmPipelineRequest](../../models/operations/removecrmpipelinerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RemoveCrmPipelineResponse](../../models/operations/removecrmpipelineresponse.md), error**


## RemoveCrmTeam

Remove a team

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmTeam(ctx, operations.RemoveCrmTeamRequest{
        ConnectionID: "Sol",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmTeamRequest](../../models/operations/removecrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmTeamResponse](../../models/operations/removecrmteamresponse.md), error**


## RemoveCrmUser

Remove a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmUser(ctx, operations.RemoveCrmUserRequest{
        ConnectionID: "Southeast",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmUserRequest](../../models/operations/removecrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmUserResponse](../../models/operations/removecrmuserresponse.md), error**


## UpdateCrmCompany

Update a company

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmCompany(ctx, operations.UpdateCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "SMS",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "barrel",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Account alarm infrastructure",
                },
            },
            Websites: []string{
                "Visionary",
            },
        },
        ConnectionID: "Southeast ad",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCrmCompanyRequest](../../models/operations/updatecrmcompanyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateCrmCompanyResponse](../../models/operations/updatecrmcompanyresponse.md), error**


## UpdateCrmContact

Update a contact

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmContact(ctx, operations.UpdateCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "Universal",
            },
            DealIds: []string{
                "Harbors",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Analyst Des green",
                },
            },
        },
        ConnectionID: "man panel",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCrmContactRequest](../../models/operations/updatecrmcontactrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateCrmContactResponse](../../models/operations/updatecrmcontactresponse.md), error**


## UpdateCrmDeal

Update a deal

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmDeal(ctx, operations.UpdateCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "South",
            },
        },
        ConnectionID: "Shirt",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmDealRequest](../../models/operations/updatecrmdealrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmDealResponse](../../models/operations/updatecrmdealresponse.md), error**


## UpdateCrmEvent

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmEvent(ctx, operations.UpdateCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "Account",
            },
            ContactIds: []string{
                "DRAM",
            },
            DealIds: []string{
                "input",
            },
            Email: &shared.PropertyCrmEventEmail{
                Cc: []string{
                    "Bicycle",
                },
                To: []string{
                    "Wagon",
                },
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "Accountability",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateCrmEventRequest](../../models/operations/updatecrmeventrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateCrmEventResponse](../../models/operations/updatecrmeventresponse.md), error**


## UpdateCrmFile

Update a file

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmFile(ctx, operations.UpdateCrmFileRequest{
        CrmFile: &shared.CrmFile{
            Raw: &shared.PropertyCrmFileRaw{},
        },
        ConnectionID: "Orchestrator",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmFileRequest](../../models/operations/updatecrmfilerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmFileResponse](../../models/operations/updatecrmfileresponse.md), error**


## UpdateCrmLead

Update a lead

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmLead(ctx, operations.UpdateCrmLeadRequest{
        CrmLead: &shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "compelling",
                },
            },
        },
        ConnectionID: "Pickup Polestar Checking",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmLeadRequest](../../models/operations/updatecrmleadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmLeadResponse](../../models/operations/updatecrmleadresponse.md), error**


## UpdateCrmPipeline

Update a pipeline

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmPipeline(ctx, operations.UpdateCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "needily",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateCrmPipelineRequest](../../models/operations/updatecrmpipelinerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateCrmPipelineResponse](../../models/operations/updatecrmpipelineresponse.md), error**


## UpdateCrmTeam

Update a team

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmTeam(ctx, operations.UpdateCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "Carbon",
            },
        },
        ConnectionID: "Dakota",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmTeamRequest](../../models/operations/updatecrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmTeamResponse](../../models/operations/updatecrmteamresponse.md), error**


## UpdateCrmUser

Update a user

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmUser(ctx, operations.UpdateCrmUserRequest{
        CrmUser: &shared.CrmUser{
            Address: &shared.PropertyCrmUserAddress{},
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Idaho green",
                },
            },
        },
        ConnectionID: "Savings",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmUserRequest](../../models/operations/updatecrmuserrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmUserResponse](../../models/operations/updatecrmuserresponse.md), error**

