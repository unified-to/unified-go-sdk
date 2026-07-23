# CdpMetadataSchemas5


## Supported Types

### CdpMetadataSchemas1

```go
cdpMetadataSchemas5 := shared.CreateCdpMetadataSchemas5CdpMetadataSchemas1(shared.CdpMetadataSchemas1{/* values here */})
```

### 

```go
cdpMetadataSchemas5 := shared.CreateCdpMetadataSchemas5Str(string{/* values here */})
```

### 

```go
cdpMetadataSchemas5 := shared.CreateCdpMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
cdpMetadataSchemas5 := shared.CreateCdpMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cdpMetadataSchemas5.Type {
	case shared.CdpMetadataSchemas5TypeCdpMetadataSchemas1:
		// cdpMetadataSchemas5.CdpMetadataSchemas1 is populated
	case shared.CdpMetadataSchemas5TypeStr:
		// cdpMetadataSchemas5.Str is populated
	case shared.CdpMetadataSchemas5TypeNumber:
		// cdpMetadataSchemas5.Number is populated
	case shared.CdpMetadataSchemas5TypeBoolean:
		// cdpMetadataSchemas5.Boolean is populated
}
```
