# Integration5


## Supported Types

### Integration1

```go
integration5 := shared.CreateIntegration5Integration1(shared.Integration1{/* values here */})
```

### 

```go
integration5 := shared.CreateIntegration5Str(string{/* values here */})
```

### 

```go
integration5 := shared.CreateIntegration5Number(float64{/* values here */})
```

### 

```go
integration5 := shared.CreateIntegration5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integration5.Type {
	case shared.Integration5TypeIntegration1:
		// integration5.Integration1 is populated
	case shared.Integration5TypeStr:
		// integration5.Str is populated
	case shared.Integration5TypeNumber:
		// integration5.Number is populated
	case shared.Integration5TypeBoolean:
		// integration5.Boolean is populated
}
```
