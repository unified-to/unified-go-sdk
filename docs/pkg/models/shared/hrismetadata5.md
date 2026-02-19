# HrisMetadata5


## Supported Types

### HrisMetadata1

```go
hrisMetadata5 := shared.CreateHrisMetadata5HrisMetadata1(shared.HrisMetadata1{/* values here */})
```

### 

```go
hrisMetadata5 := shared.CreateHrisMetadata5Str(string{/* values here */})
```

### 

```go
hrisMetadata5 := shared.CreateHrisMetadata5Number(float64{/* values here */})
```

### 

```go
hrisMetadata5 := shared.CreateHrisMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hrisMetadata5.Type {
	case shared.HrisMetadata5TypeHrisMetadata1:
		// hrisMetadata5.HrisMetadata1 is populated
	case shared.HrisMetadata5TypeStr:
		// hrisMetadata5.Str is populated
	case shared.HrisMetadata5TypeNumber:
		// hrisMetadata5.Number is populated
	case shared.HrisMetadata5TypeBoolean:
		// hrisMetadata5.Boolean is populated
}
```
