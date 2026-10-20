# AdsMetadata5


## Supported Types

### AdsMetadata1

```go
adsMetadata5 := shared.CreateAdsMetadata5AdsMetadata1(shared.AdsMetadata1{/* values here */})
```

### 

```go
adsMetadata5 := shared.CreateAdsMetadata5Str(string{/* values here */})
```

### 

```go
adsMetadata5 := shared.CreateAdsMetadata5Number(float64{/* values here */})
```

### 

```go
adsMetadata5 := shared.CreateAdsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch adsMetadata5.Type {
	case shared.AdsMetadata5TypeAdsMetadata1:
		// adsMetadata5.AdsMetadata1 is populated
	case shared.AdsMetadata5TypeStr:
		// adsMetadata5.Str is populated
	case shared.AdsMetadata5TypeNumber:
		// adsMetadata5.Number is populated
	case shared.AdsMetadata5TypeBoolean:
		// adsMetadata5.Boolean is populated
}
```
