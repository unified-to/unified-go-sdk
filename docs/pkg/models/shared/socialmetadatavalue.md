# SocialMetadataValue


## Supported Types

### 

```go
socialMetadataValue := shared.CreateSocialMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
socialMetadataValue := shared.CreateSocialMetadataValueStr(string{/* values here */})
```

### 

```go
socialMetadataValue := shared.CreateSocialMetadataValueNumber(float64{/* values here */})
```

### 

```go
socialMetadataValue := shared.CreateSocialMetadataValueBoolean(bool{/* values here */})
```

### 

```go
socialMetadataValue := shared.CreateSocialMetadataValueArrayOfSocialMetadataSchemas5([]shared.SocialMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch socialMetadataValue.Type {
	case shared.SocialMetadataValueTypeMapOfAny:
		// socialMetadataValue.MapOfAny is populated
	case shared.SocialMetadataValueTypeStr:
		// socialMetadataValue.Str is populated
	case shared.SocialMetadataValueTypeNumber:
		// socialMetadataValue.Number is populated
	case shared.SocialMetadataValueTypeBoolean:
		// socialMetadataValue.Boolean is populated
	case shared.SocialMetadataValueTypeArrayOfSocialMetadataSchemas5:
		// socialMetadataValue.ArrayOfSocialMetadataSchemas5 is populated
}
```
