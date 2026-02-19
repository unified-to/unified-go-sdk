# CrmMetadataSchemas5


## Supported Types

### CrmMetadataSchemas1

```go
crmMetadataSchemas5 := shared.CreateCrmMetadataSchemas5CrmMetadataSchemas1(shared.CrmMetadataSchemas1{/* values here */})
```

### 

```go
crmMetadataSchemas5 := shared.CreateCrmMetadataSchemas5Str(string{/* values here */})
```

### 

```go
crmMetadataSchemas5 := shared.CreateCrmMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
crmMetadataSchemas5 := shared.CreateCrmMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch crmMetadataSchemas5.Type {
	case shared.CrmMetadataSchemas5TypeCrmMetadataSchemas1:
		// crmMetadataSchemas5.CrmMetadataSchemas1 is populated
	case shared.CrmMetadataSchemas5TypeStr:
		// crmMetadataSchemas5.Str is populated
	case shared.CrmMetadataSchemas5TypeNumber:
		// crmMetadataSchemas5.Number is populated
	case shared.CrmMetadataSchemas5TypeBoolean:
		// crmMetadataSchemas5.Boolean is populated
}
```
