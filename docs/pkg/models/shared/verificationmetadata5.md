# VerificationMetadata5


## Supported Types

### VerificationMetadata1

```go
verificationMetadata5 := shared.CreateVerificationMetadata5VerificationMetadata1(shared.VerificationMetadata1{/* values here */})
```

### 

```go
verificationMetadata5 := shared.CreateVerificationMetadata5Str(string{/* values here */})
```

### 

```go
verificationMetadata5 := shared.CreateVerificationMetadata5Number(float64{/* values here */})
```

### 

```go
verificationMetadata5 := shared.CreateVerificationMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch verificationMetadata5.Type {
	case shared.VerificationMetadata5TypeVerificationMetadata1:
		// verificationMetadata5.VerificationMetadata1 is populated
	case shared.VerificationMetadata5TypeStr:
		// verificationMetadata5.Str is populated
	case shared.VerificationMetadata5TypeNumber:
		// verificationMetadata5.Number is populated
	case shared.VerificationMetadata5TypeBoolean:
		// verificationMetadata5.Boolean is populated
}
```
