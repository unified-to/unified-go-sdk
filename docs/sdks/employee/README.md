# Employee

### Available Operations

* [DeleteHrisConnectionIDEmployeeID](#deletehrisconnectionidemployeeid) - Remove a Employee
* [GetHrisConnectionIDEmployee](#gethrisconnectionidemployee) - List all Employees
* [GetHrisConnectionIDEmployeeID](#gethrisconnectionidemployeeid) - Retrieve a Employee
* [PatchHrisConnectionIDEmployeeID](#patchhrisconnectionidemployeeid) - Update a Employee
* [PostHrisConnectionIDEmployee](#posthrisconnectionidemployee) - Create a Employee
* [PutHrisConnectionIDEmployeeID](#puthrisconnectionidemployeeid) - Update a Employee

## DeleteHrisConnectionIDEmployeeID

Remove a Employee

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
    operationSecurity := operations.DeleteHrisConnectionIDEmployeeIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.DeleteHrisConnectionIDEmployeeID(ctx, operations.DeleteHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "delectus",
        ID: "4844225e-75b7-4960-a5c0-efa6f93b90a1",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteHrisConnectionIDEmployeeIDRequest](../../models/operations/deletehrisconnectionidemployeeidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeleteHrisConnectionIDEmployeeIDSecurity](../../models/operations/deletehrisconnectionidemployeeidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.DeleteHrisConnectionIDEmployeeIDResponse](../../models/operations/deletehrisconnectionidemployeeidresponse.md), error**


## GetHrisConnectionIDEmployee

List all Employees

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
    operationSecurity := operations.GetHrisConnectionIDEmployeeSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.GetHrisConnectionIDEmployee(ctx, operations.GetHrisConnectionIDEmployeeRequest{
        ConnectionID: "distinctio",
        Limit: unifiedto.Float64(5362.02),
        Offset: unifiedto.Float64(7558.48),
        Order: unifiedto.String("unde"),
        Query: unifiedto.String("veniam"),
        Sort: unifiedto.String("nam"),
        UpdatedGte: types.MustTimeFromString("2022-08-28T08:17:40.334Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployees != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetHrisConnectionIDEmployeeRequest](../../models/operations/gethrisconnectionidemployeerequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetHrisConnectionIDEmployeeSecurity](../../models/operations/gethrisconnectionidemployeesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.GetHrisConnectionIDEmployeeResponse](../../models/operations/gethrisconnectionidemployeeresponse.md), error**


## GetHrisConnectionIDEmployeeID

Retrieve a Employee

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
    operationSecurity := operations.GetHrisConnectionIDEmployeeIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.GetHrisConnectionIDEmployeeID(ctx, operations.GetHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "explicabo",
        ID: "54b739f4-fe77-4210-91f6-558c99c722d2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetHrisConnectionIDEmployeeIDRequest](../../models/operations/gethrisconnectionidemployeeidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetHrisConnectionIDEmployeeIDSecurity](../../models/operations/gethrisconnectionidemployeeidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.GetHrisConnectionIDEmployeeIDResponse](../../models/operations/gethrisconnectionidemployeeidresponse.md), error**


## PatchHrisConnectionIDEmployeeID

Update a Employee

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
    operationSecurity := operations.PatchHrisConnectionIDEmployeeIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.PatchHrisConnectionIDEmployeeID(ctx, operations.PatchHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("nam"),
                Address2: unifiedto.String("optio"),
                City: unifiedto.String("Fort Kentonburgh"),
                Country: unifiedto.String("Anguilla"),
                CountryCode: unifiedto.String("LS"),
                PostalCode: unifiedto.String("85866"),
                Region: unifiedto.String("recusandae"),
                RegionCode: unifiedto.String("doloremque"),
            },
            CreatedAt: types.MustTimeFromString("2022-11-04T08:22:44.750Z"),
            DateOfBirth: types.MustTimeFromString("2020-06-24T16:30:44.804Z"),
            Department: unifiedto.String("voluptate"),
            Division: unifiedto.String("placeat"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Manuela_Schowalter@hotmail.com",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("minus"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeCasual.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HiredAt: types.MustTimeFromString("2020-02-21T17:23:34.909Z"),
            ID: unifiedto.String("e9e15df9-0390-47f3-b831-983d42e54a85"),
            Location: unifiedto.String("dolore"),
            ManagerID: unifiedto.String("commodi"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedto.String("Ramona Kub"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "consequatur",
                    Type: shared.HrisTelephoneTypeWork.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2022-10-20T10:23:58.834Z"),
            Title: unifiedto.String("Dr."),
            UpdatedAt: types.MustTimeFromString("2022-09-26T09:29:33.100Z"),
        },
        ConnectionID: "nihil",
        ID: "1d51aaa6-ddf5-4abd-a487-c5fc2b862a00",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchHrisConnectionIDEmployeeIDRequest](../../models/operations/patchhrisconnectionidemployeeidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PatchHrisConnectionIDEmployeeIDSecurity](../../models/operations/patchhrisconnectionidemployeeidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.PatchHrisConnectionIDEmployeeIDResponse](../../models/operations/patchhrisconnectionidemployeeidresponse.md), error**


## PostHrisConnectionIDEmployee

Create a Employee

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
    operationSecurity := operations.PostHrisConnectionIDEmployeeSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.PostHrisConnectionIDEmployee(ctx, operations.PostHrisConnectionIDEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("nobis"),
                Address2: unifiedto.String("saepe"),
                City: unifiedto.String("Irvine"),
                Country: unifiedto.String("Montserrat"),
                CountryCode: unifiedto.String("TT"),
                PostalCode: unifiedto.String("00034"),
                Region: unifiedto.String("suscipit"),
                RegionCode: unifiedto.String("dolor"),
            },
            CreatedAt: types.MustTimeFromString("2022-04-09T11:38:39.555Z"),
            DateOfBirth: types.MustTimeFromString("2020-12-14T05:42:01.818Z"),
            Department: unifiedto.String("nihil"),
            Division: unifiedto.String("similique"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Roman_Torp27@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("dolor"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeCasual.ToPointer(),
            Gender: shared.HrisEmployeeGenderFemale.ToPointer(),
            HiredAt: types.MustTimeFromString("2022-11-14T11:24:15.321Z"),
            ID: unifiedto.String("38e1a735-ac26-4ae3-bbef-971a8f46bca1"),
            Location: unifiedto.String("quae"),
            ManagerID: unifiedto.String("aut"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedto.String("Frankie Mohr"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "soluta",
                    Type: shared.HrisTelephoneTypeOther.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2022-11-23T11:50:41.638Z"),
            Title: unifiedto.String("Dr."),
            UpdatedAt: types.MustTimeFromString("2022-07-01T14:41:51.539Z"),
        },
        ConnectionID: "eligendi",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostHrisConnectionIDEmployeeRequest](../../models/operations/posthrisconnectionidemployeerequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PostHrisConnectionIDEmployeeSecurity](../../models/operations/posthrisconnectionidemployeesecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PostHrisConnectionIDEmployeeResponse](../../models/operations/posthrisconnectionidemployeeresponse.md), error**


## PutHrisConnectionIDEmployeeID

Update a Employee

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
    operationSecurity := operations.PutHrisConnectionIDEmployeeIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Employee.PutHrisConnectionIDEmployeeID(ctx, operations.PutHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("tenetur"),
                Address2: unifiedto.String("deleniti"),
                City: unifiedto.String("Stantonfort"),
                Country: unifiedto.String("Namibia"),
                CountryCode: unifiedto.String("WF"),
                PostalCode: unifiedto.String("75563"),
                Region: unifiedto.String("quis"),
                RegionCode: unifiedto.String("doloremque"),
            },
            CreatedAt: types.MustTimeFromString("2022-02-19T05:47:56.169Z"),
            DateOfBirth: types.MustTimeFromString("2022-08-03T10:40:15.780Z"),
            Department: unifiedto.String("eveniet"),
            Division: unifiedto.String("possimus"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Claire_Frami3@gmail.com",
                    Type: shared.HrisEmailTypeOther.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("officiis"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeCasual.ToPointer(),
            Gender: shared.HrisEmployeeGenderTrans.ToPointer(),
            HiredAt: types.MustTimeFromString("2022-08-31T15:51:12.518Z"),
            ID: unifiedto.String("432a986e-b7e1-44ca-9640-89150097019a"),
            Location: unifiedto.String("eius"),
            ManagerID: unifiedto.String("rem"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusSingle.ToPointer(),
            Name: unifiedto.String("Armando Wehner"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "reprehenderit",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2021-04-08T14:13:47.569Z"),
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustTimeFromString("2022-02-05T17:52:52.343Z"),
        },
        ConnectionID: "consequatur",
        ID: "1105d389-0816-42c6-beb6-8a0f657b7d03",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutHrisConnectionIDEmployeeIDRequest](../../models/operations/puthrisconnectionidemployeeidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PutHrisConnectionIDEmployeeIDSecurity](../../models/operations/puthrisconnectionidemployeeidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.PutHrisConnectionIDEmployeeIDResponse](../../models/operations/puthrisconnectionidemployeeidresponse.md), error**

