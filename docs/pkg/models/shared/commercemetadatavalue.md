# CommerceMetadataValue


## Supported Types

### 

```go
commerceMetadataValue := shared.CreateCommerceMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
commerceMetadataValue := shared.CreateCommerceMetadataValueStr(string{/* values here */})
```

### 

```go
commerceMetadataValue := shared.CreateCommerceMetadataValueNumber(float64{/* values here */})
```

### 

```go
commerceMetadataValue := shared.CreateCommerceMetadataValueBoolean(bool{/* values here */})
```

### 

```go
commerceMetadataValue := shared.CreateCommerceMetadataValueArrayOfCommerceMetadataSchemas5([]shared.CommerceMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch commerceMetadataValue.Type {
	case shared.CommerceMetadataValueTypeMapOfAny:
		// commerceMetadataValue.MapOfAny is populated
	case shared.CommerceMetadataValueTypeStr:
		// commerceMetadataValue.Str is populated
	case shared.CommerceMetadataValueTypeNumber:
		// commerceMetadataValue.Number is populated
	case shared.CommerceMetadataValueTypeBoolean:
		// commerceMetadataValue.Boolean is populated
	case shared.CommerceMetadataValueTypeArrayOfCommerceMetadataSchemas5:
		// commerceMetadataValue.ArrayOfCommerceMetadataSchemas5 is populated
}
```
