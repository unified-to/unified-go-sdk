# IntegrationSupportSchemas5


## Supported Types

### IntegrationSupportSchemas1

```go
integrationSupportSchemas5 := shared.CreateIntegrationSupportSchemas5IntegrationSupportSchemas1(shared.IntegrationSupportSchemas1{/* values here */})
```

### 

```go
integrationSupportSchemas5 := shared.CreateIntegrationSupportSchemas5Str(string{/* values here */})
```

### 

```go
integrationSupportSchemas5 := shared.CreateIntegrationSupportSchemas5Number(float64{/* values here */})
```

### 

```go
integrationSupportSchemas5 := shared.CreateIntegrationSupportSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationSupportSchemas5.Type {
	case shared.IntegrationSupportSchemas5TypeIntegrationSupportSchemas1:
		// integrationSupportSchemas5.IntegrationSupportSchemas1 is populated
	case shared.IntegrationSupportSchemas5TypeStr:
		// integrationSupportSchemas5.Str is populated
	case shared.IntegrationSupportSchemas5TypeNumber:
		// integrationSupportSchemas5.Number is populated
	case shared.IntegrationSupportSchemas5TypeBoolean:
		// integrationSupportSchemas5.Boolean is populated
}
```
