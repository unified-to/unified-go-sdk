# Scorecard
(*Scorecard*)

### Available Operations

* [CreateAtsScorecard](#createatsscorecard) - Create a scorecard
* [GetAtsScorecard](#getatsscorecard) - Retrieve a scorecard
* [ListAtsScorecards](#listatsscorecards) - List all scorecards
* [PatchAtsScorecard](#patchatsscorecard) - Update a scorecard
* [RemoveAtsScorecard](#removeatsscorecard) - Remove a scorecard
* [UpdateAtsScorecard](#updateatsscorecard) - Update a scorecard

## CreateAtsScorecard

Create a scorecard

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
    res, err := s.Scorecard.CreateAtsScorecard(ctx, operations.CreateAtsScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            Raw: shared.PropertyAtsScorecardRaw{},
        },
        ConnectionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateAtsScorecardRequest](../../models/operations/createatsscorecardrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateAtsScorecardResponse](../../models/operations/createatsscorecardresponse.md), error**


## GetAtsScorecard

Retrieve a scorecard

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
    res, err := s.Scorecard.GetAtsScorecard(ctx, operations.GetAtsScorecardRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAtsScorecardRequest](../../models/operations/getatsscorecardrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetAtsScorecardResponse](../../models/operations/getatsscorecardresponse.md), error**


## ListAtsScorecards

List all scorecards

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
    res, err := s.Scorecard.ListAtsScorecards(ctx, operations.ListAtsScorecardsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListAtsScorecardsRequest](../../models/operations/listatsscorecardsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListAtsScorecardsResponse](../../models/operations/listatsscorecardsresponse.md), error**


## PatchAtsScorecard

Update a scorecard

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
    res, err := s.Scorecard.PatchAtsScorecard(ctx, operations.PatchAtsScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            Raw: shared.PropertyAtsScorecardRaw{},
        },
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchAtsScorecardRequest](../../models/operations/patchatsscorecardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PatchAtsScorecardResponse](../../models/operations/patchatsscorecardresponse.md), error**


## RemoveAtsScorecard

Remove a scorecard

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
    res, err := s.Scorecard.RemoveAtsScorecard(ctx, operations.RemoveAtsScorecardRequest{
        ConnectionID: "string",
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
| `request`                                                                                    | [operations.RemoveAtsScorecardRequest](../../models/operations/removeatsscorecardrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RemoveAtsScorecardResponse](../../models/operations/removeatsscorecardresponse.md), error**


## UpdateAtsScorecard

Update a scorecard

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
    res, err := s.Scorecard.UpdateAtsScorecard(ctx, operations.UpdateAtsScorecardRequest{
        AtsScorecard: &shared.AtsScorecard{
            Raw: shared.PropertyAtsScorecardRaw{},
        },
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateAtsScorecardRequest](../../models/operations/updateatsscorecardrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateAtsScorecardResponse](../../models/operations/updateatsscorecardresponse.md), error**

