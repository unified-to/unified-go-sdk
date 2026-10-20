# LmsMetadata5


## Supported Types

### LmsMetadata1

```go
lmsMetadata5 := shared.CreateLmsMetadata5LmsMetadata1(shared.LmsMetadata1{/* values here */})
```

### 

```go
lmsMetadata5 := shared.CreateLmsMetadata5Str(string{/* values here */})
```

### 

```go
lmsMetadata5 := shared.CreateLmsMetadata5Number(float64{/* values here */})
```

### 

```go
lmsMetadata5 := shared.CreateLmsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lmsMetadata5.Type {
	case shared.LmsMetadata5TypeLmsMetadata1:
		// lmsMetadata5.LmsMetadata1 is populated
	case shared.LmsMetadata5TypeStr:
		// lmsMetadata5.Str is populated
	case shared.LmsMetadata5TypeNumber:
		// lmsMetadata5.Number is populated
	case shared.LmsMetadata5TypeBoolean:
		// lmsMetadata5.Boolean is populated
}
```
