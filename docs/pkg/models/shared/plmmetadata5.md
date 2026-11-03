# PlmMetadata5


## Supported Types

### PlmMetadata1

```go
plmMetadata5 := shared.CreatePlmMetadata5PlmMetadata1(shared.PlmMetadata1{/* values here */})
```

### 

```go
plmMetadata5 := shared.CreatePlmMetadata5Str(string{/* values here */})
```

### 

```go
plmMetadata5 := shared.CreatePlmMetadata5Number(float64{/* values here */})
```

### 

```go
plmMetadata5 := shared.CreatePlmMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch plmMetadata5.Type {
	case shared.PlmMetadata5TypePlmMetadata1:
		// plmMetadata5.PlmMetadata1 is populated
	case shared.PlmMetadata5TypeStr:
		// plmMetadata5.Str is populated
	case shared.PlmMetadata5TypeNumber:
		// plmMetadata5.Number is populated
	case shared.PlmMetadata5TypeBoolean:
		// plmMetadata5.Boolean is populated
}
```
