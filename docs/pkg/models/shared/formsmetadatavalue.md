# FormsMetadataValue


## Supported Types

### 

```go
formsMetadataValue := shared.CreateFormsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
formsMetadataValue := shared.CreateFormsMetadataValueStr(string{/* values here */})
```

### 

```go
formsMetadataValue := shared.CreateFormsMetadataValueNumber(float64{/* values here */})
```

### 

```go
formsMetadataValue := shared.CreateFormsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
formsMetadataValue := shared.CreateFormsMetadataValueArrayOfFormsMetadataSchemas5([]shared.FormsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formsMetadataValue.Type {
	case shared.FormsMetadataValueTypeMapOfAny:
		// formsMetadataValue.MapOfAny is populated
	case shared.FormsMetadataValueTypeStr:
		// formsMetadataValue.Str is populated
	case shared.FormsMetadataValueTypeNumber:
		// formsMetadataValue.Number is populated
	case shared.FormsMetadataValueTypeBoolean:
		// formsMetadataValue.Boolean is populated
	case shared.FormsMetadataValueTypeArrayOfFormsMetadataSchemas5:
		// formsMetadataValue.ArrayOfFormsMetadataSchemas5 is populated
}
```
