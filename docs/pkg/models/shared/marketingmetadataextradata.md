# MarketingMetadataExtraData


## Supported Types

### 

```go
marketingMetadataExtraData := shared.CreateMarketingMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
marketingMetadataExtraData := shared.CreateMarketingMetadataExtraDataStr(string{/* values here */})
```

### 

```go
marketingMetadataExtraData := shared.CreateMarketingMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
marketingMetadataExtraData := shared.CreateMarketingMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
marketingMetadataExtraData := shared.CreateMarketingMetadataExtraDataArrayOfMarketingMetadata5([]shared.MarketingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch marketingMetadataExtraData.Type {
	case shared.MarketingMetadataExtraDataTypeMapOfAny:
		// marketingMetadataExtraData.MapOfAny is populated
	case shared.MarketingMetadataExtraDataTypeStr:
		// marketingMetadataExtraData.Str is populated
	case shared.MarketingMetadataExtraDataTypeNumber:
		// marketingMetadataExtraData.Number is populated
	case shared.MarketingMetadataExtraDataTypeBoolean:
		// marketingMetadataExtraData.Boolean is populated
	case shared.MarketingMetadataExtraDataTypeArrayOfMarketingMetadata5:
		// marketingMetadataExtraData.ArrayOfMarketingMetadata5 is populated
}
```
