# Crm
(*Crm*)

### Available Operations

* [DeleteCrmConnectionIDCompanyID](#deletecrmconnectionidcompanyid) - Remove a company
* [DeleteCrmConnectionIDContactID](#deletecrmconnectionidcontactid) - Remove a contact
* [DeleteCrmConnectionIDDealID](#deletecrmconnectioniddealid) - Remove a deal
* [DeleteCrmConnectionIDEventID](#deletecrmconnectionideventid) - Remove a event
* [DeleteCrmConnectionIDFileID](#deletecrmconnectionidfileid) - Remove a file
* [DeleteCrmConnectionIDLeadID](#deletecrmconnectionidleadid) - Remove a lead
* [DeleteCrmConnectionIDPipelineID](#deletecrmconnectionidpipelineid) - Remove a pipeline
* [DeleteCrmConnectionIDTeamID](#deletecrmconnectionidteamid) - Remove a team
* [DeleteCrmConnectionIDUserID](#deletecrmconnectioniduserid) - Remove a user
* [GetCrmConnectionIDCompany](#getcrmconnectionidcompany) - List all companies
* [GetCrmConnectionIDCompanyID](#getcrmconnectionidcompanyid) - Retrieve a company
* [GetCrmConnectionIDContact](#getcrmconnectionidcontact) - List all contacts
* [GetCrmConnectionIDContactID](#getcrmconnectionidcontactid) - Retrieve a contact
* [GetCrmConnectionIDDeal](#getcrmconnectioniddeal) - List all deals
* [GetCrmConnectionIDDealID](#getcrmconnectioniddealid) - Retrieve a deal
* [GetCrmConnectionIDEvent](#getcrmconnectionidevent) - List all events
* [GetCrmConnectionIDEventID](#getcrmconnectionideventid) - Retrieve a event
* [GetCrmConnectionIDFile](#getcrmconnectionidfile) - List all files
* [GetCrmConnectionIDFileID](#getcrmconnectionidfileid) - Retrieve a file
* [GetCrmConnectionIDLead](#getcrmconnectionidlead) - List all leads
* [GetCrmConnectionIDLeadID](#getcrmconnectionidleadid) - Retrieve a lead
* [GetCrmConnectionIDPipeline](#getcrmconnectionidpipeline) - List all pipelines
* [GetCrmConnectionIDPipelineID](#getcrmconnectionidpipelineid) - Retrieve a pipeline
* [GetCrmConnectionIDTeam](#getcrmconnectionidteam) - List all teams
* [GetCrmConnectionIDTeamID](#getcrmconnectionidteamid) - Retrieve a team
* [GetCrmConnectionIDUser](#getcrmconnectioniduser) - List all users
* [GetCrmConnectionIDUserID](#getcrmconnectioniduserid) - Retrieve a user
* [PatchCrmConnectionIDCompanyID](#patchcrmconnectionidcompanyid) - Update a company
* [PatchCrmConnectionIDContactID](#patchcrmconnectionidcontactid) - Update a contact
* [PatchCrmConnectionIDDealID](#patchcrmconnectioniddealid) - Update a deal
* [PatchCrmConnectionIDEventID](#patchcrmconnectionideventid) - Update a event
* [PatchCrmConnectionIDFileID](#patchcrmconnectionidfileid) - Update a file
* [PatchCrmConnectionIDLeadID](#patchcrmconnectionidleadid) - Update a lead
* [PatchCrmConnectionIDPipelineID](#patchcrmconnectionidpipelineid) - Update a pipeline
* [PatchCrmConnectionIDTeamID](#patchcrmconnectionidteamid) - Update a team
* [PatchCrmConnectionIDUserID](#patchcrmconnectioniduserid) - Update a user
* [PostCrmConnectionIDCompany](#postcrmconnectionidcompany) - Create a company
* [PostCrmConnectionIDContact](#postcrmconnectionidcontact) - Create a contact
* [PostCrmConnectionIDDeal](#postcrmconnectioniddeal) - Create a deal
* [PostCrmConnectionIDEvent](#postcrmconnectionidevent) - Create a event
* [PostCrmConnectionIDFile](#postcrmconnectionidfile) - Create a file
* [PostCrmConnectionIDLead](#postcrmconnectionidlead) - Create a lead
* [PostCrmConnectionIDPipeline](#postcrmconnectionidpipeline) - Create a pipeline
* [PostCrmConnectionIDTeam](#postcrmconnectionidteam) - Create a team
* [PostCrmConnectionIDUser](#postcrmconnectioniduser) - Create a user
* [PutCrmConnectionIDCompanyID](#putcrmconnectionidcompanyid) - Update a company
* [PutCrmConnectionIDContactID](#putcrmconnectionidcontactid) - Update a contact
* [PutCrmConnectionIDDealID](#putcrmconnectioniddealid) - Update a deal
* [PutCrmConnectionIDEventID](#putcrmconnectionideventid) - Update a event
* [PutCrmConnectionIDFileID](#putcrmconnectionidfileid) - Update a file
* [PutCrmConnectionIDLeadID](#putcrmconnectionidleadid) - Update a lead
* [PutCrmConnectionIDPipelineID](#putcrmconnectionidpipelineid) - Update a pipeline
* [PutCrmConnectionIDTeamID](#putcrmconnectionidteamid) - Update a team
* [PutCrmConnectionIDUserID](#putcrmconnectioniduserid) - Update a user

## DeleteCrmConnectionIDCompanyID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDCompanyID(ctx, operations.DeleteCrmConnectionIDCompanyIDRequest{
        ConnectionID: "hertz morph",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteCrmConnectionIDCompanyIDRequest](../../models/operations/deletecrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteCrmConnectionIDCompanyIDResponse](../../models/operations/deletecrmconnectionidcompanyidresponse.md), error**


## DeleteCrmConnectionIDContactID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDContactID(ctx, operations.DeleteCrmConnectionIDContactIDRequest{
        ConnectionID: "chargesheet",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteCrmConnectionIDContactIDRequest](../../models/operations/deletecrmconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteCrmConnectionIDContactIDResponse](../../models/operations/deletecrmconnectionidcontactidresponse.md), error**


## DeleteCrmConnectionIDDealID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDDealID(ctx, operations.DeleteCrmConnectionIDDealIDRequest{
        ConnectionID: "Fresh",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDDealIDRequest](../../models/operations/deletecrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDDealIDResponse](../../models/operations/deletecrmconnectioniddealidresponse.md), error**


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
    res, err := s.Crm.DeleteCrmConnectionIDEventID(ctx, operations.DeleteCrmConnectionIDEventIDRequest{
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


## DeleteCrmConnectionIDFileID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDFileID(ctx, operations.DeleteCrmConnectionIDFileIDRequest{
        ConnectionID: "Bicycle",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDFileIDRequest](../../models/operations/deletecrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDFileIDResponse](../../models/operations/deletecrmconnectionidfileidresponse.md), error**


## DeleteCrmConnectionIDLeadID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDLeadID(ctx, operations.DeleteCrmConnectionIDLeadIDRequest{
        ConnectionID: "Senior azure",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDLeadIDRequest](../../models/operations/deletecrmconnectionidleadidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDLeadIDResponse](../../models/operations/deletecrmconnectionidleadidresponse.md), error**


## DeleteCrmConnectionIDPipelineID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDPipelineID(ctx, operations.DeleteCrmConnectionIDPipelineIDRequest{
        ConnectionID: "Customer",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteCrmConnectionIDPipelineIDRequest](../../models/operations/deletecrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.DeleteCrmConnectionIDPipelineIDResponse](../../models/operations/deletecrmconnectionidpipelineidresponse.md), error**


## DeleteCrmConnectionIDTeamID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDTeamID(ctx, operations.DeleteCrmConnectionIDTeamIDRequest{
        ConnectionID: "Diverse",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDTeamIDRequest](../../models/operations/deletecrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDTeamIDResponse](../../models/operations/deletecrmconnectionidteamidresponse.md), error**


## DeleteCrmConnectionIDUserID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.DeleteCrmConnectionIDUserID(ctx, operations.DeleteCrmConnectionIDUserIDRequest{
        ConnectionID: "Intranet Data",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDUserIDRequest](../../models/operations/deletecrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDUserIDResponse](../../models/operations/deletecrmconnectioniduseridresponse.md), error**


## GetCrmConnectionIDCompany

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDCompany(ctx, operations.GetCrmConnectionIDCompanyRequest{
        ConnectionID: "indexing",
        ContactID: unifiedgosdk.String("Porsche firewall"),
        DealID: unifiedgosdk.String("Hafnium Computers"),
        Limit: unifiedgosdk.Float64(902.85),
        Offset: unifiedgosdk.Float64(2893.88),
        Order: unifiedgosdk.String("Interactions relationships juxtapose"),
        Query: unifiedgosdk.String("newton Luxembourg"),
        Sort: unifiedgosdk.String("Dakota quantifying Actinium"),
        UpdatedGte: types.MustTimeFromString("2022-09-27T07:42:48.074Z"),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDCompanyRequest](../../models/operations/getcrmconnectionidcompanyrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetCrmConnectionIDCompanyResponse](../../models/operations/getcrmconnectionidcompanyresponse.md), error**


## GetCrmConnectionIDCompanyID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDCompanyID(ctx, operations.GetCrmConnectionIDCompanyIDRequest{
        ConnectionID: "Netherlands",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCrmConnectionIDCompanyIDRequest](../../models/operations/getcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCrmConnectionIDCompanyIDResponse](../../models/operations/getcrmconnectionidcompanyidresponse.md), error**


## GetCrmConnectionIDContact

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDContact(ctx, operations.GetCrmConnectionIDContactRequest{
        CompanyID: unifiedgosdk.String("Southeast Human Southeast"),
        ConnectionID: "magenta loose",
        DealID: unifiedgosdk.String("intuitive"),
        Limit: unifiedgosdk.Float64(9605),
        Offset: unifiedgosdk.Float64(8572.44),
        Order: unifiedgosdk.String("Music Electronics"),
        Query: unifiedgosdk.String("Elegant"),
        Sort: unifiedgosdk.String("North Analyst Otis"),
        UpdatedGte: types.MustTimeFromString("2022-09-18T15:42:24.943Z"),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDContactRequest](../../models/operations/getcrmconnectionidcontactrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetCrmConnectionIDContactResponse](../../models/operations/getcrmconnectionidcontactresponse.md), error**


## GetCrmConnectionIDContactID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDContactID(ctx, operations.GetCrmConnectionIDContactIDRequest{
        ConnectionID: "Account fountain visionary",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCrmConnectionIDContactIDRequest](../../models/operations/getcrmconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCrmConnectionIDContactIDResponse](../../models/operations/getcrmconnectionidcontactidresponse.md), error**


## GetCrmConnectionIDDeal

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDDeal(ctx, operations.GetCrmConnectionIDDealRequest{
        CompanyID: unifiedgosdk.String("Tools Card copying"),
        ConnectionID: "Renminbi",
        ContactID: unifiedgosdk.String("till payment World"),
        Limit: unifiedgosdk.Float64(8656.16),
        Offset: unifiedgosdk.Float64(4455.8),
        Order: unifiedgosdk.String("global"),
        Query: unifiedgosdk.String("Program Bespoke Wisconsin"),
        Sort: unifiedgosdk.String("Netherlands under"),
        UpdatedGte: types.MustTimeFromString("2022-12-23T01:47:21.816Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDDealRequest](../../models/operations/getcrmconnectioniddealrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDDealResponse](../../models/operations/getcrmconnectioniddealresponse.md), error**


## GetCrmConnectionIDDealID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDDealID(ctx, operations.GetCrmConnectionIDDealIDRequest{
        ConnectionID: "Concrete Director",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDDealIDRequest](../../models/operations/getcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDDealIDResponse](../../models/operations/getcrmconnectioniddealidresponse.md), error**


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
    res, err := s.Crm.GetCrmConnectionIDEvent(ctx, operations.GetCrmConnectionIDEventRequest{
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
    res, err := s.Crm.GetCrmConnectionIDEventID(ctx, operations.GetCrmConnectionIDEventIDRequest{
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


## GetCrmConnectionIDFile

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDFile(ctx, operations.GetCrmConnectionIDFileRequest{
        CompanyID: unifiedgosdk.String("reboot"),
        ConnectionID: "customise far",
        ContactID: unifiedgosdk.String("Electronic proactive"),
        DealID: unifiedgosdk.String("withdrawal deposit Gloves"),
        Limit: unifiedgosdk.Float64(1588.79),
        Offset: unifiedgosdk.Float64(3754.81),
        Order: unifiedgosdk.String("Implemented fairly meh"),
        Query: unifiedgosdk.String("FTP Producer"),
        Sort: unifiedgosdk.String("soprano deliverables"),
        UpdatedGte: types.MustTimeFromString("2022-03-02T03:00:09.711Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDFileRequest](../../models/operations/getcrmconnectionidfilerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDFileResponse](../../models/operations/getcrmconnectionidfileresponse.md), error**


## GetCrmConnectionIDFileID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDFileID(ctx, operations.GetCrmConnectionIDFileIDRequest{
        ConnectionID: "Bicycle",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDFileIDRequest](../../models/operations/getcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDFileIDResponse](../../models/operations/getcrmconnectionidfileidresponse.md), error**


## GetCrmConnectionIDLead

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDLead(ctx, operations.GetCrmConnectionIDLeadRequest{
        ConnectionID: "Computer Hop",
        Limit: unifiedgosdk.Float64(7411.81),
        Offset: unifiedgosdk.Float64(9004.32),
        Order: unifiedgosdk.String("Operations candela Integration"),
        Query: unifiedgosdk.String("impactful transform"),
        Sort: unifiedgosdk.String("Tala defense Southwest"),
        UpdatedGte: types.MustTimeFromString("2021-09-29T00:37:32.184Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDLeadRequest](../../models/operations/getcrmconnectionidleadrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDLeadResponse](../../models/operations/getcrmconnectionidleadresponse.md), error**


## GetCrmConnectionIDLeadID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDLeadID(ctx, operations.GetCrmConnectionIDLeadIDRequest{
        ConnectionID: "users Minnesota Bypass",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDLeadIDRequest](../../models/operations/getcrmconnectionidleadidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDLeadIDResponse](../../models/operations/getcrmconnectionidleadidresponse.md), error**


## GetCrmConnectionIDPipeline

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDPipeline(ctx, operations.GetCrmConnectionIDPipelineRequest{
        ConnectionID: "dirty Awesome Checking",
        Limit: unifiedgosdk.Float64(9055.88),
        Offset: unifiedgosdk.Float64(3443.76),
        Order: unifiedgosdk.String("glom"),
        Query: unifiedgosdk.String("panel"),
        Sort: unifiedgosdk.String("Latin tightly"),
        UpdatedGte: types.MustTimeFromString("2022-03-01T15:47:43.244Z"),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCrmConnectionIDPipelineRequest](../../models/operations/getcrmconnectionidpipelinerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetCrmConnectionIDPipelineResponse](../../models/operations/getcrmconnectionidpipelineresponse.md), error**


## GetCrmConnectionIDPipelineID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDPipelineID(ctx, operations.GetCrmConnectionIDPipelineIDRequest{
        ConnectionID: "Tricycle roughly markets",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCrmConnectionIDPipelineIDRequest](../../models/operations/getcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetCrmConnectionIDPipelineIDResponse](../../models/operations/getcrmconnectionidpipelineidresponse.md), error**


## GetCrmConnectionIDTeam

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDTeam(ctx, operations.GetCrmConnectionIDTeamRequest{
        ConnectionID: "bath Lamborghini",
        Limit: unifiedgosdk.Float64(1042.31),
        Offset: unifiedgosdk.Float64(1586.42),
        Order: unifiedgosdk.String("Diesel Bike virtual"),
        Query: unifiedgosdk.String("bakery"),
        Sort: unifiedgosdk.String("Senior"),
        UpdatedGte: types.MustTimeFromString("2021-12-04T23:56:00.028Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDTeamRequest](../../models/operations/getcrmconnectionidteamrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDTeamResponse](../../models/operations/getcrmconnectionidteamresponse.md), error**


## GetCrmConnectionIDTeamID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDTeamID(ctx, operations.GetCrmConnectionIDTeamIDRequest{
        ConnectionID: "Intelligent invoice Tesla",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDTeamIDRequest](../../models/operations/getcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDTeamIDResponse](../../models/operations/getcrmconnectionidteamidresponse.md), error**


## GetCrmConnectionIDUser

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDUser(ctx, operations.GetCrmConnectionIDUserRequest{
        ConnectionID: "suit Electronic Tampa",
        Limit: unifiedgosdk.Float64(2883.34),
        Offset: unifiedgosdk.Float64(8886.55),
        Order: unifiedgosdk.String("despite"),
        Query: unifiedgosdk.String("frightfully Fitness"),
        Sort: unifiedgosdk.String("success servant"),
        UpdatedGte: types.MustTimeFromString("2023-02-23T05:53:04.259Z"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDUserRequest](../../models/operations/getcrmconnectioniduserrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDUserResponse](../../models/operations/getcrmconnectioniduserresponse.md), error**


## GetCrmConnectionIDUserID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmConnectionIDUserID(ctx, operations.GetCrmConnectionIDUserIDRequest{
        ConnectionID: "connecting Program",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDUserIDRequest](../../models/operations/getcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDUserIDResponse](../../models/operations/getcrmconnectioniduseridresponse.md), error**


## PatchCrmConnectionIDCompanyID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDCompanyID(ctx, operations.PatchCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("invoice"),
                Address2: unifiedgosdk.String("indexing Ford"),
                City: unifiedgosdk.String("McAllen"),
                Country: unifiedgosdk.String("Netherlands"),
                CountryCode: unifiedgosdk.String("PF"),
                PostalCode: unifiedgosdk.String("93486"),
                Region: unifiedgosdk.String("Steel impactful"),
                RegionCode: unifiedgosdk.String("Dong"),
            },
            CreatedAt: types.MustTimeFromString("2023-07-25T08:43:38.995Z"),
            DealIds: []string{
                "usefully",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Annabel31@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Toys Vermont Astatine"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "Trigender",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "female",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-12-22T11:39:56.432Z"),
            Websites: []string{
                "Latin",
            },
        },
        ConnectionID: "North kilogram connecting",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchCrmConnectionIDCompanyIDRequest](../../models/operations/patchcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PatchCrmConnectionIDCompanyIDResponse](../../models/operations/patchcrmconnectionidcompanyidresponse.md), error**


## PatchCrmConnectionIDContactID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDContactID(ctx, operations.PatchCrmConnectionIDContactIDRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.String("until instantly Taiwan"),
                Address2: unifiedgosdk.String("disintermediate ah Southwest"),
                City: unifiedgosdk.String("San Antonio"),
                Country: unifiedgosdk.String("Djibouti"),
                CountryCode: unifiedgosdk.String("LA"),
                PostalCode: unifiedgosdk.String("23695"),
                Region: unifiedgosdk.String("grey around"),
                RegionCode: unifiedgosdk.String("Folding"),
            },
            Company: unifiedgosdk.String("Johnson - Gerlach"),
            CompanyIds: []string{
                "Personal",
            },
            CreatedAt: types.MustTimeFromString("2022-07-24T05:16:20.203Z"),
            DealIds: []string{
                "generation",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Leora_Konopelski27@hotmail.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Innovative indeed brand"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "unsung Borders",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("withdrawal"),
            UpdatedAt: types.MustTimeFromString("2022-05-05T23:37:21.563Z"),
        },
        ConnectionID: "markets radian",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchCrmConnectionIDContactIDRequest](../../models/operations/patchcrmconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PatchCrmConnectionIDContactIDResponse](../../models/operations/patchcrmconnectionidcontactidresponse.md), error**


## PatchCrmConnectionIDDealID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDDealID(ctx, operations.PatchCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(7725.78),
            ClosedAt: types.MustTimeFromString("2021-10-28T08:42:49.591Z"),
            CreatedAt: types.MustTimeFromString("2023-04-23T15:03:53.999Z"),
            Currency: unifiedgosdk.String("Afghani"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("North"),
            Name: unifiedgosdk.String("midnight"),
            Pipeline: unifiedgosdk.String("envisioneer Functionality Loan"),
            Probability: unifiedgosdk.Float64(7051.73),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("Krone"),
            Stage: unifiedgosdk.String("pascal aliquam gripping"),
            Tags: []string{
                "where",
            },
            UpdatedAt: types.MustTimeFromString("2022-04-05T10:21:22.505Z"),
            WonReason: unifiedgosdk.String("Savings kilogram"),
        },
        ConnectionID: "Chair weber silver",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDDealIDRequest](../../models/operations/patchcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDDealIDResponse](../../models/operations/patchcrmconnectioniddealidresponse.md), error**


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
    res, err := s.Crm.PatchCrmConnectionIDEventID(ctx, operations.PatchCrmConnectionIDEventIDRequest{
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


## PatchCrmConnectionIDFileID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDFileID(ctx, operations.PatchCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("duh Handmade harness"),
            CompanyID: unifiedgosdk.String("CFP"),
            ContactID: unifiedgosdk.String("unaware yellow generating"),
            CreatedAt: types.MustTimeFromString("2021-05-04T04:54:33.785Z"),
            DealID: unifiedgosdk.String("channels SUV"),
            Description: unifiedgosdk.String("De-engineered didactic hardware"),
            FileName: unifiedgosdk.String("metical_silver_yellow.html"),
            FileSize: unifiedgosdk.Float64(6861.53),
            FileType: unifiedgosdk.String("video"),
            FileURL: unifiedgosdk.String("navigate Funk"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("internal"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-02-21T13:46:42.012Z"),
            UserID: unifiedgosdk.String("Interactions"),
        },
        ConnectionID: "Handcrafted",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDFileIDRequest](../../models/operations/patchcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDFileIDResponse](../../models/operations/patchcrmconnectionidfileidresponse.md), error**


## PatchCrmConnectionIDLeadID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDLeadID(ctx, operations.PatchCrmConnectionIDLeadIDRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.String("Cambridgeshire"),
                Address2: unifiedgosdk.String("Oriental farad male"),
                City: unifiedgosdk.String("D'Amorebury"),
                Country: unifiedgosdk.String("Reunion"),
                CountryCode: unifiedgosdk.String("UY"),
                PostalCode: unifiedgosdk.String("87017-9001"),
                Region: unifiedgosdk.String("Buckinghamshire Electric"),
                RegionCode: unifiedgosdk.String("South gee"),
            },
            CompanyID: unifiedgosdk.String("Gasoline conglomeration Tennessine"),
            ContactID: unifiedgosdk.String("grow hub"),
            CreatedAt: types.MustTimeFromString("2023-06-09T15:23:12.644Z"),
            CreatorUserID: unifiedgosdk.String("voluptates"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jeffrey.Denesik52@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("whiteboard lumen"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Cheese before against",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-07-12T14:19:50.007Z"),
            UserID: unifiedgosdk.String("Games yellow Towels"),
        },
        ConnectionID: "brr misuse",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDLeadIDRequest](../../models/operations/patchcrmconnectionidleadidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDLeadIDResponse](../../models/operations/patchcrmconnectionidleadidresponse.md), error**


## PatchCrmConnectionIDPipelineID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDPipelineID(ctx, operations.PatchCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2023-08-24T17:39:51.183Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(664.58),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("bandwidth"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2023-11-27T01:55:15.440Z"),
        },
        ConnectionID: "Chips",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchCrmConnectionIDPipelineIDRequest](../../models/operations/patchcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PatchCrmConnectionIDPipelineIDResponse](../../models/operations/patchcrmconnectionidpipelineidresponse.md), error**


## PatchCrmConnectionIDTeamID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDTeamID(ctx, operations.PatchCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2021-05-20T12:47:48.451Z"),
            Description: unifiedgosdk.String("Automated executive emulation"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Internal experiences"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-05-22T09:41:53.599Z"),
            UserIds: []string{
                "lumen",
            },
        },
        ConnectionID: "up Candace",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDTeamIDRequest](../../models/operations/patchcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDTeamIDResponse](../../models/operations/patchcrmconnectionidteamidresponse.md), error**


## PatchCrmConnectionIDUserID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmConnectionIDUserID(ctx, operations.PatchCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("Customer"),
                Address2: unifiedgosdk.String("violet groupware blanditiis"),
                City: unifiedgosdk.String("South Phoebeshire"),
                Country: unifiedgosdk.String("Thailand"),
                CountryCode: unifiedgosdk.String("NO"),
                PostalCode: unifiedgosdk.String("30801-4594"),
                Region: unifiedgosdk.String("portals Vanadium"),
                RegionCode: unifiedgosdk.String("Future"),
            },
            CreatedAt: types.MustTimeFromString("2023-01-04T02:42:28.788Z"),
            Currency: unifiedgosdk.String("Guinea Franc"),
            Department: unifiedgosdk.String("Gloves global rosin"),
            Division: unifiedgosdk.String("Berkshire Europium"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Wade.Dach@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Checking"),
            LanguageLocale: unifiedgosdk.String("Sedan Porsche matrix"),
            Name: unifiedgosdk.String("superstructure Nissan sedately"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "unto ubiquitous input",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("America/Tijuana"),
            Title: unifiedgosdk.String("Computer Bicycle"),
            UpdatedAt: types.MustTimeFromString("2021-12-13T16:36:33.886Z"),
        },
        ConnectionID: "gold generating",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDUserIDRequest](../../models/operations/patchcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDUserIDResponse](../../models/operations/patchcrmconnectioniduseridresponse.md), error**


## PostCrmConnectionIDCompany

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDCompany(ctx, operations.PostCrmConnectionIDCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("consequently gosh"),
                Address2: unifiedgosdk.String("phooey"),
                City: unifiedgosdk.String("Antonettaville"),
                Country: unifiedgosdk.String("Lebanon"),
                CountryCode: unifiedgosdk.String("SI"),
                PostalCode: unifiedgosdk.String("79462"),
                Region: unifiedgosdk.String("orchid Oxygen Kids"),
                RegionCode: unifiedgosdk.String("Electric utilisation"),
            },
            CreatedAt: types.MustTimeFromString("2021-10-01T08:46:18.197Z"),
            DealIds: []string{
                "Tennessee",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jaida.McDermott26@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Hydrogen Wooden"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "CSS",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Account invoice",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-09-12T17:31:24.634Z"),
            Websites: []string{
                "Intuitive",
            },
        },
        ConnectionID: "Gasoline",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostCrmConnectionIDCompanyRequest](../../models/operations/postcrmconnectionidcompanyrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostCrmConnectionIDCompanyResponse](../../models/operations/postcrmconnectionidcompanyresponse.md), error**


## PostCrmConnectionIDContact

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDContact(ctx, operations.PostCrmConnectionIDContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.String("orchid"),
                Address2: unifiedgosdk.String("invoice wherever watt"),
                City: unifiedgosdk.String("Rempelcester"),
                Country: unifiedgosdk.String("Nepal"),
                CountryCode: unifiedgosdk.String("FI"),
                PostalCode: unifiedgosdk.String("27896-6482"),
                Region: unifiedgosdk.String("swig"),
                RegionCode: unifiedgosdk.String("Recumbent"),
            },
            Company: unifiedgosdk.String("Fritsch - Bernhard"),
            CompanyIds: []string{
                "Executive",
            },
            CreatedAt: types.MustTimeFromString("2021-07-26T17:34:53.280Z"),
            DealIds: []string{
                "Southwest",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Colby24@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("farad"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Dynamic withdrawal",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("second Fresh"),
            UpdatedAt: types.MustTimeFromString("2023-01-03T09:41:22.581Z"),
        },
        ConnectionID: "what",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostCrmConnectionIDContactRequest](../../models/operations/postcrmconnectionidcontactrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostCrmConnectionIDContactResponse](../../models/operations/postcrmconnectionidcontactresponse.md), error**


## PostCrmConnectionIDDeal

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDDeal(ctx, operations.PostCrmConnectionIDDealRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(6144.41),
            ClosedAt: types.MustTimeFromString("2022-07-10T09:55:59.977Z"),
            CreatedAt: types.MustTimeFromString("2022-01-20T07:28:03.436Z"),
            Currency: unifiedgosdk.String("Convertible Marks"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("pfft female"),
            Name: unifiedgosdk.String("Expressway"),
            Pipeline: unifiedgosdk.String("withdrawal Extended busily"),
            Probability: unifiedgosdk.Float64(7998.22),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("spiffy sometimes"),
            Stage: unifiedgosdk.String("transmitter"),
            Tags: []string{
                "intermediate",
            },
            UpdatedAt: types.MustTimeFromString("2022-10-06T18:34:11.762Z"),
            WonReason: unifiedgosdk.String("Cisgender input HTTP"),
        },
        ConnectionID: "accusantium Checking",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDDealRequest](../../models/operations/postcrmconnectioniddealrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDDealResponse](../../models/operations/postcrmconnectioniddealresponse.md), error**


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
    res, err := s.Crm.PostCrmConnectionIDEvent(ctx, operations.PostCrmConnectionIDEventRequest{
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


## PostCrmConnectionIDFile

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDFile(ctx, operations.PostCrmConnectionIDFileRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("tan impedit Pickup"),
            CompanyID: unifiedgosdk.String("Manager"),
            ContactID: unifiedgosdk.String("Florida Shoes East"),
            CreatedAt: types.MustTimeFromString("2023-01-08T11:37:24.708Z"),
            DealID: unifiedgosdk.String("Agent"),
            Description: unifiedgosdk.String("Multi-lateral well-modulated portal"),
            FileName: unifiedgosdk.String("panel_city.wav"),
            FileSize: unifiedgosdk.Float64(1401.73),
            FileType: unifiedgosdk.String("application"),
            FileURL: unifiedgosdk.String("for Chips under"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("abaft Checking"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-02-25T09:46:59.608Z"),
            UserID: unifiedgosdk.String("Mexico withdrawal"),
        },
        ConnectionID: "national Lead",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDFileRequest](../../models/operations/postcrmconnectionidfilerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDFileResponse](../../models/operations/postcrmconnectionidfileresponse.md), error**


## PostCrmConnectionIDLead

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDLead(ctx, operations.PostCrmConnectionIDLeadRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.String("XSS Country knowledge"),
                Address2: unifiedgosdk.String("structure"),
                City: unifiedgosdk.String("Giovaniton"),
                Country: unifiedgosdk.String("Ghana"),
                CountryCode: unifiedgosdk.String("CO"),
                PostalCode: unifiedgosdk.String("34495-0585"),
                Region: unifiedgosdk.String("Modern"),
                RegionCode: unifiedgosdk.String("Diesel"),
            },
            CompanyID: unifiedgosdk.String("yuppify"),
            ContactID: unifiedgosdk.String("demanding scratch male"),
            CreatedAt: types.MustTimeFromString("2023-03-07T11:22:05.657Z"),
            CreatorUserID: unifiedgosdk.String("masticate South"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Gregorio37@gmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Granite Tools"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Hassium Balanced male",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-06-26T11:56:58.926Z"),
            UserID: unifiedgosdk.String("Consultant"),
        },
        ConnectionID: "solutions gosh",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDLeadRequest](../../models/operations/postcrmconnectionidleadrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDLeadResponse](../../models/operations/postcrmconnectionidleadresponse.md), error**


## PostCrmConnectionIDPipeline

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDPipeline(ctx, operations.PostCrmConnectionIDPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2023-12-10T23:55:22.206Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(3879.73),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("upward Mayaguez"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2021-09-25T10:43:23.679Z"),
        },
        ConnectionID: "Lead Health",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostCrmConnectionIDPipelineRequest](../../models/operations/postcrmconnectionidpipelinerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostCrmConnectionIDPipelineResponse](../../models/operations/postcrmconnectionidpipelineresponse.md), error**


## PostCrmConnectionIDTeam

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDTeam(ctx, operations.PostCrmConnectionIDTeamRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2022-02-12T08:57:03.070Z"),
            Description: unifiedgosdk.String("Organic transitional portal"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("male bandwidth"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-12-29T15:50:04.365Z"),
            UserIds: []string{
                "meter",
            },
        },
        ConnectionID: "Guaynabo AGP East",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDTeamRequest](../../models/operations/postcrmconnectionidteamrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDTeamResponse](../../models/operations/postcrmconnectionidteamresponse.md), error**


## PostCrmConnectionIDUser

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PostCrmConnectionIDUser(ctx, operations.PostCrmConnectionIDUserRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("driver East"),
                Address2: unifiedgosdk.String("relationships Computer navigate"),
                City: unifiedgosdk.String("Homestead"),
                Country: unifiedgosdk.String("Cayman Islands"),
                CountryCode: unifiedgosdk.String("BW"),
                PostalCode: unifiedgosdk.String("34958"),
                Region: unifiedgosdk.String("South"),
                RegionCode: unifiedgosdk.String("morph an colossal"),
            },
            CreatedAt: types.MustTimeFromString("2021-02-02T08:27:21.693Z"),
            Currency: unifiedgosdk.String("Cayman Islands Dollar"),
            Department: unifiedgosdk.String("um"),
            Division: unifiedgosdk.String("West sievert generating"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jadon_Schumm45@gmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("radian Borders Southeast"),
            LanguageLocale: unifiedgosdk.String("Silicon Awesome Industrial"),
            Name: unifiedgosdk.String("Mouse Accounts"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "ohm Southeast ROI",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("Europe/Bratislava"),
            Title: unifiedgosdk.String("Money"),
            UpdatedAt: types.MustTimeFromString("2023-12-09T10:24:50.054Z"),
        },
        ConnectionID: "competent calculate",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDUserRequest](../../models/operations/postcrmconnectioniduserrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDUserResponse](../../models/operations/postcrmconnectioniduserresponse.md), error**


## PutCrmConnectionIDCompanyID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDCompanyID(ctx, operations.PutCrmConnectionIDCompanyIDRequest{
        CrmCompany: &shared.CrmCompany{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.String("Northwest Northwest"),
                Address2: unifiedgosdk.String("portals Diesel"),
                City: unifiedgosdk.String("Azusa"),
                Country: unifiedgosdk.String("Qatar"),
                CountryCode: unifiedgosdk.String("CG"),
                PostalCode: unifiedgosdk.String("52396"),
                Region: unifiedgosdk.String("Tuna sticky lest"),
                RegionCode: unifiedgosdk.String("Soft boo Missoula"),
            },
            CreatedAt: types.MustTimeFromString("2022-05-14T19:17:30.970Z"),
            DealIds: []string{
                "Hybrid",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Vance_Cruickshank93@gmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("lest Northwest"),
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "East",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Bronze round",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-02-21T09:41:36.696Z"),
            Websites: []string{
                "Keyboard",
            },
        },
        ConnectionID: "orange Bespoke",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutCrmConnectionIDCompanyIDRequest](../../models/operations/putcrmconnectionidcompanyidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutCrmConnectionIDCompanyIDResponse](../../models/operations/putcrmconnectionidcompanyidresponse.md), error**


## PutCrmConnectionIDContactID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDContactID(ctx, operations.PutCrmConnectionIDContactIDRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.String("idolized"),
                Address2: unifiedgosdk.String("Southeast Specialist background"),
                City: unifiedgosdk.String("New Orlando"),
                Country: unifiedgosdk.String("Switzerland"),
                CountryCode: unifiedgosdk.String("GL"),
                PostalCode: unifiedgosdk.String("95864"),
                Region: unifiedgosdk.String("Intersex mmm"),
                RegionCode: unifiedgosdk.String("Specialist"),
            },
            Company: unifiedgosdk.String("Mann and Sons"),
            CompanyIds: []string{
                "impedit",
            },
            CreatedAt: types.MustTimeFromString("2023-10-28T10:36:29.710Z"),
            DealIds: []string{
                "transmitting",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Marjorie.Feeney14@hotmail.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("shuttle"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Bolivia",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Austria reinvent"),
            UpdatedAt: types.MustTimeFromString("2023-03-20T11:49:01.796Z"),
        },
        ConnectionID: "hic truck",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutCrmConnectionIDContactIDRequest](../../models/operations/putcrmconnectionidcontactidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutCrmConnectionIDContactIDResponse](../../models/operations/putcrmconnectionidcontactidresponse.md), error**


## PutCrmConnectionIDDealID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDDealID(ctx, operations.PutCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(4050.98),
            ClosedAt: types.MustTimeFromString("2022-01-15T04:05:31.641Z"),
            CreatedAt: types.MustTimeFromString("2023-06-04T01:28:32.466Z"),
            Currency: unifiedgosdk.String("Bermudian Dollar (customarily known as Bermuda Dollar)"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("laudantium Southwest"),
            Name: unifiedgosdk.String("wail Developer"),
            Pipeline: unifiedgosdk.String("male Samarium Gourde"),
            Probability: unifiedgosdk.Float64(6728.74),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("Stage Gasoline Metal"),
            Stage: unifiedgosdk.String("Corporate withdrawal Tasty"),
            Tags: []string{
                "extranet",
            },
            UpdatedAt: types.MustTimeFromString("2021-10-16T22:38:02.052Z"),
            WonReason: unifiedgosdk.String("phooey"),
        },
        ConnectionID: "Jazz",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDDealIDRequest](../../models/operations/putcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDDealIDResponse](../../models/operations/putcrmconnectioniddealidresponse.md), error**


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
    res, err := s.Crm.PutCrmConnectionIDEventID(ctx, operations.PutCrmConnectionIDEventIDRequest{
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


## PutCrmConnectionIDFileID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDFileID(ctx, operations.PutCrmConnectionIDFileIDRequest{
        CrmFile: &shared.CrmFile{
            Active: unifiedgosdk.Bool(false),
            ActivityID: unifiedgosdk.String("Cotton"),
            CompanyID: unifiedgosdk.String("Northeast"),
            ContactID: unifiedgosdk.String("Computer"),
            CreatedAt: types.MustTimeFromString("2021-04-09T13:10:01.367Z"),
            DealID: unifiedgosdk.String("toward confiscate East"),
            Description: unifiedgosdk.String("Devolved upward-trending matrices"),
            FileName: unifiedgosdk.String("generation_tactics.wav"),
            FileSize: unifiedgosdk.Float64(4770.09),
            FileType: unifiedgosdk.String("audio"),
            FileURL: unifiedgosdk.String("framework azure Metal"),
            ID: unifiedgosdk.String("<ID>"),
            LeadID: unifiedgosdk.String("ampere costume"),
            Raw: &shared.PropertyCrmFileRaw{},
            UpdatedAt: types.MustTimeFromString("2023-05-15T05:04:24.130Z"),
            UserID: unifiedgosdk.String("Research payment"),
        },
        ConnectionID: "East Associate Mazda",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDFileIDRequest](../../models/operations/putcrmconnectionidfileidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDFileIDResponse](../../models/operations/putcrmconnectionidfileidresponse.md), error**


## PutCrmConnectionIDLeadID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDLeadID(ctx, operations.PutCrmConnectionIDLeadIDRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.String("Extension"),
                Address2: unifiedgosdk.String("supposing Dorado Assistant"),
                City: unifiedgosdk.String("South Gate"),
                Country: unifiedgosdk.String("Reunion"),
                CountryCode: unifiedgosdk.String("IS"),
                PostalCode: unifiedgosdk.String("73732-2192"),
                Region: unifiedgosdk.String("JBOD phew"),
                RegionCode: unifiedgosdk.String("Southeast Framingham female"),
            },
            CompanyID: unifiedgosdk.String("deposit male"),
            ContactID: unifiedgosdk.String("bunch edge"),
            CreatedAt: types.MustTimeFromString("2021-04-03T18:08:02.798Z"),
            CreatorUserID: unifiedgosdk.String("East Panama"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Jamal20@yahoo.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("pianist"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "caricature female",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-08-09T07:11:50.077Z"),
            UserID: unifiedgosdk.String("Designer Folding"),
        },
        ConnectionID: "Lanthanum wink Regional",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDLeadIDRequest](../../models/operations/putcrmconnectionidleadidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDLeadIDResponse](../../models/operations/putcrmconnectionidleadidresponse.md), error**


## PutCrmConnectionIDPipelineID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDPipelineID(ctx, operations.PutCrmConnectionIDPipelineIDRequest{
        CrmPipeline: &shared.CrmPipeline{
            Active: unifiedgosdk.Bool(false),
            CreatedAt: types.MustTimeFromString("2021-05-16T17:24:47.805Z"),
            DealProbability: unifiedgosdk.Bool(false),
            DisplayOrder: unifiedgosdk.Float64(5470.76),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("West"),
            Raw: &shared.PropertyCrmPipelineRaw{},
            UpdatedAt: types.MustTimeFromString("2022-02-28T07:49:31.151Z"),
        },
        ConnectionID: "optimizing",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutCrmConnectionIDPipelineIDRequest](../../models/operations/putcrmconnectionidpipelineidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PutCrmConnectionIDPipelineIDResponse](../../models/operations/putcrmconnectionidpipelineidresponse.md), error**


## PutCrmConnectionIDTeamID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDTeamID(ctx, operations.PutCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2023-08-14T23:28:53.515Z"),
            Description: unifiedgosdk.String("Inverse multi-tasking task-force"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Indonesia Orchestrator Division"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-10-23T23:13:25.973Z"),
            UserIds: []string{
                "thoroughly",
            },
        },
        ConnectionID: "delectus",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDTeamIDRequest](../../models/operations/putcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDTeamIDResponse](../../models/operations/putcrmconnectionidteamidresponse.md), error**


## PutCrmConnectionIDUserID

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
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Crm.PutCrmConnectionIDUserID(ctx, operations.PutCrmConnectionIDUserIDRequest{
        CrmUser: &shared.CrmUser{
            Active: unifiedgosdk.Bool(false),
            Address: &shared.PropertyCrmUserAddress{
                Address1: unifiedgosdk.String("Honduras"),
                Address2: unifiedgosdk.String("Oxygen Libyan Burundi"),
                City: unifiedgosdk.String("North Gertrudefield"),
                Country: unifiedgosdk.String("Belgium"),
                CountryCode: unifiedgosdk.String("DK"),
                PostalCode: unifiedgosdk.String("00781"),
                Region: unifiedgosdk.String("Wagon"),
                RegionCode: unifiedgosdk.String("how overriding"),
            },
            CreatedAt: types.MustTimeFromString("2023-03-13T00:47:15.649Z"),
            Currency: unifiedgosdk.String("Pakistan Rupee"),
            Department: unifiedgosdk.String("Tricycle vaguely"),
            Division: unifiedgosdk.String("Severn bluetooth Argon"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Karl_Stehr4@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("AGP romance didactic"),
            LanguageLocale: unifiedgosdk.String("ROI Polarised"),
            Name: unifiedgosdk.String("olive synthesizing"),
            Raw: &shared.PropertyCrmUserRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "Honda Indiana",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.String("Asia/Novosibirsk"),
            Title: unifiedgosdk.String("compelling red compressing"),
            UpdatedAt: types.MustTimeFromString("2022-09-03T15:59:05.095Z"),
        },
        ConnectionID: "relationships",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDUserIDRequest](../../models/operations/putcrmconnectioniduseridrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDUserIDResponse](../../models/operations/putcrmconnectioniduseridresponse.md), error**

