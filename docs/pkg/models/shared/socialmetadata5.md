# SocialMetadata5


## Supported Types

### SocialMetadata1

```go
socialMetadata5 := shared.CreateSocialMetadata5SocialMetadata1(shared.SocialMetadata1{/* values here */})
```

### 

```go
socialMetadata5 := shared.CreateSocialMetadata5Str(string{/* values here */})
```

### 

```go
socialMetadata5 := shared.CreateSocialMetadata5Number(float64{/* values here */})
```

### 

```go
socialMetadata5 := shared.CreateSocialMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch socialMetadata5.Type {
	case shared.SocialMetadata5TypeSocialMetadata1:
		// socialMetadata5.SocialMetadata1 is populated
	case shared.SocialMetadata5TypeStr:
		// socialMetadata5.Str is populated
	case shared.SocialMetadata5TypeNumber:
		// socialMetadata5.Number is populated
	case shared.SocialMetadata5TypeBoolean:
		// socialMetadata5.Boolean is populated
}
```
