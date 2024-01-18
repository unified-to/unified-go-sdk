# Crm
(*Crm*)

### Available Operations

* [CreateCrmCompany](#createcrmcompany) - Create a company
* [CreateCrmContact](#createcrmcontact) - Create a contact
* [CreateCrmDeal](#createcrmdeal) - Create a deal
* [CreateCrmEvent](#createcrmevent) - Create a event
* [CreateCrmLead](#createcrmlead) - Create a lead
* [CreateCrmPipeline](#createcrmpipeline) - Create a pipeline
* [GetCrmCompany](#getcrmcompany) - Retrieve a company
* [GetCrmContact](#getcrmcontact) - Retrieve a contact
* [GetCrmDeal](#getcrmdeal) - Retrieve a deal
* [GetCrmEvent](#getcrmevent) - Retrieve a event
* [GetCrmLead](#getcrmlead) - Retrieve a lead
* [GetCrmPipeline](#getcrmpipeline) - Retrieve a pipeline
* [ListCrmCompanies](#listcrmcompanies) - List all companies
* [ListCrmContacts](#listcrmcontacts) - List all contacts
* [ListCrmDeals](#listcrmdeals) - List all deals
* [ListCrmEvents](#listcrmevents) - List all events
* [ListCrmLeads](#listcrmleads) - List all leads
* [ListCrmPipelines](#listcrmpipelines) - List all pipelines
* [PatchCrmCompany](#patchcrmcompany) - Update a company
* [PatchCrmContact](#patchcrmcontact) - Update a contact
* [PatchCrmDeal](#patchcrmdeal) - Update a deal
* [PatchCrmEvent](#patchcrmevent) - Update a event
* [PatchCrmLead](#patchcrmlead) - Update a lead
* [PatchCrmPipeline](#patchcrmpipeline) - Update a pipeline
* [RemoveCrmCompany](#removecrmcompany) - Remove a company
* [RemoveCrmContact](#removecrmcontact) - Remove a contact
* [RemoveCrmDeal](#removecrmdeal) - Remove a deal
* [RemoveCrmEvent](#removecrmevent) - Remove a event
* [RemoveCrmLead](#removecrmlead) - Remove a lead
* [RemoveCrmPipeline](#removecrmpipeline) - Remove a pipeline
* [UpdateCrmCompany](#updatecrmcompany) - Update a company
* [UpdateCrmContact](#updatecrmcontact) - Update a contact
* [UpdateCrmDeal](#updatecrmdeal) - Update a deal
* [UpdateCrmEvent](#updatecrmevent) - Update a event
* [UpdateCrmLead](#updatecrmlead) - Update a lead
* [UpdateCrmPipeline](#updatecrmpipeline) - Update a pipeline

## CreateCrmCompany

Create a company

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmCompany(ctx, operations.CreateCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            LinkUrls: []string{
                "string",
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "string",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
            Websites: []string{
                "string",
            },
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmCompanyResponse](../../pkg/models/operations/createcrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCrmContact

Create a contact

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmContact(ctx, operations.CreateCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmContactResponse](../../pkg/models/operations/createcrmcontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCrmDeal

Create a deal

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmDeal(ctx, operations.CreateCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "string",
            },
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmDealResponse](../../pkg/models/operations/createcrmdealresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCrmEvent

Create a event

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmEvent(ctx, operations.CreateCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "string",
            },
            ContactIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Email: &shared.PropertyCrmEventEmail{
                AttachmentFileIds: []string{
                    "string",
                },
                Cc: []string{
                    "string",
                },
                To: []string{
                    "string",
                },
            },
            LeadIds: []string{
                "string",
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmEventResponse](../../pkg/models/operations/createcrmeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCrmLead

Create a lead

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
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
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmLeadResponse](../../pkg/models/operations/createcrmleadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCrmPipeline

Create a pipeline

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.CreateCrmPipeline(ctx, operations.CreateCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateCrmPipelineResponse](../../pkg/models/operations/createcrmpipelineresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmCompany

Retrieve a company

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmCompany(ctx, operations.GetCrmCompanyRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                              | [operations.GetCrmCompanyRequest](../../pkg/models/operations/getcrmcompanyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetCrmCompanyResponse](../../pkg/models/operations/getcrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmContact

Retrieve a contact

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmContact(ctx, operations.GetCrmContactRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                              | [operations.GetCrmContactRequest](../../pkg/models/operations/getcrmcontactrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetCrmContactResponse](../../pkg/models/operations/getcrmcontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmDeal

Retrieve a deal

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmDeal(ctx, operations.GetCrmDealRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                        | [operations.GetCrmDealRequest](../../pkg/models/operations/getcrmdealrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetCrmDealResponse](../../pkg/models/operations/getcrmdealresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmEvent

Retrieve a event

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmEvent(ctx, operations.GetCrmEventRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                          | [operations.GetCrmEventRequest](../../pkg/models/operations/getcrmeventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetCrmEventResponse](../../pkg/models/operations/getcrmeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmLead

Retrieve a lead

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmLead(ctx, operations.GetCrmLeadRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                        | [operations.GetCrmLeadRequest](../../pkg/models/operations/getcrmleadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetCrmLeadResponse](../../pkg/models/operations/getcrmleadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmPipeline

Retrieve a pipeline

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.GetCrmPipeline(ctx, operations.GetCrmPipelineRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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
| `request`                                                                                | [operations.GetCrmPipelineRequest](../../pkg/models/operations/getcrmpipelinerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetCrmPipelineResponse](../../pkg/models/operations/getcrmpipelineresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmCompanies

List all companies

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmCompanies(ctx, operations.ListCrmCompaniesRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmCompaniesResponse](../../pkg/models/operations/listcrmcompaniesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmContacts

List all contacts

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmContacts(ctx, operations.ListCrmContactsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmContactsResponse](../../pkg/models/operations/listcrmcontactsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmDeals

List all deals

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmDeals(ctx, operations.ListCrmDealsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmDealsResponse](../../pkg/models/operations/listcrmdealsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmEvents

List all events

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmEvents(ctx, operations.ListCrmEventsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmEventsResponse](../../pkg/models/operations/listcrmeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmLeads

List all leads

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmLeads(ctx, operations.ListCrmLeadsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmLeadsResponse](../../pkg/models/operations/listcrmleadsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCrmPipelines

List all pipelines

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.ListCrmPipelines(ctx, operations.ListCrmPipelinesRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListCrmPipelinesResponse](../../pkg/models/operations/listcrmpipelinesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmCompany

Update a company

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmCompany(ctx, operations.PatchCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            LinkUrls: []string{
                "string",
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "string",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
            Websites: []string{
                "string",
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchCrmCompanyRequest](../../pkg/models/operations/patchcrmcompanyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PatchCrmCompanyResponse](../../pkg/models/operations/patchcrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmContact

Update a contact

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmContact(ctx, operations.PatchCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchCrmContactRequest](../../pkg/models/operations/patchcrmcontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PatchCrmContactResponse](../../pkg/models/operations/patchcrmcontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmDeal

Update a deal

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmDeal(ctx, operations.PatchCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "string",
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmDealRequest](../../pkg/models/operations/patchcrmdealrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PatchCrmDealResponse](../../pkg/models/operations/patchcrmdealresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmEvent

Update a event

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmEvent(ctx, operations.PatchCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "string",
            },
            ContactIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Email: &shared.PropertyCrmEventEmail{
                AttachmentFileIds: []string{
                    "string",
                },
                Cc: []string{
                    "string",
                },
                To: []string{
                    "string",
                },
            },
            LeadIds: []string{
                "string",
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchCrmEventRequest](../../pkg/models/operations/patchcrmeventrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchCrmEventResponse](../../pkg/models/operations/patchcrmeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmLead

Update a lead

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
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
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmLeadRequest](../../pkg/models/operations/patchcrmleadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PatchCrmLeadResponse](../../pkg/models/operations/patchcrmleadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmPipeline

Update a pipeline

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.PatchCrmPipeline(ctx, operations.PatchCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchCrmPipelineRequest](../../pkg/models/operations/patchcrmpipelinerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PatchCrmPipelineResponse](../../pkg/models/operations/patchcrmpipelineresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmCompany

Remove a company

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmCompany(ctx, operations.RemoveCrmCompanyRequest{
        ConnectionID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveCrmCompanyRequest](../../pkg/models/operations/removecrmcompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RemoveCrmCompanyResponse](../../pkg/models/operations/removecrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmContact

Remove a contact

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmContact(ctx, operations.RemoveCrmContactRequest{
        ConnectionID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveCrmContactRequest](../../pkg/models/operations/removecrmcontactrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RemoveCrmContactResponse](../../pkg/models/operations/removecrmcontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmDeal

Remove a deal

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmDeal(ctx, operations.RemoveCrmDealRequest{
        ConnectionID: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmDealRequest](../../pkg/models/operations/removecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.RemoveCrmDealResponse](../../pkg/models/operations/removecrmdealresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmEvent

Remove a event

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmEvent(ctx, operations.RemoveCrmEventRequest{
        ConnectionID: "string",
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
| `request`                                                                                | [operations.RemoveCrmEventRequest](../../pkg/models/operations/removecrmeventrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RemoveCrmEventResponse](../../pkg/models/operations/removecrmeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmLead

Remove a lead

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmLead(ctx, operations.RemoveCrmLeadRequest{
        ConnectionID: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmLeadRequest](../../pkg/models/operations/removecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.RemoveCrmLeadResponse](../../pkg/models/operations/removecrmleadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCrmPipeline

Remove a pipeline

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.RemoveCrmPipeline(ctx, operations.RemoveCrmPipelineRequest{
        ConnectionID: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveCrmPipelineRequest](../../pkg/models/operations/removecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.RemoveCrmPipelineResponse](../../pkg/models/operations/removecrmpipelineresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmCompany

Update a company

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmCompany(ctx, operations.UpdateCrmCompanyRequest{
        CrmCompany: &shared.CrmCompany{
            Address: &shared.PropertyCrmCompanyAddress{},
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            LinkUrls: []string{
                "string",
            },
            Raw: &shared.PropertyCrmCompanyRaw{},
            Tags: []string{
                "string",
            },
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
            Websites: []string{
                "string",
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCrmCompanyRequest](../../pkg/models/operations/updatecrmcompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateCrmCompanyResponse](../../pkg/models/operations/updatecrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmContact

Update a contact

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmContact(ctx, operations.UpdateCrmContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{},
            CompanyIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{},
            },
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCrmContactRequest](../../pkg/models/operations/updatecrmcontactrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateCrmContactResponse](../../pkg/models/operations/updatecrmcontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmDeal

Update a deal

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmDeal(ctx, operations.UpdateCrmDealRequest{
        CrmDeal: &shared.CrmDeal{
            Raw: &shared.PropertyCrmDealRaw{},
            Tags: []string{
                "string",
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmDealRequest](../../pkg/models/operations/updatecrmdealrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateCrmDealResponse](../../pkg/models/operations/updatecrmdealresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmEvent

Update a event

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmEvent(ctx, operations.UpdateCrmEventRequest{
        CrmEvent: &shared.CrmEvent{
            Call: &shared.PropertyCrmEventCall{},
            CompanyIds: []string{
                "string",
            },
            ContactIds: []string{
                "string",
            },
            DealIds: []string{
                "string",
            },
            Email: &shared.PropertyCrmEventEmail{
                AttachmentFileIds: []string{
                    "string",
                },
                Cc: []string{
                    "string",
                },
                To: []string{
                    "string",
                },
            },
            LeadIds: []string{
                "string",
            },
            Meeting: &shared.PropertyCrmEventMeeting{},
            Note: &shared.PropertyCrmEventNote{},
            Raw: &shared.PropertyCrmEventRaw{},
            Task: &shared.PropertyCrmEventTask{},
        },
        ConnectionID: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCrmEventRequest](../../pkg/models/operations/updatecrmeventrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateCrmEventResponse](../../pkg/models/operations/updatecrmeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmLead

Update a lead

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
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
                    Telephone: "string",
                },
            },
        },
        ConnectionID: "string",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmLeadRequest](../../pkg/models/operations/updatecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateCrmLeadResponse](../../pkg/models/operations/updatecrmleadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCrmPipeline

Update a pipeline

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Crm.UpdateCrmPipeline(ctx, operations.UpdateCrmPipelineRequest{
        CrmPipeline: &shared.CrmPipeline{
            Raw: &shared.PropertyCrmPipelineRaw{},
        },
        ConnectionID: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateCrmPipelineRequest](../../pkg/models/operations/updatecrmpipelinerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateCrmPipelineResponse](../../pkg/models/operations/updatecrmpipelineresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
