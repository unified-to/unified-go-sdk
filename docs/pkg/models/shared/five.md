# Five


## Supported Types

### One

```go
five := shared.CreateFiveOne(shared.One{/* values here */})
```

### 

```go
five := shared.CreateFiveStr(string{/* values here */})
```

### 

```go
five := shared.CreateFiveNumber(float64{/* values here */})
```

### 

```go
five := shared.CreateFiveBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch five.Type {
	case shared.FiveTypeOne:
		// five.One is populated
	case shared.FiveTypeStr:
		// five.Str is populated
	case shared.FiveTypeNumber:
		// five.Number is populated
	case shared.FiveTypeBoolean:
		// five.Boolean is populated
}
```
