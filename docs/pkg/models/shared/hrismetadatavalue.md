# HrisMetadataValue


## Supported Types

### 

```go
hrisMetadataValue := shared.CreateHrisMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
hrisMetadataValue := shared.CreateHrisMetadataValueStr(string{/* values here */})
```

### 

```go
hrisMetadataValue := shared.CreateHrisMetadataValueNumber(float64{/* values here */})
```

### 

```go
hrisMetadataValue := shared.CreateHrisMetadataValueBoolean(bool{/* values here */})
```

### 

```go
hrisMetadataValue := shared.CreateHrisMetadataValueArrayOfHrisMetadataSchemas5([]shared.HrisMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hrisMetadataValue.Type {
	case shared.HrisMetadataValueTypeMapOfAny:
		// hrisMetadataValue.MapOfAny is populated
	case shared.HrisMetadataValueTypeStr:
		// hrisMetadataValue.Str is populated
	case shared.HrisMetadataValueTypeNumber:
		// hrisMetadataValue.Number is populated
	case shared.HrisMetadataValueTypeBoolean:
		// hrisMetadataValue.Boolean is populated
	case shared.HrisMetadataValueTypeArrayOfHrisMetadataSchemas5:
		// hrisMetadataValue.ArrayOfHrisMetadataSchemas5 is populated
}
```
