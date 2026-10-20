# VerificationMetadataValue


## Supported Types

### 

```go
verificationMetadataValue := shared.CreateVerificationMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
verificationMetadataValue := shared.CreateVerificationMetadataValueStr(string{/* values here */})
```

### 

```go
verificationMetadataValue := shared.CreateVerificationMetadataValueNumber(float64{/* values here */})
```

### 

```go
verificationMetadataValue := shared.CreateVerificationMetadataValueBoolean(bool{/* values here */})
```

### 

```go
verificationMetadataValue := shared.CreateVerificationMetadataValueArrayOfVerificationMetadataSchemas5([]shared.VerificationMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch verificationMetadataValue.Type {
	case shared.VerificationMetadataValueTypeMapOfAny:
		// verificationMetadataValue.MapOfAny is populated
	case shared.VerificationMetadataValueTypeStr:
		// verificationMetadataValue.Str is populated
	case shared.VerificationMetadataValueTypeNumber:
		// verificationMetadataValue.Number is populated
	case shared.VerificationMetadataValueTypeBoolean:
		// verificationMetadataValue.Boolean is populated
	case shared.VerificationMetadataValueTypeArrayOfVerificationMetadataSchemas5:
		// verificationMetadataValue.ArrayOfVerificationMetadataSchemas5 is populated
}
```
