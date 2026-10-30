# PropertyMetadataValue

the custom field value


## Supported Types

### 

```go
propertyMetadataValue := shared.CreatePropertyMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
propertyMetadataValue := shared.CreatePropertyMetadataValueStr(string{/* values here */})
```

### 

```go
propertyMetadataValue := shared.CreatePropertyMetadataValueNumber(float64{/* values here */})
```

### 

```go
propertyMetadataValue := shared.CreatePropertyMetadataValueBoolean(bool{/* values here */})
```

### 

```go
propertyMetadataValue := shared.CreatePropertyMetadataValueArrayOfPropertyMetadataSchemas5([]shared.PropertyMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch propertyMetadataValue.Type {
	case shared.PropertyMetadataValueTypeMapOfAny:
		// propertyMetadataValue.MapOfAny is populated
	case shared.PropertyMetadataValueTypeStr:
		// propertyMetadataValue.Str is populated
	case shared.PropertyMetadataValueTypeNumber:
		// propertyMetadataValue.Number is populated
	case shared.PropertyMetadataValueTypeBoolean:
		// propertyMetadataValue.Boolean is populated
	case shared.PropertyMetadataValueTypeArrayOfPropertyMetadataSchemas5:
		// propertyMetadataValue.ArrayOfPropertyMetadataSchemas5 is populated
}
```
