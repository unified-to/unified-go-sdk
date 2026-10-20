# SocialMetadataExtraData


## Supported Types

### 

```go
socialMetadataExtraData := shared.CreateSocialMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
socialMetadataExtraData := shared.CreateSocialMetadataExtraDataStr(string{/* values here */})
```

### 

```go
socialMetadataExtraData := shared.CreateSocialMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
socialMetadataExtraData := shared.CreateSocialMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
socialMetadataExtraData := shared.CreateSocialMetadataExtraDataArrayOfSocialMetadata5([]shared.SocialMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch socialMetadataExtraData.Type {
	case shared.SocialMetadataExtraDataTypeMapOfAny:
		// socialMetadataExtraData.MapOfAny is populated
	case shared.SocialMetadataExtraDataTypeStr:
		// socialMetadataExtraData.Str is populated
	case shared.SocialMetadataExtraDataTypeNumber:
		// socialMetadataExtraData.Number is populated
	case shared.SocialMetadataExtraDataTypeBoolean:
		// socialMetadataExtraData.Boolean is populated
	case shared.SocialMetadataExtraDataTypeArrayOfSocialMetadata5:
		// socialMetadataExtraData.ArrayOfSocialMetadata5 is populated
}
```
