# API


## Supported Types

### 

```go
api := shared.CreateAPIMapOfAny(map[string]any{/* values here */})
```

### 

```go
api := shared.CreateAPIStr(string{/* values here */})
```

### 

```go
api := shared.CreateAPINumber(float64{/* values here */})
```

### 

```go
api := shared.CreateAPIBoolean(bool{/* values here */})
```

### 

```go
api := shared.CreateAPIArrayOfIntegration5([]shared.Integration5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch api.Type {
	case shared.APITypeMapOfAny:
		// api.MapOfAny is populated
	case shared.APITypeStr:
		// api.Str is populated
	case shared.APITypeNumber:
		// api.Number is populated
	case shared.APITypeBoolean:
		// api.Boolean is populated
	case shared.APITypeArrayOfIntegration5:
		// api.ArrayOfIntegration5 is populated
}
```
