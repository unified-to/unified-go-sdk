# VirtualWebhookPaymentID


## Supported Types

### 

```go
virtualWebhookPaymentID := shared.CreateVirtualWebhookPaymentIDMapOfAny(map[string]any{/* values here */})
```

### 

```go
virtualWebhookPaymentID := shared.CreateVirtualWebhookPaymentIDStr(string{/* values here */})
```

### 

```go
virtualWebhookPaymentID := shared.CreateVirtualWebhookPaymentIDNumber(float64{/* values here */})
```

### 

```go
virtualWebhookPaymentID := shared.CreateVirtualWebhookPaymentIDBoolean(bool{/* values here */})
```

### 

```go
virtualWebhookPaymentID := shared.CreateVirtualWebhookPaymentIDArrayOfIntegrationSupportSchemas5([]shared.IntegrationSupportSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualWebhookPaymentID.Type {
	case shared.VirtualWebhookPaymentIDTypeMapOfAny:
		// virtualWebhookPaymentID.MapOfAny is populated
	case shared.VirtualWebhookPaymentIDTypeStr:
		// virtualWebhookPaymentID.Str is populated
	case shared.VirtualWebhookPaymentIDTypeNumber:
		// virtualWebhookPaymentID.Number is populated
	case shared.VirtualWebhookPaymentIDTypeBoolean:
		// virtualWebhookPaymentID.Boolean is populated
	case shared.VirtualWebhookPaymentIDTypeArrayOfIntegrationSupportSchemas5:
		// virtualWebhookPaymentID.ArrayOfIntegrationSupportSchemas5 is populated
}
```
