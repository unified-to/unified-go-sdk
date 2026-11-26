# PrintingMetadataValue

The field's value, typed by `format`.


## Supported Types

### 

```go
printingMetadataValue := shared.CreatePrintingMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
printingMetadataValue := shared.CreatePrintingMetadataValueStr(string{/* values here */})
```

### 

```go
printingMetadataValue := shared.CreatePrintingMetadataValueNumber(float64{/* values here */})
```

### 

```go
printingMetadataValue := shared.CreatePrintingMetadataValueBoolean(bool{/* values here */})
```

### 

```go
printingMetadataValue := shared.CreatePrintingMetadataValueArrayOfPrintingMetadataSchemas5([]shared.PrintingMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch printingMetadataValue.Type {
	case shared.PrintingMetadataValueTypeMapOfAny:
		// printingMetadataValue.MapOfAny is populated
	case shared.PrintingMetadataValueTypeStr:
		// printingMetadataValue.Str is populated
	case shared.PrintingMetadataValueTypeNumber:
		// printingMetadataValue.Number is populated
	case shared.PrintingMetadataValueTypeBoolean:
		// printingMetadataValue.Boolean is populated
	case shared.PrintingMetadataValueTypeArrayOfPrintingMetadataSchemas5:
		// printingMetadataValue.ArrayOfPrintingMetadataSchemas5 is populated
}
```
