# AdsMetadataSchemas5


## Supported Types

### AdsMetadataSchemas1

```go
adsMetadataSchemas5 := shared.CreateAdsMetadataSchemas5AdsMetadataSchemas1(shared.AdsMetadataSchemas1{/* values here */})
```

### 

```go
adsMetadataSchemas5 := shared.CreateAdsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
adsMetadataSchemas5 := shared.CreateAdsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
adsMetadataSchemas5 := shared.CreateAdsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch adsMetadataSchemas5.Type {
	case shared.AdsMetadataSchemas5TypeAdsMetadataSchemas1:
		// adsMetadataSchemas5.AdsMetadataSchemas1 is populated
	case shared.AdsMetadataSchemas5TypeStr:
		// adsMetadataSchemas5.Str is populated
	case shared.AdsMetadataSchemas5TypeNumber:
		// adsMetadataSchemas5.Number is populated
	case shared.AdsMetadataSchemas5TypeBoolean:
		// adsMetadataSchemas5.Boolean is populated
}
```
