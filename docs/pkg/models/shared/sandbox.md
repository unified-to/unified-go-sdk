# Sandbox


## Supported Types

### 

```go
sandbox := shared.CreateSandboxMapOfAny(map[string]any{/* values here */})
```

### 

```go
sandbox := shared.CreateSandboxStr(string{/* values here */})
```

### 

```go
sandbox := shared.CreateSandboxNumber(float64{/* values here */})
```

### 

```go
sandbox := shared.CreateSandboxBoolean(bool{/* values here */})
```

### 

```go
sandbox := shared.CreateSandboxArrayOfIntegrationSchemasSandbox5([]shared.IntegrationSchemasSandbox5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sandbox.Type {
	case shared.SandboxTypeMapOfAny:
		// sandbox.MapOfAny is populated
	case shared.SandboxTypeStr:
		// sandbox.Str is populated
	case shared.SandboxTypeNumber:
		// sandbox.Number is populated
	case shared.SandboxTypeBoolean:
		// sandbox.Boolean is populated
	case shared.SandboxTypeArrayOfIntegrationSchemasSandbox5:
		// sandbox.ArrayOfIntegrationSchemasSandbox5 is populated
}
```
