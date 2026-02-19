# IntegrationSchemasSandbox5


## Supported Types

### IntegrationSchemasSandbox1

```go
integrationSchemasSandbox5 := shared.CreateIntegrationSchemasSandbox5IntegrationSchemasSandbox1(shared.IntegrationSchemasSandbox1{/* values here */})
```

### 

```go
integrationSchemasSandbox5 := shared.CreateIntegrationSchemasSandbox5Str(string{/* values here */})
```

### 

```go
integrationSchemasSandbox5 := shared.CreateIntegrationSchemasSandbox5Number(float64{/* values here */})
```

### 

```go
integrationSchemasSandbox5 := shared.CreateIntegrationSchemasSandbox5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationSchemasSandbox5.Type {
	case shared.IntegrationSchemasSandbox5TypeIntegrationSchemasSandbox1:
		// integrationSchemasSandbox5.IntegrationSchemasSandbox1 is populated
	case shared.IntegrationSchemasSandbox5TypeStr:
		// integrationSchemasSandbox5.Str is populated
	case shared.IntegrationSchemasSandbox5TypeNumber:
		// integrationSchemasSandbox5.Number is populated
	case shared.IntegrationSchemasSandbox5TypeBoolean:
		// integrationSchemasSandbox5.Boolean is populated
}
```
