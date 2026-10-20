# MarketingMetadataValue


## Supported Types

### 

```go
marketingMetadataValue := shared.CreateMarketingMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
marketingMetadataValue := shared.CreateMarketingMetadataValueStr(string{/* values here */})
```

### 

```go
marketingMetadataValue := shared.CreateMarketingMetadataValueNumber(float64{/* values here */})
```

### 

```go
marketingMetadataValue := shared.CreateMarketingMetadataValueBoolean(bool{/* values here */})
```

### 

```go
marketingMetadataValue := shared.CreateMarketingMetadataValueArrayOfMarketingMetadataSchemas5([]shared.MarketingMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch marketingMetadataValue.Type {
	case shared.MarketingMetadataValueTypeMapOfAny:
		// marketingMetadataValue.MapOfAny is populated
	case shared.MarketingMetadataValueTypeStr:
		// marketingMetadataValue.Str is populated
	case shared.MarketingMetadataValueTypeNumber:
		// marketingMetadataValue.Number is populated
	case shared.MarketingMetadataValueTypeBoolean:
		// marketingMetadataValue.Boolean is populated
	case shared.MarketingMetadataValueTypeArrayOfMarketingMetadataSchemas5:
		// marketingMetadataValue.ArrayOfMarketingMetadataSchemas5 is populated
}
```
