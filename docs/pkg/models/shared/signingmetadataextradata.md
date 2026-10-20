# SigningMetadataExtraData


## Supported Types

### 

```go
signingMetadataExtraData := shared.CreateSigningMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
signingMetadataExtraData := shared.CreateSigningMetadataExtraDataStr(string{/* values here */})
```

### 

```go
signingMetadataExtraData := shared.CreateSigningMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
signingMetadataExtraData := shared.CreateSigningMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
signingMetadataExtraData := shared.CreateSigningMetadataExtraDataArrayOfSigningMetadata5([]shared.SigningMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signingMetadataExtraData.Type {
	case shared.SigningMetadataExtraDataTypeMapOfAny:
		// signingMetadataExtraData.MapOfAny is populated
	case shared.SigningMetadataExtraDataTypeStr:
		// signingMetadataExtraData.Str is populated
	case shared.SigningMetadataExtraDataTypeNumber:
		// signingMetadataExtraData.Number is populated
	case shared.SigningMetadataExtraDataTypeBoolean:
		// signingMetadataExtraData.Boolean is populated
	case shared.SigningMetadataExtraDataTypeArrayOfSigningMetadata5:
		// signingMetadataExtraData.ArrayOfSigningMetadata5 is populated
}
```
