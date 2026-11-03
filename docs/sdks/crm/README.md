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

Creates a CrmCompany. Provide the unified CrmCompany fields in the request body; they are mapped to the provider's API. Returns the created CrmCompany.

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
            ID: unifiedgosdk.Pointer("ef474737-c201-4161-a700-901ac48d12a4"),
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
                    ID: unifiedgosdk.Pointer("598d2e62-c411-4ddc-89d5-e901b41f13f0"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-12T20:05:24.373Z"),
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

Creates a CrmContact. Provide the unified CrmContact fields in the request body; they are mapped to the provider's API. Returns the created CrmContact.

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
            ID: unifiedgosdk.Pointer("f1f3892e-513e-499a-9e2f-dcc8bf68a871"),
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
                    ID: unifiedgosdk.Pointer("253f34ba-dcf2-463d-a428-c10bcafce97e"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T12:55:33.900Z"),
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

Creates a CrmDeal. Provide the unified CrmDeal fields in the request body; they are mapped to the provider's API. Returns the created CrmDeal.

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
            ClosedAt: types.MustNewTimeFromString("2024-03-13T02:41:45.600Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-09T00:33:18.481Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("be0e3d5a-bf9b-4235-946d-c209f84c28be"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("077abf4b-e542-4ee0-bf68-73f0ffe27ec2"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-10-16T21:29:27.458Z"),
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

Creates a CrmEvent. Provide the unified CrmEvent fields in the request body; they are mapped to the provider's API. Returns the created CrmEvent.

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
                StartAt: types.MustNewTimeFromString("2024-12-20T03:48:11.828Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("3cd5e602-53c3-41fb-9928-0ef144156b16"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-24T10:56:53.011Z"),
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

Creates a CrmLead. Provide the unified CrmLead fields in the request body; they are mapped to the provider's API. Returns the created CrmLead.

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
            ID: unifiedgosdk.Pointer("1cd8916f-db77-4742-a4eb-1a6859ba0cbb"),
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
                    ID: unifiedgosdk.Pointer("05a1399b-8cc3-4cf7-b94d-5cc053d67e15"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-05-19T00:05:03.222Z"),
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

Creates a CrmPipeline. Provide the unified CrmPipeline fields in the request body; they are mapped to the provider's API. Returns the created CrmPipeline.

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
            ID: unifiedgosdk.Pointer("fe5b32d5-f6ef-4264-9916-eb5876761d74"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("8874d26b-289b-4422-a5d7-130f91320719"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-19T23:50:08.255Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-10T22:07:54.700Z"),
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

Retrieves a single CrmCompany by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CrmContact by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CrmDeal by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CrmEvent by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CrmLead by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single CrmPipeline by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of CrmCompanies. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmContacts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmDeals. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmEvents. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmLeads. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmPipelines. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of CrmTaxonomies. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the CrmCompany identified by `id`. Provide the unified CrmCompany fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("612323be-ceb5-42f0-9c4d-3d59ba09aa85"),
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
                    ID: unifiedgosdk.Pointer("88b84323-4eeb-4c9d-b78f-7538dbfabee6"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-12T20:05:24.401Z"),
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

Updates the CrmContact identified by `id`. Provide the unified CrmContact fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("2d9298af-e922-43a2-8579-cdec8e056db4"),
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
                    ID: unifiedgosdk.Pointer("24dbc12c-ed78-4cba-8316-9347178d9226"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T12:55:33.900Z"),
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

Updates the CrmDeal identified by `id`. Provide the unified CrmDeal fields to change in the request body.

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
            ClosedAt: types.MustNewTimeFromString("2024-03-13T02:41:45.605Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-09T00:33:18.498Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("a363d93c-1442-42ad-af79-7350660957a5"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("12df402a-02fe-4dec-b963-b2ead2eed1d1"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-10-16T21:29:27.467Z"),
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

Updates the CrmEvent identified by `id`. Provide the unified CrmEvent fields to change in the request body.

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
                StartAt: types.MustNewTimeFromString("2024-12-20T03:48:11.868Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("2aa52f86-09d8-4101-a41b-1226a0383b75"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-24T10:56:53.066Z"),
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

Updates the CrmLead identified by `id`. Provide the unified CrmLead fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("d3c4a023-31ec-4952-b601-1bb575a0fa06"),
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
                    ID: unifiedgosdk.Pointer("a62e75b7-b6ad-47f0-acdf-7fb34b624abd"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-05-19T00:05:03.224Z"),
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

Updates the CrmPipeline identified by `id`. Provide the unified CrmPipeline fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("d81bd107-8945-4f7a-8cdd-8bf9b8d327ed"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("c36a3197-faf1-45b2-ba0b-ec5c4ab0a4fb"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-19T23:50:08.267Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-10T22:07:54.713Z"),
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

Deletes the CrmCompany identified by `id`.

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

Deletes the CrmContact identified by `id`.

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

Deletes the CrmDeal identified by `id`.

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

Deletes the CrmEvent identified by `id`.

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

Deletes the CrmLead identified by `id`.

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

Deletes the CrmPipeline identified by `id`.

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

Updates the CrmCompany identified by `id`. Provide the unified CrmCompany fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("612323be-ceb5-42f0-9c4d-3d59ba09aa85"),
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
                    ID: unifiedgosdk.Pointer("88b84323-4eeb-4c9d-b78f-7538dbfabee6"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-03-12T20:05:24.401Z"),
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

Updates the CrmContact identified by `id`. Provide the unified CrmContact fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("2d9298af-e922-43a2-8579-cdec8e056db4"),
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
                    ID: unifiedgosdk.Pointer("24dbc12c-ed78-4cba-8316-9347178d9226"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T12:55:33.900Z"),
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

Updates the CrmDeal identified by `id`. Provide the unified CrmDeal fields to change in the request body.

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
            ClosedAt: types.MustNewTimeFromString("2024-03-13T02:41:45.605Z"),
            ClosingAt: types.MustNewTimeFromString("2025-09-09T00:33:18.498Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Comis tepidus sit cavus crustulum trans tubineus adfectus."),
            ID: unifiedgosdk.Pointer("a363d93c-1442-42ad-af79-7350660957a5"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("12df402a-02fe-4dec-b963-b2ead2eed1d1"),
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
            UpdatedAt: types.MustNewTimeFromString("2024-10-16T21:29:27.467Z"),
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

Updates the CrmEvent identified by `id`. Provide the unified CrmEvent fields to change in the request body.

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
                StartAt: types.MustNewTimeFromString("2024-12-20T03:48:11.868Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("2aa52f86-09d8-4101-a41b-1226a0383b75"),
            Metadata: []shared.CrmMetadata{},
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-24T10:56:53.066Z"),
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

Updates the CrmLead identified by `id`. Provide the unified CrmLead fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("d3c4a023-31ec-4952-b601-1bb575a0fa06"),
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
                    ID: unifiedgosdk.Pointer("a62e75b7-b6ad-47f0-acdf-7fb34b624abd"),
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
            UpdatedAt: types.MustNewTimeFromString("2020-05-19T00:05:03.224Z"),
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

Updates the CrmPipeline identified by `id`. Provide the unified CrmPipeline fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("d81bd107-8945-4f7a-8cdd-8bf9b8d327ed"),
            IsActive: unifiedgosdk.Pointer(true),
            Metadata: []shared.CrmMetadata{},
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("c36a3197-faf1-45b2-ba0b-ec5c4ab0a4fb"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-19T23:50:08.267Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-11-10T22:07:54.713Z"),
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