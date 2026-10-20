# FormsMetadata5


## Supported Types

### FormsMetadata1

```go
formsMetadata5 := shared.CreateFormsMetadata5FormsMetadata1(shared.FormsMetadata1{/* values here */})
```

### 

```go
formsMetadata5 := shared.CreateFormsMetadata5Str(string{/* values here */})
```

### 

```go
formsMetadata5 := shared.CreateFormsMetadata5Number(float64{/* values here */})
```

### 

```go
formsMetadata5 := shared.CreateFormsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formsMetadata5.Type {
	case shared.FormsMetadata5TypeFormsMetadata1:
		// formsMetadata5.FormsMetadata1 is populated
	case shared.FormsMetadata5TypeStr:
		// formsMetadata5.Str is populated
	case shared.FormsMetadata5TypeNumber:
		// formsMetadata5.Number is populated
	case shared.FormsMetadata5TypeBoolean:
		// formsMetadata5.Boolean is populated
}
```
