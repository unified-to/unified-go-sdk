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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.DeleteHrisConnectionIDEmployeeID(ctx, operations.DeleteHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "excepturi",
        ID: "3e34316c-f55b-4431-b553-ccf1c204c4ad",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteHrisConnectionIDEmployeeIDRequest](../../models/operations/deletehrisconnectionidemployeeidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.DeleteHrisConnectionIDGroupID(ctx, operations.DeleteHrisConnectionIDGroupIDRequest{
        ConnectionID: "quod",
        ID: "c9904c51-95b8-4648-8efa-78f1e2d3b901",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteHrisConnectionIDGroupIDRequest](../../models/operations/deletehrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDEmployee(ctx, operations.GetHrisConnectionIDEmployeeRequest{
        ConnectionID: "saepe",
        Limit: unifiedgosdk.Float64(289.94),
        Offset: unifiedgosdk.Float64(5725.89),
        Order: unifiedgosdk.String("corporis"),
        Query: unifiedgosdk.String("explicabo"),
        Sort: unifiedgosdk.String("distinctio"),
        UpdatedGte: types.MustTimeFromString("2021-08-12T22:16:16.253Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployees != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetHrisConnectionIDEmployeeRequest](../../models/operations/gethrisconnectionidemployeerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDEmployeeID(ctx, operations.GetHrisConnectionIDEmployeeIDRequest{
        ConnectionID: "tempora",
        ID: "cbb19f71-3d95-4a41-a9c1-387271e18ea9",
    })
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
| `request`                                                                                                          | [operations.GetHrisConnectionIDEmployeeIDRequest](../../models/operations/gethrisconnectionidemployeeidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDGroup(ctx, operations.GetHrisConnectionIDGroupRequest{
        ConnectionID: "debitis",
        Limit: unifiedgosdk.Float64(2501.58),
        Offset: unifiedgosdk.Float64(3333.54),
        Order: unifiedgosdk.String("illo"),
        Query: unifiedgosdk.String("illo"),
        Sort: unifiedgosdk.String("deleniti"),
        UpdatedGte: types.MustTimeFromString("2022-07-26T10:34:40.123Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetHrisConnectionIDGroupRequest](../../models/operations/gethrisconnectionidgrouprequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.GetHrisConnectionIDGroupID(ctx, operations.GetHrisConnectionIDGroupIDRequest{
        ConnectionID: "optio",
        ID: "c57fbd60-b1a7-48ed-a9a9-d4eea85658c2",
    })
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
| `request`                                                                                                    | [operations.GetHrisConnectionIDGroupIDRequest](../../models/operations/gethrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PatchHrisConnectionIDEmployeeID(ctx, operations.PatchHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.String("at"),
                Address2: unifiedgosdk.String("modi"),
                City: unifiedgosdk.String("Findlay"),
                Country: unifiedgosdk.String("Saint Martin"),
                CountryCode: unifiedgosdk.String("LI"),
                PostalCode: unifiedgosdk.String("69291-4598"),
                Region: unifiedgosdk.String("unde"),
                RegionCode: unifiedgosdk.String("autem"),
            },
            CreatedAt: types.MustTimeFromString("2022-07-11T08:13:45.102Z"),
            DateOfBirth: types.MustTimeFromString("2022-02-18T07:36:00.073Z"),
            Department: unifiedgosdk.String("autem"),
            Division: unifiedgosdk.String("placeat"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Blaze97@yahoo.com",
                    Type: shared.HrisEmailTypeOther.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("id"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeIntern.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustTimeFromString("2020-08-13T08:33:47.369Z"),
            ID: unifiedgosdk.String("ef234c95-5b9b-4df2-990a-bd9bbcc2725e"),
            Location: unifiedgosdk.String("impedit"),
            ManagerID: unifiedgosdk.String("magni"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedgosdk.String("Sabrina Schamberger Sr."),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "totam",
                    Type: shared.HrisTelephoneTypeWork.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2022-06-27T13:37:26.116Z"),
            Title: unifiedgosdk.String("Mr."),
            UpdatedAt: types.MustTimeFromString("2021-10-27T12:07:56.119Z"),
        },
        ConnectionID: "excepturi",
        ID: "ef68e45c-8add-4fac-b545-00430c6632b4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchHrisConnectionIDEmployeeIDRequest](../../models/operations/patchhrisconnectionidemployeeidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PatchHrisConnectionIDGroupID(ctx, operations.PatchHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2022-05-31T02:06:03.542Z"),
            Description: unifiedgosdk.String("inventore"),
            EmployeeIds: []string{
                "sapiente",
            },
            ID: unifiedgosdk.String("df01c3e9-1e8f-47bc-a9d4-60a77eceb26d"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "architecto",
            },
            Name: unifiedgosdk.String("Lorene Bosco"),
            ParentID: unifiedgosdk.String("qui"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDepartment.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-12-06T03:55:58.108Z"),
        },
        ConnectionID: "quisquam",
        ID: "7c0f0f87-3f9d-45c2-9fd3-e0b4a4a4253c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchHrisConnectionIDGroupIDRequest](../../models/operations/patchhrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PostHrisConnectionIDEmployee(ctx, operations.PostHrisConnectionIDEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.String("amet"),
                Address2: unifiedgosdk.String("ipsa"),
                City: unifiedgosdk.String("West Jaunitaland"),
                Country: unifiedgosdk.String("Bahrain"),
                CountryCode: unifiedgosdk.String("YE"),
                PostalCode: unifiedgosdk.String("18494"),
                Region: unifiedgosdk.String("pariatur"),
                RegionCode: unifiedgosdk.String("porro"),
            },
            CreatedAt: types.MustTimeFromString("2022-09-23T16:00:07.090Z"),
            DateOfBirth: types.MustTimeFromString("2021-07-11T10:12:07.009Z"),
            Department: unifiedgosdk.String("itaque"),
            Division: unifiedgosdk.String("sit"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Timmothy68@yahoo.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("culpa"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeFullTime.ToPointer(),
            Gender: shared.HrisEmployeeGenderFemale.ToPointer(),
            HiredAt: types.MustTimeFromString("2021-01-27T23:42:02.589Z"),
            ID: unifiedgosdk.String("12a4ba9d-5998-4819-acfd-0c77c53e7e7d"),
            Location: unifiedgosdk.String("eius"),
            ManagerID: unifiedgosdk.String("accusamus"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusSingle.ToPointer(),
            Name: unifiedgosdk.String("Sophie Lesch"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "consequatur",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2021-06-09T00:38:41.199Z"),
            Title: unifiedgosdk.String("Mrs."),
            UpdatedAt: types.MustTimeFromString("2022-06-19T22:55:28.813Z"),
        },
        ConnectionID: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostHrisConnectionIDEmployeeRequest](../../models/operations/posthrisconnectionidemployeerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PostHrisConnectionIDGroup(ctx, operations.PostHrisConnectionIDGroupRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2022-10-11T00:06:05.180Z"),
            Description: unifiedgosdk.String("sint"),
            EmployeeIds: []string{
                "ea",
            },
            ID: unifiedgosdk.String("703fec31-c508-424d-989a-36a6b2d27eb7"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "accusantium",
            },
            Name: unifiedgosdk.String("Harriet Orn DDS"),
            ParentID: unifiedgosdk.String("voluptatum"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeSubDepartment.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2022-02-14T18:55:03.442Z"),
        },
        ConnectionID: "commodi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostHrisConnectionIDGroupRequest](../../models/operations/posthrisconnectionidgrouprequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PutHrisConnectionIDEmployeeID(ctx, operations.PutHrisConnectionIDEmployeeIDRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.String("debitis"),
                Address2: unifiedgosdk.String("commodi"),
                City: unifiedgosdk.String("South Izaiahhaven"),
                Country: unifiedgosdk.String("Paraguay"),
                CountryCode: unifiedgosdk.String("MS"),
                PostalCode: unifiedgosdk.String("72740-9977"),
                Region: unifiedgosdk.String("nisi"),
                RegionCode: unifiedgosdk.String("occaecati"),
            },
            CreatedAt: types.MustTimeFromString("2022-12-15T12:15:10.431Z"),
            DateOfBirth: types.MustTimeFromString("2020-03-31T16:49:34.939Z"),
            Department: unifiedgosdk.String("odio"),
            Division: unifiedgosdk.String("nihil"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Tamia.Doyle@gmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("iusto"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeOther.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustTimeFromString("2022-06-18T09:54:51.929Z"),
            ID: unifiedgosdk.String("06e61b0d-3087-414c-a0a3-d98637ca85c3"),
            Location: unifiedgosdk.String("delectus"),
            ManagerID: unifiedgosdk.String("repudiandae"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedgosdk.String("Erin Kris"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "harum",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2021-02-20T15:13:08.610Z"),
            Title: unifiedgosdk.String("Mrs."),
            UpdatedAt: types.MustTimeFromString("2022-01-28T12:53:07.770Z"),
        },
        ConnectionID: "placeat",
        ID: "98f13af2-8db2-4cf2-bf4f-3ded356d7e14",
    })
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
| `request`                                                                                                          | [operations.PutHrisConnectionIDEmployeeIDRequest](../../models/operations/puthrisconnectionidemployeeidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hris.PutHrisConnectionIDGroupID(ctx, operations.PutHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2022-08-29T10:06:09.751Z"),
            Description: unifiedgosdk.String("beatae"),
            EmployeeIds: []string{
                "eligendi",
            },
            ID: unifiedgosdk.String("d98196d5-5af6-49a1-84b7-9ae33681c23c"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "dolorem",
            },
            Name: unifiedgosdk.String("Grant Klein PhD"),
            ParentID: unifiedgosdk.String("ab"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDivision.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2020-11-07T15:44:18.546Z"),
        },
        ConnectionID: "quasi",
        ID: "2c5ba825-fe22-4cd5-8ba6-fbfec932af68",
    })
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
| `request`                                                                                                    | [operations.PutHrisConnectionIDGroupIDRequest](../../models/operations/puthrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutHrisConnectionIDGroupIDResponse](../../models/operations/puthrisconnectionidgroupidresponse.md), error**

