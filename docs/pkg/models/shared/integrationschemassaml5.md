# IntegrationSchemasSaml5


## Supported Types

### IntegrationSchemasSaml1

```go
integrationSchemasSaml5 := shared.CreateIntegrationSchemasSaml5IntegrationSchemasSaml1(shared.IntegrationSchemasSaml1{/* values here */})
```

### 

```go
integrationSchemasSaml5 := shared.CreateIntegrationSchemasSaml5Str(string{/* values here */})
```

### 

```go
integrationSchemasSaml5 := shared.CreateIntegrationSchemasSaml5Number(float64{/* values here */})
```

### 

```go
integrationSchemasSaml5 := shared.CreateIntegrationSchemasSaml5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationSchemasSaml5.Type {
	case shared.IntegrationSchemasSaml5TypeIntegrationSchemasSaml1:
		// integrationSchemasSaml5.IntegrationSchemasSaml1 is populated
	case shared.IntegrationSchemasSaml5TypeStr:
		// integrationSchemasSaml5.Str is populated
	case shared.IntegrationSchemasSaml5TypeNumber:
		// integrationSchemasSaml5.Number is populated
	case shared.IntegrationSchemasSaml5TypeBoolean:
		// integrationSchemasSaml5.Boolean is populated
}
```
