# PrintingMetadata5


## Supported Types

### PrintingMetadata1

```go
printingMetadata5 := shared.CreatePrintingMetadata5PrintingMetadata1(shared.PrintingMetadata1{/* values here */})
```

### 

```go
printingMetadata5 := shared.CreatePrintingMetadata5Str(string{/* values here */})
```

### 

```go
printingMetadata5 := shared.CreatePrintingMetadata5Number(float64{/* values here */})
```

### 

```go
printingMetadata5 := shared.CreatePrintingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch printingMetadata5.Type {
	case shared.PrintingMetadata5TypePrintingMetadata1:
		// printingMetadata5.PrintingMetadata1 is populated
	case shared.PrintingMetadata5TypeStr:
		// printingMetadata5.Str is populated
	case shared.PrintingMetadata5TypeNumber:
		// printingMetadata5.Number is populated
	case shared.PrintingMetadata5TypeBoolean:
		// printingMetadata5.Boolean is populated
}
```
