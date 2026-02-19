# CommerceMetadataExtraData


## Supported Types

### 

```go
commerceMetadataExtraData := shared.CreateCommerceMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
commerceMetadataExtraData := shared.CreateCommerceMetadataExtraDataStr(string{/* values here */})
```

### 

```go
commerceMetadataExtraData := shared.CreateCommerceMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
commerceMetadataExtraData := shared.CreateCommerceMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
commerceMetadataExtraData := shared.CreateCommerceMetadataExtraDataArrayOfCommerceMetadata5([]shared.CommerceMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch commerceMetadataExtraData.Type {
	case shared.CommerceMetadataExtraDataTypeMapOfAny:
		// commerceMetadataExtraData.MapOfAny is populated
	case shared.CommerceMetadataExtraDataTypeStr:
		// commerceMetadataExtraData.Str is populated
	case shared.CommerceMetadataExtraDataTypeNumber:
		// commerceMetadataExtraData.Number is populated
	case shared.CommerceMetadataExtraDataTypeBoolean:
		// commerceMetadataExtraData.Boolean is populated
	case shared.CommerceMetadataExtraDataTypeArrayOfCommerceMetadata5:
		// commerceMetadataExtraData.ArrayOfCommerceMetadata5 is populated
}
```
