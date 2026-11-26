# DmsMetadata5


## Supported Types

### DmsMetadata1

```go
dmsMetadata5 := shared.CreateDmsMetadata5DmsMetadata1(shared.DmsMetadata1{/* values here */})
```

### 

```go
dmsMetadata5 := shared.CreateDmsMetadata5Str(string{/* values here */})
```

### 

```go
dmsMetadata5 := shared.CreateDmsMetadata5Number(float64{/* values here */})
```

### 

```go
dmsMetadata5 := shared.CreateDmsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dmsMetadata5.Type {
	case shared.DmsMetadata5TypeDmsMetadata1:
		// dmsMetadata5.DmsMetadata1 is populated
	case shared.DmsMetadata5TypeStr:
		// dmsMetadata5.Str is populated
	case shared.DmsMetadata5TypeNumber:
		// dmsMetadata5.Number is populated
	case shared.DmsMetadata5TypeBoolean:
		// dmsMetadata5.Boolean is populated
}
```
