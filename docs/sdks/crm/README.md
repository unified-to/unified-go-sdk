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
                Address1: unifiedgosdk.Pointer("7261 Salisbury Road"),
                Address2: unifiedgosdk.Pointer("Apt. 778"),
                City: unifiedgosdk.Pointer("Harrisburg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("56293-3678"),
                Region: unifiedgosdk.Pointer("Pennsylvania"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Balbus crapula spiculum."),
            Domains: []string{
                "fussy-nerve.info",
                "sturdy-lobster.org",
                "greedy-offset.name",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine.Jacobi@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](967.0),
            ID: unifiedgosdk.Pointer("d56641bc-d8ca-4fe5-9ff2-b3dee82d5b70"),
            Industry: unifiedgosdk.Pointer("Infrastructure"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://blue-license.org",
                "https://minor-formation.com",
                "https://ecstatic-hammock.com",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("06b20837-8835-4c6c-85b2-0cd918aa8ec4"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "esse",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "quaerat",
                "valeo",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(432) 849-2690",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(606) 871-2046",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(842) 258-9395",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Europe/San_Marino"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-24T08:27:37.566Z"),
            Websites: []string{
                "https://wise-possession.org",
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
                Address1: unifiedgosdk.Pointer("518 Brannon Burg"),
                City: unifiedgosdk.Pointer("East Helenebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("92622-2406"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            Company: unifiedgosdk.Pointer("Lowe - Jakubowski"),
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("systematic"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell45@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell90@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad_Bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("975518ba-49ce-400b-9d2a-35c7000c5a83"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/zmbPeg/2905/378"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://limited-parade.info",
                "https://faint-papa.com/",
                "https://windy-accountability.name",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("007a4fe9-3ee5-4aa5-a771-1a04e3863eaa"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "autem",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(975) 986-1658",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(489) 332-3509",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(205) 880-8886",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-23T23:36:56.134Z"),
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
            ClosedAt: types.MustNewTimeFromString("2024-03-08T12:43:23.470Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-25T13:19:40.116Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("IQD"),
            Description: unifiedgosdk.Pointer("Tabula cicuta sophismata comis tepidus sit cavus."),
            ID: unifiedgosdk.Pointer("fdf28ba0-974e-4f63-b805-b2826f981404"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("d0cea94f-26bb-42b2-b596-a2d888e70738"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Frozen Silk Chicken"),
            Pipelines: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("e5fadaa2-30c8-46ef-af75-72a5899b02a6"),
                    Name: unifiedgosdk.Pointer("trans"),
                },
            },
            Probability: unifiedgosdk.Pointer[float64](65.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Stages: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("4a713c4e-9d82-44af-8dbe-8e11dfe8bb45"),
                    Name: unifiedgosdk.Pointer("tubineus"),
                },
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("7749e3da-d8d7-482f-a06e-e1c208eb72bf"),
                    Name: unifiedgosdk.Pointer("adfectus"),
                },
            },
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-08T08:41:42.847Z"),
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
                StartAt: types.MustNewTimeFromString("2024-12-04T14:54:31.265Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("44c71a42-fc89-4911-b375-c5451eb9370a"),
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-02T11:04:05.572Z"),
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
                Address1: unifiedgosdk.Pointer("528 Forest Road"),
                Address2: unifiedgosdk.Pointer("Apt. 643"),
                City: unifiedgosdk.Pointer("Palm Springs"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55624-6499"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer16@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("02dbf5e6-f641-4f1f-84ed-00e8a8dd4c2d"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://ruddy-unblinking.info/",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("79296067-280f-4933-a004-3fde6b915e53"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("aetas"),
            Status: unifiedgosdk.Pointer("vesco"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(955) 643-9849",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(621) 811-8800",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Research Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T02:57:20.043Z"),
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
            ID: unifiedgosdk.Pointer("0b0fe66a-f844-4c41-84f5-6d2c4befc144"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("2097bb1f-e677-4900-8fb9-547ba17cc0c6"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-03T20:42:07.698Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-25T10:43:38.768Z"),
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
                Address1: unifiedgosdk.Pointer("7261 Salisbury Road"),
                Address2: unifiedgosdk.Pointer("Apt. 778"),
                City: unifiedgosdk.Pointer("Harrisburg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("56293-3678"),
                Region: unifiedgosdk.Pointer("Pennsylvania"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Balbus crapula spiculum."),
            Domains: []string{
                "fussy-nerve.info",
                "sturdy-lobster.org",
                "greedy-offset.name",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine.Jacobi@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](967.0),
            ID: unifiedgosdk.Pointer("a6a83f76-920e-4045-9781-5040b9ddec67"),
            Industry: unifiedgosdk.Pointer("Infrastructure"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://blue-license.org",
                "https://minor-formation.com",
                "https://ecstatic-hammock.com",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("00e72570-ff48-4e7a-82e5-24d572654988"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "esse",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "quaerat",
                "valeo",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(432) 849-2690",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(606) 871-2046",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(842) 258-9395",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Europe/San_Marino"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-24T08:27:37.589Z"),
            Websites: []string{
                "https://wise-possession.org",
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
                Address1: unifiedgosdk.Pointer("518 Brannon Burg"),
                City: unifiedgosdk.Pointer("East Helenebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("92622-2406"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            Company: unifiedgosdk.Pointer("Lowe - Jakubowski"),
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("systematic"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell45@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell90@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad_Bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("bcde9992-0ecc-436f-b831-b41d12fbd7bf"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/zmbPeg/2905/378"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://limited-parade.info",
                "https://faint-papa.com/",
                "https://windy-accountability.name",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("17b66175-7ce3-4dec-af9d-6a35906ac93d"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "autem",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(975) 986-1658",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(489) 332-3509",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(205) 880-8886",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-23T23:36:56.135Z"),
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
            ClosedAt: types.MustNewTimeFromString("2024-03-08T12:43:23.474Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-25T13:19:40.131Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("IQD"),
            Description: unifiedgosdk.Pointer("Tabula cicuta sophismata comis tepidus sit cavus."),
            ID: unifiedgosdk.Pointer("e3d53c4b-ecc0-485e-8a43-c0585eb93fd3"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("d908440f-3257-430f-b24e-5005e7c68c41"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Frozen Silk Chicken"),
            Pipelines: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("b2b4cd81-c8b7-40e4-b41a-c00beda94fb5"),
                    Name: unifiedgosdk.Pointer("trans"),
                },
            },
            Probability: unifiedgosdk.Pointer[float64](65.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Stages: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("ced5831d-c91d-45cb-9c03-490e845ce5d9"),
                    Name: unifiedgosdk.Pointer("tubineus"),
                },
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("c5e83990-40da-4da2-971b-8fc356a04a90"),
                    Name: unifiedgosdk.Pointer("adfectus"),
                },
            },
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-08T08:41:42.855Z"),
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
                StartAt: types.MustNewTimeFromString("2024-12-04T14:54:31.299Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("b92e1e6a-4fff-4cdf-9fc3-b0316e18bd55"),
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-02T11:04:05.620Z"),
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
                Address1: unifiedgosdk.Pointer("528 Forest Road"),
                Address2: unifiedgosdk.Pointer("Apt. 643"),
                City: unifiedgosdk.Pointer("Palm Springs"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55624-6499"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer16@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("6a93af72-3f80-4626-82ec-ddd304163896"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://ruddy-unblinking.info/",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("97bafe7b-7b3c-4686-97b4-373cf4c096c3"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("aetas"),
            Status: unifiedgosdk.Pointer("vesco"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(955) 643-9849",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(621) 811-8800",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Research Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T02:57:20.045Z"),
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
            ID: unifiedgosdk.Pointer("3e338f13-8770-4c28-b0c5-0cb090ce9d70"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("73be8808-0f6b-4241-9a1c-a4e91a486c49"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-03T20:42:07.705Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-25T10:43:38.775Z"),
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
                Address1: unifiedgosdk.Pointer("7261 Salisbury Road"),
                Address2: unifiedgosdk.Pointer("Apt. 778"),
                City: unifiedgosdk.Pointer("Harrisburg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("56293-3678"),
                Region: unifiedgosdk.Pointer("Pennsylvania"),
                RegionCode: unifiedgosdk.Pointer("ID"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-05-11T18:26:32.925Z"),
            Description: unifiedgosdk.Pointer("Balbus crapula spiculum."),
            Domains: []string{
                "fussy-nerve.info",
                "sturdy-lobster.org",
                "greedy-offset.name",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine_Jacobi@gmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Sandrine.Jacobi@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            Employees: unifiedgosdk.Pointer[float64](967.0),
            ID: unifiedgosdk.Pointer("a6a83f76-920e-4045-9781-5040b9ddec67"),
            Industry: unifiedgosdk.Pointer("Infrastructure"),
            IsActive: unifiedgosdk.Pointer(true),
            LinkUrls: []string{
                "https://blue-license.org",
                "https://minor-formation.com",
                "https://ecstatic-hammock.com",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("00e72570-ff48-4e7a-82e5-24d572654988"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "esse",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Goodwin and Sons"),
            Tags: []string{
                "quaerat",
                "valeo",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(432) 849-2690",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(606) 871-2046",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(842) 258-9395",
                    Type: shared.CrmTelephoneTypeMobile.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Europe/San_Marino"),
            UpdatedAt: types.MustNewTimeFromString("2025-02-24T08:27:37.589Z"),
            Websites: []string{
                "https://wise-possession.org",
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
                Address1: unifiedgosdk.Pointer("518 Brannon Burg"),
                City: unifiedgosdk.Pointer("East Helenebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("92622-2406"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("AZ"),
            },
            Company: unifiedgosdk.Pointer("Lowe - Jakubowski"),
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("systematic"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell45@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad.Bartell90@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Mohammad_Bartell@hotmail.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Mohammad"),
            ID: unifiedgosdk.Pointer("bcde9992-0ecc-436f-b831-b41d12fbd7bf"),
            ImageURL: unifiedgosdk.Pointer("https://picsum.photos/seed/zmbPeg/2905/378"),
            LastName: unifiedgosdk.Pointer("Bartell"),
            LinkUrls: []string{
                "https://limited-parade.info",
                "https://faint-papa.com/",
                "https://windy-accountability.name",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("17b66175-7ce3-4dec-af9d-6a35906ac93d"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "autem",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Mohammad Bartell"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(975) 986-1658",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(489) 332-3509",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(205) 880-8886",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("National Tactics Analyst"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-23T23:36:56.135Z"),
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
            ClosedAt: types.MustNewTimeFromString("2024-03-08T12:43:23.474Z"),
            ClosingAt: types.MustNewTimeFromString("2025-08-25T13:19:40.131Z"),
            CreatedAt: types.MustNewTimeFromString("2023-07-04T12:48:48.470Z"),
            Currency: unifiedgosdk.Pointer("IQD"),
            Description: unifiedgosdk.Pointer("Tabula cicuta sophismata comis tepidus sit cavus."),
            ID: unifiedgosdk.Pointer("e3d53c4b-ecc0-485e-8a43-c0585eb93fd3"),
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("d908440f-3257-430f-b24e-5005e7c68c41"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "conatus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Frozen Silk Chicken"),
            Pipelines: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("b2b4cd81-c8b7-40e4-b41a-c00beda94fb5"),
                    Name: unifiedgosdk.Pointer("trans"),
                },
            },
            Probability: unifiedgosdk.Pointer[float64](65.0),
            Source: unifiedgosdk.Pointer("cubo"),
            Stages: []shared.CrmReference{
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("ced5831d-c91d-45cb-9c03-490e845ce5d9"),
                    Name: unifiedgosdk.Pointer("tubineus"),
                },
                shared.CrmReference{
                    ID: unifiedgosdk.Pointer("c5e83990-40da-4da2-971b-8fc356a04a90"),
                    Name: unifiedgosdk.Pointer("adfectus"),
                },
            },
            Tags: []string{
                "causa",
                "suus",
            },
            UpdatedAt: types.MustNewTimeFromString("2024-10-08T08:41:42.855Z"),
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
                StartAt: types.MustNewTimeFromString("2024-12-04T14:54:31.299Z"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-07-14T04:53:23.784Z"),
            ID: unifiedgosdk.Pointer("b92e1e6a-4fff-4cdf-9fc3-b0316e18bd55"),
            Type: shared.CrmEventTypeCall.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-10-02T11:04:05.620Z"),
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
                Address1: unifiedgosdk.Pointer("528 Forest Road"),
                Address2: unifiedgosdk.Pointer("Apt. 643"),
                City: unifiedgosdk.Pointer("Palm Springs"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55624-6499"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("LA"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer16@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("Velda.Sporer@yahoo.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("6a93af72-3f80-4626-82ec-ddd304163896"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://ruddy-unblinking.info/",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("97bafe7b-7b3c-4686-97b4-373cf4c096c3"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("aetas"),
            Status: unifiedgosdk.Pointer("vesco"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(955) 643-9849",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(621) 811-8800",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Research Assistant"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-17T02:57:20.045Z"),
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
            ID: unifiedgosdk.Pointer("3e338f13-8770-4c28-b0c5-0cb090ce9d70"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Small Steel Bacon"),
            Stages: []shared.CrmStage{
                shared.CrmStage{
                    Active: unifiedgosdk.Pointer(false),
                    CreatedAt: types.MustNewTimeFromString("2022-12-28T13:45:38.446Z"),
                    DealProbability: unifiedgosdk.Pointer[float64](84.0),
                    DisplayOrder: unifiedgosdk.Pointer[float64](72.0),
                    ID: unifiedgosdk.Pointer("73be8808-0f6b-4241-9a1c-a4e91a486c49"),
                    IsClosed: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Veniam."),
                    UpdatedAt: types.MustNewTimeFromString("2025-10-03T20:42:07.705Z"),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2025-10-25T10:43:38.775Z"),
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