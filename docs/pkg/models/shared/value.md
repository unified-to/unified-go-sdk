# Value


## Supported Types

### 

```go
value := shared.CreateValueNumber(float64{/* values here */})
```

### 

```go
value := shared.CreateValueStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value.Type {
	case shared.ValueTypeNumber:
		// value.Number is populated
	case shared.ValueTypeStr:
		// value.Str is populated
}
```
