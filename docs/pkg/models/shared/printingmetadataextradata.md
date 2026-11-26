# PrintingMetadataExtraData


## Supported Types

### 

```go
printingMetadataExtraData := shared.CreatePrintingMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
printingMetadataExtraData := shared.CreatePrintingMetadataExtraDataStr(string{/* values here */})
```

### 

```go
printingMetadataExtraData := shared.CreatePrintingMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
printingMetadataExtraData := shared.CreatePrintingMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
printingMetadataExtraData := shared.CreatePrintingMetadataExtraDataArrayOfPrintingMetadata5([]shared.PrintingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch printingMetadataExtraData.Type {
	case shared.PrintingMetadataExtraDataTypeMapOfAny:
		// printingMetadataExtraData.MapOfAny is populated
	case shared.PrintingMetadataExtraDataTypeStr:
		// printingMetadataExtraData.Str is populated
	case shared.PrintingMetadataExtraDataTypeNumber:
		// printingMetadataExtraData.Number is populated
	case shared.PrintingMetadataExtraDataTypeBoolean:
		// printingMetadataExtraData.Boolean is populated
	case shared.PrintingMetadataExtraDataTypeArrayOfPrintingMetadata5:
		// printingMetadataExtraData.ArrayOfPrintingMetadata5 is populated
}
```
