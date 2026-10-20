# AdsMetadataExtraData


## Supported Types

### 

```go
adsMetadataExtraData := shared.CreateAdsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
adsMetadataExtraData := shared.CreateAdsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
adsMetadataExtraData := shared.CreateAdsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
adsMetadataExtraData := shared.CreateAdsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
adsMetadataExtraData := shared.CreateAdsMetadataExtraDataArrayOfAdsMetadata5([]shared.AdsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch adsMetadataExtraData.Type {
	case shared.AdsMetadataExtraDataTypeMapOfAny:
		// adsMetadataExtraData.MapOfAny is populated
	case shared.AdsMetadataExtraDataTypeStr:
		// adsMetadataExtraData.Str is populated
	case shared.AdsMetadataExtraDataTypeNumber:
		// adsMetadataExtraData.Number is populated
	case shared.AdsMetadataExtraDataTypeBoolean:
		// adsMetadataExtraData.Boolean is populated
	case shared.AdsMetadataExtraDataTypeArrayOfAdsMetadata5:
		// adsMetadataExtraData.ArrayOfAdsMetadata5 is populated
}
```
