# Employee

## Overview

### Available Operations

* [CreateHrisEmployee](#createhrisemployee) - Create an employee
* [GetHrisEmployee](#gethrisemployee) - Retrieve an employee
* [ListHrisEmployees](#listhrisemployees) - List all employees
* [PatchHrisEmployee](#patchhrisemployee) - Update an employee
* [RemoveHrisEmployee](#removehrisemployee) - Remove an employee
* [UpdateHrisEmployee](#updatehrisemployee) - Update an employee

## CreateHrisEmployee

Creates a HrisEmployee. Provide the unified HrisEmployee fields in the request body; they are mapped to the provider's API. Returns the created HrisEmployee.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisEmployee" method="post" path="/hris/{connection_id}/employee" example="hris_employee" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.CreateHrisEmployee(ctx, operations.CreateHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("89ece618-3681-468a-8f17-b6da53e1d01b"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b870637f-d229-4b09-9ebe-d6c2baa4179f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.156Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisEmployeeRequest](../../pkg/models/operations/createhrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisEmployeeResponse](../../pkg/models/operations/createhrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisEmployee

Retrieves a single HrisEmployee by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisEmployee" method="get" path="/hris/{connection_id}/employee/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.GetHrisEmployee(ctx, operations.GetHrisEmployeeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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
| `request`                                                                                  | [operations.GetHrisEmployeeRequest](../../pkg/models/operations/gethrisemployeerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisEmployeeResponse](../../pkg/models/operations/gethrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisEmployees

Returns a list of HrisEmployees. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisEmployees" method="get" path="/hris/{connection_id}/employee" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.ListHrisEmployees(ctx, operations.ListHrisEmployeesRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisEmployeesRequest](../../pkg/models/operations/listhrisemployeesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisEmployeesResponse](../../pkg/models/operations/listhrisemployeesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisEmployee

Updates the HrisEmployee identified by `id`. Provide the unified HrisEmployee fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisEmployee" method="patch" path="/hris/{connection_id}/employee/{id}" example="hris_employee" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.PatchHrisEmployee(ctx, operations.PatchHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("ffeb673d-ef5e-4f46-bf0f-3d93b4b2ad4d"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5c436a00-8fde-4c33-b3a7-bfa83a5bad6b"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.164Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisEmployeeRequest](../../pkg/models/operations/patchhrisemployeerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisEmployeeResponse](../../pkg/models/operations/patchhrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisEmployee

Deletes the HrisEmployee identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisEmployee" method="delete" path="/hris/{connection_id}/employee/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.RemoveHrisEmployee(ctx, operations.RemoveHrisEmployeeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveHrisEmployeeRequest](../../pkg/models/operations/removehrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisEmployeeResponse](../../pkg/models/operations/removehrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisEmployee

Updates the HrisEmployee identified by `id`. Provide the unified HrisEmployee fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisEmployee" method="put" path="/hris/{connection_id}/employee/{id}" example="hris_employee" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Employee.UpdateHrisEmployee(ctx, operations.UpdateHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("ffeb673d-ef5e-4f46-bf0f-3d93b4b2ad4d"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5c436a00-8fde-4c33-b3a7-bfa83a5bad6b"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.164Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisEmployeeRequest](../../pkg/models/operations/updatehrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisEmployeeResponse](../../pkg/models/operations/updatehrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |