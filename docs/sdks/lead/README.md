# Lead
(*Lead*)

### Available Operations

* [DeleteCrmConnectionIDLeadID](#deletecrmconnectionidleadid) - Remove a lead
* [GetCrmConnectionIDLead](#getcrmconnectionidlead) - List all leads
* [GetCrmConnectionIDLeadID](#getcrmconnectionidleadid) - Retrieve a lead
* [PatchCrmConnectionIDLeadID](#patchcrmconnectionidleadid) - Update a lead
* [PostCrmConnectionIDLead](#postcrmconnectionidlead) - Create a lead
* [PutCrmConnectionIDLeadID](#putcrmconnectionidleadid) - Update a lead

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
    res, err := s.Lead.DeleteCrmConnectionIDLeadID(ctx, operations.DeleteCrmConnectionIDLeadIDRequest{
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
    res, err := s.Lead.GetCrmConnectionIDLead(ctx, operations.GetCrmConnectionIDLeadRequest{
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
    res, err := s.Lead.GetCrmConnectionIDLeadID(ctx, operations.GetCrmConnectionIDLeadIDRequest{
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
    res, err := s.Lead.PatchCrmConnectionIDLeadID(ctx, operations.PatchCrmConnectionIDLeadIDRequest{
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
    res, err := s.Lead.PostCrmConnectionIDLead(ctx, operations.PostCrmConnectionIDLeadRequest{
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
    res, err := s.Lead.PutCrmConnectionIDLeadID(ctx, operations.PutCrmConnectionIDLeadIDRequest{
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

