# PrintingMetadataSchemas5


## Supported Types

### PrintingMetadataSchemas1

```go
printingMetadataSchemas5 := shared.CreatePrintingMetadataSchemas5PrintingMetadataSchemas1(shared.PrintingMetadataSchemas1{/* values here */})
```

### 

```go
printingMetadataSchemas5 := shared.CreatePrintingMetadataSchemas5Str(string{/* values here */})
```

### 

```go
printingMetadataSchemas5 := shared.CreatePrintingMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
printingMetadataSchemas5 := shared.CreatePrintingMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch printingMetadataSchemas5.Type {
	case shared.PrintingMetadataSchemas5TypePrintingMetadataSchemas1:
		// printingMetadataSchemas5.PrintingMetadataSchemas1 is populated
	case shared.PrintingMetadataSchemas5TypeStr:
		// printingMetadataSchemas5.Str is populated
	case shared.PrintingMetadataSchemas5TypeNumber:
		// printingMetadataSchemas5.Number is populated
	case shared.PrintingMetadataSchemas5TypeBoolean:
		// printingMetadataSchemas5.Boolean is populated
}
```
