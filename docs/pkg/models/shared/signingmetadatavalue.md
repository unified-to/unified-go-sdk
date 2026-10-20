# SigningMetadataValue


## Supported Types

### 

```go
signingMetadataValue := shared.CreateSigningMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
signingMetadataValue := shared.CreateSigningMetadataValueStr(string{/* values here */})
```

### 

```go
signingMetadataValue := shared.CreateSigningMetadataValueNumber(float64{/* values here */})
```

### 

```go
signingMetadataValue := shared.CreateSigningMetadataValueBoolean(bool{/* values here */})
```

### 

```go
signingMetadataValue := shared.CreateSigningMetadataValueArrayOfSigningMetadataSchemas5([]shared.SigningMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signingMetadataValue.Type {
	case shared.SigningMetadataValueTypeMapOfAny:
		// signingMetadataValue.MapOfAny is populated
	case shared.SigningMetadataValueTypeStr:
		// signingMetadataValue.Str is populated
	case shared.SigningMetadataValueTypeNumber:
		// signingMetadataValue.Number is populated
	case shared.SigningMetadataValueTypeBoolean:
		// signingMetadataValue.Boolean is populated
	case shared.SigningMetadataValueTypeArrayOfSigningMetadataSchemas5:
		// signingMetadataValue.ArrayOfSigningMetadataSchemas5 is populated
}
```
