# MarketingMetadata5


## Supported Types

### MarketingMetadata1

```go
marketingMetadata5 := shared.CreateMarketingMetadata5MarketingMetadata1(shared.MarketingMetadata1{/* values here */})
```

### 

```go
marketingMetadata5 := shared.CreateMarketingMetadata5Str(string{/* values here */})
```

### 

```go
marketingMetadata5 := shared.CreateMarketingMetadata5Number(float64{/* values here */})
```

### 

```go
marketingMetadata5 := shared.CreateMarketingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch marketingMetadata5.Type {
	case shared.MarketingMetadata5TypeMarketingMetadata1:
		// marketingMetadata5.MarketingMetadata1 is populated
	case shared.MarketingMetadata5TypeStr:
		// marketingMetadata5.Str is populated
	case shared.MarketingMetadata5TypeNumber:
		// marketingMetadata5.Number is populated
	case shared.MarketingMetadata5TypeBoolean:
		// marketingMetadata5.Boolean is populated
}
```
