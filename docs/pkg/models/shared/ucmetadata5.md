# UcMetadata5


## Supported Types

### UcMetadata1

```go
ucMetadata5 := shared.CreateUcMetadata5UcMetadata1(shared.UcMetadata1{/* values here */})
```

### 

```go
ucMetadata5 := shared.CreateUcMetadata5Str(string{/* values here */})
```

### 

```go
ucMetadata5 := shared.CreateUcMetadata5Number(float64{/* values here */})
```

### 

```go
ucMetadata5 := shared.CreateUcMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ucMetadata5.Type {
	case shared.UcMetadata5TypeUcMetadata1:
		// ucMetadata5.UcMetadata1 is populated
	case shared.UcMetadata5TypeStr:
		// ucMetadata5.Str is populated
	case shared.UcMetadata5TypeNumber:
		// ucMetadata5.Number is populated
	case shared.UcMetadata5TypeBoolean:
		// ucMetadata5.Boolean is populated
}
```
