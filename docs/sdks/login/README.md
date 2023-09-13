# Login

### Available Operations

* [GetUnifiedIntegrationLoginWorkspaceIDIntegrationType](#getunifiedintegrationloginworkspaceidintegrationtype) - Sign in a user

## GetUnifiedIntegrationLoginWorkspaceIDIntegrationType

Returns an authentication URL for the specified integration.  Once a successful authentication occurs, the name and emails are returned.

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

    ctx := context.Background()
    res, err := s.Login.GetUnifiedIntegrationLoginWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeRequest{
        Env: unifiedto.String("veniam"),
        FailureRedirect: unifiedto.String("nemo"),
        IntegrationType: "voluptatum",
        Redirect: unifiedto.Bool(false),
        State: unifiedto.String("quia"),
        SuccessRedirect: unifiedto.String("quisquam"),
        WorkspaceID: "et",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedIntegrationLoginWorkspaceIDIntegrationType200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeRequest](../../models/operations/getunifiedintegrationloginworkspaceidintegrationtyperequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeResponse](../../models/operations/getunifiedintegrationloginworkspaceidintegrationtyperesponse.md), error**

