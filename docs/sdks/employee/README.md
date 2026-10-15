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

Create an employee

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
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](60136.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeStockOptions.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusSalariedNonexempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(false),
            HiredAt: types.MustNewTimeFromString("2023-05-24T16:00:07.841Z"),
            ID: unifiedgosdk.Pointer("cea8919c-01b7-4fd0-8a41-647854da3816"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f505b9c1-5cae-41b6-a4db-c346d0f6b134"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "appono",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyBiweekly.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Margot11@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Jaden_Miller@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Electronics"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](1590710.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](730343.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-28T14:31:33.603Z"),
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

Retrieve an employee

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

List all employees

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

Update an employee

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
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](60136.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeStockOptions.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusSalariedNonexempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(false),
            HiredAt: types.MustNewTimeFromString("2023-05-24T16:00:07.895Z"),
            ID: unifiedgosdk.Pointer("cda9f7dc-db25-488a-b95d-7b18bccf4f17"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5b35d173-1d68-40f5-bec4-8809d598a9af"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "appono",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyBiweekly.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Margot11@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Jaden_Miller@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Electronics"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](1590710.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](730343.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-28T14:31:33.639Z"),
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

Remove an employee

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

Update an employee

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
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](60136.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeStockOptions.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusSalariedNonexempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(false),
            HiredAt: types.MustNewTimeFromString("2023-05-24T16:00:07.895Z"),
            ID: unifiedgosdk.Pointer("cda9f7dc-db25-488a-b95d-7b18bccf4f17"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5b35d173-1d68-40f5-bec4-8809d598a9af"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "appono",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyBiweekly.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Margot11@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Jaden_Miller@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Electronics"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](1590710.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](730343.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-28T14:31:33.639Z"),
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