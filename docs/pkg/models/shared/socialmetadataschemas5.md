# SocialMetadataSchemas5


## Supported Types

### SocialMetadataSchemas1

```go
socialMetadataSchemas5 := shared.CreateSocialMetadataSchemas5SocialMetadataSchemas1(shared.SocialMetadataSchemas1{/* values here */})
```

### 

```go
socialMetadataSchemas5 := shared.CreateSocialMetadataSchemas5Str(string{/* values here */})
```

### 

```go
socialMetadataSchemas5 := shared.CreateSocialMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
socialMetadataSchemas5 := shared.CreateSocialMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch socialMetadataSchemas5.Type {
	case shared.SocialMetadataSchemas5TypeSocialMetadataSchemas1:
		// socialMetadataSchemas5.SocialMetadataSchemas1 is populated
	case shared.SocialMetadataSchemas5TypeStr:
		// socialMetadataSchemas5.Str is populated
	case shared.SocialMetadataSchemas5TypeNumber:
		// socialMetadataSchemas5.Number is populated
	case shared.SocialMetadataSchemas5TypeBoolean:
		// socialMetadataSchemas5.Boolean is populated
}
```
