# QmsMetadataExtraData


## Supported Types

### 

```go
qmsMetadataExtraData := shared.CreateQmsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
qmsMetadataExtraData := shared.CreateQmsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
qmsMetadataExtraData := shared.CreateQmsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
qmsMetadataExtraData := shared.CreateQmsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
qmsMetadataExtraData := shared.CreateQmsMetadataExtraDataArrayOfQmsMetadata5([]shared.QmsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch qmsMetadataExtraData.Type {
	case shared.QmsMetadataExtraDataTypeMapOfAny:
		// qmsMetadataExtraData.MapOfAny is populated
	case shared.QmsMetadataExtraDataTypeStr:
		// qmsMetadataExtraData.Str is populated
	case shared.QmsMetadataExtraDataTypeNumber:
		// qmsMetadataExtraData.Number is populated
	case shared.QmsMetadataExtraDataTypeBoolean:
		// qmsMetadataExtraData.Boolean is populated
	case shared.QmsMetadataExtraDataTypeArrayOfQmsMetadata5:
		// qmsMetadataExtraData.ArrayOfQmsMetadata5 is populated
}
```
