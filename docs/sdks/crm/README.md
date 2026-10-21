# Crm

## Overview

### Available Operations

* [CreateCrmCompany](#createcrmcompany) - Create a company
* [CreateCrmContact](#createcrmcontact) - Create a contact
* [CreateCrmDeal](#createcrmdeal) - Create a deal
* [CreateCrmEvent](#createcrmevent) - Create an event
* [CreateCrmLead](#createcrmlead) - Create a lead
* [CreateCrmPipeline](#createcrmpipeline) - Create a pipeline
* [GetCrmCompany](#getcrmcompany) - Retrieve a company
* [GetCrmContact](#getcrmcontact) - Retrieve a contact
* [GetCrmDeal](#getcrmdeal) - Retrieve a deal
* [GetCrmEvent](#getcrmevent) - Retrieve an event
* [GetCrmLead](#getcrmlead) - Retrieve a lead
* [GetCrmPipeline](#getcrmpipeline) - Retrieve a pipeline
* [ListCrmCompanies](#listcrmcompanies) - List all companies
* [ListCrmContacts](#listcrmcontacts) - List all contacts
* [ListCrmDeals](#listcrmdeals) - List all deals
* [ListCrmEvents](#listcrmevents) - List all events
* [ListCrmLeads](#listcrmleads) - List all leads
* [ListCrmPipelines](#listcrmpipelines) - List all pipelines
* [ListCrmTaxonomies](#listcrmtaxonomies) - List all taxonomies
* [PatchCrmCompany](#patchcrmcompany) - Update a company
* [PatchCrmContact](#patchcrmcontact) - Update a contact
* [PatchCrmDeal](#patchcrmdeal) - Update a deal
* [PatchCrmEvent](#patchcrmevent) - Update an event
* [PatchCrmLead](#patchcrmlead) - Update a lead
* [PatchCrmPipeline](#patchcrmpipeline) - Update a pipeline
* [RemoveCrmCompany](#removecrmcompany) - Remove a company
* [RemoveCrmContact](#removecrmcontact) - Remove a contact
* [RemoveCrmDeal](#removecrmdeal) - Remove a deal
* [RemoveCrmEvent](#removecrmevent) - Remove an event
* [RemoveCrmLead](#removecrmlead) - Remove a lead
* [RemoveCrmPipeline](#removecrmpipeline) - Remove a pipeline
* [UpdateCrmCompany](#updatecrmcompany) - Update a company
* [UpdateCrmContact](#updatecrmcontact) - Update a contact
* [UpdateCrmDeal](#updatecrmdeal) - Update a deal
* [UpdateCrmEvent](#updatecrmevent) - Update an event
* [UpdateCrmLead](#updatecrmlead) - Update a lead
* [UpdateCrmPipeline](#updatecrmpipeline) - Update a pipeline

## CreateCrmCompany

Create a company

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmCompany" method="post" path="/crm/{connection_id}/company" example="crm_company" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmCompany(ctx, operations.CreateCrmCompanyRequest{
        CrmCompany: shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.Pointer("3328 Wunsch Flat"),
                Address2: unifiedgosdk.Pointer("Suite 905"),
                City: unifiedgosdk.Pointer("New Ryderfort"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("04674"),
                Region: unifiedgosdk.Pointer("Delaware"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Venustas cruentus votum patruus terebro."),
            Domains: []string{
                "goodwinsons.com",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("info@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("sales@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("support@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](4177.0),
            ID: unifiedgosdk.Pointer("76899288-87bf-4646-974b-44a0fd1b3acd"),
            Industry: unifiedgosdk.Pointer("Metrics"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://www.linkedin.com/company/goodwin-and-sons",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("31a67396-12ff-4ae0-b56e-f60d44ae7d5d"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "unde",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "censura",
                "abstergo",
                "ipsam",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(667) 607-5872",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Atlantic/Azores"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-03T03:44:26.054Z"),
            Websites: []string{
                "https://www.goodwinsons.com",
            },
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateCrmCompanyRequest](../../pkg/models/operations/createcrmcompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateCrmCompanyResponse](../../pkg/models/operations/createcrmcompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCrmContact

Create a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmContact" method="post" path="/crm/{connection_id}/contact" example="crm_contact" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmContact(ctx, operations.CreateCrmContactRequest{
        CrmContact: shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.Pointer("65838 College Avenue"),
                Address2: unifiedgosdk.Pointer("Apt. 335"),
                City: unifiedgosdk.Pointer("Fort Alexandratown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80888-6094"),
                Region: unifiedgosdk.Pointer("Montana"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Beauty"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("mohammad.bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("300f2e85-3d9e-4075-840a-4fd9fe523962"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26803724"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/mohammad-bartell",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f5f35195-7412-4359-8cd1-f82139af7291"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "antiquus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(254) 404-3529",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(469) 989-7588",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T05:06:32.253Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateCrmContactRequest](../../pkg/models/operations/createcrmcontactrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateCrmContactResponse](../../pkg/models/operations/createcrmcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCrmDeal

Create a deal

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmDeal" method="post" path="/crm/{connection_id}/deal" example="crm_deal" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmDeal(ctx, operations.CreateCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T10:06:36.060Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-31T12:37:54.634Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("7d6cd0b2-f9c5-42c1-a68b-d59f7680ea80"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("de172c9b-9f28-4356-b8cd-d329d4ff5974"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-11T21:12:58.993Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateCrmDealRequest](../../pkg/models/operations/createcrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateCrmDealResponse](../../pkg/models/operations/createcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCrmEvent

Create an event

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmEvent" method="post" path="/crm/{connection_id}/event" example="crm_event" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmEvent(ctx, operations.CreateCrmEventRequest{
        CrmEvent: shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.Pointer("Arbitro aptus."),
                Duration: unifiedgosdk.Pointer[float64](64.0),
                StartAt: types.MustNewTimeFromString("2024-12-11T00:48:21.026Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("2861a115-b973-4d8a-8ef6-7c9a7f2efb64"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-11T12:55:45.327Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCrmEventRequest](../../pkg/models/operations/createcrmeventrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateCrmEventResponse](../../pkg/models/operations/createcrmeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCrmLead

Create a lead

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmLead" method="post" path="/crm/{connection_id}/lead" example="crm_lead" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmLead(ctx, operations.CreateCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("fdf5c447-a048-4324-91ea-918517eaf80b"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f9248f76-97df-44eb-9efe-20e1c22d7ac4"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T21:34:50.146Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateCrmLeadRequest](../../pkg/models/operations/createcrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateCrmLeadResponse](../../pkg/models/operations/createcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCrmPipeline

Create a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmPipeline" method="post" path="/crm/{connection_id}/pipeline" example="crm_pipeline" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.CreateCrmPipeline(ctx, operations.CreateCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("020faeaa-7aef-48eb-8503-9e5b08726b8e"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("431f199f-0e23-4baa-a7b9-9c583698c712"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-10T12:28:32.951Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-01T05:54:52.602Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateCrmPipelineRequest](../../pkg/models/operations/createcrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateCrmPipelineResponse](../../pkg/models/operations/createcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmCompany

Retrieve a company

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmCompany" method="get" path="/crm/{connection_id}/company/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmCompany(ctx, operations.GetCrmCompanyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                              | [operations.GetCrmCompanyRequest](../../pkg/models/operations/getcrmcompanyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetCrmCompanyResponse](../../pkg/models/operations/getcrmcompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmContact

Retrieve a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmContact" method="get" path="/crm/{connection_id}/contact/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmContact(ctx, operations.GetCrmContactRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                              | [operations.GetCrmContactRequest](../../pkg/models/operations/getcrmcontactrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetCrmContactResponse](../../pkg/models/operations/getcrmcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmDeal

Retrieve a deal

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmDeal" method="get" path="/crm/{connection_id}/deal/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmDeal(ctx, operations.GetCrmDealRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                        | [operations.GetCrmDealRequest](../../pkg/models/operations/getcrmdealrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetCrmDealResponse](../../pkg/models/operations/getcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmEvent

Retrieve an event

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmEvent" method="get" path="/crm/{connection_id}/event/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmEvent(ctx, operations.GetCrmEventRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                          | [operations.GetCrmEventRequest](../../pkg/models/operations/getcrmeventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetCrmEventResponse](../../pkg/models/operations/getcrmeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmLead

Retrieve a lead

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmLead" method="get" path="/crm/{connection_id}/lead/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmLead(ctx, operations.GetCrmLeadRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                        | [operations.GetCrmLeadRequest](../../pkg/models/operations/getcrmleadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetCrmLeadResponse](../../pkg/models/operations/getcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmPipeline

Retrieve a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmPipeline" method="get" path="/crm/{connection_id}/pipeline/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.GetCrmPipeline(ctx, operations.GetCrmPipelineRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                                | [operations.GetCrmPipelineRequest](../../pkg/models/operations/getcrmpipelinerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetCrmPipelineResponse](../../pkg/models/operations/getcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmCompanies

List all companies

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmCompanies" method="get" path="/crm/{connection_id}/company" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmCompanies(ctx, operations.ListCrmCompaniesRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListCrmCompaniesRequest](../../pkg/models/operations/listcrmcompaniesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListCrmCompaniesResponse](../../pkg/models/operations/listcrmcompaniesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmContacts

List all contacts

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmContacts" method="get" path="/crm/{connection_id}/contact" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmContacts(ctx, operations.ListCrmContactsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCrmContactsRequest](../../pkg/models/operations/listcrmcontactsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListCrmContactsResponse](../../pkg/models/operations/listcrmcontactsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmDeals

List all deals

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmDeals" method="get" path="/crm/{connection_id}/deal" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmDeals(ctx, operations.ListCrmDealsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCrmDealsRequest](../../pkg/models/operations/listcrmdealsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCrmDealsResponse](../../pkg/models/operations/listcrmdealsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmEvents

List all events

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmEvents" method="get" path="/crm/{connection_id}/event" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmEvents(ctx, operations.ListCrmEventsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCrmEventsRequest](../../pkg/models/operations/listcrmeventsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListCrmEventsResponse](../../pkg/models/operations/listcrmeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmLeads

List all leads

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmLeads" method="get" path="/crm/{connection_id}/lead" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmLeads(ctx, operations.ListCrmLeadsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCrmLeadsRequest](../../pkg/models/operations/listcrmleadsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCrmLeadsResponse](../../pkg/models/operations/listcrmleadsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmPipelines

List all pipelines

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmPipelines" method="get" path="/crm/{connection_id}/pipeline" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmPipelines(ctx, operations.ListCrmPipelinesRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListCrmPipelinesRequest](../../pkg/models/operations/listcrmpipelinesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListCrmPipelinesResponse](../../pkg/models/operations/listcrmpipelinesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmTaxonomies

List all taxonomies

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmTaxonomies" method="get" path="/crm/{connection_id}/taxonomy" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.ListCrmTaxonomies(ctx, operations.ListCrmTaxonomiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmTaxonomies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCrmTaxonomiesRequest](../../pkg/models/operations/listcrmtaxonomiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCrmTaxonomiesResponse](../../pkg/models/operations/listcrmtaxonomiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmCompany" method="patch" path="/crm/{connection_id}/company/{id}" example="crm_company" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmCompany(ctx, operations.PatchCrmCompanyRequest{
        CrmCompany: shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.Pointer("3328 Wunsch Flat"),
                Address2: unifiedgosdk.Pointer("Suite 905"),
                City: unifiedgosdk.Pointer("New Ryderfort"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("04674"),
                Region: unifiedgosdk.Pointer("Delaware"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Venustas cruentus votum patruus terebro."),
            Domains: []string{
                "goodwinsons.com",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("info@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("sales@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("support@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](4177.0),
            ID: unifiedgosdk.Pointer("cc28f72c-f444-453c-b4f6-2697156e25e2"),
            Industry: unifiedgosdk.Pointer("Metrics"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://www.linkedin.com/company/goodwin-and-sons",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("01f0ae9f-23c2-49fe-a397-7f5bda43cc3e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "unde",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "censura",
                "abstergo",
                "ipsam",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(667) 607-5872",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Atlantic/Azores"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-03T03:44:26.077Z"),
            Websites: []string{
                "https://www.goodwinsons.com",
            },
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchCrmCompanyRequest](../../pkg/models/operations/patchcrmcompanyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchCrmCompanyResponse](../../pkg/models/operations/patchcrmcompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmContact" method="patch" path="/crm/{connection_id}/contact/{id}" example="crm_contact" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmContact(ctx, operations.PatchCrmContactRequest{
        CrmContact: shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.Pointer("65838 College Avenue"),
                Address2: unifiedgosdk.Pointer("Apt. 335"),
                City: unifiedgosdk.Pointer("Fort Alexandratown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80888-6094"),
                Region: unifiedgosdk.Pointer("Montana"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Beauty"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("mohammad.bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("6ed2fb9f-37d9-4e9b-94e4-71fc8efcf77f"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26803724"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/mohammad-bartell",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("ae413485-8c49-40a3-833a-1a5f212c344f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "antiquus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(254) 404-3529",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(469) 989-7588",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T05:06:32.254Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchCrmContactRequest](../../pkg/models/operations/patchcrmcontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchCrmContactResponse](../../pkg/models/operations/patchcrmcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmDeal

Update a deal

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmDeal" method="patch" path="/crm/{connection_id}/deal/{id}" example="crm_deal" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmDeal(ctx, operations.PatchCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T10:06:36.063Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-31T12:37:54.643Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("25cb1f3c-58e5-463c-b7e6-ab7a5e2dc1f3"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("4ce364e7-3ae1-4d8c-8353-0bcb57c73575"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-11T21:12:58.999Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmDealRequest](../../pkg/models/operations/patchcrmdealrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchCrmDealResponse](../../pkg/models/operations/patchcrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmEvent" method="patch" path="/crm/{connection_id}/event/{id}" example="crm_event" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmEvent(ctx, operations.PatchCrmEventRequest{
        CrmEvent: shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.Pointer("Arbitro aptus."),
                Duration: unifiedgosdk.Pointer[float64](64.0),
                StartAt: types.MustNewTimeFromString("2024-12-11T00:48:21.064Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("fc1b8c52-fc3c-4ad5-8c05-6d3e8eb071aa"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-11T12:55:45.381Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchCrmEventRequest](../../pkg/models/operations/patchcrmeventrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchCrmEventResponse](../../pkg/models/operations/patchcrmeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmLead

Update a lead

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmLead" method="patch" path="/crm/{connection_id}/lead/{id}" example="crm_lead" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmLead(ctx, operations.PatchCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("6c8c585e-50d2-4a85-a7c8-7909fac94b4c"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("8f740a58-690e-4ddc-9d90-630a91b6caa2"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T21:34:50.148Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmLeadRequest](../../pkg/models/operations/patchcrmleadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchCrmLeadResponse](../../pkg/models/operations/patchcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmPipeline

Update a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmPipeline" method="patch" path="/crm/{connection_id}/pipeline/{id}" example="crm_pipeline" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.PatchCrmPipeline(ctx, operations.PatchCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("e646ed6a-5891-4590-920f-56531ac2c297"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("adb7be55-98ab-4e90-9d73-31fb8021e556"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-10T12:28:32.963Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-01T05:54:52.614Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchCrmPipelineRequest](../../pkg/models/operations/patchcrmpipelinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchCrmPipelineResponse](../../pkg/models/operations/patchcrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmCompany

Remove a company

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmCompany" method="delete" path="/crm/{connection_id}/company/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmCompany(ctx, operations.RemoveCrmCompanyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveCrmCompanyRequest](../../pkg/models/operations/removecrmcompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveCrmCompanyResponse](../../pkg/models/operations/removecrmcompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmContact

Remove a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmContact" method="delete" path="/crm/{connection_id}/contact/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmContact(ctx, operations.RemoveCrmContactRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveCrmContactRequest](../../pkg/models/operations/removecrmcontactrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveCrmContactResponse](../../pkg/models/operations/removecrmcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmDeal

Remove a deal

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmDeal" method="delete" path="/crm/{connection_id}/deal/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmDeal(ctx, operations.RemoveCrmDealRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmDealRequest](../../pkg/models/operations/removecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveCrmDealResponse](../../pkg/models/operations/removecrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmEvent

Remove an event

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmEvent" method="delete" path="/crm/{connection_id}/event/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmEvent(ctx, operations.RemoveCrmEventRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveCrmEventRequest](../../pkg/models/operations/removecrmeventrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveCrmEventResponse](../../pkg/models/operations/removecrmeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmLead

Remove a lead

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmLead" method="delete" path="/crm/{connection_id}/lead/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmLead(ctx, operations.RemoveCrmLeadRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmLeadRequest](../../pkg/models/operations/removecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveCrmLeadResponse](../../pkg/models/operations/removecrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmPipeline

Remove a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmPipeline" method="delete" path="/crm/{connection_id}/pipeline/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.RemoveCrmPipeline(ctx, operations.RemoveCrmPipelineRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveCrmPipelineRequest](../../pkg/models/operations/removecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveCrmPipelineResponse](../../pkg/models/operations/removecrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmCompany" method="put" path="/crm/{connection_id}/company/{id}" example="crm_company" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmCompany(ctx, operations.UpdateCrmCompanyRequest{
        CrmCompany: shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{
                Address1: unifiedgosdk.Pointer("3328 Wunsch Flat"),
                Address2: unifiedgosdk.Pointer("Suite 905"),
                City: unifiedgosdk.Pointer("New Ryderfort"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("04674"),
                Region: unifiedgosdk.Pointer("Delaware"),
                RegionCode: unifiedgosdk.Pointer("NY"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Venustas cruentus votum patruus terebro."),
            Domains: []string{
                "goodwinsons.com",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("info@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("sales@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("support@goodwinsons.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](4177.0),
            ID: unifiedgosdk.Pointer("cc28f72c-f444-453c-b4f6-2697156e25e2"),
            Industry: unifiedgosdk.Pointer("Metrics"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://www.linkedin.com/company/goodwin-and-sons",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("01f0ae9f-23c2-49fe-a397-7f5bda43cc3e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "unde",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "censura",
                "abstergo",
                "ipsam",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(667) 607-5872",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Atlantic/Azores"),
            UpdatedAt: types.MustNewTimeFromString("2025-03-03T03:44:26.077Z"),
            Websites: []string{
                "https://www.goodwinsons.com",
            },
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCrmCompanyRequest](../../pkg/models/operations/updatecrmcompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateCrmCompanyResponse](../../pkg/models/operations/updatecrmcompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmContact" method="put" path="/crm/{connection_id}/contact/{id}" example="crm_contact" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmContact(ctx, operations.UpdateCrmContactRequest{
        CrmContact: shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedgosdk.Pointer("65838 College Avenue"),
                Address2: unifiedgosdk.Pointer("Apt. 335"),
                City: unifiedgosdk.Pointer("Fort Alexandratown"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80888-6094"),
                Region: unifiedgosdk.Pointer("Montana"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Beauty"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("mohammad.bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("6ed2fb9f-37d9-4e9b-94e4-71fc8efcf77f"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26803724"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/mohammad-bartell",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("ae413485-8c49-40a3-833a-1a5f212c344f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "antiquus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(254) 404-3529",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(469) 989-7588",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T05:06:32.254Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCrmContactRequest](../../pkg/models/operations/updatecrmcontactrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateCrmContactResponse](../../pkg/models/operations/updatecrmcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmDeal

Update a deal

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmDeal" method="put" path="/crm/{connection_id}/deal/{id}" example="crm_deal" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmDeal(ctx, operations.UpdateCrmDealRequest{
        CrmDeal: shared.CrmDeal{
            Amount: unifiedgosdk.Pointer[float64](98162.0),
            ClosedAt: types.MustNewTimeFromString("2024-03-10T10:06:36.063Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-31T12:37:54.643Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("25cb1f3c-58e5-463c-b7e6-ab7a5e2dc1f3"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("4ce364e7-3ae1-4d8c-8353-0bcb57c73575"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Stoltenberg - Funk - Luxurious Rubber Chicken"),
            Probability: unifiedgosdk.Pointer[float64](42.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-11T21:12:58.999Z"),
            WonReason: unifiedgosdk.Pointer("Usque libero soleo."),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmDealRequest](../../pkg/models/operations/updatecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateCrmDealResponse](../../pkg/models/operations/updatecrmdealresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmEvent" method="put" path="/crm/{connection_id}/event/{id}" example="crm_event" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmEvent(ctx, operations.UpdateCrmEventRequest{
        CrmEvent: shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{
                Description: unifiedgosdk.Pointer("Arbitro aptus."),
                Duration: unifiedgosdk.Pointer[float64](64.0),
                StartAt: types.MustNewTimeFromString("2024-12-11T00:48:21.064Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("fc1b8c52-fc3c-4ad5-8c05-6d3e8eb071aa"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-11T12:55:45.381Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCrmEventRequest](../../pkg/models/operations/updatecrmeventrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateCrmEventResponse](../../pkg/models/operations/updatecrmeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmLead

Update a lead

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmLead" method="put" path="/crm/{connection_id}/lead/{id}" example="crm_lead" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmLead(ctx, operations.UpdateCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("6c8c585e-50d2-4a85-a7c8-7909fac94b4c"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("8f740a58-690e-4ddc-9d90-630a91b6caa2"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T21:34:50.148Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmLeadRequest](../../pkg/models/operations/updatecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateCrmLeadResponse](../../pkg/models/operations/updatecrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmPipeline

Update a pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmPipeline" method="put" path="/crm/{connection_id}/pipeline/{id}" example="crm_pipeline" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Crm.UpdateCrmPipeline(ctx, operations.UpdateCrmPipelineRequest{
        CrmPipeline: shared.CrmPipeline{
            CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
            DealProbability: unifiedgosdk.Pointer[float64](99.0),
            DisplayOrder: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("e646ed6a-5891-4590-920f-56531ac2c297"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("adb7be55-98ab-4e90-9d73-31fb8021e556"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-10T12:28:32.963Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-01T05:54:52.614Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateCrmPipelineRequest](../../pkg/models/operations/updatecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateCrmPipelineResponse](../../pkg/models/operations/updatecrmpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |