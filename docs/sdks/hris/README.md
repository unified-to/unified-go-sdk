# Hris
(*Hris*)

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
        ConnectionID: "Laredo turquoise port",
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
        ConnectionID: "consequently platforms Metal",
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
        ConnectionID: "initiatives greedily project",
        Limit: unifiedgosdk.Float64(1798.52),
        Offset: unifiedgosdk.Float64(6683.19),
        Order: unifiedgosdk.String("Sports"),
        Query: unifiedgosdk.String("TLS"),
        Sort: unifiedgosdk.String("Jazz Trans"),
        UpdatedGte: types.MustTimeFromString("2021-04-09T17:32:06.988Z"),
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
        ConnectionID: "Keyboard cleverly Rubber",
        ID: "<ID>",
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
        ConnectionID: "Loan",
        Limit: unifiedgosdk.Float64(3486.96),
        Offset: unifiedgosdk.Float64(9705.73),
        Order: unifiedgosdk.String("Coordinator"),
        Query: unifiedgosdk.String("World"),
        Sort: unifiedgosdk.String("Dollar"),
        UpdatedGte: types.MustTimeFromString("2021-01-15T16:06:13.340Z"),
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
        ConnectionID: "behind",
        ID: "<ID>",
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
                Address1: unifiedgosdk.String("Tennessee virtual"),
                Address2: unifiedgosdk.String("Assurance forecast"),
                City: unifiedgosdk.String("Baton Rouge"),
                Country: unifiedgosdk.String("Bahamas"),
                CountryCode: unifiedgosdk.String("TW"),
                PostalCode: unifiedgosdk.String("93632"),
                Region: unifiedgosdk.String("Bailey"),
                RegionCode: unifiedgosdk.String("navigating Oregon"),
            },
            CreatedAt: types.MustTimeFromString("2023-01-31T08:11:49.561Z"),
            DateOfBirth: types.MustTimeFromString("2021-10-08T23:14:10.860Z"),
            Department: unifiedgosdk.String("soupy web Robust"),
            Division: unifiedgosdk.String("Corporate loudly quantify"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Wendy_Kovacek@hotmail.com",
                    Type: shared.HrisEmailTypeOther.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("hydrate indigo transmit"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeFullTime.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustTimeFromString("2021-07-20T11:37:42.486Z"),
            ID: unifiedgosdk.String("<ID>"),
            Location: unifiedgosdk.String("withdrawal wonderfully"),
            ManagerID: unifiedgosdk.String("molestias white Gainesville"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedgosdk.String("tensely technologies"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "Northeast Music Hassium",
                    Type: shared.HrisTelephoneTypeMobile.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2023-08-07T01:58:28.622Z"),
            Title: unifiedgosdk.String("failing Southwest Kuhn"),
            UpdatedAt: types.MustTimeFromString("2023-05-18T01:32:56.083Z"),
        },
        ConnectionID: "South Money past",
        ID: "<ID>",
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
            CreatedAt: types.MustTimeFromString("2023-10-19T05:30:26.390Z"),
            Description: unifiedgosdk.String("Stand-alone asymmetric orchestration"),
            EmployeeIds: []string{
                "shootdown",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "24/7",
            },
            Name: unifiedgosdk.String("Agender trainer"),
            ParentID: unifiedgosdk.String("Configuration Kids Sedan"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDivision.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-08-18T16:48:12.885Z"),
        },
        ConnectionID: "Intersex",
        ID: "<ID>",
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
                Address1: unifiedgosdk.String("Southeast ea withdrawal"),
                Address2: unifiedgosdk.String("Developer"),
                City: unifiedgosdk.String("Grand Forks"),
                Country: unifiedgosdk.String("Cayman Islands"),
                CountryCode: unifiedgosdk.String("BM"),
                PostalCode: unifiedgosdk.String("63867-8134"),
                Region: unifiedgosdk.String("teal Northwest firewall"),
                RegionCode: unifiedgosdk.String("doubt Diesel COM"),
            },
            CreatedAt: types.MustTimeFromString("2022-09-25T20:39:21.870Z"),
            DateOfBirth: types.MustTimeFromString("2022-04-06T20:53:56.362Z"),
            Department: unifiedgosdk.String("payment mull"),
            Division: unifiedgosdk.String("Blues red"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Jacquelyn8@hotmail.com",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("North Southeast"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusInactive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HiredAt: types.MustTimeFromString("2021-03-26T14:20:42.258Z"),
            ID: unifiedgosdk.String("<ID>"),
            Location: unifiedgosdk.String("East"),
            ManagerID: unifiedgosdk.String("Maserati"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusSingle.ToPointer(),
            Name: unifiedgosdk.String("Xenogender copy"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "hmph",
                    Type: shared.HrisTelephoneTypeWork.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2022-08-16T03:13:22.861Z"),
            Title: unifiedgosdk.String("Regional synthesize"),
            UpdatedAt: types.MustTimeFromString("2022-06-15T02:35:02.446Z"),
        },
        ConnectionID: "past",
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
            CreatedAt: types.MustTimeFromString("2021-02-23T15:35:38.483Z"),
            Description: unifiedgosdk.String("Configurable stable product"),
            EmployeeIds: []string{
                "Auto",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "JSON",
            },
            Name: unifiedgosdk.String("whereas Usability transmitting"),
            ParentID: unifiedgosdk.String("invoice Cyclocross Electric"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeGroup.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-04-30T12:40:50.129Z"),
        },
        ConnectionID: "Hybrid Schenectady",
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
                Address1: unifiedgosdk.String("Northwest"),
                Address2: unifiedgosdk.String("and"),
                City: unifiedgosdk.String("Uniquefield"),
                Country: unifiedgosdk.String("Virgin Islands, British"),
                CountryCode: unifiedgosdk.String("AE"),
                PostalCode: unifiedgosdk.String("41682"),
                Region: unifiedgosdk.String("Adventure Avon"),
                RegionCode: unifiedgosdk.String("bah South"),
            },
            CreatedAt: types.MustTimeFromString("2023-07-18T13:59:47.040Z"),
            DateOfBirth: types.MustTimeFromString("2022-04-19T17:38:57.783Z"),
            Department: unifiedgosdk.String("West auxiliary"),
            Division: unifiedgosdk.String("volt"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Kenton_Turcotte@gmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.String("DNS coulomb Berkshire"),
            EmploymentStatus: shared.HrisEmployeeEmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeCasual.ToPointer(),
            Gender: shared.HrisEmployeeGenderFemale.ToPointer(),
            HiredAt: types.MustTimeFromString("2022-10-08T23:22:26.211Z"),
            ID: unifiedgosdk.String("<ID>"),
            Location: unifiedgosdk.String("East primary"),
            ManagerID: unifiedgosdk.String("Tokelau"),
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Name: unifiedgosdk.String("Bespoke Investment"),
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "East Investment",
                    Type: shared.HrisTelephoneTypeOther.ToPointer(),
                },
            },
            TerminatedAt: types.MustTimeFromString("2022-03-28T08:29:30.386Z"),
            Title: unifiedgosdk.String("unless"),
            UpdatedAt: types.MustTimeFromString("2022-06-29T10:38:14.570Z"),
        },
        ConnectionID: "Designer Tennessine",
        ID: "<ID>",
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
            CreatedAt: types.MustTimeFromString("2022-08-10T12:11:42.375Z"),
            Description: unifiedgosdk.String("Decentralized methodical projection"),
            EmployeeIds: []string{
                "Credit",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "South",
            },
            Name: unifiedgosdk.String("Jeep brr Northwest"),
            ParentID: unifiedgosdk.String("quickly Licensed"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-11-08T00:11:45.458Z"),
        },
        ConnectionID: "vortals interface Gasoline",
        ID: "<ID>",
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

