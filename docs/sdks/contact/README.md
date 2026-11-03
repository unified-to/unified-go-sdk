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

Creates a AccountingContact. Provide the unified AccountingContact fields in the request body; they are mapped to the provider's API. Returns the created AccountingContact.

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
            AssociatedContacts: []shared.AccountingAssociatedContact{},
            BalanceAmount: unifiedgosdk.Pointer[float64](55574.0),
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("4091 Tromp Coves"),
                City: unifiedgosdk.Pointer("Lake Emmybury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("50764-5226"),
                Region: unifiedgosdk.Pointer("Washington"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            CompanyName: unifiedgosdk.Pointer("Kertzmann Group"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("delores.reynolds@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("96e76755-4955-40e7-9fd3-d7a1d82162b3"),
            Identification: unifiedgosdk.Pointer("531850872"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Kertzmann Group"),
            Notes: unifiedgosdk.Pointer("Et videlicet pauci."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("d154ebb9-ae44-4870-ab0e-c30c3ef8344e"),
                    Name: unifiedgosdk.Pointer("Visa 7844"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://kertzmanngroup.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("23129 Rebekah Mews"),
                City: unifiedgosdk.Pointer("Mortimerhaven"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55878-6806"),
                Region: unifiedgosdk.Pointer("Florida"),
                RegionCode: unifiedgosdk.Pointer("OK"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("867556410"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(432) 760-1716"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-25T17:42:49.966Z"),
            Website: unifiedgosdk.Pointer("https://kertzmanngroup.com"),
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

    res, err := s.Contact.CreateCrmContact(ctx, operations.CreateCrmContactRequest{
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

## CreateUcContact

Creates a UcContact. Provide the unified UcContact fields in the request body; they are mapped to the provider's API. Returns the created UcContact.

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
            ID: unifiedgosdk.Pointer("6aabfa05-d631-4d33-8e11-fe1ca3d169c6"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
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
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T00:29:55.939Z"),
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

Retrieves a single AccountingContact by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single UcContact by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of AccountingContacts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of UcContacts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the AccountingContact identified by `id`. Provide the unified AccountingContact fields to change in the request body.

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
            AssociatedContacts: []shared.AccountingAssociatedContact{},
            BalanceAmount: unifiedgosdk.Pointer[float64](55574.0),
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("4091 Tromp Coves"),
                City: unifiedgosdk.Pointer("Lake Emmybury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("50764-5226"),
                Region: unifiedgosdk.Pointer("Washington"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            CompanyName: unifiedgosdk.Pointer("Kertzmann Group"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("delores.reynolds@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("b995c64f-b334-4d71-9499-fe483fa80c6e"),
            Identification: unifiedgosdk.Pointer("531850872"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Kertzmann Group"),
            Notes: unifiedgosdk.Pointer("Et videlicet pauci."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("0834d938-340a-4eb6-8629-f356a1340133"),
                    Name: unifiedgosdk.Pointer("Visa 7844"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://kertzmanngroup.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("23129 Rebekah Mews"),
                City: unifiedgosdk.Pointer("Mortimerhaven"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55878-6806"),
                Region: unifiedgosdk.Pointer("Florida"),
                RegionCode: unifiedgosdk.Pointer("OK"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("867556410"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(432) 760-1716"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-25T17:42:49.991Z"),
            Website: unifiedgosdk.Pointer("https://kertzmanngroup.com"),
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

    res, err := s.Contact.PatchCrmContact(ctx, operations.PatchCrmContactRequest{
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

## PatchUcContact

Updates the UcContact identified by `id`. Provide the unified UcContact fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("9382c547-dae9-4d0d-b495-8ae21f1a5635"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
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
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T00:29:55.950Z"),
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

Deletes the AccountingContact identified by `id`.

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

Deletes the UcContact identified by `id`.

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

Updates the AccountingContact identified by `id`. Provide the unified AccountingContact fields to change in the request body.

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
            AssociatedContacts: []shared.AccountingAssociatedContact{},
            BalanceAmount: unifiedgosdk.Pointer[float64](55574.0),
            BillingAddress: &shared.PropertyAccountingContactBillingAddress{
                Address1: unifiedgosdk.Pointer("4091 Tromp Coves"),
                City: unifiedgosdk.Pointer("Lake Emmybury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("50764-5226"),
                Region: unifiedgosdk.Pointer("Washington"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            CompanyName: unifiedgosdk.Pointer("Kertzmann Group"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("delores.reynolds@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@kertzmanngroup.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Delores"),
            ID: unifiedgosdk.Pointer("b995c64f-b334-4d71-9499-fe483fa80c6e"),
            Identification: unifiedgosdk.Pointer("531850872"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Reynolds"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Kertzmann Group"),
            Notes: unifiedgosdk.Pointer("Et videlicet pauci."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("0834d938-340a-4eb6-8629-f356a1340133"),
                    Name: unifiedgosdk.Pointer("Visa 7844"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://kertzmanngroup.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("23129 Rebekah Mews"),
                City: unifiedgosdk.Pointer("Mortimerhaven"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("55878-6806"),
                Region: unifiedgosdk.Pointer("Florida"),
                RegionCode: unifiedgosdk.Pointer("OK"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("867556410"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(432) 760-1716"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2023-12-25T17:42:49.991Z"),
            Website: unifiedgosdk.Pointer("https://kertzmanngroup.com"),
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

    res, err := s.Contact.UpdateCrmContact(ctx, operations.UpdateCrmContactRequest{
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

## UpdateUcContact

Updates the UcContact identified by `id`. Provide the unified UcContact fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("9382c547-dae9-4d0d-b495-8ae21f1a5635"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
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
            UpdatedAt: types.MustNewTimeFromString("2023-12-16T00:29:55.950Z"),
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