# LegalMetadata5


## Supported Types

### LegalMetadata1

```go
legalMetadata5 := shared.CreateLegalMetadata5LegalMetadata1(shared.LegalMetadata1{/* values here */})
```

### 

```go
legalMetadata5 := shared.CreateLegalMetadata5Str(string{/* values here */})
```

### 

```go
legalMetadata5 := shared.CreateLegalMetadata5Number(float64{/* values here */})
```

### 

```go
legalMetadata5 := shared.CreateLegalMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legalMetadata5.Type {
	case shared.LegalMetadata5TypeLegalMetadata1:
		// legalMetadata5.LegalMetadata1 is populated
	case shared.LegalMetadata5TypeStr:
		// legalMetadata5.Str is populated
	case shared.LegalMetadata5TypeNumber:
		// legalMetadata5.Number is populated
	case shared.LegalMetadata5TypeBoolean:
		// legalMetadata5.Boolean is populated
}
```
