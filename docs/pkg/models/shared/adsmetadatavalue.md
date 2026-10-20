# AdsMetadataValue


## Supported Types

### 

```go
adsMetadataValue := shared.CreateAdsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
adsMetadataValue := shared.CreateAdsMetadataValueStr(string{/* values here */})
```

### 

```go
adsMetadataValue := shared.CreateAdsMetadataValueNumber(float64{/* values here */})
```

### 

```go
adsMetadataValue := shared.CreateAdsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
adsMetadataValue := shared.CreateAdsMetadataValueArrayOfAdsMetadataSchemas5([]shared.AdsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch adsMetadataValue.Type {
	case shared.AdsMetadataValueTypeMapOfAny:
		// adsMetadataValue.MapOfAny is populated
	case shared.AdsMetadataValueTypeStr:
		// adsMetadataValue.Str is populated
	case shared.AdsMetadataValueTypeNumber:
		// adsMetadataValue.Number is populated
	case shared.AdsMetadataValueTypeBoolean:
		// adsMetadataValue.Boolean is populated
	case shared.AdsMetadataValueTypeArrayOfAdsMetadataSchemas5:
		// adsMetadataValue.ArrayOfAdsMetadataSchemas5 is populated
}
```
