# PropertyMetadataExtraData

additional data about the custom field


## Supported Types

### 

```go
propertyMetadataExtraData := shared.CreatePropertyMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
propertyMetadataExtraData := shared.CreatePropertyMetadataExtraDataStr(string{/* values here */})
```

### 

```go
propertyMetadataExtraData := shared.CreatePropertyMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
propertyMetadataExtraData := shared.CreatePropertyMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
propertyMetadataExtraData := shared.CreatePropertyMetadataExtraDataArrayOfPropertyMetadata5([]shared.PropertyMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch propertyMetadataExtraData.Type {
	case shared.PropertyMetadataExtraDataTypeMapOfAny:
		// propertyMetadataExtraData.MapOfAny is populated
	case shared.PropertyMetadataExtraDataTypeStr:
		// propertyMetadataExtraData.Str is populated
	case shared.PropertyMetadataExtraDataTypeNumber:
		// propertyMetadataExtraData.Number is populated
	case shared.PropertyMetadataExtraDataTypeBoolean:
		// propertyMetadataExtraData.Boolean is populated
	case shared.PropertyMetadataExtraDataTypeArrayOfPropertyMetadata5:
		// propertyMetadataExtraData.ArrayOfPropertyMetadata5 is populated
}
```
