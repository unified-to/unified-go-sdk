# Employee
(*Employee*)

### Available Operations

* [CreateHrisEmployee](#createhrisemployee) - Create an employee
* [GetHrisEmployee](#gethrisemployee) - Retrieve an employee
* [ListHrisEmployees](#listhrisemployees) - List all employees
* [PatchHrisEmployee](#patchhrisemployee) - Update an employee
* [RemoveHrisEmployee](#removehrisemployee) - Remove an employee
* [UpdateHrisEmployee](#updatehrisemployee) - Update an employee

## CreateHrisEmployee

Create an employee

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.CreateHrisEmployee(ctx, operations.CreateHrisEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{},
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Adriel_Hansen@hotmail.com",
                },
            },
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "Gloves with",
                },
            },
        },
        ConnectionID: "ack Recycled",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateHrisEmployeeRequest](../../models/operations/createhrisemployeerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateHrisEmployeeResponse](../../models/operations/createhrisemployeeresponse.md), error**


## GetHrisEmployee

Retrieve an employee

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.GetHrisEmployee(ctx, operations.GetHrisEmployeeRequest{
        ConnectionID: "Automated",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetHrisEmployeeRequest](../../models/operations/gethrisemployeerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetHrisEmployeeResponse](../../models/operations/gethrisemployeeresponse.md), error**


## ListHrisEmployees

List all employees

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.ListHrisEmployees(ctx, operations.ListHrisEmployeesRequest{
        ConnectionID: "Table moratorium",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListHrisEmployeesRequest](../../models/operations/listhrisemployeesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListHrisEmployeesResponse](../../models/operations/listhrisemployeesresponse.md), error**


## PatchHrisEmployee

Update an employee

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.PatchHrisEmployee(ctx, operations.PatchHrisEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{},
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta.Cassin@yahoo.com",
                },
            },
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "shameful",
                },
            },
        },
        ConnectionID: "barring transmitting Hybrid",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchHrisEmployeeRequest](../../models/operations/patchhrisemployeerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PatchHrisEmployeeResponse](../../models/operations/patchhrisemployeeresponse.md), error**


## RemoveHrisEmployee

Remove an employee

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.RemoveHrisEmployee(ctx, operations.RemoveHrisEmployeeRequest{
        ConnectionID: "Architect",
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
| `request`                                                                                    | [operations.RemoveHrisEmployeeRequest](../../models/operations/removehrisemployeerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RemoveHrisEmployeeResponse](../../models/operations/removehrisemployeeresponse.md), error**


## UpdateHrisEmployee

Update an employee

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Employee.UpdateHrisEmployee(ctx, operations.UpdateHrisEmployeeRequest{
        HrisEmployee: &shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{},
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Abbie_Aufderhar@gmail.com",
                },
            },
            Raw: &shared.PropertyHrisEmployeeRaw{},
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "Web Borders North",
                },
            },
        },
        ConnectionID: "Jewelery slap",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateHrisEmployeeRequest](../../models/operations/updatehrisemployeerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateHrisEmployeeResponse](../../models/operations/updatehrisemployeeresponse.md), error**

