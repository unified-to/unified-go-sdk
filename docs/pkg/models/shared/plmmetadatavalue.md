# PlmMetadataValue

The field's value, typed by `format`.


## Supported Types

### 

```go
plmMetadataValue := shared.CreatePlmMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
plmMetadataValue := shared.CreatePlmMetadataValueStr(string{/* values here */})
```

### 

```go
plmMetadataValue := shared.CreatePlmMetadataValueNumber(float64{/* values here */})
```

### 

```go
plmMetadataValue := shared.CreatePlmMetadataValueBoolean(bool{/* values here */})
```

### 

```go
plmMetadataValue := shared.CreatePlmMetadataValueArrayOfPlmMetadataSchemas5([]shared.PlmMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch plmMetadataValue.Type {
	case shared.PlmMetadataValueTypeMapOfAny:
		// plmMetadataValue.MapOfAny is populated
	case shared.PlmMetadataValueTypeStr:
		// plmMetadataValue.Str is populated
	case shared.PlmMetadataValueTypeNumber:
		// plmMetadataValue.Number is populated
	case shared.PlmMetadataValueTypeBoolean:
		// plmMetadataValue.Boolean is populated
	case shared.PlmMetadataValueTypeArrayOfPlmMetadataSchemas5:
		// plmMetadataValue.ArrayOfPlmMetadataSchemas5 is populated
}
```
