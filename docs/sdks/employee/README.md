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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.CreateHrisEmployeeSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.CreateHrisEmployee(ctx, operations.CreateHrisEmployeeRequest{
        ConnectionID: "<value>",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisEmployeeRequest](../../pkg/models/operations/createhrisemployeerequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CreateHrisEmployeeSecurity](../../pkg/models/operations/createhrisemployeesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CreateHrisEmployeeResponse](../../pkg/models/operations/createhrisemployeeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetHrisEmployee

Retrieve an employee

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.GetHrisEmployeeSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.GetHrisEmployee(ctx, operations.GetHrisEmployeeRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisEmployeeRequest](../../pkg/models/operations/gethrisemployeerequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetHrisEmployeeSecurity](../../pkg/models/operations/gethrisemployeesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.GetHrisEmployeeResponse](../../pkg/models/operations/gethrisemployeeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListHrisEmployees

List all employees

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.ListHrisEmployeesSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.ListHrisEmployees(ctx, operations.ListHrisEmployeesRequest{
        ConnectionID: "<value>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisEmployeesRequest](../../pkg/models/operations/listhrisemployeesrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ListHrisEmployeesSecurity](../../pkg/models/operations/listhrisemployeessecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ListHrisEmployeesResponse](../../pkg/models/operations/listhrisemployeesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchHrisEmployee

Update an employee

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.PatchHrisEmployeeSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.PatchHrisEmployee(ctx, operations.PatchHrisEmployeeRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisEmployeeRequest](../../pkg/models/operations/patchhrisemployeerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.PatchHrisEmployeeSecurity](../../pkg/models/operations/patchhrisemployeesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.PatchHrisEmployeeResponse](../../pkg/models/operations/patchhrisemployeeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveHrisEmployee

Remove an employee

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.RemoveHrisEmployeeSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.RemoveHrisEmployee(ctx, operations.RemoveHrisEmployeeRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveHrisEmployeeRequest](../../pkg/models/operations/removehrisemployeerequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.RemoveHrisEmployeeSecurity](../../pkg/models/operations/removehrisemployeesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.RemoveHrisEmployeeResponse](../../pkg/models/operations/removehrisemployeeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateHrisEmployee

Update an employee

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.UpdateHrisEmployeeSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Employee.UpdateHrisEmployee(ctx, operations.UpdateHrisEmployeeRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisEmployeeRequest](../../pkg/models/operations/updatehrisemployeerequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpdateHrisEmployeeSecurity](../../pkg/models/operations/updatehrisemployeesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpdateHrisEmployeeResponse](../../pkg/models/operations/updatehrisemployeeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
