# CrmMetadataValue


## Supported Types

### 

```go
crmMetadataValue := shared.CreateCrmMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
crmMetadataValue := shared.CreateCrmMetadataValueStr(string{/* values here */})
```

### 

```go
crmMetadataValue := shared.CreateCrmMetadataValueNumber(float64{/* values here */})
```

### 

```go
crmMetadataValue := shared.CreateCrmMetadataValueBoolean(bool{/* values here */})
```

### 

```go
crmMetadataValue := shared.CreateCrmMetadataValueArrayOfCrmMetadataSchemas5([]shared.CrmMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch crmMetadataValue.Type {
	case shared.CrmMetadataValueTypeMapOfAny:
		// crmMetadataValue.MapOfAny is populated
	case shared.CrmMetadataValueTypeStr:
		// crmMetadataValue.Str is populated
	case shared.CrmMetadataValueTypeNumber:
		// crmMetadataValue.Number is populated
	case shared.CrmMetadataValueTypeBoolean:
		// crmMetadataValue.Boolean is populated
	case shared.CrmMetadataValueTypeArrayOfCrmMetadataSchemas5:
		// crmMetadataValue.ArrayOfCrmMetadataSchemas5 is populated
}
```
