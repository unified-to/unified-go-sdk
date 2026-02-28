# Value


## Supported Types

### 

```go
value := shared.CreateValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
value := shared.CreateValueStr(string{/* values here */})
```

### 

```go
value := shared.CreateValueNumber(float64{/* values here */})
```

### 

```go
value := shared.CreateValueBoolean(bool{/* values here */})
```

### 

```go
value := shared.CreateValueArrayOf5([]shared.Five{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value.Type {
	case shared.ValueTypeMapOfAny:
		// value.MapOfAny is populated
	case shared.ValueTypeStr:
		// value.Str is populated
	case shared.ValueTypeNumber:
		// value.Number is populated
	case shared.ValueTypeBoolean:
		// value.Boolean is populated
	case shared.ValueTypeArrayOf5:
		// value.ArrayOf5 is populated
}
```
