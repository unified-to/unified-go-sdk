# HrisMetadataExtraData


## Supported Types

### 

```go
hrisMetadataExtraData := shared.CreateHrisMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
hrisMetadataExtraData := shared.CreateHrisMetadataExtraDataStr(string{/* values here */})
```

### 

```go
hrisMetadataExtraData := shared.CreateHrisMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
hrisMetadataExtraData := shared.CreateHrisMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
hrisMetadataExtraData := shared.CreateHrisMetadataExtraDataArrayOfHrisMetadata5([]shared.HrisMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hrisMetadataExtraData.Type {
	case shared.HrisMetadataExtraDataTypeMapOfAny:
		// hrisMetadataExtraData.MapOfAny is populated
	case shared.HrisMetadataExtraDataTypeStr:
		// hrisMetadataExtraData.Str is populated
	case shared.HrisMetadataExtraDataTypeNumber:
		// hrisMetadataExtraData.Number is populated
	case shared.HrisMetadataExtraDataTypeBoolean:
		// hrisMetadataExtraData.Boolean is populated
	case shared.HrisMetadataExtraDataTypeArrayOfHrisMetadata5:
		// hrisMetadataExtraData.ArrayOfHrisMetadata5 is populated
}
```
