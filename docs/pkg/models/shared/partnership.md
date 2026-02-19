# Partnership


## Supported Types

### 

```go
partnership := shared.CreatePartnershipMapOfAny(map[string]any{/* values here */})
```

### 

```go
partnership := shared.CreatePartnershipStr(string{/* values here */})
```

### 

```go
partnership := shared.CreatePartnershipNumber(float64{/* values here */})
```

### 

```go
partnership := shared.CreatePartnershipBoolean(bool{/* values here */})
```

### 

```go
partnership := shared.CreatePartnershipArrayOfIntegrationSchemas5([]shared.IntegrationSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch partnership.Type {
	case shared.PartnershipTypeMapOfAny:
		// partnership.MapOfAny is populated
	case shared.PartnershipTypeStr:
		// partnership.Str is populated
	case shared.PartnershipTypeNumber:
		// partnership.Number is populated
	case shared.PartnershipTypeBoolean:
		// partnership.Boolean is populated
	case shared.PartnershipTypeArrayOfIntegrationSchemas5:
		// partnership.ArrayOfIntegrationSchemas5 is populated
}
```
