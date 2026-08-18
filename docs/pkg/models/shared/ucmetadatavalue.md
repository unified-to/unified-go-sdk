# UcMetadataValue


## Supported Types

### 

```go
ucMetadataValue := shared.CreateUcMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
ucMetadataValue := shared.CreateUcMetadataValueStr(string{/* values here */})
```

### 

```go
ucMetadataValue := shared.CreateUcMetadataValueNumber(float64{/* values here */})
```

### 

```go
ucMetadataValue := shared.CreateUcMetadataValueBoolean(bool{/* values here */})
```

### 

```go
ucMetadataValue := shared.CreateUcMetadataValueArrayOfUcMetadataSchemas5([]shared.UcMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ucMetadataValue.Type {
	case shared.UcMetadataValueTypeMapOfAny:
		// ucMetadataValue.MapOfAny is populated
	case shared.UcMetadataValueTypeStr:
		// ucMetadataValue.Str is populated
	case shared.UcMetadataValueTypeNumber:
		// ucMetadataValue.Number is populated
	case shared.UcMetadataValueTypeBoolean:
		// ucMetadataValue.Boolean is populated
	case shared.UcMetadataValueTypeArrayOfUcMetadataSchemas5:
		// ucMetadataValue.ArrayOfUcMetadataSchemas5 is populated
}
```
