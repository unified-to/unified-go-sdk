# IntegrationSupport5


## Supported Types

### IntegrationSupport1

```go
integrationSupport5 := shared.CreateIntegrationSupport5IntegrationSupport1(shared.IntegrationSupport1{/* values here */})
```

### 

```go
integrationSupport5 := shared.CreateIntegrationSupport5Str(string{/* values here */})
```

### 

```go
integrationSupport5 := shared.CreateIntegrationSupport5Number(float64{/* values here */})
```

### 

```go
integrationSupport5 := shared.CreateIntegrationSupport5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationSupport5.Type {
	case shared.IntegrationSupport5TypeIntegrationSupport1:
		// integrationSupport5.IntegrationSupport1 is populated
	case shared.IntegrationSupport5TypeStr:
		// integrationSupport5.Str is populated
	case shared.IntegrationSupport5TypeNumber:
		// integrationSupport5.Number is populated
	case shared.IntegrationSupport5TypeBoolean:
		// integrationSupport5.Boolean is populated
}
```
