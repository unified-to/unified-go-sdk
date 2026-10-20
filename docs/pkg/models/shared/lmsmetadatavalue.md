# LmsMetadataValue


## Supported Types

### 

```go
lmsMetadataValue := shared.CreateLmsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
lmsMetadataValue := shared.CreateLmsMetadataValueStr(string{/* values here */})
```

### 

```go
lmsMetadataValue := shared.CreateLmsMetadataValueNumber(float64{/* values here */})
```

### 

```go
lmsMetadataValue := shared.CreateLmsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
lmsMetadataValue := shared.CreateLmsMetadataValueArrayOfLmsMetadataSchemas5([]shared.LmsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lmsMetadataValue.Type {
	case shared.LmsMetadataValueTypeMapOfAny:
		// lmsMetadataValue.MapOfAny is populated
	case shared.LmsMetadataValueTypeStr:
		// lmsMetadataValue.Str is populated
	case shared.LmsMetadataValueTypeNumber:
		// lmsMetadataValue.Number is populated
	case shared.LmsMetadataValueTypeBoolean:
		// lmsMetadataValue.Boolean is populated
	case shared.LmsMetadataValueTypeArrayOfLmsMetadataSchemas5:
		// lmsMetadataValue.ArrayOfLmsMetadataSchemas5 is populated
}
```
