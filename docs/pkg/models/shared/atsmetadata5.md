# AtsMetadata5


## Supported Types

### AtsMetadata1

```go
atsMetadata5 := shared.CreateAtsMetadata5AtsMetadata1(shared.AtsMetadata1{/* values here */})
```

### 

```go
atsMetadata5 := shared.CreateAtsMetadata5Str(string{/* values here */})
```

### 

```go
atsMetadata5 := shared.CreateAtsMetadata5Number(float64{/* values here */})
```

### 

```go
atsMetadata5 := shared.CreateAtsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch atsMetadata5.Type {
	case shared.AtsMetadata5TypeAtsMetadata1:
		// atsMetadata5.AtsMetadata1 is populated
	case shared.AtsMetadata5TypeStr:
		// atsMetadata5.Str is populated
	case shared.AtsMetadata5TypeNumber:
		// atsMetadata5.Number is populated
	case shared.AtsMetadata5TypeBoolean:
		// atsMetadata5.Boolean is populated
}
```
