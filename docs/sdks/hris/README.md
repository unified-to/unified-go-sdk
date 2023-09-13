# Hris

### Available Operations

* [DeleteHrisConnectionIDEmployeeID](#deletehrisconnectionidemployeeid) - Remove a Employee
* [DeleteHrisConnectionIDGroupID](#deletehrisconnectionidgroupid) - Remove a Group
* [GetHrisConnectionIDEmployee](#gethrisconnectionidemployee) - List all Employees
* [GetHrisConnectionIDEmployeeID](#gethrisconnectionidemployeeid) - Retrieve a Employee
* [GetHrisConnectionIDGroup](#gethrisconnectionidgroup) - List all Groups
* [GetHrisConnectionIDGroupID](#gethrisconnectionidgroupid) - Retrieve a Group
* [PatchHrisConnectionIDEmployeeID](#patchhrisconnectionidemployeeid) - Update a Employee
* [PatchHrisConnectionIDGroupID](#patchhrisconnectionidgroupid) - Update a Group
* [PostHrisConnectionIDEmployee](#posthrisconnectionidemployee) - Create a Employee
* [PostHrisConnectionIDGroup](#posthrisconnectionidgroup) - Create a Group
* [PutHrisConnectionIDEmployeeID](#puthrisconnectionidemployeeid) - Update a Employee
* [PutHrisConnectionIDGroupID](#puthrisconnectionidgroupid) - Update a Group

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
    res, err := s.Hris.DeleteHrisConnectionIDEmployeeID(ctx, operations.DeleteHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "excepturi",
        ID: "3e34316c-f55b-4431-b553-ccf1c204c4ad",
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


## DeleteHrisConnectionIDGroupID

Remove a Group

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
    operationSecurity := operations.DeleteHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.DeleteHrisConnectionIDGroupID(ctx, operations.DeleteHrisConnectionIDGroupIDRequest{
        ConnectionID: "quod",
        ID: "c9904c51-95b8-4648-8efa-78f1e2d3b901",
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
| `request`                                                                                                            | [operations.DeleteHrisConnectionIDGroupIDRequest](../../models/operations/deletehrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.DeleteHrisConnectionIDGroupIDSecurity](../../models/operations/deletehrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.DeleteHrisConnectionIDGroupIDResponse](../../models/operations/deletehrisconnectionidgroupidresponse.md), error**


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
    res, err := s.Hris.GetHrisConnectionIDEmployee(ctx, operations.GetHrisConnectionIDEmployeeRequest{
        ConnectionID: "saepe",
        Limit: unifiedto.Float64(289.94),
        Offset: unifiedto.Float64(5725.89),
        Order: unifiedto.String("corporis"),
        Query: unifiedto.String("explicabo"),
        Sort: unifiedto.String("distinctio"),
        UpdatedGte: types.MustDateFromString("2021-08-12"),
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
    res, err := s.Hris.GetHrisConnectionIDEmployeeID(ctx, operations.GetHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "tempora",
        ID: "cbb19f71-3d95-4a41-a9c1-387271e18ea9",
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


## GetHrisConnectionIDGroup

List all Groups

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
    operationSecurity := operations.GetHrisConnectionIDGroupSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDGroup(ctx, operations.GetHrisConnectionIDGroupRequest{
        ConnectionID: "debitis",
        Limit: unifiedto.Float64(2501.58),
        Offset: unifiedto.Float64(3333.54),
        Order: unifiedto.String("illo"),
        Query: unifiedto.String("illo"),
        Sort: unifiedto.String("deleniti"),
        UpdatedGte: types.MustDateFromString("2022-07-26"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetHrisConnectionIDGroupRequest](../../models/operations/gethrisconnectionidgrouprequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetHrisConnectionIDGroupSecurity](../../models/operations/gethrisconnectionidgroupsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.GetHrisConnectionIDGroupResponse](../../models/operations/gethrisconnectionidgroupresponse.md), error**


## GetHrisConnectionIDGroupID

Retrieve a Group

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
    operationSecurity := operations.GetHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDGroupID(ctx, operations.GetHrisConnectionIDGroupIDRequest{
        ConnectionID: "optio",
        ID: "c57fbd60-b1a7-48ed-a9a9-d4eea85658c2",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetHrisConnectionIDGroupIDRequest](../../models/operations/gethrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetHrisConnectionIDGroupIDSecurity](../../models/operations/gethrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.GetHrisConnectionIDGroupIDResponse](../../models/operations/gethrisconnectionidgroupidresponse.md), error**


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
    res, err := s.Hris.PatchHrisConnectionIDEmployeeID(ctx, operations.PatchHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("at"),
                Address2: unifiedto.String("modi"),
                City: unifiedto.String("Findlay"),
                Country: unifiedto.String("Saint Martin"),
                CountryCode: unifiedto.String("LI"),
                PostalCode: unifiedto.String("69291-4598"),
                Region: unifiedto.String("unde"),
                RegionCode: unifiedto.String("autem"),
            },
            CreatedAt: types.MustDateFromString("2022-07-11"),
            DateOfBirth: types.MustDateFromString("2022-02-18"),
            Department: unifiedto.String("autem"),
            Division: unifiedto.String("placeat"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Blaze97@yahoo.com",
                    Type: shared.HrisEmailTypeOther.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("id"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeIntern.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustDateFromString("2020-08-13"),
            ID: unifiedto.String("ef234c95-5b9b-4df2-990a-bd9bbcc2725e"),
            Location: unifiedto.String("impedit"),
            ManagerID: unifiedto.String("magni"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedto.String("Sabrina Schamberger Sr."),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "totam",
                    Type: shared.HrisTelephoneTypeWork.ToPointer(),
                },
            },
            TerminatedAt: types.MustDateFromString("2022-06-27"),
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustDateFromString("2021-10-27"),
        },
        ConnectionID: "excepturi",
        ID: "ef68e45c-8add-4fac-b545-00430c6632b4",
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


## PatchHrisConnectionIDGroupID

Update a Group

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
    operationSecurity := operations.PatchHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.PatchHrisConnectionIDGroupID(ctx, operations.PatchHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2022-05-31"),
            Description: unifiedto.String("inventore"),
            EmployeeIds: []string{
                "sapiente",
            },
            ID: unifiedto.String("df01c3e9-1e8f-47bc-a9d4-60a77eceb26d"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "architecto",
            },
            Name: unifiedto.String("Lorene Bosco"),
            ParentID: unifiedto.String("qui"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDepartment.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-12-06"),
        },
        ConnectionID: "quisquam",
        ID: "7c0f0f87-3f9d-45c2-9fd3-e0b4a4a4253c",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchHrisConnectionIDGroupIDRequest](../../models/operations/patchhrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PatchHrisConnectionIDGroupIDSecurity](../../models/operations/patchhrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PatchHrisConnectionIDGroupIDResponse](../../models/operations/patchhrisconnectionidgroupidresponse.md), error**


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
    res, err := s.Hris.PostHrisConnectionIDEmployee(ctx, operations.PostHrisConnectionIDEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("amet"),
                Address2: unifiedto.String("ipsa"),
                City: unifiedto.String("West Jaunitaland"),
                Country: unifiedto.String("Bahrain"),
                CountryCode: unifiedto.String("YE"),
                PostalCode: unifiedto.String("18494"),
                Region: unifiedto.String("pariatur"),
                RegionCode: unifiedto.String("porro"),
            },
            CreatedAt: types.MustDateFromString("2022-09-23"),
            DateOfBirth: types.MustDateFromString("2021-07-11"),
            Department: unifiedto.String("itaque"),
            Division: unifiedto.String("sit"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Timmothy68@yahoo.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("culpa"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeFullTime.ToPointer(),
            Gender: shared.HrisEmployeeGenderFemale.ToPointer(),
            HiredAt: types.MustDateFromString("2021-01-27"),
            ID: unifiedto.String("12a4ba9d-5998-4819-acfd-0c77c53e7e7d"),
            Location: unifiedto.String("eius"),
            ManagerID: unifiedto.String("accusamus"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusSingle.ToPointer(),
            Name: unifiedto.String("Sophie Lesch"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "consequatur",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminatedAt: types.MustDateFromString("2021-06-09"),
            Title: unifiedto.String("Mrs."),
            UpdatedAt: types.MustDateFromString("2022-06-19"),
        },
        ConnectionID: "saepe",
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


## PostHrisConnectionIDGroup

Create a Group

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
    operationSecurity := operations.PostHrisConnectionIDGroupSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.PostHrisConnectionIDGroup(ctx, operations.PostHrisConnectionIDGroupRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2022-10-11"),
            Description: unifiedto.String("sint"),
            EmployeeIds: []string{
                "ea",
            },
            ID: unifiedto.String("703fec31-c508-424d-989a-36a6b2d27eb7"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "accusantium",
            },
            Name: unifiedto.String("Harriet Orn DDS"),
            ParentID: unifiedto.String("voluptatum"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeSubDepartment.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-02-14"),
        },
        ConnectionID: "commodi",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostHrisConnectionIDGroupRequest](../../models/operations/posthrisconnectionidgrouprequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PostHrisConnectionIDGroupSecurity](../../models/operations/posthrisconnectionidgroupsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.PostHrisConnectionIDGroupResponse](../../models/operations/posthrisconnectionidgroupresponse.md), error**


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
    res, err := s.Hris.PutHrisConnectionIDEmployeeID(ctx, operations.PutHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedto.String("debitis"),
                Address2: unifiedto.String("commodi"),
                City: unifiedto.String("South Izaiahhaven"),
                Country: unifiedto.String("Paraguay"),
                CountryCode: unifiedto.String("MS"),
                PostalCode: unifiedto.String("72740-9977"),
                Region: unifiedto.String("nisi"),
                RegionCode: unifiedto.String("occaecati"),
            },
            CreatedAt: types.MustDateFromString("2022-12-15"),
            DateOfBirth: types.MustDateFromString("2020-03-31"),
            Department: unifiedto.String("odio"),
            Division: unifiedto.String("nihil"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Tamia.Doyle@gmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedto.String("iusto"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeOther.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustDateFromString("2022-06-18"),
            ID: unifiedto.String("06e61b0d-3087-414c-a0a3-d98637ca85c3"),
            Location: unifiedto.String("delectus"),
            ManagerID: unifiedto.String("repudiandae"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedto.String("Erin Kris"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "harum",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminatedAt: types.MustDateFromString("2021-02-20"),
            Title: unifiedto.String("Mrs."),
            UpdatedAt: types.MustDateFromString("2022-01-28"),
        },
        ConnectionID: "placeat",
        ID: "98f13af2-8db2-4cf2-bf4f-3ded356d7e14",
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


## PutHrisConnectionIDGroupID

Update a Group

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
    operationSecurity := operations.PutHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Hris.PutHrisConnectionIDGroupID(ctx, operations.PutHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2022-08-29"),
            Description: unifiedto.String("beatae"),
            EmployeeIds: []string{
                "eligendi",
            },
            ID: unifiedto.String("d98196d5-5af6-49a1-84b7-9ae33681c23c"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "dolorem",
            },
            Name: unifiedto.String("Grant Klein PhD"),
            ParentID: unifiedto.String("ab"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDivision.ToPointer(),
            UpdatedAt: types.MustDateFromString("2020-11-07"),
        },
        ConnectionID: "quasi",
        ID: "2c5ba825-fe22-4cd5-8ba6-fbfec932af68",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutHrisConnectionIDGroupIDRequest](../../models/operations/puthrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PutHrisConnectionIDGroupIDSecurity](../../models/operations/puthrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PutHrisConnectionIDGroupIDResponse](../../models/operations/puthrisconnectionidgroupidresponse.md), error**

