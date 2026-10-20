# MarketingMetadataSchemas5


## Supported Types

### MarketingMetadataSchemas1

```go
marketingMetadataSchemas5 := shared.CreateMarketingMetadataSchemas5MarketingMetadataSchemas1(shared.MarketingMetadataSchemas1{/* values here */})
```

### 

```go
marketingMetadataSchemas5 := shared.CreateMarketingMetadataSchemas5Str(string{/* values here */})
```

### 

```go
marketingMetadataSchemas5 := shared.CreateMarketingMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
marketingMetadataSchemas5 := shared.CreateMarketingMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch marketingMetadataSchemas5.Type {
	case shared.MarketingMetadataSchemas5TypeMarketingMetadataSchemas1:
		// marketingMetadataSchemas5.MarketingMetadataSchemas1 is populated
	case shared.MarketingMetadataSchemas5TypeStr:
		// marketingMetadataSchemas5.Str is populated
	case shared.MarketingMetadataSchemas5TypeNumber:
		// marketingMetadataSchemas5.Number is populated
	case shared.MarketingMetadataSchemas5TypeBoolean:
		// marketingMetadataSchemas5.Boolean is populated
}
```
