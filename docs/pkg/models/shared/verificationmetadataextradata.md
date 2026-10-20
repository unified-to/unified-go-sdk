# VerificationMetadataExtraData


## Supported Types

### 

```go
verificationMetadataExtraData := shared.CreateVerificationMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
verificationMetadataExtraData := shared.CreateVerificationMetadataExtraDataStr(string{/* values here */})
```

### 

```go
verificationMetadataExtraData := shared.CreateVerificationMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
verificationMetadataExtraData := shared.CreateVerificationMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
verificationMetadataExtraData := shared.CreateVerificationMetadataExtraDataArrayOfVerificationMetadata5([]shared.VerificationMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch verificationMetadataExtraData.Type {
	case shared.VerificationMetadataExtraDataTypeMapOfAny:
		// verificationMetadataExtraData.MapOfAny is populated
	case shared.VerificationMetadataExtraDataTypeStr:
		// verificationMetadataExtraData.Str is populated
	case shared.VerificationMetadataExtraDataTypeNumber:
		// verificationMetadataExtraData.Number is populated
	case shared.VerificationMetadataExtraDataTypeBoolean:
		// verificationMetadataExtraData.Boolean is populated
	case shared.VerificationMetadataExtraDataTypeArrayOfVerificationMetadata5:
		// verificationMetadataExtraData.ArrayOfVerificationMetadata5 is populated
}
```
