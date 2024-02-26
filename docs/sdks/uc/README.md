# Uc
(*Uc*)

### Available Operations

* [CreateUcContact](#createuccontact) - Create a contact
* [GetUcContact](#getuccontact) - Retrieve a contact
* [ListUcCalls](#listuccalls) - List all calls
* [ListUcContacts](#listuccontacts) - List all contacts
* [PatchUcContact](#patchuccontact) - Update a contact
* [RemoveUcContact](#removeuccontact) - Remove a contact
* [UpdateUcContact](#updateuccontact) - Update a contact

## CreateUcContact

Create a contact

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


    operationSecurity := operations.CreateUcContactSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.CreateUcContact(ctx, operations.CreateUcContactRequest{
        ConnectionID: "<value>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateUcContactRequest](../../pkg/models/operations/createuccontactrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateUcContactSecurity](../../pkg/models/operations/createuccontactsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateUcContactResponse](../../pkg/models/operations/createuccontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUcContact

Retrieve a contact

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


    operationSecurity := operations.GetUcContactSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.GetUcContact(ctx, operations.GetUcContactRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUcContactRequest](../../pkg/models/operations/getuccontactrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetUcContactSecurity](../../pkg/models/operations/getuccontactsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetUcContactResponse](../../pkg/models/operations/getuccontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUcCalls

List all calls

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


    operationSecurity := operations.ListUcCallsSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.ListUcCalls(ctx, operations.ListUcCallsRequest{
        ConnectionID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListUcCallsRequest](../../pkg/models/operations/listuccallsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListUcCallsSecurity](../../pkg/models/operations/listuccallssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListUcCallsResponse](../../pkg/models/operations/listuccallsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUcContacts

List all contacts

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


    operationSecurity := operations.ListUcContactsSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.ListUcContacts(ctx, operations.ListUcContactsRequest{
        ConnectionID: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListUcContactsRequest](../../pkg/models/operations/listuccontactsrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListUcContactsSecurity](../../pkg/models/operations/listuccontactssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListUcContactsResponse](../../pkg/models/operations/listuccontactsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchUcContact

Update a contact

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


    operationSecurity := operations.PatchUcContactSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.PatchUcContact(ctx, operations.PatchUcContactRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchUcContactRequest](../../pkg/models/operations/patchuccontactrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.PatchUcContactSecurity](../../pkg/models/operations/patchuccontactsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.PatchUcContactResponse](../../pkg/models/operations/patchuccontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveUcContact

Remove a contact

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


    operationSecurity := operations.RemoveUcContactSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.RemoveUcContact(ctx, operations.RemoveUcContactRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveUcContactRequest](../../pkg/models/operations/removeuccontactrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.RemoveUcContactSecurity](../../pkg/models/operations/removeuccontactsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.RemoveUcContactResponse](../../pkg/models/operations/removeuccontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUcContact

Update a contact

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


    operationSecurity := operations.UpdateUcContactSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Uc.UpdateUcContact(ctx, operations.UpdateUcContactRequest{
        ConnectionID: "<value>",
        ID: "<id>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateUcContactRequest](../../pkg/models/operations/updateuccontactrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateUcContactSecurity](../../pkg/models/operations/updateuccontactsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateUcContactResponse](../../pkg/models/operations/updateuccontactresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
