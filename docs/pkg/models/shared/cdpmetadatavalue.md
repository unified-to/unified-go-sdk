# CdpMetadataValue


## Supported Types

### 

```go
cdpMetadataValue := shared.CreateCdpMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
cdpMetadataValue := shared.CreateCdpMetadataValueStr(string{/* values here */})
```

### 

```go
cdpMetadataValue := shared.CreateCdpMetadataValueNumber(float64{/* values here */})
```

### 

```go
cdpMetadataValue := shared.CreateCdpMetadataValueBoolean(bool{/* values here */})
```

### 

```go
cdpMetadataValue := shared.CreateCdpMetadataValueArrayOfCdpMetadataSchemas5([]shared.CdpMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cdpMetadataValue.Type {
	case shared.CdpMetadataValueTypeMapOfAny:
		// cdpMetadataValue.MapOfAny is populated
	case shared.CdpMetadataValueTypeStr:
		// cdpMetadataValue.Str is populated
	case shared.CdpMetadataValueTypeNumber:
		// cdpMetadataValue.Number is populated
	case shared.CdpMetadataValueTypeBoolean:
		// cdpMetadataValue.Boolean is populated
	case shared.CdpMetadataValueTypeArrayOfCdpMetadataSchemas5:
		// cdpMetadataValue.ArrayOfCdpMetadataSchemas5 is populated
}
```
