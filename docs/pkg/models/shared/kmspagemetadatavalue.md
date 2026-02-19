# KmsPageMetadataValue


## Supported Types

### 

```go
kmsPageMetadataValue := shared.CreateKmsPageMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
kmsPageMetadataValue := shared.CreateKmsPageMetadataValueStr(string{/* values here */})
```

### 

```go
kmsPageMetadataValue := shared.CreateKmsPageMetadataValueNumber(float64{/* values here */})
```

### 

```go
kmsPageMetadataValue := shared.CreateKmsPageMetadataValueBoolean(bool{/* values here */})
```

### 

```go
kmsPageMetadataValue := shared.CreateKmsPageMetadataValueArrayOfKmsPageMetadataSchemas5([]shared.KmsPageMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kmsPageMetadataValue.Type {
	case shared.KmsPageMetadataValueTypeMapOfAny:
		// kmsPageMetadataValue.MapOfAny is populated
	case shared.KmsPageMetadataValueTypeStr:
		// kmsPageMetadataValue.Str is populated
	case shared.KmsPageMetadataValueTypeNumber:
		// kmsPageMetadataValue.Number is populated
	case shared.KmsPageMetadataValueTypeBoolean:
		// kmsPageMetadataValue.Boolean is populated
	case shared.KmsPageMetadataValueTypeArrayOfKmsPageMetadataSchemas5:
		// kmsPageMetadataValue.ArrayOfKmsPageMetadataSchemas5 is populated
}
```
