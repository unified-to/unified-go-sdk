# KmsPageMetadata5


## Supported Types

### KmsPageMetadata1

```go
kmsPageMetadata5 := shared.CreateKmsPageMetadata5KmsPageMetadata1(shared.KmsPageMetadata1{/* values here */})
```

### 

```go
kmsPageMetadata5 := shared.CreateKmsPageMetadata5Str(string{/* values here */})
```

### 

```go
kmsPageMetadata5 := shared.CreateKmsPageMetadata5Number(float64{/* values here */})
```

### 

```go
kmsPageMetadata5 := shared.CreateKmsPageMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kmsPageMetadata5.Type {
	case shared.KmsPageMetadata5TypeKmsPageMetadata1:
		// kmsPageMetadata5.KmsPageMetadata1 is populated
	case shared.KmsPageMetadata5TypeStr:
		// kmsPageMetadata5.Str is populated
	case shared.KmsPageMetadata5TypeNumber:
		// kmsPageMetadata5.Number is populated
	case shared.KmsPageMetadata5TypeBoolean:
		// kmsPageMetadata5.Boolean is populated
}
```
