# SigningMetadata5


## Supported Types

### SigningMetadata1

```go
signingMetadata5 := shared.CreateSigningMetadata5SigningMetadata1(shared.SigningMetadata1{/* values here */})
```

### 

```go
signingMetadata5 := shared.CreateSigningMetadata5Str(string{/* values here */})
```

### 

```go
signingMetadata5 := shared.CreateSigningMetadata5Number(float64{/* values here */})
```

### 

```go
signingMetadata5 := shared.CreateSigningMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signingMetadata5.Type {
	case shared.SigningMetadata5TypeSigningMetadata1:
		// signingMetadata5.SigningMetadata1 is populated
	case shared.SigningMetadata5TypeStr:
		// signingMetadata5.Str is populated
	case shared.SigningMetadata5TypeNumber:
		// signingMetadata5.Number is populated
	case shared.SigningMetadata5TypeBoolean:
		// signingMetadata5.Boolean is populated
}
```
