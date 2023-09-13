# Contact

### Available Operations

* [DeleteCrmConnectionIDContactID](#deletecrmconnectionidcontactid) - Remove a contact
* [DeleteCrmConnectionIDContactIDCompanyCompanyID](#deletecrmconnectionidcontactidcompanycompanyid) - Remove company association from a contact
* [DeleteCrmConnectionIDContactIDDealDealID](#deletecrmconnectionidcontactiddealdealid) - Remove deal association from a contact
* [DeleteUcConnectionIDContactID](#deleteucconnectionidcontactid) - Remove a contact
* [GetCrmConnectionIDContact](#getcrmconnectionidcontact) - List all contacts
* [GetCrmConnectionIDContactID](#getcrmconnectionidcontactid) - Retrieve a contact
* [GetUcConnectionIDContact](#getucconnectionidcontact) - List all contacts
* [GetUcConnectionIDContactID](#getucconnectionidcontactid) - Retrieve a contact
* [PatchCrmConnectionIDContactID](#patchcrmconnectionidcontactid) - Update a contact
* [PatchCrmConnectionIDContactIDCompanyCompanyID](#patchcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PatchCrmConnectionIDContactIDDealDealID](#patchcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PatchUcConnectionIDContactID](#patchucconnectionidcontactid) - Update a contact
* [PostCrmConnectionIDContact](#postcrmconnectionidcontact) - Create a contact
* [PostUcConnectionIDContact](#postucconnectionidcontact) - Create a contact
* [PutCrmConnectionIDContactID](#putcrmconnectionidcontactid) - Update a contact
* [PutCrmConnectionIDContactIDCompanyCompanyID](#putcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PutCrmConnectionIDContactIDDealDealID](#putcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PutUcConnectionIDContactID](#putucconnectionidcontactid) - Update a contact

## DeleteCrmConnectionIDContactID

Remove a contact

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
    operationSecurity := operations.DeleteCrmConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.DeleteCrmConnectionIDContactID(ctx, operations.DeleteCrmConnectionIDContactIDRequest{
        ConnectionID: "molestias",
        ID: "57a9e618-76c6-4ab2-9d29-dfc94d6fecd7",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteCrmConnectionIDContactIDRequest](../../models/operations/deletecrmconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.DeleteCrmConnectionIDContactIDSecurity](../../models/operations/deletecrmconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.DeleteCrmConnectionIDContactIDResponse](../../models/operations/deletecrmconnectionidcontactidresponse.md), error**


## DeleteCrmConnectionIDContactIDCompanyCompanyID

Remove company association from a contact

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
    operationSecurity := operations.DeleteCrmConnectionIDContactIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.DeleteCrmConnectionIDContactIDCompanyCompanyID(ctx, operations.DeleteCrmConnectionIDContactIDCompanyCompanyIDRequest{
        CompanyID: "iste",
        ConnectionID: "provident",
        ID: "390066a6-d2d0-4003-9533-8cec086fa21e",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.DeleteCrmConnectionIDContactIDCompanyCompanyIDRequest](../../models/operations/deletecrmconnectionidcontactidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `security`                                                                                                                                             | [operations.DeleteCrmConnectionIDContactIDCompanyCompanyIDSecurity](../../models/operations/deletecrmconnectionidcontactidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                                     | The security requirements to use for the request.                                                                                                      |


### Response

**[*operations.DeleteCrmConnectionIDContactIDCompanyCompanyIDResponse](../../models/operations/deletecrmconnectionidcontactidcompanycompanyidresponse.md), error**


## DeleteCrmConnectionIDContactIDDealDealID

Remove deal association from a contact

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
    operationSecurity := operations.DeleteCrmConnectionIDContactIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.DeleteCrmConnectionIDContactIDDealDealID(ctx, operations.DeleteCrmConnectionIDContactIDDealDealIDRequest{
        ConnectionID: "iste",
        DealID: "dicta",
        ID: "52cb3119-167b-48e3-88db-03408d6d364f",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.DeleteCrmConnectionIDContactIDDealDealIDRequest](../../models/operations/deletecrmconnectionidcontactiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `security`                                                                                                                                 | [operations.DeleteCrmConnectionIDContactIDDealDealIDSecurity](../../models/operations/deletecrmconnectionidcontactiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                         | The security requirements to use for the request.                                                                                          |


### Response

**[*operations.DeleteCrmConnectionIDContactIDDealDealIDResponse](../../models/operations/deletecrmconnectionidcontactiddealdealidresponse.md), error**


## DeleteUcConnectionIDContactID

Remove a contact

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
    operationSecurity := operations.DeleteUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.DeleteUcConnectionIDContactID(ctx, operations.DeleteUcConnectionIDContactIDRequest{
        ConnectionID: "tenetur",
        ID: "d455906d-1263-4d48-a935-c2c9e81f30be",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteUcConnectionIDContactIDRequest](../../models/operations/deleteucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.DeleteUcConnectionIDContactIDSecurity](../../models/operations/deleteucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.DeleteUcConnectionIDContactIDResponse](../../models/operations/deleteucconnectionidcontactidresponse.md), error**


## GetCrmConnectionIDContact

List all contacts

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
    operationSecurity := operations.GetCrmConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.GetCrmConnectionIDContact(ctx, operations.GetCrmConnectionIDContactRequest{
        CompanyID: unifiedto.String("sequi"),
        ConnectionID: "recusandae",
        DealID: unifiedto.String("labore"),
        Limit: unifiedto.Float64(2406.24),
        Offset: unifiedto.Float64(1667.41),
        Order: unifiedto.String("aperiam"),
        Query: unifiedto.String("dolores"),
        Sort: unifiedto.String("illum"),
        UpdatedGte: types.MustDateFromString("2022-10-31"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCrmConnectionIDContactRequest](../../models/operations/getcrmconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.GetCrmConnectionIDContactSecurity](../../models/operations/getcrmconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.GetCrmConnectionIDContactResponse](../../models/operations/getcrmconnectionidcontactresponse.md), error**


## GetCrmConnectionIDContactID

Retrieve a contact

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
    operationSecurity := operations.GetCrmConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.GetCrmConnectionIDContactID(ctx, operations.GetCrmConnectionIDContactIDRequest{
        ConnectionID: "beatae",
        ID: "65765066-4187-40d9-921f-9ad030c4ecc1",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCrmConnectionIDContactIDRequest](../../models/operations/getcrmconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetCrmConnectionIDContactIDSecurity](../../models/operations/getcrmconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.GetCrmConnectionIDContactIDResponse](../../models/operations/getcrmconnectionidcontactidresponse.md), error**


## GetUcConnectionIDContact

List all contacts

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
    operationSecurity := operations.GetUcConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.GetUcConnectionIDContact(ctx, operations.GetUcConnectionIDContactRequest{
        AgentID: unifiedto.String("beatae"),
        ConnectionID: "id",
        Limit: unifiedto.Float64(90.6),
        Offset: unifiedto.Float64(5515.76),
        Order: unifiedto.String("ratione"),
        Query: unifiedto.String("iure"),
        Sort: unifiedto.String("tempora"),
        UpdatedGte: types.MustDateFromString("2022-05-21"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetUcConnectionIDContactRequest](../../models/operations/getucconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetUcConnectionIDContactSecurity](../../models/operations/getucconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.GetUcConnectionIDContactResponse](../../models/operations/getucconnectionidcontactresponse.md), error**


## GetUcConnectionIDContactID

Retrieve a contact

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
    operationSecurity := operations.GetUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.GetUcConnectionIDContactID(ctx, operations.GetUcConnectionIDContactIDRequest{
        ConnectionID: "voluptatem",
        ID: "68b8502a-55e7-4f73-bc84-5e320a319f4b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetUcConnectionIDContactIDRequest](../../models/operations/getucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetUcConnectionIDContactIDSecurity](../../models/operations/getucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.GetUcConnectionIDContactIDResponse](../../models/operations/getucconnectionidcontactidresponse.md), error**


## PatchCrmConnectionIDContactID

Update a contact

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
    operationSecurity := operations.PatchCrmConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PatchCrmConnectionIDContactID(ctx, operations.PatchCrmConnectionIDContactIDRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedto.String("dolorum"),
                Address2: unifiedto.String("quibusdam"),
                City: unifiedto.String("Mission"),
                Country: unifiedto.String("El Salvador"),
                CountryCode: unifiedto.String("KM"),
                PostalCode: unifiedto.String("66534-7721"),
                Region: unifiedto.String("magnam"),
                RegionCode: unifiedto.String("dolores"),
            },
            Company: unifiedto.String("Kemmer - Kassulke"),
            CompanyIds: []string{
                "ad",
            },
            CreatedAt: types.MustDateFromString("2022-11-07"),
            DealIds: []string{
                "suscipit",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Rupert_Russel@hotmail.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("f51fcb4c-593e-4c12-8daa-d0ec7afedbd8"),
            Name: unifiedto.String("Paulette White"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "corrupti",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedto.String("Mrs."),
            UpdatedAt: types.MustDateFromString("2022-01-12"),
        },
        ConnectionID: "iste",
        ID: "390c5888-0983-4dab-b9ef-3ffdd9f7f079",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchCrmConnectionIDContactIDRequest](../../models/operations/patchcrmconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PatchCrmConnectionIDContactIDSecurity](../../models/operations/patchcrmconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.PatchCrmConnectionIDContactIDResponse](../../models/operations/patchcrmconnectionidcontactidresponse.md), error**


## PatchCrmConnectionIDContactIDCompanyCompanyID

Associate a company with a contact

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
    operationSecurity := operations.PatchCrmConnectionIDContactIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PatchCrmConnectionIDContactIDCompanyCompanyID(ctx, operations.PatchCrmConnectionIDContactIDCompanyCompanyIDRequest{
        CompanyID: "similique",
        ConnectionID: "asperiores",
        ID: "4d35724c-db0f-44d2-8118-7d56844eded8",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PatchCrmConnectionIDContactIDCompanyCompanyIDRequest](../../models/operations/patchcrmconnectionidcontactidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `security`                                                                                                                                           | [operations.PatchCrmConnectionIDContactIDCompanyCompanyIDSecurity](../../models/operations/patchcrmconnectionidcontactidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                                   | The security requirements to use for the request.                                                                                                    |


### Response

**[*operations.PatchCrmConnectionIDContactIDCompanyCompanyIDResponse](../../models/operations/patchcrmconnectionidcontactidcompanycompanyidresponse.md), error**


## PatchCrmConnectionIDContactIDDealDealID

Associate a deal with a contact

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
    operationSecurity := operations.PatchCrmConnectionIDContactIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PatchCrmConnectionIDContactIDDealDealID(ctx, operations.PatchCrmConnectionIDContactIDDealDealIDRequest{
        ConnectionID: "enim",
        DealID: "animi",
        ID: "9065e628-bdfc-4203-ab6c-879923b7e135",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PatchCrmConnectionIDContactIDDealDealIDRequest](../../models/operations/patchcrmconnectionidcontactiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.PatchCrmConnectionIDContactIDDealDealIDSecurity](../../models/operations/patchcrmconnectionidcontactiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |


### Response

**[*operations.PatchCrmConnectionIDContactIDDealDealIDResponse](../../models/operations/patchcrmconnectionidcontactiddealdealidresponse.md), error**


## PatchUcConnectionIDContactID

Update a contact

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
    operationSecurity := operations.PatchUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PatchUcConnectionIDContactID(ctx, operations.PatchUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Graham - Zboncak"),
            CreatedAt: types.MustDateFromString("2022-05-10"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Augustine.Davis50@hotmail.com",
                    Type: shared.UcEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("1f82ce11-5717-4230-9377-dcfa89df975e"),
            Name: unifiedto.String("Erica Jerde"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "aliquid",
                    Type: shared.UcTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedto.String("Ms."),
            UpdatedAt: types.MustDateFromString("2022-01-25"),
        },
        ConnectionID: "unde",
        ID: "c3ddc5f1-11de-4a10-a6d5-41a4d190feb2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchUcConnectionIDContactIDRequest](../../models/operations/patchucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PatchUcConnectionIDContactIDSecurity](../../models/operations/patchucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PatchUcConnectionIDContactIDResponse](../../models/operations/patchucconnectionidcontactidresponse.md), error**


## PostCrmConnectionIDContact

Create a contact

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
    operationSecurity := operations.PostCrmConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PostCrmConnectionIDContact(ctx, operations.PostCrmConnectionIDContactRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedto.String("vitae"),
                Address2: unifiedto.String("odio"),
                City: unifiedto.String("Amyacester"),
                Country: unifiedto.String("Slovenia"),
                CountryCode: unifiedto.String("SI"),
                PostalCode: unifiedto.String("08778-8725"),
                Region: unifiedto.String("tempora"),
                RegionCode: unifiedto.String("esse"),
            },
            Company: unifiedto.String("Lowe Group"),
            CompanyIds: []string{
                "facilis",
            },
            CreatedAt: types.MustDateFromString("2022-01-29"),
            DealIds: []string{
                "amet",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Annetta_Torphy@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("158c4c4e-5459-49ea-b422-60e9b200ce78"),
            Name: unifiedto.String("Carl Rath"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "maiores",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedto.String("Ms."),
            UpdatedAt: types.MustDateFromString("2022-11-30"),
        },
        ConnectionID: "fuga",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostCrmConnectionIDContactRequest](../../models/operations/postcrmconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PostCrmConnectionIDContactSecurity](../../models/operations/postcrmconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PostCrmConnectionIDContactResponse](../../models/operations/postcrmconnectionidcontactresponse.md), error**


## PostUcConnectionIDContact

Create a contact

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
    operationSecurity := operations.PostUcConnectionIDContactSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PostUcConnectionIDContact(ctx, operations.PostUcConnectionIDContactRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Braun and Sons"),
            CreatedAt: types.MustDateFromString("2020-05-06"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Caleb.Dickens1@yahoo.com",
                    Type: shared.UcEmailTypeHome.ToPointer(),
                },
            },
            ID: unifiedto.String("7fa30e9a-f725-4b29-9220-30d83f5aeb77"),
            Name: unifiedto.String("Mr. Tracy Shields"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "deleniti",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
            },
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustDateFromString("2021-05-16"),
        },
        ConnectionID: "magnam",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostUcConnectionIDContactRequest](../../models/operations/postucconnectionidcontactrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PostUcConnectionIDContactSecurity](../../models/operations/postucconnectionidcontactsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.PostUcConnectionIDContactResponse](../../models/operations/postucconnectionidcontactresponse.md), error**


## PutCrmConnectionIDContactID

Update a contact

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
    operationSecurity := operations.PutCrmConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PutCrmConnectionIDContactID(ctx, operations.PutCrmConnectionIDContactIDRequest{
        CrmContact: &shared.CrmContact{
            Address: &shared.PropertyCrmContactAddress{
                Address1: unifiedto.String("perspiciatis"),
                Address2: unifiedto.String("amet"),
                City: unifiedto.String("Braedenbury"),
                Country: unifiedto.String("Venezuela"),
                CountryCode: unifiedto.String("SY"),
                PostalCode: unifiedto.String("31754-4817"),
                Region: unifiedto.String("consequuntur"),
                RegionCode: unifiedto.String("possimus"),
            },
            Company: unifiedto.String("Pollich, Haag and Rosenbaum"),
            CompanyIds: []string{
                "hic",
            },
            CreatedAt: types.MustDateFromString("2022-10-10"),
            DealIds: []string{
                "nobis",
            },
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedto.String("Gladyce0@yahoo.com"),
                    Type: shared.CrmEmailTypeOther.ToPointer(),
                },
            },
            ID: unifiedto.String("841fb1bd-23fd-4b14-9b6b-e5a685998e22"),
            Name: unifiedto.String("Dr. Hugo Daugherty"),
            Raw: &shared.PropertyCrmContactRaw{},
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "quae",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedto.String("Dr."),
            UpdatedAt: types.MustDateFromString("2022-07-21"),
        },
        ConnectionID: "nam",
        ID: "271a289c-57e8-454e-9043-9d2224656946",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutCrmConnectionIDContactIDRequest](../../models/operations/putcrmconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PutCrmConnectionIDContactIDSecurity](../../models/operations/putcrmconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.PutCrmConnectionIDContactIDResponse](../../models/operations/putcrmconnectionidcontactidresponse.md), error**


## PutCrmConnectionIDContactIDCompanyCompanyID

Associate a company with a contact

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
    operationSecurity := operations.PutCrmConnectionIDContactIDCompanyCompanyIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PutCrmConnectionIDContactIDCompanyCompanyID(ctx, operations.PutCrmConnectionIDContactIDCompanyCompanyIDRequest{
        CompanyID: "explicabo",
        ConnectionID: "modi",
        ID: "07084f7a-b37c-4ef0-a225-194db55410ad",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.PutCrmConnectionIDContactIDCompanyCompanyIDRequest](../../models/operations/putcrmconnectionidcontactidcompanycompanyidrequest.md)   | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `security`                                                                                                                                       | [operations.PutCrmConnectionIDContactIDCompanyCompanyIDSecurity](../../models/operations/putcrmconnectionidcontactidcompanycompanyidsecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |


### Response

**[*operations.PutCrmConnectionIDContactIDCompanyCompanyIDResponse](../../models/operations/putcrmconnectionidcontactidcompanycompanyidresponse.md), error**


## PutCrmConnectionIDContactIDDealDealID

Associate a deal with a contact

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
    operationSecurity := operations.PutCrmConnectionIDContactIDDealDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PutCrmConnectionIDContactIDDealDealID(ctx, operations.PutCrmConnectionIDContactIDDealDealIDRequest{
        ConnectionID: "quo",
        DealID: "suscipit",
        ID: "69af90a2-6c7c-4dc9-81f0-68981d6bb33c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PutCrmConnectionIDContactIDDealDealIDRequest](../../models/operations/putcrmconnectionidcontactiddealdealidrequest.md)   | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `security`                                                                                                                           | [operations.PutCrmConnectionIDContactIDDealDealIDSecurity](../../models/operations/putcrmconnectionidcontactiddealdealidsecurity.md) | :heavy_check_mark:                                                                                                                   | The security requirements to use for the request.                                                                                    |


### Response

**[*operations.PutCrmConnectionIDContactIDDealDealIDResponse](../../models/operations/putcrmconnectionidcontactiddealdealidresponse.md), error**


## PutUcConnectionIDContactID

Update a contact

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
    operationSecurity := operations.PutUcConnectionIDContactIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Contact.PutUcConnectionIDContactID(ctx, operations.PutUcConnectionIDContactIDRequest{
        UcContact: &shared.UcContact{
            Company: unifiedto.String("Murphy, Pfannerstill and Feest"),
            CreatedAt: types.MustDateFromString("2022-06-26"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Dax_Boehm30@hotmail.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            ID: unifiedto.String("7ee4fcf0-c42b-478f-9562-6398a0dc7663"),
            Name: unifiedto.String("Jamie Schneider"),
            Raw: &shared.PropertyUcContactRaw{},
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "accusantium",
                    Type: shared.UcTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedto.String("Miss"),
            UpdatedAt: types.MustDateFromString("2021-06-23"),
        },
        ConnectionID: "est",
        ID: "12d02529-270b-48d5-b22d-d895b8bcf24d",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutUcConnectionIDContactIDRequest](../../models/operations/putucconnectionidcontactidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PutUcConnectionIDContactIDSecurity](../../models/operations/putucconnectionidcontactidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PutUcConnectionIDContactIDResponse](../../models/operations/putucconnectionidcontactidresponse.md), error**

