# UcMetadataExtraData


## Supported Types

### 

```go
ucMetadataExtraData := shared.CreateUcMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
ucMetadataExtraData := shared.CreateUcMetadataExtraDataStr(string{/* values here */})
```

### 

```go
ucMetadataExtraData := shared.CreateUcMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
ucMetadataExtraData := shared.CreateUcMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
ucMetadataExtraData := shared.CreateUcMetadataExtraDataArrayOfUcMetadata5([]shared.UcMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ucMetadataExtraData.Type {
	case shared.UcMetadataExtraDataTypeMapOfAny:
		// ucMetadataExtraData.MapOfAny is populated
	case shared.UcMetadataExtraDataTypeStr:
		// ucMetadataExtraData.Str is populated
	case shared.UcMetadataExtraDataTypeNumber:
		// ucMetadataExtraData.Number is populated
	case shared.UcMetadataExtraDataTypeBoolean:
		// ucMetadataExtraData.Boolean is populated
	case shared.UcMetadataExtraDataTypeArrayOfUcMetadata5:
		// ucMetadataExtraData.ArrayOfUcMetadata5 is populated
}
```
