# Lead

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
        ConnectionID: "totam",
        ID: "6c3ae7d7-b67f-4eef-9e14-2d95b1dbecef",
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
        ConnectionID: "voluptatibus",
        Limit: unifiedgosdk.Float64(4809.57),
        Offset: unifiedgosdk.Float64(7789.75),
        Order: unifiedgosdk.String("non"),
        Query: unifiedgosdk.String("tempore"),
        Sort: unifiedgosdk.String("quae"),
        UpdatedGte: types.MustTimeFromString("2022-08-03T04:30:42.588Z"),
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
        ConnectionID: "itaque",
        ID: "9278275e-ea76-4817-8680-63f799b7956c",
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
                Address1: unifiedgosdk.String("quae"),
                Address2: unifiedgosdk.String("quidem"),
                City: unifiedgosdk.String("Fort Maddisontown"),
                Country: unifiedgosdk.String("Paraguay"),
                CountryCode: unifiedgosdk.String("PY"),
                PostalCode: unifiedgosdk.String("06209"),
                Region: unifiedgosdk.String("reprehenderit"),
                RegionCode: unifiedgosdk.String("quo"),
            },
            CompanyID: unifiedgosdk.String("incidunt"),
            ContactID: unifiedgosdk.String("id"),
            CreatedAt: types.MustTimeFromString("2021-10-20T07:58:35.149Z"),
            CreatorUserID: unifiedgosdk.String("quaerat"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Giovanni48@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("657b01a0-7c08-4fd3-921c-257930d6f093"),
            Name: unifiedgosdk.String("Earl Towne"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "aliquam",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-13T09:05:51.677Z"),
            UserID: unifiedgosdk.String("autem"),
        },
        ConnectionID: "ea",
        ID: "dfa1011a-091b-43ec-8b53-862de1a9d14f",
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
                Address1: unifiedgosdk.String("debitis"),
                Address2: unifiedgosdk.String("reprehenderit"),
                City: unifiedgosdk.String("Fort Fritzville"),
                Country: unifiedgosdk.String("Bangladesh"),
                CountryCode: unifiedgosdk.String("ZA"),
                PostalCode: unifiedgosdk.String("02028-9722"),
                Region: unifiedgosdk.String("laudantium"),
                RegionCode: unifiedgosdk.String("velit"),
            },
            CompanyID: unifiedgosdk.String("natus"),
            ContactID: unifiedgosdk.String("molestiae"),
            CreatedAt: types.MustTimeFromString("2020-03-03T18:07:31.494Z"),
            CreatorUserID: unifiedgosdk.String("hic"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Hanna86@yahoo.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("2090fc15-7ac9-4fe1-961c-e9be41c869dd"),
            Name: unifiedgosdk.String("Rosemarie Moen V"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "repellendus",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-24T12:06:20.344Z"),
            UserID: unifiedgosdk.String("odit"),
        },
        ConnectionID: "aut",
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
                Address1: unifiedgosdk.String("eaque"),
                Address2: unifiedgosdk.String("deserunt"),
                City: unifiedgosdk.String("New Vito"),
                Country: unifiedgosdk.String("Uruguay"),
                CountryCode: unifiedgosdk.String("ST"),
                PostalCode: unifiedgosdk.String("63489"),
                Region: unifiedgosdk.String("laudantium"),
                RegionCode: unifiedgosdk.String("sapiente"),
            },
            CompanyID: unifiedgosdk.String("facere"),
            ContactID: unifiedgosdk.String("laudantium"),
            CreatedAt: types.MustTimeFromString("2022-09-11T02:42:21.444Z"),
            CreatorUserID: unifiedgosdk.String("fuga"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.String("Shannon73@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedgosdk.String("620cd9c5-afdd-404c-b752-512beae1d87e"),
            Name: unifiedgosdk.String("Roosevelt Hessel"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "quod",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-12-04T00:59:07.070Z"),
            UserID: unifiedgosdk.String("eveniet"),
        },
        ConnectionID: "molestiae",
        ID: "a8831166-2cda-46d7-bc1d-86066237d422",
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

