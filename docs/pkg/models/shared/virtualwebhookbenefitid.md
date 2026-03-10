# VirtualWebhookBenefitID


## Supported Types

### 

```go
virtualWebhookBenefitID := shared.CreateVirtualWebhookBenefitIDMapOfAny(map[string]any{/* values here */})
```

### 

```go
virtualWebhookBenefitID := shared.CreateVirtualWebhookBenefitIDStr(string{/* values here */})
```

### 

```go
virtualWebhookBenefitID := shared.CreateVirtualWebhookBenefitIDNumber(float64{/* values here */})
```

### 

```go
virtualWebhookBenefitID := shared.CreateVirtualWebhookBenefitIDBoolean(bool{/* values here */})
```

### 

```go
virtualWebhookBenefitID := shared.CreateVirtualWebhookBenefitIDArrayOfIntegrationSupport5([]shared.IntegrationSupport5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualWebhookBenefitID.Type {
	case shared.VirtualWebhookBenefitIDTypeMapOfAny:
		// virtualWebhookBenefitID.MapOfAny is populated
	case shared.VirtualWebhookBenefitIDTypeStr:
		// virtualWebhookBenefitID.Str is populated
	case shared.VirtualWebhookBenefitIDTypeNumber:
		// virtualWebhookBenefitID.Number is populated
	case shared.VirtualWebhookBenefitIDTypeBoolean:
		// virtualWebhookBenefitID.Boolean is populated
	case shared.VirtualWebhookBenefitIDTypeArrayOfIntegrationSupport5:
		// virtualWebhookBenefitID.ArrayOfIntegrationSupport5 is populated
}
```
