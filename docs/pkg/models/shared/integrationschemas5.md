# IntegrationSchemas5


## Supported Types

### IntegrationSchemas1

```go
integrationSchemas5 := shared.CreateIntegrationSchemas5IntegrationSchemas1(shared.IntegrationSchemas1{/* values here */})
```

### 

```go
integrationSchemas5 := shared.CreateIntegrationSchemas5Str(string{/* values here */})
```

### 

```go
integrationSchemas5 := shared.CreateIntegrationSchemas5Number(float64{/* values here */})
```

### 

```go
integrationSchemas5 := shared.CreateIntegrationSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationSchemas5.Type {
	case shared.IntegrationSchemas5TypeIntegrationSchemas1:
		// integrationSchemas5.IntegrationSchemas1 is populated
	case shared.IntegrationSchemas5TypeStr:
		// integrationSchemas5.Str is populated
	case shared.IntegrationSchemas5TypeNumber:
		// integrationSchemas5.Number is populated
	case shared.IntegrationSchemas5TypeBoolean:
		// integrationSchemas5.Boolean is populated
}
```
