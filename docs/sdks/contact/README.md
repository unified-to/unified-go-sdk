# Contact

## Overview

### Available Operations

* [CreateAccountingContact](#createaccountingcontact) - Create a contact
* [CreateCrmContact](#createcrmcontact) - Create a contact
* [CreateUcContact](#createuccontact) - Create a contact
* [GetAccountingContact](#getaccountingcontact) - Retrieve a contact
* [GetCrmContact](#getcrmcontact) - Retrieve a contact
* [GetUcContact](#getuccontact) - Retrieve a contact
* [ListAccountingContacts](#listaccountingcontacts) - List all contacts
* [ListCrmContacts](#listcrmcontacts) - List all contacts
* [ListUcContacts](#listuccontacts) - List all contacts
* [PatchAccountingContact](#patchaccountingcontact) - Update a contact
* [PatchCrmContact](#patchcrmcontact) - Update a contact
* [PatchUcContact](#patchuccontact) - Update a contact
* [RemoveAccountingContact](#removeaccountingcontact) - Remove a contact
* [RemoveCrmContact](#removecrmcontact) - Remove a contact
* [RemoveUcContact](#removeuccontact) - Remove a contact
* [UpdateAccountingContact](#updateaccountingcontact) - Update a contact
* [UpdateCrmContact](#updatecrmcontact) - Update a contact
* [UpdateUcContact](#updateuccontact) - Update a contact

## CreateAccountingContact

Create a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingContact" method="post" path="/accounting/{connection_id}/contact" example="accounting_contact" -->
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

    res, err := s.Contact.CreateAccountingContact(ctx, operations.CreateAccountingContactRequest{
        AccountingContact: shared.AccountingContact{
            AssociatedContacts: []shared.AccountingAssociatedContact{
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("86f0d281-6653-4045-b59d-ecaa1314660e"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("c12c5093-9ff7-4bd5-90f7-40c5432788a3"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
            },
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("2633 Stoney Lane"),
                Address2: unifiedgosdk.Pointer("Suite 176"),
                City: unifiedgosdk.Pointer("Ladariusboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Marquardt Inc"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("ISK"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("Delores.Reynolds10@hotmail.com"),
                    Type: shared.AccountingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("e33b5080-ca73-4466-9c34-8863505a6e12"),
            Identification: unifiedgosdk.Pointer("amicitia"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Name: unifiedgosdk.Pointer("Delores Reynolds"),
            Notes: unifiedgosdk.Pointer("Caput accusamus et videlicet."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("97d70fc5-0792-4e19-9b40-4bbb9bcbc2d6"),
                    Name: unifiedgosdk.Pointer("Visa 1234"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://scented-t-shirt.info/"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Renner Stream"),
                Address2: unifiedgosdk.Pointer("Apt. 555"),
                City: unifiedgosdk.Pointer("Roanoke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80686-7556"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionResale.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("amplexus"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(427) 701-7160"),
                    Type: shared.AccountingTelephoneTypeHome.ToPointer(),
                },
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(540) 913-9171"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T13:07:53.829Z"),
            Website: unifiedgosdk.Pointer("https://noxious-advertisement.org"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingContactRequest](../../pkg/models/operations/createaccountingcontactrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingContactResponse](../../pkg/models/operations/createaccountingcontactresponse.md), error**

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

    res, err := s.Contact.CreateCrmContact(ctx, operations.CreateCrmContactRequest{
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
            ID: unifiedgosdk.Pointer("e5e6187a-df1b-405e-bd64-3409f1989572"),
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
                    ID: unifiedgosdk.Pointer("2fbd4b85-41b9-4a4d-9fcc-253cc4a1f426"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T02:00:58.918Z"),
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

## CreateUcContact

Create a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcContact" method="post" path="/uc/{connection_id}/contact" example="uc_contact" -->
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

    res, err := s.Contact.CreateUcContact(ctx, operations.CreateUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("d48cba89-d1d1-4c18-bea4-50564c537034"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-05T08:33:34.515Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateUcContactRequest](../../pkg/models/operations/createuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateUcContactResponse](../../pkg/models/operations/createuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingContact

Retrieve a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingContact" method="get" path="/accounting/{connection_id}/contact/{id}" -->
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

    res, err := s.Contact.GetAccountingContact(ctx, operations.GetAccountingContactRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAccountingContactRequest](../../pkg/models/operations/getaccountingcontactrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAccountingContactResponse](../../pkg/models/operations/getaccountingcontactresponse.md), error**

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

    res, err := s.Contact.GetCrmContact(ctx, operations.GetCrmContactRequest{
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

## GetUcContact

Retrieve a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcContact" method="get" path="/uc/{connection_id}/contact/{id}" -->
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

    res, err := s.Contact.GetUcContact(ctx, operations.GetUcContactRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetUcContactRequest](../../pkg/models/operations/getuccontactrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetUcContactResponse](../../pkg/models/operations/getuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingContacts

List all contacts

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingContacts" method="get" path="/accounting/{connection_id}/contact" -->
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

    res, err := s.Contact.ListAccountingContacts(ctx, operations.ListAccountingContactsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAccountingContactsRequest](../../pkg/models/operations/listaccountingcontactsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAccountingContactsResponse](../../pkg/models/operations/listaccountingcontactsresponse.md), error**

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

    res, err := s.Contact.ListCrmContacts(ctx, operations.ListCrmContactsRequest{
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

## ListUcContacts

List all contacts

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcContacts" method="get" path="/uc/{connection_id}/contact" -->
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

    res, err := s.Contact.ListUcContacts(ctx, operations.ListUcContactsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListUcContactsRequest](../../pkg/models/operations/listuccontactsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListUcContactsResponse](../../pkg/models/operations/listuccontactsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingContact" method="patch" path="/accounting/{connection_id}/contact/{id}" example="accounting_contact" -->
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

    res, err := s.Contact.PatchAccountingContact(ctx, operations.PatchAccountingContactRequest{
        AccountingContact: shared.AccountingContact{
            AssociatedContacts: []shared.AccountingAssociatedContact{
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("5711d916-2111-477a-a562-25d41fd5b2bb"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("731c34d2-3a7d-40ee-9b56-352b72062b8a"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
            },
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("2633 Stoney Lane"),
                Address2: unifiedgosdk.Pointer("Suite 176"),
                City: unifiedgosdk.Pointer("Ladariusboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Marquardt Inc"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("ISK"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("Delores.Reynolds10@hotmail.com"),
                    Type: shared.AccountingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("3c0ebe8c-1b78-452d-be0a-2de8a632eeee"),
            Identification: unifiedgosdk.Pointer("amicitia"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Name: unifiedgosdk.Pointer("Delores Reynolds"),
            Notes: unifiedgosdk.Pointer("Caput accusamus et videlicet."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("f62e0b00-e61e-437d-8b6d-0070c2519c02"),
                    Name: unifiedgosdk.Pointer("Visa 1234"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://scented-t-shirt.info/"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Renner Stream"),
                Address2: unifiedgosdk.Pointer("Apt. 555"),
                City: unifiedgosdk.Pointer("Roanoke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80686-7556"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionResale.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("amplexus"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(427) 701-7160"),
                    Type: shared.AccountingTelephoneTypeHome.ToPointer(),
                },
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(540) 913-9171"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T13:07:53.853Z"),
            Website: unifiedgosdk.Pointer("https://noxious-advertisement.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAccountingContactRequest](../../pkg/models/operations/patchaccountingcontactrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAccountingContactResponse](../../pkg/models/operations/patchaccountingcontactresponse.md), error**

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

    res, err := s.Contact.PatchCrmContact(ctx, operations.PatchCrmContactRequest{
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
            ID: unifiedgosdk.Pointer("e10b7b3c-9687-448c-b29b-54392feb8e80"),
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
                    ID: unifiedgosdk.Pointer("3c745edc-9ced-48e3-9d4c-f8235fb4a78d"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T02:00:58.918Z"),
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

## PatchUcContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcContact" method="patch" path="/uc/{connection_id}/contact/{id}" example="uc_contact" -->
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

    res, err := s.Contact.PatchUcContact(ctx, operations.PatchUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("d6d33397-d74c-4da5-ac43-49b9a7e35dc6"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-05T08:33:34.527Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchUcContactRequest](../../pkg/models/operations/patchuccontactrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchUcContactResponse](../../pkg/models/operations/patchuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingContact

Remove a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingContact" method="delete" path="/accounting/{connection_id}/contact/{id}" -->
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

    res, err := s.Contact.RemoveAccountingContact(ctx, operations.RemoveAccountingContactRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAccountingContactRequest](../../pkg/models/operations/removeaccountingcontactrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAccountingContactResponse](../../pkg/models/operations/removeaccountingcontactresponse.md), error**

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

    res, err := s.Contact.RemoveCrmContact(ctx, operations.RemoveCrmContactRequest{
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

## RemoveUcContact

Remove a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUcContact" method="delete" path="/uc/{connection_id}/contact/{id}" -->
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

    res, err := s.Contact.RemoveUcContact(ctx, operations.RemoveUcContactRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveUcContactRequest](../../pkg/models/operations/removeuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveUcContactResponse](../../pkg/models/operations/removeuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingContact" method="put" path="/accounting/{connection_id}/contact/{id}" example="accounting_contact" -->
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

    res, err := s.Contact.UpdateAccountingContact(ctx, operations.UpdateAccountingContactRequest{
        AccountingContact: shared.AccountingContact{
            AssociatedContacts: []shared.AccountingAssociatedContact{
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("5711d916-2111-477a-a562-25d41fd5b2bb"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
                shared.AccountingAssociatedContact{
                    ID: unifiedgosdk.Pointer("731c34d2-3a7d-40ee-9b56-352b72062b8a"),
                    Name: unifiedgosdk.Pointer("Delores Reynolds"),
                },
            },
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("2633 Stoney Lane"),
                Address2: unifiedgosdk.Pointer("Suite 176"),
                City: unifiedgosdk.Pointer("Ladariusboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("NV"),
            },
            CompanyName: unifiedgosdk.Pointer("Marquardt Inc"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("ISK"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("Delores.Reynolds10@hotmail.com"),
                    Type: shared.AccountingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("3c0ebe8c-1b78-452d-be0a-2de8a632eeee"),
            Identification: unifiedgosdk.Pointer("amicitia"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Name: unifiedgosdk.Pointer("Delores Reynolds"),
            Notes: unifiedgosdk.Pointer("Caput accusamus et videlicet."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("f62e0b00-e61e-437d-8b6d-0070c2519c02"),
                    Name: unifiedgosdk.Pointer("Visa 1234"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://scented-t-shirt.info/"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Renner Stream"),
                Address2: unifiedgosdk.Pointer("Apt. 555"),
                City: unifiedgosdk.Pointer("Roanoke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("80686-7556"),
                Region: unifiedgosdk.Pointer("Vermont"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionResale.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("amplexus"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(427) 701-7160"),
                    Type: shared.AccountingTelephoneTypeHome.ToPointer(),
                },
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(540) 913-9171"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T13:07:53.853Z"),
            Website: unifiedgosdk.Pointer("https://noxious-advertisement.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAccountingContactRequest](../../pkg/models/operations/updateaccountingcontactrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAccountingContactResponse](../../pkg/models/operations/updateaccountingcontactresponse.md), error**

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

    res, err := s.Contact.UpdateCrmContact(ctx, operations.UpdateCrmContactRequest{
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
            ID: unifiedgosdk.Pointer("e10b7b3c-9687-448c-b29b-54392feb8e80"),
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
                    ID: unifiedgosdk.Pointer("3c745edc-9ced-48e3-9d4c-f8235fb4a78d"),
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
            UpdatedAt: types.MustNewTimeFromString("2021-02-24T02:00:58.918Z"),
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

## UpdateUcContact

Update a contact

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcContact" method="put" path="/uc/{connection_id}/contact/{id}" example="uc_contact" -->
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

    res, err := s.Contact.UpdateUcContact(ctx, operations.UpdateUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("d6d33397-d74c-4da5-ac43-49b9a7e35dc6"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-05T08:33:34.527Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateUcContactRequest](../../pkg/models/operations/updateuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateUcContactResponse](../../pkg/models/operations/updateuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |