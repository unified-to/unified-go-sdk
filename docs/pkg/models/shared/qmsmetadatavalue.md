# QmsMetadataValue

The field's value, typed by `format`.


## Supported Types

### 

```go
qmsMetadataValue := shared.CreateQmsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
qmsMetadataValue := shared.CreateQmsMetadataValueStr(string{/* values here */})
```

### 

```go
qmsMetadataValue := shared.CreateQmsMetadataValueNumber(float64{/* values here */})
```

### 

```go
qmsMetadataValue := shared.CreateQmsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
qmsMetadataValue := shared.CreateQmsMetadataValueArrayOfQmsMetadataSchemas5([]shared.QmsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch qmsMetadataValue.Type {
	case shared.QmsMetadataValueTypeMapOfAny:
		// qmsMetadataValue.MapOfAny is populated
	case shared.QmsMetadataValueTypeStr:
		// qmsMetadataValue.Str is populated
	case shared.QmsMetadataValueTypeNumber:
		// qmsMetadataValue.Number is populated
	case shared.QmsMetadataValueTypeBoolean:
		// qmsMetadataValue.Boolean is populated
	case shared.QmsMetadataValueTypeArrayOfQmsMetadataSchemas5:
		// qmsMetadataValue.ArrayOfQmsMetadataSchemas5 is populated
}
```
