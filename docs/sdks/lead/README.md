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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDLeadIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.DeleteCrmConnectionIDLeadID(ctx, operations.DeleteCrmConnectionIDLeadIDRequest{
        ConnectionID: "totam",
        ID: "6c3ae7d7-b67f-4eef-9e14-2d95b1dbecef",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDLeadIDRequest](../../models/operations/deletecrmconnectionidleadidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.DeleteCrmConnectionIDLeadIDSecurity](../../models/operations/deletecrmconnectionidleadidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDLeadSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.GetCrmConnectionIDLead(ctx, operations.GetCrmConnectionIDLeadRequest{
        ConnectionID: "voluptatibus",
        Limit: unifiedto.Float64(4809.57),
        Offset: unifiedto.Float64(7789.75),
        Order: unifiedto.String("non"),
        Query: unifiedto.String("tempore"),
        Sort: unifiedto.String("quae"),
        UpdatedGte: types.MustTimeFromString("2022-08-03T04:30:42.588Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLeads != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDLeadRequest](../../models/operations/getcrmconnectionidleadrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCrmConnectionIDLeadSecurity](../../models/operations/getcrmconnectionidleadsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDLeadIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.GetCrmConnectionIDLeadID(ctx, operations.GetCrmConnectionIDLeadIDRequest{
        ConnectionID: "itaque",
        ID: "9278275e-ea76-4817-8680-63f799b7956c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDLeadIDRequest](../../models/operations/getcrmconnectionidleadidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetCrmConnectionIDLeadIDSecurity](../../models/operations/getcrmconnectionidleadidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDLeadIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.PatchCrmConnectionIDLeadID(ctx, operations.PatchCrmConnectionIDLeadIDRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedto.String("quae"),
                Address2: unifiedto.String("quidem"),
                City: unifiedto.String("Fort Maddisontown"),
                Country: unifiedto.String("Paraguay"),
                CountryCode: unifiedto.String("PY"),
                PostalCode: unifiedto.String("06209"),
                Region: unifiedto.String("reprehenderit"),
                RegionCode: unifiedto.String("quo"),
            },
            CompanyID: unifiedto.String("incidunt"),
            ContactID: unifiedto.String("id"),
            CreatedAt: types.MustTimeFromString("2021-10-20T07:58:35.149Z"),
            CreatorUserID: unifiedto.String("quaerat"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Giovanni48@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("657b01a0-7c08-4fd3-921c-257930d6f093"),
            Name: unifiedto.String("Earl Towne"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "aliquam",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-13T09:05:51.677Z"),
            UserID: unifiedto.String("autem"),
        },
        ConnectionID: "ea",
        ID: "dfa1011a-091b-43ec-8b53-862de1a9d14f",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDLeadIDRequest](../../models/operations/patchcrmconnectionidleadidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PatchCrmConnectionIDLeadIDSecurity](../../models/operations/patchcrmconnectionidleadidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDLeadSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.PostCrmConnectionIDLead(ctx, operations.PostCrmConnectionIDLeadRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedto.String("debitis"),
                Address2: unifiedto.String("reprehenderit"),
                City: unifiedto.String("Fort Fritzville"),
                Country: unifiedto.String("Bangladesh"),
                CountryCode: unifiedto.String("ZA"),
                PostalCode: unifiedto.String("02028-9722"),
                Region: unifiedto.String("laudantium"),
                RegionCode: unifiedto.String("velit"),
            },
            CompanyID: unifiedto.String("natus"),
            ContactID: unifiedto.String("molestiae"),
            CreatedAt: types.MustTimeFromString("2020-03-03T18:07:31.494Z"),
            CreatorUserID: unifiedto.String("hic"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Hanna86@yahoo.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("2090fc15-7ac9-4fe1-961c-e9be41c869dd"),
            Name: unifiedto.String("Rosemarie Moen V"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "repellendus",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2022-04-24T12:06:20.344Z"),
            UserID: unifiedto.String("odit"),
        },
        ConnectionID: "aut",
    }, operationSecurity)
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
| `request`                                                                                                | [operations.PostCrmConnectionIDLeadRequest](../../models/operations/postcrmconnectionidleadrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PostCrmConnectionIDLeadSecurity](../../models/operations/postcrmconnectionidleadsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDLeadIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Lead.PutCrmConnectionIDLeadID(ctx, operations.PutCrmConnectionIDLeadIDRequest{
        CrmLead: &shared.CrmLead{
            Active: unifiedto.Bool(false),
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedto.String("eaque"),
                Address2: unifiedto.String("deserunt"),
                City: unifiedto.String("New Vito"),
                Country: unifiedto.String("Uruguay"),
                CountryCode: unifiedto.String("ST"),
                PostalCode: unifiedto.String("63489"),
                Region: unifiedto.String("laudantium"),
                RegionCode: unifiedto.String("sapiente"),
            },
            CompanyID: unifiedto.String("facere"),
            ContactID: unifiedto.String("laudantium"),
            CreatedAt: types.MustTimeFromString("2022-09-11T02:42:21.444Z"),
            CreatorUserID: unifiedto.String("fuga"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Shannon73@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("620cd9c5-afdd-404c-b752-512beae1d87e"),
            Name: unifiedto.String("Roosevelt Hessel"),
            Raw: &shared.PropertyCrmLeadRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "quod",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            UpdatedAt: types.MustTimeFromString("2021-12-04T00:59:07.070Z"),
            UserID: unifiedto.String("eveniet"),
        },
        ConnectionID: "molestiae",
        ID: "a8831166-2cda-46d7-bc1d-86066237d422",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDLeadIDRequest](../../models/operations/putcrmconnectionidleadidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PutCrmConnectionIDLeadIDSecurity](../../models/operations/putcrmconnectionidleadidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PutCrmConnectionIDLeadIDResponse](../../models/operations/putcrmconnectionidleadidresponse.md), error**

