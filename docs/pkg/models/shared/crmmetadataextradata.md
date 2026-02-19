# CrmMetadataExtraData


## Supported Types

### 

```go
crmMetadataExtraData := shared.CreateCrmMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
crmMetadataExtraData := shared.CreateCrmMetadataExtraDataStr(string{/* values here */})
```

### 

```go
crmMetadataExtraData := shared.CreateCrmMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
crmMetadataExtraData := shared.CreateCrmMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
crmMetadataExtraData := shared.CreateCrmMetadataExtraDataArrayOfCrmMetadata5([]shared.CrmMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch crmMetadataExtraData.Type {
	case shared.CrmMetadataExtraDataTypeMapOfAny:
		// crmMetadataExtraData.MapOfAny is populated
	case shared.CrmMetadataExtraDataTypeStr:
		// crmMetadataExtraData.Str is populated
	case shared.CrmMetadataExtraDataTypeNumber:
		// crmMetadataExtraData.Number is populated
	case shared.CrmMetadataExtraDataTypeBoolean:
		// crmMetadataExtraData.Boolean is populated
	case shared.CrmMetadataExtraDataTypeArrayOfCrmMetadata5:
		// crmMetadataExtraData.ArrayOfCrmMetadata5 is populated
}
```
