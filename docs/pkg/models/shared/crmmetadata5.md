# CrmMetadata5


## Supported Types

### CrmMetadata1

```go
crmMetadata5 := shared.CreateCrmMetadata5CrmMetadata1(shared.CrmMetadata1{/* values here */})
```

### 

```go
crmMetadata5 := shared.CreateCrmMetadata5Str(string{/* values here */})
```

### 

```go
crmMetadata5 := shared.CreateCrmMetadata5Number(float64{/* values here */})
```

### 

```go
crmMetadata5 := shared.CreateCrmMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch crmMetadata5.Type {
	case shared.CrmMetadata5TypeCrmMetadata1:
		// crmMetadata5.CrmMetadata1 is populated
	case shared.CrmMetadata5TypeStr:
		// crmMetadata5.Str is populated
	case shared.CrmMetadata5TypeNumber:
		// crmMetadata5.Number is populated
	case shared.CrmMetadata5TypeBoolean:
		// crmMetadata5.Boolean is populated
}
```
