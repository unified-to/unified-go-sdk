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
                Address1: unifiedgosdk.Pointer("9171 Parisian Gateway"),
                City: unifiedgosdk.Pointer("Lake Keltonboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("CO"),
            },
            CompanyName: unifiedgosdk.Pointer("Rice LLC"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("michele.breitenberg@ricellc.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@ricellc.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Michele"),
            ID: unifiedgosdk.Pointer("6ffdfe98-c45a-42be-8eac-998457211a68"),
            Identification: unifiedgosdk.Pointer("508720595"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Breitenberg"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rice LLC"),
            Notes: unifiedgosdk.Pointer("Tabella reprehenderit sono corrupti approbo crastinus sopor ultra."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("43c7beab-abc7-4d66-ad33-b212013439d3"),
                    Name: unifiedgosdk.Pointer("Visa 5318"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://ricellc.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Dare Meadows"),
                City: unifiedgosdk.Pointer("Phoenix"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("68675"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("564107844"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(217) 260-4409"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-01-05T03:13:26.378Z"),
            Website: unifiedgosdk.Pointer("https://ricellc.com"),
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
                Address1: unifiedgosdk.Pointer("489 Eula Park"),
                Address2: unifiedgosdk.Pointer("Apt. 900"),
                City: unifiedgosdk.Pointer("Archfield"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("60941-8210"),
                Region: unifiedgosdk.Pointer("Texas"),
                RegionCode: unifiedgosdk.Pointer("NH"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Health"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("alexandro.kuphal@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alexandro"),
            ID: unifiedgosdk.Pointer("0cb9e92a-f41d-4c8d-8e65-b3925244f5ff"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26594214"),
            LastName: unifiedgosdk.Pointer("Kuphal"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/alexandro-kuphal",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("37cef42c-e539-4916-b606-cdacb8c627e8"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "amplus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Alexandro Kuphal"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(404) 452-9269",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(989) 758-8616",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Infrastructure Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-25T02:46:23.329Z"),
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
            Company: unifiedgosdk.Pointer("Olson - Stanton"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Norval_Brekke@gmail.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Norval"),
            ID: unifiedgosdk.Pointer("2c39d126-e669-4218-8db5-e84c7d941baf"),
            LastName: unifiedgosdk.Pointer("Brekke"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Norval Brekke"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(992) 368-1868",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(238) 374-6019",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(736) 235-7480",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Internal Creative Director"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-29T13:20:51.045Z"),
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
                Address1: unifiedgosdk.Pointer("9171 Parisian Gateway"),
                City: unifiedgosdk.Pointer("Lake Keltonboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("CO"),
            },
            CompanyName: unifiedgosdk.Pointer("Rice LLC"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("michele.breitenberg@ricellc.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@ricellc.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Michele"),
            ID: unifiedgosdk.Pointer("b3cd5556-0414-4043-a19a-55642194b29c"),
            Identification: unifiedgosdk.Pointer("508720595"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Breitenberg"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rice LLC"),
            Notes: unifiedgosdk.Pointer("Tabella reprehenderit sono corrupti approbo crastinus sopor ultra."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("e6fbdbad-13fd-424a-9552-2928b73efd2c"),
                    Name: unifiedgosdk.Pointer("Visa 5318"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://ricellc.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Dare Meadows"),
                City: unifiedgosdk.Pointer("Phoenix"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("68675"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("564107844"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(217) 260-4409"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-01-05T03:13:26.393Z"),
            Website: unifiedgosdk.Pointer("https://ricellc.com"),
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
                Address1: unifiedgosdk.Pointer("489 Eula Park"),
                Address2: unifiedgosdk.Pointer("Apt. 900"),
                City: unifiedgosdk.Pointer("Archfield"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("60941-8210"),
                Region: unifiedgosdk.Pointer("Texas"),
                RegionCode: unifiedgosdk.Pointer("NH"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Health"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("alexandro.kuphal@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alexandro"),
            ID: unifiedgosdk.Pointer("0f0cb0d9-601f-4bb8-808f-e234cc0fd51e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26594214"),
            LastName: unifiedgosdk.Pointer("Kuphal"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/alexandro-kuphal",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("1feda212-b228-449d-bcf3-e4a0b60b1689"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "amplus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Alexandro Kuphal"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(404) 452-9269",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(989) 758-8616",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Infrastructure Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-25T02:46:23.330Z"),
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
            Company: unifiedgosdk.Pointer("Olson - Stanton"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Norval_Brekke@gmail.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Norval"),
            ID: unifiedgosdk.Pointer("ee028c4e-42f3-448f-b77e-fc5a6aa90ee3"),
            LastName: unifiedgosdk.Pointer("Brekke"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Norval Brekke"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(992) 368-1868",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(238) 374-6019",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(736) 235-7480",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Internal Creative Director"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-29T13:20:51.053Z"),
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
                Address1: unifiedgosdk.Pointer("9171 Parisian Gateway"),
                City: unifiedgosdk.Pointer("Lake Keltonboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("70131-2908"),
                Region: unifiedgosdk.Pointer("Illinois"),
                RegionCode: unifiedgosdk.Pointer("CO"),
            },
            CompanyName: unifiedgosdk.Pointer("Rice LLC"),
            CreatedAt: types.MustNewTimeFromString("2021-08-15T14:56:50.258Z"),
            Currency: unifiedgosdk.Pointer("AUD"),
            Emails: []shared.AccountingEmail{
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("michele.breitenberg@ricellc.com"),
                    Type: shared.AccountingEmailTypeWork.ToPointer(),
                },
                shared.AccountingEmail{
                    Email: unifiedgosdk.Pointer("accounts@ricellc.com"),
                    Type: shared.AccountingEmailTypeOther.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Michele"),
            ID: unifiedgosdk.Pointer("b3cd5556-0414-4043-a19a-55642194b29c"),
            Identification: unifiedgosdk.Pointer("508720595"),
            IsActive: unifiedgosdk.Pointer(true),
            IsCustomer: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Breitenberg"),
            Metadata: []shared.AccountingMetadata{},
            Name: unifiedgosdk.Pointer("Rice LLC"),
            Notes: unifiedgosdk.Pointer("Tabella reprehenderit sono corrupti approbo crastinus sopor ultra."),
            PaymentMethods: []shared.AccountingContactPaymentMethod{
                shared.AccountingContactPaymentMethod{
                    Default: unifiedgosdk.Pointer(true),
                    ID: unifiedgosdk.Pointer("e6fbdbad-13fd-424a-9552-2928b73efd2c"),
                    Name: unifiedgosdk.Pointer("Visa 5318"),
                    Type: shared.AccountingContactPaymentMethodTypeCard,
                },
            },
            PortalURL: unifiedgosdk.Pointer("https://ricellc.com/billing"),
            ShippingAddress: &shared.PropertyAccountingContactShippingAddress{
                Address1: unifiedgosdk.Pointer("786 Dare Meadows"),
                City: unifiedgosdk.Pointer("Phoenix"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("68675"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("NE"),
            },
            TaxExemption: shared.TaxExemptionForeign.ToPointer(),
            TaxNumber: unifiedgosdk.Pointer("564107844"),
            Telephones: []shared.AccountingTelephone{
                shared.AccountingTelephone{
                    Telephone: unifiedgosdk.Pointer("(217) 260-4409"),
                    Type: shared.AccountingTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2024-01-05T03:13:26.393Z"),
            Website: unifiedgosdk.Pointer("https://ricellc.com"),
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
                Address1: unifiedgosdk.Pointer("489 Eula Park"),
                Address2: unifiedgosdk.Pointer("Apt. 900"),
                City: unifiedgosdk.Pointer("Archfield"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("60941-8210"),
                Region: unifiedgosdk.Pointer("Texas"),
                RegionCode: unifiedgosdk.Pointer("NH"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-01-02T00:41:38.885Z"),
            Department: unifiedgosdk.Pointer("Health"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("alexandro.kuphal@hotmail.com"),
                    Type: shared.CrmEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Alexandro"),
            ID: unifiedgosdk.Pointer("0f0cb0d9-601f-4bb8-808f-e234cc0fd51e"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/26594214"),
            LastName: unifiedgosdk.Pointer("Kuphal"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/alexandro-kuphal",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("1feda212-b228-449d-bcf3-e4a0b60b1689"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "amplus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Alexandro Kuphal"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(404) 452-9269",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(989) 758-8616",
                    Type: shared.CrmTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Direct Infrastructure Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2021-02-25T02:46:23.330Z"),
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
            Company: unifiedgosdk.Pointer("Olson - Stanton"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Norval_Brekke@gmail.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Norval"),
            ID: unifiedgosdk.Pointer("ee028c4e-42f3-448f-b77e-fc5a6aa90ee3"),
            LastName: unifiedgosdk.Pointer("Brekke"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Norval Brekke"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(992) 368-1868",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(238) 374-6019",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(736) 235-7480",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Internal Creative Director"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-29T13:20:51.053Z"),
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