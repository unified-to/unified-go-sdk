# GenaiMetadataValue


## Supported Types

### 

```go
genaiMetadataValue := shared.CreateGenaiMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
genaiMetadataValue := shared.CreateGenaiMetadataValueStr(string{/* values here */})
```

### 

```go
genaiMetadataValue := shared.CreateGenaiMetadataValueNumber(float64{/* values here */})
```

### 

```go
genaiMetadataValue := shared.CreateGenaiMetadataValueBoolean(bool{/* values here */})
```

### 

```go
genaiMetadataValue := shared.CreateGenaiMetadataValueArrayOfGenaiMetadataSchemas5([]shared.GenaiMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch genaiMetadataValue.Type {
	case shared.GenaiMetadataValueTypeMapOfAny:
		// genaiMetadataValue.MapOfAny is populated
	case shared.GenaiMetadataValueTypeStr:
		// genaiMetadataValue.Str is populated
	case shared.GenaiMetadataValueTypeNumber:
		// genaiMetadataValue.Number is populated
	case shared.GenaiMetadataValueTypeBoolean:
		// genaiMetadataValue.Boolean is populated
	case shared.GenaiMetadataValueTypeArrayOfGenaiMetadataSchemas5:
		// genaiMetadataValue.ArrayOfGenaiMetadataSchemas5 is populated
}
```
