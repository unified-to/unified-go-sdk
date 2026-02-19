# KmsPageMetadataExtraData


## Supported Types

### 

```go
kmsPageMetadataExtraData := shared.CreateKmsPageMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
kmsPageMetadataExtraData := shared.CreateKmsPageMetadataExtraDataStr(string{/* values here */})
```

### 

```go
kmsPageMetadataExtraData := shared.CreateKmsPageMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
kmsPageMetadataExtraData := shared.CreateKmsPageMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
kmsPageMetadataExtraData := shared.CreateKmsPageMetadataExtraDataArrayOfKmsPageMetadata5([]shared.KmsPageMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kmsPageMetadataExtraData.Type {
	case shared.KmsPageMetadataExtraDataTypeMapOfAny:
		// kmsPageMetadataExtraData.MapOfAny is populated
	case shared.KmsPageMetadataExtraDataTypeStr:
		// kmsPageMetadataExtraData.Str is populated
	case shared.KmsPageMetadataExtraDataTypeNumber:
		// kmsPageMetadataExtraData.Number is populated
	case shared.KmsPageMetadataExtraDataTypeBoolean:
		// kmsPageMetadataExtraData.Boolean is populated
	case shared.KmsPageMetadataExtraDataTypeArrayOfKmsPageMetadata5:
		// kmsPageMetadataExtraData.ArrayOfKmsPageMetadata5 is populated
}
```
