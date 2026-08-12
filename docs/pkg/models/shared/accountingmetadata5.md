# AccountingMetadata5


## Supported Types

### AccountingMetadata1

```go
accountingMetadata5 := shared.CreateAccountingMetadata5AccountingMetadata1(shared.AccountingMetadata1{/* values here */})
```

### 

```go
accountingMetadata5 := shared.CreateAccountingMetadata5Str(string{/* values here */})
```

### 

```go
accountingMetadata5 := shared.CreateAccountingMetadata5Number(float64{/* values here */})
```

### 

```go
accountingMetadata5 := shared.CreateAccountingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch accountingMetadata5.Type {
	case shared.AccountingMetadata5TypeAccountingMetadata1:
		// accountingMetadata5.AccountingMetadata1 is populated
	case shared.AccountingMetadata5TypeStr:
		// accountingMetadata5.Str is populated
	case shared.AccountingMetadata5TypeNumber:
		// accountingMetadata5.Number is populated
	case shared.AccountingMetadata5TypeBoolean:
		// accountingMetadata5.Boolean is populated
}
```
