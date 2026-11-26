# DmsMetadataValue

The field's value, typed by `format`.


## Supported Types

### 

```go
dmsMetadataValue := shared.CreateDmsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
dmsMetadataValue := shared.CreateDmsMetadataValueStr(string{/* values here */})
```

### 

```go
dmsMetadataValue := shared.CreateDmsMetadataValueNumber(float64{/* values here */})
```

### 

```go
dmsMetadataValue := shared.CreateDmsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
dmsMetadataValue := shared.CreateDmsMetadataValueArrayOfDmsMetadataSchemas5([]shared.DmsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dmsMetadataValue.Type {
	case shared.DmsMetadataValueTypeMapOfAny:
		// dmsMetadataValue.MapOfAny is populated
	case shared.DmsMetadataValueTypeStr:
		// dmsMetadataValue.Str is populated
	case shared.DmsMetadataValueTypeNumber:
		// dmsMetadataValue.Number is populated
	case shared.DmsMetadataValueTypeBoolean:
		// dmsMetadataValue.Boolean is populated
	case shared.DmsMetadataValueTypeArrayOfDmsMetadataSchemas5:
		// dmsMetadataValue.ArrayOfDmsMetadataSchemas5 is populated
}
```
