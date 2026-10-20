# GenaiMetadata5


## Supported Types

### GenaiMetadata1

```go
genaiMetadata5 := shared.CreateGenaiMetadata5GenaiMetadata1(shared.GenaiMetadata1{/* values here */})
```

### 

```go
genaiMetadata5 := shared.CreateGenaiMetadata5Str(string{/* values here */})
```

### 

```go
genaiMetadata5 := shared.CreateGenaiMetadata5Number(float64{/* values here */})
```

### 

```go
genaiMetadata5 := shared.CreateGenaiMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch genaiMetadata5.Type {
	case shared.GenaiMetadata5TypeGenaiMetadata1:
		// genaiMetadata5.GenaiMetadata1 is populated
	case shared.GenaiMetadata5TypeStr:
		// genaiMetadata5.Str is populated
	case shared.GenaiMetadata5TypeNumber:
		// genaiMetadata5.Number is populated
	case shared.GenaiMetadata5TypeBoolean:
		// genaiMetadata5.Boolean is populated
}
```
