# VerificationMetadataSchemas5


## Supported Types

### VerificationMetadataSchemas1

```go
verificationMetadataSchemas5 := shared.CreateVerificationMetadataSchemas5VerificationMetadataSchemas1(shared.VerificationMetadataSchemas1{/* values here */})
```

### 

```go
verificationMetadataSchemas5 := shared.CreateVerificationMetadataSchemas5Str(string{/* values here */})
```

### 

```go
verificationMetadataSchemas5 := shared.CreateVerificationMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
verificationMetadataSchemas5 := shared.CreateVerificationMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch verificationMetadataSchemas5.Type {
	case shared.VerificationMetadataSchemas5TypeVerificationMetadataSchemas1:
		// verificationMetadataSchemas5.VerificationMetadataSchemas1 is populated
	case shared.VerificationMetadataSchemas5TypeStr:
		// verificationMetadataSchemas5.Str is populated
	case shared.VerificationMetadataSchemas5TypeNumber:
		// verificationMetadataSchemas5.Number is populated
	case shared.VerificationMetadataSchemas5TypeBoolean:
		// verificationMetadataSchemas5.Boolean is populated
}
```
