# Saml


## Supported Types

### 

```go
saml := shared.CreateSamlMapOfAny(map[string]any{/* values here */})
```

### 

```go
saml := shared.CreateSamlStr(string{/* values here */})
```

### 

```go
saml := shared.CreateSamlNumber(float64{/* values here */})
```

### 

```go
saml := shared.CreateSamlBoolean(bool{/* values here */})
```

### 

```go
saml := shared.CreateSamlArrayOfIntegrationSchemasSaml5([]shared.IntegrationSchemasSaml5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch saml.Type {
	case shared.SamlTypeMapOfAny:
		// saml.MapOfAny is populated
	case shared.SamlTypeStr:
		// saml.Str is populated
	case shared.SamlTypeNumber:
		// saml.Number is populated
	case shared.SamlTypeBoolean:
		// saml.Boolean is populated
	case shared.SamlTypeArrayOfIntegrationSchemasSaml5:
		// saml.ArrayOfIntegrationSchemasSaml5 is populated
}
```
