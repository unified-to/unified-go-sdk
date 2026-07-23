# CdpMetadataExtraData


## Supported Types

### 

```go
cdpMetadataExtraData := shared.CreateCdpMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
cdpMetadataExtraData := shared.CreateCdpMetadataExtraDataStr(string{/* values here */})
```

### 

```go
cdpMetadataExtraData := shared.CreateCdpMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
cdpMetadataExtraData := shared.CreateCdpMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
cdpMetadataExtraData := shared.CreateCdpMetadataExtraDataArrayOfCdpMetadata5([]shared.CdpMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cdpMetadataExtraData.Type {
	case shared.CdpMetadataExtraDataTypeMapOfAny:
		// cdpMetadataExtraData.MapOfAny is populated
	case shared.CdpMetadataExtraDataTypeStr:
		// cdpMetadataExtraData.Str is populated
	case shared.CdpMetadataExtraDataTypeNumber:
		// cdpMetadataExtraData.Number is populated
	case shared.CdpMetadataExtraDataTypeBoolean:
		// cdpMetadataExtraData.Boolean is populated
	case shared.CdpMetadataExtraDataTypeArrayOfCdpMetadata5:
		// cdpMetadataExtraData.ArrayOfCdpMetadata5 is populated
}
```
