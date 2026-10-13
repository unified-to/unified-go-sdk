# Hris

## Overview

### Available Operations

* [CreateHrisAttendance](#createhrisattendance) - Create an attendance
* [CreateHrisBankaccount](#createhrisbankaccount) - Create a bankaccount
* [CreateHrisBenefit](#createhrisbenefit) - Create a benefit
* [CreateHrisCompany](#createhriscompany) - Create a company
* [CreateHrisDeduction](#createhrisdeduction) - Create a deduction
* [CreateHrisDevice](#createhrisdevice) - Create a device
* [CreateHrisDocument](#createhrisdocument) - Create a document
* [CreateHrisEarning](#createhrisearning) - Create an earning
* [CreateHrisEmployee](#createhrisemployee) - Create an employee
* [CreateHrisGroup](#createhrisgroup) - Create a group
* [CreateHrisJob](#createhrisjob) - Create a job
* [CreateHrisLocation](#createhrislocation) - Create a location
* [CreateHrisTaxonomy](#createhristaxonomy) - Create a taxonomy
* [CreateHrisTimeoff](#createhristimeoff) - Create a timeoff
* [CreateHrisTimeshift](#createhristimeshift) - Create a timeshift
* [GetHrisAttendance](#gethrisattendance) - Retrieve an attendance
* [GetHrisBankaccount](#gethrisbankaccount) - Retrieve a bankaccount
* [GetHrisBenefit](#gethrisbenefit) - Retrieve a benefit
* [GetHrisCompany](#gethriscompany) - Retrieve a company
* [GetHrisDeduction](#gethrisdeduction) - Retrieve a deduction
* [GetHrisDevice](#gethrisdevice) - Retrieve a device
* [GetHrisDocument](#gethrisdocument) - Retrieve a document
* [GetHrisEarning](#gethrisearning) - Retrieve an earning
* [GetHrisEmployee](#gethrisemployee) - Retrieve an employee
* [GetHrisGroup](#gethrisgroup) - Retrieve a group
* [GetHrisJob](#gethrisjob) - Retrieve a job
* [GetHrisLocation](#gethrislocation) - Retrieve a location
* [GetHrisPaycode](#gethrispaycode) - Retrieve a paycode
* [GetHrisPayroll](#gethrispayroll) - Retrieve a payroll
* [GetHrisPayslip](#gethrispayslip) - Retrieve a payslip
* [GetHrisTaxonomy](#gethristaxonomy) - Retrieve a taxonomy
* [GetHrisTimeoff](#gethristimeoff) - Retrieve a timeoff
* [GetHrisTimeshift](#gethristimeshift) - Retrieve a timeshift
* [ListHrisAttendances](#listhrisattendances) - List all attendances
* [ListHrisBankaccounts](#listhrisbankaccounts) - List all bankaccounts
* [ListHrisBenefits](#listhrisbenefits) - List all benefits
* [ListHrisCompanies](#listhriscompanies) - List all companies
* [ListHrisDeductions](#listhrisdeductions) - List all deductions
* [ListHrisDevices](#listhrisdevices) - List all devices
* [ListHrisDocuments](#listhrisdocuments) - List all documents
* [ListHrisEarnings](#listhrisearnings) - List all earnings
* [ListHrisEmployees](#listhrisemployees) - List all employees
* [ListHrisGroups](#listhrisgroups) - List all groups
* [ListHrisJobs](#listhrisjobs) - List all jobs
* [ListHrisLocations](#listhrislocations) - List all locations
* [ListHrisPaycodes](#listhrispaycodes) - List all paycodes
* [ListHrisPayrolls](#listhrispayrolls) - List all payrolls
* [ListHrisPayslips](#listhrispayslips) - List all payslips
* [ListHrisTaxonomies](#listhristaxonomies) - List all taxonomies
* [ListHrisTimeoffs](#listhristimeoffs) - List all timeoffs
* [ListHrisTimeshifts](#listhristimeshifts) - List all timeshifts
* [PatchHrisAttendance](#patchhrisattendance) - Update an attendance
* [PatchHrisBankaccount](#patchhrisbankaccount) - Update a bankaccount
* [PatchHrisBenefit](#patchhrisbenefit) - Update a benefit
* [PatchHrisCompany](#patchhriscompany) - Update a company
* [PatchHrisDeduction](#patchhrisdeduction) - Update a deduction
* [PatchHrisDevice](#patchhrisdevice) - Update a device
* [PatchHrisDocument](#patchhrisdocument) - Update a document
* [PatchHrisEarning](#patchhrisearning) - Update an earning
* [PatchHrisEmployee](#patchhrisemployee) - Update an employee
* [PatchHrisGroup](#patchhrisgroup) - Update a group
* [PatchHrisJob](#patchhrisjob) - Update a job
* [PatchHrisLocation](#patchhrislocation) - Update a location
* [PatchHrisTimeoff](#patchhristimeoff) - Update a timeoff
* [PatchHrisTimeshift](#patchhristimeshift) - Update a timeshift
* [RemoveHrisAttendance](#removehrisattendance) - Remove an attendance
* [RemoveHrisBankaccount](#removehrisbankaccount) - Remove a bankaccount
* [RemoveHrisBenefit](#removehrisbenefit) - Remove a benefit
* [RemoveHrisCompany](#removehriscompany) - Remove a company
* [RemoveHrisDeduction](#removehrisdeduction) - Remove a deduction
* [RemoveHrisDevice](#removehrisdevice) - Remove a device
* [RemoveHrisDocument](#removehrisdocument) - Remove a document
* [RemoveHrisEarning](#removehrisearning) - Remove an earning
* [RemoveHrisEmployee](#removehrisemployee) - Remove an employee
* [RemoveHrisGroup](#removehrisgroup) - Remove a group
* [RemoveHrisJob](#removehrisjob) - Remove a job
* [RemoveHrisLocation](#removehrislocation) - Remove a location
* [RemoveHrisTimeoff](#removehristimeoff) - Remove a timeoff
* [RemoveHrisTimeshift](#removehristimeshift) - Remove a timeshift
* [UpdateHrisAttendance](#updatehrisattendance) - Update an attendance
* [UpdateHrisBankaccount](#updatehrisbankaccount) - Update a bankaccount
* [UpdateHrisBenefit](#updatehrisbenefit) - Update a benefit
* [UpdateHrisCompany](#updatehriscompany) - Update a company
* [UpdateHrisDeduction](#updatehrisdeduction) - Update a deduction
* [UpdateHrisDevice](#updatehrisdevice) - Update a device
* [UpdateHrisDocument](#updatehrisdocument) - Update a document
* [UpdateHrisEarning](#updatehrisearning) - Update an earning
* [UpdateHrisEmployee](#updatehrisemployee) - Update an employee
* [UpdateHrisGroup](#updatehrisgroup) - Update a group
* [UpdateHrisJob](#updatehrisjob) - Update a job
* [UpdateHrisLocation](#updatehrislocation) - Update a location
* [UpdateHrisTimeoff](#updatehristimeoff) - Update a timeoff
* [UpdateHrisTimeshift](#updatehristimeshift) - Update a timeshift

## CreateHrisAttendance

Create an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisAttendance" method="post" path="/hris/{connection_id}/attendance" example="hris_attendance" -->
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

    res, err := s.Hris.CreateHrisAttendance(ctx, operations.CreateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T11:25:58.839Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-11-02T07:25:15.837Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-26T09:41:41.516Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-18T23:51:59.281Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("769ce4c1-27d0-4585-8eb5-30d157525f24"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-10T15:13:04.626Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-19T04:00:08.145Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateHrisAttendanceRequest](../../pkg/models/operations/createhrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateHrisAttendanceResponse](../../pkg/models/operations/createhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisBankaccount

Create a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBankaccount" method="post" path="/hris/{connection_id}/bankaccount" example="hris_bankaccount" -->
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

    res, err := s.Hris.CreateHrisBankaccount(ctx, operations.CreateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("fd554d0c-f06f-4f21-9db0-d8f822ef1dad"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            Priority: unifiedgosdk.Pointer[float64](3.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](23.0),
            SplitType: shared.SplitTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-24T12:07:54.523Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateHrisBankaccountRequest](../../pkg/models/operations/createhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateHrisBankaccountResponse](../../pkg/models/operations/createhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisBenefit

Create a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBenefit" method="post" path="/hris/{connection_id}/benefit" example="hris_benefit" -->
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

    res, err := s.Hris.CreateHrisBenefit(ctx, operations.CreateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("64edbbe5-bec2-4fd5-851d-d306a33d0147"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-17T05:12:07.142Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisBenefitRequest](../../pkg/models/operations/createhrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisBenefitResponse](../../pkg/models/operations/createhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisCompany

Create a company

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisCompany" method="post" path="/hris/{connection_id}/company" example="hris_company" -->
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

    res, err := s.Hris.CreateHrisCompany(ctx, operations.CreateHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("0d43f2b8-9866-4f91-bde9-3113a5c99e0e"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-30T09:43:50.889Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisCompanyRequest](../../pkg/models/operations/createhriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisCompanyResponse](../../pkg/models/operations/createhriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDeduction

Create a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDeduction" method="post" path="/hris/{connection_id}/deduction" example="hris_deduction" -->
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

    res, err := s.Hris.CreateHrisDeduction(ctx, operations.CreateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-15T10:08:18.557Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("8203f6a9-4ad8-409d-855a-3e4d128220b9"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-08T23:29:54.628Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-17T04:01:19.048Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisDeductionRequest](../../pkg/models/operations/createhrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisDeductionResponse](../../pkg/models/operations/createhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDevice

Create a device

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDevice" method="post" path="/hris/{connection_id}/device" example="hris_device" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisDevice(ctx, operations.CreateHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("bb5f11eb-138b-4465-ba05-2d1987082665"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-03T16:42:52.438Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateHrisDeviceRequest](../../pkg/models/operations/createhrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateHrisDeviceResponse](../../pkg/models/operations/createhrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDocument

Create a document

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDocument" method="post" path="/hris/{connection_id}/document" example="hris_document" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisDocument(ctx, operations.CreateHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("698ad094-deb5-43df-aa4f-c96e8bbe29c6"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-05T09:40:12.143Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisDocumentRequest](../../pkg/models/operations/createhrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisDocumentResponse](../../pkg/models/operations/createhrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisEarning

Create an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisEarning" method="post" path="/hris/{connection_id}/earning" example="hris_earning" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisEarning(ctx, operations.CreateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("PKR"),
            Frequency: shared.HrisEarningFrequencyBiweekly.ToPointer(),
            ID: unifiedgosdk.Pointer("4623ea3b-83ef-4751-9b0f-19e40d66239f"),
            Notes: unifiedgosdk.Pointer("Truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-21T17:30:55.265Z"),
            Status: shared.HrisEarningStatusProcessed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-07T17:46:27.510Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisEarningRequest](../../pkg/models/operations/createhrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisEarningResponse](../../pkg/models/operations/createhrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Hris.CreateHrisEmployee(ctx, operations.CreateHrisEmployeeRequest{
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
            HiredAt: types.MustNewTimeFromString("2023-05-23T12:34:35.369Z"),
            ID: unifiedgosdk.Pointer("2e32136f-ca1e-43b8-b4f2-f058401f01f3"),
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
                    ID: unifiedgosdk.Pointer("e1e75dab-d9aa-4b9f-86b9-9a95391748e2"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-02-27T20:16:13.397Z"),
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

## CreateHrisGroup

Create a group

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisGroup" method="post" path="/hris/{connection_id}/group" example="hris_group" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisGroup(ctx, operations.CreateHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("db0018e3-a3b3-49e9-8c5d-9d7b73a83a10"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T19:35:48.923Z"),
        },
        ConnectionID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateHrisGroupRequest](../../pkg/models/operations/createhrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateHrisGroupResponse](../../pkg/models/operations/createhrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisJob

Create a job

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisJob" method="post" path="/hris/{connection_id}/job" example="hris_job" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisJob(ctx, operations.CreateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("55bc90a5-1fe2-475f-84fd-b067394047b7"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T01:52:47.134Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateHrisJobRequest](../../pkg/models/operations/createhrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateHrisJobResponse](../../pkg/models/operations/createhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisLocation

Create a location

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisLocation" method="post" path="/hris/{connection_id}/location" example="hris_location" -->
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

    res, err := s.Hris.CreateHrisLocation(ctx, operations.CreateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("8c0652db-e4ef-4abf-b7e6-30921ce6ee90"),
            ID: unifiedgosdk.Pointer("1bcadcf6-a720-49a2-8060-de2c3f503d2b"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-18T01:19:45.380Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisLocationRequest](../../pkg/models/operations/createhrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisLocationResponse](../../pkg/models/operations/createhrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTaxonomy

Create a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTaxonomy" method="post" path="/hris/{connection_id}/taxonomy" example="hris_taxonomy" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisTaxonomy(ctx, operations.CreateHrisTaxonomyRequest{
        HrisTaxonomy: shared.HrisTaxonomy{
            CreatedAt: types.MustNewTimeFromString("2022-06-23T02:10:00.789Z"),
            Description: unifiedgosdk.Pointer("Apto demonstro audacia adstringo cursim tristis solio careo."),
            Domain: unifiedgosdk.Pointer("Electronics"),
            ID: unifiedgosdk.Pointer("ede085db-5709-4d53-a490-746f3de5be17"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("International Functionality Architect"),
            ParentID: unifiedgosdk.Pointer("6524b2a7-6520-4e15-8c4e-1aa6793db837"),
            RoleIds: []string{
                "2b1ef757-eb4c-4207-8af1-929afe49cd65",
            },
            Subcategory: unifiedgosdk.Pointer("Bamboo"),
            Type: shared.HrisTaxonomyTypeKnowledge.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-27T12:24:05.647Z"),
            URL: unifiedgosdk.Pointer("https://our-polarisation.name"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisTaxonomyRequest](../../pkg/models/operations/createhristaxonomyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisTaxonomyResponse](../../pkg/models/operations/createhristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTimeoff

Create a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeoff" method="post" path="/hris/{connection_id}/timeoff" example="hris_timeoff" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisTimeoff(ctx, operations.CreateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T17:58:39.588Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-23T09:14:50.175Z"),
            ID: unifiedgosdk.Pointer("524dcb2e-6a65-4824-af67-a0e20e9f38a4"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-01T15:35:03.613Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-11T13:23:16.842Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisTimeoffRequest](../../pkg/models/operations/createhristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisTimeoffResponse](../../pkg/models/operations/createhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTimeshift

Create a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeshift" method="post" path="/hris/{connection_id}/timeshift" example="hris_timeshift" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.CreateHrisTimeshift(ctx, operations.CreateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-19T00:56:48.953Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyWeek.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-18T19:21:22.235Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("c9aea5e9-14c3-4ed4-bc41-e543ea8fab91"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-07-08T10:13:38.989Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-29T15:07:27.023Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisTimeshiftRequest](../../pkg/models/operations/createhristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisTimeshiftResponse](../../pkg/models/operations/createhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisAttendance

Retrieve an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisAttendance" method="get" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Hris.GetHrisAttendance(ctx, operations.GetHrisAttendanceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHrisAttendanceRequest](../../pkg/models/operations/gethrisattendancerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetHrisAttendanceResponse](../../pkg/models/operations/gethrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBankaccount

Retrieve a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBankaccount" method="get" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Hris.GetHrisBankaccount(ctx, operations.GetHrisBankaccountRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetHrisBankaccountRequest](../../pkg/models/operations/gethrisbankaccountrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetHrisBankaccountResponse](../../pkg/models/operations/gethrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBenefit

Retrieve a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBenefit" method="get" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Hris.GetHrisBenefit(ctx, operations.GetHrisBenefitRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisBenefitRequest](../../pkg/models/operations/gethrisbenefitrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisBenefitResponse](../../pkg/models/operations/gethrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisCompany

Retrieve a company

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisCompany" method="get" path="/hris/{connection_id}/company/{id}" -->
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

    res, err := s.Hris.GetHrisCompany(ctx, operations.GetHrisCompanyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisCompanyRequest](../../pkg/models/operations/gethriscompanyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisCompanyResponse](../../pkg/models/operations/gethriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDeduction

Retrieve a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDeduction" method="get" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Hris.GetHrisDeduction(ctx, operations.GetHrisDeductionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisDeductionRequest](../../pkg/models/operations/gethrisdeductionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisDeductionResponse](../../pkg/models/operations/gethrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDevice

Retrieve a device

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDevice" method="get" path="/hris/{connection_id}/device/{id}" -->
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

    res, err := s.Hris.GetHrisDevice(ctx, operations.GetHrisDeviceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetHrisDeviceRequest](../../pkg/models/operations/gethrisdevicerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetHrisDeviceResponse](../../pkg/models/operations/gethrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDocument

Retrieve a document

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDocument" method="get" path="/hris/{connection_id}/document/{id}" -->
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

    res, err := s.Hris.GetHrisDocument(ctx, operations.GetHrisDocumentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisDocumentRequest](../../pkg/models/operations/gethrisdocumentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisDocumentResponse](../../pkg/models/operations/gethrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisEarning

Retrieve an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisEarning" method="get" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Hris.GetHrisEarning(ctx, operations.GetHrisEarningRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisEarningRequest](../../pkg/models/operations/gethrisearningrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisEarningResponse](../../pkg/models/operations/gethrisearningresponse.md), error**

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

    res, err := s.Hris.GetHrisEmployee(ctx, operations.GetHrisEmployeeRequest{
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

## GetHrisGroup

Retrieve a group

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisGroup" method="get" path="/hris/{connection_id}/group/{id}" -->
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

    res, err := s.Hris.GetHrisGroup(ctx, operations.GetHrisGroupRequest{
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetHrisGroupRequest](../../pkg/models/operations/gethrisgrouprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetHrisGroupResponse](../../pkg/models/operations/gethrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisJob

Retrieve a job

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisJob" method="get" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Hris.GetHrisJob(ctx, operations.GetHrisJobRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetHrisJobRequest](../../pkg/models/operations/gethrisjobrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetHrisJobResponse](../../pkg/models/operations/gethrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisLocation" method="get" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Hris.GetHrisLocation(ctx, operations.GetHrisLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisLocationRequest](../../pkg/models/operations/gethrislocationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisLocationResponse](../../pkg/models/operations/gethrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPaycode

Retrieve a paycode

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPaycode" method="get" path="/hris/{connection_id}/paycode/{id}" -->
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

    res, err := s.Hris.GetHrisPaycode(ctx, operations.GetHrisPaycodeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPaycode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPaycodeRequest](../../pkg/models/operations/gethrispaycoderequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPaycodeResponse](../../pkg/models/operations/gethrispaycoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPayroll

Retrieve a payroll

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPayroll" method="get" path="/hris/{connection_id}/payroll/{id}" -->
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

    res, err := s.Hris.GetHrisPayroll(ctx, operations.GetHrisPayrollRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayroll != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPayrollRequest](../../pkg/models/operations/gethrispayrollrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPayrollResponse](../../pkg/models/operations/gethrispayrollresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPayslip

Retrieve a payslip

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPayslip" method="get" path="/hris/{connection_id}/payslip/{id}" -->
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

    res, err := s.Hris.GetHrisPayslip(ctx, operations.GetHrisPayslipRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslip != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPayslipRequest](../../pkg/models/operations/gethrispaysliprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPayslipResponse](../../pkg/models/operations/gethrispayslipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTaxonomy

Retrieve a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTaxonomy" method="get" path="/hris/{connection_id}/taxonomy/{id}" -->
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

    res, err := s.Hris.GetHrisTaxonomy(ctx, operations.GetHrisTaxonomyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisTaxonomyRequest](../../pkg/models/operations/gethristaxonomyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisTaxonomyResponse](../../pkg/models/operations/gethristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeoff

Retrieve a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeoff" method="get" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Hris.GetHrisTimeoff(ctx, operations.GetHrisTimeoffRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisTimeoffRequest](../../pkg/models/operations/gethristimeoffrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisTimeoffResponse](../../pkg/models/operations/gethristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeshift

Retrieve a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeshift" method="get" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Hris.GetHrisTimeshift(ctx, operations.GetHrisTimeshiftRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisTimeshiftRequest](../../pkg/models/operations/gethristimeshiftrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisTimeshiftResponse](../../pkg/models/operations/gethristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisAttendances

List all attendances

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisAttendances" method="get" path="/hris/{connection_id}/attendance" -->
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

    res, err := s.Hris.ListHrisAttendances(ctx, operations.ListHrisAttendancesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListHrisAttendancesRequest](../../pkg/models/operations/listhrisattendancesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListHrisAttendancesResponse](../../pkg/models/operations/listhrisattendancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBankaccounts

List all bankaccounts

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBankaccounts" method="get" path="/hris/{connection_id}/bankaccount" -->
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

    res, err := s.Hris.ListHrisBankaccounts(ctx, operations.ListHrisBankaccountsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListHrisBankaccountsRequest](../../pkg/models/operations/listhrisbankaccountsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListHrisBankaccountsResponse](../../pkg/models/operations/listhrisbankaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBenefits

List all benefits

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBenefits" method="get" path="/hris/{connection_id}/benefit" -->
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

    res, err := s.Hris.ListHrisBenefits(ctx, operations.ListHrisBenefitsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisBenefitsRequest](../../pkg/models/operations/listhrisbenefitsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisBenefitsResponse](../../pkg/models/operations/listhrisbenefitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisCompanies

List all companies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisCompanies" method="get" path="/hris/{connection_id}/company" -->
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

    res, err := s.Hris.ListHrisCompanies(ctx, operations.ListHrisCompaniesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisCompaniesRequest](../../pkg/models/operations/listhriscompaniesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisCompaniesResponse](../../pkg/models/operations/listhriscompaniesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDeductions

List all deductions

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDeductions" method="get" path="/hris/{connection_id}/deduction" -->
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

    res, err := s.Hris.ListHrisDeductions(ctx, operations.ListHrisDeductionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeductions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisDeductionsRequest](../../pkg/models/operations/listhrisdeductionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisDeductionsResponse](../../pkg/models/operations/listhrisdeductionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDevices

List all devices

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDevices" method="get" path="/hris/{connection_id}/device" -->
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

    res, err := s.Hris.ListHrisDevices(ctx, operations.ListHrisDevicesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListHrisDevicesRequest](../../pkg/models/operations/listhrisdevicesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListHrisDevicesResponse](../../pkg/models/operations/listhrisdevicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDocuments

List all documents

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDocuments" method="get" path="/hris/{connection_id}/document" -->
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

    res, err := s.Hris.ListHrisDocuments(ctx, operations.ListHrisDocumentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocuments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisDocumentsRequest](../../pkg/models/operations/listhrisdocumentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisDocumentsResponse](../../pkg/models/operations/listhrisdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisEarnings

List all earnings

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisEarnings" method="get" path="/hris/{connection_id}/earning" -->
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

    res, err := s.Hris.ListHrisEarnings(ctx, operations.ListHrisEarningsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarnings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisEarningsRequest](../../pkg/models/operations/listhrisearningsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisEarningsResponse](../../pkg/models/operations/listhrisearningsresponse.md), error**

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

    res, err := s.Hris.ListHrisEmployees(ctx, operations.ListHrisEmployeesRequest{
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

## ListHrisGroups

List all groups

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisGroups" method="get" path="/hris/{connection_id}/group" -->
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

    res, err := s.Hris.ListHrisGroups(ctx, operations.ListHrisGroupsRequest{
        ConnectionID: "<id>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListHrisGroupsRequest](../../pkg/models/operations/listhrisgroupsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListHrisGroupsResponse](../../pkg/models/operations/listhrisgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisJobs

List all jobs

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisJobs" method="get" path="/hris/{connection_id}/job" -->
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

    res, err := s.Hris.ListHrisJobs(ctx, operations.ListHrisJobsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListHrisJobsRequest](../../pkg/models/operations/listhrisjobsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListHrisJobsResponse](../../pkg/models/operations/listhrisjobsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisLocations" method="get" path="/hris/{connection_id}/location" -->
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

    res, err := s.Hris.ListHrisLocations(ctx, operations.ListHrisLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisLocationsRequest](../../pkg/models/operations/listhrislocationsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisLocationsResponse](../../pkg/models/operations/listhrislocationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPaycodes

List all paycodes

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPaycodes" method="get" path="/hris/{connection_id}/paycode" -->
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

    res, err := s.Hris.ListHrisPaycodes(ctx, operations.ListHrisPaycodesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPaycodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPaycodesRequest](../../pkg/models/operations/listhrispaycodesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPaycodesResponse](../../pkg/models/operations/listhrispaycodesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPayrolls

List all payrolls

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPayrolls" method="get" path="/hris/{connection_id}/payroll" -->
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

    res, err := s.Hris.ListHrisPayrolls(ctx, operations.ListHrisPayrollsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayrolls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPayrollsRequest](../../pkg/models/operations/listhrispayrollsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPayrollsResponse](../../pkg/models/operations/listhrispayrollsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPayslips

List all payslips

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPayslips" method="get" path="/hris/{connection_id}/payslip" -->
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

    res, err := s.Hris.ListHrisPayslips(ctx, operations.ListHrisPayslipsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslips != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPayslipsRequest](../../pkg/models/operations/listhrispayslipsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPayslipsResponse](../../pkg/models/operations/listhrispayslipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTaxonomies

List all taxonomies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTaxonomies" method="get" path="/hris/{connection_id}/taxonomy" -->
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

    res, err := s.Hris.ListHrisTaxonomies(ctx, operations.ListHrisTaxonomiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTaxonomiesRequest](../../pkg/models/operations/listhristaxonomiesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTaxonomiesResponse](../../pkg/models/operations/listhristaxonomiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeoffs

List all timeoffs

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeoffs" method="get" path="/hris/{connection_id}/timeoff" -->
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

    res, err := s.Hris.ListHrisTimeoffs(ctx, operations.ListHrisTimeoffsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoffs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisTimeoffsRequest](../../pkg/models/operations/listhristimeoffsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisTimeoffsResponse](../../pkg/models/operations/listhristimeoffsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeshifts

List all timeshifts

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeshifts" method="get" path="/hris/{connection_id}/timeshift" -->
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

    res, err := s.Hris.ListHrisTimeshifts(ctx, operations.ListHrisTimeshiftsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshifts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTimeshiftsRequest](../../pkg/models/operations/listhristimeshiftsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTimeshiftsResponse](../../pkg/models/operations/listhristimeshiftsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisAttendance

Update an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisAttendance" method="patch" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Hris.PatchHrisAttendance(ctx, operations.PatchHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T11:25:58.839Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-11-02T07:25:15.849Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-26T09:41:41.528Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-18T23:51:59.295Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("721b5c9e-8462-4ea8-bf98-888ef9a317d6"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-10T15:13:04.628Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-19T04:00:08.147Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchHrisAttendanceRequest](../../pkg/models/operations/patchhrisattendancerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchHrisAttendanceResponse](../../pkg/models/operations/patchhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBankaccount

Update a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBankaccount" method="patch" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Hris.PatchHrisBankaccount(ctx, operations.PatchHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("f4113c97-5bc1-4f82-a9c9-332cf3efb69c"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            Priority: unifiedgosdk.Pointer[float64](3.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](23.0),
            SplitType: shared.SplitTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-24T12:07:54.535Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchHrisBankaccountRequest](../../pkg/models/operations/patchhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchHrisBankaccountResponse](../../pkg/models/operations/patchhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBenefit

Update a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBenefit" method="patch" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Hris.PatchHrisBenefit(ctx, operations.PatchHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("f0a1631a-ba83-4473-9893-f9bab90a70de"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-17T05:12:07.152Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisBenefitRequest](../../pkg/models/operations/patchhrisbenefitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisBenefitResponse](../../pkg/models/operations/patchhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisCompany" method="patch" path="/hris/{connection_id}/company/{id}" example="hris_company" -->
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

    res, err := s.Hris.PatchHrisCompany(ctx, operations.PatchHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("4c0a0aa3-7aec-4f96-8b39-5f5039eb1ab1"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-30T09:43:50.904Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisCompanyRequest](../../pkg/models/operations/patchhriscompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisCompanyResponse](../../pkg/models/operations/patchhriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDeduction" method="patch" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Hris.PatchHrisDeduction(ctx, operations.PatchHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-15T10:08:18.575Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("5b8ba2a5-5f77-4bd9-a159-37fcdbc12736"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-08T23:29:54.643Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-17T04:01:19.061Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisDeductionRequest](../../pkg/models/operations/patchhrisdeductionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisDeductionResponse](../../pkg/models/operations/patchhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDevice

Update a device

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDevice" method="patch" path="/hris/{connection_id}/device/{id}" example="hris_device" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisDevice(ctx, operations.PatchHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("023adc95-9d8b-4702-b9ac-4cbddb8396a8"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-03T16:42:52.446Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchHrisDeviceRequest](../../pkg/models/operations/patchhrisdevicerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchHrisDeviceResponse](../../pkg/models/operations/patchhrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDocument" method="patch" path="/hris/{connection_id}/document/{id}" example="hris_document" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisDocument(ctx, operations.PatchHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("a37fa394-b40c-461d-b58b-f241c83c1d71"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-05T09:40:12.151Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisDocumentRequest](../../pkg/models/operations/patchhrisdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisDocumentResponse](../../pkg/models/operations/patchhrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisEarning

Update an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisEarning" method="patch" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisEarning(ctx, operations.PatchHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("PKR"),
            Frequency: shared.HrisEarningFrequencyBiweekly.ToPointer(),
            ID: unifiedgosdk.Pointer("bebe286e-07ab-4a30-833e-d57bebc02130"),
            Notes: unifiedgosdk.Pointer("Truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-21T17:30:55.271Z"),
            Status: shared.HrisEarningStatusProcessed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-07T17:46:27.512Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisEarningRequest](../../pkg/models/operations/patchhrisearningrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisEarningResponse](../../pkg/models/operations/patchhrisearningresponse.md), error**

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

    res, err := s.Hris.PatchHrisEmployee(ctx, operations.PatchHrisEmployeeRequest{
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
            HiredAt: types.MustNewTimeFromString("2023-05-23T12:34:35.425Z"),
            ID: unifiedgosdk.Pointer("d1fbd084-cb37-42cf-b71d-7284b1e43b81"),
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
                    ID: unifiedgosdk.Pointer("0f7dd336-708c-400e-9dd5-2f60af834985"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-02-27T20:16:13.435Z"),
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

## PatchHrisGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisGroup" method="patch" path="/hris/{connection_id}/group/{id}" example="hris_group" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisGroup(ctx, operations.PatchHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("2990d098-f7a6-4a65-a1a1-667094a42001"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T19:35:48.934Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchHrisGroupRequest](../../pkg/models/operations/patchhrisgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchHrisGroupResponse](../../pkg/models/operations/patchhrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisJob

Update a job

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisJob" method="patch" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisJob(ctx, operations.PatchHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("9d2649c6-530c-427e-8c1c-3326ca3d1346"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T01:52:47.190Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchHrisJobRequest](../../pkg/models/operations/patchhrisjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchHrisJobResponse](../../pkg/models/operations/patchhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisLocation" method="patch" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Hris.PatchHrisLocation(ctx, operations.PatchHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("3a6676d7-494d-454e-8d80-b4a4f16da150"),
            ID: unifiedgosdk.Pointer("4457b4b2-35c6-411b-8945-d59c9e99260c"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-18T01:19:45.389Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisLocationRequest](../../pkg/models/operations/patchhrislocationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisLocationResponse](../../pkg/models/operations/patchhrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeoff" method="patch" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisTimeoff(ctx, operations.PatchHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T17:58:39.589Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-23T09:14:50.185Z"),
            ID: unifiedgosdk.Pointer("9b3d4ae1-a0d1-48c4-9b44-087963b7a041"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-01T15:35:03.619Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-11T13:23:16.845Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisTimeoffRequest](../../pkg/models/operations/patchhristimeoffrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisTimeoffResponse](../../pkg/models/operations/patchhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeshift

Update a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeshift" method="patch" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.PatchHrisTimeshift(ctx, operations.PatchHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-19T00:56:48.963Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyWeek.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-18T19:21:22.251Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("67e7a295-d916-4f9b-a4c1-7982136c37ce"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-07-08T10:13:38.999Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-29T15:07:27.027Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisTimeshiftRequest](../../pkg/models/operations/patchhristimeshiftrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisTimeshiftResponse](../../pkg/models/operations/patchhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisAttendance

Remove an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisAttendance" method="delete" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Hris.RemoveHrisAttendance(ctx, operations.RemoveHrisAttendanceRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveHrisAttendanceRequest](../../pkg/models/operations/removehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveHrisAttendanceResponse](../../pkg/models/operations/removehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBankaccount

Remove a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBankaccount" method="delete" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Hris.RemoveHrisBankaccount(ctx, operations.RemoveHrisBankaccountRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveHrisBankaccountRequest](../../pkg/models/operations/removehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveHrisBankaccountResponse](../../pkg/models/operations/removehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBenefit

Remove a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBenefit" method="delete" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Hris.RemoveHrisBenefit(ctx, operations.RemoveHrisBenefitRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveHrisBenefitRequest](../../pkg/models/operations/removehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisBenefitResponse](../../pkg/models/operations/removehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisCompany

Remove a company

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisCompany" method="delete" path="/hris/{connection_id}/company/{id}" -->
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

    res, err := s.Hris.RemoveHrisCompany(ctx, operations.RemoveHrisCompanyRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveHrisCompanyRequest](../../pkg/models/operations/removehriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisCompanyResponse](../../pkg/models/operations/removehriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDeduction

Remove a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDeduction" method="delete" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Hris.RemoveHrisDeduction(ctx, operations.RemoveHrisDeductionRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveHrisDeductionRequest](../../pkg/models/operations/removehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisDeductionResponse](../../pkg/models/operations/removehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDevice

Remove a device

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDevice" method="delete" path="/hris/{connection_id}/device/{id}" -->
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

    res, err := s.Hris.RemoveHrisDevice(ctx, operations.RemoveHrisDeviceRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveHrisDeviceRequest](../../pkg/models/operations/removehrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveHrisDeviceResponse](../../pkg/models/operations/removehrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDocument

Remove a document

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDocument" method="delete" path="/hris/{connection_id}/document/{id}" -->
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

    res, err := s.Hris.RemoveHrisDocument(ctx, operations.RemoveHrisDocumentRequest{
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
| `request`                                                                                        | [operations.RemoveHrisDocumentRequest](../../pkg/models/operations/removehrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisDocumentResponse](../../pkg/models/operations/removehrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisEarning

Remove an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisEarning" method="delete" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Hris.RemoveHrisEarning(ctx, operations.RemoveHrisEarningRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveHrisEarningRequest](../../pkg/models/operations/removehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisEarningResponse](../../pkg/models/operations/removehrisearningresponse.md), error**

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

    res, err := s.Hris.RemoveHrisEmployee(ctx, operations.RemoveHrisEmployeeRequest{
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

## RemoveHrisGroup

Remove a group

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisGroup" method="delete" path="/hris/{connection_id}/group/{id}" -->
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

    res, err := s.Hris.RemoveHrisGroup(ctx, operations.RemoveHrisGroupRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveHrisGroupRequest](../../pkg/models/operations/removehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveHrisGroupResponse](../../pkg/models/operations/removehrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisJob

Remove a job

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisJob" method="delete" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Hris.RemoveHrisJob(ctx, operations.RemoveHrisJobRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveHrisJobRequest](../../pkg/models/operations/removehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveHrisJobResponse](../../pkg/models/operations/removehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisLocation

Remove a location

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisLocation" method="delete" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Hris.RemoveHrisLocation(ctx, operations.RemoveHrisLocationRequest{
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
| `request`                                                                                        | [operations.RemoveHrisLocationRequest](../../pkg/models/operations/removehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisLocationResponse](../../pkg/models/operations/removehrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeoff

Remove a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeoff" method="delete" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Hris.RemoveHrisTimeoff(ctx, operations.RemoveHrisTimeoffRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveHrisTimeoffRequest](../../pkg/models/operations/removehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisTimeoffResponse](../../pkg/models/operations/removehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeshift

Remove a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeshift" method="delete" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Hris.RemoveHrisTimeshift(ctx, operations.RemoveHrisTimeshiftRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveHrisTimeshiftRequest](../../pkg/models/operations/removehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisTimeshiftResponse](../../pkg/models/operations/removehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisAttendance

Update an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisAttendance" method="put" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Hris.UpdateHrisAttendance(ctx, operations.UpdateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T11:25:58.839Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-11-02T07:25:15.849Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-26T09:41:41.528Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-18T23:51:59.295Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("721b5c9e-8462-4ea8-bf98-888ef9a317d6"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-10T15:13:04.628Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-19T04:00:08.147Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateHrisAttendanceRequest](../../pkg/models/operations/updatehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateHrisAttendanceResponse](../../pkg/models/operations/updatehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBankaccount

Update a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBankaccount" method="put" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Hris.UpdateHrisBankaccount(ctx, operations.UpdateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("f4113c97-5bc1-4f82-a9c9-332cf3efb69c"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            Priority: unifiedgosdk.Pointer[float64](3.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](23.0),
            SplitType: shared.SplitTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-06-24T12:07:54.535Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateHrisBankaccountRequest](../../pkg/models/operations/updatehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateHrisBankaccountResponse](../../pkg/models/operations/updatehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBenefit

Update a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBenefit" method="put" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Hris.UpdateHrisBenefit(ctx, operations.UpdateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("f0a1631a-ba83-4473-9893-f9bab90a70de"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-17T05:12:07.152Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisBenefitRequest](../../pkg/models/operations/updatehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisBenefitResponse](../../pkg/models/operations/updatehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisCompany" method="put" path="/hris/{connection_id}/company/{id}" example="hris_company" -->
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

    res, err := s.Hris.UpdateHrisCompany(ctx, operations.UpdateHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("4c0a0aa3-7aec-4f96-8b39-5f5039eb1ab1"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-30T09:43:50.904Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisCompanyRequest](../../pkg/models/operations/updatehriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisCompanyResponse](../../pkg/models/operations/updatehriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDeduction" method="put" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Hris.UpdateHrisDeduction(ctx, operations.UpdateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](22611.0),
            EmployerType: shared.EmployerTypeFixed.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-06-15T10:08:18.575Z"),
            Frequency: shared.HrisDeductionFrequencyBiweekly.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeOther.ToPointer(),
            ID: unifiedgosdk.Pointer("5b8ba2a5-5f77-4bd9-a159-37fcdbc12736"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-08T23:29:54.643Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-17T04:01:19.061Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisDeductionRequest](../../pkg/models/operations/updatehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisDeductionResponse](../../pkg/models/operations/updatehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDevice

Update a device

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDevice" method="put" path="/hris/{connection_id}/device/{id}" example="hris_device" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisDevice(ctx, operations.UpdateHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("023adc95-9d8b-4702-b9ac-4cbddb8396a8"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-03T16:42:52.446Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateHrisDeviceRequest](../../pkg/models/operations/updatehrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateHrisDeviceResponse](../../pkg/models/operations/updatehrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDocument" method="put" path="/hris/{connection_id}/document/{id}" example="hris_document" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisDocument(ctx, operations.UpdateHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("a37fa394-b40c-461d-b58b-f241c83c1d71"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-05T09:40:12.151Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisDocumentRequest](../../pkg/models/operations/updatehrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisDocumentResponse](../../pkg/models/operations/updatehrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisEarning

Update an earning

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisEarning" method="put" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisEarning(ctx, operations.UpdateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("PKR"),
            Frequency: shared.HrisEarningFrequencyBiweekly.ToPointer(),
            ID: unifiedgosdk.Pointer("bebe286e-07ab-4a30-833e-d57bebc02130"),
            Notes: unifiedgosdk.Pointer("Truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-21T17:30:55.271Z"),
            Status: shared.HrisEarningStatusProcessed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-07T17:46:27.512Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisEarningRequest](../../pkg/models/operations/updatehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisEarningResponse](../../pkg/models/operations/updatehrisearningresponse.md), error**

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

    res, err := s.Hris.UpdateHrisEmployee(ctx, operations.UpdateHrisEmployeeRequest{
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
            HiredAt: types.MustNewTimeFromString("2023-05-23T12:34:35.425Z"),
            ID: unifiedgosdk.Pointer("d1fbd084-cb37-42cf-b71d-7284b1e43b81"),
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
                    ID: unifiedgosdk.Pointer("0f7dd336-708c-400e-9dd5-2f60af834985"),
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
            UpdatedAt: types.MustNewTimeFromString("2022-02-27T20:16:13.435Z"),
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

## UpdateHrisGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisGroup" method="put" path="/hris/{connection_id}/group/{id}" example="hris_group" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisGroup(ctx, operations.UpdateHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("2990d098-f7a6-4a65-a1a1-667094a42001"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-14T19:35:48.934Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateHrisGroupRequest](../../pkg/models/operations/updatehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateHrisGroupResponse](../../pkg/models/operations/updatehrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisJob

Update a job

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisJob" method="put" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisJob(ctx, operations.UpdateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("9d2649c6-530c-427e-8c1c-3326ca3d1346"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T01:52:47.190Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateHrisJobRequest](../../pkg/models/operations/updatehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateHrisJobResponse](../../pkg/models/operations/updatehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisLocation" method="put" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Hris.UpdateHrisLocation(ctx, operations.UpdateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("3a6676d7-494d-454e-8d80-b4a4f16da150"),
            ID: unifiedgosdk.Pointer("4457b4b2-35c6-411b-8945-d59c9e99260c"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-18T01:19:45.389Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisLocationRequest](../../pkg/models/operations/updatehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisLocationResponse](../../pkg/models/operations/updatehrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeoff" method="put" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisTimeoff(ctx, operations.UpdateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-22T17:58:39.589Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-23T09:14:50.185Z"),
            ID: unifiedgosdk.Pointer("9b3d4ae1-a0d1-48c4-9b44-087963b7a041"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-01T15:35:03.619Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-11T13:23:16.845Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisTimeoffRequest](../../pkg/models/operations/updatehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisTimeoffResponse](../../pkg/models/operations/updatehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeshift

Update a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeshift" method="put" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Hris.UpdateHrisTimeshift(ctx, operations.UpdateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-19T00:56:48.963Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyWeek.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-09-18T19:21:22.251Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("67e7a295-d916-4f9b-a4c1-7982136c37ce"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-07-08T10:13:38.999Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-29T15:07:27.027Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisTimeshiftRequest](../../pkg/models/operations/updatehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisTimeshiftResponse](../../pkg/models/operations/updatehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |